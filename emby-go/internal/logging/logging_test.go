package logging

import (
	"testing"
)

func TestNewLogger(t *testing.T) {
	logger, err := NewLogger("info", "json")
	if err != nil {
		t.Fatalf("NewLogger returned error: %v", err)
	}
	if logger == nil {
		t.Fatal("NewLogger returned nil")
	}
}

func TestNewLogger_WithLevel(t *testing.T) {
	tests := []struct {
		level string
	}{
		{"debug"},
		{"info"},
		{"warn"},
		{"error"},
		{"invalid"}, // defaults to info
	}

	for _, tt := range tests {
		logger, err := NewLogger(tt.level, "json")
		if err != nil {
			t.Errorf("NewLogger(%q) returned error: %v", tt.level, err)
		}
		if logger == nil {
			t.Errorf("NewLogger(%q) returned nil", tt.level)
		}
	}
}

func TestNewLogger_ConsoleFormat(t *testing.T) {
	logger, err := NewLogger("info", "console")
	if err != nil {
		t.Fatalf("NewLogger returned error: %v", err)
	}
	if logger == nil {
		t.Fatal("NewLogger returned nil")
	}
}

func TestNewSilentLogger(t *testing.T) {
	logger := NewSilentLogger()
	if logger == nil {
		t.Fatal("NewSilentLogger returned nil")
	}
}

func TestSilentLogger_Info(t *testing.T) {
	logger := NewSilentLogger()
	// Should not panic
	logger.Info("test info")
}

func TestSilentLogger_Debug(t *testing.T) {
	logger := NewSilentLogger()
	logger.Debug("test debug")
}

func TestSilentLogger_Error(t *testing.T) {
	logger := NewSilentLogger()
	logger.Error("test error")
}