package livetv

import (
	"sync"
	"time"

	"go.uber.org/zap"
)

type Manager struct {
	logger          *zap.Logger
	mu              sync.RWMutex
	channels        map[string]*Channel
	programs        map[string]*Program
	recordings     map[string]*Recording
	timers          map[string]*Timer
	isEnabled       bool
}

type Channel struct {
	ID            string    `json:"Id"`
	Name          string    `json:"Name"`
	Number        int       `json:"ChannelNumber"`
	CallSign      string    `json:"CallSign"`
	Network       string    `json:"Network"`
	ImageURL      string    `json:"ImageUrl"`
	ChannelType   string    `json:"ChannelType"`
	IsFavorite    bool      `json:"IsFavorite"`
	ChannelNumber string     `json:"ChannelNumber"`
}

type Program struct {
	ID             string    `json:"Id"`
	ChannelID      string    `json:"ChannelId"`
	Name           string    `json:"Name"`
	Overview       string    `json:"Overview"`
	StartDate      time.Time `json:"StartDate"`
	EndDate        time.Time `json:"EndDate"`
	ServiceName    string    `json:"ServiceName"`
	ChannelName    string    `json:"ChannelName"`
	EpisodeTitle   string    `json:"EpisodeTitle"`
	IsKids         bool      `json:"IsKids"`
	IsMovie        bool      `json:"IsMovie"`
	IsSports       bool      `json:"IsSports"`
	IsNews         bool      `json:"IsNews"`
	IsSeries       bool      `json:"IsSeries"`
	IsLive         bool      `json:"IsLive"`
	IsPremiere     bool      `json:"IsPremiere"`
	AudioCodec     string    `json:"AudioCodec"`
	VideoCodec     string    `json:"VideoCodec"`
	Width          int       `json:"Width"`
	Height         int       `json:"Height"`
	AspectRatio    string    `json:"AspectRatio"`
	RunTimeTicks   int64     `json:"RunTimeTicks"`
}

type Recording struct {
	ID           string    `json:"Id"`
	ChannelID    string    `json:"ChannelId"`
	ProgramID    string    `json:"ProgramId"`
	Name         string    `json:"Name"`
	Overview     string    `json:"Overview"`
	StartDate    time.Time `json:"StartDate"`
	EndDate      time.Time `json:"EndDate"`
	Status       string    `json:"Status"`
	FilePath     string    `json:"FilePath"`
	Duration     int64     `json:"Duration"`
	Priority     int       `json:"Priority"`
	RecordingUID string    `json:"RecordingUid"`
}

type Timer struct {
	ID           string    `json:"Id"`
	ChannelID    string    `json:"ChannelId"`
	ProgramID    string    `json:"ProgramId"`
	Name         string    `json:"Name"`
	StartDate    time.Time `json:"StartDate"`
	EndDate      time.Time `json:"EndDate"`
	Priority     int       `json:"Priority"`
	PrePadding   int       `json:"PrePadding"`
	PostPadding  int       `json:"PostPadding"`
	DayPattern   string    `json:"DayPattern"`
	TimerType    string    `json:"TimerType"`
	IsPending    bool      `json:"IsPending"`
}

type SeriesTimer struct {
	ID              string    `json:"Id"`
	ChannelID      string    `json:"ChannelId"`
	ProgramName     string    `json:"ProgramName"`
	StartDate       time.Time `json:"StartDate"`
	EndDate         time.Time `json:"EndDate"`
	Priority        int       `json:"Priority"`
	PrePadding      int       `json:"PrePadding"`
	PostPadding     int       `json:"PostPadding"`
	DayPattern      string    `json:"DayPattern"`
	RecordAnyTime   bool      `json:"RecordAnyTime"`
	RecordAnyChannel bool     `json:"RecordAnyChannel"`
	SeriesTimerID   string    `json:"SeriesTimerId"`
}

func NewManager(logger *zap.Logger) *Manager {
	if logger == nil {
		logger = zap.NewNop()
	}
	return &Manager{
		logger:     logger,
		channels:   make(map[string]*Channel),
		programs:   make(map[string]*Program),
		recordings: make(map[string]*Recording),
		timers:     make(map[string]*Timer),
		isEnabled:  false,
	}
}

func (m *Manager) Enable() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.isEnabled = true
	m.logger.Info("LiveTV manager enabled")
}

func (m *Manager) Disable() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.isEnabled = false
	m.logger.Info("LiveTV manager disabled")
}

func (m *Manager) IsEnabled() bool {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.isEnabled
}

func (m *Manager) GetChannels() []*Channel {
	m.mu.RLock()
	defer m.mu.RUnlock()

	channels := make([]*Channel, 0, len(m.channels))
	for _, ch := range m.channels {
		channels = append(channels, ch)
	}
	return channels
}

func (m *Manager) GetChannel(id string) (*Channel, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	ch, ok := m.channels[id]
	return ch, ok
}

func (m *Manager) AddChannel(ch *Channel) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.channels[ch.ID] = ch
}

func (m *Manager) GetPrograms(channelID string) []*Program {
	m.mu.RLock()
	defer m.mu.RUnlock()

	var programs []*Program
	for _, p := range m.programs {
		if channelID == "" || p.ChannelID == channelID {
			programs = append(programs, p)
		}
	}
	return programs
}

func (m *Manager) GetProgram(id string) (*Program, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	p, ok := m.programs[id]
	return p, ok
}

func (m *Manager) GetRecordings() []*Recording {
	m.mu.RLock()
	defer m.mu.RUnlock()

	recordings := make([]*Recording, 0, len(m.recordings))
	for _, r := range m.recordings {
		recordings = append(recordings, r)
	}
	return recordings
}

func (m *Manager) GetRecording(id string) (*Recording, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	r, ok := m.recordings[id]
	return r, ok
}

func (m *Manager) AddRecording(r *Recording) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.recordings[r.ID] = r
}

func (m *Manager) DeleteRecording(id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, ok := m.recordings[id]; !ok {
		return nil
	}
	delete(m.recordings, id)
	return nil
}

func (m *Manager) GetTimers() []*Timer {
	m.mu.RLock()
	defer m.mu.RUnlock()

	timers := make([]*Timer, 0, len(m.timers))
	for _, t := range m.timers {
		timers = append(timers, t)
	}
	return timers
}

func (m *Manager) AddTimer(t *Timer) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.timers[t.ID] = t
}

func (m *Manager) DeleteTimer(id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.timers, id)
	return nil
}
