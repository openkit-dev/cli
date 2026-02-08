package targets

import (
	"strings"
	"testing"
)

func TestGenerateAgentsMD_ContainsExpectedSections(t *testing.T) {
	content := generateAgentsMD()
	text := string(content)

	expectedSections := []string{
		"# OpenKit - Agents Configuration",
		"## What is OpenKit?",
		"## Project Structure",
		"## SDD Workflow",
		"## Skills Reference",
		"## Command Policies",
		"## Working Agreements",
		"## Verification Scripts",
	}

	for _, section := range expectedSections {
		if !strings.Contains(text, section) {
			t.Errorf("Expected section %q not found in AGENTS.md", section)
		}
	}
}

func TestGenerateAgentsMD_ContainsSDDSteps(t *testing.T) {
	content := generateAgentsMD()
	text := string(content)

	expectedSteps := []string{
		"### 1. Specify",
		"### 2. Clarify",
		"### 3. Plan",
		"### 4. Tasks",
		"### 5. Implement",
		"### 6. Test",
	}

	for _, step := range expectedSteps {
		if !strings.Contains(text, step) {
			t.Errorf("Expected SDD step %q not found in AGENTS.md", step)
		}
	}
}

func TestGenerateAgentsMD_ReferencesSkillsPath(t *testing.T) {
	content := generateAgentsMD()
	text := string(content)

	if !strings.Contains(text, ".agents/skills/") {
		t.Error("Expected AGENTS.md to reference .agents/skills/ path")
	}
}

func TestGenerateAgentsMD_SizeUnder32KB(t *testing.T) {
	content := generateAgentsMD()
	// Codex default limit is 32KB (32768 bytes)
	if len(content) > 32768 {
		t.Errorf("AGENTS.md size (%d bytes) exceeds Codex's 32KB limit", len(content))
	}
}

func TestGenerateCodexRules_ValidStarlarkSyntax(t *testing.T) {
	content := generateCodexRules()
	text := string(content)

	// Check for basic Starlark structure
	if !strings.Contains(text, "prefix_rule(") {
		t.Error("Expected rules file to contain prefix_rule() calls")
	}

	// Check for required fields
	requiredFields := []string{
		"pattern =",
		"decision =",
		"justification =",
	}

	for _, field := range requiredFields {
		if !strings.Contains(text, field) {
			t.Errorf("Expected rules file to contain field %q", field)
		}
	}
}

func TestGenerateCodexRules_ContainsCommonCommands(t *testing.T) {
	content := generateCodexRules()
	text := string(content)

	expectedCommands := []string{
		"git\", \"status",
		"git\", \"log",
		"git\", \"diff",
		"npm\", \"test",
		"npm\", \"run\", \"lint",
	}

	for _, cmd := range expectedCommands {
		if !strings.Contains(text, cmd) {
			t.Errorf("Expected rules to include command pattern %q", cmd)
		}
	}
}

func TestGenerateCodexRules_ContainsForbiddenCommands(t *testing.T) {
	content := generateCodexRules()
	text := string(content)

	// Check for forbidden rules
	if !strings.Contains(text, "decision = \"forbidden\"") {
		t.Error("Expected rules to contain at least one forbidden decision")
	}

	// Check for specific dangerous patterns
	dangerousPatterns := []string{
		"rm\", \"-rf",
		"git\", \"push\", \"--force",
	}

	for _, pattern := range dangerousPatterns {
		if !strings.Contains(text, pattern) {
			t.Errorf("Expected rules to forbid dangerous pattern %q", pattern)
		}
	}
}

func TestGenerateCodexRules_HasComments(t *testing.T) {
	content := generateCodexRules()
	text := string(content)

	if !strings.Contains(text, "#") {
		t.Error("Expected rules file to contain comments")
	}

	if !strings.Contains(text, "# OpenKit Command Policies") {
		t.Error("Expected rules file to have header comment")
	}
}

func TestGenerateAgentsMD_NotEmpty(t *testing.T) {
	content := generateAgentsMD()
	if len(content) == 0 {
		t.Error("Expected generateAgentsMD to return non-empty content")
	}
	if len(content) < 1000 {
		t.Errorf("Expected generateAgentsMD to return substantial content, got %d bytes", len(content))
	}
}

func TestGenerateCodexRules_NotEmpty(t *testing.T) {
	content := generateCodexRules()
	if len(content) == 0 {
		t.Error("Expected generateCodexRules to return non-empty content")
	}
	if len(content) < 500 {
		t.Errorf("Expected generateCodexRules to return substantial content, got %d bytes", len(content))
	}
}
