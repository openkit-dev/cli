package detection

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDetectCLITool(t *testing.T) {
	dir := t.TempDir()
	mustWriteFile(t, filepath.Join(dir, "go.mod"), "module example.com/cli\n")
	mustWriteFile(t, filepath.Join(dir, "cmd", "app", "main.go"), "package main\nfunc main() {}\n")

	reg, err := LoadRegistry(dir)
	if err != nil {
		t.Fatalf("LoadRegistry failed: %v", err)
	}

	res, err := reg.Detect()
	if err != nil {
		t.Fatalf("Detect failed: %v", err)
	}

	if res.ProjectType != "cli-tool" {
		t.Fatalf("expected project type cli-tool, got %q", res.ProjectType)
	}
}

func TestDetectWebFullstack(t *testing.T) {
	dir := t.TempDir()
	mustMkdir(t, filepath.Join(dir, "frontend"))
	mustMkdir(t, filepath.Join(dir, "backend"))

	reg, err := LoadRegistry(dir)
	if err != nil {
		t.Fatalf("LoadRegistry failed: %v", err)
	}

	res, err := reg.Detect()
	if err != nil {
		t.Fatalf("Detect failed: %v", err)
	}

	if res.ProjectType != "web-fullstack" {
		t.Fatalf("expected project type web-fullstack, got %q", res.ProjectType)
	}
}

func TestDetectUnknownForEmptyProject(t *testing.T) {
	dir := t.TempDir()

	reg, err := LoadRegistry(dir)
	if err != nil {
		t.Fatalf("LoadRegistry failed: %v", err)
	}

	res, err := reg.Detect()
	if err != nil {
		t.Fatalf("Detect failed: %v", err)
	}

	if res.ProjectType != "unknown" {
		t.Fatalf("expected project type unknown, got %q", res.ProjectType)
	}
}

func TestHasMatchingContent(t *testing.T) {
	dir := t.TempDir()
	mustWriteFile(t, filepath.Join(dir, "README.md"), "this project uses special-pattern-123")

	reg := &Registry{projectDir: dir}
	if !reg.hasMatchingContent("special-pattern-123") {
		t.Fatal("expected hasMatchingContent to return true")
	}
	if reg.hasMatchingContent("non-existent-pattern") {
		t.Fatal("expected hasMatchingContent to return false")
	}
}

func TestOverlayDefaultsLoaded(t *testing.T) {
	dir := t.TempDir()
	reg, err := LoadRegistry(dir)
	if err != nil {
		t.Fatalf("LoadRegistry failed: %v", err)
	}

	for _, id := range []string{"testing-overlay", "security-overlay", "ci-cd-overlay", "documentation-overlay"} {
		if _, ok := reg.GetOverlayConfig(id); !ok {
			t.Fatalf("expected default overlay %q to be loaded", id)
		}
	}
}

func mustWriteFile(t *testing.T, path, content string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatalf("mkdir failed: %v", err)
	}
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("write file failed: %v", err)
	}
}

func mustMkdir(t *testing.T, path string) {
	t.Helper()
	if err := os.MkdirAll(path, 0o755); err != nil {
		t.Fatalf("mkdir failed: %v", err)
	}
}
