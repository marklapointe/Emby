package websocket

import (
	"fmt"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

// Client represents a connected WebSocket client.
type Client struct {
	ID        string          `json:"id"`
	Conn      *websocket.Conn `json:"-"`
	Send      chan []byte     `json:"-"`
	LastActive time.Time       `json:"lastActive"`
	IsActive  bool            `json:"isActive"`
}

// Manager handles WebSocket connections.
type Manager struct {
	clients   map[string]*Client
	mu        sync.RWMutex
	broadcast chan []byte
	logger    *zap.Logger
}

// NewManager creates a new WebSocket manager.
func NewManager(logger *zap.Logger) *Manager {
	return &Manager{
		clients:   make(map[string]*Client),
		broadcast: make(chan []byte),
		logger:    logger,
	}
}

// AddClient adds a new WebSocket client.
func (m *Manager) AddClient(conn *websocket.Conn) (*Client, error) {
	clientID := fmt.Sprintf("ws-%d", time.Now().UnixNano())
	client := &Client{
		ID:         clientID,
		Conn:       conn,
		Send:       make(chan []byte, 256),
		LastActive: time.Now(),
		IsActive:   true,
	}

	m.mu.Lock()
	m.clients[clientID] = client
	m.mu.Unlock()

	m.logger.Info("WebSocket client connected", zap.String("clientID", clientID))

	// Start read and write loops
	go m.readLoop(client)
	go m.writeLoop(client)

	return client, nil
}

// RemoveClient removes a WebSocket client.
func (m *Manager) RemoveClient(clientID string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if client, ok := m.clients[clientID]; ok {
		client.IsActive = false
		close(client.Send)
		delete(m.clients, clientID)
		m.logger.Info("WebSocket client disconnected", zap.String("clientID", clientID))
	}
}

// Broadcast sends a message to all connected clients.
func (m *Manager) Broadcast(message []byte) {
	m.broadcast <- message
}

// GetClientCount returns the number of connected clients.
func (m *Manager) GetClientCount() int {
	m.mu.RLock()
	defer m.mu.RUnlock()

	count := 0
	for _, client := range m.clients {
		if client.IsActive {
			count++
		}
	}

	return count
}

// GetActiveClients returns all active clients.
func (m *Manager) GetActiveClients() []*Client {
	m.mu.RLock()
	defer m.mu.RUnlock()

	var clients []*Client
	for _, client := range m.clients {
		if client.IsActive {
			clients = append(clients, client)
		}
	}

	return clients
}

// readLoop reads messages from a client.
func (m *Manager) readLoop(client *Client) {
	defer func() {
		m.RemoveClient(client.ID)
	}()

	for {
		_, message, err := client.Conn.ReadMessage()
		if err != nil {
			m.logger.Error("WebSocket read error", zap.Error(err))
			break
		}

		client.LastActive = time.Now()
		m.logger.Debug("WebSocket message received", zap.String("clientID", client.ID))

		// Process the message
		m.processMessage(client, message)
	}
}

// writeLoop writes messages to a client.
func (m *Manager) writeLoop(client *Client) {
	for {
		message, ok := <-client.Send
		if !ok {
			client.Conn.Close()
			return
		}

		client.LastActive = time.Now()

		err := client.Conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			m.logger.Error("WebSocket write error", zap.Error(err))
			break
		}
	}
}

// processMessage processes a message from a client.
func (m *Manager) processMessage(client *Client, message []byte) {
	// For now, just log the message
	_ = client
	_ = message
}

// StartBroadcastLoop starts the broadcast loop.
func (m *Manager) StartBroadcastLoop() {
	go func() {
		for {
			message := <-m.broadcast
			m.mu.RLock()
			for _, client := range m.clients {
				if client.IsActive {
					select {
					case client.Send <- message:
					default:
						close(client.Send)
						client.IsActive = false
					}
				}
			}
			m.mu.RUnlock()
		}
	}()
}
