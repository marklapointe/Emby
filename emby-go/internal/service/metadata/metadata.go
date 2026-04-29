package metadata

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"go.uber.org/zap"
)

// Provider represents a metadata provider.
type Provider struct {
	ID            string
	Name          string
	Enabled       bool
	BaseURL       string
	APIKey        string
	LastUpdated   time.Time
	Items         []MediaMetadata
	Logger        *zap.Logger
}

// MediaMetadata represents metadata for a media item.
type MediaMetadata struct {
	ID              string    `json:"id"`
	Name            string    `json:"name"`
	Overview        string    `json:"overview"`
	Tagline         string    `json:"tagline"`
	Rating          float32   `json:"rating"`
	Year            int       `json:"year"`
	Genres          []string  `json:"genres"`
	Studios         []string  `json:"studios"`
	Actors          []Person  `json:"actors"`
	Directors       []Person  `json:"directors"`
	Writers         []Person  `json:"writers"`
	PosterURL       string    `json:"posterUrl"`
	BackdropURL     string    `json:"backdropUrl"`
	TrailerURL      string    `json:"trailerUrl"`
	IMDBID          string    `json:"imdbId"`
	TMDBID          string    `json:"tmdbId"`
	Language        string    `json:"language"`
	Certification   string    `json:"certification"`
	Runtime         int       `json:"runtime"`
	ContentType     string    `json:"contentType"`
}

// Person represents a person in the media.
type Person struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
	Type string `json:"type"`
}

// Manager manages metadata providers.
type Manager struct {
	mu         sync.RWMutex
	providers  map[string]*Provider
	httpClient *http.Client
	logger     *zap.Logger
}

// NewManager creates a new metadata manager.
func NewManager(logger *zap.Logger) *Manager {
	return &Manager{
		providers:  make(map[string]*Provider),
		httpClient: &http.Client{Timeout: 30 * time.Second},
		logger:     logger,
	}
}

// RegisterProvider registers a new metadata provider.
func (m *Manager) RegisterProvider(p *Provider) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.providers[p.ID]; exists {
		return fmt.Errorf("provider already registered: %s", p.ID)
	}

	p.Logger = m.logger
	m.providers[p.ID] = p
	m.logger.Info("metadata provider registered", zap.String("id", p.ID))
	return nil
}

// GetProvider returns a provider by ID.
func (m *Manager) GetProvider(id string) (*Provider, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	p, exists := m.providers[id]
	return p, exists
}

// GetProviders returns all registered providers.
func (m *Manager) GetProviders() []*Provider {
	m.mu.RLock()
	defer m.mu.RUnlock()

	providers := make([]*Provider, 0, len(m.providers))
	for _, p := range m.providers {
		providers = append(providers, p)
	}
	return providers
}

// SearchMetadata searches for metadata using all enabled providers.
func (m *Manager) SearchMetadata(query string) ([]MediaMetadata, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	var results []MediaMetadata

	for _, provider := range m.providers {
		if !provider.Enabled {
			continue
		}

		metadata, err := m.searchProvider(provider, query)
		if err != nil {
			m.logger.Error("provider search error", zap.String("provider", provider.ID), zap.Error(err))
			continue
		}

		results = append(results, metadata...)
	}

	return results, nil
}

// searchProvider searches a single provider.
func (m *Manager) searchProvider(provider *Provider, query string) ([]MediaMetadata, error) {
	url := fmt.Sprintf("%s/Search?query=%s&apiKey=%s", provider.BaseURL, query, provider.APIKey)

	resp, err := m.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("search request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search error: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	var results []MediaMetadata
	if err := json.Unmarshal(body, &results); err != nil {
		return nil, fmt.Errorf("parse response: %w", err)
	}

	return results, nil
}

// GetMetadata retrieves metadata for a specific item.
func (m *Manager) GetMetadata(providerID, itemID string) (*MediaMetadata, error) {
	provider, exists := m.GetProvider(providerID)
	if !exists {
		return nil, fmt.Errorf("provider not found: %s", providerID)
	}

	url := fmt.Sprintf("%s/Items/%s?apiKey=%s", provider.BaseURL, itemID, provider.APIKey)

	resp, err := m.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("metadata request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("metadata error: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	var metadata MediaMetadata
	if err := json.Unmarshal(body, &metadata); err != nil {
		return nil, fmt.Errorf("parse response: %w", err)
	}

	return &metadata, nil
}

// UpdateProvider updates a provider's configuration.
func (m *Manager) UpdateProvider(id string, name, baseURL, apiKey string, enabled bool) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	p, exists := m.providers[id]
	if !exists {
		return fmt.Errorf("provider not found: %s", id)
	}

	if name != "" {
		p.Name = name
	}
	if baseURL != "" {
		p.BaseURL = baseURL
	}
	if apiKey != "" {
		p.APIKey = apiKey
	}
	p.Enabled = enabled
	p.LastUpdated = time.Now()

	return nil
}

// EnableProvider enables a provider.
func (m *Manager) EnableProvider(id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	p, exists := m.providers[id]
	if !exists {
		return fmt.Errorf("provider not found: %s", id)
	}

	p.Enabled = true
	return nil
}

// DisableProvider disables a provider.
func (m *Manager) DisableProvider(id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	p, exists := m.providers[id]
	if !exists {
		return fmt.Errorf("provider not found: %s", id)
	}

	p.Enabled = false
	return nil
}
