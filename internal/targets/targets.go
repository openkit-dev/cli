package targets

import (
	"fmt"
	"strings"

	"github.com/openkit-devtools/openkit/internal/agents"
	"github.com/openkit-devtools/openkit/internal/syncer"
	"github.com/openkit-devtools/openkit/internal/templates"
)

type DesiredResult struct {
	Files []syncer.DesiredFile

	PackID      string
	PackVersion string
}

func BuildEmbeddedDesired(agent *agents.Agent, cliVersion string) (DesiredResult, error) {
	if agent == nil {
		return DesiredResult{}, fmt.Errorf("nil agent")
	}

	switch agent.ID {
	case "opencode":
		return buildOpenCode(agent, cliVersion)
	case "claude":
		return buildClaude(agent, cliVersion)
	case "gemini":
		return buildGemini(agent, cliVersion)
	case "codex":
		return buildCodex(agent, cliVersion)
	case "cursor":
		return buildCursor(agent, cliVersion)
	default:
		return DesiredResult{}, fmt.Errorf("unsupported agent: %s", agent.ID)
	}
}

func buildOpenCode(agent *agents.Agent, cliVersion string) (DesiredResult, error) {
	files, err := syncer.DesiredFromEmbeddedBase(templates.BaseFS(), agent.Folder)
	if err != nil {
		return DesiredResult{}, err
	}

	rootCfg, err := syncer.DesiredRootFile(templates.RootFS(), "root/opencode.json", "opencode.json", "embedded/root/opencode.json")
	if err == nil {
		files = append(files, rootCfg)
	}

	files = append(files, syncer.DesiredFile{
		OutputPath: "AGENTS.md",
		Bytes:      []byte("# Agents\n\nSee docs/AGENTS.md for the complete agent reference.\n"),
		ArtifactID: "embedded/extra/AGENTS.md",
		Mode:       "copy",
	})

	return DesiredResult{Files: files, PackID: "embedded", PackVersion: cliVersion}, nil
}

func buildClaude(agent *agents.Agent, cliVersion string) (DesiredResult, error) {
	var files []syncer.DesiredFile

	rules, err := syncer.DesiredFromEmbeddedSubdir(templates.BaseFS(), "base/rules", ".claude/rules")
	if err != nil {
		return DesiredResult{}, err
	}
	skills, err := syncer.DesiredFromEmbeddedSubdir(templates.BaseFS(), "base/skills", ".claude/skills")
	if err != nil {
		return DesiredResult{}, err
	}
	files = append(files, rules...)
	files = append(files, skills...)

	// Root prompt: reuse orchestrator prompt as the primary instruction file.
	rootPrompt, err := syncer.DesiredFromEmbeddedSubdir(templates.BaseFS(), "base/prompts", ".claude/agents")
	if err == nil {
		files = append(files, rootPrompt...)
	}

	// Preferred entrypoint.
	claudeMd := []byte("# OpenKit (Claude Code)\n\nThis project uses OpenKit content synced by OpenKit CLI.\n\n- Rules: .claude/rules/\n- Skills: .claude/skills/\n- Agents: .claude/agents/\n\nRun Claude Code and follow the project instructions.\n")
	files = append(files, syncer.DesiredFile{
		OutputPath: ".claude/CLAUDE.md",
		Bytes:      claudeMd,
		ArtifactID: "embedded/extra/.claude/CLAUDE.md",
		Mode:       "copy",
	})

	// Project settings (syncable). Do NOT manage settings.local.json.
	settings := []byte("{\n  \"permissions\": {\n    \"allow\": [\"Read\", \"Write\", \"Edit\", \"Bash\", \"WebFetch\"]\n  },\n  \"env\": {}\n}\n")
	files = append(files, syncer.DesiredFile{
		OutputPath: ".claude/settings.json",
		Bytes:      settings,
		ArtifactID: "embedded/extra/.claude/settings.json",
		Mode:       "copy",
	})

	return DesiredResult{Files: files, PackID: "embedded", PackVersion: cliVersion}, nil
}

func buildGemini(agent *agents.Agent, cliVersion string) (DesiredResult, error) {
	geminiMd := []byte(strings.Join([]string{
		"# OpenKit (Gemini CLI)",
		"",
		"This project uses OpenKit content synced by OpenKit CLI.",
		"",
		"## What OpenKit installs",
		"",
		"- Commands: `.gemini/commands/openkit/*.toml` (run as `/openkit:<command>`)",
		"- Rules: `.gemini/rules/*.md` (referenced, not embedded)",
		"- Skills: `.gemini/skills/<name>/SKILL.md` (content for reference)",
		"",
		"## Notes",
		"",
		"- Gemini CLI may ignore project commands unless the repo is trusted.",
		"- Sync never executes commands; it only writes files tracked in `.openkit/managed.json`.",
		"",
	}, "\n"))
	settings := []byte("{\n  \"version\": 1\n}\n")

	var files []syncer.DesiredFile
	files = append(files,
		syncer.DesiredFile{OutputPath: "GEMINI.md", Bytes: geminiMd, ArtifactID: "embedded/extra/GEMINI.md", Mode: "copy"},
		syncer.DesiredFile{OutputPath: ".gemini/settings.json", Bytes: settings, ArtifactID: "embedded/extra/.gemini/settings.json", Mode: "copy"},
	)

	rules, err := syncer.DesiredFromEmbeddedSubdir(templates.BaseFS(), "base/rules", ".gemini/rules")
	if err != nil {
		return DesiredResult{}, err
	}
	skills, err := syncer.DesiredFromEmbeddedSubdir(templates.BaseFS(), "base/skills", ".gemini/skills")
	if err != nil {
		return DesiredResult{}, err
	}
	cmds, err := desiredGeminiCommands(templates.BaseFS())
	if err != nil {
		return DesiredResult{}, err
	}

	files = append(files, rules...)
	files = append(files, skills...)
	files = append(files, cmds...)

	return DesiredResult{Files: files, PackID: "embedded", PackVersion: cliVersion}, nil
}

func buildCodex(agent *agents.Agent, cliVersion string) (DesiredResult, error) {
	var files []syncer.DesiredFile

	// Generate comprehensive AGENTS.md
	agentsMd := generateAgentsMD()
	files = append(files, syncer.DesiredFile{
		OutputPath: "AGENTS.md",
		Bytes:      agentsMd,
		ArtifactID: "embedded/extra/AGENTS.md",
		Mode:       "copy",
	})

	// Generate Codex rules file
	rulesContent := generateCodexRules()
	files = append(files, syncer.DesiredFile{
		OutputPath: ".codex/rules/openkit.rules",
		Bytes:      rulesContent,
		ArtifactID: "embedded/extra/.codex/rules/openkit.rules",
		Mode:       "copy",
	})

	// Copy skills
	skills, err := syncer.DesiredFromEmbeddedSubdir(templates.BaseFS(), "base/skills", ".agents/skills")
	if err != nil {
		return DesiredResult{}, err
	}
	files = append(files, skills...)

	return DesiredResult{Files: files, PackID: "embedded", PackVersion: cliVersion}, nil
}

func buildCursor(agent *agents.Agent, cliVersion string) (DesiredResult, error) {
	var files []syncer.DesiredFile

	// Generate .cursorrules (backward compatibility)
	cursorrules := generateCursorRules()
	files = append(files, syncer.DesiredFile{
		OutputPath: ".cursorrules",
		Bytes:      cursorrules,
		ArtifactID: "embedded/extra/.cursorrules",
		Mode:       "copy",
	})

	// Generate .cursor/rules/openkit.mdc (modular rules)
	mdcContent := generateCursorMDC()
	files = append(files, syncer.DesiredFile{
		OutputPath: ".cursor/rules/openkit.mdc",
		Bytes:      mdcContent,
		ArtifactID: "embedded/extra/.cursor/rules/openkit.mdc",
		Mode:       "copy",
	})

	// Copy skills
	skills, err := syncer.DesiredFromEmbeddedSubdir(templates.BaseFS(), "base/skills", ".cursor/skills")
	if err != nil {
		return DesiredResult{}, err
	}
	files = append(files, skills...)

	return DesiredResult{Files: files, PackID: "embedded", PackVersion: cliVersion}, nil
}
