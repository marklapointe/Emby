package repository

import (
	"testing"

	"github.com/emby/emby-go/internal/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func createMemoryDB() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
}

func setupTestDB(t *testing.T) (*gorm.DB, func()) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to create db: %v", err)
	}
	db.AutoMigrate(
		&model.GORMItem{},
		&model.GORMMediaSource{},
		&model.GORMUser{},
		&model.GORMUserItem{},
		&model.GORMSession{},
	)
	cleanup := func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}
	return db, cleanup
}