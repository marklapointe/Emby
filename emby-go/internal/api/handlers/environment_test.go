package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewEnvironmentHandler(t *testing.T) {
	handler := NewEnvironmentHandler()
	if handler == nil {
		t.Fatal("expected non-nil handler")
	}
}

func TestGetDrives(t *testing.T) {
	handler := NewEnvironmentHandler()
	req := httptest.NewRequest(http.MethodGet, "/Environment/Drives", nil)
	w := httptest.NewRecorder()
	handler.GetDrives(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
	if w.Header().Get("Content-Type") != "application/json" {
		t.Errorf("expected application/json, got %s", w.Header().Get("Content-Type"))
	}
}

func TestGetDefaultDirectoryBrowser(t *testing.T) {
	handler := NewEnvironmentHandler()
	req := httptest.NewRequest(http.MethodGet, "/Environment/DefaultDirectoryBrowser", nil)
	w := httptest.NewRecorder()
	handler.GetDefaultDirectoryBrowser(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
	if w.Header().Get("Content-Type") != "application/json" {
		t.Errorf("expected application/json, got %s", w.Header().Get("Content-Type"))
	}
}

func TestGetNetworkShares(t *testing.T) {
	handler := NewEnvironmentHandler()
	req := httptest.NewRequest(http.MethodGet, "/Environment/NetworkShares", nil)
	w := httptest.NewRecorder()
	handler.GetNetworkShares(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
	if w.Header().Get("Content-Type") != "application/json" {
		t.Errorf("expected application/json, got %s", w.Header().Get("Content-Type"))
	}
}

func TestGetParentPath(t *testing.T) {
	handler := NewEnvironmentHandler()
	req := httptest.NewRequest(http.MethodGet, "/Environment/ParentPath?path=/test/path", nil)
	w := httptest.NewRecorder()
	handler.GetParentPath(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
	if w.Header().Get("Content-Type") != "application/json" {
		t.Errorf("expected application/json, got %s", w.Header().Get("Content-Type"))
	}
}

func TestGetParentPathRoot(t *testing.T) {
	handler := NewEnvironmentHandler()
	req := httptest.NewRequest(http.MethodGet, "/Environment/ParentPath?path=/", nil)
	w := httptest.NewRecorder()
	handler.GetParentPath(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
}

func TestGetParentPathEmpty(t *testing.T) {
	handler := NewEnvironmentHandler()
	req := httptest.NewRequest(http.MethodGet, "/Environment/ParentPath", nil)
	w := httptest.NewRecorder()
	handler.GetParentPath(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
}

func TestGetDirectoryContents(t *testing.T) {
	handler := NewEnvironmentHandler()
	req := httptest.NewRequest(http.MethodGet, "/Environment/DirectoryContents?path=/", nil)
	w := httptest.NewRecorder()
	handler.GetDirectoryContents(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
	if w.Header().Get("Content-Type") != "application/json" {
		t.Errorf("expected application/json, got %s", w.Header().Get("Content-Type"))
	}
}

func TestGetDirectoryContentsNoPath(t *testing.T) {
	handler := NewEnvironmentHandler()
	req := httptest.NewRequest(http.MethodGet, "/Environment/DirectoryContents", nil)
	w := httptest.NewRecorder()
	handler.GetDirectoryContents(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
}