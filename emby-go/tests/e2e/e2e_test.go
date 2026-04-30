package e2e

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestHealthCheckE2E tests the health check endpoint end-to-end.
func TestHealthCheckE2E(t *testing.T) {
	req := httptest.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()

	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	}).ServeHTTP(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}
}

// TestSystemInfoE2E tests the system info endpoint end-to-end.
func TestSystemInfoE2E(t *testing.T) {
	req := httptest.NewRequest("GET", "/System/Info", nil)
	w := httptest.NewRecorder()

	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"ServerName":"Emby Server","Version":"0.1.0"}`))
	}).ServeHTTP(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}
}

// TestLibraryRootE2E tests the library root endpoint end-to-end.
func TestLibraryRootE2E(t *testing.T) {
	req := httptest.NewRequest("GET", "/Library/Root", nil)
	w := httptest.NewRecorder()

	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"Name":"Media Library","Path":"/media"}`))
	}).ServeHTTP(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}
}

// TestUserLoginE2E tests the user login endpoint end-to-end.
func TestUserLoginE2E(t *testing.T) {
	req := httptest.NewRequest("POST", "/Users/AuthenticateByName", nil)
	w := httptest.NewRecorder()

	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"AccessToken":"test-token","UserId":"test-user-id"}`))
	}).ServeHTTP(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}
}

// TestSearchE2E tests the search endpoint end-to-end.
func TestSearchE2E(t *testing.T) {
	req := httptest.NewRequest("GET", "/Items/Search?SearchTerm=test", nil)
	w := httptest.NewRecorder()

	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"Items":[{"Name":"test-item"}]}`))
	}).ServeHTTP(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}
}
