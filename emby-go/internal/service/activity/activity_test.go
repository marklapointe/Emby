package activity

import (
	"testing"
)

func TestNewManager(t *testing.T) {
	m := NewManager()
	if m == nil {
		t.Fatal("NewManager returned nil")
	}
	if m.activities == nil {
		t.Error("activities map not initialized")
	}
}

func TestGetActivities_Empty(t *testing.T) {
	m := NewManager()
	activities := m.GetActivities()
	if len(activities) != 0 {
		t.Errorf("expected 0 activities, got %d", len(activities))
	}
}

func TestLogActivity(t *testing.T) {
	m := NewManager()

	activity := m.LogActivity("Test Activity", "Test Overview", "user-1", "TestUser", "test", "info")

	if activity == nil {
		t.Fatal("LogActivity returned nil")
	}
	if activity.ID == "" {
		t.Error("activity ID is empty")
	}
	if activity.Name != "Test Activity" {
		t.Errorf("expected name 'Test Activity', got '%s'", activity.Name)
	}
	if activity.Overview != "Test Overview" {
		t.Errorf("expected overview 'Test Overview', got '%s'", activity.Overview)
	}
	if activity.UserID != "user-1" {
		t.Errorf("expected userID 'user-1', got '%s'", activity.UserID)
	}
	if activity.UserName != "TestUser" {
		t.Errorf("expected userName 'TestUser', got '%s'", activity.UserName)
	}
	if activity.Type != "test" {
		t.Errorf("expected type 'test', got '%s'", activity.Type)
	}
	if activity.Severity != "info" {
		t.Errorf("expected severity 'info', got '%s'", activity.Severity)
	}
}

func TestGetActivity(t *testing.T) {
	m := NewManager()

	activity := m.LogActivity("Test", "Overview", "user-1", "User", "type", "info")
	got, ok := m.GetActivity(activity.ID)

	if !ok {
		t.Error("GetActivity returned false for existing activity")
	}
	if got.ID != activity.ID {
		t.Errorf("expected ID '%s', got '%s'", activity.ID, got.ID)
	}
}

func TestGetActivity_NotFound(t *testing.T) {
	m := NewManager()
	_, ok := m.GetActivity("non-existent")
	if ok {
		t.Error("GetActivity returned true for non-existent activity")
	}
}

func TestGetActivities_Multiple(t *testing.T) {
	m := NewManager()

	m.LogActivity("Activity 1", "Overview 1", "user-1", "User1", "type1", "info")
	m.LogActivity("Activity 2", "Overview 2", "user-2", "User2", "type2", "warning")

	activities := m.GetActivities()
	if len(activities) != 2 {
		t.Errorf("expected 2 activities, got %d", len(activities))
	}
}

func TestDeleteActivity(t *testing.T) {
	m := NewManager()

	activity := m.LogActivity("Test", "Overview", "user-1", "User", "type", "info")
	err := m.DeleteActivity(activity.ID)

	if err != nil {
		t.Errorf("DeleteActivity returned error: %v", err)
	}

	_, ok := m.GetActivity(activity.ID)
	if ok {
		t.Error("activity still exists after deletion")
	}
}

func TestDeleteActivity_NotFound(t *testing.T) {
	m := NewManager()
	err := m.DeleteActivity("non-existent")
	if err != nil {
		t.Errorf("DeleteActivity returned error for non-existent: %v", err)
	}
}