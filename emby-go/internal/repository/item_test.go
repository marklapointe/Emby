package repository

import (
	"database/sql"
	"testing"

	_ "modernc.org/sqlite"
)

func TestCreateSchema(t *testing.T) {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatalf("failed to open database: %v", err)
	}
	defer db.Close()
	
	repo := NewItemRepository(db)
	
	if err := repo.CreateSchema(); err != nil {
		t.Fatalf("failed to create schema: %v", err)
	}
}

func TestInsertAndGetItem(t *testing.T) {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatalf("failed to open database: %v", err)
	}
	defer db.Close()
	
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
	
	// Verify item was inserted by querying directly
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM Items WHERE Id = ?", itemID).Scan(&count)
	if err != nil {
		t.Fatalf("failed to query item: %v", err)
	}
	
	if count != 1 {
		t.Errorf("expected 1 item, got %d", count)
	}
}

func TestSearchItems(t *testing.T) {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatalf("failed to open database: %v", err)
	}
	defer db.Close()
	
	repo := NewItemRepository(db)
	
	if err := repo.CreateSchema(); err != nil {
		t.Fatalf("failed to create schema: %v", err)
	}
	
	// Insert test items
	testItems := []struct {
		id     string
		name   string
		path   string
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
	
	// Search for items
	results, err := repo.SearchItems("Movie", 10, 0)
	if err != nil {
		t.Fatalf("failed to search items: %v", err)
	}
	
	if len(results) != 2 {
		t.Errorf("expected 2 results, got %d", len(results))
	}
}

func TestGetTotalItemCounts(t *testing.T) {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatalf("failed to open database: %v", err)
	}
	defer db.Close()
	
	repo := NewItemRepository(db)
	
	if err := repo.CreateSchema(); err != nil {
		t.Fatalf("failed to create schema: %v", err)
	}
	
	// Insert test items
	testItems := []struct {
		id     string
		name   string
		path   string
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
