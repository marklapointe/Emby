package transcoding

import (
	"testing"
)

func TestNewManager(t *testing.T) {
	m := NewManager(nil, nil)
	if m == nil {
		t.Fatal("NewManager returned nil")
	}
	if m.activeStreams == nil {
		t.Error("activeStreams map not initialized")
	}
}

func TestGetTranscodingProfiles(t *testing.T) {
	m := NewManager(nil, nil)
	profiles := m.GetTranscodingProfiles()
	if len(profiles) == 0 {
		t.Error("expected at least one profile")
	}
	if profiles[0].Container != "ts" {
		t.Errorf("expected first profile container 'ts', got '%s'", profiles[0].Container)
	}
}

func TestGetActiveStreamCount(t *testing.T) {
	m := NewManager(nil, nil)
	count := m.GetActiveStreamCount()
	if count != 0 {
		t.Errorf("expected 0 active streams, got %d", count)
	}
}

func TestGetStreamURL(t *testing.T) {
	m := NewManager(nil, nil)
	info, err := m.GetStreamURL("item-1", "mobile")
	if err != nil {
		t.Fatalf("GetStreamURL returned error: %v", err)
	}
	if info.URL == "" {
		t.Error("expected non-empty URL")
	}
	if !info.IsLive {
		t.Error("expected IsLive to be true")
	}
}

func TestBuildTranscodeCommand(t *testing.T) {
	m := NewManager(nil, nil)
	config := TranscodeConfig{
		VideoCodec:      "h264",
		AudioCodec:      "aac",
		MaxVideoBitrate: "5000k",
		MaxAudioBitrate: "192k",
		Container:       "ts",
	}
	cmd, err := m.BuildTranscodeCommand("item-1", "source-1", config)
	if err != nil {
		t.Fatalf("BuildTranscodeCommand returned error: %v", err)
	}
	if cmd.Path != "ffmpeg" {
		t.Errorf("expected command 'ffmpeg', got '%s'", cmd.Path)
	}
}

func TestBuildAudioTranscodeCommand(t *testing.T) {
	m := NewManager(nil, nil)
	config := AudioTranscodeConfig{
		AudioCodec:    "mp3",
		MaxAudioBitrate: "320k",
		Container:     "mp3",
		SampleRate:    48000,
		Channels:      2,
	}
	cmd, err := m.BuildAudioTranscodeCommand("item-1", "source-1", config)
	if err != nil {
		t.Fatalf("BuildAudioTranscodeCommand returned error: %v", err)
	}
	if cmd.Path != "ffmpeg" {
		t.Errorf("expected command 'ffmpeg', got '%s'", cmd.Path)
	}
}

func TestGetSubtitleStream(t *testing.T) {
	m := NewManager(nil, nil)
	data, err := m.GetSubtitleStream("item-1", "0", "vtt")
	if err != nil {
		t.Fatalf("GetSubtitleStream returned error: %v", err)
	}
	if len(data) == 0 {
		t.Error("expected non-empty subtitle data")
	}
}

func TestStopStream_NotFound(t *testing.T) {
	m := NewManager(nil, nil)
	err := m.StopStream("non-existent")
	if err == nil {
		t.Error("expected error for non-existent stream")
	}
}

func TestTranscodeConfig(t *testing.T) {
	config := TranscodeConfig{
		VideoCodec:      "h264",
		AudioCodec:      "aac",
		MaxVideoBitrate: "5000k",
		Container:       "ts",
	}
	if config.VideoCodec != "h264" {
		t.Errorf("expected VideoCodec 'h264', got '%s'", config.VideoCodec)
	}
}