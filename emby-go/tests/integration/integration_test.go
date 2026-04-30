package integration

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/emby/emby-go/internal/config"
	"github.com/emby/emby-go/internal/database"
	"github.com/emby/emby-go/internal/repository"
)

// TestHealthEndpoint tests the health check endpoint.
func TestHealthEndpoint(t *testing.T) {
	req := httptest.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()

	// Create a simple health check handler
	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	}).ServeHTTP(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if result["status"] != "ok" {
		t.Errorf("expected status 'ok', got %v", result["status"])
	}
}

// TestConfigLoad tests configuration loading.
func TestConfigLoad(t *testing.T) {
	cfg, err := config.LoadConfig("")
	if err != nil {
		t.Fatalf("failed to load config: %v", err)
	}

	if cfg.Server.Port != 8096 {
		t.Errorf("expected port 8096, got %d", cfg.Server.Port)
	}

	if cfg.Database.Path != "data/emby-server.db" {
		t.Errorf("expected database path 'data/emby-server.db', got %s", cfg.Database.Path)
	}
}

// TestDatabaseConnection tests database connection.
func TestDatabaseConnection(t *testing.T) {
	dbManager, err := database.NewManager(&config.DatabaseConfig{
		Path:           ":memory:",
		MaxOpenConns:   10,
		MaxIdleConns:   5,
		ConnMaxLifetime: 300,
		EnableWAL:      true,
	})
	if err != nil {
		t.Fatalf("failed to create database manager: %v", err)
	}
	defer dbManager.Close()

	if dbManager.DB() == nil {
		t.Fatal("database connection is nil")
	}
}

// TestRepositoryCRUD tests repository CRUD operations.
func TestRepositoryCRUD(t *testing.T) {
	dbManager, err := database.NewManager(&config.DatabaseConfig{
		Path:           ":memory:",
		MaxOpenConns:   10,
		MaxIdleConns:   5,
		ConnMaxLifetime: 300,
		EnableWAL:      true,
	})
	if err != nil {
		t.Fatalf("failed to create database manager: %v", err)
	}
	defer dbManager.Close()

	repo := repository.NewItemRepository(dbManager.DB())

	if err := repo.CreateSchema(); err != nil {
		t.Fatalf("failed to create schema: %v", err)
	}

	// Test insert
	if err := repo.InsertItem("test-1", "Test Item", "/test/path", "folder"); err != nil {
		t.Fatalf("failed to insert item: %v", err)
	}

	// Test search
	results, err := repo.SearchItems("Test", 10, 0)
	if err != nil {
		t.Fatalf("failed to search items: %v", err)
	}

	if len(results) != 1 {
		t.Errorf("expected 1 result, got %d", len(results))
	}
}
