package cli

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

var lookPath = exec.LookPath
var executablePath = os.Executable

var memoryRuntimeCmd = &cobra.Command{
	Use:   "memory",
	Short: "Manage docs-first memory commands",
	Long:  "Run Memory Kernel commands using the Rust runtime (openkit-rs).",
}

var memoryInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize Memory Kernel directories and contracts",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runMemoryRuntime("init", args)
	},
}

var memoryDoctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Validate memory health and documentation graph",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runMemoryRuntime("doctor", args)
	},
}

var memoryCaptureCmd = &cobra.Command{
	Use:   "capture",
	Short: "Capture a memory session snapshot",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runMemoryRuntime("capture", args)
	},
}

var memoryReviewCmd = &cobra.Command{
	Use:   "review",
	Short: "Review observations, tensions, and session state",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runMemoryRuntime("review", args)
	},
}

func init() {
	for _, c := range []*cobra.Command{memoryInitCmd, memoryDoctorCmd, memoryCaptureCmd, memoryReviewCmd} {
		c.Args = cobra.ArbitraryArgs
		c.DisableFlagParsing = true
	}
	memoryRuntimeCmd.AddCommand(memoryInitCmd, memoryDoctorCmd, memoryCaptureCmd, memoryReviewCmd)
	rootCmd.AddCommand(memoryRuntimeCmd)
}

func runMemoryRuntime(subcommand string, args []string) error {
	projectDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err)
	}

	bin, binArgs, err := resolveMemoryRuntime(projectDir)
	if err != nil {
		return err
	}

	fullArgs := append(binArgs, "memory", subcommand)
	fullArgs = append(fullArgs, args...)

	cmd := exec.Command(bin, fullArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Dir = projectDir

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("memory runtime failed: %w", err)
	}
	return nil
}

func resolveMemoryRuntime(projectDir string) (string, []string, error) {
	if explicit := strings.TrimSpace(os.Getenv("OPENKIT_MEMORY_RUNTIME_PATH")); explicit != "" {
		if _, err := os.Stat(explicit); err != nil {
			return "", nil, fmt.Errorf("OPENKIT_MEMORY_RUNTIME_PATH does not exist: %s", explicit)
		}
		return explicit, nil, nil
	}

	mode := strings.ToLower(strings.TrimSpace(os.Getenv("OPENKIT_MEMORY_RUNTIME")))
	switch mode {
	case "cargo":
		return resolveCargoRuntime(projectDir)
	case "binary", "":
		if sidecar, ok := resolveSidecarRuntime(); ok {
			return sidecar, nil, nil
		}
		if _, err := lookPath("openkit-rs"); err == nil {
			return "openkit-rs", nil, nil
		}
		if mode == "binary" {
			return "", nil, fmt.Errorf("openkit-rs binary not found in PATH")
		}
		return resolveCargoRuntime(projectDir)
	default:
		return "", nil, fmt.Errorf("invalid OPENKIT_MEMORY_RUNTIME value %q (use: binary or cargo)", mode)
	}
}

func resolveSidecarRuntime() (string, bool) {
	exe, err := executablePath()
	if err != nil {
		return "", false
	}
	dir := filepath.Dir(exe)
	name := "openkit-rs"
	if runtime.GOOS == "windows" {
		name += ".exe"
	}
	candidate := filepath.Join(dir, name)
	if _, err := os.Stat(candidate); err == nil {
		return candidate, true
	}
	return "", false
}

func resolveCargoRuntime(projectDir string) (string, []string, error) {
	if _, err := lookPath("cargo"); err != nil {
		return "", nil, fmt.Errorf("memory runtime unavailable: neither openkit-rs nor cargo is installed")
	}

	manifest := filepath.Join(projectDir, "rust-cli", "Cargo.toml")
	if _, err := os.Stat(manifest); err != nil {
		return "", nil, fmt.Errorf("memory runtime unavailable: rust manifest not found at %s", manifest)
	}

	args := []string{"run", "--manifest-path", manifest, "--"}
	return "cargo", args, nil
}
