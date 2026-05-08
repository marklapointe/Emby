package metadata

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"go.uber.org/zap"
)

// Manager tests

func TestNewManager(t *testing.T) {
	logger := zap.NewNop()
	m := NewManager(logger)
	if m == nil {
		t.Fatal("NewManager returned nil")
	}
	if m.providers == nil {
		t.Error("providers map is nil")
	}
}

func TestManager_RegisterProvider(t *testing.T) {
	logger := zap.NewNop()
	m := NewManager(logger)

	provider := &MetadataProvider{
		ID:   "test-provider",
		Name: "Test Provider",
		Type: "movie",
	}

	err := m.RegisterProvider(provider)
	if err != nil {
		t.Errorf("RegisterProvider failed: %v", err)
	}

	// Verify it's registered
	p, exists := m.GetProvider("test-provider")
	if !exists {
		t.Error("provider not found after registration")
	}
	if p.Name != "Test Provider" {
		t.Errorf("expected name 'Test Provider', got '%s'", p.Name)
	}
}

func TestManager_RegisterProvider_Duplicate(t *testing.T) {
	logger := zap.NewNop()
	m := NewManager(logger)

	provider := &MetadataProvider{ID: "dup-provider", Name: "Dup"}
	err := m.RegisterProvider(provider)
	if err != nil {
		t.Errorf("first registration failed: %v", err)
	}

	err = m.RegisterProvider(provider)
	if err == nil {
		t.Error("expected error for duplicate registration")
	}
}

func TestManager_GetProvider(t *testing.T) {
	logger := zap.NewNop()
	m := NewManager(logger)

	m.RegisterProvider(&MetadataProvider{ID: "provider1", Name: "Provider 1"})

	p, exists := m.GetProvider("provider1")
	if !exists {
		t.Error("expected provider to exist")
	}
	if p.Name != "Provider 1" {
		t.Errorf("expected 'Provider 1', got '%s'", p.Name)
	}

	_, exists = m.GetProvider("nonexistent")
	if exists {
		t.Error("expected provider to not exist")
	}
}

func TestManager_GetAllProviders(t *testing.T) {
	logger := zap.NewNop()
	m := NewManager(logger)

	m.RegisterProvider(&MetadataProvider{ID: "p1", Name: "P1"})
	m.RegisterProvider(&MetadataProvider{ID: "p2", Name: "P2", Enabled: true})

	providers := m.GetAllProviders()
	if len(providers) != 2 {
		t.Errorf("expected 2 providers, got %d", len(providers))
	}
}

func TestManager_GetEnabledProviders(t *testing.T) {
	logger := zap.NewNop()
	m := NewManager(logger)

	m.RegisterProvider(&MetadataProvider{ID: "p1", Name: "P1", Enabled: false})
	m.RegisterProvider(&MetadataProvider{ID: "p2", Name: "P2", Enabled: true})
	m.RegisterProvider(&MetadataProvider{ID: "p3", Name: "P3", Enabled: true})

	providers := m.GetEnabledProviders()
	if len(providers) != 2 {
		t.Errorf("expected 2 enabled providers, got %d", len(providers))
	}
}

func TestManager_GetProvidersByType(t *testing.T) {
	logger := zap.NewNop()
	m := NewManager(logger)

	m.RegisterProvider(&MetadataProvider{ID: "m1", Type: "movie"})
	m.RegisterProvider(&MetadataProvider{ID: "m2", Type: "movie"})
	m.RegisterProvider(&MetadataProvider{ID: "tv1", Type: "tv"})

	movieProviders := m.GetProvidersByType("movie")
	if len(movieProviders) != 2 {
		t.Errorf("expected 2 movie providers, got %d", len(movieProviders))
	}

	allProviders := m.GetProvidersByType("")
	if len(allProviders) != 3 {
		t.Errorf("expected 3 providers with empty type filter, got %d", len(allProviders))
	}
}

func TestManager_EnableProvider(t *testing.T) {
	logger := zap.NewNop()
	m := NewManager(logger)

	m.RegisterProvider(&MetadataProvider{ID: "p1", Enabled: false})

	err := m.EnableProvider("p1")
	if err != nil {
		t.Errorf("EnableProvider failed: %v", err)
	}

	p, _ := m.GetProvider("p1")
	if !p.Enabled {
		t.Error("expected provider to be enabled")
	}

	err = m.EnableProvider("nonexistent")
	if err == nil {
		t.Error("expected error for nonexistent provider")
	}
}

func TestManager_DisableProvider(t *testing.T) {
	logger := zap.NewNop()
	m := NewManager(logger)

	m.RegisterProvider(&MetadataProvider{ID: "p1", Enabled: true})

	err := m.DisableProvider("p1")
	if err != nil {
		t.Errorf("DisableProvider failed: %v", err)
	}

	p, _ := m.GetProvider("p1")
	if p.Enabled {
		t.Error("expected provider to be disabled")
	}

	err = m.DisableProvider("nonexistent")
	if err == nil {
		t.Error("expected error for nonexistent provider")
	}
}

func TestManager_GetProviderCount(t *testing.T) {
	logger := zap.NewNop()
	m := NewManager(logger)

	if m.GetProviderCount() != 0 {
		t.Errorf("expected 0, got %d", m.GetProviderCount())
	}

	m.RegisterProvider(&MetadataProvider{ID: "p1"})
	m.RegisterProvider(&MetadataProvider{ID: "p2"})

	if m.GetProviderCount() != 2 {
		t.Errorf("expected 2, got %d", m.GetProviderCount())
	}
}

func TestManager_GetEnabledProviderCount(t *testing.T) {
	logger := zap.NewNop()
	m := NewManager(logger)

	m.RegisterProvider(&MetadataProvider{ID: "p1", Enabled: true})
	m.RegisterProvider(&MetadataProvider{ID: "p2", Enabled: false})
	m.RegisterProvider(&MetadataProvider{ID: "p3", Enabled: true})

	if m.GetEnabledProviderCount() != 2 {
		t.Errorf("expected 2, got %d", m.GetEnabledProviderCount())
	}
}

// Fetcher tests

func TestFetcher_NewFetcher(t *testing.T) {
	f := NewFetcher("/tmp/cache")
	if f == nil {
		t.Fatal("NewFetcher returned nil")
	}
	if f.cacheDir != "/tmp/cache" {
		t.Errorf("expected cacheDir '/tmp/cache', got '%s'", f.cacheDir)
	}
}

func TestFetcher_RegisterProvider(t *testing.T) {
	f := NewFetcher("/tmp/cache")
	provider := &Provider{ID: "test", Name: "Test"}
	f.RegisterProvider(provider)

	providers := f.GetProviders()
	if len(providers) != 1 {
		t.Errorf("expected 1 provider, got %d", len(providers))
	}
}

func TestFetcher_FetchMetadata_ProviderNotFound(t *testing.T) {
	f := NewFetcher("/tmp/cache")
	_, err := f.FetchMetadata("item1", "nonexistent")
	if err == nil {
		t.Error("expected error for nonexistent provider")
	}
}

func TestFetcher_FetchMetadata_ProviderDisabled(t *testing.T) {
	f := NewFetcher("/tmp/cache")
	f.RegisterProvider(&Provider{ID: "disabled", Enabled: false})

	_, err := f.FetchMetadata("item1", "disabled")
	if err == nil {
		t.Error("expected error for disabled provider")
	}
}

func TestFetcher_FetchMetadata_CacheHit(t *testing.T) {
	tmpDir := t.TempDir()
	f := NewFetcher(tmpDir)

	f.RegisterProvider(&Provider{ID: "tmdb", Enabled: true})

	cachePath := tmpDir + "/item1-tmdb.json"
	cacheData := `{"title":"Cached Title"}`
	err := os.WriteFile(cachePath, []byte(cacheData), 0644)
	if err != nil {
		t.Fatalf("failed to write cache file: %v", err)
	}

	metadata, err := f.FetchMetadata("item1", "tmdb")
	if err != nil {
		t.Errorf("FetchMetadata failed: %v", err)
	}
	if title, ok := metadata["title"].(string); !ok || title != "Cached Title" {
		t.Errorf("expected cached title 'Cached Title', got '%v'", metadata["title"])
	}
}

func TestFetcher_FetchMetadata_FetchFromProvider(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"title":"Fetched Title","year":2024}`))
	}))
	defer server.Close()

	tmpDir := t.TempDir()
	f := NewFetcher(tmpDir)

	f.RegisterProvider(&Provider{
		ID:      "test-provider",
		Enabled: true,
		Config:  map[string]string{"api_url": server.URL},
	})

	metadata, err := f.FetchMetadata("item1", "test-provider")
	if err != nil {
		t.Errorf("FetchMetadata failed: %v", err)
	}
	if title, ok := metadata["title"].(string); !ok || title != "Fetched Title" {
		t.Errorf("expected 'Fetched Title', got '%v'", metadata["title"])
	}
}

func TestFetcher_FetchMetadata_FetchError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	tmpDir := t.TempDir()
	f := NewFetcher(tmpDir)

	f.RegisterProvider(&Provider{
		ID:      "err-provider",
		Enabled: true,
		Config:  map[string]string{"api_url": server.URL},
	})

	_, err := f.FetchMetadata("item1", "err-provider")
	if err == nil {
		t.Error("expected error for failed fetch")
	}
}

func TestFetcher_DownloadImage(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("fake image data"))
	}))
	defer server.Close()

	tmpDir := t.TempDir()
	destPath := tmpDir + "/image.jpg"

	f := NewFetcher(tmpDir)
	err := f.DownloadImage(server.URL, destPath)
	if err != nil {
		t.Errorf("DownloadImage failed: %v", err)
	}

	data, err := os.ReadFile(destPath)
	if err != nil {
		t.Errorf("failed to read downloaded file: %v", err)
	}
	if string(data) != "fake image data" {
		t.Errorf("expected 'fake image data', got '%s'", string(data))
	}
}

func TestFetcher_DownloadImage_HTTPError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	tmpDir := t.TempDir()
	f := NewFetcher(tmpDir)

	err := f.DownloadImage(server.URL, tmpDir+"/image.jpg")
	if err == nil {
		t.Error("expected error for HTTP 404")
	}
}

func TestFetcher_ClearCache(t *testing.T) {
	tmpDir := t.TempDir()
	f := NewFetcher(tmpDir)

	// Create a cache file
	cachePath := tmpDir + "/cache.json"
	os.WriteFile(cachePath, []byte("{}"), 0644)

	err := f.ClearCache()
	if err != nil {
		t.Errorf("ClearCache failed: %v", err)
	}

	if _, err := os.Stat(cachePath); !os.IsNotExist(err) {
		t.Error("expected cache file to be deleted")
	}
}

// RateLimiter tests

func TestRateLimiter_Allow(t *testing.T) {
	rl := NewRateLimiter(2, time.Minute)

	if !rl.Allow("key1") {
		t.Error("first request should be allowed")
	}
	if !rl.Allow("key1") {
		t.Error("second request should be allowed")
	}
	if rl.Allow("key1") {
		t.Error("third request should be denied")
	}

	// Different key should be allowed
	if !rl.Allow("key2") {
		t.Error("different key should be allowed")
	}
}

func TestRateLimiter_WaitTime(t *testing.T) {
	rl := NewRateLimiter(2, time.Minute)

	// No requests
	wait := rl.WaitTime("key1")
	if wait != 0 {
		t.Errorf("expected 0 wait time, got %v", wait)
	}

	rl.Allow("key1")
	rl.Allow("key1")

	// Should have wait time now
	wait = rl.WaitTime("key1")
	if wait <= 0 || wait > time.Minute {
		t.Errorf("expected positive wait time, got %v", wait)
	}
}

func TestRateLimiter_Clear(t *testing.T) {
	rl := NewRateLimiter(2, time.Minute)

	rl.Allow("key1")
	rl.Allow("key1")
	rl.Clear()

	// Should be able to make requests again
	if !rl.Allow("key1") {
		t.Error("should be allowed after clear")
	}
}

// HTTPClient tests

func TestHTTPClient_Get_RateLimited(t *testing.T) {
	rl := NewRateLimiter(1, time.Hour)
	hc := NewHTTPClient(rl, "test")

	// First request should succeed
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	_, err := hc.Get(server.URL)
	if err != nil {
		t.Errorf("first request failed: %v", err)
	}

	// Second request should be rate limited
	_, err = hc.Get(server.URL)
	if err == nil {
		t.Error("expected rate limit error")
	}
}

func TestNewTMDbProvider(t *testing.T) {
	p := NewTMDbProvider("test-api-key")
	if p == nil {
		t.Fatal("NewTMDbProvider returned nil")
	}
	if p.apiKey != "test-api-key" {
		t.Errorf("expected apiKey 'test-api-key', got '%s'", p.apiKey)
	}
	if p.httpClient == nil {
		t.Error("httpClient is nil")
	}
}

func TestTMDbProvider_GetImageURL(t *testing.T) {
	p := NewTMDbProvider("test-api-key")
	url := p.GetImageURL("/poster.jpg", "w500")
	expected := "https://image.tmdb.org/t/p/w500/poster.jpg"
	if url != expected {
		t.Errorf("expected '%s', got '%s'", expected, url)
	}
}

func TestNewTVDbProvider(t *testing.T) {
	p := NewTVDbProvider("test-api-key")
	if p == nil {
		t.Fatal("NewTVDbProvider returned nil")
	}
	if p.apiKey != "test-api-key" {
		t.Errorf("expected apiKey 'test-api-key', got '%s'", p.apiKey)
	}
}

func TestTVDbProvider_GetImageURL(t *testing.T) {
	p := NewTVDbProvider("test-api-key")
	url := p.GetImageURL("poster", "/banners/posters/123.jpg")
	expected := "https://artworks.thetvdb.com/banners//banners/posters/123.jpg"
	if url != expected {
		t.Errorf("expected '%s', got '%s'", expected, url)
	}
}

func TestTMDbProvider_SearchMovie_Error(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	p := &TMDbProvider{
		apiKey: "test",
		httpClient: server.Client(),
	}

	_, err := p.SearchMovie("test")
	if err == nil {
		t.Error("expected error for HTTP 500")
	}
}

func TestTVDbProvider_SearchSeries_Error(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	p := &TVDbProvider{
		apiKey: "test",
		httpClient: server.Client(),
	}

	_, err := p.SearchSeries("test")
	if err == nil {
		t.Error("expected error for HTTP 500")
	}
}

func TestTVDbProvider_GetSeries_Error(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	p := &TVDbProvider{
		apiKey: "test",
		httpClient: server.Client(),
	}

	_, err := p.GetSeries(0)
	if err == nil {
		t.Error("expected error for HTTP 404")
	}
}