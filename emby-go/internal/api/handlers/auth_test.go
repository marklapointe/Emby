package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewAuthHandler(t *testing.T) {
	h := NewAuthHandler()
	if h == nil {
		t.Fatal("NewAuthHandler returned nil")
	}
}

func TestGetAuthProviders(t *testing.T) {
	h := NewAuthHandler()

	req := httptest.NewRequest("GET", "/Auth/Providers", nil)
	w := httptest.NewRecorder()

	h.GetAuthProviders(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	var result interface{}
	if err := json.NewDecoder(w.Body).Decode(&result); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
}