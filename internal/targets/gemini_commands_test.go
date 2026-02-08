package targets

import (
	"strings"
	"testing"
	"testing/fstest"
)

func TestParseMarkdownFrontmatter_ParsesDescription(t *testing.T) {
	md := "---\n" +
		"description: Hello world\n" +
		"subtask: false\n" +
		"---\n\n" +
		"# Title\n\nBody\n"

	desc, body := parseMarkdownFrontmatter([]byte(md))
	if desc != "Hello world" {
		t.Fatalf("desc = %q, want %q", desc, "Hello world")
	}
	if !strings.HasPrefix(string(body), "\n# Title") {
		t.Fatalf("body did not start with expected content: %q", string(body)[:min(20, len(body))])
	}
}

func TestGeminiCommandToml_UsesMultilinePrompt(t *testing.T) {
	b, err := geminiCommandToml("My desc", "line1\nline2\n")
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	s := string(b)
	if !strings.Contains(s, "description = ") {
		t.Fatalf("missing description: %q", s)
	}
	if !strings.Contains(s, "prompt = \"\"\"\n") {
		t.Fatalf("expected multiline prompt: %q", s)
	}
}

func TestGeminiCommandToml_FallsBackToBasicString(t *testing.T) {
	prompt := "has triple quotes: \"\"\" inside"
	b, err := geminiCommandToml("", prompt)
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	s := string(b)
	if strings.Contains(s, "\"\"\"\n") {
		t.Fatalf("expected basic string fallback, got multiline: %q", s)
	}
	if !strings.Contains(s, "prompt = \"") {
		t.Fatalf("expected basic prompt string: %q", s)
	}
}

func TestDesiredGeminiCommands_GeneratesTomlAndReplacesArgs(t *testing.T) {
	fsys := fstest.MapFS{
		"base/commands/specify.md": {Data: []byte("---\ndescription: Spec\n---\n\nHello $ARGUMENTS\n")},
		"base/commands/README.md":  {Data: []byte("ignore")},
	}

	files, err := desiredGeminiCommands(fsys)
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	if len(files) != 1 {
		t.Fatalf("files = %d, want 1", len(files))
	}

	f := files[0]
	if f.OutputPath != ".gemini/commands/openkit/specify.toml" {
		t.Fatalf("OutputPath = %q", f.OutputPath)
	}
	s := string(f.Bytes)
	if !strings.Contains(s, "description = ") {
		t.Fatalf("missing description in toml: %q", s)
	}
	if !strings.Contains(s, "Hello {{args}}") {
		t.Fatalf("did not replace $ARGUMENTS: %q", s)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
