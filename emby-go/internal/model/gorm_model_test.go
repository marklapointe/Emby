package model

import (
	"testing"
	"time"
)

func TestGORMItemTableName(t *testing.T) {
	item := GORMItem{}
	if item.TableName() != "Items" {
		t.Errorf("expected Items, got %s", item.TableName())
	}
}

func TestGORMMediaSourceTableName(t *testing.T) {
	ms := GORMMediaSource{}
	if ms.TableName() != "MediaSources" {
		t.Errorf("expected MediaSources, got %s", ms.TableName())
	}
}

func TestGORMUserTableName(t *testing.T) {
	user := GORMUser{}
	if user.TableName() != "Users" {
		t.Errorf("expected Users, got %s", user.TableName())
	}
}

func TestGORMUserItemTableName(t *testing.T) {
	ui := GORMUserItem{}
	if ui.TableName() != "UserItems" {
		t.Errorf("expected UserItems, got %s", ui.TableName())
	}
}

func TestGORMSessionTableName(t *testing.T) {
	sess := GORMSession{}
	if sess.TableName() != "Sessions" {
		t.Errorf("expected Sessions, got %s", sess.TableName())
	}
}

func TestGORMItemFields(t *testing.T) {
	item := GORMItem{
		Id:              "test-id",
		Name:            "Test Movie",
		Overview:        "A test overview",
		RunTimeTicks:    7200000000,
		ProductionYear:  2024,
		MediaType:       "Video",
		Path:            "/media/test.mp4",
		IsMovie:         true,
		CreatedDate:     time.Now(),
	}
	if item.Id != "test-id" {
		t.Errorf("expected Id test-id, got %s", item.Id)
	}
	if item.Name != "Test Movie" {
		t.Errorf("expected Name Test Movie, got %s", item.Name)
	}
	if !item.IsMovie {
		t.Error("expected IsMovie to be true")
	}
}

func TestGORMMediaSourceFields(t *testing.T) {
	ms := GORMMediaSource{
		Id:          "ms-1",
		ItemId:      "item-1",
		Name:        "Main",
		Container:   "mkv",
		Path:        "/media/test.mkv",
		Protocol:    "File",
		VideoCodec:  "h264",
		AudioCodec:  "aac",
		Size:        1024000,
	}
	if ms.Id != "ms-1" {
		t.Errorf("expected Id ms-1, got %s", ms.Id)
	}
	if ms.Container != "mkv" {
		t.Errorf("expected Container mkv, got %s", ms.Container)
	}
}

func TestGORMUserFields(t *testing.T) {
	user := GORMUser{
		Id:        "user-1",
		Name:      "Test User",
		EmailAddress: "test@example.com",
	}
	if user.Id != "user-1" {
		t.Errorf("expected Id user-1, got %s", user.Id)
	}
	if user.Name != "Test User" {
		t.Errorf("expected Name Test User, got %s", user.Name)
	}
}

func TestGORMUserItemFields(t *testing.T) {
	ui := GORMUserItem{
		Id:                  "ui-1",
		UserId:              "user-1",
		ItemID:              "item-1",
		PlaybackPositionTicks: 3600000000,
		PlayCount:           3,
		IsFavorite:          true,
		Played:              true,
	}
	if ui.Id != "ui-1" {
		t.Errorf("expected Id ui-1, got %s", ui.Id)
	}
	if ui.PlayCount != 3 {
		t.Errorf("expected PlayCount 3, got %d", ui.PlayCount)
	}
}

func TestGORMSessionFields(t *testing.T) {
	sess := GORMSession{
		Id:          "sess-1",
		Client:      "Emby iOS",
		DeviceName:  "iPhone 15",
		DisplayName: "John's iPhone",
		MachineId:   "abc123",
	}
	if sess.Id != "sess-1" {
		t.Errorf("expected Id sess-1, got %s", sess.Id)
	}
	if sess.DeviceName != "iPhone 15" {
		t.Errorf("expected DeviceName iPhone 15, got %s", sess.DeviceName)
	}
}