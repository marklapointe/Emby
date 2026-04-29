package image

import (
	"fmt"
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
			// Return image data (placeholder for now)
			return []byte("image_data"), getContentType(imageType), nil
		}
	}

	return nil, "", fmt.Errorf("image not found: %s for type %s", itemID, imageType)
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

	return []byte("image_data"), getContentType(imageType), nil
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
			return []byte("image_data"), getContentType(imageType), nil
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
			return []byte("cropped_image_data"), getContentType(imageType), nil
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
			return []byte("resized_image_data"), getContentType(imageType), nil
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
			return []byte("rotated_image_data"), getContentType(imageType), nil
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
