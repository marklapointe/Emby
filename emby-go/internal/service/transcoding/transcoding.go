package transcoding

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"sync"
	"time"

	"go.uber.org/zap"
)

// Transcoder handles media transcoding operations.
type Transcoder struct {
	mu         sync.RWMutex
	activeJobs map[string]*TranscodeJob
	logger     *zap.Logger
	ffmpegPath string
}

// TranscodeJob represents a transcoding job.
type TranscodeJob struct {
	ID           string
	SourcePath   string
	OutputPath   string
	Profile      *TranscodeProfile
	Status       string // pending, running, completed, failed
	Progress     float64
	StartTime    time.Time
	EndTime      time.Time
	FFmpegPID    int
	CancelFunc   context.CancelFunc
	Logger       *zap.Logger
}

// TranscodeProfile represents a transcoding profile.
type TranscodeProfile struct {
	Container       string
	VideoCodec      string
	VideoBitrate    int
	MaxVideoBitrate int
	AudioCodec      string
	AudioBitrate    int
	MaxAudioBitrate int
	MaxAudioChannels string
	Width           int
	Height          int
	MaxFrameRate    int
	Protocol        string
	AnalyzeDuration int
	ProbeSize       string
}

// NewTranscoder creates a new transcoder.
func NewTranscoder(logger *zap.Logger) *Transcoder {
	return &Transcoder{
		activeJobs: make(map[string]*TranscodeJob),
		logger:     logger,
		ffmpegPath: "ffmpeg",
	}
}

// SetFFmpegPath sets the path to the ffmpeg binary.
func (t *Transcoder) SetFFmpegPath(path string) {
	t.ffmpegPath = path
}

// StartTranscode starts a new transcoding job.
func (t *Transcoder) StartTranscode(ctx context.Context, sourcePath, outputPath string, profile *TranscodeProfile) (*TranscodeJob, error) {
	jobID := fmt.Sprintf("transcode-%d", time.Now().UnixNano())

	job := &TranscodeJob{
		ID:         jobID,
		SourcePath: sourcePath,
		OutputPath: outputPath,
		Profile:    profile,
		Status:     "pending",
		StartTime:  time.Now(),
		Logger:     t.logger,
	}

	t.mu.Lock()
	t.activeJobs[jobID] = job
	t.mu.Unlock()

	// Build ffmpeg command
	cmd := t.buildFFmpegCommand(sourcePath, outputPath, profile)

	// Start ffmpeg
	cmdCtx, cancel := context.WithCancel(ctx)
	job.CancelFunc = cancel

	cmd = exec.CommandContext(cmdCtx, t.ffmpegPath, cmd.Args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		t.mu.Lock()
		job.Status = "failed"
		t.mu.Unlock()
		return nil, fmt.Errorf("start ffmpeg: %w", err)
	}

	t.mu.Lock()
	job.Status = "running"
	job.FFmpegPID = cmd.Process.Pid
	t.mu.Unlock()

	// Wait for completion in background
	go func() {
		err := cmd.Wait()
		t.mu.Lock()
		job.EndTime = time.Now()
		if err != nil {
			job.Status = "failed"
			t.logger.Error("transcode failed", zap.String("jobId", jobID), zap.Error(err))
		} else {
			job.Status = "completed"
			t.logger.Info("transcode completed", zap.String("jobId", jobID))
		}
		t.mu.Unlock()
	}()

	t.logger.Info("transcode started",
		zap.String("jobId", jobID),
		zap.String("source", sourcePath),
		zap.String("output", outputPath),
	)

	return job, nil
}

// buildFFmpegCommand builds the ffmpeg command for transcoding.
func (t *Transcoder) buildFFmpegCommand(source, output string, profile *TranscodeProfile) *exec.Cmd {
	var args []string

	// Input
	args = append(args, "-i", source)

	// Analyze duration
	if profile.AnalyzeDuration > 0 {
		args = append(args, "-analyzeduration", fmt.Sprintf("%d", profile.AnalyzeDuration*1000000))
	}

	// Probe size
	if profile.ProbeSize != "" {
		args = append(args, "-probesize", profile.ProbeSize)
	}

	// Video settings
	if profile.VideoCodec != "" {
		args = append(args, "-c:v", profile.VideoCodec)
	}
	if profile.VideoBitrate > 0 {
		args = append(args, "-b:v", fmt.Sprintf("%dk", profile.VideoBitrate))
	}
	if profile.MaxVideoBitrate > 0 {
		args = append(args, "-maxrate", fmt.Sprintf("%dk", profile.MaxVideoBitrate))
	}
	if profile.Width > 0 || profile.Height > 0 {
		args = append(args, "-s", fmt.Sprintf("%dx%d", profile.Width, profile.Height))
	}
	if profile.MaxFrameRate > 0 {
		args = append(args, "-r", fmt.Sprintf("%d", profile.MaxFrameRate))
	}

	// Audio settings
	if profile.AudioCodec != "" {
		args = append(args, "-c:a", profile.AudioCodec)
	}
	if profile.AudioBitrate > 0 {
		args = append(args, "-b:a", fmt.Sprintf("%dk", profile.AudioBitrate))
	}
	if profile.MaxAudioBitrate > 0 {
		args = append(args, "-maxrate:a", fmt.Sprintf("%dk", profile.MaxAudioBitrate))
	}
	if profile.MaxAudioChannels != "" {
		args = append(args, "-ac", profile.MaxAudioChannels)
	}

	// Output
	args = append(args, "-f", profile.Container, output)

	return exec.Command(args[0], args[1:]...)
}

// GetJob returns a transcode job by ID.
func (t *Transcoder) GetJob(id string) (*TranscodeJob, bool) {
	t.mu.RLock()
	defer t.mu.RUnlock()

	job, exists := t.activeJobs[id]
	return job, exists
}

// GetActiveJobs returns all active transcode jobs.
func (t *Transcoder) GetActiveJobs() []*TranscodeJob {
	t.mu.RLock()
	defer t.mu.RUnlock()

	jobs := make([]*TranscodeJob, 0, len(t.activeJobs))
	for _, job := range t.activeJobs {
		jobs = append(jobs, job)
	}
	return jobs
}

// CancelJob cancels a transcode job.
func (t *Transcoder) CancelJob(id string) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	job, exists := t.activeJobs[id]
	if !exists {
		return fmt.Errorf("job not found: %s", id)
	}

	if job.CancelFunc != nil {
		job.CancelFunc()
		job.Status = "cancelled"
	}

	return nil
}

// RemoveJob removes a completed job.
func (t *Transcoder) RemoveJob(id string) {
	t.mu.Lock()
	defer t.mu.Unlock()

	delete(t.activeJobs, id)
}

// GetActiveJobCount returns the number of active transcode jobs.
func (t *Transcoder) GetActiveJobCount() int {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return len(t.activeJobs)
}
