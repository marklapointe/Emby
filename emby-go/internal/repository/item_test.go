package repository

import (
	"testing"

	"github.com/emby/emby-go/internal/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupGORM(t *testing.T) (*gorm.DB, func()) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open database: %v", err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(1)
	cleanup := func() {
		sqlDB.Close()
	}
	return db, cleanup
}

func TestCreateSchema(t *testing.T) {
	db, cleanup := setupGORM(t)
	defer cleanup()

	repo := NewItemRepository(db)

	if err := repo.CreateSchema(); err != nil {
		t.Fatalf("failed to create schema: %v", err)
	}
}

func TestInsertAndGetItem(t *testing.T) {
	db, cleanup := setupGORM(t)
	defer cleanup()

	repo := NewItemRepository(db)

	if err := repo.CreateSchema(); err != nil {
		t.Fatalf("failed to create schema: %v", err)
	}

	itemID := "test-item-1"
	itemName := "Test Movie"
	itemPath := "/media/movies/test.mp4"
	itemType := "Video"

	if err := repo.InsertItem(itemID, itemName, itemPath, itemType); err != nil {
		t.Fatalf("failed to insert item: %v", err)
	}

	var count int64
	err := db.Model(&model.GORMItem{}).Where("Id = ?", itemID).Count(&count).Error
	if err != nil {
		t.Fatalf("failed to query item: %v", err)
	}

	if count != 1 {
		t.Errorf("expected 1 item, got %d", count)
	}
}

func TestSearchItems(t *testing.T) {
	db, cleanup := setupGORM(t)
	defer cleanup()

	repo := NewItemRepository(db)

	if err := repo.CreateSchema(); err != nil {
		t.Fatalf("failed to create schema: %v", err)
	}

	testItems := []struct {
		id       string
		name     string
		path     string
		mediaType string
	}{
		{"item-1", "Movie A", "/media/movies/a.mp4", "Video"},
		{"item-2", "Movie B", "/media/movies/b.mp4", "Video"},
		{"item-3", "Song A", "/media/music/a.mp3", "Music"},
	}

	for _, item := range testItems {
		if err := repo.InsertItem(item.id, item.name, item.path, item.mediaType); err != nil {
			t.Fatalf("failed to insert item %s: %v", item.id, err)
		}
	}

	results, err := repo.SearchItems("Movie", 10, 0)
	if err != nil {
		t.Fatalf("failed to search items: %v", err)
	}

	if len(results) != 2 {
		t.Errorf("expected 2 results, got %d", len(results))
	}
}

func TestGetTotalItemCounts(t *testing.T) {
	db, cleanup := setupGORM(t)
	defer cleanup()

	repo := NewItemRepository(db)

	if err := repo.CreateSchema(); err != nil {
		t.Fatalf("failed to create schema: %v", err)
	}

	testItems := []struct {
		id       string
		name     string
		path     string
		mediaType string
	}{
		{"item-1", "Movie A", "/media/movies/a.mp4", "Video"},
		{"item-2", "Movie B", "/media/movies/b.mp4", "Video"},
		{"item-3", "Song A", "/media/music/a.mp3", "Music"},
	}

	for _, item := range testItems {
		if err := repo.InsertItem(item.id, item.name, item.path, item.mediaType); err != nil {
			t.Fatalf("failed to insert item %s: %v", item.id, err)
		}
	}

	counts, err := repo.GetTotalItemCounts()
	if err != nil {
		t.Fatalf("failed to get item counts: %v", err)
	}

	if counts["Video"] != 2 {
		t.Errorf("expected 2 Video items, got %d", counts["Video"])
	}

	if counts["Music"] != 1 {
		t.Errorf("expected 1 Music item, got %d", counts["Music"])
	}
}

func TestBaseRepository_DB(t *testing.T) {
	db, cleanup := setupGORM(t)
	defer cleanup()

	repo := NewBaseRepository(db)
	if repo.DB() == nil {
		t.Error("DB() returned nil")
	}
}

func TestBaseRepository_SQL(t *testing.T) {
	db, cleanup := setupGORM(t)
	defer cleanup()

	repo := NewBaseRepository(db)
	sqlDB := repo.SQL()
	if sqlDB == nil {
		t.Error("SQL() returned nil")
	}
}

func TestBaseRepository_Ping(t *testing.T) {
	db, cleanup := setupGORM(t)
	defer cleanup()

	repo := NewBaseRepository(db)
	err := repo.Ping()
	if err != nil {
		t.Errorf("Ping failed: %v", err)
	}
}

func TestItemRepository_GetAllItems(t *testing.T) {
	db, cleanup := setupGORM(t)
	defer cleanup()

	repo := NewItemRepository(db)
	if err := repo.CreateSchema(); err != nil {
		t.Fatalf("failed to create schema: %v", err)
	}

	repo.InsertItem("all1", "All Items 1", "/path1", "Video")
	repo.InsertItem("all2", "All Items 2", "/path2", "Music")

	items, err := repo.GetAllItems()
	if err != nil {
		t.Fatalf("GetAllItems failed: %v", err)
	}
	if len(items) != 2 {
		t.Errorf("expected 2 items, got %d", len(items))
	}
}