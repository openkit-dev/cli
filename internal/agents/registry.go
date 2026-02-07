package agents

import "strings"

// Agent represents a supported AI coding agent
type Agent struct {
	ID          string   // Unique identifier (e.g., "opencode", "claude")
	Name        string   // Display name
	Folder      string   // Configuration folder (e.g., ".opencode/", ".claude/")
	ExtraFiles  []string // Additional files to create (e.g., ["CLAUDE.md"])
	RequiresCLI bool     // Whether CLI tool must be installed
	CLICommand  string   // Command to run the agent
	InstallURL  string   // Documentation URL for installation
}

// Registry of supported agents
var registry = map[string]*Agent{
	"opencode": {
		ID:          "opencode",
		Name:        "OpenCode",
		Folder:      ".opencode",
		ExtraFiles:  []string{"AGENTS.md"},
		RequiresCLI: true,
		CLICommand:  "opencode",
		InstallURL:  "https://opencode.ai",
	},
	"claude": {
		ID:          "claude",
		Name:        "Claude Code",
		Folder:      ".claude",
		ExtraFiles:  []string{"CLAUDE.md"},
		RequiresCLI: true,
		CLICommand:  "claude",
		InstallURL:  "https://docs.anthropic.com/claude-code",
	},
	"cursor": {
		ID:          "cursor",
		Name:        "Cursor",
		Folder:      ".cursor",
		ExtraFiles:  []string{".cursorrules"},
		RequiresCLI: false,
		CLICommand:  "cursor",
		InstallURL:  "https://cursor.com",
	},
	"gemini": {
		ID:          "gemini",
		Name:        "Gemini CLI",
		Folder:      ".gemini",
		ExtraFiles:  []string{"GEMINI.md"},
		RequiresCLI: true,
		CLICommand:  "gemini",
		InstallURL:  "https://github.com/google-gemini/gemini-cli",
	},
	"codex": {
		ID:          "codex",
		Name:        "Codex CLI",
		Folder:      ".codex",
		ExtraFiles:  []string{},
		RequiresCLI: true,
		CLICommand:  "codex",
		InstallURL:  "https://github.com/openai/codex",
	},
	"windsurf": {
		ID:          "windsurf",
		Name:        "Windsurf",
		Folder:      ".windsurf",
		ExtraFiles:  []string{},
		RequiresCLI: false,
		CLICommand:  "windsurf",
		InstallURL:  "https://codeium.com/windsurf",
	},
}

// Get returns an agent by ID (case-insensitive)
func Get(id string) *Agent {
	return registry[strings.ToLower(id)]
}

// All returns all registered agents
func All() []*Agent {
	agents := make([]*Agent, 0, len(registry))
	// Return in specific order for consistent display
	order := []string{"opencode", "claude", "cursor", "gemini", "codex", "windsurf"}
	for _, id := range order {
		if agent, ok := registry[id]; ok {
			agents = append(agents, agent)
		}
	}
	return agents
}

// IDs returns all agent IDs
func IDs() []string {
	return []string{"opencode", "claude", "cursor", "gemini", "codex", "windsurf"}
}
