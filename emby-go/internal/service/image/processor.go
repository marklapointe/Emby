package image

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"io"
	"math"
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

// GenerateBlurHash generates a blurhash string for an image.
// BlurHash is a compact representation of a placeholder image.
func (p *Processor) GenerateBlurHash(img image.Image, componentsX, componentsY int) (string, error) {
	if componentsX < 1 || componentsX > 9 || componentsY < 1 || componentsY > 9 {
		return "", fmt.Errorf("components must be between 1 and 9")
	}

	width := img.Bounds().Dx()
	height := img.Bounds().Dy()

	// Calculate scaling factors
	scaleX := float64(width) / float64(componentsX*32)
	scaleY := float64(height) / float64(componentsY*32)

	// Downsample image
	sampled := make([]float64, componentsX*componentsY*3)
	for y := 0; y < componentsY*32; y++ {
		for x := 0; x < componentsX*32; x++ {
			srcX := int(float64(x) * scaleX)
			srcY := int(float64(y) * scaleY)
			if srcX >= width {
				srcX = width - 1
			}
			if srcY >= height {
				srcY = height - 1
			}
			r, g, b, _ := img.At(srcX, srcY).RGBA()
			// Convert to 0-1 range
			factor := float64(y*componentsX*32 + x)
			sampled[0] += float64(r>>8) * math.Cos(math.Pi*float64(x)*factor/float64(componentsX)) * math.Cos(math.Pi*float64(y)*factor/float64(componentsY))
			sampled[1] += float64(g>>8) * math.Cos(math.Pi*float64(x)*factor/float64(componentsX)) * math.Cos(math.Pi*float64(y)*factor/float64(componentsY))
			sampled[2] += float64(b>>8) * math.Cos(math.Pi*float64(x)*factor/float64(componentsX)) * math.Cos(math.Pi*float64(y)*factor/float64(componentsY))
		}
	}

	// Calculate DC component (average color)
	dc := [3]float64{}
	for i := 0; i < 3; i++ {
		dc[i] = sampled[i] / float64(width*height)
	}

	// Calculate AC components
	var ac [][3]float64
	for c := 0; c < componentsX*componentsY-1; c++ {
		ac = append(ac, [3]float64{
			sampled[3*(c+1)] / float64(width*height),
			sampled[3*(c+1)+1] / float64(width*height),
			sampled[3*(c+1)+2] / float64(width*height),
		})
	}

	// Encode to base83
	return encodeBlurHash(dc, ac, componentsX, componentsY)
}

func encodeBlurHash(dc [3]float64, ac [][3]float64, width, height int) (string, error) {
	// Base83 encoding characters
	base83Chars := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz#$%*+,-.:;=?@[]^_{|}~"

	// Calculate size flag
	sizeFlag := (width - 1) + (height-1)*9

	// Encode DC component
	dcValue := int(math.Round(dc[0]))<<16 | int(math.Round(dc[1]))<<8 | int(math.Round(dc[2]))

	// Build result
	var result bytes.Buffer
	result.WriteByte(base83Chars[sizeFlag/83])
	result.WriteByte(base83Chars[sizeFlag%83])

	// Encode DC
	result.WriteByte(base83Chars[dcValue/(83*83)])
	result.WriteByte(base83Chars[(dcValue/83)%83])
	result.WriteByte(base83Chars[dcValue%83])

	// Encode AC components
	for _, v := range ac {
		for i := 0; i < 3; i++ {
			value := int(math.Round(signPow(v[i]/float64(len(ac)), 0.5) * 166 + 9))
			if value < 0 {
				value = 0
			}
			if value > 242 {
				value = 242
			}
			result.WriteByte(base83Chars[value/83])
			result.WriteByte(base83Chars[value%83])
		}
	}

	return result.String(), nil
}

func signPow(val, exp float64) float64 {
	sign := 1.0
	if val < 0 {
		sign = -1
	}
	return sign * math.Pow(math.Abs(val), exp)
}

// DecodeBlurHash decodes a blurhash string back to an image.
func (p *Processor) DecodeBlurHash(hash string, width, height int) (image.Image, error) {
	if len(hash) < 6 {
		return nil, fmt.Errorf("blurhash too short")
	}

	// Parse size flag
	sizeFlag := decodeBase83(string(hash[0]))
	widthParam := sizeFlag%9 + 1
	heightParam := sizeFlag/9 + 1

	// Parse DC component
	dcValue := decodeBase83(string(hash[1]))*83*83 + decodeBase83(string(hash[2]))*83 + decodeBase83(string(hash[3]))
	dc := [3]float64{
		float64(dcValue >> 16),
		float64((dcValue >> 8) & 0xFF),
		float64(dcValue & 0xFF),
	}

	// Parse AC components
	ac := make([][3]float64, widthParam*heightParam-1)
	for i := 0; i < len(ac); i++ {
		value := decodeBase83(string(hash[4+i*2]))*83 + decodeBase83(string(hash[5+i*2]))
		ac[i] = [3]float64{
			signPow(float64(value/361)-9, 2) / 166,
			signPow(float64((value/19)%19)-9, 2) / 166,
			signPow(float64(value%19)-9, 2) / 166,
		}
	}

	// Create image
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			var r, g, b float64

			// Add DC component
			r += dc[0]
			g += dc[1]
			b += dc[2]

			// Add AC components
			for j := 0; j < heightParam; j++ {
				for i := 0; i < widthParam; i++ {
					if i+j == 0 {
						continue
					}
					freqX := math.Cos(math.Pi * float64(i) * float64(x) / float64(width))
					freqY := math.Cos(math.Pi * float64(j) * float64(y) / float64(height))
					idx := i + j*widthParam - 1
					if idx < len(ac) {
						r += ac[idx][0] * freqX * freqY
						g += ac[idx][1] * freqX * freqY
						b += ac[idx][2] * freqX * freqY
					}
				}
			}

			img.Set(x, y, color.RGBA{
				R: uint8(math.Max(0, math.Min(255, r))),
				G: uint8(math.Max(0, math.Min(255, g))),
				B: uint8(math.Max(0, math.Min(255, b))),
				A: 255,
			})
		}
	}

	return img, nil
}

func decodeBase83(s string) int {
	base83Chars := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz#$%*+,-.:;=?@[]^_{|}~"
	val := 0
	for _, c := range s {
		idx := bytes.IndexByte([]byte(base83Chars), byte(c))
		if idx == -1 {
			return 0
		}
		val = val*83 + idx
	}
	return val
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
