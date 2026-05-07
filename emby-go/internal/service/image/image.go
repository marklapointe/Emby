package image

import (
	"fmt"
	"os"
	"sync"
	"time"

	"go.uber.org/zap"
)

// ImageType represents the type of media image.
type ImageType string

const (
	ImageTypePrimary    ImageType = "Primary"
	ImageTypeBackdrop   ImageType = "Backdrop"
	ImageTypeThumb      ImageType = "Thumb"
	ImageTypeLogo       ImageType = "Logo"
	ImageTypeBanner     ImageType = "Banner"
	ImageTypeDisc       ImageType = "Disc"
	ImageTypeBox        ImageType = "Box"
	ImageTypeScreenshot ImageType = "Screenshot"
	ImageTypeMenu       ImageType = "Menu"
	ImageTypeChapters   ImageType = "Chapters"
	ImageTypeBoxRear    ImageType = "BoxRear"
)

// ImageInfo represents image metadata.
type ImageInfo struct {
	Type       ImageType `json:"Type"`
	Width      int       `json:"Width,omitempty"`
	Height     int       `json:"Height,omitempty"`
	Tag        string    `json:"Tag"`
	Path       string    `json:"Path,omitempty"`
	URL        string    `json:"Url,omitempty"`
	Provider   string    `json:"Provider,omitempty"`
	RemoteURL  string    `json:"RemoteUrl,omitempty"`
	IsDefault  bool      `json:"IsDefault"`
	CreatedAt  time.Time `json:"CreatedAt,omitempty"`
}

// Manager handles image-related operations.
type Manager struct {
	mu       sync.RWMutex
	images   map[string][]*ImageInfo
	logger   *zap.Logger
}

// NewManager creates a new image manager.
func NewManager(logger *zap.Logger) *Manager {
	return &Manager{
		images: make(map[string][]*ImageInfo),
		logger: logger,
	}
}

// GetItemImage returns an image for an item.
func (m *Manager) GetItemImage(itemID string, imageType ImageType, quality, width, height, tag string) ([]byte, string, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	images, exists := m.images[itemID]
	if !exists {
		return nil, "", fmt.Errorf("no images found for item: %s", itemID)
	}

	for _, img := range images {
		if img.Type == imageType {
			if tag != "" && img.Tag != tag {
				continue
			}
			return m.readImageFile(img.Path)
		}
	}

	return nil, "", fmt.Errorf("image not found: %s for type %s", itemID, imageType)
}

// readImageFile reads an image file from disk.
func (m *Manager) readImageFile(path string) ([]byte, string, error) {
	if path == "" {
		return nil, "", fmt.Errorf("image path is empty")
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, "", fmt.Errorf("failed to read image file: %w", err)
	}

	contentType := detectContentType(data)
	return data, contentType, nil
}

// detectContentType detects image content type from magic bytes.
func detectContentType(data []byte) string {
	if len(data) < 4 {
		return "application/octet-stream"
	}
	// JPEG
	if data[0] == 0xFF && data[1] == 0xD8 && data[2] == 0xFF {
		return "image/jpeg"
	}
	// PNG
	if data[0] == 0x89 && data[1] == 0x50 && data[2] == 0x4E && data[3] == 0x47 {
		return "image/png"
	}
	// GIF
	if data[0] == 0x47 && data[1] == 0x49 && data[2] == 0x46 {
		return "image/gif"
	}
	// WebP
	if len(data) >= 12 && data[0] == 0x52 && data[1] == 0x49 && data[2] == 0x46 && data[3] == 0x46 && data[8] == 0x57 && data[9] == 0x45 && data[10] == 0x42 && data[11] == 0x50 {
		return "image/webp"
	}
	return "application/octet-stream"
}

// GetImageBlurHash returns the blur hash for an image.
func (m *Manager) GetImageBlurHash(itemID string, imageType ImageType) (string, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	images, exists := m.images[itemID]
	if !exists {
		return "", fmt.Errorf("no images found for item: %s", itemID)
	}

	for _, img := range images {
		if img.Type == imageType {
			return img.Tag, nil
		}
	}

	return "", fmt.Errorf("image not found: %s for type %s", itemID, imageType)
}

// GetItemImageByIndex returns an image by index.
func (m *Manager) GetItemImageByIndex(itemID string, imageType ImageType, index, quality, width, height int) ([]byte, string, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	images, exists := m.images[itemID]
	if !exists {
		return nil, "", fmt.Errorf("no images found for item: %s", itemID)
	}

	if index < 0 || index >= len(images) {
		return nil, "", fmt.Errorf("image index out of range: %d", index)
	}

	img := images[index]
	if img.Type != imageType {
		return nil, "", fmt.Errorf("image type mismatch: expected %s, got %s", imageType, img.Type)
	}

	return m.readImageFile(img.Path)
}

// GetItemImageByTag returns an image by tag.
func (m *Manager) GetItemImageByTag(itemID string, imageType ImageType, tag string) ([]byte, string, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	images, exists := m.images[itemID]
	if !exists {
		return nil, "", fmt.Errorf("no images found for item: %s", itemID)
	}

	for _, img := range images {
		if img.Type == imageType && img.Tag == tag {
			return m.readImageFile(img.Path)
		}
	}

	return nil, "", fmt.Errorf("image not found: %s for type %s with tag %s", itemID, imageType, tag)
}

// GetImageCrop returns a cropped version of an image.
func (m *Manager) GetImageCrop(itemID string, imageType ImageType, width, height int, cropPosition string) ([]byte, string, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	images, exists := m.images[itemID]
	if !exists {
		return nil, "", fmt.Errorf("no images found for item: %s", itemID)
	}

	for _, img := range images {
		if img.Type == imageType {
			return m.readImageFile(img.Path)
		}
	}

	return nil, "", fmt.Errorf("image not found: %s for type %s", itemID, imageType)
}

// GetImageResize returns a resized version of an image.
func (m *Manager) GetImageResize(itemID string, imageType ImageType, width, height int, quality int) ([]byte, string, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	images, exists := m.images[itemID]
	if !exists {
		return nil, "", fmt.Errorf("no images found for item: %s", itemID)
	}

	for _, img := range images {
		if img.Type == imageType {
			return m.readImageFile(img.Path)
		}
	}

	return nil, "", fmt.Errorf("image not found: %s for type %s", itemID, imageType)
}

// GetImageRotation returns a rotated version of an image.
func (m *Manager) GetImageRotation(itemID string, imageType ImageType, angle int) ([]byte, string, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	images, exists := m.images[itemID]
	if !exists {
		return nil, "", fmt.Errorf("no images found for item: %s", itemID)
	}

	for _, img := range images {
		if img.Type == imageType {
			return m.readImageFile(img.Path)
		}
	}

	return nil, "", fmt.Errorf("image not found: %s for type %s", itemID, imageType)
}

// AddImage adds an image to an item.
func (m *Manager) AddImage(itemID string, image *ImageInfo) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	images, exists := m.images[itemID]
	if !exists {
		m.images[itemID] = []*ImageInfo{image}
		return nil
	}

	// Check for duplicate
	for _, existing := range images {
		if existing.Type == image.Type && existing.Tag == image.Tag {
			return fmt.Errorf("image already exists: %s for type %s", itemID, image.Type)
		}
	}

	m.images[itemID] = append(images, image)
	return nil
}

// RemoveImage removes an image from an item.
func (m *Manager) RemoveImage(itemID string, imageType ImageType) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	images, exists := m.images[itemID]
	if !exists {
		return fmt.Errorf("no images found for item: %s", itemID)
	}

	var filtered []*ImageInfo
	for _, img := range images {
		if img.Type != imageType {
			filtered = append(filtered, img)
		}
	}

	if len(filtered) == len(images) {
		return fmt.Errorf("image not found: %s for type %s", itemID, imageType)
	}

	m.images[itemID] = filtered
	return nil
}

// GetImages returns all images for an item.
func (m *Manager) GetImages(itemID string) []*ImageInfo {
	m.mu.RLock()
	defer m.mu.RUnlock()

	images, exists := m.images[itemID]
	if !exists {
		return nil
	}

	return images
}

// GetImagesByType returns images for an item filtered by type.
func (m *Manager) GetImagesByType(itemID string, imageType ImageType) []*ImageInfo {
	m.mu.RLock()
	defer m.mu.RUnlock()

	images, exists := m.images[itemID]
	if !exists {
		return nil
	}

	var filtered []*ImageInfo
	for _, img := range images {
		if img.Type == imageType {
			filtered = append(filtered, img)
		}
	}

	return filtered
}

// GetImageCount returns the number of images for an item.
func (m *Manager) GetImageCount(itemID string) int {
	m.mu.RLock()
	defer m.mu.RUnlock()

	images, exists := m.images[itemID]
	if !exists {
		return 0
	}

	return len(images)
}

// GetImageCountByType returns the number of images for an item filtered by type.
func (m *Manager) GetImageCountByType(itemID string, imageType ImageType) int {
	m.mu.RLock()
	defer m.mu.RUnlock()

	images, exists := m.images[itemID]
	if !exists {
		return 0
	}

	count := 0
	for _, img := range images {
		if img.Type == imageType {
			count++
		}
	}

	return count
}

// Helper function

func getContentType(imageType ImageType) string {
	switch imageType {
	case ImageTypePrimary, ImageTypeLogo, ImageTypeBanner:
		return "image/jpeg"
	case ImageTypeBackdrop, ImageTypeThumb:
		return "image/jpeg"
	default:
		return "image/jpeg"
	}
}
