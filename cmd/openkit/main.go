package main

import (
	"os"

	"github.com/openkit-devtools/openkit/internal/cli"
)

// Version information (set by goreleaser)
var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	cli.SetVersionInfo(version, commit, date)

	if err := cli.Execute(); err != nil {
		os.Exit(1)
	}
}
