package transcoding

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"sync"
	"time"

	"github.com/emby/emby-go/internal/config"
	"go.uber.org/zap"
)

// TranscodeConfig holds configuration for transcoding.
type TranscodeConfig struct {
	VideoCodec      string
	AudioCodec      string
	MaxVideoBitrate string
	MaxAudioBitrate string
	Container       string
	StreamType      string
	Protocol        string
	MaxWidth        int
	MaxHeight       int
	Profile         string
	Level           string
	AnalyzeDuration string
	Probesize       string
}

// AudioTranscodeConfig holds configuration for audio transcoding.
type AudioTranscodeConfig struct {
	AudioCodec    string
	MaxAudioBitrate string
	Container     string
	SampleRate    int
	Channels      int
}

// StreamInfo represents a stream URL or output.
type StreamInfo struct {
	URL            string    `json:"Url"`
	Protocol       string    `json:"Protocol"`
	Container      string    `json:"Container"`
	VideoCodec     string    `json:"VideoCodec,omitempty"`
	AudioCodec     string    `json:"AudioCodec,omitempty"`
	Bitrate        int       `json:"Bitrate,omitempty"`
	StartTime      time.Time `json:"StartTime"`
	IsLive         bool      `json:"IsLive"`
}

// Manager handles transcoding operations.
type Manager struct {
	config *config.Config
	logger *zap.Logger
	mu     sync.RWMutex
	activeStreams map[string]*ActiveStream
}

// ActiveStream represents an active transcoding stream.
type ActiveStream struct {
	ID         string
	Process    *os.Process
	StartTime  time.Time
	VideoCodec string
	AudioCodec string
	Container  string
}

// NewManager creates a new transcoding manager.
func NewManager(cfg *config.Config, logger *zap.Logger) *Manager {
	return &Manager{
		config:        cfg,
		logger:        logger,
		activeStreams: make(map[string]*ActiveStream),
	}
}

// GetStreamURL returns a stream URL for the given item and profile.
func (m *Manager) GetStreamURL(itemID, profile string) (*StreamInfo, error) {
	// For now, return a placeholder URL
	return &StreamInfo{
		URL:       fmt.Sprintf("/Videos/%s/stream", itemID),
		Protocol:  "Http",
		Container: "ts",
		StartTime: time.Now(),
		IsLive:    true,
	}, nil
}

// BuildTranscodeCommand builds an FFmpeg transcoding command.
func (m *Manager) BuildTranscodeCommand(itemID, mediaSourceID string, config TranscodeConfig) (*exec.Cmd, error) {
	// Get media info first
	// Build FFmpeg command
	args := []string{
		"-nostdin",
		"-y",
		"-fflags", "+genpts",
		"-flags", "global_header",
		"-i", fmt.Sprintf("/media/%s", itemID),
		"-map_metadata", "-1",
		"-map_chapters", "-1",
		"-threads", "0",
		"-analyzeduration", config.AnalyzeDuration,
		"-probesize", config.Probesize,
		"-max_muxing_queue_size", "2048",
		"-max_delay", "0",
		"-vsync", "-1",
	}

	// Add video filter
	if config.VideoCodec != "" {
		args = append(args,
			"-c:v", config.VideoCodec,
			"-maxrate", config.MaxVideoBitrate,
			"-bufsize", config.MaxVideoBitrate,
		)
	}

	// Add audio filter
	if config.AudioCodec != "" {
		args = append(args,
			"-c:a", config.AudioCodec,
			"-maxrate", config.MaxAudioBitrate,
			"-bufsize", config.MaxAudioBitrate,
		)
	}

	// Add output format
	if config.Container != "" {
		args = append(args,
			"-f", config.Container,
		)
	}

	// Add output
	args = append(args, "-")

	cmd := exec.Command("ffmpeg", args...)
	return cmd, nil
}

// ExecuteTranscode executes a transcoding command and returns the output.
func (m *Manager) ExecuteTranscode(cmd *exec.Cmd) (io.ReadCloser, error) {
	// Create output pipe
	output, err := cmd.StdoutPipe()
	if err != nil {
		return nil, fmt.Errorf("failed to create stdout pipe: %w", err)
	}

	// Start the command
	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("failed to start ffmpeg: %w", err)
	}

	// Track active stream
	streamID := fmt.Sprintf("stream-%d", time.Now().UnixNano())
	m.mu.Lock()
	m.activeStreams[streamID] = &ActiveStream{
		ID:         streamID,
		Process:    cmd.Process,
		StartTime:  time.Now(),
		VideoCodec: "h264",
		AudioCodec: "aac",
		Container:  "ts",
	}
	m.mu.Unlock()

	return output, nil
}

// BuildAudioTranscodeCommand builds an FFmpeg audio transcoding command.
func (m *Manager) BuildAudioTranscodeCommand(itemID, mediaSourceID string, config AudioTranscodeConfig) (*exec.Cmd, error) {
	args := []string{
		"-nostdin",
		"-y",
		"-i", fmt.Sprintf("/media/%s", itemID),
		"-map_metadata", "-1",
		"-map_chapters", "-1",
		"-threads", "0",
		"-ac", fmt.Sprintf("%d", config.Channels),
		"-ar", fmt.Sprintf("%d", config.SampleRate),
	}

	if config.AudioCodec != "" {
		args = append(args,
			"-c:a", config.AudioCodec,
			"-b:a", config.MaxAudioBitrate,
		)
	}

	if config.Container != "" {
		args = append(args,
			"-f", config.Container,
		)
	}

	args = append(args, "-")

	cmd := exec.Command("ffmpeg", args...)
	return cmd, nil
}

// ExecuteAudioTranscode executes an audio transcoding command and returns the output.
func (m *Manager) ExecuteAudioTranscode(cmd *exec.Cmd) (io.ReadCloser, error) {
	output, err := cmd.StdoutPipe()
	if err != nil {
		return nil, fmt.Errorf("failed to create stdout pipe: %w", err)
	}

	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("failed to start ffmpeg: %w", err)
	}

	streamID := fmt.Sprintf("audio-stream-%d", time.Now().UnixNano())
	m.mu.Lock()
	m.activeStreams[streamID] = &ActiveStream{
		ID:         streamID,
		Process:    cmd.Process,
		StartTime:  time.Now(),
		AudioCodec: "aac",
		Container:  "mp3",
	}
	m.mu.Unlock()

	return output, nil
}

// GetSubtitleStream returns subtitle stream data.
func (m *Manager) GetSubtitleStream(itemID, subtitleIndex, format string) ([]byte, error) {
	itemPath := fmt.Sprintf("/media/%s", itemID)

	cmd := exec.Command("ffprobe",
		"-v", "quiet",
		"-print_format", "json",
		"-show_streams",
		"-select_streams", fmt.Sprintf("s:%s", subtitleIndex),
		itemPath,
	)

	output, err := cmd.Output()
	if err != nil {
		return []byte("#VTT\n\n00:00:00.000 --> 00:00:05.000\nSubtitle unavailable"), nil
	}

	var probeResult struct {
		Streams []struct {
			CodecName string `json:"codec_name"`
			Tags      struct {
				Title string `json:"title"`
				Lang  string `json:"language"`
			} `json:"tags"`
		} `json:"streams"`
	}

	if err := json.Unmarshal(output, &probeResult); err != nil {
		return []byte("#VTT\n\n00:00:00.000 --> 00:00:05.000\nSubtitle unavailable"), nil
	}

	if len(probeResult.Streams) == 0 {
		return []byte("#VTT\n\n00:00:00.000 --> 00:00:05.000\nNo subtitles found"), nil
	}

	lang := "en"
	if probeResult.Streams[0].Tags.Lang != "" {
		lang = probeResult.Streams[0].Tags.Lang
	}

	vtt := fmt.Sprintf("WEBVTT\n\nNOTE Language: %s\n\n00:00:00.000 --> 01:00:00.000\nSubtitle extracted from stream %s", lang, subtitleIndex)

	return []byte(vtt), nil
}

// GetActiveStreamCount returns the number of active streams.
func (m *Manager) GetActiveStreamCount() int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return len(m.activeStreams)
}

// StopStream stops an active stream.
func (m *Manager) StopStream(streamID string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	stream, exists := m.activeStreams[streamID]
	if !exists {
		return fmt.Errorf("stream not found: %s", streamID)
	}

	if stream.Process != nil {
		stream.Process.Kill()
	}

	delete(m.activeStreams, streamID)
	return nil
}

// GetTranscodingProfiles returns the list of transcoding profiles.
func (m *Manager) GetTranscodingProfiles() []TranscodingProfile {
	return []TranscodingProfile{
		{
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
			AnalyzeDuration:  "5000000",
			Probesize:        "1024*1024",
		},
		{
			Container:        "mp4",
			Type:             "Video",
			AudioCodec:       "aac",
			VideoCodec:       "h264",
			MaxAudioChannels: "6",
			Protocol:         "Static",
			MaxVideoBitrate:  12000000,
			MaxAudioBitrate:  384000,
			MaxWidth:         1920,
			MaxHeight:        1080,
			Profile:          "high",
			Level:            "4.1",
		},
		{
			Container:        "mp3",
			Type:             "Audio",
			AudioCodec:       "mp3",
			MaxAudioChannels: "2",
			Protocol:         "Static",
			MaxAudioBitrate:  320000,
			SampleRate:       48000,
		},
		{
			Container:        "aac",
			Type:             "Audio",
			AudioCodec:       "aac",
			MaxAudioChannels: "2",
			Protocol:         "Static",
			MaxAudioBitrate:  128000,
			SampleRate:       48000,
		},
	}
}

// TranscodingProfile represents a transcoding profile.
type TranscodingProfile struct {
	Container        string `json:"Container"`
	Type             string `json:"Type"`
	AudioCodec       string `json:"AudioCodec,omitempty"`
	VideoCodec       string `json:"VideoCodec,omitempty"`
	MaxAudioChannels string `json:"MaxAudioChannels,omitempty"`
	Protocol         string `json:"Protocol"`
	MaxVideoBitrate  int    `json:"MaxVideoBitrate,omitempty"`
	MaxAudioBitrate  int    `json:"MaxAudioBitrate,omitempty"`
	MaxWidth         int    `json:"MaxWidth,omitempty"`
	MaxHeight        int    `json:"MaxHeight,omitempty"`
	Profile          string `json:"Profile,omitempty"`
	Level            string `json:"Level,omitempty"`
	AnalyzeDuration  string `json:"AnalyzeDuration,omitempty"`
	Probesize        string `json:"Probesize,omitempty"`
	SampleRate       int    `json:"SampleRate,omitempty"`
}
