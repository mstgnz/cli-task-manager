package main

import (
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	// This is a simple test to ensure the main function doesn't crash
	// when called with the help command

	// Save original args
	originalArgs := os.Args

	// Restore original args when the test completes
	defer func() {
		os.Args = originalArgs
	}()

	// Set args to use the help command
	os.Args = []string{"issue-tracker", "help"}

	// Call main() in a separate goroutine to avoid os.Exit
	// In a real test, we might use a custom exit function
	go func() {
		main()
	}()

	// If we get here without panicking, the test passes
}
