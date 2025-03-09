package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/mstgnz/cli-task-manager/models"
)

// JSONStorage implements the Storage interface using a JSON file
type JSONStorage struct {
	filePath string
	mutex    sync.RWMutex
}

// NewJSONStorage creates a new JSONStorage instance
func NewJSONStorage(filePath string) (*JSONStorage, error) {
	// Create directory if it doesn't exist
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create directory: %w", err)
	}

	// Create file if it doesn't exist
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		if err := os.WriteFile(filePath, []byte("[]"), 0644); err != nil {
			return nil, fmt.Errorf("failed to create file: %w", err)
		}
	}

	return &JSONStorage{
		filePath: filePath,
	}, nil
}

// GetTasks returns all tasks from the JSON file
func (s *JSONStorage) GetTasks() ([]models.Task, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	data, err := os.ReadFile(s.filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var tasks []models.Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, fmt.Errorf("failed to unmarshal tasks: %w", err)
	}

	return tasks, nil
}

// AddTask adds a new task to the JSON file
func (s *JSONStorage) AddTask(task models.Task) (models.Task, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	tasks, err := s.readTasks()
	if err != nil {
		return models.Task{}, err
	}

	// Generate a new ID
	maxID := 0
	for _, t := range tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}
	task.ID = maxID + 1

	tasks = append(tasks, task)

	if err := s.writeTasks(tasks); err != nil {
		return models.Task{}, err
	}

	return task, nil
}

// UpdateTask updates an existing task in the JSON file
func (s *JSONStorage) UpdateTask(task models.Task) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	tasks, err := s.readTasks()
	if err != nil {
		return err
	}

	found := false
	for i, t := range tasks {
		if t.ID == task.ID {
			tasks[i] = task
			found = true
			break
		}
	}

	if !found {
		return errors.New("task not found")
	}

	return s.writeTasks(tasks)
}

// DeleteTask removes a task by ID from the JSON file
func (s *JSONStorage) DeleteTask(id int) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	tasks, err := s.readTasks()
	if err != nil {
		return err
	}

	found := false
	var updatedTasks []models.Task
	for _, t := range tasks {
		if t.ID != id {
			updatedTasks = append(updatedTasks, t)
		} else {
			found = true
		}
	}

	if !found {
		return errors.New("task not found")
	}

	return s.writeTasks(updatedTasks)
}

// GetTaskByID retrieves a task by its ID
func (s *JSONStorage) GetTaskByID(id int) (models.Task, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	tasks, err := s.readTasks()
	if err != nil {
		return models.Task{}, err
	}

	for _, t := range tasks {
		if t.ID == id {
			return t, nil
		}
	}

	return models.Task{}, errors.New("task not found")
}

// readTasks reads all tasks from the JSON file
func (s *JSONStorage) readTasks() ([]models.Task, error) {
	data, err := os.ReadFile(s.filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var tasks []models.Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, fmt.Errorf("failed to unmarshal tasks: %w", err)
	}

	return tasks, nil
}

// writeTasks writes all tasks to the JSON file
func (s *JSONStorage) writeTasks(tasks []models.Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal tasks: %w", err)
	}

	if err := os.WriteFile(s.filePath, data, 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}
