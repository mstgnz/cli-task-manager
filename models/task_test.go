package models

import (
	"testing"
	"time"
)

func TestNewTask(t *testing.T) {
	title := "Test Task"
	label := "test"

	task := NewTask(title, label)

	if task.Title != title {
		t.Errorf("Expected title to be %s, got %s", title, task.Title)
	}

	if task.Label != label {
		t.Errorf("Expected label to be %s, got %s", label, task.Label)
	}

	if task.Status != StatusTodo {
		t.Errorf("Expected status to be %s, got %s", StatusTodo, task.Status)
	}

	// Check if created_at and updated_at are set to approximately now
	now := time.Now()
	if task.CreatedAt.Sub(now).Abs() > time.Second {
		t.Errorf("Expected CreatedAt to be close to now, got %v", task.CreatedAt)
	}

	if task.UpdatedAt.Sub(now).Abs() > time.Second {
		t.Errorf("Expected UpdatedAt to be close to now, got %v", task.UpdatedAt)
	}
}

func TestTaskString(t *testing.T) {
	task := Task{
		ID:     1,
		Title:  "Test Task",
		Status: StatusInProgress,
		Label:  "feature",
	}

	expected := "1. [feature] Test Task [Status: in-progress]"
	if task.String() != expected {
		t.Errorf("Expected string representation to be %s, got %s", expected, task.String())
	}
}
