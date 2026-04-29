package session

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/emby/emby-go/internal/config"
	"go.uber.org/zap"
)

// Upgrader handles WebSocket connections.
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// WebSocketManager manages WebSocket connections.
type WebSocketManager struct {
	mu          sync.RWMutex
	connections map[string]*WebSocketConnection
	logger      *zap.Logger
	config      *config.Config
}

// WebSocketConnection represents a single WebSocket connection.
type WebSocketConnection struct {
	mu       sync.RWMutex
	conn     *websocket.Conn
	id       string
	userID   string
	lastPing time.Time
	logger   *zap.Logger
}

// NewWebSocketManager creates a new WebSocket manager.
func NewWebSocketManager(cfg *config.Config, logger *zap.Logger) *WebSocketManager {
	return &WebSocketManager{
		connections: make(map[string]*WebSocketConnection),
		logger:      logger,
		config:      cfg,
	}
}

// Upgrade upgrades an HTTP connection to WebSocket.
func (m *WebSocketManager) Upgrade(w http.ResponseWriter, r *http.Request) (*WebSocketConnection, error) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return nil, fmt.Errorf("upgrade websocket: %w", err)
	}

	wsConn := &WebSocketConnection{
		conn:     conn,
		id:       fmt.Sprintf("ws-%d", time.Now().UnixNano()),
		lastPing: time.Now(),
		logger:   m.logger,
	}

	m.mu.Lock()
	m.connections[wsConn.id] = wsConn
	m.mu.Unlock()

	m.logger.Info("websocket connected", zap.String("id", wsConn.id))
	return wsConn, nil
}

// Send sends a message to a specific WebSocket connection.
func (m *WebSocketManager) Send(id string, message interface{}) error {
	m.mu.RLock()
	conn, exists := m.connections[id]
	m.mu.RUnlock()

	if !exists {
		return fmt.Errorf("connection not found: %s", id)
	}

	data, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("marshal message: %w", err)
	}

	conn.mu.Lock()
	defer conn.mu.Unlock()

	return conn.conn.WriteMessage(websocket.TextMessage, data)
}

// Broadcast sends a message to all connected WebSocket clients.
func (m *WebSocketManager) Broadcast(message interface{}) {
	data, err := json.Marshal(message)
	if err != nil {
		m.logger.Error("marshal broadcast message", zap.Error(err))
		return
	}

	m.mu.RLock()
	defer m.mu.RUnlock()

	for id, conn := range m.connections {
		conn.mu.Lock()
		if err := conn.conn.WriteMessage(websocket.TextMessage, data); err != nil {
			m.logger.Error("broadcast to websocket", zap.String("id", id), zap.Error(err))
		}
		conn.mu.Unlock()
	}
}

// RemoveConnection removes a WebSocket connection.
func (m *WebSocketManager) RemoveConnection(id string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if conn, exists := m.connections[id]; exists {
		conn.conn.Close()
		delete(m.connections, id)
		m.logger.Info("websocket disconnected", zap.String("id", id))
	}
}

// GetConnectionCount returns the number of active WebSocket connections.
func (m *WebSocketManager) GetConnectionCount() int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return len(m.connections)
}

// StartHeartbeat starts the WebSocket heartbeat loop.
func (m *WebSocketManager) StartHeartbeat() {
	go func() {
		ticker := time.NewTicker(30 * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			m.mu.RLock()
			for id, conn := range m.connections {
				conn.mu.Lock()
				if time.Since(conn.lastPing) > 5*time.Minute {
					if err := conn.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
						m.logger.Error("websocket ping failed", zap.String("id", id), zap.Error(err))
						conn.mu.Unlock()
						m.RemoveConnection(id)
						continue
					}
				}
				conn.mu.Unlock()
			}
			m.mu.RUnlock()
		}
	}()
}
