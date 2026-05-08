package media

import (
	"context"
	"testing"

	"go.uber.org/zap"
)

func TestNewManager(t *testing.T) {
	m := NewManager(nil, nil)
	if m == nil {
		t.Fatal("NewManager returned nil")
	}
	if m.mediaDB == nil {
		t.Error("mediaDB map not initialized")
	}
}

func TestGetMediaSource_NotFound(t *testing.T) {
	m := NewManager(nil, nil)
	_, err := m.GetMediaSource("non-existent")
	if err == nil {
		t.Error("expected error for non-existent media source")
	}
}

func TestGetMediaSources_NotFound(t *testing.T) {
	m := NewManager(nil, nil)
	_, err := m.GetMediaSources("non-existent")
	if err == nil {
		t.Error("expected error for non-existent media sources")
	}
}

func TestGetMediaInfo_NotFound(t *testing.T) {
	m := NewManager(nil, nil)
	_, err := m.GetMediaInfo("/non/existent/path.mp4")
	if err == nil {
		t.Error("expected error for non-existent file")
	}
}

func TestAddMediaSource(t *testing.T) {
	m := NewManager(nil, nil)
	info := &MediaInfo{
		ID:   "media-1",
		Name: "Test Media",
		MediaSources: []MediaSource{
			{ID: "source-1", Name: "Source 1"},
		},
	}
	m.mediaDB["media-1"] = info

	ms, err := m.GetMediaSource("media-1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if ms.ID != "source-1" {
		t.Errorf("expected source ID 'source-1', got '%s'", ms.ID)
	}
}

func TestGetMediaSources(t *testing.T) {
	m := NewManager(nil, nil)
	info := &MediaInfo{
		ID:   "media-1",
		Name: "Test Media",
		MediaSources: []MediaSource{
			{ID: "source-1", Name: "Source 1"},
			{ID: "source-2", Name: "Source 2"},
		},
	}
	m.mediaDB["media-1"] = info

	sources, err := m.GetMediaSources("media-1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(sources) != 2 {
		t.Errorf("expected 2 sources, got %d", len(sources))
	}
}

func TestRemoveMediaSource(t *testing.T) {
	m := NewManager(nil, nil)
	info := &MediaInfo{
		ID:   "media-1",
		Name: "Test Media",
		MediaSources: []MediaSource{
			{ID: "source-1", Name: "Source 1"},
		},
	}
	m.mediaDB["media-1"] = info
	delete(m.mediaDB, "media-1")

	_, err := m.GetMediaSource("media-1")
	if err == nil {
		t.Error("expected error after removing media source")
	}
}

func TestAudioStream_Fields(t *testing.T) {
	stream := AudioStream{
		Index:          0,
		Codecs:         "aac",
		Language:       "en",
		DisplayLanguage: "English",
		ChannelLayout:  "stereo",
		SampleRate:     48000,
		BitRate:        128000,
		IsDefault:      true,
		IsForced:       false,
		IsExternal:     false,
	}
	if stream.Index != 0 {
		t.Errorf("expected Index 0, got %d", stream.Index)
	}
	if stream.Codecs != "aac" {
		t.Errorf("expected Codecs 'aac', got '%s'", stream.Codecs)
	}
}

func TestSubtitleStream_Fields(t *testing.T) {
	stream := SubtitleStream{
		Index:          1,
		Codecs:         "srt",
		Language:       "en",
		DisplayLanguage: "English",
		IsDefault:      false,
		IsForced:       true,
		IsExternal:     true,
		Path:           "/subs/en.srt",
	}
	if stream.Index != 1 {
		t.Errorf("expected Index 1, got %d", stream.Index)
	}
	if stream.Path != "/subs/en.srt" {
		t.Errorf("expected Path '/subs/en.srt', got '%s'", stream.Path)
	}
}

func TestVideoStream_Fields(t *testing.T) {
	stream := VideoStream{
		Index:     0,
		Codecs:    "h264",
		Width:     1920,
		Height:    1080,
		FrameRate: 23.976,
		BitRate:   5000000,
		IsDefault: true,
		Profile:   "high",
		Level:     41,
	}
	if stream.Width != 1920 {
		t.Errorf("expected Width 1920, got %d", stream.Width)
	}
	if stream.Height != 1080 {
		t.Errorf("expected Height 1080, got %d", stream.Height)
	}
}

func TestChapter_Fields(t *testing.T) {
	chapter := Chapter{
		Index:              0,
		Name:               "Chapter 1",
		StartPositionTicks: 0,
	}
	if chapter.Index != 0 {
		t.Errorf("expected Index 0, got %d", chapter.Index)
	}
	if chapter.StartPositionTicks != 0 {
		t.Errorf("expected StartPositionTicks 0, got %d", chapter.StartPositionTicks)
	}
}

func TestMediaSource_Fields(t *testing.T) {
	ms := MediaSource{
		ID:                  "ms-1",
		Name:                "Test Source",
		Type:                "Default",
		Container:           "mkv",
		Path:                "/media/test.mkv",
		Protocol:            "File",
		VideoCodec:          "h264",
		AudioCodec:          "aac",
		Width:               1920,
		Height:              1080,
		Size:                1024000000,
		SupportsDirectPlay:   true,
		SupportsDirectStream: true,
		SupportsTranscoding:  false,
	}
	if ms.ID != "ms-1" {
		t.Errorf("expected ID 'ms-1', got '%s'", ms.ID)
	}
	if ms.Size != 1024000000 {
		t.Errorf("expected Size 1024000000, got %d", ms.Size)
	}
}

func TestMediaInfo_Fields(t *testing.T) {
	info := MediaInfo{
		ID:            "mi-1",
		Name:          "Test Media",
		Path:          "/media/test.mkv",
		MediaType:     "Video",
		Container:     "mkv",
		VideoCodec:    "h264",
		AudioCodec:    "aac",
		Width:         1920,
		Height:        1080,
		VideoDuration: 7200000000,
		Bitrate:       5000000,
		AudioChannels: 2,
	}
	if info.ID != "mi-1" {
		t.Errorf("expected ID 'mi-1', got '%s'", info.ID)
	}
	if info.MediaType != "Video" {
		t.Errorf("expected MediaType 'Video', got '%s'", info.MediaType)
	}
}

func TestStreamManager_NewManager(t *testing.T) {
	m := NewStreamManager(5, nil)
	if m == nil {
		t.Fatal("NewStreamManager returned nil")
	}
	if m.maxSourceStreams != 5 {
		t.Errorf("expected maxSourceStreams 5, got %d", m.maxSourceStreams)
	}
	if m.sources == nil {
		t.Error("sources map not initialized")
	}
}

func TestStreamManager_GetMetrics(t *testing.T) {
	m := NewStreamManager(5, nil)
	metrics := m.GetMetrics()
	if metrics == nil {
		t.Fatal("GetMetrics returned nil")
	}
	if metrics.TotalSourceStreamsCreated != 0 {
		t.Errorf("expected 0 TotalSourceStreamsCreated, got %d", metrics.TotalSourceStreamsCreated)
	}
}

func TestStreamManager_GetStreamViewers_NotFound(t *testing.T) {
	m := NewStreamManager(5, nil)
	profile := &TranscodingProfile{Container: "ts", VideoCodec: "h264", MaxVideoBitrate: 5000}
	viewers := m.GetStreamViewers("non-existent", profile)
	if viewers != 0 {
		t.Errorf("expected 0 viewers, got %d", viewers)
	}
}

func TestStreamError_Error(t *testing.T) {
	err := &StreamError{Message: "test message"}
	if err.Error() != "test message" {
		t.Errorf("expected 'test message', got '%s'", err.Error())
	}
}

func TestStreamManager_RemoveViewer_NotFound(t *testing.T) {
	m := NewStreamManager(1, nil)
	profile := &TranscodingProfile{Container: "ts", VideoCodec: "h264", MaxVideoBitrate: 5000}
	m.RemoveViewer("non-existent", profile, "viewer-1")
}

func TestParseInt(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"123", 123},
		{"0", 0},
		{"-456", -456},
		{"invalid", 0},
	}

	for _, tc := range tests {
		result := parseInt(tc.input)
		if result != tc.expected {
			t.Errorf("parseInt(%s) = %d, expected %d", tc.input, result, tc.expected)
		}
	}
}

func TestParseBitrate(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"5000", 5000},
		{"0", 0},
		{"-1000", -1000},
		{"invalid", 0},
	}

	for _, tc := range tests {
		result := parseBitrate(tc.input)
		if result != tc.expected {
			t.Errorf("parseBitrate(%s) = %d, expected %d", tc.input, result, tc.expected)
		}
	}
}

func TestParseSize(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"1024000", 1024000},
		{"0", 0},
		{"-500", -500},
		{"invalid", 0},
	}

	for _, tc := range tests {
		result := parseSize(tc.input)
		if result != tc.expected {
			t.Errorf("parseSize(%s) = %d, expected %d", tc.input, result, tc.expected)
		}
	}
}

func TestItemIDFromPath(t *testing.T) {
	id1 := itemIDFromPath("/path/to/file.mp4")
	id2 := itemIDFromPath("/path/to/file.mp4")

	if id1 == "" {
		t.Error("itemIDFromPath returned empty string")
	}
	if id1 == id2 {
		t.Error("itemIDFromPath should generate unique IDs")
	}
}

func TestStreamManager_GetOrCreateStream_New(t *testing.T) {
	m := NewStreamManager(5, zap.NewNop())

	profile := &TranscodingProfile{
		Container: "mp4",
		VideoCodec: "h264",
	}

	src, out, err := m.GetOrCreateStream(context.Background(), "content-1", profile, "viewer-1")
	if err != nil {
		t.Fatalf("GetOrCreateStream failed: %v", err)
	}
	if src == nil {
		t.Fatal("GetOrCreateStream returned nil source")
	}
	if out == nil {
		t.Fatal("GetOrCreateStream returned nil output")
	}
	if src.ContentID != "content-1" {
		t.Errorf("expected ContentID 'content-1', got '%s'", src.ContentID)
	}
}

func TestStreamManager_GetOrCreateStream_Existing(t *testing.T) {
	m := NewStreamManager(5, zap.NewNop())

	profile := &TranscodingProfile{
		Container: "mp4",
		VideoCodec: "h264",
	}

	src1, out1, err := m.GetOrCreateStream(context.Background(), "content-1", profile, "viewer-1")
	if err != nil {
		t.Fatalf("GetOrCreateStream failed: %v", err)
	}

	src2, out2, err := m.GetOrCreateStream(context.Background(), "content-1", profile, "viewer-2")
	if err != nil {
		t.Fatalf("GetOrCreateStream failed: %v", err)
	}

	if src1 != src2 {
		t.Error("GetOrCreateStream should return same source for same content")
	}

	if out1 != out2 {
		t.Error("GetOrCreateStream should return same output for same content+profile")
	}

	if len(out2.Viewers) != 2 {
		t.Errorf("expected 2 viewers, got %d", len(out2.Viewers))
	}
}

func TestStreamManager_GetOrCreateStream_MaxStreams(t *testing.T) {
	m := NewStreamManager(1, zap.NewNop())

	profile := &TranscodingProfile{
		Container: "mp4",
		VideoCodec: "h264",
	}

	_, _, err := m.GetOrCreateStream(context.Background(), "content-1", profile, "viewer-1")
	if err != nil {
		t.Fatalf("GetOrCreateStream failed: %v", err)
	}

	_, _, err = m.GetOrCreateStream(context.Background(), "content-2", profile, "viewer-2")
	if err == nil {
		t.Error("expected error when max streams exceeded")
	}
}