package media

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"time"

	"github.com/emby/emby-go/internal/config"
	"github.com/emby/emby-go/internal/model"
	"go.uber.org/zap"
)

// MediaSource represents a media file source.
type MediaSource struct {
	ID               string                 `json:"Id"`
	Name             string                 `json:"Name"`
	Type             string                 `json:"Type"`
	Container        string                 `json:"Container"`
	Path             string                 `json:"Path"`
	Protocol         string                 `json:"Protocol"`
	VideoCodec       string                 `json:"VideoCodec,omitempty"`
	AudioCodec       string                 `json:"AudioCodec,omitempty"`
	Width            int                    `json:"Width,omitempty"`
	Height           int                    `json:"Height,omitempty"`
	VideoDuration    int64                  `json:"VideoDuration,omitempty"`
	AudioDuration    int64                  `json:"AudioDuration,omitempty"`
	VideoBitrate     int                    `json:"VideoBitrate,omitempty"`
	AudioBitrate     int                    `json:"AudioBitrate,omitempty"`
	AudioChannels    int                    `json:"AudioChannels,omitempty"`
	FrameRate        float64                `json:"FrameRate,omitempty"`
	IsDefault        bool                   `json:"IsDefault"`
	IsSecondary      bool                   `json:"IsSecondary"`
	SupportsDirectPlay bool                `json:"SupportsDirectPlay"`
	SupportsDirectStream bool              `json:"SupportsDirectStream"`
	SupportsTranscoding bool               `json:"SupportsTranscoding"`
	Bitrate          int                    `json:"Bitrate,omitempty"`
	RequiredHttpHeaders map[string]string `json:"RequiredHttpHeaders,omitempty"`
	Size             int64                  `json:"Size,omitempty"`
	MediaStreams     []model.MediaStream    `json:"MediaStreams,omitempty"`
	Formats          []string               `json:"Formats,omitempty"`
	DefaultAudioStreamIndex *int             `json:"DefaultAudioStreamIndex,omitempty"`
	DefaultSubtitleStreamIndex *int            `json:"DefaultSubtitleStreamIndex,omitempty"`
}

// MediaInfo represents media file information.
type MediaInfo struct {
	ID            string        `json:"Id"`
	Name          string        `json:"Name"`
	Path          string        `json:"Path"`
	MediaType     string        `json:"MediaType"`
	Container     string        `json:"Container"`
	VideoCodec    string        `json:"VideoCodec,omitempty"`
	AudioCodec    string        `json:"AudioCodec,omitempty"`
	Width         int           `json:"Width,omitempty"`
	Height        int           `json:"Height,omitempty"`
	FrameRate     float64       `json:"FrameRate,omitempty"`
	VideoDuration int64         `json:"VideoDuration,omitempty"`
	AudioDuration int64         `json:"AudioDuration,omitempty"`
	Bitrate       int           `json:"Bitrate,omitempty"`
	AudioChannels int           `json:"AudioChannels,omitempty"`
	AudioStreams  []AudioStream `json:"AudioStreams,omitempty"`
	SubtitleStreams []SubtitleStream `json:"SubtitleStreams,omitempty"`
	VideoStreams  []VideoStream `json:"VideoStreams,omitempty"`
	Chapters      []Chapter     `json:"Chapters,omitempty"`
	MediaSources  []MediaSource `json:"MediaSources,omitempty"`
}

// AudioStream represents an audio stream.
type AudioStream struct {
	Index          int    `json:"Index"`
	Codecs         string `json:"Codecs"`
	Language       string `json:"Language,omitempty"`
	DisplayLanguage string `json:"DisplayLanguage,omitempty"`
	ChannelLayout  string `json:"ChannelLayout,omitempty"`
	SampleRate     int    `json:"SampleRate,omitempty"`
	BitRate        int    `json:"BitRate,omitempty"`
	IsDefault      bool   `json:"IsDefault"`
	IsForced       bool   `json:"IsForced"`
	IsExternal     bool   `json:"IsExternal"`
}

// SubtitleStream represents a subtitle stream.
type SubtitleStream struct {
	Index          int    `json:"Index"`
	Codecs         string `json:"Codecs"`
	Language       string `json:"Language,omitempty"`
	DisplayLanguage string `json:"DisplayLanguage,omitempty"`
	IsDefault      bool   `json:"IsDefault"`
	IsForced       bool   `json:"IsForced"`
	IsExternal     bool   `json:"IsExternal"`
	Path           string `json:"Path,omitempty"`
}

// VideoStream represents a video stream.
type VideoStream struct {
	Index          int     `json:"Index"`
	Codecs         string  `json:"Codecs"`
	Width          int     `json:"Width,omitempty"`
	Height         int     `json:"Height,omitempty"`
	FrameRate      float64 `json:"FrameRate,omitempty"`
	BitRate        int     `json:"BitRate,omitempty"`
	IsDefault      bool    `json:"IsDefault"`
	IsForced       bool    `json:"IsForced"`
	IsExternal     bool    `json:"IsExternal"`
	Profile        string  `json:"Profile,omitempty"`
	Level          int     `json:"Level,omitempty"`
}

// Chapter represents a media chapter.
type Chapter struct {
	Index    int     `json:"Index"`
	Name     string  `json:"Name,omitempty"`
	StartPositionTicks int64 `json:"StartPositionTicks"`
}

// Manager handles media-related operations.
type Manager struct {
	config *config.Config
	logger *zap.Logger
	mu     sync.RWMutex
	mediaDB map[string]*MediaInfo
}

// NewManager creates a new media manager.
func NewManager(cfg *config.Config, logger *zap.Logger) *Manager {
	return &Manager{
		config:  cfg,
		logger:  logger,
		mediaDB: make(map[string]*MediaInfo),
	}
}

// GetMediaSource returns media source information for an item.
func (m *Manager) GetMediaSource(itemID string) (*MediaSource, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	mediaInfo, exists := m.mediaDB[itemID]
	if !exists {
		return nil, fmt.Errorf("media info not found for item: %s", itemID)
	}

	if len(mediaInfo.MediaSources) == 0 {
		return nil, fmt.Errorf("no media sources found for item: %s", itemID)
	}

	return &mediaInfo.MediaSources[0], nil
}

// GetMediaSources returns all media sources for an item.
func (m *Manager) GetMediaSources(itemID string) ([]MediaSource, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	mediaInfo, exists := m.mediaDB[itemID]
	if !exists {
		return nil, fmt.Errorf("media info not found for item: %s", itemID)
	}

	return mediaInfo.MediaSources, nil
}

// GetMediaInfo returns media information for a file path.
func (m *Manager) GetMediaInfo(filePath string) (*MediaInfo, error) {
	// Check cache first
	m.mu.RLock()
	for _, info := range m.mediaDB {
		if info.Path == filePath {
			m.mu.RUnlock()
			return info, nil
		}
	}
	m.mu.RUnlock()

	// Extract media info using ffprobe
	info, err := m.extractMediaInfo(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to extract media info: %w", err)
	}

	// Cache the result
	m.mu.Lock()
	m.mediaDB[filePath] = info
	m.mu.Unlock()

	return info, nil
}

// extractMediaInfo extracts media information using ffprobe.
func (m *Manager) extractMediaInfo(filePath string) (*MediaInfo, error) {
	// Check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("file not found: %s", filePath)
	}

	// Run ffprobe
	cmd := exec.Command("ffprobe",
		"-v", "quiet",
		"-print_format", "json",
		"-show_format",
		"-show_streams",
		filePath,
	)

	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("ffprobe failed: %w", err)
	}

	// Parse ffprobe output
	var probeResult struct {
		Format struct {
			FileName    string  `json:"filename"`
			NbStreams   int     `json:"nb_streams"`
			FormatName  string  `json:"format_name"`
			FormatLong  string  `json:"format_long_name"`
			StartTime   float64 `json:"start_time"`
			Duration    float64 `json:"duration"`
			Size        string  `json:"size"`
			BitRate     string  `json:"bit_rate"`
		} `json:"format"`
		Streams []struct {
			Index          int     `json:"index"`
			Codecs         string  `json:"codec_name"`
			CodecType      string  `json:"codec_type"`
			Width          int     `json:"width,omitempty"`
			Height         int     `json:"height,omitempty"`
			FrameRate      string  `json:"r_frame_rate"`
			BitRate        string  `json:"bit_rate"`
			SampleRate     string  `json:"sample_rate"`
			Channels       int     `json:"channels"`
			ChannelLayout  string  `json:"channel_layout"`
			Language       string  `json:"language"`
			Disposition    struct {
				Default int `json:"default"`
				Forced  int `json:"forced"`
			} `json:"disposition"`
		} `json:"streams"`
	}

	if err := json.Unmarshal(output, &probeResult); err != nil {
		return nil, fmt.Errorf("failed to parse ffprobe output: %w", err)
	}

	// Build MediaInfo
	info := &MediaInfo{
		Path:  filePath,
		Name:  filepath.Base(filePath),
		Container: probeResult.Format.FormatName,
	}

	// Parse streams
	for _, stream := range probeResult.Streams {
		switch stream.CodecType {
		case "video":
			info.VideoCodec = stream.Codecs
			info.Width = stream.Width
			info.Height = stream.Height
			info.VideoStreams = append(info.VideoStreams, VideoStream{
				Index:     stream.Index,
				Codecs:    stream.Codecs,
				Width:     stream.Width,
				Height:    stream.Height,
				IsDefault: stream.Disposition.Default == 1,
				IsForced:  stream.Disposition.Forced == 1,
			})
		case "audio":
			info.AudioCodec = stream.Codecs
			info.AudioChannels = stream.Channels
			info.AudioStreams = append(info.AudioStreams, AudioStream{
				Index:       stream.Index,
				Codecs:      stream.Codecs,
				Language:    stream.Language,
				SampleRate:  parseInt(stream.SampleRate),
				ChannelLayout: stream.ChannelLayout,
				IsDefault:   stream.Disposition.Default == 1,
				IsForced:    stream.Disposition.Forced == 1,
			})
		case "subtitle":
			info.SubtitleStreams = append(info.SubtitleStreams, SubtitleStream{
				Index:     stream.Index,
				Codecs:    stream.Codecs,
				Language:  stream.Language,
				IsDefault: stream.Disposition.Default == 1,
				IsForced:  stream.Disposition.Forced == 1,
			})
		}
	}

	// Parse duration
	info.VideoDuration = int64(probeResult.Format.Duration * 10000000) // Convert to ticks
	info.AudioDuration = info.VideoDuration

	// Parse bitrate
	info.Bitrate = parseBitrate(probeResult.Format.BitRate)

	// Create default media source
	info.MediaSources = []MediaSource{
		{
			ID:               itemIDFromPath(filePath),
			Name:             info.Name,
			Type:             "Default",
			Container:        info.Container,
			Path:             info.Path,
			Protocol:         "File",
			SupportsDirectPlay: true,
			SupportsDirectStream: true,
			SupportsTranscoding: true,
			Bitrate:          info.Bitrate,
			Size:             parseSize(probeResult.Format.Size),
		},
	}

	return info, nil
}

// itemIDFromPath generates an item ID from a file path.
func itemIDFromPath(path string) string {
	return fmt.Sprintf("item-%d", time.Now().UnixNano())
}

// Helper functions

func parseInt(s string) int {
	var n int
	fmt.Sscanf(s, "%d", &n)
	return n
}

func parseBitrate(s string) int {
	var n int
	fmt.Sscanf(s, "%d", &n)
	return n
}

func parseSize(s string) int64 {
	var n int64
	fmt.Sscanf(s, "%d", &n)
	return n
}
