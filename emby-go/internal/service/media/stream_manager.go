package media

import (
	"context"
	"sync"
	"time"

	"go.uber.org/zap"
)

// StreamManager manages shared content streams.
type StreamManager struct {
	mu           sync.RWMutex
	activeStreams map[string]*ActiveStream
	maxStreams   int
	metrics      *StreamMetrics
	logger       *zap.Logger
}

// ActiveStream represents a single content stream being shared.
type ActiveStream struct {
	ContentID      string
	Profile        *TranscodingProfile
	Viewers        map[string]*Viewer
	Position       time.Duration
	LastAccessTime time.Time
	Health         *StreamHealth
	CancelFunc     context.CancelFunc
}

// Viewer represents a per-user session watching a stream.
type Viewer struct {
	SessionID        string
	UserID           string
	ConnectedAt      time.Time
	LastPingTime     time.Time
	PlaybackPosition time.Duration
}

// StreamHealth represents the health status of a stream.
type StreamHealth struct {
	IsHealthy    bool
	LastCheck    time.Time
	ErrorCount   int
	FFmpegPID    int
	OutputBitrate int
}

// TranscodingProfile represents a transcoding profile for a stream.
type TranscodingProfile struct {
	Container       string
	Type            string
	AudioCodec      string
	VideoCodec      string
	MaxAudioChannels string
	Protocol        string
	MaxVideoBitrate int
	MaxAudioBitrate int
}

// StreamMetrics tracks stream pool statistics.
type StreamMetrics struct {
	TotalStreamsCreated int
	TotalStreamsClosed int
	TotalViewersServed int
	MaxConcurrentStreams int
}

// NewStreamManager creates a new stream manager.
func NewStreamManager(maxStreams int, logger *zap.Logger) *StreamManager {
	return &StreamManager{
		activeStreams: make(map[string]*ActiveStream),
		maxStreams:    maxStreams,
		metrics:       &StreamMetrics{},
		logger:        logger,
	}
}

// GetOrCreateStream gets an existing stream or creates a new one.
func (m *StreamManager) GetOrCreateStream(ctx context.Context, contentID string, profile *TranscodingProfile, viewerID string) (*ActiveStream, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	key := contentID + ":" + profile.Container + ":" + profile.VideoCodec

	// Check if stream exists
	if stream, exists := m.activeStreams[key]; exists {
		stream.Viewers[viewerID] = &Viewer{
			SessionID:        viewerID,
			ConnectedAt:      time.Now(),
			LastPingTime:     time.Now(),
			PlaybackPosition: stream.Position,
		}
		stream.LastAccessTime = time.Now()
		m.metrics.TotalViewersServed++
		return stream, nil
	}

	// Check if we can create a new stream
	if len(m.activeStreams) >= m.maxStreams {
		// Evict idle streams
		m.evictIdleStreams()
		if len(m.activeStreams) >= m.maxStreams {
			return nil, ErrMaxStreamsExceeded
		}
	}

	// Create new stream
	streamCtx, cancel := context.WithCancel(ctx)
	_ = streamCtx
	stream := &ActiveStream{
		ContentID:      contentID,
		Profile:        profile,
		Viewers:        map[string]*Viewer{viewerID: {SessionID: viewerID, ConnectedAt: time.Now(), LastPingTime: time.Now()}},
		LastAccessTime: time.Now(),
		Health: &StreamHealth{
			IsHealthy: true,
			LastCheck: time.Now(),
		},
		CancelFunc: cancel,
	}
	m.activeStreams[key] = stream
	m.metrics.TotalStreamsCreated++

	if m.metrics.MaxConcurrentStreams < len(m.activeStreams) {
		m.metrics.MaxConcurrentStreams = len(m.activeStreams)
	}

	m.logger.Info("created new stream", zap.String("contentID", contentID), zap.Int("viewers", len(stream.Viewers)))
	return stream, nil
}

// RemoveViewer removes a viewer from a stream.
func (m *StreamManager) RemoveViewer(contentID, profileKey, viewerID string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	key := contentID + ":" + profileKey
	if stream, exists := m.activeStreams[key]; exists {
		delete(stream.Viewers, viewerID)
		if len(stream.Viewers) == 0 {
			stream.CancelFunc()
			delete(m.activeStreams, key)
			m.metrics.TotalStreamsClosed++
			m.logger.Info("stream closed, no viewers remaining", zap.String("contentID", contentID))
		}
	}
}

// GetStreamViewers returns the number of viewers for a stream.
func (m *StreamManager) GetStreamViewers(contentID, profileKey string) int {
	m.mu.RLock()
	defer m.mu.RUnlock()

	key := contentID + ":" + profileKey
	if stream, exists := m.activeStreams[key]; exists {
		return len(stream.Viewers)
	}
	return 0
}

// GetMetrics returns the current stream metrics.
func (m *StreamManager) GetMetrics() *StreamMetrics {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return &StreamMetrics{
		TotalStreamsCreated:  m.metrics.TotalStreamsCreated,
		TotalStreamsClosed:   m.metrics.TotalStreamsClosed,
		TotalViewersServed:   m.metrics.TotalViewersServed,
		MaxConcurrentStreams: m.metrics.MaxConcurrentStreams,
	}
}

// evictIdleStreams removes the most idle streams to make room.
func (m *StreamManager) evictIdleStreams() {
	// Simple LRU eviction - remove streams with no recent access
	for key, stream := range m.activeStreams {
		if time.Since(stream.LastAccessTime) > 5*time.Minute {
			stream.CancelFunc()
			delete(m.activeStreams, key)
			m.metrics.TotalStreamsClosed++
		}
	}
}

// ErrMaxStreamsExceeded is returned when the max stream limit is reached.
var ErrMaxStreamsExceeded = &StreamError{"max concurrent streams exceeded"}

// StreamError represents a stream-related error.
type StreamError struct {
	Message string
}

func (e *StreamError) Error() string {
	return e.Message
}
