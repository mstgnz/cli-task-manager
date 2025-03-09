package storage

import (
	"github.com/mstgnz/cli-task-manager/models"
)

// Storage defines the interface for task storage operations
type Storage interface {
	// GetTasks returns all tasks
	GetTasks() ([]models.Task, error)

	// AddTask adds a new task and returns the updated task with ID
	AddTask(task models.Task) (models.Task, error)

	// UpdateTask updates an existing task
	UpdateTask(task models.Task) error

	// DeleteTask removes a task by ID
	DeleteTask(id int) error

	// GetTaskByID retrieves a task by its ID
	GetTaskByID(id int) (models.Task, error)
}
