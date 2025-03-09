package commands

import (
	"testing"

	"github.com/mstgnz/cli-task-manager/models"
	"github.com/mstgnz/cli-task-manager/storage"
)

func TestHandleAdd(t *testing.T) {
	// Create a mock app with mock storage
	app := &App{
		storage: storage.NewMockStorage(),
	}

	// Test adding a task with title and label
	err := app.handleAdd([]string{"Test Task", "--label", "feature"})
	if err != nil {
		t.Errorf("Expected no error when adding task, got %v", err)
	}

	// Check if the task was added
	tasks, err := app.storage.GetTasks()
	if err != nil {
		t.Fatalf("Failed to get tasks: %v", err)
	}

	if len(tasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(tasks))
	}

	if tasks[0].Title != "Test Task" {
		t.Errorf("Expected task title to be 'Test Task', got %s", tasks[0].Title)
	}

	if tasks[0].Label != "feature" {
		t.Errorf("Expected task label to be 'feature', got %s", tasks[0].Label)
	}

	// Test adding a task without a label (should use default)
	err = app.handleAdd([]string{"Another Task"})
	if err != nil {
		t.Errorf("Expected no error when adding task without label, got %v", err)
	}

	// Check if the task was added with default label
	tasks, err = app.storage.GetTasks()
	if err != nil {
		t.Fatalf("Failed to get tasks: %v", err)
	}

	if len(tasks) != 2 {
		t.Errorf("Expected 2 tasks, got %d", len(tasks))
	}

	// Find the second task
	var secondTask models.Task
	for _, task := range tasks {
		if task.Title == "Another Task" {
			secondTask = task
			break
		}
	}

	if secondTask.Title == "" {
		t.Fatal("Second task not found")
	}

	if secondTask.Label != "task" {
		t.Errorf("Expected default label to be 'task', got %s", secondTask.Label)
	}

	// Test adding a task without a title
	err = app.handleAdd([]string{})
	if err != nil {
		t.Errorf("Expected no error when adding task without title, got %v", err)
	}

	// Check that no new task was added
	tasks, err = app.storage.GetTasks()
	if err != nil {
		t.Fatalf("Failed to get tasks: %v", err)
	}

	if len(tasks) != 2 {
		t.Errorf("Expected still 2 tasks, got %d", len(tasks))
	}
}

func TestHandleList(t *testing.T) {
	// Create a mock app with mock storage
	app := &App{
		storage: storage.NewMockStorage(),
	}

	// Test listing tasks when there are none
	err := app.handleList([]string{})
	if err != nil {
		t.Errorf("Expected no error when listing empty tasks, got %v", err)
	}

	// Add a task
	task := models.Task{
		Title:  "Test Task",
		Label:  "test",
		Status: models.StatusTodo,
	}

	_, err = app.storage.AddTask(task)
	if err != nil {
		t.Fatalf("Failed to add task: %v", err)
	}

	// Test listing tasks when there is one
	err = app.handleList([]string{})
	if err != nil {
		t.Errorf("Expected no error when listing tasks, got %v", err)
	}
}

func TestHandleUpdate(t *testing.T) {
	// Create a mock app with mock storage
	app := &App{
		storage: storage.NewMockStorage(),
	}

	// Add a task
	task := models.Task{
		Title:  "Test Task",
		Label:  "test",
		Status: models.StatusTodo,
	}

	addedTask, err := app.storage.AddTask(task)
	if err != nil {
		t.Fatalf("Failed to add task: %v", err)
	}

	// Test updating task status
	err = app.handleUpdate([]string{
		"1", // Task ID
		"--status", "in-progress",
	})
	if err != nil {
		t.Errorf("Expected no error when updating task status, got %v", err)
	}

	// Check if the task was updated
	updatedTask, err := app.storage.GetTaskByID(addedTask.ID)
	if err != nil {
		t.Fatalf("Failed to get updated task: %v", err)
	}

	if updatedTask.Status != models.StatusInProgress {
		t.Errorf("Expected status to be %s, got %s", models.StatusInProgress, updatedTask.Status)
	}

	// Test updating task label
	err = app.handleUpdate([]string{
		"1", // Task ID
		"--label", "feature",
	})
	if err != nil {
		t.Errorf("Expected no error when updating task label, got %v", err)
	}

	// Check if the task was updated
	updatedTask, err = app.storage.GetTaskByID(addedTask.ID)
	if err != nil {
		t.Fatalf("Failed to get updated task: %v", err)
	}

	if updatedTask.Label != "feature" {
		t.Errorf("Expected label to be 'feature', got %s", updatedTask.Label)
	}

	// Test updating task title
	err = app.handleUpdate([]string{
		"1", // Task ID
		"--title", "Updated Task",
	})
	if err != nil {
		t.Errorf("Expected no error when updating task title, got %v", err)
	}

	// Check if the task was updated
	updatedTask, err = app.storage.GetTaskByID(addedTask.ID)
	if err != nil {
		t.Fatalf("Failed to get updated task: %v", err)
	}

	if updatedTask.Title != "Updated Task" {
		t.Errorf("Expected title to be 'Updated Task', got %s", updatedTask.Title)
	}

	// Test updating non-existent task
	err = app.handleUpdate([]string{
		"999", // Non-existent task ID
		"--status", "done",
	})
	if err == nil {
		t.Error("Expected error when updating non-existent task, got nil")
	}

	// Test updating with invalid task ID
	err = app.handleUpdate([]string{
		"invalid", // Invalid task ID
		"--status", "done",
	})
	if err == nil {
		t.Error("Expected error when updating with invalid task ID, got nil")
	}

	// Test updating with no task ID
	err = app.handleUpdate([]string{})
	if err != nil {
		t.Errorf("Expected no error when updating with no task ID, got %v", err)
	}
}

func TestHandleFilter(t *testing.T) {
	// Create a mock app with mock storage
	app := &App{
		storage: storage.NewMockStorage(),
	}

	// Add tasks with different labels and statuses
	task1 := models.Task{
		Title:  "Task 1",
		Label:  "feature",
		Status: models.StatusTodo,
	}

	task2 := models.Task{
		Title:  "Task 2",
		Label:  "bug",
		Status: models.StatusInProgress,
	}

	task3 := models.Task{
		Title:  "Task 3",
		Label:  "feature",
		Status: models.StatusDone,
	}

	_, err := app.storage.AddTask(task1)
	if err != nil {
		t.Fatalf("Failed to add task 1: %v", err)
	}

	_, err = app.storage.AddTask(task2)
	if err != nil {
		t.Fatalf("Failed to add task 2: %v", err)
	}

	_, err = app.storage.AddTask(task3)
	if err != nil {
		t.Fatalf("Failed to add task 3: %v", err)
	}

	// Test filtering by label
	err = app.handleFilter([]string{"--label", "feature"})
	if err != nil {
		t.Errorf("Expected no error when filtering by label, got %v", err)
	}

	// Test filtering by status
	err = app.handleFilter([]string{"--status", "in-progress"})
	if err != nil {
		t.Errorf("Expected no error when filtering by status, got %v", err)
	}

	// Test filtering by label and status
	err = app.handleFilter([]string{"--label", "feature", "--status", "done"})
	if err != nil {
		t.Errorf("Expected no error when filtering by label and status, got %v", err)
	}

	// Test filtering with no matches
	err = app.handleFilter([]string{"--label", "nonexistent"})
	if err != nil {
		t.Errorf("Expected no error when filtering with no matches, got %v", err)
	}

	// Test filtering with no filters (should show all tasks)
	err = app.handleFilter([]string{})
	if err != nil {
		t.Errorf("Expected no error when filtering with no filters, got %v", err)
	}
}

func TestHandleRemove(t *testing.T) {
	// Create a mock app with mock storage
	app := &App{
		storage: storage.NewMockStorage(),
	}

	// Add a task
	task := models.Task{
		Title:  "Test Task",
		Label:  "test",
		Status: models.StatusTodo,
	}

	addedTask, err := app.storage.AddTask(task)
	if err != nil {
		t.Fatalf("Failed to add task: %v", err)
	}

	// Test removing the task
	err = app.handleRemove([]string{
		"1", // Task ID
	})
	if err != nil {
		t.Errorf("Expected no error when removing task, got %v", err)
	}

	// Check if the task was removed
	_, err = app.storage.GetTaskByID(addedTask.ID)
	if err == nil {
		t.Error("Expected error when getting removed task, got nil")
	}

	// Test removing non-existent task
	err = app.handleRemove([]string{
		"999", // Non-existent task ID
	})
	if err == nil {
		t.Error("Expected error when removing non-existent task, got nil")
	}

	// Test removing with invalid task ID
	err = app.handleRemove([]string{
		"invalid", // Invalid task ID
	})
	if err == nil {
		t.Error("Expected error when removing with invalid task ID, got nil")
	}

	// Test removing with no task ID
	err = app.handleRemove([]string{})
	if err != nil {
		t.Errorf("Expected no error when removing with no task ID, got %v", err)
	}
}
