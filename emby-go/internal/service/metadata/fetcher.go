package metadata

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// Provider represents a metadata provider.
type Provider struct {
	Name        string            `json:"name"`
	ID          string            `json:"id"`
	Enabled     bool              `json:"enabled"`
	Priority    int               `json:"priority"`
	LastUpdated time.Time         `json:"lastUpdated"`
	Config      map[string]string `json:"config"`
}

// Fetcher handles metadata fetching.
type Fetcher struct {
	providers map[string]*Provider
	cacheDir  string
	mu        sync.RWMutex
	httpClient *http.Client
}

// NewFetcher creates a new metadata fetcher.
func NewFetcher(cacheDir string) *Fetcher {
	return &Fetcher{
		providers: make(map[string]*Provider),
		cacheDir:  cacheDir,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// RegisterProvider registers a metadata provider.
func (f *Fetcher) RegisterProvider(provider *Provider) {
	f.mu.Lock()
	defer f.mu.Unlock()

	f.providers[provider.ID] = provider
}

// GetProviders returns all registered providers.
func (f *Fetcher) GetProviders() []*Provider {
	f.mu.RLock()
	defer f.mu.RUnlock()

	providers := make([]*Provider, 0, len(f.providers))
	for _, p := range f.providers {
		providers = append(providers, p)
	}

	return providers
}

// FetchMetadata fetches metadata for an item.
func (f *Fetcher) FetchMetadata(itemID string, providerID string) (map[string]interface{}, error) {
	provider, ok := f.providers[providerID]
	if !ok {
		return nil, fmt.Errorf("provider not found: %s", providerID)
	}

	if !provider.Enabled {
		return nil, fmt.Errorf("provider disabled: %s", providerID)
	}

	// Check cache first
	cachePath := filepath.Join(f.cacheDir, fmt.Sprintf("%s-%s.json", itemID, providerID))
	if data, err := os.ReadFile(cachePath); err == nil {
		var metadata map[string]interface{}
		if err := json.Unmarshal(data, &metadata); err == nil {
			return metadata, nil
		}
	}

	// Fetch from provider
	metadata, err := f.fetchFromProvider(itemID, providerID)
	if err != nil {
		return nil, err
	}

	// Cache the result
	if err := os.MkdirAll(filepath.Dir(cachePath), 0755); err == nil {
		if data, err := json.Marshal(metadata); err == nil {
			os.WriteFile(cachePath, data, 0644)
		}
	}

	return metadata, nil
}

// fetchFromProvider fetches metadata from a specific provider.
func (f *Fetcher) fetchFromProvider(itemID, providerID string) (map[string]interface{}, error) {
	provider, ok := f.providers[providerID]
	if !ok {
		return nil, fmt.Errorf("provider not registered: %s", providerID)
	}

	baseURL, ok := provider.Config["api_url"]
	if !ok {
		baseURL = fmt.Sprintf("https://api.%s.com", providerID)
	}

	apiKey, _ := provider.Config["api_key"]

	url := fmt.Sprintf("%s/items/%s/metadata", baseURL, itemID)
	if apiKey != "" {
		url += "?apikey=" + apiKey
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	if apiKey != "" {
		req.Header.Set("Authorization", "Bearer "+apiKey)
	}

	resp, err := f.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch metadata: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("provider returned status: %d", resp.StatusCode)
	}

	var metadata map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&metadata); err != nil {
		return nil, fmt.Errorf("failed to decode metadata: %w", err)
	}

	return metadata, nil
}

// DownloadImage downloads an image from a URL.
func (f *Fetcher) DownloadImage(url, destPath string) error {
	resp, err := f.httpClient.Get(url)
	if err != nil {
		return fmt.Errorf("failed to download image: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download image: HTTP %d", resp.StatusCode)
	}

	if err := os.MkdirAll(filepath.Dir(destPath), 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	out, err := os.Create(destPath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer out.Close()

	if _, err := io.Copy(out, resp.Body); err != nil {
		return fmt.Errorf("failed to write image: %w", err)
	}

	return nil
}

// ClearCache clears the metadata cache.
func (f *Fetcher) ClearCache() error {
	return os.RemoveAll(f.cacheDir)
}
