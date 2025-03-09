package storage

import (
	"errors"
	"sync"
	"time"

	"github.com/mstgnz/cli-task-manager/models"
)

// MockStorage implements the Storage interface for testing
type MockStorage struct {
	tasks []models.Task
	mutex sync.RWMutex
}

// NewMockStorage creates a new MockStorage instance
func NewMockStorage() *MockStorage {
	return &MockStorage{
		tasks: []models.Task{},
	}
}

// GetTasks returns all tasks
func (s *MockStorage) GetTasks() ([]models.Task, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.tasks, nil
}

// AddTask adds a new task and returns the updated task with ID
func (s *MockStorage) AddTask(task models.Task) (models.Task, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Generate a new ID
	maxID := 0
	for _, t := range s.tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}
	task.ID = maxID + 1

	// Set timestamps if not already set
	if task.CreatedAt.IsZero() {
		now := time.Now()
		task.CreatedAt = now
		task.UpdatedAt = now
	}

	s.tasks = append(s.tasks, task)

	return task, nil
}

// UpdateTask updates an existing task
func (s *MockStorage) UpdateTask(task models.Task) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for i, t := range s.tasks {
		if t.ID == task.ID {
			// Update timestamp
			task.UpdatedAt = time.Now()
			s.tasks[i] = task
			return nil
		}
	}

	return errors.New("task not found")
}

// DeleteTask removes a task by ID
func (s *MockStorage) DeleteTask(id int) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for i, t := range s.tasks {
		if t.ID == id {
			// Remove task by replacing it with the last one and truncating
			s.tasks[i] = s.tasks[len(s.tasks)-1]
			s.tasks = s.tasks[:len(s.tasks)-1]
			return nil
		}
	}

	return errors.New("task not found")
}

// GetTaskByID retrieves a task by its ID
func (s *MockStorage) GetTaskByID(id int) (models.Task, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	for _, t := range s.tasks {
		if t.ID == id {
			return t, nil
		}
	}

	return models.Task{}, errors.New("task not found")
}
