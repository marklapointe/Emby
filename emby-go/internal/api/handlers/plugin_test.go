package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewPluginHandler(t *testing.T) {
	h := NewPluginHandler()
	if h == nil {
		t.Fatal("NewPluginHandler returned nil")
	}
}

func TestGetPlugins(t *testing.T) {
	h := NewPluginHandler()

	req := httptest.NewRequest("GET", "/Plugins", nil)
	w := httptest.NewRecorder()

	h.GetPlugins(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	var result interface{}
	if err := json.NewDecoder(w.Body).Decode(&result); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
}

func TestDeletePlugin(t *testing.T) {
	h := NewPluginHandler()

	req := httptest.NewRequest("DELETE", "/Plugins/plugin-123", nil)
	w := httptest.NewRecorder()

	h.DeletePlugin(w, req)

	if w.Code != http.StatusNoContent {
		t.Errorf("expected status 204, got %d", w.Code)
	}
}

func TestGetPluginConfiguration(t *testing.T) {
	h := NewPluginHandler()

	req := httptest.NewRequest("GET", "/Plugins/plugin-123/Configuration", nil)
	w := httptest.NewRecorder()

	h.GetPluginConfiguration(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
}

func TestGetSecurityInfo(t *testing.T) {
	h := NewPluginHandler()

	req := httptest.NewRequest("GET", "/Plugins/SecurityInfo", nil)
	w := httptest.NewRecorder()

	h.GetSecurityInfo(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(w.Body).Decode(&result); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if result["IsMBSupporter"] != true {
		t.Errorf("expected IsMBSupporter to be true")
	}

	if result["IsSuperUser"] != true {
		t.Errorf("expected IsSuperUser to be true")
	}
}

func TestGetReleased(t *testing.T) {
	h := NewPluginHandler()

	req := httptest.NewRequest("GET", "/Plugins/Released", nil)
	w := httptest.NewRecorder()

	h.GetReleased(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
}