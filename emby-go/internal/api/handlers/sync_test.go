package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/emby/emby-go/internal/service/sync"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func TestNewSyncHandler(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	syncSvc := sync.NewManager(logger)
	h := NewSyncHandler(syncSvc)
	if h == nil {
		t.Fatal("NewSyncHandler returned nil")
	}
	if h.syncSvc == nil {
		t.Error("syncSvc should not be nil")
	}
}

func TestGetJobs(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	syncSvc := sync.NewManager(logger)
	h := NewSyncHandler(syncSvc)

	req := httptest.NewRequest("GET", "/Sync/Jobs", nil)
	w := httptest.NewRecorder()

	h.GetJobs(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	var result interface{}
	if err := json.NewDecoder(w.Body).Decode(&result); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
}

func TestCreateJob(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	syncSvc := sync.NewManager(logger)
	h := NewSyncHandler(syncSvc)

	req := httptest.NewRequest("POST", "/Sync/Jobs", nil)
	w := httptest.NewRecorder()

	h.CreateJob(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
}

func TestGetJob_WithChi(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	syncSvc := sync.NewManager(logger)
	h := NewSyncHandler(syncSvc)

	r := chi.NewRouter()
	r.Get("/Sync/Jobs/{id}", h.GetJob)

	req := httptest.NewRequest("GET", "/Sync/Jobs/job-123", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(w.Body).Decode(&result); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if result["Id"] != "job-123" {
		t.Errorf("expected Id 'job-123', got %v", result["Id"])
	}
}

func TestDeleteJob(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	syncSvc := sync.NewManager(logger)
	h := NewSyncHandler(syncSvc)

	r := chi.NewRouter()
	r.Delete("/Sync/Jobs/{id}", h.DeleteJob)

	req := httptest.NewRequest("DELETE", "/Sync/Jobs/job-123", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusNoContent {
		t.Errorf("expected status 204, got %d", w.Code)
	}
}

func TestAddItemToJob_WithChi(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	syncSvc := sync.NewManager(logger)
	h := NewSyncHandler(syncSvc)

	r := chi.NewRouter()
	r.Post("/Sync/Jobs/{id}/Items/{itemId}", h.AddItemToJob)

	req := httptest.NewRequest("POST", "/Sync/Jobs/job-123/Items/item-456", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusNoContent {
		t.Errorf("expected status 204, got %d", w.Code)
	}
}