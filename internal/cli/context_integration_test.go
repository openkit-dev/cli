package cli

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRunContext_DetectsCLIAndGeneratesDocs(t *testing.T) {
	projectDir := t.TempDir()
	mustWriteFile(t, filepath.Join(projectDir, "go.mod"), "module example.com/cli\n")
	mustWriteFile(t, filepath.Join(projectDir, "cmd", "app", "main.go"), "package main\nfunc main() {}\n")

	originalWD, err := os.Getwd()
	if err != nil {
		t.Fatalf("getwd failed: %v", err)
	}
	t.Cleanup(func() { _ = os.Chdir(originalWD) })
	if err := os.Chdir(projectDir); err != nil {
		t.Fatalf("chdir failed: %v", err)
	}

	resetContextFlags()
	contextAutoYes = true

	runContext()

	for _, f := range []string{"CONTEXT.md", "CLI_ARCHITECTURE.md", "SECURITY.md", "QUALITY_GATES.md", "ACTION_ITEMS.md"} {
		if _, err := os.Stat(filepath.Join(projectDir, "docs", f)); err != nil {
			t.Fatalf("expected %s to be generated: %v", f, err)
		}
	}
}

func TestRunContext_ForcedTypeWebFullstack(t *testing.T) {
	projectDir := t.TempDir()
	originalWD, err := os.Getwd()
	if err != nil {
		t.Fatalf("getwd failed: %v", err)
	}
	t.Cleanup(func() { _ = os.Chdir(originalWD) })
	if err := os.Chdir(projectDir); err != nil {
		t.Fatalf("chdir failed: %v", err)
	}

	resetContextFlags()
	contextAutoYes = true
	contextType = "web-fullstack"

	runContext()

	for _, f := range []string{"CONTEXT.md", "BACKEND.md", "FRONTEND.md"} {
		if _, err := os.Stat(filepath.Join(projectDir, "docs", f)); err != nil {
			t.Fatalf("expected %s to be generated: %v", f, err)
		}
	}
}

func TestRunContext_WithOverlaySelection(t *testing.T) {
	projectDir := t.TempDir()
	mustWriteFile(t, filepath.Join(projectDir, "go.mod"), "module example.com/cli\n")
	mustWriteFile(t, filepath.Join(projectDir, "cmd", "main.go"), "package main\nfunc main() {}\n")

	originalWD, err := os.Getwd()
	if err != nil {
		t.Fatalf("getwd failed: %v", err)
	}
	t.Cleanup(func() { _ = os.Chdir(originalWD) })
	if err := os.Chdir(projectDir); err != nil {
		t.Fatalf("chdir failed: %v", err)
	}

	resetContextFlags()
	contextAutoYes = true
	contextOverlays = []string{"testing-overlay"}

	runContext()

	if _, err := os.Stat(filepath.Join(projectDir, "docs", "TESTING.md")); err != nil {
		t.Fatalf("expected TESTING.md to be generated from overlay: %v", err)
	}
}

func resetContextFlags() {
	contextAutoYes = false
	contextType = ""
	contextOverlays = nil
}

func mustWriteFile(t *testing.T, path, content string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatalf("mkdir failed: %v", err)
	}
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("write failed: %v", err)
	}
}
