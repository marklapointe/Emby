package metadata

import (
	"fmt"
	"sync"
	"time"

	"go.uber.org/zap"
)

// MetadataProvider represents a metadata provider.
type MetadataProvider struct {
	ID          string `json:"Id"`
	Name        string `json:"Name"`
	Type        string `json:"Type"`
	Enabled     bool   `json:"Enabled"`
	Order       int    `json:"Order"`
	Capabilities struct {
		SupportsAlternateTitles bool `json:"SupportsAlternateTitles"`
		CanRetrieveImages       bool `json:"CanRetrieveImages"`
		ImageProviders          []string `json:"ImageProviders"`
		SubtitleProviders       []string `json:"SubtitleProviders"`
	} `json:"Capabilities"`
}

// MetadataResult represents the result of a metadata search.
type MetadataResult struct {
	ItemID       string        `json:"ItemId"`
	Name         string        `json:"Name"`
	Overview     string        `json:"Overview,omitempty"`
	Tagline      string        `json:"Tagline,omitempty"`
	Taglines     []string      `json:"Taglines,omitempty"`
	HomePageURL  string        `json:"HomePageUrl,omitempty"`
	IMDBID       string        `json:"ImdbId,omitempty"`
	TVDBID       string        `json:"TvdbId,omitempty"`
	TVMAZEID     string        `json:"TvMazeId,omitempty"`
	IMDBRating   float64       `json:"ImdbRating,omitempty"`
	Year         int           `json:"Year,omitempty"`
	Runtime      int           `json:"Runtime,omitempty"`
	Studios      []string      `json:"Studios,omitempty"`
	Genres       []string      `json:"Genres,omitempty"`
	ProductionLocations []string `json:"ProductionLocations,omitempty"`
	Certification string       `json:"Certification,omitempty"`
	Networks     []string      `json:"Networks,omitempty"`
	Trailers     []Trailer     `json:"Trailers,omitempty"`
	People       []Person      `json:"People,omitempty"`
	Images       []MediaImage  `json:"Images,omitempty"`
	BackdropImages []MediaImage `json:"BackdropImages,omitempty"`
	LocalImages  []MediaImage  `json:"LocalImages,omitempty"`
	Providers    []ProviderID  `json:"ProviderIds,omitempty"`
	ParentIndex  int           `json:"ParentIndex,omitempty"`
	IndexNumberEnd int          `json:"IndexNumberEnd,omitempty"`
	IndexNumber  int           `json:"IndexNumber,omitempty"`
	IndexNumberStart int        `json:"IndexNumberStart,omitempty"`
	SeasonName   string        `json:"SeasonName,omitempty"`
	EpisodeName  string        `json:"EpisodeName,omitempty"`
	AirTime      string        `json:"AirTime,omitempty"`
	Aired        time.Time     `json:"Aired,omitempty"`
	OverviewType string        `json:"OverviewType,omitempty"`
	ProviderTypes []string      `json:"ProviderTypes,omitempty"`
}

// Trailer represents a video trailer.
type Trailer struct {
	Name   string `json:"Name"`
	Type   string `json:"Type"`
	URL    string `json:"Url"`
	Path   string `json:"Path,omitempty"`
	Internal bool  `json:"Internal"`
}

// Person represents a person associated with media.
type Person struct {
	Name     string `json:"Name"`
	Role     string `json:"Role,omitempty"`
	Character []string `json:"Characters,omitempty"`
	Type    string `json:"Type"`
	ID      string `json:"Id,omitempty"`
	Primary bool   `json:"Primary"`
	Order   int    `json:"Order,omitempty"`
	Profile string `json:"Profile,omitempty"`
}

// MediaImage represents a media image.
type MediaImage struct {
	Type       string `json:"Type"`
	URL        string `json:"Url,omitempty"`
	Path       string `json:"Path,omitempty"`
	Width      int    `json:"Width,omitempty"`
	Height     int    `json:"Height,omitempty"`
	HeightRaw  int    `json:"HeightRaw,omitempty"`
	WidthRaw   int    `json:"WidthRaw,omitempty"`
	Index      int    `json:"Index,omitempty"`
	Tag        string `json:"Tag,omitempty"`
	Provider   string `json:"Provider,omitempty"`
	RemoteURL  string `json:"RemoteUrl,omitempty"`
	AspectRatio string `json:"AspectRatio,omitempty"`
}

// ProviderID represents a provider identifier.
type ProviderID struct {
	ProviderName string `json:"ProviderName"`
	ID           string `json:"Id"`
}

// Manager handles metadata operations.
type Manager struct {
	mu          sync.RWMutex
	providers   map[string]*MetadataProvider
	logger      *zap.Logger
}

// NewManager creates a new metadata manager.
func NewManager(logger *zap.Logger) *Manager {
	return &Manager{
		providers: make(map[string]*MetadataProvider),
		logger:    logger,
	}
}

// RegisterProvider registers a new metadata provider.
func (m *Manager) RegisterProvider(provider *MetadataProvider) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.providers[provider.ID]; exists {
		return fmt.Errorf("provider already registered: %s", provider.ID)
	}

	m.providers[provider.ID] = provider
	if m.logger != nil {
		m.logger.Info("metadata provider registered", zap.String("id", provider.ID), zap.String("name", provider.Name))
	}
	return nil
}

// GetProvider returns a provider by ID.
func (m *Manager) GetProvider(id string) (*MetadataProvider, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	provider, exists := m.providers[id]
	return provider, exists
}

// GetAllProviders returns all registered providers.
func (m *Manager) GetAllProviders() []*MetadataProvider {
	m.mu.RLock()
	defer m.mu.RUnlock()

	providers := make([]*MetadataProvider, 0, len(m.providers))
	for _, provider := range m.providers {
		providers = append(providers, provider)
	}
	return providers
}

// GetEnabledProviders returns all enabled providers.
func (m *Manager) GetEnabledProviders() []*MetadataProvider {
	m.mu.RLock()
	defer m.mu.RUnlock()

	var providers []*MetadataProvider
	for _, provider := range m.providers {
		if provider.Enabled {
			providers = append(providers, provider)
		}
	}
	return providers
}

// GetProvidersByType returns providers filtered by type.
func (m *Manager) GetProvidersByType(providerType string) []*MetadataProvider {
	m.mu.RLock()
	defer m.mu.RUnlock()

	var providers []*MetadataProvider
	for _, provider := range m.providers {
		if providerType == "" || provider.Type == providerType {
			providers = append(providers, provider)
		}
	}
	return providers
}

// EnableProvider enables a provider.
func (m *Manager) EnableProvider(id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	provider, exists := m.providers[id]
	if !exists {
		return fmt.Errorf("provider not found: %s", id)
	}

	provider.Enabled = true
	return nil
}

// DisableProvider disables a provider.
func (m *Manager) DisableProvider(id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	provider, exists := m.providers[id]
	if !exists {
		return fmt.Errorf("provider not found: %s", id)
	}

	provider.Enabled = false
	return nil
}

// GetProviderCount returns the total number of providers.
func (m *Manager) GetProviderCount() int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return len(m.providers)
}

// GetEnabledProviderCount returns the number of enabled providers.
func (m *Manager) GetEnabledProviderCount() int {
	m.mu.RLock()
	defer m.mu.RUnlock()

	count := 0
	for _, provider := range m.providers {
		if provider.Enabled {
			count++
		}
	}
	return count
}
