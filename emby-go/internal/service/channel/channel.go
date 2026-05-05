package channel

import (
	"sync"

	"go.uber.org/zap"
)

type Manager struct {
	logger   *zap.Logger
	mu       sync.RWMutex
	channels map[string]*Channel
}

type Channel struct {
	ID          string            `json:"Id"`
	Name        string            `json:"Name"`
	Type        string            `json:"Type"`
	ChannelType string            `json:"ChannelType"`
	Number      int               `json:"Number"`
	ImageURL    string            `json:"ImageUrl"`
	ChannelNumber string          `json:"ChannelNumber"`
	Folders     []*ChannelFolder  `json:"Folders,omitempty"`
	Items       []*ChannelItem    `json:"Items,omitempty"`
}

type ChannelFolder struct {
	ID          string `json:"Id"`
	Name        string `json:"Name"`
	ImageURL    string `json:"ImageUrl,omitempty"`
	Description string `json:"Description,omitempty"`
}

type ChannelItem struct {
	ID          string `json:"Id"`
	Name        string `json:"Name"`
	Type        string `json:"Type"`
	ChannelID   string `json:"ChannelId"`
	ContentType string `json:"ContentType"`
	MediaType   string `json:"MediaType"`
	ImageURL    string `json:"ImageUrl,omitempty"`
	Overview    string `json:"Overview,omitempty"`
	Duration    int64  `json:"Duration,omitempty"`
}

func NewManager(logger *zap.Logger) *Manager {
	return &Manager{
		logger:   logger,
		channels: make(map[string]*Channel),
	}
}

func (m *Manager) GetChannels() []*Channel {
	m.mu.RLock()
	defer m.mu.RUnlock()

	channels := make([]*Channel, 0, len(m.channels))
	for _, ch := range m.channels {
		channels = append(channels, ch)
	}
	return channels
}

func (m *Manager) GetChannel(id string) (*Channel, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	ch, ok := m.channels[id]
	return ch, ok
}

func (m *Manager) AddChannel(ch *Channel) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.channels[ch.ID] = ch
}

func (m *Manager) RemoveChannel(id string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.channels, id)
}

func (m *Manager) GetChannelFolders(channelID string) []*ChannelFolder {
	m.mu.RLock()
	defer m.mu.RUnlock()

	ch, ok := m.channels[channelID]
	if !ok {
		return nil
	}
	return ch.Folders
}

func (m *Manager) GetChannelItems(channelID, userID string) []*ChannelItem {
	m.mu.RLock()
	defer m.mu.RUnlock()

	ch, ok := m.channels[channelID]
	if !ok {
		return nil
	}
	return ch.Items
}
