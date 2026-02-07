package templates

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/openkit-dev/cli/internal/agents"
)

//go:embed base/*
var baseTemplates embed.FS

//go:embed agents/*
var agentTemplates embed.FS

// Extract copies templates to the target directory for the specified agent
func Extract(targetDir string, agent *agents.Agent) error {
	// Create agent configuration folder
	agentDir := filepath.Join(targetDir, agent.Folder)
	if err := os.MkdirAll(agentDir, 0755); err != nil {
		return fmt.Errorf("failed to create agent directory: %w", err)
	}

	// Extract base templates to agent folder
	if err := extractFS(baseTemplates, "base", agentDir); err != nil {
		return fmt.Errorf("failed to extract base templates: %w", err)
	}

	// Extract agent-specific templates (if they exist)
	agentTemplatePath := filepath.Join("agents", agent.ID)
	if err := extractFS(agentTemplates, agentTemplatePath, agentDir); err != nil {
		// Not an error if agent-specific templates don't exist
		if !os.IsNotExist(err) {
			// Log but don't fail - agent templates are optional
		}
	}

	// Create extra files for the agent
	for _, extraFile := range agent.ExtraFiles {
		extraPath := filepath.Join(targetDir, extraFile)
		if err := createExtraFile(extraPath, agent); err != nil {
			return fmt.Errorf("failed to create %s: %w", extraFile, err)
		}
	}

	// Create docs folder
	docsDir := filepath.Join(targetDir, "docs")
	if err := os.MkdirAll(docsDir, 0755); err != nil {
		return fmt.Errorf("failed to create docs directory: %w", err)
	}

	// Create basic docs structure
	if err := createDocsStructure(docsDir); err != nil {
		return fmt.Errorf("failed to create docs structure: %w", err)
	}

	return nil
}

// extractFS extracts files from an embedded filesystem to a target directory
func extractFS(efs embed.FS, root, targetDir string) error {
	return fs.WalkDir(efs, root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Calculate relative path
		relPath, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}

		// Skip root
		if relPath == "." {
			return nil
		}

		targetPath := filepath.Join(targetDir, relPath)

		if d.IsDir() {
			return os.MkdirAll(targetPath, 0755)
		}

		// Read and write file
		content, err := efs.ReadFile(path)
		if err != nil {
			return err
		}

		return os.WriteFile(targetPath, content, 0644)
	})
}

// createExtraFile creates agent-specific extra files
func createExtraFile(path string, agent *agents.Agent) error {
	filename := filepath.Base(path)
	
	var content string
	switch filename {
	case "AGENTS.md":
		content = fmt.Sprintf(`# Agents

See docs/AGENTS.md for the complete agent reference.

## Quick Start

This project uses OpenKit with %s.

Run %s to start your AI coding session.
`, agent.Name, agent.CLICommand)
	
	case "CLAUDE.md":
		content = `# Claude Code Instructions

This project uses OpenKit's Spec-Driven Development methodology.

## Available Commands

Use the slash commands in .claude/commands/ for SDD workflow:
- /openkit.specify - Create feature specification
- /openkit.plan - Create implementation plan
- /openkit.implement - Execute implementation
- /openkit.analyze - Validate spec/plan

## Skills

Specialized skills are available in .claude/skills/.

## Rules

Follow the rules in .claude/rules/ for consistent behavior.
`

	case "GEMINI.md":
		content = `# Gemini CLI Instructions

This project uses OpenKit's Spec-Driven Development methodology.

See .gemini/ for commands, skills, and rules.
`

	case ".cursorrules":
		content = `# Cursor Rules

This project uses OpenKit's Spec-Driven Development methodology.

See .cursor/rules/ for detailed rules.
See .cursor/skills/ for specialized domain knowledge.
`

	default:
		content = fmt.Sprintf("# %s Configuration\n\nThis project uses OpenKit.\n", agent.Name)
	}

	return os.WriteFile(path, []byte(content), 0644)
}

// createDocsStructure creates the basic docs folder structure
func createDocsStructure(docsDir string) error {
	// Create subdirectories
	dirs := []string{
		"requirements",
		"sprint",
		"adr",
	}

	for _, dir := range dirs {
		path := filepath.Join(docsDir, dir)
		if err := os.MkdirAll(path, 0755); err != nil {
			return err
		}
	}

	// Create basic docs files
	files := map[string]string{
		"ARCHITECTURE.md": `# Architecture

Document your project architecture here.
`,
		"COMMANDS.md": `# Commands Reference

OpenKit provides the following slash commands for SDD workflow:

| Command | Purpose |
|---------|---------|
| /specify | Create feature specification |
| /clarify | Resolve spec ambiguities |
| /plan | Create implementation plan |
| /tasks | Generate executable tasks |
| /analyze | Validate spec/plan/tasks |
| /impl | Execute implementation |
| /test | Generate or run tests |
| /debug | Investigate complex errors |
`,
		"SKILLS.md": `# Skills Reference

Skills provide domain-specific knowledge for specialized tasks.

See the skills/ folder in your agent configuration for available skills.
`,
		"WORKFLOW.md": `# Development Workflow

## Spec-Driven Development (SDD)

1. **Specify**: Create feature specification
2. **Clarify**: Resolve ambiguities
3. **Plan**: Create implementation plan
4. **Tasks**: Break down into executable tasks
5. **Implement**: Execute with AI assistance
6. **Verify**: Run tests and validation
`,
	}

	for filename, content := range files {
		path := filepath.Join(docsDir, filename)
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			return err
		}
	}

	return nil
}

// ListEmbedded returns a list of all embedded template files
func ListEmbedded() []string {
	var files []string
	
	fs.WalkDir(baseTemplates, "base", func(path string, d fs.DirEntry, err error) error {
		if err == nil && !d.IsDir() {
			files = append(files, strings.TrimPrefix(path, "base/"))
		}
		return nil
	})
	
	return files
}
