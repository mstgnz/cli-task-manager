package storage

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/mstgnz/cli-task-manager/models"
)

func TestJSONStorage(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "json-storage-test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir) // Clean up after test

	// Create a JSON storage with a file in the temp directory
	filePath := filepath.Join(tempDir, "tasks.json")
	storage, err := NewJSONStorage(filePath)
	if err != nil {
		t.Fatalf("Failed to create JSON storage: %v", err)
	}

	// Test adding a task
	task := models.Task{
		Title:  "Test Task",
		Label:  "test",
		Status: models.StatusTodo,
	}

	addedTask, err := storage.AddTask(task)
	if err != nil {
		t.Fatalf("Failed to add task: %v", err)
	}

	if addedTask.ID != 1 {
		t.Errorf("Expected task ID to be 1, got %d", addedTask.ID)
	}

	// Test getting all tasks
	tasks, err := storage.GetTasks()
	if err != nil {
		t.Fatalf("Failed to get tasks: %v", err)
	}

	if len(tasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(tasks))
	}

	// Test getting a task by ID
	retrievedTask, err := storage.GetTaskByID(addedTask.ID)
	if err != nil {
		t.Fatalf("Failed to get task by ID: %v", err)
	}

	if retrievedTask.ID != addedTask.ID {
		t.Errorf("Expected task ID to be %d, got %d", addedTask.ID, retrievedTask.ID)
	}

	// Test updating a task
	retrievedTask.Title = "Updated Task"
	retrievedTask.Status = models.StatusInProgress

	err = storage.UpdateTask(retrievedTask)
	if err != nil {
		t.Fatalf("Failed to update task: %v", err)
	}

	// Get the updated task
	updatedTask, err := storage.GetTaskByID(retrievedTask.ID)
	if err != nil {
		t.Fatalf("Failed to get updated task: %v", err)
	}

	if updatedTask.Title != "Updated Task" {
		t.Errorf("Expected updated title to be 'Updated Task', got %s", updatedTask.Title)
	}

	if updatedTask.Status != models.StatusInProgress {
		t.Errorf("Expected updated status to be %s, got %s", models.StatusInProgress, updatedTask.Status)
	}

	// Test deleting a task
	err = storage.DeleteTask(updatedTask.ID)
	if err != nil {
		t.Fatalf("Failed to delete task: %v", err)
	}

	// Check if the task was deleted
	tasks, err = storage.GetTasks()
	if err != nil {
		t.Fatalf("Failed to get tasks: %v", err)
	}

	if len(tasks) != 0 {
		t.Errorf("Expected 0 tasks after deletion, got %d", len(tasks))
	}

	// Try to get the deleted task
	_, err = storage.GetTaskByID(updatedTask.ID)
	if err == nil {
		t.Error("Expected error when getting deleted task, got nil")
	}

	// Test file persistence by creating a new storage instance
	newStorage, err := NewJSONStorage(filePath)
	if err != nil {
		t.Fatalf("Failed to create new JSON storage: %v", err)
	}

	// Check if the tasks are still empty (since we deleted the only task)
	tasks, err = newStorage.GetTasks()
	if err != nil {
		t.Fatalf("Failed to get tasks from new storage: %v", err)
	}

	if len(tasks) != 0 {
		t.Errorf("Expected 0 tasks in new storage, got %d", len(tasks))
	}
}
