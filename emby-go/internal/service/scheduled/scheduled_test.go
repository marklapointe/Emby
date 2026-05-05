package scheduled

import (
	"testing"
)

func TestNewManager(t *testing.T) {
	m := NewManager(nil)
	if m == nil {
		t.Fatal("NewManager returned nil")
	}
	if m.tasks == nil {
		t.Error("tasks map not initialized")
	}
}

func TestRegisterTask(t *testing.T) {
	m := NewManager(nil)

	task := &Task{
		ID:   "test-task",
		Name: "Test Task",
	}

	err := m.RegisterTask(task)
	if err != nil {
		t.Errorf("RegisterTask returned error: %v", err)
	}

	if m.tasks["test-task"] == nil {
		t.Error("task not registered")
	}
}

func TestRegisterTask_Duplicate(t *testing.T) {
	m := NewManager(nil)

	task := &Task{ID: "test-task", Name: "Test Task"}
	m.RegisterTask(task)

	err := m.RegisterTask(task)
	if err == nil {
		t.Error("expected error for duplicate task registration")
	}
}

func TestGetTask(t *testing.T) {
	m := NewManager(nil)

	task := &Task{ID: "test-task", Name: "Test Task"}
	m.RegisterTask(task)

	got, ok := m.GetTask("test-task")
	if !ok {
		t.Error("GetTask returned false for existing task")
	}
	if got.Name != "Test Task" {
		t.Errorf("expected name 'Test Task', got '%s'", got.Name)
	}
}

func TestGetTask_NotFound(t *testing.T) {
	m := NewManager(nil)
	_, ok := m.GetTask("non-existent")
	if ok {
		t.Error("GetTask returned true for non-existent task")
	}
}

func TestGetAllTasks(t *testing.T) {
	m := NewManager(nil)

	m.RegisterTask(&Task{ID: "task-1", Name: "Task 1"})
	m.RegisterTask(&Task{ID: "task-2", Name: "Task 2"})

	tasks := m.GetAllTasks()
	if len(tasks) != 2 {
		t.Errorf("expected 2 tasks, got %d", len(tasks))
	}
}

func TestGetRunningTasks(t *testing.T) {
	m := NewManager(nil)

	m.RegisterTask(&Task{ID: "task-1", Name: "Task 1", IsRunning: true})
	m.RegisterTask(&Task{ID: "task-2", Name: "Task 2", IsRunning: false})

	tasks := m.GetRunningTasks()
	if len(tasks) != 1 {
		t.Errorf("expected 1 running task, got %d", len(tasks))
	}
}

func TestExecuteTask(t *testing.T) {
	m := NewManager(nil)

	task := &Task{ID: "test-task", Name: "Test Task", IsRunning: false}
	m.RegisterTask(task)

	err := m.ExecuteTask("test-task")
	if err != nil {
		t.Errorf("ExecuteTask returned error: %v", err)
	}

	if !m.tasks["test-task"].IsRunning {
		t.Error("task.IsRunning should be true after ExecuteTask")
	}
}

func TestExecuteTask_NotFound(t *testing.T) {
	m := NewManager(nil)
	err := m.ExecuteTask("non-existent")
	if err == nil {
		t.Error("expected error for non-existent task")
	}
}

func TestExecuteTask_AlreadyRunning(t *testing.T) {
	m := NewManager(nil)

	task := &Task{ID: "test-task", Name: "Test Task", IsRunning: true}
	m.RegisterTask(task)

	err := m.ExecuteTask("test-task")
	if err == nil {
		t.Error("expected error for already running task")
	}
}

func TestCancelTask(t *testing.T) {
	m := NewManager(nil)

	task := &Task{ID: "test-task", Name: "Test Task", IsRunning: true}
	m.RegisterTask(task)

	err := m.CancelTask("test-task")
	if err != nil {
		t.Errorf("CancelTask returned error: %v", err)
	}

	if m.tasks["test-task"].IsRunning {
		t.Error("task.IsRunning should be false after CancelTask")
	}
}

func TestCompleteTask(t *testing.T) {
	m := NewManager(nil)

	task := &Task{ID: "test-task", Name: "Test Task", IsRunning: true}
	m.RegisterTask(task)

	err := m.CompleteTask("test-task")
	if err != nil {
		t.Errorf("CompleteTask returned error: %v", err)
	}

	if m.tasks["test-task"].IsRunning {
		t.Error("task.IsRunning should be false after CompleteTask")
	}
	if m.tasks["test-task"].ProgressPercent != 100 {
		t.Error("task.ProgressPercent should be 100 after CompleteTask")
	}
}

func TestUpdateTaskProgress(t *testing.T) {
	m := NewManager(nil)

	task := &Task{ID: "test-task", Name: "Test Task"}
	m.RegisterTask(task)

	err := m.UpdateTaskProgress("test-task", 50)
	if err != nil {
		t.Errorf("UpdateTaskProgress returned error: %v", err)
	}

	if m.tasks["test-task"].ProgressPercent != 50 {
		t.Errorf("expected progress 50, got %d", m.tasks["test-task"].ProgressPercent)
	}
}

func TestGetTaskCount(t *testing.T) {
	m := NewManager(nil)

	m.RegisterTask(&Task{ID: "task-1", Name: "Task 1"})
	m.RegisterTask(&Task{ID: "task-2", Name: "Task 2"})

	count := m.GetTaskCount()
	if count != 2 {
		t.Errorf("expected 2, got %d", count)
	}
}

func TestGetRunningTaskCount(t *testing.T) {
	m := NewManager(nil)

	m.RegisterTask(&Task{ID: "task-1", IsRunning: true})
	m.RegisterTask(&Task{ID: "task-2", IsRunning: false})
	m.RegisterTask(&Task{ID: "task-3", IsRunning: true})

	count := m.GetRunningTaskCount()
	if count != 2 {
		t.Errorf("expected 2, got %d", count)
	}
}

func TestGetTasksByCategory(t *testing.T) {
	m := NewManager(nil)

	m.RegisterTask(&Task{ID: "task-1", Category: "Library"})
	m.RegisterTask(&Task{ID: "task-2", Category: "System"})
	m.RegisterTask(&Task{ID: "task-3", Category: "Library"})

	tasks := m.GetTasksByCategory("Library")
	if len(tasks) != 2 {
		t.Errorf("expected 2 Library tasks, got %d", len(tasks))
	}
}

func TestGetTasksByStatus(t *testing.T) {
	m := NewManager(nil)

	m.RegisterTask(&Task{ID: "task-1", IsRunning: true})
	m.RegisterTask(&Task{ID: "task-2", IsRunning: false})

	running := m.GetTasksByStatus(true)
	if len(running) != 1 {
		t.Errorf("expected 1 running task, got %d", len(running))
	}
}