package cli

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	version   = "dev"
	commit    = "none"
	buildDate = "unknown"
)

// SetVersionInfo sets version information from build flags
func SetVersionInfo(v, c, d string) {
	version = v
	commit = c
	buildDate = d
}

// GetVersion returns the current version
func GetVersion() string {
	return version
}

var rootCmd = &cobra.Command{
	Use:   "openkit",
	Short: "Universal Spec-Driven Development toolkit",
	Long: `OpenKit is a universal Spec-Driven Development toolkit that works with
multiple AI coding agents including OpenCode, Claude, Cursor, and Gemini.

It provides a consistent SDD workflow across all supported agents with
embedded templates, commands, skills, and prompts.`,
	Run: func(cmd *cobra.Command, args []string) {
		printBanner()
		fmt.Println()
		cmd.Help()
	},
}

func printBanner() {
	cyan := color.New(color.FgCyan, color.Bold)
	white := color.New(color.FgWhite)
	
	banner := `
   ___                   _  ___ _   
  / _ \ _ __   ___ _ __ | |/ (_) |_ 
 | | | | '_ \ / _ \ '_ \| ' /| | __|
 | |_| | |_) |  __/ | | | . \| | |_ 
  \___/| .__/ \___|_| |_|_|\_\_|\__|
       |_|                          `

	cyan.Println(banner)
	white.Printf("  Universal Spec-Driven Development Toolkit v%s\n", version)
}

// Execute runs the root command
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Add subcommands
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(checkCmd)
	rootCmd.AddCommand(initCmd)
	
	// Disable completion command for cleaner help
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}

// exitWithError prints an error message and exits
func exitWithError(msg string) {
	red := color.New(color.FgRed, color.Bold)
	red.Fprintf(os.Stderr, "Error: %s\n", msg)
	os.Exit(1)
}

// printSuccess prints a success message
func printSuccess(msg string) {
	green := color.New(color.FgGreen, color.Bold)
	green.Printf("  %s\n", msg)
}

// printInfo prints an info message
func printInfo(msg string) {
	cyan := color.New(color.FgCyan)
	cyan.Printf("  %s\n", msg)
}

// printWarning prints a warning message
func printWarning(msg string) {
	yellow := color.New(color.FgYellow)
	yellow.Printf("  %s\n", msg)
}
