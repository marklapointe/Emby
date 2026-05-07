package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/emby/emby-go/internal/repository"
)

func TestNewFilterHandler(t *testing.T) {
	h := NewFilterHandler(nil)
	if h == nil {
		t.Fatal("NewFilterHandler returned nil")
	}
}

func TestGetCultures(t *testing.T) {
	h := NewFilterHandler(nil)

	req := httptest.NewRequest("GET", "/Cultures", nil)
	w := httptest.NewRecorder()

	h.GetCultures(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	var result []map[string]interface{}
	if err := json.NewDecoder(w.Body).Decode(&result); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if len(result) == 0 {
		t.Error("expected non-empty cultures list")
	}

	found := false
	for _, culture := range result {
		if culture["Name"] == "en-US" && culture["DisplayName"] == "English (United States)" {
			found = true
			break
		}
	}
	if !found {
		t.Error("expected to find en-US culture")
	}
}

func TestGetCountries(t *testing.T) {
	h := NewFilterHandler(nil)

	req := httptest.NewRequest("GET", "/Countries", nil)
	w := httptest.NewRecorder()

	h.GetCountries(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	var result []map[string]interface{}
	if err := json.NewDecoder(w.Body).Decode(&result); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if len(result) == 0 {
		t.Error("expected non-empty countries list")
	}

	found := false
	for _, country := range result {
		if country["Name"] == "US" && country["DisplayName"] == "United States" {
			found = true
			break
		}
	}
	if !found {
		t.Error("expected to find US country")
	}
}

func TestGetMusicGenres(t *testing.T) {
	h := NewFilterHandler(nil)

	req := httptest.NewRequest("GET", "/MusicGenres", nil)
	w := httptest.NewRecorder()

	h.GetMusicGenres(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	var result interface{}
	if err := json.NewDecoder(w.Body).Decode(&result); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
}

func TestGetGenres_WithRepo(t *testing.T) {
	itemRepo := &repository.ItemRepository{}
	h := NewFilterHandler(itemRepo)
	if h == nil {
		t.Fatal("NewFilterHandler returned nil")
	}
	if h.repo == nil {
		t.Error("repo should not be nil when passed a valid repo")
	}
}