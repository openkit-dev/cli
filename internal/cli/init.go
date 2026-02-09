package cli

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/openkit-devtools/openkit/internal/agents"
	"github.com/openkit-devtools/openkit/internal/templates"
	"github.com/spf13/cobra"
)

var (
	flagAgent  string
	flagHere   bool
	flagForce  bool
	flagNoGit  bool
	flagMemory bool
)

var initCmd = &cobra.Command{
	Use:   "init [project-name]",
	Short: "Initialize a new project with OpenKit SDD templates",
	Long: `Initialize a new project directory with OpenKit's Spec-Driven Development
templates, configured for your preferred AI coding agent.

Examples:
  openkit init my-app                  # Create new project with interactive agent selection
  openkit init my-app --ai opencode    # Create project for OpenCode
  openkit init my-app --ai claude      # Create project for Claude Code
  openkit init --here                  # Initialize in current directory`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runInit(args)
	},
}

func init() {
	initCmd.Flags().StringVar(&flagAgent, "ai", "", "AI agent to configure (opencode, claude, cursor, gemini)")
	initCmd.Flags().BoolVar(&flagHere, "here", false, "Initialize in current directory")
	initCmd.Flags().BoolVar(&flagForce, "force", false, "Overwrite existing files")
	initCmd.Flags().BoolVar(&flagNoGit, "no-git", false, "Skip git initialization")
	initCmd.Flags().BoolVar(&flagMemory, "memory", false, "Enable semantic memory plugin")
}

func runInit(args []string) {
	cyan := color.New(color.FgCyan, color.Bold)
	green := color.New(color.FgGreen, color.Bold)

	// Determine project directory
	var projectDir string
	var projectName string

	if flagHere {
		cwd, err := os.Getwd()
		if err != nil {
			exitWithError(fmt.Sprintf("Failed to get current directory: %v", err))
		}
		projectDir = cwd
		projectName = filepath.Base(cwd)
	} else if len(args) > 0 {
		projectName = args[0]
		cwd, err := os.Getwd()
		if err != nil {
			exitWithError(fmt.Sprintf("Failed to get current directory: %v", err))
		}
		projectDir = filepath.Join(cwd, projectName)
	} else {
		exitWithError("Project name required. Use 'openkit init <name>' or 'openkit init --here'")
	}

	// Check if directory exists
	if !flagHere {
		if _, err := os.Stat(projectDir); err == nil {
			if !flagForce {
				exitWithError(fmt.Sprintf("Directory '%s' already exists. Use --force to overwrite.", projectName))
			}
		}
	}

	// Get agent configuration
	var agent *agents.Agent
	if flagAgent != "" {
		a := agents.Get(flagAgent)
		if a == nil {
			exitWithError(fmt.Sprintf("Unknown agent '%s'. Available: opencode, claude, cursor, gemini", flagAgent))
		}
		agent = a
	} else {
		// Interactive selection (TODO: implement Bubble Tea selector)
		// For now, default to OpenCode
		agent = agents.Get("opencode")
		printInfo("No agent specified, using OpenCode. Use --ai to specify an agent.")
	}

	cyan.Printf("\nInitializing OpenKit project: %s\n", projectName)
	cyan.Printf("Agent: %s\n\n", agent.Name)

	// Create project directory
	if !flagHere {
		if err := os.MkdirAll(projectDir, 0755); err != nil {
			exitWithError(fmt.Sprintf("Failed to create directory: %v", err))
		}
	}

	// Extract templates
	printInfo("Extracting templates...")
	if err := templates.Extract(projectDir, agent); err != nil {
		exitWithError(fmt.Sprintf("Failed to extract templates: %v", err))
	}

	// Initialize git
	if !flagNoGit {
		printInfo("Initializing git repository...")
		if err := initGit(projectDir); err != nil {
			printWarning(fmt.Sprintf("Git initialization failed: %v", err))
		}
	}

	// Install memory plugin if requested
	if flagMemory {
		printInfo("Installing semantic memory plugin...")
		if err := installMemoryPlugin(projectDir); err != nil {
			printWarning(fmt.Sprintf("Memory plugin installation failed: %v", err))
		} else {
			cyan.Println("✓ Semantic memory enabled")

			// Install npm dependencies
			printInfo("Installing dependencies (this may take a moment)...")
			if err := installMemoryDependencies(projectDir); err != nil {
				printWarning(fmt.Sprintf("Dependency installation failed: %v", err))
				printInfo("You can install manually: cd .opencode && npm install")
			} else {
				cyan.Println("✓ Dependencies installed")
			}
		}
	}

	fmt.Println()
	green.Println("Project initialized successfully!")
	fmt.Println()
	printInfo("Next steps:")
	if !flagHere {
		fmt.Printf("    cd %s\n", projectName)
	}
	fmt.Printf("    %s   # Start your AI agent\n", agent.CLICommand)
	if flagMemory {
		fmt.Println("    # Memory plugin will automatically capture context across sessions")
	}
	fmt.Println()
}

func initGit(dir string) error {
	// Check if already a git repo
	gitDir := filepath.Join(dir, ".git")
	if _, err := os.Stat(gitDir); err == nil {
		return nil // Already initialized
	}

	// Initialize git
	cmd := newCommand("git", "init")
	cmd.Dir = dir
	return cmd.Run()
}

func installMemoryPlugin(projectDir string) error {
	// 1. Create plugin directory
	pluginDir := filepath.Join(projectDir, ".opencode", "plugins", "semantic-memory")
	if err := os.MkdirAll(pluginDir, 0755); err != nil {
		return fmt.Errorf("failed to create plugin directory: %w", err)
	}

	// 2. Create memory data directory
	memoryDir := filepath.Join(projectDir, ".opencode", "memory")
	if err := os.MkdirAll(memoryDir, 0755); err != nil {
		return fmt.Errorf("failed to create memory directory: %w", err)
	}

	// 3. Extract memory plugin from embedded templates
	if err := templates.ExtractMemoryPlugin(pluginDir); err != nil {
		return fmt.Errorf("failed to extract plugin: %w", err)
	}

	// 4. Create default config.json
	configPath := filepath.Join(memoryDir, "config.json")
	configContent := `{
  "version": "1.0.0",
  "embedding": {
    "model": "nomic-embed-text",
    "runtime": "onnx"
  },
  "retrieval": {
    "max_results": 10,
    "min_similarity": 0.7,
    "token_budget": 4000
  },
  "curation": {
    "ttl_days": 90,
    "max_per_project": 500,
    "prune_unused_after_days": 30
  },
  "extraction": {
    "on_session_idle": true,
    "patterns": ["decision", "architecture", "pattern", "fix", "solution"]
  },
  "debug": {
    "verbose": false,
    "show_injection_indicator": true
  }
}
`
	if err := os.WriteFile(configPath, []byte(configContent), 0644); err != nil {
		return fmt.Errorf("failed to write config: %w", err)
	}

	// 5. Create .gitignore for memory data
	gitignorePath := filepath.Join(memoryDir, ".gitignore")
	gitignoreContent := "index.lance/\nmodel/\nmetrics.json\n"
	if err := os.WriteFile(gitignorePath, []byte(gitignoreContent), 0644); err != nil {
		return fmt.Errorf("failed to write .gitignore: %w", err)
	}

	// 6. Extract memory rules to .opencode/rules/
	rulesDir := filepath.Join(projectDir, ".opencode", "rules")
	if err := templates.ExtractMemoryRules(rulesDir); err != nil {
		return fmt.Errorf("failed to extract memory rules: %w", err)
	}

	// 7. Create/update .opencode/package.json with memory plugin dependencies
	if err := updateOpencodePackageJson(projectDir); err != nil {
		return fmt.Errorf("failed to update package.json: %w", err)
	}

	// 8. Update opencode.json
	if err := updateOpencodeJsonMemory(projectDir, true); err != nil {
		return fmt.Errorf("failed to update opencode.json: %w", err)
	}

	return nil
}

func updateOpencodePackageJson(projectDir string) error {
	packagePath := filepath.Join(projectDir, ".opencode", "package.json")

	// Package.json content with memory plugin dependencies
	// OpenCode reads this file and runs bun/npm install automatically
	// "type": "module" is required for ES module imports in the plugin
	packageContent := `{
  "type": "module",
  "dependencies": {
    "@opencode-ai/plugin": "^1.1.0",
    "@lancedb/lancedb": "^0.26.0",
    "onnxruntime-node": "^1.24.0"
  }
}
`
	return os.WriteFile(packagePath, []byte(packageContent), 0644)
}

func updateOpencodeJsonMemory(projectDir string, enable bool) error {
	// Local plugins in .opencode/plugins/ are loaded automatically by OpenCode
	// No need to modify opencode.json - the "plugin" array is only for npm packages
	// See: https://opencode.ai/docs/plugins#from-local-files
	return nil
}

func installMemoryDependencies(projectDir string) error {
	opencodeDir := filepath.Join(projectDir, ".opencode")

	// Check if package.json exists
	packagePath := filepath.Join(opencodeDir, "package.json")
	if _, err := os.Stat(packagePath); os.IsNotExist(err) {
		return fmt.Errorf("package.json not found")
	}

	// Try npm install (most common)
	cmd := newCommand("npm", "install", "--silent")
	cmd.Dir = opencodeDir

	// Suppress output unless there's an error
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		// If npm fails, try bun (OpenCode may use bun)
		cmd = newCommand("bun", "install")
		cmd.Dir = opencodeDir
		cmd.Stderr = &stderr

		if err := cmd.Run(); err != nil {
			return fmt.Errorf("npm/bun install failed: %s", stderr.String())
		}
	}

	return nil
}
