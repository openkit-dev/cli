package targets

import (
	"strings"
	"testing"
)

func TestGenerateCursorRules_ContainsExpectedSections(t *testing.T) {
	content := generateCursorRules()
	text := string(content)

	expectedSections := []string{
		"# Cursor Rules - OpenKit",
		"## Core Principles",
		"## Project Structure",
		"## Commands",
		"## Skills",
	}

	for _, section := range expectedSections {
		if !strings.Contains(text, section) {
			t.Errorf("Expected section %q not found in generated .cursorrules", section)
		}
	}
}

func TestGenerateCursorRules_ContainsSDDWorkflow(t *testing.T) {
	content := generateCursorRules()
	text := string(content)

	if !strings.Contains(text, "/specify") {
		t.Error("Expected SDD workflow to mention /specify")
	}
	if !strings.Contains(text, "/clarify") {
		t.Error("Expected SDD workflow to mention /clarify")
	}
	if !strings.Contains(text, "/plan") {
		t.Error("Expected SDD workflow to mention /plan")
	}
}

func TestGenerateCursorMDC_HasFrontmatter(t *testing.T) {
	content := generateCursorMDC()
	text := string(content)

	if !strings.HasPrefix(text, "---") {
		t.Error("Expected .mdc file to start with frontmatter (---)")
	}

	if !strings.Contains(text, "name:") {
		t.Error("Expected frontmatter to contain 'name:' field")
	}
	if !strings.Contains(text, "description:") {
		t.Error("Expected frontmatter to contain 'description:' field")
	}
}

func TestGenerateCursorMDC_ContainsWorkflowSections(t *testing.T) {
	content := generateCursorMDC()
	text := string(content)

	expectedSections := []string{
		"## Workflow Overview",
		"## Before Writing Code",
		"## During Implementation",
		"## Code Quality Standards",
		"## Verification Scripts",
		"## Skills Reference",
	}

	for _, section := range expectedSections {
		if !strings.Contains(text, section) {
			t.Errorf("Expected section %q not found in .mdc file", section)
		}
	}
}

func TestGenerateCursorMDC_ReferencesSkillsPath(t *testing.T) {
	content := generateCursorMDC()
	text := string(content)

	if !strings.Contains(text, ".cursor/skills/") {
		t.Error("Expected .mdc to reference .cursor/skills/ path")
	}
}

func TestGenerateCursorRules_NotEmpty(t *testing.T) {
	content := generateCursorRules()
	if len(content) == 0 {
		t.Error("Expected generateCursorRules to return non-empty content")
	}
	if len(content) < 100 {
		t.Errorf("Expected generateCursorRules to return substantial content, got %d bytes", len(content))
	}
}

func TestGenerateCursorMDC_NotEmpty(t *testing.T) {
	content := generateCursorMDC()
	if len(content) == 0 {
		t.Error("Expected generateCursorMDC to return non-empty content")
	}
	if len(content) < 500 {
		t.Errorf("Expected generateCursorMDC to return substantial content, got %d bytes", len(content))
	}
}
