package ui

import (
	"testing"
)

func TestOutputs(t *testing.T) {
	// Simple smoke test to ensure no panics
	Success("Test success %s", "msg")
	Error("Test error %s", "msg")
	Warning("Test warning %s", "msg")
	Info("Test info %s", "msg")
}
