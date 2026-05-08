package media

import (
	"context"
	"fmt"
	"sync"
	"time"

	"go.uber.org/zap"
)

// StreamManager manages shared content streams with pooled transcoding.
type StreamManager struct {
	mu            sync.RWMutex
	sources       map[string]*SourceStream
	outputs       map[string]*OutputStream
	maxSourceStreams int
	maxOutputStreams int
	metrics       *StreamMetrics
	logger        *zap.Logger
}

// SourceStream represents a single remote content source (one per channel).
type SourceStream struct {
	ContentID      string
	OutputStreams  map[string]*OutputStream
	Position       time.Duration
	LastAccessTime time.Time
	Health         *StreamHealth
	CancelFunc     context.CancelFunc
}

// OutputStream represents a transcoded output at a specific resolution/profile.
type OutputStream struct {
	SourceKey      string
	Profile        *TranscodingProfile
	Viewers        map[string]*Viewer
	LastAccessTime time.Time
	Health         *StreamHealth
	CancelFunc     context.CancelFunc
}

// Viewer represents a per-user session watching an output stream.
type Viewer struct {
	SessionID        string
	UserID           string
	ConnectedAt      time.Time
	LastPingTime     time.Time
	PlaybackPosition time.Duration
}

// StreamHealth represents the health status of a stream.
type StreamHealth struct {
	IsHealthy     bool
	LastCheck     time.Time
	ErrorCount    int
	FFmpegPID     int
	OutputBitrate int
}

// TranscodingProfile represents a transcoding profile for a stream.
type TranscodingProfile struct {
	Container        string
	Type             string
	AudioCodec       string
	VideoCodec       string
	MaxAudioChannels string
	Protocol         string
	MaxVideoBitrate  int
	MaxAudioBitrate  int
}

// StreamMetrics tracks stream pool statistics.
type StreamMetrics struct {
	TotalSourceStreamsCreated int
	TotalOutputStreamsCreated int
	TotalSourceStreamsClosed  int
	TotalOutputStreamsClosed  int
	TotalViewersServed        int
	MaxConcurrentSources      int
	MaxConcurrentOutputs     int
}

func NewStreamManager(maxStreams int, logger *zap.Logger) *StreamManager {
	return &StreamManager{
		sources:           make(map[string]*SourceStream),
		outputs:           make(map[string]*OutputStream),
		maxSourceStreams:  maxStreams,
		maxOutputStreams:  maxStreams * 4,
		metrics:           &StreamMetrics{},
		logger:            logger,
	}
}

func profileKey(p *TranscodingProfile) string {
	return fmt.Sprintf("%s:%s:%d", p.Container, p.VideoCodec, p.MaxVideoBitrate)
}

func outputKey(contentID string, p *TranscodingProfile) string {
	return contentID + ":" + profileKey(p)
}

func (m *StreamManager) GetOrCreateStream(ctx context.Context, contentID string, profile *TranscodingProfile, viewerID string) (*SourceStream, *OutputStream, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	outKey := outputKey(contentID, profile)

	if out, exists := m.outputs[outKey]; exists {
		out.Viewers[viewerID] = &Viewer{
			SessionID:        viewerID,
			ConnectedAt:      time.Now(),
			LastPingTime:     time.Now(),
		}
		out.LastAccessTime = time.Now()
		m.metrics.TotalViewersServed++

		src, _ := m.sources[out.SourceKey]
		return src, out, nil
	}

	var src *SourceStream
	srcKey := contentID

	if source, exists := m.sources[srcKey]; exists {
		src = source
	} else {
		if len(m.sources) >= m.maxSourceStreams {
			m.evictIdleSources()
			if len(m.sources) >= m.maxSourceStreams {
				return nil, nil, ErrMaxSourcesExceeded
			}
		}

		_, cancel := context.WithCancel(ctx)
		src = &SourceStream{
			ContentID:      contentID,
			OutputStreams:  make(map[string]*OutputStream),
			LastAccessTime: time.Now(),
			Health: &StreamHealth{
				IsHealthy: true,
				LastCheck: time.Now(),
			},
			CancelFunc: cancel,
		}
		m.sources[srcKey] = src
		m.metrics.TotalSourceStreamsCreated++
		if m.metrics.MaxConcurrentSources < len(m.sources) {
			m.metrics.MaxConcurrentSources = len(m.sources)
		}
		m.logger.Info("created source stream", zap.String("contentID", contentID))
	}

	if len(m.outputs) >= m.maxOutputStreams {
		m.evictIdleOutputs()
		if len(m.outputs) >= m.maxOutputStreams {
			return nil, nil, ErrMaxOutputsExceeded
		}
	}

	outCtx, outCancel := context.WithCancel(ctx)
	_ = outCtx
	out := &OutputStream{
		SourceKey:      srcKey,
		Profile:        profile,
		Viewers:        map[string]*Viewer{viewerID: {SessionID: viewerID, ConnectedAt: time.Now(), LastPingTime: time.Now()}},
		LastAccessTime: time.Now(),
		Health: &StreamHealth{
			IsHealthy: true,
			LastCheck: time.Now(),
		},
		CancelFunc: outCancel,
	}
	m.outputs[outKey] = out
	src.OutputStreams[outKey] = out
	m.metrics.TotalOutputStreamsCreated++
	if m.metrics.MaxConcurrentOutputs < len(m.outputs) {
		m.metrics.MaxConcurrentOutputs = len(m.outputs)
	}

	m.metrics.TotalViewersServed++
	m.logger.Info("created output stream", zap.String("contentID", contentID), zap.String("profile", profileKey(profile)), zap.Int("viewers", len(out.Viewers)))

	return src, out, nil
}

func (m *StreamManager) RemoveViewer(contentID string, profile *TranscodingProfile, viewerID string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	outKey := outputKey(contentID, profile)

	out, exists := m.outputs[outKey]
	if !exists {
		return
	}

	delete(out.Viewers, viewerID)

	if len(out.Viewers) == 0 {
		out.CancelFunc()
		delete(m.outputs, outKey)
		m.metrics.TotalOutputStreamsClosed++

		if src, srcExists := m.sources[out.SourceKey]; srcExists {
			delete(src.OutputStreams, outKey)
			if len(src.OutputStreams) == 0 {
				src.CancelFunc()
				delete(m.sources, out.SourceKey)
				m.metrics.TotalSourceStreamsClosed++
				m.logger.Info("source stream closed, no outputs remaining", zap.String("contentID", contentID))
			}
		}
	}
}

func (m *StreamManager) GetStreamViewers(contentID string, profile *TranscodingProfile) int {
	m.mu.RLock()
	defer m.mu.RUnlock()

	outKey := outputKey(contentID, profile)
	if out, exists := m.outputs[outKey]; exists {
		return len(out.Viewers)
	}
	return 0
}

func (m *StreamManager) GetSourceViewerCount(contentID string) int {
	m.mu.RLock()
	defer m.mu.RUnlock()

	src, exists := m.sources[contentID]
	if !exists {
		return 0
	}
	count := 0
	for _, out := range src.OutputStreams {
		count += len(out.Viewers)
	}
	return count
}

func (m *StreamManager) GetMetrics() *StreamMetrics {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return &StreamMetrics{
		TotalSourceStreamsCreated: m.metrics.TotalSourceStreamsCreated,
		TotalOutputStreamsCreated: m.metrics.TotalOutputStreamsCreated,
		TotalSourceStreamsClosed:  m.metrics.TotalSourceStreamsClosed,
		TotalOutputStreamsClosed:  m.metrics.TotalOutputStreamsClosed,
		TotalViewersServed:        m.metrics.TotalViewersServed,
		MaxConcurrentSources:      m.metrics.MaxConcurrentSources,
		MaxConcurrentOutputs:      m.metrics.MaxConcurrentOutputs,
	}
}

func (m *StreamManager) evictIdleSources() {
	for key, src := range m.sources {
		if time.Since(src.LastAccessTime) > 5*time.Minute && len(src.OutputStreams) == 0 {
			src.CancelFunc()
			delete(m.sources, key)
			m.metrics.TotalSourceStreamsClosed++
		}
	}
}

func (m *StreamManager) evictIdleOutputs() {
	for key, out := range m.outputs {
		if time.Since(out.LastAccessTime) > 5*time.Minute && len(out.Viewers) == 0 {
			out.CancelFunc()
			delete(m.outputs, key)
			m.metrics.TotalOutputStreamsClosed++
		}
	}
}

var (
	ErrMaxSourcesExceeded = &StreamError{Message: "max concurrent source streams exceeded"}
	ErrMaxOutputsExceeded = &StreamError{Message: "max concurrent output streams exceeded"}
)

type StreamError struct {
	Message string
}

func (e *StreamError) Error() string {
	return e.Message
}