package cli

import (
	"context"
	"fmt"
	"net/http"
	"runtime"
	"time"

	"github.com/openkit-devtools/openkit/internal/platform"
	"github.com/openkit-devtools/openkit/internal/selfupdate"
	"github.com/spf13/cobra"
)

var (
	flagUpgradeCheck  bool
	flagUpgradeDryRun bool
)

var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Download and install the latest OpenKit CLI release",
	Long:  "Downloads the latest release from GitHub, verifies checksums, and installs into your OpenKit home directory.",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runUpgrade(cmd.Context(), upgradeRunOptions{CheckOnly: flagUpgradeCheck || flagUpgradeDryRun})
	},
}

func init() {
	upgradeCmd.Flags().BoolVar(&flagUpgradeCheck, "check", false, "Check for updates without downloading/installing")
	upgradeCmd.Flags().BoolVar(&flagUpgradeDryRun, "dry-run", false, "Alias for --check")
}

type upgradeRunOptions struct {
	CheckOnly bool
}

func runUpgrade(ctx context.Context, opt upgradeRunOptions) error {
	installDir, err := platform.OpenKitBinDir()
	if err != nil {
		exitWithError(fmt.Sprintf("Failed to determine install dir: %v", err))
	}

	client := &http.Client{Timeout: 60 * time.Second}
	latestURL := "https://api.github.com/repos/openkit-devtools/openkit/releases/latest"
	tag, _, _, err := selfupdate.FetchLatestTag(ctx, client, latestURL, "")
	if err != nil {
		exitWithError(fmt.Sprintf("Failed to fetch latest release: %v", err))
	}

	if opt.CheckOnly {
		artifact, err := selfupdate.ArtifactFilename(runtime.GOOS, runtime.GOARCH)
		if err != nil {
			exitWithError(fmt.Sprintf("Failed to compute artifact filename: %v", err))
		}
		base := fmt.Sprintf("https://github.com/%s/%s/releases/download/%s", "openkit-devtools", "openkit", tag)
		artifactURL := base + "/" + artifact
		checksumsURL := base + "/checksums.txt"

		printInfo(fmt.Sprintf("Current: %s", GetVersion()))
		printInfo(fmt.Sprintf("Latest:  %s", tag))
		printInfo(fmt.Sprintf("Artifact: %s", artifact))
		printInfo(fmt.Sprintf("URL:      %s", artifactURL))

		if err := probeURL(ctx, client, artifactURL); err != nil {
			exitWithError(fmt.Sprintf("Artifact check failed: %v", err))
		}
		if err := probeURL(ctx, client, checksumsURL); err != nil {
			exitWithError(fmt.Sprintf("Checksums check failed: %v", err))
		}

		if GetVersion() == tag {
			printSuccess("Already up to date")
		} else {
			printSuccess("Update available")
			printInfo("Run: openkit upgrade")
		}

		statePath, err := platform.OpenKitStatePath()
		if err == nil {
			_ = selfupdate.SaveState(statePath, selfupdate.State{LastCheckedUnix: time.Now().Unix(), LatestTag: tag})
		}
		return nil
	}

	printInfo(fmt.Sprintf("Upgrading to %s...", tag))
	path, err := selfupdate.Upgrade(ctx, client, selfupdate.UpgradeOptions{
		RepoOwner:  "openkit-devtools",
		RepoName:   "openkit",
		Tag:        tag,
		InstallDir: installDir,
	})
	if err != nil {
		exitWithError(fmt.Sprintf("Upgrade failed: %v", err))
	}

	printSuccess(fmt.Sprintf("Installed to %s", path))
	printInfo("Restart your terminal session if needed")

	// Best-effort: update state so next runs don't recheck immediately.
	statePath, err := platform.OpenKitStatePath()
	if err == nil {
		_ = selfupdate.SaveState(statePath, selfupdate.State{LastCheckedUnix: time.Now().Unix(), LatestTag: tag})
	}

	return nil
}

func probeURL(ctx context.Context, client *http.Client, url string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodHead, url, nil)
	if err != nil {
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 400 {
		return fmt.Errorf("%s returned %s", url, resp.Status)
	}
	return nil
}
