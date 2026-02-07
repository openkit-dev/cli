package cli

import (
	"os/exec"
)

// newCommand creates a new exec.Command with proper setup
func newCommand(name string, args ...string) *exec.Cmd {
	return exec.Command(name, args...)
}
