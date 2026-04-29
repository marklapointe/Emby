package scheduled

import (
	"context"
	"sync"
	"time"

	"go.uber.org/zap"
)

// Task represents a scheduled task.
type Task struct {
	ID          string
	Name        string
	Description string
	Interval    time.Duration
	LastRun     time.Time
	NextRun     time.Time
	IsRunning   bool
	CancelFunc  context.CancelFunc
	fn          func(context.Context) error
}

// Manager manages scheduled tasks.
type Manager struct {
	mu       sync.RWMutex
	tasks    map[string]*Task
	logger   *zap.Logger
	running  bool
}

// NewManager creates a new scheduled task manager.
func NewManager(logger *zap.Logger) *Manager {
	return &Manager{
		tasks:  make(map[string]*Task),
		logger: logger,
	}
}

// RegisterTask registers a new scheduled task.
func (m *Manager) RegisterTask(id, name, description string, interval time.Duration, fn func(context.Context) error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	task := &Task{
		ID:          id,
		Name:        name,
		Description: description,
		Interval:    interval,
		LastRun:     time.Time{},
		NextRun:     time.Now().Add(interval),
		fn:          fn,
	}
	m.tasks[id] = task
	m.logger.Info("registered scheduled task", zap.String("id", id), zap.Duration("interval", interval))
}

// Start begins executing all registered tasks.
func (m *Manager) Start() {
	m.running = true
	m.logger.Info("starting scheduled task manager")

	go m.runLoop()
}

// Stop stops all scheduled tasks.
func (m *Manager) Stop() {
	m.mu.Lock()
	m.running = false
	for _, task := range m.tasks {
		if task.CancelFunc != nil {
			task.CancelFunc()
		}
	}
	m.mu.Unlock()
	m.logger.Info("stopped scheduled task manager")
}

// runLoop executes the task scheduling loop.
func (m *Manager) runLoop() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for m.running {
		select {
		case <-ticker.C:
			m.mu.RLock()
			for _, task := range m.tasks {
				if time.Now().After(task.NextRun) && !task.IsRunning {
					go m.executeTask(task)
				}
			}
			m.mu.RUnlock()
		}
	}
}

// executeTask runs a single task.
func (m *Manager) executeTask(task *Task) {
	m.mu.Lock()
	task.IsRunning = true
	m.mu.Unlock()

	m.logger.Info("executing scheduled task", zap.String("id", task.ID))

	ctx, cancel := context.WithCancel(context.Background())
	task.CancelFunc = cancel

	err := task.fn(ctx)

	m.mu.Lock()
	task.LastRun = time.Now()
	task.NextRun = time.Now().Add(task.Interval)
	task.IsRunning = false
	m.mu.Unlock()

	if err != nil {
		m.logger.Error("scheduled task failed", zap.String("id", task.ID), zap.Error(err))
	} else {
		m.logger.Info("scheduled task completed", zap.String("id", task.ID))
	}
}

// GetTasks returns all registered tasks.
func (m *Manager) GetTasks() []*Task {
	m.mu.RLock()
	defer m.mu.RUnlock()

	tasks := make([]*Task, 0, len(m.tasks))
	for _, task := range m.tasks {
		tasks = append(tasks, task)
	}
	return tasks
}

// GetTask returns a task by ID.
func (m *Manager) GetTask(id string) (*Task, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	task, exists := m.tasks[id]
	return task, exists
}
