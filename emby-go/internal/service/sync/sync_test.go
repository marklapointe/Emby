package sync

import (
	"testing"

	"go.uber.org/zap"
)

func TestNewManager(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	m := NewManager(logger)
	if m == nil {
		t.Fatal("NewManager returned nil")
	}
	if m.logger == nil {
		t.Error("logger should not be nil")
	}
}

func TestNewManager_NilLogger(t *testing.T) {
	m := NewManager(nil)
	if m == nil {
		t.Fatal("NewManager returned nil")
	}
	if m.logger != nil {
		t.Error("logger should be nil when passed nil")
	}
}