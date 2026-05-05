package media

import (
	"testing"
)

func TestNewManager(t *testing.T) {
	m := NewManager(nil, nil)
	if m == nil {
		t.Fatal("NewManager returned nil")
	}
	if m.mediaDB == nil {
		t.Error("mediaDB map not initialized")
	}
}

func TestGetMediaSource_NotFound(t *testing.T) {
	m := NewManager(nil, nil)
	_, err := m.GetMediaSource("non-existent")
	if err == nil {
		t.Error("expected error for non-existent media source")
	}
}

func TestGetMediaSources_NotFound(t *testing.T) {
	m := NewManager(nil, nil)
	_, err := m.GetMediaSources("non-existent")
	if err == nil {
		t.Error("expected error for non-existent media sources")
	}
}

func TestGetMediaInfo_NotFound(t *testing.T) {
	m := NewManager(nil, nil)
	_, err := m.GetMediaInfo("/non/existent/path.mp4")
	if err == nil {
		t.Error("expected error for non-existent file")
	}
}

func TestStreamManager_NewManager(t *testing.T) {
	m := NewStreamManager(5, nil)
	if m == nil {
		t.Fatal("NewStreamManager returned nil")
	}
	if m.maxStreams != 5 {
		t.Errorf("expected maxStreams 5, got %d", m.maxStreams)
	}
	if m.activeStreams == nil {
		t.Error("activeStreams map not initialized")
	}
}

func TestStreamManager_GetMetrics(t *testing.T) {
	m := NewStreamManager(5, nil)
	metrics := m.GetMetrics()
	if metrics == nil {
		t.Fatal("GetMetrics returned nil")
	}
	if metrics.TotalStreamsCreated != 0 {
		t.Errorf("expected 0 TotalStreamsCreated, got %d", metrics.TotalStreamsCreated)
	}
}

func TestStreamManager_GetStreamViewers_NotFound(t *testing.T) {
	m := NewStreamManager(5, nil)
	viewers := m.GetStreamViewers("non-existent", "profile")
	if viewers != 0 {
		t.Errorf("expected 0 viewers, got %d", viewers)
	}
}

func TestStreamError_Error(t *testing.T) {
	err := &StreamError{Message: "test message"}
	if err.Error() != "test message" {
		t.Errorf("expected 'test message', got '%s'", err.Error())
	}
}

func TestStreamManager_EvictIdleStreams(t *testing.T) {
	m := NewStreamManager(1, nil)
	// Should not panic
	m.evictIdleStreams()
}

func TestStreamManager_RemoveViewer_NotFound(t *testing.T) {
	m := NewStreamManager(5, nil)
	// Should not panic
	m.RemoveViewer("non-existent", "profile", "viewer-1")
}