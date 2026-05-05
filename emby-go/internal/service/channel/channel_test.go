package channel

import (
	"testing"
)

func TestNewManager(t *testing.T) {
	m := NewManager(nil)
	if m == nil {
		t.Fatal("NewManager returned nil")
	}
	if m.channels == nil {
		t.Error("channels map not initialized")
	}
}

func TestGetChannels_Empty(t *testing.T) {
	m := NewManager(nil)
	channels := m.GetChannels()
	if len(channels) != 0 {
		t.Errorf("expected 0 channels, got %d", len(channels))
	}
}

func TestAddChannel(t *testing.T) {
	m := NewManager(nil)

	ch := &Channel{
		ID:   "ch-1",
		Name: "Test Channel",
	}

	m.AddChannel(ch)

	channels := m.GetChannels()
	if len(channels) != 1 {
		t.Errorf("expected 1 channel, got %d", len(channels))
	}
}

func TestGetChannel(t *testing.T) {
	m := NewManager(nil)

	ch := &Channel{
		ID:   "ch-1",
		Name: "Test Channel",
	}
	m.AddChannel(ch)

	got, ok := m.GetChannel("ch-1")
	if !ok {
		t.Error("GetChannel returned false for existing channel")
	}
	if got.ID != "ch-1" {
		t.Errorf("expected ID 'ch-1', got '%s'", got.ID)
	}
}

func TestGetChannel_NotFound(t *testing.T) {
	m := NewManager(nil)
	_, ok := m.GetChannel("non-existent")
	if ok {
		t.Error("GetChannel returned true for non-existent channel")
	}
}

func TestRemoveChannel(t *testing.T) {
	m := NewManager(nil)

	ch := &Channel{ID: "ch-1", Name: "Test"}
	m.AddChannel(ch)

	m.RemoveChannel("ch-1")

	channels := m.GetChannels()
	if len(channels) != 0 {
		t.Errorf("expected 0 channels after removal, got %d", len(channels))
	}
}

func TestGetChannelFolders(t *testing.T) {
	m := NewManager(nil)

	ch := &Channel{
		ID:   "ch-1",
		Name: "Test",
		Folders: []*ChannelFolder{
			{ID: "f1", Name: "Folder 1"},
			{ID: "f2", Name: "Folder 2"},
		},
	}
	m.AddChannel(ch)

	folders := m.GetChannelFolders("ch-1")
	if len(folders) != 2 {
		t.Errorf("expected 2 folders, got %d", len(folders))
	}
}

func TestGetChannelFolders_NotFound(t *testing.T) {
	m := NewManager(nil)
	folders := m.GetChannelFolders("non-existent")
	if folders != nil {
		t.Error("expected nil for non-existent channel")
	}
}

func TestGetChannelItems(t *testing.T) {
	m := NewManager(nil)

	ch := &Channel{
		ID:   "ch-1",
		Name: "Test",
		Items: []*ChannelItem{
			{ID: "item-1", Name: "Item 1"},
			{ID: "item-2", Name: "Item 2"},
		},
	}
	m.AddChannel(ch)

	items := m.GetChannelItems("ch-1", "user-1")
	if len(items) != 2 {
		t.Errorf("expected 2 items, got %d", len(items))
	}
}

func TestGetChannelItems_NotFound(t *testing.T) {
	m := NewManager(nil)
	items := m.GetChannelItems("non-existent", "user-1")
	if items != nil {
		t.Error("expected nil for non-existent channel")
	}
}