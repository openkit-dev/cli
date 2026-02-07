package cli

import (
	"fmt"
	"runtime"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	Long:  `Display the version, build date, and runtime information for OpenKit.`,
	Run: func(cmd *cobra.Command, args []string) {
		printVersionInfo()
	},
}

func printVersionInfo() {
	cyan := color.New(color.FgCyan, color.Bold)
	white := color.New(color.FgWhite)
	dim := color.New(color.FgHiBlack)

	cyan.Println("OpenKit CLI")
	fmt.Println()
	
	white.Print("  Version:    ")
	fmt.Println(version)
	
	white.Print("  Commit:     ")
	dim.Println(commit)
	
	white.Print("  Built:      ")
	dim.Println(buildDate)
	
	white.Print("  Go version: ")
	dim.Println(runtime.Version())
	
	white.Print("  OS/Arch:    ")
	dim.Printf("%s/%s\n", runtime.GOOS, runtime.GOARCH)
}
