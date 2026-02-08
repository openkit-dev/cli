package selfupdate

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

type UpgradeOptions struct {
	RepoOwner  string
	RepoName   string
	Tag        string
	InstallDir string
}

func Upgrade(ctx context.Context, client *http.Client, opt UpgradeOptions) (installedPath string, err error) {
	if opt.RepoOwner == "" || opt.RepoName == "" {
		return "", fmt.Errorf("missing repo owner/name")
	}
	if opt.Tag == "" {
		return "", fmt.Errorf("missing tag")
	}
	if opt.InstallDir == "" {
		return "", fmt.Errorf("missing install dir")
	}

	artifact, err := artifactFilename(runtime.GOOS, runtime.GOARCH)
	if err != nil {
		return "", err
	}
	base := fmt.Sprintf("https://github.com/%s/%s/releases/download/%s", opt.RepoOwner, opt.RepoName, opt.Tag)
	artifactURL := base + "/" + artifact
	checksumsURL := base + "/checksums.txt"

	tmpDir, err := os.MkdirTemp("", "openkit-upgrade-*")
	if err != nil {
		return "", err
	}
	defer os.RemoveAll(tmpDir)

	artifactPath := filepath.Join(tmpDir, artifact)
	if err := downloadToFile(ctx, client, artifactURL, artifactPath); err != nil {
		return "", err
	}

	checksumsPath := filepath.Join(tmpDir, "checksums.txt")
	if err := downloadToFile(ctx, client, checksumsURL, checksumsPath); err != nil {
		return "", err
	}

	expected, err := checksumForFile(checksumsPath, artifact)
	if err != nil {
		return "", err
	}
	actual, err := sha256File(artifactPath)
	if err != nil {
		return "", err
	}
	if !strings.EqualFold(expected, actual) {
		return "", fmt.Errorf("checksum mismatch for %s: expected %s got %s", artifact, expected, actual)
	}

	binaryName := "openkit"
	if runtime.GOOS == "windows" {
		binaryName = "openkit.exe"
	}

	extracted, err := extractBinary(tmpDir, artifactPath, binaryName)
	if err != nil {
		return "", err
	}

	if err := os.MkdirAll(opt.InstallDir, 0755); err != nil {
		return "", err
	}
	installedPath = filepath.Join(opt.InstallDir, binaryName)
	if err := atomicReplaceFile(extracted, installedPath, 0755); err != nil {
		return "", err
	}

	return installedPath, nil
}

func artifactFilename(goos, goarch string) (string, error) {
	osName, archName, ext, err := platformArtifact(goos, goarch)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("openkit_%s_%s.%s", osName, archName, ext), nil
}

func ArtifactFilename(goos, goarch string) (string, error) {
	return artifactFilename(goos, goarch)
}

func platformArtifact(goos, goarch string) (osName string, archName string, ext string, err error) {
	switch goos {
	case "darwin":
		osName = "Darwin"
	case "linux":
		osName = "Linux"
	case "windows":
		osName = "Windows"
	default:
		return "", "", "", fmt.Errorf("unsupported OS: %s", goos)
	}

	switch goarch {
	case "amd64":
		archName = "x86_64"
	case "arm64":
		archName = "arm64"
	default:
		return "", "", "", fmt.Errorf("unsupported arch: %s", goarch)
	}

	if goos == "windows" {
		ext = "zip"
		return osName, archName, ext, nil
	}
	ext = "tar.gz"
	return osName, archName, ext, nil
}

func downloadToFile(ctx context.Context, client *http.Client, url string, path string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		b, _ := io.ReadAll(io.LimitReader(resp.Body, 8<<10))
		return fmt.Errorf("download failed: %s: %s", resp.Status, string(b))
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	return err
}

func checksumForFile(checksumsPath string, filename string) (string, error) {
	b, err := os.ReadFile(checksumsPath)
	if err != nil {
		return "", err
	}
	lines := strings.Split(string(b), "\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) != 2 {
			continue
		}
		if fields[1] == filename {
			return fields[0], nil
		}
	}
	return "", fmt.Errorf("checksum not found for %s", filename)
}

func sha256File(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

func extractBinary(tmpDir, archivePath, binaryName string) (string, error) {
	if strings.HasSuffix(archivePath, ".zip") {
		return extractZip(tmpDir, archivePath, binaryName)
	}
	if strings.HasSuffix(archivePath, ".tar.gz") {
		return extractTarGz(tmpDir, archivePath, binaryName)
	}
	return "", fmt.Errorf("unsupported archive: %s", archivePath)
}

func extractZip(tmpDir, zipPath, binaryName string) (string, error) {
	r, err := zip.OpenReader(zipPath)
	if err != nil {
		return "", err
	}
	defer r.Close()

	for _, f := range r.File {
		if filepath.Base(f.Name) != binaryName {
			continue
		}
		rc, err := f.Open()
		if err != nil {
			return "", err
		}
		defer rc.Close()

		out := filepath.Join(tmpDir, binaryName)
		wf, err := os.Create(out)
		if err != nil {
			return "", err
		}
		if _, err := io.Copy(wf, rc); err != nil {
			wf.Close()
			return "", err
		}
		if err := wf.Close(); err != nil {
			return "", err
		}
		return out, nil
	}
	return "", fmt.Errorf("binary %s not found in zip", binaryName)
}

func extractTarGz(tmpDir, tarGzPath, binaryName string) (string, error) {
	f, err := os.Open(tarGzPath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	gzr, err := gzip.NewReader(f)
	if err != nil {
		return "", err
	}
	defer gzr.Close()

	tr := tar.NewReader(gzr)
	for {
		hdr, err := tr.Next()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return "", err
		}
		if hdr.Typeflag != tar.TypeReg {
			continue
		}
		if filepath.Base(hdr.Name) != binaryName {
			continue
		}

		out := filepath.Join(tmpDir, binaryName)
		wf, err := os.Create(out)
		if err != nil {
			return "", err
		}
		if _, err := io.Copy(wf, tr); err != nil {
			wf.Close()
			return "", err
		}
		if err := wf.Close(); err != nil {
			return "", err
		}
		return out, nil
	}
	return "", fmt.Errorf("binary %s not found in tar.gz", binaryName)
}

func atomicReplaceFile(src string, dst string, mode os.FileMode) error {
	dir := filepath.Dir(dst)
	tmp := filepath.Join(dir, fmt.Sprintf(".%s.tmp.%d", filepath.Base(dst), time.Now().UnixNano()))

	if err := copyFile(src, tmp, mode); err != nil {
		return err
	}

	if err := os.Rename(tmp, dst); err != nil {
		_ = os.Remove(tmp)
		return err
	}
	return nil
}

func copyFile(src, dst string, mode os.FileMode) error {
	r, err := os.Open(src)
	if err != nil {
		return err
	}
	defer r.Close()

	w, err := os.OpenFile(dst, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, mode)
	if err != nil {
		return err
	}
	defer w.Close()

	_, err = io.Copy(w, r)
	return err
}
