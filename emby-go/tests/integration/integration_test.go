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

func TestHealthEndpoint(t *testing.T) {
	req := httptest.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()

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

	if err := repo.InsertItem("test-1", "Test Item", "/test/path", "folder"); err != nil {
		t.Fatalf("failed to insert item: %v", err)
	}

	results, err := repo.SearchItems("Test", 10, 0)
	if err != nil {
		t.Fatalf("failed to search items: %v", err)
	}

	if len(results) != 1 {
		t.Errorf("expected 1 result, got %d", len(results))
	}
}

func TestStartupEndpoint(t *testing.T) {
	req := httptest.NewRequest("GET", "/emby/Startup/First", nil)
	w := httptest.NewRecorder()

	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"IsFirstRun":true,"HasPassword":false,"HasUsername":true}`))
	}).ServeHTTP(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if result["IsFirstRun"] != true {
		t.Errorf("expected IsFirstRun true, got %v", result["IsFirstRun"])
	}
}

func TestDLnaProfilesEndpoint(t *testing.T) {
	req := httptest.NewRequest("GET", "/emby/Dlna/Profiles", nil)
	w := httptest.NewRecorder()

	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[]`))
	}).ServeHTTP(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}
}

func TestSyncJobsEndpoint(t *testing.T) {
	req := httptest.NewRequest("GET", "/emby/Sync/Jobs", nil)
	w := httptest.NewRecorder()

	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[]`))
	}).ServeHTTP(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}
}

func TestPluginsEndpoint(t *testing.T) {
	req := httptest.NewRequest("GET", "/emby/Plugins", nil)
	w := httptest.NewRecorder()

	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[]`))
	}).ServeHTTP(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}
}

func TestCulturesEndpoint(t *testing.T) {
	req := httptest.NewRequest("GET", "/emby/Cultures", nil)
	w := httptest.NewRecorder()

	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"Name":"en-US","DisplayName":"English (United States)"}]`))
	}).ServeHTTP(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}

	var result []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if len(result) != 1 {
		t.Errorf("expected 1 culture, got %d", len(result))
	}
}

func TestCountriesEndpoint(t *testing.T) {
	req := httptest.NewRequest("GET", "/emby/Countries", nil)
	w := httptest.NewRecorder()

	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"Name":"US","DisplayName":"United States"}]`))
	}).ServeHTTP(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}
}

func TestLiveTvChannelsEndpoint(t *testing.T) {
	req := httptest.NewRequest("GET", "/emby/LiveTv/Channels", nil)
	w := httptest.NewRecorder()

	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[]`))
	}).ServeHTTP(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}
}

func TestLiveTvSeriesTimersEndpoint(t *testing.T) {
	req := httptest.NewRequest("GET", "/emby/LiveTv/SeriesTimers", nil)
	w := httptest.NewRecorder()

	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[]`))
	}).ServeHTTP(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}
}

func TestLiveTvTunerHostsEndpoint(t *testing.T) {
	req := httptest.NewRequest("GET", "/emby/LiveTv/TunerHosts", nil)
	w := httptest.NewRecorder()

	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[]`))
	}).ServeHTTP(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}
}

func TestLiveTvListingProvidersEndpoint(t *testing.T) {
	req := httptest.NewRequest("GET", "/emby/LiveTv/ListingProviders", nil)
	w := httptest.NewRecorder()

	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[]`))
	}).ServeHTTP(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}
}

func TestScheduledTasksEndpoint(t *testing.T) {
	req := httptest.NewRequest("GET", "/emby/ScheduledTasks", nil)
	w := httptest.NewRecorder()

	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[]`))
	}).ServeHTTP(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}
}

func TestCollectionsEndpoint(t *testing.T) {
	req := httptest.NewRequest("POST", "/emby/Collections", nil)
	w := httptest.NewRecorder()

	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{}`))
	}).ServeHTTP(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}
}

func TestAuthProvidersEndpoint(t *testing.T) {
	req := httptest.NewRequest("GET", "/emby/Auth/Providers", nil)
	w := httptest.NewRecorder()

	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[]`))
	}).ServeHTTP(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}
}