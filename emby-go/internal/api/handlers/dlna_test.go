package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/emby/emby-go/internal/service/dlna"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func TestNewDLNAHandler(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	dlnaSvc := dlna.NewManager(logger)
	h := NewDLNAHandler(dlnaSvc)
	if h == nil {
		t.Fatal("NewDLNAHandler returned nil")
	}
	if h.dlnaSvc == nil {
		t.Error("dlnaSvc should not be nil")
	}
}

func TestGetProfiles(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	dlnaSvc := dlna.NewManager(logger)
	h := NewDLNAHandler(dlnaSvc)

	req := httptest.NewRequest("GET", "/Dlna/Profiles", nil)
	w := httptest.NewRecorder()

	h.GetProfiles(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	var result interface{}
	if err := json.NewDecoder(w.Body).Decode(&result); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if result == nil {
		t.Error("expected non-nil result")
	}
}

func TestGetProfile_WithChi(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	dlnaSvc := dlna.NewManager(logger)
	h := NewDLNAHandler(dlnaSvc)

	r := chi.NewRouter()
	r.Get("/Dlna/Profiles/{id}", h.GetProfile)

	req := httptest.NewRequest("GET", "/Dlna/Profiles/test-id", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(w.Body).Decode(&result); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if result["Id"] != "test-id" {
		t.Errorf("expected Id 'test-id', got %v", result["Id"])
	}
}

func TestGetProfileInfos(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	dlnaSvc := dlna.NewManager(logger)
	h := NewDLNAHandler(dlnaSvc)

	req := httptest.NewRequest("GET", "/Dlna/ProfileInfos", nil)
	w := httptest.NewRecorder()

	h.GetProfileInfos(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
}

func TestGetDefaultProfile(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	dlnaSvc := dlna.NewManager(logger)
	h := NewDLNAHandler(dlnaSvc)

	req := httptest.NewRequest("GET", "/Dlna/Profiles/Default", nil)
	w := httptest.NewRecorder()

	h.GetDefaultProfile(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
}