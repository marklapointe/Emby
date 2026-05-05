package device

import (
	"fmt"
	"sync"
	"time"

	"go.uber.org/zap"
)

// Device represents a connected device.
type Device struct {
	ID            string
	Name          string
	ProductName   string
	DeviceClass   string
	OS            string
	DeviceProfile *DeviceProfile
	LastSeen      time.Time
	IsAuthorized  bool
	Capabilities  map[string]interface{}
}

// DeviceProfile represents a device's capabilities profile.
type DeviceProfile struct {
	Name              string
	ID                string
	MaxStreamingBitrate int
	MaxStaticBitrate  int
	MaxChannels       int
	SubtitleProfiles  []SubtitleProfile
	TranscodingProfiles []TranscodingProfile
	DirectPlayProfiles []DirectPlayProfile
}

// SubtitleProfile represents subtitle support.
type SubtitleProfile struct {
	Format  string
	Method  string
}

// TranscodingProfile represents transcoding support.
type TranscodingProfile struct {
	Container       string
	Type            string
	AudioCodec      string
	VideoCodec      string
	MaxAudioChannels string
	Protocol        string
	MaxVideoBitrate int
}

// DirectPlayProfile represents direct play support.
type DirectPlayProfile struct {
	Container  string
	Type       string
	AudioCodec string
	VideoCodec string
}

// Manager manages connected devices.
type Manager struct {
	mu       sync.RWMutex
	devices  map[string]*Device
	logger   *zap.Logger
}

// NewManager creates a new device manager.
func NewManager(logger *zap.Logger) *Manager {
	if logger == nil {
		logger = zap.NewNop()
	}
	return &Manager{
		devices: make(map[string]*Device),
		logger:  logger,
	}
}

// RegisterDevice registers a new device.
func (m *Manager) RegisterDevice(d *Device) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.devices[d.ID]; exists {
		return fmt.Errorf("device already registered: %s", d.ID)
	}

	d.LastSeen = time.Now()
	m.devices[d.ID] = d
	m.logger.Info("device registered", zap.String("id", d.ID), zap.String("name", d.Name))
	return nil
}

// UpdateDevice updates a device's information.
func (m *Manager) UpdateDevice(id string, name, productName string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	d, exists := m.devices[id]
	if !exists {
		return fmt.Errorf("device not found: %s", id)
	}

	if name != "" {
		d.Name = name
	}
	if productName != "" {
		d.ProductName = productName
	}
	d.LastSeen = time.Now()
	return nil
}

// GetDevice returns a device by ID.
func (m *Manager) GetDevice(id string) (*Device, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	d, exists := m.devices[id]
	return d, exists
}

// GetDevices returns all registered devices.
func (m *Manager) GetDevices() []*Device {
	m.mu.RLock()
	defer m.mu.RUnlock()

	devices := make([]*Device, 0, len(m.devices))
	for _, d := range m.devices {
		devices = append(devices, d)
	}
	return devices
}

// RemoveDevice removes a device.
func (m *Manager) RemoveDevice(id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.devices[id]; !exists {
		return fmt.Errorf("device not found: %s", id)
	}

	delete(m.devices, id)
	m.logger.Info("device removed", zap.String("id", id))
	return nil
}

// GetDeviceProfile returns the device profile for a device.
func (m *Manager) GetDeviceProfile(id string) (*DeviceProfile, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	d, exists := m.devices[id]
	if !exists {
		return nil, false
	}

	return d.DeviceProfile, true
}

// GetActiveDeviceCount returns the number of devices seen recently.
func (m *Manager) GetActiveDeviceCount() int {
	m.mu.RLock()
	defer m.mu.RUnlock()

	count := 0
	for _, d := range m.devices {
		if time.Since(d.LastSeen) < 24*time.Hour {
			count++
		}
	}
	return count
}
