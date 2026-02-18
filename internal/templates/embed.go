package templates

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/openkit-devtools/openkit/internal/agents"
)

//go:embed base/*
var baseTemplates embed.FS

//go:embed root/*
var rootTemplates embed.FS

// BaseFS exposes the embedded base template filesystem.
func BaseFS() fs.FS {
	return baseTemplates
}

// RootFS exposes the embedded root template filesystem.
func RootFS() fs.FS {
	return rootTemplates
}

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

	// Create agent-specific config files in project root
	if err := createAgentConfig(targetDir, agent); err != nil {
		return fmt.Errorf("failed to create agent config: %w", err)
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

func embeddedExists(efs embed.FS, path string) bool {
	_, err := fs.Stat(efs, path)
	return err == nil
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

		// Skip README.md in commands directory
		if root == "base/commands" && d.Name() == "README.md" {
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

// createAgentConfig creates agent-specific configuration files in project root
func createAgentConfig(targetDir string, agent *agents.Agent) error {
	switch agent.ID {
	case "opencode":
		// Create opencode.json in project root (project config)
		if !embeddedExists(rootTemplates, "root/opencode.json") {
			return fmt.Errorf("missing embedded template root/opencode.json")
		}
		content, err := rootTemplates.ReadFile("root/opencode.json")
		if err != nil {
			return err
		}
		configPath := filepath.Join(targetDir, "opencode.json")
		return os.WriteFile(configPath, content, 0644)

	case "claude":
		// Create .claude/settings.local.json
		settingsDir := filepath.Join(targetDir, ".claude")
		if err := os.MkdirAll(settingsDir, 0755); err != nil {
			return err
		}
		settingsPath := filepath.Join(settingsDir, "settings.local.json")
		settings := `{
  "permissions": {
    "allow": ["Read", "Write", "Edit", "Bash", "WebFetch"]
  },
  "env": {}
}
`
		return os.WriteFile(settingsPath, []byte(settings), 0644)

	case "cursor":
		// Cursor uses .cursorrules in root, handled by ExtraFiles
		return nil

	case "gemini":
		// Gemini config (if needed)
		return nil

	default:
		return nil
	}
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
		"runbooks",
	}

	for _, dir := range dirs {
		path := filepath.Join(docsDir, dir)
		if err := os.MkdirAll(path, 0755); err != nil {
			return err
		}
	}

	// Create basic docs files
	files := map[string]string{
		"HUB-DOCS.md": `# Documentation Index

This file is the primary hub for project documentation.

## Context

Central hub for project documentation. Use this page to understand scope, find domain docs, and navigate to requirements and sprint execution artifacts.

## Status Overview

- Documentation graph: Active
- Canonical hub naming: Enforced (HUB-*.md)
- Last major migration: [YYYY-MM-DD]

## Navigation

- [[CONTEXT.md]]
- [[SECURITY.md]]
- [[QUALITY_GATES.md]]
- [[ACTION_ITEMS.md]]
- [[ARCHITECTURE.md]]
- [[WORKFLOW.md]]
- [[COMMANDS.md]]
- [[SKILLS.md]]
- [[GLOSSARY.md]]
- [[requirements/HUB-REQUIREMENTS.md]]
- [[sprint/HUB-SPRINTS.md]]
- [[adr/HUB-ADR.md]]
- [[runbooks/HUB-RUNBOOKS.md]]

## Related

- [[CONTEXT.md]]
- [[WORKFLOW.md]]
`,
		"GLOSSARY.md": `# Glossary

Standard terminology used across this project.

| Term | Definition |
|---|---|
| [TERM] | [Definition] |

## Related

- [[HUB-DOCS.md]]
- [[CONTEXT.md]]
`,
		"MIGRATION_CHECKLIST.md": `# Migration Checklist

Use this file to migrate legacy documentation to the Obsidian-compatible standard.

## Filename Standard

- [ ] Rename docs files to canonical uppercase names from [[.opencode/rules/DOCS_FILE_GLOSSARY.md]]

## Link Standard

- [ ] Replace internal Markdown links with wikilinks [[...]]
- [ ] Keep external URLs as standard Markdown links

## Hub Notes

- [ ] Ensure [[HUB-DOCS.md]] exists and links major sections
- [ ] Ensure [[requirements/HUB-REQUIREMENTS.md]] exists
- [ ] Ensure [[sprint/HUB-SPRINTS.md]] exists

## Related

- [[HUB-DOCS.md]]
- [[GLOSSARY.md]]
`,
		"ARCHITECTURE.md": `# Architecture

Document your project architecture here.

## Related

- [[WORKFLOW.md]]
- [[COMMANDS.md]]
- [[SKILLS.md]]
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

## Related

- [[WORKFLOW.md]]
- [[ARCHITECTURE.md]]
- [[SKILLS.md]]
`,
		"SKILLS.md": `# Skills Reference

Skills provide domain-specific knowledge for specialized tasks.

See the skills/ folder in your agent configuration for available skills.

## Related

- [[COMMANDS.md]]
- [[WORKFLOW.md]]
- [[ARCHITECTURE.md]]
`,
		"WORKFLOW.md": `# Development Workflow

## Spec-Driven Development (SDD)

1. **Specify**: Create feature specification
2. **Clarify**: Resolve ambiguities
3. **Plan**: Create implementation plan
4. **Tasks**: Break down into executable tasks
5. **Implement**: Execute with AI assistance
6. **Verify**: Run tests and validation

## Related

- [[COMMANDS.md]]
- [[ARCHITECTURE.md]]
- [[SKILLS.md]]
`,
		"requirements/HUB-REQUIREMENTS.md": `# Requirements Index

This hub links feature-level requirements packs.

## Context

Hub for all feature requirement packs. It organizes discovery-to-plan artifacts used as the source of truth before implementation.

## Status Overview

- Feature packs tracked: [N]
- Ready for implementation (PLAN.md + TASKS.md): [N]
- At-risk feature packs: [N]

## Structure

- [[requirements/<feature>/HUB-<FEATURE>.md]]
- [[requirements/<feature>/PROBLEM_STATEMENT.md]]
- [[requirements/<feature>/USER_STORIES.md]]
- [[requirements/<feature>/ACCEPTANCE_CRITERIA.md]]
- [[requirements/<feature>/DATA_CONTRACTS.md]]
- [[requirements/<feature>/RISKS.md]]
- [[requirements/<feature>/PLAN.md]]

## Related

- [[HUB-DOCS.md]]
- [[WORKFLOW.md]]
`,
		"sprint/HUB-SPRINTS.md": `# Sprint Index

This hub links sprint folders and execution artifacts.

## Context

Portfolio-level sprint hub used to navigate all sprint execution packs.

## Status Overview

- Active sprint: [Sprint-XX]
- Completed sprints: [N]
- Delivery health: On Track | At Risk

## Structure

- [[sprint/Sprint-XX/HUB-SPRINT-XX.md]]
- [[sprint/Sprint-XX/SPRINT_GOAL.md]]
- [[sprint/Sprint-XX/BACKLOG.md]]
- [[sprint/Sprint-XX/TASKS.md]]
- [[sprint/Sprint-XX/RISK_REGISTER.md]]

## Related

- [[HUB-DOCS.md]]
- [[WORKFLOW.md]]
`,
		"adr/HUB-ADR.md": `# ADR Index

Architecture Decision Records index.

## Context

Hub for architectural decisions and their rationale.

## Status Overview

- Accepted ADRs: [N]
- Superseded ADRs: [N]
- Pending decisions: [N]

## Related

- [[HUB-DOCS.md]]
- [[ARCHITECTURE.md]]
`,
		"runbooks/HUB-RUNBOOKS.md": `# Runbooks Index

Operational playbooks and incident procedures.

## Context

Operational hub for runbooks used during incidents and routine operations.

## Status Overview

- Critical runbooks covered: [N]
- Last review date: [YYYY-MM-DD]
- Escalation readiness: Ready | Needs Review

## Related

- [[HUB-DOCS.md]]
- [[SECURITY.md]]
- [[QUALITY_GATES.md]]
`,
	}

	for filename, content := range files {
		path := filepath.Join(docsDir, filename)
		if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
			return err
		}
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			return err
		}
	}

	return nil
}

// ListEmbedded returns a list of all embedded template files
func ListEmbedded() []string {
	var files []string

	if err := fs.WalkDir(baseTemplates, "base", func(path string, d fs.DirEntry, err error) error {
		if err == nil && !d.IsDir() {
			files = append(files, strings.TrimPrefix(path, "base/"))
		}
		return nil
	}); err != nil {
		return files
	}

	return files
}
