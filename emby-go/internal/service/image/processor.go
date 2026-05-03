package image

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// Processor handles image processing operations.
type Processor struct {
	cacheDir string
	mu       sync.RWMutex
	cache    map[string]*CachedImage
}

// CachedImage represents a cached image.
type CachedImage struct {
	Data       []byte
	Format     string
	Width      int
	Height     int
	LastAccess time.Time
}

// NewProcessor creates a new image processor.
func NewProcessor(cacheDir string) *Processor {
	return &Processor{
		cacheDir: cacheDir,
		cache:    make(map[string]*CachedImage),
	}
}

// ResizeImage resizes an image to the specified dimensions.
func (p *Processor) ResizeImage(src image.Image, width, height int) (image.Image, error) {
	// For now, return the original image
	_ = width
	_ = height
	return src, nil
}

// ConvertFormat converts an image to the specified format.
func (p *Processor) ConvertFormat(src image.Image, format string) ([]byte, error) {
	var buf []byte
	var err error

	switch format {
	case "jpeg":
		buf = make([]byte, 0)
		writer := &bytesWriter{buf: &buf}
		err = jpeg.Encode(writer, src, &jpeg.Options{Quality: 85})
	case "png":
		buf = make([]byte, 0)
		writer := &bytesWriter{buf: &buf}
		err = png.Encode(writer, src)
	default:
		return nil, fmt.Errorf("unsupported format: %s", format)
	}

	return buf, err
}

// GetImageFormat returns the format of an image.
func GetImageFormat(data []byte) string {
	if len(data) >= 3 && data[0] == 0xFF && data[1] == 0xD8 && data[2] == 0xFF {
		return "jpeg"
	}
	if len(data) >= 8 && data[0] == 0x89 && data[1] == 0x50 && data[2] == 0x4E && data[3] == 0x47 {
		return "png"
	}
	return "unknown"
}

// CacheImage caches an image.
func (p *Processor) CacheImage(key string, data []byte, format string, width, height int) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.cache[key] = &CachedImage{
		Data:       data,
		Format:     format,
		Width:      width,
		Height:     height,
		LastAccess: time.Now(),
	}

	return nil
}

// GetCachedImage retrieves a cached image.
func (p *Processor) GetCachedImage(key string) (*CachedImage, error) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	image, ok := p.cache[key]
	if !ok {
		return nil, fmt.Errorf("image not found in cache: %s", key)
	}

	image.LastAccess = time.Now()
	return image, nil
}

// ClearCache clears the image cache.
func (p *Processor) ClearCache() error {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.cache = make(map[string]*CachedImage)
	return os.RemoveAll(p.cacheDir)
}

// bytesWriter is a helper type for writing to a byte slice.
type bytesWriter struct {
	buf *[]byte
}

func (w *bytesWriter) Write(p []byte) (n int, err error) {
	*w.buf = append(*w.buf, p...)
	return len(p), nil
}

// GetImageDimensions returns the dimensions of an image.
func GetImageDimensions(data []byte) (int, int, error) {
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return 0, 0, err
	}
	bounds := img.Bounds()
	return bounds.Dx(), bounds.Dy(), nil
}

// SaveImage saves an image to a file.
func SaveImage(data []byte, path string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	out, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer out.Close()

	_, err = io.Copy(out, bytes.NewReader(data))
	return err
}
