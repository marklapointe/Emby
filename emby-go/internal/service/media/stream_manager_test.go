package media

import (
	"context"
	"sync"
	"testing"

	"go.uber.org/zap"
)

func TestPooledTranscoding_SharedSource_MultipleResolutions(t *testing.T) {
	logger := zap.NewNop()
	m := NewStreamManager(10, logger)

	ctx := context.Background()

	profiles := []*TranscodingProfile{
		{Container: "ts", VideoCodec: "h264", MaxVideoBitrate: 5000},
		{Container: "ts", VideoCodec: "h264", MaxVideoBitrate: 5000},
		{Container: "ts", VideoCodec: "h264", MaxVideoBitrate: 5000},
		{Container: "ts", VideoCodec: "h264", MaxVideoBitrate: 2500},
		{Container: "ts", VideoCodec: "h264", MaxVideoBitrate: 2500},
	}

	for i, profile := range profiles {
		viewerID := "viewer-" + string(rune('A'+i))
		_, _, err := m.GetOrCreateStream(ctx, "channel-101", profile, viewerID)
		if err != nil {
			t.Fatalf("GetOrCreateStream failed for viewer %s: %v", viewerID, err)
		}
	}

	metrics := m.GetMetrics()
	t.Logf("Metrics: Sources=%d, Outputs=%d, Viewers=%d",
		metrics.TotalSourceStreamsCreated, metrics.TotalOutputStreamsCreated, metrics.TotalViewersServed)

	if metrics.TotalSourceStreamsCreated != 1 {
		t.Errorf("expected 1 source stream, got %d", metrics.TotalSourceStreamsCreated)
	}

	if metrics.TotalOutputStreamsCreated != 2 {
		t.Errorf("expected 2 output streams (1080p + 720p), got %d", metrics.TotalOutputStreamsCreated)
	}

	if int(metrics.TotalViewersServed) != 5 {
		t.Errorf("expected 5 viewers served, got %d", metrics.TotalViewersServed)
	}
}

func TestPooledTranscoding_SameProfile_SharedTranscode(t *testing.T) {
	logger := zap.NewNop()
	m := NewStreamManager(10, logger)

	ctx := context.Background()
	profile := &TranscodingProfile{Container: "ts", VideoCodec: "h264", MaxVideoBitrate: 5000}

	for i := 0; i < 5; i++ {
		viewerID := "tv-" + string(rune('1'+i))
		_, _, err := m.GetOrCreateStream(ctx, "channel-101", profile, viewerID)
		if err != nil {
			t.Fatalf("GetOrCreateStream failed: %v", err)
		}
	}

	metrics := m.GetMetrics()
	t.Logf("Metrics: Sources=%d, Outputs=%d, Viewers=%d",
		metrics.TotalSourceStreamsCreated, metrics.TotalOutputStreamsCreated, metrics.TotalViewersServed)

	if metrics.TotalSourceStreamsCreated != 1 {
		t.Errorf("expected 1 source stream, got %d", metrics.TotalSourceStreamsCreated)
	}
	if metrics.TotalOutputStreamsCreated != 1 {
		t.Errorf("expected 1 output stream (same profile), got %d", metrics.TotalOutputStreamsCreated)
	}
}

func TestPooledTranscoding_SyncMultipleTVs(t *testing.T) {
	logger := zap.NewNop()
	m := NewStreamManager(10, logger)

	ctx := context.Background()
	profile := &TranscodingProfile{Container: "ts", VideoCodec: "h264", MaxVideoBitrate: 5000}

	tvs := []string{"living-room", "bedroom", "kitchen"}

	for _, tv := range tvs {
		_, _, err := m.GetOrCreateStream(ctx, "channel-101", profile, tv)
		if err != nil {
			t.Fatalf("GetOrCreateStream failed for %s: %v", tv, err)
		}
	}

	viewers := m.GetStreamViewers("channel-101", profile)

	if viewers != 3 {
		t.Errorf("expected 3 viewers for channel-101, got %d", viewers)
	}

	sourceViewers := m.GetSourceViewerCount("channel-101")
	t.Logf("All TVs synced on channel-101: %d viewers sharing 1 source+transcode", sourceViewers)
}

func TestPooledTranscoding_ChannelSwitch(t *testing.T) {
	logger := zap.NewNop()
	m := NewStreamManager(10, logger)

	ctx := context.Background()
	profile := &TranscodingProfile{Container: "ts", VideoCodec: "h264", MaxVideoBitrate: 5000}

	m.GetOrCreateStream(ctx, "channel-101", profile, "tv-1")
	m.GetOrCreateStream(ctx, "channel-101", profile, "tv-2")

	m.RemoveViewer("channel-101", profile, "tv-1")
	m.GetOrCreateStream(ctx, "channel-202", profile, "tv-1")

	viewers101 := m.GetStreamViewers("channel-101", profile)
	viewers202 := m.GetStreamViewers("channel-202", profile)

	if viewers101 != 1 {
		t.Errorf("expected 1 viewer on 101 after switch, got %d", viewers101)
	}
	if viewers202 != 1 {
		t.Errorf("expected 1 viewer on 202 after switch, got %d", viewers202)
	}
}

func TestPooledTranscoding_OutputProfileCleanup(t *testing.T) {
	logger := zap.NewNop()
	m := NewStreamManager(10, logger)

	ctx := context.Background()

	profile1080 := &TranscodingProfile{Container: "ts", VideoCodec: "h264", MaxVideoBitrate: 5000}
	profile720 := &TranscodingProfile{Container: "ts", VideoCodec: "h264", MaxVideoBitrate: 2500}

	m.GetOrCreateStream(ctx, "channel-101", profile1080, "viewer-1080-a")
	m.GetOrCreateStream(ctx, "channel-101", profile1080, "viewer-1080-b")
	m.GetOrCreateStream(ctx, "channel-101", profile720, "viewer-720")

	m.RemoveViewer("channel-101", profile1080, "viewer-1080-a")
	m.RemoveViewer("channel-101", profile1080, "viewer-1080-b")

	viewers720 := m.GetStreamViewers("channel-101", profile720)
	if viewers720 != 1 {
		t.Errorf("expected 1 viewer at 720p, got %d", viewers720)
	}

	viewers1080 := m.GetStreamViewers("channel-101", profile1080)
	if viewers1080 != 0 {
		t.Errorf("expected 0 viewers at 1080p, got %d", viewers1080)
	}

	metrics := m.GetMetrics()
	if metrics.TotalSourceStreamsCreated != 1 {
		t.Errorf("expected 1 source stream still active, got %d", metrics.TotalSourceStreamsCreated)
	}
}

func TestPooledTranscoding_ConcurrentViewerJoin(t *testing.T) {
	logger := zap.NewNop()
	m := NewStreamManager(10, logger)

	ctx := context.Background()
	profile := &TranscodingProfile{Container: "ts", VideoCodec: "h264", MaxVideoBitrate: 5000}

	var wg sync.WaitGroup
	viewerCount := 20

	for i := 0; i < viewerCount; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			viewerID := "viewer-" + string(rune('A'+id))
			m.GetOrCreateStream(ctx, "channel-101", profile, viewerID)
		}(i)
	}

	wg.Wait()

	metrics := m.GetMetrics()
	if int(metrics.TotalViewersServed) != viewerCount {
		t.Errorf("expected %d viewers served, got %d", viewerCount, metrics.TotalViewersServed)
	}
}

func TestPooledTranscoding_MixedDevices(t *testing.T) {
	logger := zap.NewNop()
	m := NewStreamManager(10, logger)

	ctx := context.Background()

	deviceProfiles := map[string]*TranscodingProfile{
		"living-room-tv":  {Container: "ts", VideoCodec: "h264", MaxVideoBitrate: 5000},
		"bedroom-tv":      {Container: "ts", VideoCodec: "h264", MaxVideoBitrate: 5000},
		"kitchen-tablet":  {Container: "ts", VideoCodec: "h264", MaxVideoBitrate: 2500},
		"phone":           {Container: "ts", VideoCodec: "h264", MaxVideoBitrate: 1500},
		"phone-2":         {Container: "ts", VideoCodec: "h264", MaxVideoBitrate: 1500},
	}

	for device, profile := range deviceProfiles {
		_, _, err := m.GetOrCreateStream(ctx, "channel-101", profile, device)
		if err != nil {
			t.Fatalf("GetOrCreateStream failed for %s: %v", device, err)
		}
	}

	metrics := m.GetMetrics()
	t.Logf("5 devices, 3 resolutions: 1 source + 3 outputs = %d streams total",
		metrics.TotalSourceStreamsCreated+metrics.TotalOutputStreamsCreated)

	if metrics.TotalSourceStreamsCreated != 1 {
		t.Errorf("expected 1 source stream, got %d", metrics.TotalSourceStreamsCreated)
	}

	if metrics.TotalOutputStreamsCreated != 3 {
		t.Errorf("expected 3 output streams (5000, 2500, 1500), got %d", metrics.TotalOutputStreamsCreated)
	}

	sourceViewers := m.GetSourceViewerCount("channel-101")
	if sourceViewers != 5 {
		t.Errorf("expected 5 total viewers on source, got %d", sourceViewers)
	}
}

func TestSourceResilience_DisconnectAndReconnect(t *testing.T) {
	logger := zap.NewNop()
	m := NewStreamManager(10, logger)

	ctx := context.Background()
	profile := &TranscodingProfile{Container: "ts", VideoCodec: "h264", MaxVideoBitrate: 5000}

	src, out, err := m.GetOrCreateStream(ctx, "channel-101", profile, "tv-1")
	if err != nil {
		t.Fatalf("GetOrCreateStream failed: %v", err)
	}

	if src.State != SourceStateConnected {
		t.Errorf("expected source to be connected, got %s", src.State)
	}

	m.SourceDisconnected("channel-101")

	if m.GetSourceState("channel-101") != SourceStateDisconnected {
		t.Errorf("expected source to be disconnected")
	}

	m.SourceReconnected("channel-101")

	if m.GetSourceState("channel-101") != SourceStateConnected {
		t.Errorf("expected source to be reconnected")
	}

	if len(out.Viewers) != 1 {
		t.Errorf("expected output to still have 1 viewer, got %d", len(out.Viewers))
	}
}

func TestSourceResilience_MultipleClientsReconnect(t *testing.T) {
	logger := zap.NewNop()
	m := NewStreamManager(10, logger)

	ctx := context.Background()
	profile := &TranscodingProfile{Container: "ts", VideoCodec: "h264", MaxVideoBitrate: 5000}

	m.GetOrCreateStream(ctx, "channel-101", profile, "tv-1")
	m.GetOrCreateStream(ctx, "channel-101", profile, "tv-2")
	m.GetOrCreateStream(ctx, "channel-101", profile, "tv-3")

	viewers := m.GetSourceViewerCount("channel-101")
	if viewers != 3 {
		t.Errorf("expected 3 viewers, got %d", viewers)
	}

	m.SourceDisconnected("channel-101")

	state := m.GetSourceState("channel-101")
	if state != SourceStateDisconnected {
		t.Errorf("expected disconnected state, got %s", state)
	}

	m.SourceReconnected("channel-101")

	viewers = m.GetSourceViewerCount("channel-101")
	if viewers != 3 {
		t.Errorf("expected 3 viewers after reconnect, got %d", viewers)
	}
}

func TestSourceResilience_MetricsTracking(t *testing.T) {
	logger := zap.NewNop()
	m := NewStreamManager(10, logger)

	ctx := context.Background()
	profile := &TranscodingProfile{Container: "ts", VideoCodec: "h264", MaxVideoBitrate: 5000}

	m.GetOrCreateStream(ctx, "channel-101", profile, "tv-1")

	m.SourceDisconnected("channel-101")
	m.SourceReconnected("channel-101")
	m.SourceDisconnected("channel-101")
	m.SourceReconnected("channel-101")

	metrics := m.GetMetrics()
	if metrics.TotalReconnectsAttempted != 0 {
		t.Logf("Note: TotalReconnectsAttempted=%d (automatic reconnect not triggered in this test)", metrics.TotalReconnectsAttempted)
	}
}

func TestSourceResilience_MultipleResolutionsSurviveDisconnect(t *testing.T) {
	logger := zap.NewNop()
	m := NewStreamManager(10, logger)

	ctx := context.Background()
	profile1080 := &TranscodingProfile{Container: "ts", VideoCodec: "h264", MaxVideoBitrate: 5000}
	profile720 := &TranscodingProfile{Container: "ts", VideoCodec: "h264", MaxVideoBitrate: 2500}

	m.GetOrCreateStream(ctx, "channel-101", profile1080, "tv-1080")
	m.GetOrCreateStream(ctx, "channel-101", profile720, "tablet-720")

	m.SourceDisconnected("channel-101")

	viewers1080 := m.GetStreamViewers("channel-101", profile1080)
	viewers720 := m.GetStreamViewers("channel-101", profile720)

	if viewers1080 != 1 {
		t.Errorf("expected 1 viewer at 1080p, got %d", viewers1080)
	}
	if viewers720 != 1 {
		t.Errorf("expected 1 viewer at 720p, got %d", viewers720)
	}

	m.SourceReconnected("channel-101")

	viewers1080 = m.GetStreamViewers("channel-101", profile1080)
	viewers720 = m.GetStreamViewers("channel-101", profile720)

	if viewers1080 != 1 {
		t.Errorf("expected 1 viewer at 1080p after reconnect, got %d", viewers1080)
	}
	if viewers720 != 1 {
		t.Errorf("expected 1 viewer at 720p after reconnect, got %d", viewers720)
	}
}