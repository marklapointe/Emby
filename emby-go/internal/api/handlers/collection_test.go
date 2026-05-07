package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/emby/emby-go/internal/repository"
)

func TestNewCollectionHandler(t *testing.T) {
	h := NewCollectionHandler(nil)
	if h == nil {
		t.Fatal("NewCollectionHandler returned nil")
	}
}

func TestNewCollectionHandler_WithRepo(t *testing.T) {
	itemRepo := &repository.ItemRepository{}
	h := NewCollectionHandler(itemRepo)
	if h == nil {
		t.Fatal("NewCollectionHandler returned nil")
	}
	if h.itemRepo == nil {
		t.Error("itemRepo should not be nil when passed a valid repo")
	}
}

func TestCreateCollection(t *testing.T) {
	h := NewCollectionHandler(nil)

	req := httptest.NewRequest("POST", "/Collections", nil)
	w := httptest.NewRecorder()

	h.CreateCollection(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	if w.Header().Get("Content-Type") != "application/json" {
		t.Errorf("expected Content-Type application/json, got %s", w.Header().Get("Content-Type"))
	}
}

func TestAddToCollection(t *testing.T) {
	h := NewCollectionHandler(nil)

	req := httptest.NewRequest("POST", "/Collections/123/Items", nil)
	w := httptest.NewRecorder()

	h.AddToCollection(w, req)

	if w.Code != http.StatusNoContent {
		t.Errorf("expected status 204, got %d", w.Code)
	}
}