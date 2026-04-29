package image

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
	"sync"
	"time"

	"github.com/disintegration/imaging"
	"go.uber.org/zap"
)

// Manager handles image processing operations.
type Manager struct {
	mu       sync.RWMutex
	cache    map[string]*CachedImage
	logger   *zap.Logger
	maxCache int
}

// CachedImage represents a cached image.
type CachedImage struct {
	Data       []byte
	Width      int
	Height     int
	Format     string
	LastAccess time.Time
}

// NewManager creates a new image manager.
func NewManager(logger *zap.Logger) *Manager {
	return &Manager{
		cache:    make(map[string]*CachedImage),
		logger:   logger,
		maxCache: 1000,
	}
}

// ProcessImage processes an image with the given options.
func (m *Manager) ProcessImage(srcPath string, opts *ImageOptions) ([]byte, error) {
	// Open source image
	f, err := os.Open(srcPath)
	if err != nil {
		return nil, fmt.Errorf("open image: %w", err)
	}
	defer f.Close()

	// Decode image
	img, _, err := image.Decode(f)
	if err != nil {
		return nil, fmt.Errorf("decode image: %w", err)
	}

	// Apply transformations
	result := img
	if opts != nil {
		result = m.applyOptions(result, opts)
	}

	// Encode result
	var buf []byte
	switch opts.OutputFormat {
	case "jpeg", "":
		if err := jpeg.Encode(&bufferWriter{buf: &buf}, result, &jpeg.Options{Quality: opts.Quality}); err != nil {
			return nil, fmt.Errorf("encode jpeg: %w", err)
		}
	case "png":
		if err := png.Encode(&bufferWriter{buf: &buf}, result); err != nil {
			return nil, fmt.Errorf("encode png: %w", err)
		}
	default:
		return nil, fmt.Errorf("unsupported output format: %s", opts.OutputFormat)
	}

	return buf, nil
}

// applyOptions applies image processing options.
func (m *Manager) applyOptions(img image.Image, opts *ImageOptions) image.Image {
	result := img

	// Resize
	if opts.Width > 0 || opts.Height > 0 {
		result = imaging.Resize(result, opts.Width, opts.Height, imaging.Lanczos)
	}

	// Crop
	if opts.Crop {
		bounds := img.Bounds()
		size := min(bounds.Dx(), bounds.Dy())
		result = imaging.CropCenter(result, size, size)
	}

	// Rotate
	if opts.Rotate > 0 {
		result = imaging.Rotate(result, opts.Rotate, color.Transparent)
	}

	// Adjust brightness
	if opts.Brightness != 0 {
		result = imaging.AdjustBrightness(result, opts.Brightness)
	}

	// Adjust contrast
	if opts.Contrast != 0 {
		result = imaging.AdjustContrast(result, opts.Contrast)
	}

	// Adjust saturation
	if opts.Saturation != 0 {
		result = imaging.AdjustSaturation(result, opts.Saturation)
	}

	// Adjust hue (skip - not supported by imaging package)
	// if opts.Hue != 0 {
	// 	result = imaging.AdjustHue(result, opts.Hue)
	// }

	// Adjust sharpness
	if opts.Sharpness != 0 {
		result = imaging.Sharpen(result, opts.Sharpness)
	}

	return result
}

// GeneratePlaceholderImage generates a placeholder image.
func (m *Manager) GeneratePlaceholderImage(width, height int, bgColor string) ([]byte, error) {
	// Parse background color
	c := color.NRGBA{128, 128, 128, 255}
	switch bgColor {
	case "white":
		c = color.NRGBA{255, 255, 255, 255}
	case "black":
		c = color.NRGBA{0, 0, 0, 255}
	case "gray":
		c = color.NRGBA{128, 128, 128, 255}
	}

	// Create image
	img := image.NewNRGBA(image.Rect(0, 0, width, height))
	draw.Draw(img, img.Bounds(), &image.Uniform{c}, image.Point{}, draw.Src)

	// Encode as PNG
	var buf []byte
	if err := png.Encode(&bufferWriter{buf: &buf}, img); err != nil {
		return nil, fmt.Errorf("encode placeholder: %w", err)
	}

	return buf, nil
}

// GetCachedImage returns a cached image.
func (m *Manager) GetCachedImage(key string) (*CachedImage, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	img, exists := m.cache[key]
	if exists {
		img.LastAccess = time.Now()
	}
	return img, exists
}

// CacheImage caches an image.
func (m *Manager) CacheImage(key string, img *CachedImage) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if len(m.cache) >= m.maxCache {
		// Evict oldest
		m.evictOldest()
	}

	m.cache[key] = img
}

// evictOldest removes the oldest cached image.
func (m *Manager) evictOldest() {
	var oldestKey string
	var oldestTime time.Time
	first := true

	for key, img := range m.cache {
		if first || img.LastAccess.Before(oldestTime) {
			oldestKey = key
			oldestTime = img.LastAccess
			first = false
		}
	}

	if !first {
		delete(m.cache, oldestKey)
	}
}

// ImageOptions holds image processing options.
type ImageOptions struct {
	Width        int
	Height       int
	Crop         bool
	Rotate       float64
	Quality      int
	OutputFormat string
	Brightness   float64
	Contrast     float64
	Saturation   float64
	Hue          float64
	Sharpness    float64
}

// bufferWriter is a simple writer that appends to a byte slice.
type bufferWriter struct {
	buf *[]byte
}

func (w *bufferWriter) Write(p []byte) (int, error) {
	*w.buf = append(*w.buf, p...)
	return len(p), nil
}
