package livetv

import (
	"testing"
)

func TestNewManager(t *testing.T) {
	m := NewManager(nil)
	if m == nil {
		t.Fatal("NewManager returned nil")
	}
	if m.isEnabled {
		t.Error("expected isEnabled to be false initially")
	}
}

func TestIsEnabled(t *testing.T) {
	m := NewManager(nil)
	if m.IsEnabled() {
		t.Error("expected IsEnabled to be false initially")
	}
}

func TestEnableDisable(t *testing.T) {
	m := NewManager(nil)

	m.Enable()
	if !m.IsEnabled() {
		t.Error("expected IsEnabled to be true after Enable")
	}

	m.Disable()
	if m.IsEnabled() {
		t.Error("expected IsEnabled to be false after Disable")
	}
}

func TestAddChannel(t *testing.T) {
	m := NewManager(nil)

	ch := &Channel{ID: "ch-1", Name: "Test Channel"}
	m.AddChannel(ch)

	channels := m.GetChannels()
	if len(channels) != 1 {
		t.Errorf("expected 1 channel, got %d", len(channels))
	}
}

func TestGetChannel(t *testing.T) {
	m := NewManager(nil)

	ch := &Channel{ID: "ch-1", Name: "Test"}
	m.AddChannel(ch)

	got, ok := m.GetChannel("ch-1")
	if !ok {
		t.Error("GetChannel returned false for existing channel")
	}
	if got.Name != "Test" {
		t.Errorf("expected name 'Test', got '%s'", got.Name)
	}
}

func TestGetChannels_Empty(t *testing.T) {
	m := NewManager(nil)
	channels := m.GetChannels()
	if len(channels) != 0 {
		t.Errorf("expected 0 channels, got %d", len(channels))
	}
}

func TestGetPrograms_Empty(t *testing.T) {
	m := NewManager(nil)
	programs := m.GetPrograms("")
	if len(programs) != 0 {
		t.Errorf("expected 0 programs, got %d", len(programs))
	}
}

func TestGetRecordings_Empty(t *testing.T) {
	m := NewManager(nil)
	recordings := m.GetRecordings()
	if len(recordings) != 0 {
		t.Errorf("expected 0 recordings, got %d", len(recordings))
	}
}

func TestGetTimers_Empty(t *testing.T) {
	m := NewManager(nil)
	timers := m.GetTimers()
	if len(timers) != 0 {
		t.Errorf("expected 0 timers, got %d", len(timers))
	}
}

func TestAddRecording(t *testing.T) {
	m := NewManager(nil)

	rec := &Recording{ID: "rec-1", Name: "Test Recording"}
	m.AddRecording(rec)

	recordings := m.GetRecordings()
	if len(recordings) != 1 {
		t.Errorf("expected 1 recording, got %d", len(recordings))
	}
}

func TestDeleteRecording(t *testing.T) {
	m := NewManager(nil)

	m.AddRecording(&Recording{ID: "rec-1", Name: "Test"})
	err := m.DeleteRecording("rec-1")
	if err != nil {
		t.Errorf("DeleteRecording returned error: %v", err)
	}

	recordings := m.GetRecordings()
	if len(recordings) != 0 {
		t.Errorf("expected 0 recordings after delete, got %d", len(recordings))
	}
}

func TestDeleteRecording_NotFound(t *testing.T) {
	m := NewManager(nil)
	err := m.DeleteRecording("non-existent")
	if err != nil {
		t.Errorf("DeleteRecording returned error: %v", err)
	}
}

func TestAddTimer(t *testing.T) {
	m := NewManager(nil)

	timer := &Timer{ID: "timer-1", Name: "Test Timer"}
	m.AddTimer(timer)

	timers := m.GetTimers()
	if len(timers) != 1 {
		t.Errorf("expected 1 timer, got %d", len(timers))
	}
}

func TestDeleteTimer(t *testing.T) {
	m := NewManager(nil)

	m.AddTimer(&Timer{ID: "timer-1", Name: "Test"})
	err := m.DeleteTimer("timer-1")
	if err != nil {
		t.Errorf("DeleteTimer returned error: %v", err)
	}

	timers := m.GetTimers()
	if len(timers) != 0 {
		t.Errorf("expected 0 timers after delete, got %d", len(timers))
	}
}

func TestDeleteTimer_NotFound(t *testing.T) {
	m := NewManager(nil)
	err := m.DeleteTimer("non-existent")
	if err != nil {
		t.Errorf("DeleteTimer returned error: %v", err)
	}
}