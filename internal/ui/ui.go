package ui

import (
	"fmt"
	"github.com/fatih/color"
	"os"
)

var (
	green  = color.New(color.FgGreen).SprintFunc()
	red    = color.New(color.FgRed).SprintFunc()
	yellow = color.New(color.FgYellow).SprintFunc()
	cyan   = color.New(color.FgCyan).SprintFunc()
)

func Success(format string, a ...interface{}) {
	fmt.Printf("%s %s\n", green("✔"), fmt.Sprintf(format, a...))
}

func Error(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, "%s %s\n", red("✘"), fmt.Sprintf(format, a...))
}

func Warning(format string, a ...interface{}) {
	fmt.Printf("%s %s\n", yellow("!"), fmt.Sprintf(format, a...))
}

func Info(format string, a ...interface{}) {
	fmt.Printf("%s %s\n", cyan("ℹ"), fmt.Sprintf(format, a...))
}
