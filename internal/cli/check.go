package cli

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check system requirements and installed AI agents",
	Long: `Verify that your system meets the requirements for OpenKit and detect
which AI coding agents are installed and available.`,
	Run: func(cmd *cobra.Command, args []string) {
		runSystemCheck()
	},
}

type checkResult struct {
	name    string
	status  bool
	version string
	note    string
}

func runSystemCheck() {
	cyan := color.New(color.FgCyan, color.Bold)
	green := color.New(color.FgGreen)
	red := color.New(color.FgRed)
	dim := color.New(color.FgHiBlack)

	cyan.Println("\nSystem Check")
	cyan.Println("============")
	fmt.Println()

	// Platform info
	fmt.Printf("  Platform: %s/%s\n", runtime.GOOS, runtime.GOARCH)
	fmt.Println()

	// Check AI agents
	cyan.Println("AI Agents")
	cyan.Println("---------")

	agents := []struct {
		name    string
		command string
		args    []string
	}{
		{"OpenCode", "opencode", []string{"--version"}},
		{"Claude Code", "claude", []string{"--version"}},
		{"Cursor", "cursor", []string{"--version"}},
		{"Gemini CLI", "gemini", []string{"--version"}},
		{"Codex CLI", "codex", []string{"--version"}},
	}

	foundAgent := false
	for _, agent := range agents {
		result := checkCommand(agent.name, agent.command, agent.args)
		printCheckResult(result, green, red, dim)
		if result.status {
			foundAgent = true
		}
	}

	fmt.Println()

	// Check development tools
	cyan.Println("Development Tools")
	cyan.Println("-----------------")

	tools := []struct {
		name    string
		command string
		args    []string
	}{
		{"Git", "git", []string{"--version"}},
		{"Node.js", "node", []string{"--version"}},
		{"Python", "python3", []string{"--version"}},
		{"Go", "go", []string{"version"}},
	}

	for _, tool := range tools {
		result := checkCommand(tool.name, tool.command, tool.args)
		printCheckResult(result, green, red, dim)
	}

	fmt.Println()

	// Summary
	if foundAgent {
		green.Println("  Ready to use OpenKit!")
	} else {
		red.Println("  No AI agents detected. Install one of the supported agents:")
		dim.Println("    - OpenCode: curl -fsSL https://opencode.ai/install | bash")
		dim.Println("    - Claude:   npm install -g @anthropic-ai/claude-code")
	}
	fmt.Println()
}

func checkCommand(name, command string, args []string) checkResult {
	result := checkResult{
		name:   name,
		status: false,
	}

	path, err := exec.LookPath(command)
	if err != nil {
		result.note = "not found"
		return result
	}

	cmd := exec.Command(path, args...)
	output, err := cmd.Output()
	if err != nil {
		result.note = "error running command"
		return result
	}

	result.status = true
	result.version = string(output)
	if len(result.version) > 50 {
		result.version = result.version[:50] + "..."
	}
	
	// Clean up version string
	for i, c := range result.version {
		if c == '\n' || c == '\r' {
			result.version = result.version[:i]
			break
		}
	}

	return result
}

func printCheckResult(result checkResult, green, red, dim *color.Color) {
	if result.status {
		green.Print("  [OK] ")
		fmt.Printf("%-15s", result.name)
		dim.Printf(" %s\n", result.version)
	} else {
		red.Print("  [--] ")
		fmt.Printf("%-15s", result.name)
		dim.Printf(" %s\n", result.note)
	}
}

func detectShell() string {
	shell := os.Getenv("SHELL")
	if shell == "" {
		if runtime.GOOS == "windows" {
			return "powershell"
		}
		return "unknown"
	}
	return shell
}
