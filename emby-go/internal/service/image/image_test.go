package image

import (
	"bytes"
	"image"
	"image/png"
	"testing"
)

func TestNewProcessor(t *testing.T) {
	p := NewProcessor("/tmp/cache")
	if p == nil {
		t.Fatal("NewProcessor returned nil")
	}
	if p.cache == nil {
		t.Error("cache map not initialized")
	}
}

func TestGetImageFormat_JPEG(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 1, 1))
	var buf bytes.Buffer
	png.Encode(&buf, img)

	data := []byte{0xFF, 0xD8, 0xFF, 0xE0}
	format := GetImageFormat(data)
	if format != "jpeg" {
		t.Errorf("expected 'jpeg', got '%s'", format)
	}
}

func TestGetImageFormat_PNG(t *testing.T) {
	data := []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}
	format := GetImageFormat(data)
	if format != "png" {
		t.Errorf("expected 'png', got '%s'", format)
	}
}

func TestGetImageFormat_Unknown(t *testing.T) {
	data := []byte{0x00, 0x00, 0x00, 0x00}
	format := GetImageFormat(data)
	if format != "unknown" {
		t.Errorf("expected 'unknown', got '%s'", format)
	}
}

func TestGetImageDimensions(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 100, 200))
	var buf bytes.Buffer
	png.Encode(&buf, img)

	width, height, err := GetImageDimensions(buf.Bytes())
	if err != nil {
		t.Fatalf("GetImageDimensions returned error: %v", err)
	}
	if width != 100 {
		t.Errorf("expected width 100, got %d", width)
	}
	if height != 200 {
		t.Errorf("expected height 200, got %d", height)
	}
}

func TestResizeImage(t *testing.T) {
	p := NewProcessor("/tmp/cache")
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))

	resized, err := p.ResizeImage(img, 50, 50)
	if err != nil {
		t.Fatalf("ResizeImage returned error: %v", err)
	}
	if resized == nil {
		t.Fatal("resized image is nil")
	}
}

func TestConvertFormat_JPEG(t *testing.T) {
	p := NewProcessor("/tmp/cache")
	img := image.NewRGBA(image.Rect(0, 0, 10, 10))

	data, err := p.ConvertFormat(img, "jpeg")
	if err != nil {
		t.Fatalf("ConvertFormat returned error: %v", err)
	}
	if len(data) == 0 {
		t.Error("expected non-empty data")
	}
}

func TestConvertFormat_PNG(t *testing.T) {
	p := NewProcessor("/tmp/cache")
	img := image.NewRGBA(image.Rect(0, 0, 10, 10))

	data, err := p.ConvertFormat(img, "png")
	if err != nil {
		t.Fatalf("ConvertFormat returned error: %v", err)
	}
	if len(data) == 0 {
		t.Error("expected non-empty data")
	}
}

func TestConvertFormat_Unsupported(t *testing.T) {
	p := NewProcessor("/tmp/cache")
	img := image.NewRGBA(image.Rect(0, 0, 10, 10))

	_, err := p.ConvertFormat(img, "unsupported")
	if err == nil {
		t.Error("expected error for unsupported format")
	}
}

func TestCacheImage(t *testing.T) {
	p := NewProcessor("/tmp/cache")
	data := []byte{0x89, 0x50, 0x4E, 0x47}

	err := p.CacheImage("test-key", data, "png", 100, 200)
	if err != nil {
		t.Fatalf("CacheImage returned error: %v", err)
	}

	cached, err := p.GetCachedImage("test-key")
	if err != nil {
		t.Fatalf("GetCachedImage returned error: %v", err)
	}
	if cached.Width != 100 {
		t.Errorf("expected width 100, got %d", cached.Width)
	}
}

func TestGetCachedImage_NotFound(t *testing.T) {
	p := NewProcessor("/tmp/cache")
	_, err := p.GetCachedImage("non-existent")
	if err == nil {
		t.Error("expected error for non-existent key")
	}
}

func TestClearCache(t *testing.T) {
	p := NewProcessor("/tmp/cache")
	p.CacheImage("key1", []byte{0x00}, "png", 10, 10)
	p.CacheImage("key2", []byte{0x00}, "png", 10, 10)

	err := p.ClearCache()
	if err != nil {
		t.Errorf("ClearCache returned error: %v", err)
	}

	_, err = p.GetCachedImage("key1")
	if err == nil {
		t.Error("expected error after cache clear")
	}
}

func TestGenerateBlurHash(t *testing.T) {
	p := NewProcessor("/tmp/cache")
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))

	hash, err := p.GenerateBlurHash(img, 4, 4)
	if err != nil {
		t.Fatalf("GenerateBlurHash returned error: %v", err)
	}
	if hash == "" {
		t.Error("hash is empty")
	}
}

func TestGenerateBlurHash_InvalidComponents(t *testing.T) {
	p := NewProcessor("/tmp/cache")
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))

	_, err := p.GenerateBlurHash(img, 0, 4)
	if err == nil {
		t.Error("expected error for invalid components")
	}

	_, err = p.GenerateBlurHash(img, 10, 4)
	if err == nil {
		t.Error("expected error for component > 9")
	}
}

func TestDecodeBlurHash(t *testing.T) {
	p := NewProcessor("/tmp/cache")
	img := image.NewRGBA(image.Rect(0, 0, 32, 32))

	hash, err := p.GenerateBlurHash(img, 4, 4)
	if err != nil {
		t.Fatalf("GenerateBlurHash failed: %v", err)
	}

	decoded, err := p.DecodeBlurHash(hash, 32, 32)
	if err != nil {
		t.Fatalf("DecodeBlurHash returned error: %v", err)
	}
	if decoded == nil {
		t.Fatal("decoded image is nil")
	}
}

func TestDecodeBlurHash_TooShort(t *testing.T) {
	p := NewProcessor("/tmp/cache")
	_, err := p.DecodeBlurHash("abc", 32, 32)
	if err == nil {
		t.Error("expected error for too short hash")
	}
}

func TestSaveImage(t *testing.T) {
	data := []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}
	err := SaveImage(data, "/tmp/test-image.png")
	if err != nil {
		t.Errorf("SaveImage returned error: %v", err)
	}
}

func TestSaveImage_CreateDir(t *testing.T) {
	data := []byte{0x89, 0x50, 0x4E, 0x47}
	err := SaveImage(data, "/tmp/subdir/test-image.png")
	if err != nil {
		t.Errorf("SaveImage returned error: %v", err)
	}
}

func TestRotateImage(t *testing.T) {
	p := NewProcessor("/tmp/cache")
	img := image.NewRGBA(image.Rect(0, 0, 100, 50))

	rotated, err := p.RotateImage(img, 90)
	if err != nil {
		t.Fatalf("RotateImage(90) returned error: %v", err)
	}
	bounds := rotated.Bounds()
	if bounds.Dx() != 50 || bounds.Dy() != 100 {
		t.Errorf("expected 50x100, got %dx%d", bounds.Dx(), bounds.Dy())
	}

	rotated180, err := p.RotateImage(img, 180)
	if err != nil {
		t.Fatalf("RotateImage(180) returned error: %v", err)
	}
	bounds180 := rotated180.Bounds()
	if bounds180.Dx() != 100 || bounds180.Dy() != 50 {
		t.Errorf("expected 100x50, got %dx%d", bounds180.Dx(), bounds180.Dy())
	}

	rotated270, err := p.RotateImage(img, 270)
	if err != nil {
		t.Fatalf("RotateImage(270) returned error: %v", err)
	}
	bounds270 := rotated270.Bounds()
	if bounds270.Dx() != 50 || bounds270.Dy() != 100 {
		t.Errorf("expected 50x100, got %dx%d", bounds270.Dx(), bounds270.Dy())
	}

	rotated0, err := p.RotateImage(img, 0)
	if err != nil {
		t.Fatalf("RotateImage(0) returned error: %v", err)
	}
	bounds0 := rotated0.Bounds()
	if bounds0.Dx() != 100 || bounds0.Dy() != 50 {
		t.Errorf("expected 100x50, got %dx%d", bounds0.Dx(), bounds0.Dy())
	}
}

func TestRotateImage_UnsupportedAngle(t *testing.T) {
	p := NewProcessor("/tmp/cache")
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))

	_, err := p.RotateImage(img, 45)
	if err == nil {
		t.Error("expected error for unsupported angle 45")
	}
}

func TestManagerGetImageResize_Processes(t *testing.T) {
	m := NewManager(nil, "/tmp/image-cache")

	proc := NewProcessor("/tmp/image-cache")
	m.SetProcessor(proc)

	m.images["item1"] = []*ImageInfo{
		{Type: ImageTypePrimary, Path: "/tmp/test-image.png"},
	}

	data, contentType, err := m.GetImageResize("item1", ImageTypePrimary, 50, 50, 85)
	if err != nil {
		t.Fatalf("GetImageResize returned error: %v", err)
	}
	if len(data) == 0 {
		t.Error("expected processed image data")
	}
	if contentType != "image/jpeg" && contentType != "image/png" {
		t.Errorf("unexpected content type: %s", contentType)
	}
}

func TestManagerGetImageResize_NoProcessor(t *testing.T) {
	m := NewManager(nil, "/tmp/image-cache")
	m.images["item1"] = []*ImageInfo{
		{Type: ImageTypePrimary, Path: "/tmp/nonexistent.png"},
	}

	data, _, err := m.GetImageResize("item1", ImageTypePrimary, 50, 50, 85)
	if err == nil {
		t.Error("expected error for nonexistent file")
	}
	_ = data
}