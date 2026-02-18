package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestResolveMemoryRuntime_BinaryPreferred(t *testing.T) {
	t.Setenv("OPENKIT_MEMORY_RUNTIME", "")
	t.Setenv("OPENKIT_MEMORY_RUNTIME_PATH", "")
	withLookPathStub(t, func(name string) (string, error) {
		if name == "openkit-rs" {
			return "/usr/local/bin/openkit-rs", nil
		}
		return "", fmt.Errorf("not found")
	})
	withExecutablePathStub(t, func() (string, error) { return "/tmp/openkit", nil })

	bin, args, err := resolveMemoryRuntime(t.TempDir())
	if err != nil {
		t.Fatalf("resolveMemoryRuntime returned error: %v", err)
	}
	if bin != "openkit-rs" {
		t.Fatalf("expected openkit-rs binary, got %s", bin)
	}
	if len(args) != 0 {
		t.Fatalf("expected no args for binary runtime, got %v", args)
	}
}

func TestResolveMemoryRuntime_CargoFallback(t *testing.T) {
	t.Setenv("OPENKIT_MEMORY_RUNTIME", "")
	t.Setenv("OPENKIT_MEMORY_RUNTIME_PATH", "")
	projectDir := t.TempDir()
	mustWriteMemoryTestFile(t, filepath.Join(projectDir, "rust-cli", "Cargo.toml"), "[package]\nname='x'\n")

	withLookPathStub(t, func(name string) (string, error) {
		if name == "cargo" {
			return "/usr/local/bin/cargo", nil
		}
		return "", fmt.Errorf("not found")
	})
	withExecutablePathStub(t, func() (string, error) { return "/tmp/openkit", nil })

	bin, args, err := resolveMemoryRuntime(projectDir)
	if err != nil {
		t.Fatalf("resolveMemoryRuntime returned error: %v", err)
	}
	if bin != "cargo" {
		t.Fatalf("expected cargo runtime, got %s", bin)
	}
	if len(args) == 0 || args[0] != "run" {
		t.Fatalf("expected cargo run args, got %v", args)
	}
}

func TestResolveMemoryRuntime_InvalidMode(t *testing.T) {
	t.Setenv("OPENKIT_MEMORY_RUNTIME", "invalid")
	t.Setenv("OPENKIT_MEMORY_RUNTIME_PATH", "")
	withLookPathStub(t, func(name string) (string, error) {
		return "", fmt.Errorf("not found")
	})
	withExecutablePathStub(t, func() (string, error) { return "/tmp/openkit", nil })

	_, _, err := resolveMemoryRuntime(t.TempDir())
	if err == nil {
		t.Fatalf("expected error for invalid runtime mode")
	}
}

func TestResolveCargoRuntime_MissingManifest(t *testing.T) {
	projectDir := t.TempDir()
	withLookPathStub(t, func(name string) (string, error) {
		if name == "cargo" {
			return "/usr/local/bin/cargo", nil
		}
		return "", fmt.Errorf("not found")
	})

	_, _, err := resolveCargoRuntime(projectDir)
	if err == nil {
		t.Fatalf("expected error when rust-cli Cargo.toml is missing")
	}
}

func TestResolveMemoryRuntime_ExplicitPath(t *testing.T) {
	projectDir := t.TempDir()
	explicit := filepath.Join(projectDir, "openkit-rs")
	mustWriteMemoryTestFile(t, explicit, "binary")
	t.Setenv("OPENKIT_MEMORY_RUNTIME_PATH", explicit)
	t.Setenv("OPENKIT_MEMORY_RUNTIME", "")

	bin, args, err := resolveMemoryRuntime(projectDir)
	if err != nil {
		t.Fatalf("resolveMemoryRuntime returned error: %v", err)
	}
	if bin != explicit {
		t.Fatalf("expected explicit runtime path, got %s", bin)
	}
	if len(args) != 0 {
		t.Fatalf("expected no args for explicit runtime, got %v", args)
	}
}

func withLookPathStub(t *testing.T, stub func(name string) (string, error)) {
	t.Helper()
	original := lookPath
	lookPath = stub
	t.Cleanup(func() { lookPath = original })
}

func withExecutablePathStub(t *testing.T, stub func() (string, error)) {
	t.Helper()
	original := executablePath
	executablePath = stub
	t.Cleanup(func() { executablePath = original })
}

func mustWriteMemoryTestFile(t *testing.T, path, content string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatalf("mkdir failed: %v", err)
	}
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("write failed: %v", err)
	}
}
