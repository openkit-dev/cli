package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/openkit-devtools/openkit/internal/agents"
	"github.com/openkit-devtools/openkit/internal/templates"
	"github.com/spf13/cobra"
)

var (
	flagAgent string
	flagHere  bool
	flagForce bool
	flagNoGit bool
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

	fmt.Println()
	green.Println("Project initialized successfully!")
	fmt.Println()
	printInfo("Next steps:")
	if !flagHere {
		fmt.Printf("    cd %s\n", projectName)
	}
	fmt.Printf("    %s   # Start your AI agent\n", agent.CLICommand)
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
