package commands

import (
	"fmt"
	"strconv"
	"time"

	"github.com/mstgnz/cli-task-manager/models"
)

// handleAdd handles the add command
func (a *App) handleAdd(args []string) error {
	if len(args) == 0 {
		fmt.Println("Error: Title is required")
		return nil
	}

	parsedArgs := parseArgs(args)
	title := parsedArgs["main"]
	label := parsedArgs["label"]

	if label == "" {
		label = "task" // Default label
	}

	task := models.NewTask(title, label)

	addedTask, err := a.storage.AddTask(task)
	if err != nil {
		return fmt.Errorf("failed to add task: %w", err)
	}

	fmt.Printf("Task successfully added: %s\n", addedTask)
	return nil
}

// handleList handles the list command
func (a *App) handleList(args []string) error {
	tasks, err := a.storage.GetTasks()
	if err != nil {
		return fmt.Errorf("failed to get tasks: %w", err)
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return nil
	}

	fmt.Println("Tasks:")
	for _, task := range tasks {
		fmt.Println(task)
	}

	return nil
}

// handleUpdate handles the update command
func (a *App) handleUpdate(args []string) error {
	if len(args) == 0 {
		fmt.Println("Error: Task ID is required")
		return nil
	}

	parsedArgs := parseArgs(args)
	idStr := parsedArgs["main"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return fmt.Errorf("invalid task ID: %w", err)
	}

	task, err := a.storage.GetTaskByID(id)
	if err != nil {
		return fmt.Errorf("failed to get task: %w", err)
	}

	// Update status if provided
	if status, ok := parsedArgs["status"]; ok {
		switch models.Status(status) {
		case models.StatusTodo, models.StatusInProgress, models.StatusDone:
			task.Status = models.Status(status)
		default:
			fmt.Printf("Invalid status: %s. Using current status: %s\n", status, task.Status)
		}
	}

	// Update label if provided
	if label, ok := parsedArgs["label"]; ok {
		task.Label = label
	}

	// Update title if provided
	if title, ok := parsedArgs["title"]; ok {
		task.Title = title
	}

	// Update timestamp
	task.UpdatedAt = time.Now()

	if err := a.storage.UpdateTask(task); err != nil {
		return fmt.Errorf("failed to update task: %w", err)
	}

	fmt.Printf("Task successfully updated: %s\n", task)
	return nil
}

// handleFilter handles the filter command
func (a *App) handleFilter(args []string) error {
	parsedArgs := parseArgs(args)

	tasks, err := a.storage.GetTasks()
	if err != nil {
		return fmt.Errorf("failed to get tasks: %w", err)
	}

	var filteredTasks []models.Task

	// Filter by label
	if label, ok := parsedArgs["label"]; ok {
		for _, task := range tasks {
			if task.Label == label {
				filteredTasks = append(filteredTasks, task)
			}
		}
	}

	// Filter by status
	if status, ok := parsedArgs["status"]; ok {
		// If we already filtered by label, filter the filtered tasks
		if len(filteredTasks) > 0 {
			var statusFilteredTasks []models.Task
			for _, task := range filteredTasks {
				if string(task.Status) == status {
					statusFilteredTasks = append(statusFilteredTasks, task)
				}
			}
			filteredTasks = statusFilteredTasks
		} else {
			// Otherwise filter all tasks
			for _, task := range tasks {
				if string(task.Status) == status {
					filteredTasks = append(filteredTasks, task)
				}
			}
		}
	}

	// If no filters were applied, show all tasks
	if len(parsedArgs) == 0 {
		filteredTasks = tasks
	}

	if len(filteredTasks) == 0 {
		fmt.Println("No tasks found matching the filter criteria")
		return nil
	}

	fmt.Println("Filtered Tasks:")
	for _, task := range filteredTasks {
		fmt.Println(task)
	}

	return nil
}

// handleRemove handles the remove command
func (a *App) handleRemove(args []string) error {
	if len(args) == 0 {
		fmt.Println("Error: Task ID is required")
		return nil
	}

	parsedArgs := parseArgs(args)
	idStr := parsedArgs["main"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return fmt.Errorf("invalid task ID: %w", err)
	}

	if err := a.storage.DeleteTask(id); err != nil {
		return fmt.Errorf("failed to delete task: %w", err)
	}

	fmt.Printf("Task with ID %d successfully removed\n", id)
	return nil
}
