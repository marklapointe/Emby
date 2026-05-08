package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestActivityHandler_NewActivityHandler(t *testing.T) {
	h := NewActivityHandler()
	if h == nil {
		t.Fatal("NewActivityHandler returned nil")
	}
}

func TestActivityHandler_GetActivities(t *testing.T) {
	h := NewActivityHandler()

	req := httptest.NewRequest(http.MethodGet, "/Activities", nil)
	w := httptest.NewRecorder()

	h.GetActivities(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	var activities []map[string]interface{}
	if err := json.NewDecoder(w.Body).Decode(&activities); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if len(activities) != 3 {
		t.Errorf("expected 3 activities, got %d", len(activities))
	}
}

func TestActivityHandler_GetActivity(t *testing.T) {
	h := NewActivityHandler()

	req := httptest.NewRequest(http.MethodGet, "/Activities/activity-1", nil)
	w := httptest.NewRecorder()

	h.GetActivity(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	var activity map[string]interface{}
	if err := json.NewDecoder(w.Body).Decode(&activity); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if activity["Id"] != "activity-1" {
		t.Errorf("expected Id 'activity-1', got '%v'", activity["Id"])
	}
}

func TestActivityHandler_GetActivityTypes(t *testing.T) {
	h := NewActivityHandler()

	req := httptest.NewRequest(http.MethodGet, "/Activities/Types", nil)
	w := httptest.NewRecorder()

	h.GetActivityTypes(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	var types []map[string]interface{}
	if err := json.NewDecoder(w.Body).Decode(&types); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if len(types) == 0 {
		t.Error("expected at least one activity type")
	}
}

func TestActivityHandler_GetActivitySummary(t *testing.T) {
	h := NewActivityHandler()

	req := httptest.NewRequest(http.MethodGet, "/Activities/Summary", nil)
	w := httptest.NewRecorder()

	h.GetActivitySummary(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	var summary map[string]interface{}
	if err := json.NewDecoder(w.Body).Decode(&summary); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if summary["TotalActivities"] == nil {
		t.Error("expected TotalActivities in summary")
	}
}