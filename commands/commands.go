package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mstgnz/cli-task-manager/storage"
)

// App represents the CLI application
type App struct {
	storage storage.Storage
}

// NewApp creates a new CLI application
func NewApp() (*App, error) {
	// Get user's home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get home directory: %w", err)
	}

	// Create data directory in user's home directory
	dataDir := filepath.Join(homeDir, ".cli-task-manager")
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create data directory: %w", err)
	}

	// Create JSON storage
	jsonStorage, err := storage.NewJSONStorage(filepath.Join(dataDir, "tasks.json"))
	if err != nil {
		return nil, fmt.Errorf("failed to create JSON storage: %w", err)
	}

	return &App{
		storage: jsonStorage,
	}, nil
}

// Run executes the CLI application with the given arguments
func (a *App) Run(args []string) error {
	if len(args) < 2 {
		a.printUsage()
		return nil
	}

	command := args[1]

	switch command {
	case "add":
		return a.handleAdd(args[2:])
	case "list":
		return a.handleList(args[2:])
	case "update":
		return a.handleUpdate(args[2:])
	case "filter":
		return a.handleFilter(args[2:])
	case "remove":
		return a.handleRemove(args[2:])
	case "help":
		a.printUsage()
		return nil
	default:
		fmt.Printf("Unknown command: %s\n", command)
		a.printUsage()
		return nil
	}
}

// printUsage prints the usage information
func (a *App) printUsage() {
	fmt.Println("CLI Task Manager - A lightweight issue tracker for the terminal")
	fmt.Println("\nUsage:")
	fmt.Println("  issue-tracker <command> [arguments]")
	fmt.Println("\nCommands:")
	fmt.Println("  add <title> --label <label>                Add a new task")
	fmt.Println("  list                                       List all tasks")
	fmt.Println("  update <id> --status <status>              Update task status")
	fmt.Println("  filter --label <label>                     Filter tasks by label")
	fmt.Println("  remove <id>                                Remove a task")
	fmt.Println("  help                                       Show this help message")
	fmt.Println("\nExamples:")
	fmt.Println("  issue-tracker add \"Create API documentation\" --label feature")
	fmt.Println("  issue-tracker update 1 --status in-progress")
	fmt.Println("  issue-tracker filter --label bug")
}

// parseArgs parses command line arguments into a map
func parseArgs(args []string) map[string]string {
	result := make(map[string]string)

	for i := 0; i < len(args); i++ {
		if strings.HasPrefix(args[i], "--") {
			key := strings.TrimPrefix(args[i], "--")
			if i+1 < len(args) && !strings.HasPrefix(args[i+1], "--") {
				result[key] = args[i+1]
				i++
			} else {
				result[key] = "true"
			}
		} else if i == 0 && !strings.HasPrefix(args[i], "--") {
			// First non-flag argument is considered the main argument
			result["main"] = args[i]
		}
	}

	return result
}
