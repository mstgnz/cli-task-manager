package models

import (
	"fmt"
	"time"
)

// Status represents the current state of a task
type Status string

const (
	StatusTodo       Status = "to-do"
	StatusInProgress Status = "in-progress"
	StatusDone       Status = "done"
)

// Task represents a single task in the task manager
type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	Status      Status    `json:"status"`
	Label       string    `json:"label"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// String returns a formatted string representation of the task
func (t Task) String() string {
	return fmt.Sprintf("%d. [%s] %s [Status: %s]", t.ID, t.Label, t.Title, t.Status)
}

// NewTask creates a new task with the given title and label
func NewTask(title, label string) Task {
	now := time.Now()
	return Task{
		Title:     title,
		Label:     label,
		Status:    StatusTodo,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
