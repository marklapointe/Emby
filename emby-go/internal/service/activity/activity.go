package activity

import (
	"sync"
	"time"
)

type Manager struct {
	mu        sync.RWMutex
	activities map[string]*Activity
}

type Activity struct {
	ID        string    `json:"Id"`
	Name      string    `json:"Name"`
	Overview  string    `json:"Overview"`
	UserID    string    `json:"UserId,omitempty"`
	UserName  string    `json:"UserName,omitempty"`
	DateCreated time.Time `json:"DateCreated"`
	Date       time.Time `json:"Date"`
	Type       string   `json:"Type"`
	Severity   string   `json:"Severity"`
}

func NewManager() *Manager {
	return &Manager{
		activities: make(map[string]*Activity),
	}
}

func (m *Manager) GetActivities() []*Activity {
	m.mu.RLock()
	defer m.mu.RUnlock()

	activities := make([]*Activity, 0, len(m.activities))
	for _, a := range m.activities {
		activities = append(activities, a)
	}
	return activities
}

func (m *Manager) GetActivity(id string) (*Activity, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	a, ok := m.activities[id]
	return a, ok
}

func (m *Manager) LogActivity(name, overview, userID, userName, activityType, severity string) *Activity {
	m.mu.Lock()
	defer m.mu.Unlock()

	activity := &Activity{
		ID:          generateID(),
		Name:        name,
		Overview:    overview,
		UserID:      userID,
		UserName:    userName,
		DateCreated: time.Now(),
		Date:        time.Now(),
		Type:        activityType,
		Severity:    severity,
	}
	m.activities[activity.ID] = activity
	return activity
}

func (m *Manager) DeleteActivity(id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.activities, id)
	return nil
}

func generateID() string {
	return time.Now().Format("20060102150405") + "-" + randomString(8)
}

func randomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyz0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[time.Now().UnixNano()%int64(len(letters))]
	}
	return string(b)
}