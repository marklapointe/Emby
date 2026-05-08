package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/emby/emby-go/internal/config"
	"github.com/emby/emby-go/internal/database"
	"github.com/emby/emby-go/internal/repository"
	"github.com/go-chi/chi/v5"
)

func setupTestRouter(t *testing.T) (*chi.Mux, func()) {
	cfg := &config.Config{
		Database: config.DatabaseConfig{
			Path:         ":memory:",
			MaxOpenConns: 1,
			MaxIdleConns: 1,
		},
	}

	dbMgr, err := database.NewManager(&cfg.Database)
	if err != nil {
		t.Fatalf("failed to create db manager: %v", err)
	}

	itemRepo := repository.NewItemRepository(dbMgr.DB())
	if err := itemRepo.CreateSchema(); err != nil {
		t.Fatalf("failed to create schema: %v", err)
	}

	// Create a test router
	r := chi.NewRouter()

	// Register the handlers directly without using the full router setup
	// This allows us to test the handler behavior in isolation
	scanner := NewLibraryHandler(nil, itemRepo, nil)
	envHandler := NewEnvironmentHandler()

	// Register routes with the exact paths the middleware would produce after lowercasing
	r.Get("/libraries/availableoptions", scanner.GetAvailableOptions)
	r.Get("/environment/defaultdirectorybrowser", envHandler.GetDefaultDirectoryBrowser)

	cleanup := func() {
		dbMgr.Close()
	}

	return r, cleanup
}

func TestGetAvailableOptionsHandler(t *testing.T) {
	handler := NewLibraryHandler(nil, nil, nil)
	req := httptest.NewRequest(http.MethodGet, "/Libraries/AvailableOptions?LibraryContentType=movies&IsNewLibrary=false", nil)
	w := httptest.NewRecorder()

	handler.GetAvailableOptions(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	if w.Header().Get("Content-Type") != "application/json" {
		t.Errorf("expected application/json, got %s", w.Header().Get("Content-Type"))
	}

	// Verify response structure
	var result map[string]interface{}
	if err := json.NewDecoder(w.Body).Decode(&result); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	// Check that TypeOptions exists
	typeOptions, ok := result["TypeOptions"].([]interface{})
	if !ok {
		t.Fatal("expected TypeOptions to be present and array")
	}
	if len(typeOptions) == 0 {
		t.Error("expected at least one type option")
	}

	// Check first type option structure
	firstType := typeOptions[0].(map[string]interface{})
	if firstType["Type"] == nil {
		t.Error("expected Type field in type option")
	}
	if firstType["SupportedImageTypes"] == nil {
		t.Error("expected SupportedImageTypes field in type option")
	}
}

func TestGetAvailableOptionsHandlerWithQueryParams(t *testing.T) {
	handler := NewLibraryHandler(nil, nil, nil)

	tests := []struct {
		name            string
		libraryType     string
		isNew           bool
		expectTypeCount int
	}{
		{"movies", "movies", false, 1},
		{"tvshows", "tvshows", false, 1},
		{"music", "music", false, 1},
		{"new library", "movies", true, 1},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/Libraries/AvailableOptions?LibraryContentType="+tc.libraryType+"&IsNewLibrary="+boolToString(tc.isNew), nil)
			w := httptest.NewRecorder()

			handler.GetAvailableOptions(w, req)

			if w.Code != http.StatusOK {
				t.Errorf("expected status 200, got %d", w.Code)
			}
		})
	}
}

func boolToString(b bool) string {
	if b {
		return "true"
	}
	return "false"
}

func TestGetDefaultDirectoryBrowserHandler(t *testing.T) {
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

	// Verify response has path field
	var result map[string]string
	if err := json.NewDecoder(w.Body).Decode(&result); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if result["path"] == "" {
		t.Error("expected path field to be present")
	}
}

func TestAvailableOptionsRouterIntegration(t *testing.T) {
	router, cleanup := setupTestRouter(t)
	defer cleanup()

	tests := []struct {
		name       string
		path       string
		wantStatus int
	}{
		{"lowercase path", "/libraries/availableoptions", http.StatusOK},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tc.path+"?LibraryContentType=movies&IsNewLibrary=false", nil)
			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			if w.Code != tc.wantStatus {
				t.Errorf("path %s: expected status %d, got %d", tc.path, tc.wantStatus, w.Code)
			}
		})
	}
}

func TestDefaultDirectoryBrowserRouterIntegration(t *testing.T) {
	router, cleanup := setupTestRouter(t)
	defer cleanup()

	tests := []struct {
		name       string
		path       string
		wantStatus int
	}{
		{"lowercase path", "/environment/defaultdirectorybrowser", http.StatusOK},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tc.path, nil)
			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			if w.Code != tc.wantStatus {
				t.Errorf("path %s: expected status %d, got %d", tc.path, tc.wantStatus, w.Code)
			}
		})
	}
}

func TestGetAvailableOptionsResponseStructure(t *testing.T) {
	handler := NewLibraryHandler(nil, nil, nil)
	req := httptest.NewRequest(http.MethodGet, "/Libraries/AvailableOptions", nil)
	w := httptest.NewRecorder()

	handler.GetAvailableOptions(w, req)

	var result map[string]interface{}
	if err := json.NewDecoder(w.Body).Decode(&result); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	// Verify all expected fields are present
	expectedFields := []string{"MetadataSavers", "MetadataReaders", "SubtitleFetchers", "TypeOptions"}
	for _, field := range expectedFields {
		if _, ok := result[field]; !ok {
			t.Errorf("expected field %s to be present", field)
		}
	}

	// Verify MetadataSavers, MetadataReaders, SubtitleFetchers are arrays
	for _, field := range []string{"MetadataSavers", "MetadataReaders", "SubtitleFetchers"} {
		if _, ok := result[field].([]interface{}); !ok {
			t.Errorf("expected %s to be an array", field)
		}
	}

	// Verify TypeOptions is an array
	typeOptions, ok := result["TypeOptions"].([]interface{})
	if !ok {
		t.Fatal("TypeOptions should be an array")
	}

	// Each type option should have: Type, MetadataFetchers, ImageFetchers, SupportedImageTypes, DefaultImageOptions
	for i, to := range typeOptions {
		typeOpt, ok := to.(map[string]interface{})
		if !ok {
			t.Errorf("TypeOption[%d] should be a map", i)
			continue
		}

		requiredFields := []string{"Type", "MetadataFetchers", "ImageFetchers", "SupportedImageTypes", "DefaultImageOptions"}
		for _, field := range requiredFields {
			if _, ok := typeOpt[field]; !ok {
				t.Errorf("TypeOption[%d] missing field %s", i, field)
			}
		}
	}
}

func TestDefaultDirectoryBrowserReturnsRootPath(t *testing.T) {
	handler := NewEnvironmentHandler()
	req := httptest.NewRequest(http.MethodGet, "/Environment/DefaultDirectoryBrowser", nil)
	w := httptest.NewRecorder()

	handler.GetDefaultDirectoryBrowser(w, req)

	var result map[string]string
	if err := json.NewDecoder(w.Body).Decode(&result); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	// Default should be root "/"
	if result["path"] != "/" {
		t.Errorf("expected default path to be '/', got '%s'", result["path"])
	}
}
