package library

import (
	"testing"

	"github.com/emby/emby-go/internal/config"
	"github.com/emby/emby-go/internal/database"
	"github.com/emby/emby-go/internal/repository"
	"go.uber.org/zap"
)

func TestGenerateItemID(t *testing.T) {
	id1 := generateItemID("/path/to/file.mp4")
	id2 := generateItemID("/path/to/file.mp4")

	if id1 == "" {
		t.Error("generateItemID returned empty string")
	}
	if id1 == id2 {
		t.Error("generateItemID should generate unique IDs")
	}
}

func TestScanResult(t *testing.T) {
	result := &ScanResult{
		TotalItemsFound: 10,
		NewItems:        5,
		UpdatedItems:    3,
		RemovedItems:    2,
		Errors:          0,
	}

	if result.TotalItemsFound != 10 {
		t.Errorf("expected TotalItemsFound 10, got %d", result.TotalItemsFound)
	}
	if result.NewItems != 5 {
		t.Errorf("expected NewItems 5, got %d", result.NewItems)
	}
}

func TestScanner_NewScanner(t *testing.T) {
	cfg := &config.Config{
		Library: config.LibraryConfig{
			ContentTypes: []string{"movies", "tvshows"},
		},
	}

	logger := zap.NewNop()
	dbMgr, err := database.NewManager(&config.DatabaseConfig{Path: ":memory:", MaxOpenConns: 1, MaxIdleConns: 1})
	if err != nil {
		t.Fatalf("failed to create db manager: %v", err)
	}
	defer dbMgr.Close()

	repo := repository.NewItemRepository(dbMgr.DB())
	if err := repo.CreateSchema(); err != nil {
		t.Fatalf("failed to create schema: %v", err)
	}

	scanner := NewScanner(cfg, logger, repo)
	if scanner == nil {
		t.Fatal("NewScanner returned nil")
	}
}

func TestScanner_Paths(t *testing.T) {
	cfg := &config.Config{
		Library: config.LibraryConfig{
			ContentTypes: []string{"movies"},
		},
	}

	logger := zap.NewNop()
	dbMgr, err := database.NewManager(&config.DatabaseConfig{Path: ":memory:", MaxOpenConns: 1, MaxIdleConns: 1})
	if err != nil {
		t.Fatalf("failed to create db manager: %v", err)
	}
	defer dbMgr.Close()

	repo := repository.NewItemRepository(dbMgr.DB())
	repo.CreateSchema()

	scanner := NewScanner(cfg, logger, repo)

	initialLen := len(scanner.paths)
	scanner.AddPath("/new/path")
	if len(scanner.paths) != initialLen+1 {
		t.Errorf("expected %d paths, got %d", initialLen+1, len(scanner.paths))
	}

	scanner.RemovePath("/new/path")
	if len(scanner.paths) != initialLen {
		t.Errorf("expected %d paths after remove, got %d", initialLen, len(scanner.paths))
	}
}

func TestScanner_RemovePath_NotFound(t *testing.T) {
	cfg := &config.Config{
		Library: config.LibraryConfig{
			ContentTypes: []string{},
		},
	}

	logger := zap.NewNop()
	dbMgr, err := database.NewManager(&config.DatabaseConfig{Path: ":memory:", MaxOpenConns: 1, MaxIdleConns: 1})
	if err != nil {
		t.Fatalf("failed to create db manager: %v", err)
	}
	defer dbMgr.Close()

	repo := repository.NewItemRepository(dbMgr.DB())
	repo.CreateSchema()

	scanner := NewScanner(cfg, logger, repo)
	initialLen := len(scanner.paths)

	scanner.RemovePath("/nonexistent/path")
	if len(scanner.paths) != initialLen {
		t.Errorf("expected %d paths, got %d", initialLen, len(scanner.paths))
	}
}

func TestScanner_IsMediaFile(t *testing.T) {
	cfg := &config.Config{
		Library: config.LibraryConfig{
			ContentTypes: []string{},
		},
	}

	logger := zap.NewNop()
	dbMgr, err := database.NewManager(&config.DatabaseConfig{Path: ":memory:", MaxOpenConns: 1, MaxIdleConns: 1})
	if err != nil {
		t.Fatalf("failed to create db manager: %v", err)
	}
	defer dbMgr.Close()

	repo := repository.NewItemRepository(dbMgr.DB())
	repo.CreateSchema()

	scanner := NewScanner(cfg, logger, repo)

	tests := []struct {
		path     string
		expected bool
	}{
		{"/path/to/movie.mp4", true},
		{"/path/to/movie.mkv", true},
		{"/path/to/movie.avi", true},
		{"/path/to/song.mp3", true},
		{"/path/to/song.flac", true},
		{"/path/to/photo.jpg", true},
		{"/path/to/photo.png", true},
		{"/path/to/book.epub", true},
		{"/path/to/subtitle.srt", true},
		{"/path/to/subtitle.vtt", true},
		{"/path/to/file.txt", true},
		{"/path/to/file.exe", false},
		{"/path/to/file.go", false},
	}

	for _, tc := range tests {
		result := scanner.isMediaFile(tc.path)
		if result != tc.expected {
			t.Errorf("isMediaFile(%s) = %v, expected %v", tc.path, result, tc.expected)
		}
	}
}

func TestScanner_GetMediaFolderPath(t *testing.T) {
	cfg := &config.Config{
		Library: config.LibraryConfig{
			ContentTypes: []string{},
		},
	}

	logger := zap.NewNop()
	dbMgr, err := database.NewManager(&config.DatabaseConfig{Path: ":memory:", MaxOpenConns: 1, MaxIdleConns: 1})
	if err != nil {
		t.Fatalf("failed to create db manager: %v", err)
	}
	defer dbMgr.Close()

	repo := repository.NewItemRepository(dbMgr.DB())
	repo.CreateSchema()

	scanner := NewScanner(cfg, logger, repo)

	tests := []struct {
		mediaType string
		expected  string
	}{
		{"video", "media/movies"},
		{"Video", "media/movies"},
		{"VIDEO", "media/movies"},
		{"music", "media/music"},
		{"Music", "media/music"},
		{"photos", "media/photos"},
		{"Photos", "media/photos"},
		{"books", "media/books"},
		{"Books", "media/books"},
		{"unknown", "media/unknown"},
	}

	for _, tc := range tests {
		result := scanner.getMediaFolderPath(tc.mediaType)
		if result != tc.expected {
			t.Errorf("getMediaFolderPath(%s) = %s, expected %s", tc.mediaType, result, tc.expected)
		}
	}
}