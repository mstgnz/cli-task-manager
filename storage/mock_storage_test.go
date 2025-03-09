package storage

import (
	"testing"
	"time"

	"github.com/mstgnz/cli-task-manager/models"
)

func TestMockStorage_AddTask(t *testing.T) {
	storage := NewMockStorage()

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

	if addedTask.Title != task.Title {
		t.Errorf("Expected task title to be %s, got %s", task.Title, addedTask.Title)
	}

	tasks, err := storage.GetTasks()
	if err != nil {
		t.Fatalf("Failed to get tasks: %v", err)
	}

	if len(tasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(tasks))
	}
}

func TestMockStorage_GetTaskByID(t *testing.T) {
	storage := NewMockStorage()

	// Add a task
	task := models.Task{
		Title:  "Test Task",
		Label:  "test",
		Status: models.StatusTodo,
	}

	addedTask, err := storage.AddTask(task)
	if err != nil {
		t.Fatalf("Failed to add task: %v", err)
	}

	// Get the task by ID
	retrievedTask, err := storage.GetTaskByID(addedTask.ID)
	if err != nil {
		t.Fatalf("Failed to get task by ID: %v", err)
	}

	if retrievedTask.ID != addedTask.ID {
		t.Errorf("Expected task ID to be %d, got %d", addedTask.ID, retrievedTask.ID)
	}

	// Try to get a non-existent task
	_, err = storage.GetTaskByID(999)
	if err == nil {
		t.Error("Expected error when getting non-existent task, got nil")
	}
}

func TestMockStorage_UpdateTask(t *testing.T) {
	storage := NewMockStorage()

	// Add a task
	task := models.Task{
		Title:  "Test Task",
		Label:  "test",
		Status: models.StatusTodo,
	}

	addedTask, err := storage.AddTask(task)
	if err != nil {
		t.Fatalf("Failed to add task: %v", err)
	}

	// Update the task
	addedTask.Title = "Updated Task"
	addedTask.Status = models.StatusInProgress

	// Store the current update time for comparison
	beforeUpdate := time.Now()

	err = storage.UpdateTask(addedTask)
	if err != nil {
		t.Fatalf("Failed to update task: %v", err)
	}

	// Get the updated task
	updatedTask, err := storage.GetTaskByID(addedTask.ID)
	if err != nil {
		t.Fatalf("Failed to get updated task: %v", err)
	}

	if updatedTask.Title != "Updated Task" {
		t.Errorf("Expected updated title to be 'Updated Task', got %s", updatedTask.Title)
	}

	if updatedTask.Status != models.StatusInProgress {
		t.Errorf("Expected updated status to be %s, got %s", models.StatusInProgress, updatedTask.Status)
	}

	// Check if UpdatedAt was updated
	if !updatedTask.UpdatedAt.After(beforeUpdate) && !updatedTask.UpdatedAt.Equal(beforeUpdate) {
		t.Errorf("Expected UpdatedAt to be after or equal to %v, got %v", beforeUpdate, updatedTask.UpdatedAt)
	}

	// Try to update a non-existent task
	nonExistentTask := models.Task{
		ID:     999,
		Title:  "Non-existent Task",
		Status: models.StatusTodo,
	}

	err = storage.UpdateTask(nonExistentTask)
	if err == nil {
		t.Error("Expected error when updating non-existent task, got nil")
	}
}

func TestMockStorage_DeleteTask(t *testing.T) {
	storage := NewMockStorage()

	// Add tasks
	task1 := models.Task{
		Title:  "Task 1",
		Label:  "test",
		Status: models.StatusTodo,
	}

	task2 := models.Task{
		Title:  "Task 2",
		Label:  "test",
		Status: models.StatusTodo,
	}

	addedTask1, err := storage.AddTask(task1)
	if err != nil {
		t.Fatalf("Failed to add task 1: %v", err)
	}

	_, err = storage.AddTask(task2)
	if err != nil {
		t.Fatalf("Failed to add task 2: %v", err)
	}

	// Delete the first task
	err = storage.DeleteTask(addedTask1.ID)
	if err != nil {
		t.Fatalf("Failed to delete task: %v", err)
	}

	// Check if the task was deleted
	tasks, err := storage.GetTasks()
	if err != nil {
		t.Fatalf("Failed to get tasks: %v", err)
	}

	if len(tasks) != 1 {
		t.Errorf("Expected 1 task after deletion, got %d", len(tasks))
	}

	// Try to get the deleted task
	_, err = storage.GetTaskByID(addedTask1.ID)
	if err == nil {
		t.Error("Expected error when getting deleted task, got nil")
	}

	// Try to delete a non-existent task
	err = storage.DeleteTask(999)
	if err == nil {
		t.Error("Expected error when deleting non-existent task, got nil")
	}
}
