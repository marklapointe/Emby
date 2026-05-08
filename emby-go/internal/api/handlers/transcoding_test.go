package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/emby/emby-go/internal/service/transcoding"
)

func TestNewTranscodingHandler(t *testing.T) {
	svc := transcoding.NewManager(nil, nil, nil)
	handler := NewTranscodingHandler(svc)
	if handler == nil {
		t.Fatal("expected non-nil handler")
	}
	if handler.transcodingSvc != svc {
		t.Error("expected transcoding service to be set")
	}
}

func TestGetTranscodingProfile(t *testing.T) {
	svc := transcoding.NewManager(nil, nil, nil)
	handler := NewTranscodingHandler(svc)
	req := httptest.NewRequest(http.MethodGet, "/TranscodingProfiles/1", nil)
	w := httptest.NewRecorder()
	handler.GetTranscodingProfile(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
	if w.Header().Get("Content-Type") != "application/json" {
		t.Errorf("expected application/json, got %s", w.Header().Get("Content-Type"))
	}
}

func TestGetActiveTranscode(t *testing.T) {
	svc := transcoding.NewManager(nil, nil, nil)
	handler := NewTranscodingHandler(svc)
	req := httptest.NewRequest(http.MethodGet, "/ActiveTranscodes/1", nil)
	w := httptest.NewRecorder()
	handler.GetActiveTranscode(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
	if w.Header().Get("Content-Type") != "application/json" {
		t.Errorf("expected application/json, got %s", w.Header().Get("Content-Type"))
	}
}

func TestStopTranscode(t *testing.T) {
	svc := transcoding.NewManager(nil, nil, nil)
	handler := NewTranscodingHandler(svc)
	req := httptest.NewRequest(http.MethodPost, "/ActiveTranscodes/1/Stop", nil)
	w := httptest.NewRecorder()
	handler.StopTranscode(w, req)

	if w.Code != http.StatusNoContent {
		t.Errorf("expected status 204, got %d", w.Code)
	}
}