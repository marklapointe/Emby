package scheduled

import (
	"fmt"
	"sync"
	"time"

	"go.uber.org/zap"
)

// Task represents a scheduled task.
type Task struct {
	ID              string    `json:"Id"`
	Name            string    `json:"Name"`
	Description     string    `json:"Description"`
	Category        string    `json:"Category"`
	Icon            string    `json:"Icon,omitempty"`
	CommandLine     string    `json:"CommandLine,omitempty"`
	ExecutionTimeSpec string  `json:"ExecutionTimeSpec,omitempty"`
	LastExecutionTime time.Time `json:"LastExecutionTime,omitempty"`
	NextExecutionTime time.Time `json:"NextExecutionTime,omitempty"`
	IsRunning       bool      `json:"IsRunning"`
	ProgressPercent int       `json:"ProgressPercent"`
	Options         TaskOptions `json:"Options,omitempty"`
}

// TaskOptions represents task execution options.
type TaskOptions struct {
	TriggerOnce      bool `json:"TriggerOnce"`
	EnableOnStartup  bool `json:"EnableOnStartup"`
	EnableWhenEmpty  bool `json:"EnableWhenEmpty"`
}

// Manager handles scheduled task operations.
type Manager struct {
	mu       sync.RWMutex
	tasks    map[string]*Task
	logger   *zap.Logger
}

// NewManager creates a new scheduled task manager.
func NewManager(logger *zap.Logger) *Manager {
	return &Manager{
		tasks:  make(map[string]*Task),
		logger: logger,
	}
}

// RegisterTask registers a new scheduled task.
func (m *Manager) RegisterTask(task *Task) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.tasks[task.ID]; exists {
		return fmt.Errorf("task already registered: %s", task.ID)
	}

	m.tasks[task.ID] = task
	if m.logger != nil {
		m.logger.Info("task registered", zap.String("id", task.ID), zap.String("name", task.Name))
	}
	return nil
}

// GetTask returns a task by ID.
func (m *Manager) GetTask(id string) (*Task, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	task, exists := m.tasks[id]
	return task, exists
}

// GetAllTasks returns all registered tasks.
func (m *Manager) GetAllTasks() []*Task {
	m.mu.RLock()
	defer m.mu.RUnlock()

	tasks := make([]*Task, 0, len(m.tasks))
	for _, task := range m.tasks {
		tasks = append(tasks, task)
	}
	return tasks
}

// GetRunningTasks returns all running tasks.
func (m *Manager) GetRunningTasks() []*Task {
	m.mu.RLock()
	defer m.mu.RUnlock()

	var tasks []*Task
	for _, task := range m.tasks {
		if task.IsRunning {
			tasks = append(tasks, task)
		}
	}
	return tasks
}

// GetTasksByCategory returns tasks filtered by category.
func (m *Manager) GetTasksByCategory(category string) []*Task {
	m.mu.RLock()
	defer m.mu.RUnlock()

	var tasks []*Task
	for _, task := range m.tasks {
		if category == "" || task.Category == category {
			tasks = append(tasks, task)
		}
	}
	return tasks
}

// ExecuteTask executes a task immediately.
func (m *Manager) ExecuteTask(id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	task, exists := m.tasks[id]
	if !exists {
		return fmt.Errorf("task not found: %s", id)
	}

	if task.IsRunning {
		return fmt.Errorf("task already running: %s", id)
	}

	task.IsRunning = true
	task.LastExecutionTime = time.Now()
	task.ProgressPercent = 0

	if m.logger != nil {
		m.logger.Info("task executed", zap.String("id", id), zap.String("name", task.Name))
	}

	return nil
}

// CancelTask cancels a running task.
func (m *Manager) CancelTask(id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	task, exists := m.tasks[id]
	if !exists {
		return fmt.Errorf("task not found: %s", id)
	}

	task.IsRunning = false
	task.ProgressPercent = 0

	if m.logger != nil {
		m.logger.Info("task cancelled", zap.String("id", id), zap.String("name", task.Name))
	}

	return nil
}

// UpdateTaskProgress updates a task's progress.
func (m *Manager) UpdateTaskProgress(id string, progress int) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	task, exists := m.tasks[id]
	if !exists {
		return fmt.Errorf("task not found: %s", id)
	}

	task.ProgressPercent = progress
	return nil
}

// CompleteTask marks a task as completed.
func (m *Manager) CompleteTask(id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	task, exists := m.tasks[id]
	if !exists {
		return fmt.Errorf("task not found: %s", id)
	}

	task.IsRunning = false
	task.ProgressPercent = 100

	if m.logger != nil {
		m.logger.Info("task completed", zap.String("id", id), zap.String("name", task.Name))
	}

	return nil
}

// GetTaskCount returns the total number of tasks.
func (m *Manager) GetTaskCount() int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return len(m.tasks)
}

// GetRunningTaskCount returns the number of running tasks.
func (m *Manager) GetRunningTaskCount() int {
	m.mu.RLock()
	defer m.mu.RUnlock()

	count := 0
	for _, task := range m.tasks {
		if task.IsRunning {
			count++
		}
	}
	return count
}

// GetTasksByStatus returns tasks filtered by running status.
func (m *Manager) GetTasksByStatus(running bool) []*Task {
	m.mu.RLock()
	defer m.mu.RUnlock()

	var tasks []*Task
	for _, task := range m.tasks {
		if task.IsRunning == running {
			tasks = append(tasks, task)
		}
	}
	return tasks
}
