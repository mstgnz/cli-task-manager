package commands

import (
	"testing"

	"github.com/mstgnz/cli-task-manager/storage"
)

func TestParseArgs(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected map[string]string
	}{
		{
			name:     "Empty args",
			args:     []string{},
			expected: map[string]string{},
		},
		{
			name:     "Single main argument",
			args:     []string{"main-arg"},
			expected: map[string]string{"main": "main-arg"},
		},
		{
			name:     "Main argument with flag",
			args:     []string{"main-arg", "--flag", "value"},
			expected: map[string]string{"main": "main-arg", "flag": "value"},
		},
		{
			name:     "Multiple flags",
			args:     []string{"--flag1", "value1", "--flag2", "value2"},
			expected: map[string]string{"flag1": "value1", "flag2": "value2"},
		},
		{
			name:     "Boolean flag",
			args:     []string{"--flag"},
			expected: map[string]string{"flag": "true"},
		},
		{
			name:     "Mixed flags",
			args:     []string{"main-arg", "--flag1", "value1", "--flag2"},
			expected: map[string]string{"main": "main-arg", "flag1": "value1", "flag2": "true"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parseArgs(tt.args)

			// Check if the maps have the same length
			if len(result) != len(tt.expected) {
				t.Errorf("Expected map length %d, got %d", len(tt.expected), len(result))
			}

			// Check if all expected keys and values are present
			for k, v := range tt.expected {
				if result[k] != v {
					t.Errorf("Expected %s=%s, got %s=%s", k, v, k, result[k])
				}
			}
		})
	}
}

func TestNewApp(t *testing.T) {
	// This is a simple test to ensure NewApp doesn't crash
	// In a real test, we might mock the file system
	app, err := NewApp()
	if err != nil {
		t.Fatalf("Failed to create app: %v", err)
	}

	if app == nil {
		t.Fatal("Expected app to not be nil")
	}

	if app.storage == nil {
		t.Fatal("Expected app.storage to not be nil")
	}
}

func TestAppRun(t *testing.T) {
	// Create a mock app with mock storage
	app := &App{
		storage: storage.NewMockStorage(),
	}

	// Test help command
	err := app.Run([]string{"issue-tracker", "help"})
	if err != nil {
		t.Errorf("Expected no error for help command, got %v", err)
	}

	// Test unknown command
	err = app.Run([]string{"issue-tracker", "unknown"})
	if err != nil {
		t.Errorf("Expected no error for unknown command, got %v", err)
	}

	// Test no command (should show usage)
	err = app.Run([]string{"issue-tracker"})
	if err != nil {
		t.Errorf("Expected no error for no command, got %v", err)
	}
}
