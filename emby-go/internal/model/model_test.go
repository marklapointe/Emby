package model

import (
	"encoding/json"
	"testing"
	"time"
)

func TestMediaItemJSONMarshaling(t *testing.T) {
	item := MediaItem{
		ID:             "test-id-123",
		Name:           "Test Movie",
		Overview:       "A test movie overview",
		MediaType:      "Movie",
		ProductionYear: 2024,
		Path:           "/media/movies/test.mkv",
		Genres:         []string{"Action", "Sci-Fi"},
	}

	data, err := json.Marshal(item)
	if err != nil {
		t.Fatalf("failed to marshal: %v", err)
	}

	var unmarshaled MediaItem
	if err := json.Unmarshal(data, &unmarshaled); err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}

	if unmarshaled.ID != item.ID {
		t.Errorf("expected ID %s, got %s", item.ID, unmarshaled.ID)
	}
	if unmarshaled.Name != item.Name {
		t.Errorf("expected Name %s, got %s", item.Name, unmarshaled.Name)
	}
	if unmarshaled.MediaType != item.MediaType {
		t.Errorf("expected MediaType %s, got %s", item.MediaType, unmarshaled.MediaType)
	}
	if len(unmarshaled.Genres) != len(item.Genres) {
		t.Errorf("expected %d genres, got %d", len(item.Genres), len(unmarshaled.Genres))
	}
}

func TestMediaItemWithSources(t *testing.T) {
	item := MediaItem{
		ID:   "test-id",
		Name: "Test",
		MediaSources: []MediaSource{
			{
				ID:          "source-1",
				Name:        "Source 1",
				Type:        "Default",
				Container:   "mkv",
				Path:        "/media/test.mkv",
				Protocol:    "File",
				VideoCodec:  "h264",
				AudioCodec:  "aac",
				SupportsDirectPlay: true,
			},
		},
	}

	data, err := json.Marshal(item)
	if err != nil {
		t.Fatalf("failed to marshal: %v", err)
	}

	var unmarshaled MediaItem
	if err := json.Unmarshal(data, &unmarshaled); err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}

	if len(unmarshaled.MediaSources) != 1 {
		t.Errorf("expected 1 media source, got %d", len(unmarshaled.MediaSources))
	}
	if unmarshaled.MediaSources[0].ID != "source-1" {
		t.Errorf("expected source ID source-1, got %s", unmarshaled.MediaSources[0].ID)
	}
}

func TestSessionInfoJSONMarshaling(t *testing.T) {
	session := SessionInfo{
		ID:            "session-123",
		Client:        "Emby for Android",
		DeviceName:    "Test Device",
		DisplayName:   "Test User",
		MachineID:     "machine-456",
		LastActivityTime: time.Now(),
		Location: Location{
			DeviceName: "Test Device",
			IPAddress:  "192.168.1.100",
		},
		PlayState: PlayState{
			PositionTicks: 1234567890,
			VolumePercent: 80,
			IsMuted:       false,
			IsPaused:      false,
		},
		Capabilities: Capabilities{
			SupportedMediaTypes: []string{"Video", "Audio"},
			EnableMediaControl:  true,
			EnableRemoteControl: true,
		},
	}

	data, err := json.Marshal(session)
	if err != nil {
		t.Fatalf("failed to marshal: %v", err)
	}

	var unmarshaled SessionInfo
	if err := json.Unmarshal(data, &unmarshaled); err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}

	if unmarshaled.ID != session.ID {
		t.Errorf("expected ID %s, got %s", session.ID, unmarshaled.ID)
	}
	if unmarshaled.Client != session.Client {
		t.Errorf("expected Client %s, got %s", session.Client, unmarshaled.Client)
	}
	if unmarshaled.PlayState.PositionTicks != session.PlayState.PositionTicks {
		t.Errorf("expected position %d, got %d", session.PlayState.PositionTicks, unmarshaled.PlayState.PositionTicks)
	}
}

func TestPlayMethodConstants(t *testing.T) {
	if PlayMethodTranscode != "Transcode" {
		t.Errorf("expected PlayMethodTranscode to be 'Transcode', got %s", PlayMethodTranscode)
	}
	if PlayMethodDirectStream != "DirectStream" {
		t.Errorf("expected PlayMethodDirectStream to be 'DirectStream', got %s", PlayMethodDirectStream)
	}
	if PlayMethodDirectPlay != "DirectPlay" {
		t.Errorf("expected PlayMethodDirectPlay to be 'DirectPlay', got %s", PlayMethodDirectPlay)
	}
}

func TestSubtitleMethodConstants(t *testing.T) {
	if SubtitleMethodEmbed != "Embed" {
		t.Errorf("expected SubtitleMethodEmbed to be 'Embed', got %s", SubtitleMethodEmbed)
	}
	if SubtitleMethodExternal != "External" {
		t.Errorf("expected SubtitleMethodExternal to be 'External', got %s", SubtitleMethodExternal)
	}
	if SubtitleMethodHls != "Hls" {
		t.Errorf("expected SubtitleMethodHls to be 'Hls', got %s", SubtitleMethodHls)
	}
}

func TestUserDataJSONMarshaling(t *testing.T) {
	userData := UserData{
		PlaybackPositionTicks: 9876543210,
		PlayCount:             5,
		IsFavorite:            true,
		Liked:                 true,
		Played:                true,
		Rating:                4,
	}

	data, err := json.Marshal(userData)
	if err != nil {
		t.Fatalf("failed to marshal: %v", err)
	}

	var unmarshaled UserData
	if err := json.Unmarshal(data, &unmarshaled); err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}

	if unmarshaled.PlayCount != userData.PlayCount {
		t.Errorf("expected PlayCount %d, got %d", userData.PlayCount, unmarshaled.PlayCount)
	}
	if !unmarshaled.IsFavorite {
		t.Error("expected IsFavorite to be true")
	}
	if unmarshaled.Rating != userData.Rating {
		t.Errorf("expected Rating %d, got %d", userData.Rating, unmarshaled.Rating)
	}
}

func TestTranscodingProfileJSONMarshaling(t *testing.T) {
	profile := TranscodingProfile{
		Container:        "ts",
		Type:             "Video",
		AudioCodec:       "aac",
		VideoCodec:       "h264",
		MaxAudioChannels: "6",
		Protocol:         "Hls",
		MaxVideoBitrate:  12000000,
		MaxAudioBitrate:  384000,
		MaxWidth:         1920,
		MaxHeight:        1080,
		Profile:          "high",
		Level:            "4.1",
	}

	data, err := json.Marshal(profile)
	if err != nil {
		t.Fatalf("failed to marshal: %v", err)
	}

	var unmarshaled TranscodingProfile
	if err := json.Unmarshal(data, &unmarshaled); err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}

	if unmarshaled.Container != profile.Container {
		t.Errorf("expected Container %s, got %s", profile.Container, unmarshaled.Container)
	}
	if unmarshaled.MaxVideoBitrate != profile.MaxVideoBitrate {
		t.Errorf("expected MaxVideoBitrate %d, got %d", profile.MaxVideoBitrate, unmarshaled.MaxVideoBitrate)
	}
	if unmarshaled.MaxWidth != profile.MaxWidth {
		t.Errorf("expected MaxWidth %d, got %d", profile.MaxWidth, unmarshaled.MaxWidth)
	}
}

func TestDirectPlayProfileJSONMarshaling(t *testing.T) {
	profile := DirectPlayProfile{
		Container:  "mkv",
		Type:       "Video",
		VideoCodec: "h264",
		AudioCodec: "aac",
		MaxChannels: 6,
	}

	data, err := json.Marshal(profile)
	if err != nil {
		t.Fatalf("failed to marshal: %v", err)
	}

	var unmarshaled DirectPlayProfile
	if err := json.Unmarshal(data, &unmarshaled); err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}

	if unmarshaled.Container != profile.Container {
		t.Errorf("expected Container %s, got %s", profile.Container, unmarshaled.Container)
	}
	if unmarshaled.MaxChannels != profile.MaxChannels {
		t.Errorf("expected MaxChannels %d, got %d", profile.MaxChannels, unmarshaled.MaxChannels)
	}
}
