package media

import (
	"context"
	"fmt"
	"math"
	"sync"
	"time"

	"go.uber.org/zap"
)

type SourceState int

const (
	SourceStateConnecting SourceState = iota
	SourceStateConnected
	SourceStateReconnecting
	SourceStateDisconnected
	SourceStateFailed
)

func (s SourceState) String() string {
	switch s {
	case SourceStateConnecting:
		return "connecting"
	case SourceStateConnected:
		return "connected"
	case SourceStateReconnecting:
		return "reconnecting"
	case SourceStateDisconnected:
		return "disconnected"
	case SourceStateFailed:
		return "failed"
	default:
		return "unknown"
	}
}

type StreamManager struct {
	mu                 sync.RWMutex
	sources            map[string]*SourceStream
	outputs            map[string]*OutputStream
	maxSourceStreams   int
	maxOutputStreams   int
	metrics            *StreamMetrics
	logger             *zap.Logger
	reconnectInterval  time.Duration
	maxReconnectAttempts int
}

type SourceStream struct {
	ContentID      string
	State         SourceState
	OutputStreams map[string]*OutputStream
	Position      time.Duration
	LastAccessTime time.Time
	Health        *StreamHealth
	CancelFunc    context.CancelFunc

	reconnectAttempts int
	reconnectMu      sync.Mutex
	waiters          int
}

type OutputStream struct {
	SourceKey      string
	Profile        *TranscodingProfile
	Viewers        map[string]*Viewer
	LastAccessTime time.Time
	Health         *StreamHealth
	CancelFunc     context.CancelFunc
	blocked        bool
}

type Viewer struct {
	SessionID        string
	UserID           string
	ConnectedAt      time.Time
	LastPingTime     time.Time
	PlaybackPosition time.Duration
}

type StreamHealth struct {
	IsHealthy     bool
	LastCheck     time.Time
	ErrorCount    int
	FFmpegPID     int
	OutputBitrate int
}

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

type StreamMetrics struct {
	TotalSourceStreamsCreated  int
	TotalOutputStreamsCreated int
	TotalSourceStreamsClosed  int
	TotalOutputStreamsClosed  int
	TotalViewersServed        int
	TotalReconnectsAttempted  int
	TotalReconnectsSucceeded  int
	TotalReconnectsFailed     int
	MaxConcurrentSources      int
	MaxConcurrentOutputs     int
}

func NewStreamManager(maxStreams int, logger *zap.Logger) *StreamManager {
	return &StreamManager{
		sources:               make(map[string]*SourceStream),
		outputs:               make(map[string]*OutputStream),
		maxSourceStreams:      maxStreams,
		maxOutputStreams:      maxStreams * 4,
		metrics:               &StreamMetrics{},
		logger:                logger,
		reconnectInterval:     2 * time.Second,
		maxReconnectAttempts:   5,
	}
}

func profileKey(p *TranscodingProfile) string {
	return fmt.Sprintf("%s:%s:%d", p.Container, p.VideoCodec, p.MaxVideoBitrate)
}

func outputKey(contentID string, p *TranscodingProfile) string {
	return contentID + ":" + profileKey(p)
}

func (m *StreamManager) GetOrCreateStream(ctx context.Context, contentID string, profile *TranscodingProfile, viewerID string) (*SourceStream, *OutputStream, error) {
	return m.getOrCreateStream(ctx, contentID, profile, viewerID, true)
}

func (m *StreamManager) getOrCreateStream(ctx context.Context, contentID string, profile *TranscodingProfile, viewerID string, shouldReconnect bool) (*SourceStream, *OutputStream, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	outKey := outputKey(contentID, profile)

	if out, exists := m.outputs[outKey]; exists {
		out.Viewers[viewerID] = &Viewer{
			SessionID:    viewerID,
			ConnectedAt:  time.Now(),
			LastPingTime: time.Now(),
		}
		out.LastAccessTime = time.Now()
		m.metrics.TotalViewersServed++

		src := m.sources[out.SourceKey]
		if src != nil && (src.State == SourceStateReconnecting || src.State == SourceStateDisconnected) {
			if shouldReconnect {
				go m.attemptReconnect(contentID)
			}
		}
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
			State:         SourceStateConnected,
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

func (m *StreamManager) attemptReconnect(contentID string) {
	m.mu.Lock()
	src, exists := m.sources[contentID]
	if !exists {
		m.mu.Unlock()
		return
	}

	src.reconnectMu.Lock()
	if src.State == SourceStateReconnecting || src.State == SourceStateConnected {
		src.reconnectMu.Unlock()
		m.mu.Unlock()
		return
	}

	src.State = SourceStateReconnecting
	src.reconnectAttempts++
	src.waiters = len(src.OutputStreams)
	waiters := src.waiters
	m.metrics.TotalReconnectsAttempted++
	m.mu.Unlock()

	m.logger.Info("attempting source reconnection", zap.String("contentID", contentID), zap.Int("attempt", src.reconnectAttempts), zap.Int("waiters", waiters))

	backoff := time.Duration(math.Min(float64(src.reconnectAttempts)*float64(m.reconnectInterval), float64(30*time.Second)))
	time.Sleep(backoff)

	m.mu.Lock()
	defer m.mu.Unlock()

	if src.State != SourceStateReconnecting {
		return
	}

	if src.reconnectAttempts >= m.maxReconnectAttempts {
		src.State = SourceStateFailed
		m.metrics.TotalReconnectsFailed++
		m.logger.Error("source reconnection failed permanently", zap.String("contentID", contentID), zap.Int("attempts", src.reconnectAttempts))
		return
	}

	src.State = SourceStateConnected
	src.reconnectAttempts = 0
	src.Health.IsHealthy = true
	src.Health.LastCheck = time.Now()
	src.waiters = 0
	m.metrics.TotalReconnectsSucceeded++
	m.logger.Info("source reconnection succeeded", zap.String("contentID", contentID))
}

func (m *StreamManager) SourceDisconnected(contentID string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	src, exists := m.sources[contentID]
	if !exists {
		return
	}

	if src.State == SourceStateDisconnected || src.State == SourceStateFailed {
		return
	}

	src.State = SourceStateDisconnected
	src.Health.IsHealthy = false
	src.reconnectAttempts = 0
	m.logger.Warn("source disconnected", zap.String("contentID", contentID))
}

func (m *StreamManager) SourceReconnected(contentID string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	src, exists := m.sources[contentID]
	if !exists {
		return
	}

	src.State = SourceStateConnected
	src.Health.IsHealthy = true
	src.Health.LastCheck = time.Now()
	src.reconnectAttempts = 0
	m.logger.Info("source reconnected", zap.String("contentID", contentID))
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

func (m *StreamManager) GetSourceState(contentID string) SourceState {
	m.mu.RLock()
	defer m.mu.RUnlock()

	src, exists := m.sources[contentID]
	if !exists {
		return SourceStateDisconnected
	}
	return src.State
}

func (m *StreamManager) GetMetrics() *StreamMetrics {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return &StreamMetrics{
		TotalSourceStreamsCreated:  m.metrics.TotalSourceStreamsCreated,
		TotalOutputStreamsCreated: m.metrics.TotalOutputStreamsCreated,
		TotalSourceStreamsClosed:  m.metrics.TotalSourceStreamsClosed,
		TotalOutputStreamsClosed:  m.metrics.TotalOutputStreamsClosed,
		TotalViewersServed:        m.metrics.TotalViewersServed,
		TotalReconnectsAttempted:  m.metrics.TotalReconnectsAttempted,
		TotalReconnectsSucceeded:  m.metrics.TotalReconnectsSucceeded,
		TotalReconnectsFailed:     m.metrics.TotalReconnectsFailed,
		MaxConcurrentSources:      m.metrics.MaxConcurrentSources,
		MaxConcurrentOutputs:     m.metrics.MaxConcurrentOutputs,
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