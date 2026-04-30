package dlna

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"sync"

	"github.com/emby/emby-go/internal/dlna/xml"
	"go.uber.org/zap"
)

// Server represents the DLNA/SSDP server.
type Server struct {
	port     int
	logger   *zap.Logger
	mu       sync.RWMutex
	listener net.Listener
	running  bool
}

// NewServer creates a new DLNA server.
func NewServer(port int, logger *zap.Logger) *Server {
	return &Server{
		port:   port,
		logger: logger,
	}
}

// Start starts the DLNA/SSDP server.
func (s *Server) Start() error {
	addr := fmt.Sprintf("0.0.0.0:%d", s.port)

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("failed to start DLNA server: %w", err)
	}

	s.mu.Lock()
	s.listener = listener
	s.running = true
	s.mu.Unlock()

	s.logger.Info("DLNA server started", zap.String("addr", addr))

	go s.serve()
	return nil
}

// Stop stops the DLNA/SSDP server.
func (s *Server) Stop() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.running {
		return nil
	}

	s.running = false

	if s.listener != nil {
		s.listener.Close()
	}

	s.logger.Info("DLNA server stopped")
	return nil
}

// IsRunning returns whether the server is running.
func (s *Server) IsRunning() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.running
}

// serve handles incoming DLNA/SSDP requests.
func (s *Server) serve() {
	http.HandleFunc("/upnp/desc/uuid:emby-go/EmbyServer/device.xml", s.handleDeviceDescriptor)
	http.HandleFunc("/upnp/control/ConnectionManager", s.handleConnectionManager)
	http.HandleFunc("/upnp/control/ContentDirectory", s.handleContentDirectory)
	http.HandleFunc("/upnp/event/ConnectionManager", s.handleConnectionManagerEvent)
	http.HandleFunc("/upnp/event/ContentDirectory", s.handleContentDirectoryEvent)

	if err := http.Serve(s.listener, nil); err != nil {
		if !s.running {
			return
		}
		s.logger.Error("DLNA server error", zap.Error(err))
	}
}

// handleDeviceDescriptor handles device descriptor requests.
func (s *Server) handleDeviceDescriptor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/xml")
	w.Header().Set("EXT", "")
	w.Write([]byte(s.DescriptorXML()))
}

// handleConnectionManager handles ConnectionManager SOAP actions.
func (s *Server) handleConnectionManager(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/xml")
	body, _ := io.ReadAll(r.Body)
	w.Write([]byte(xml.ContentDirectoryAction("ConnectionManager", string(body))))
}

// handleContentDirectory handles ContentDirectory SOAP actions.
func (s *Server) handleContentDirectory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/xml")
	body, _ := io.ReadAll(r.Body)
	w.Write([]byte(xml.ContentDirectoryAction("ContentDirectory", string(body))))
}

// handleConnectionManagerEvent handles ConnectionManager event notifications.
func (s *Server) handleConnectionManagerEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("CONTENT-TYPE", "text/xml")
	w.Write([]byte(xml.ConnectionManagerEvent()))
}

// handleContentDirectoryEvent handles ContentDirectory event notifications.
func (s *Server) handleContentDirectoryEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("CONTENT-TYPE", "text/xml")
	w.Write([]byte(xml.ContentDirectoryEvent()))
}

// DescriptorXML returns the DLNA device descriptor XML.
func (s *Server) DescriptorXML() string {
	return `<?xml version="1.0"?>
<root xmlns="urn:schemas-upnp-org:device-1-0">
  <specVersion>
    <major>1</major>
    <minor>0</minor>
  </specVersion>
  <device>
    <deviceType>urn:schemas-upnp-org:device:MediaServer:1</deviceType>
    <friendlyName>Emby Server</friendlyName>
    <manufacturer>Emby</manufacturer>
    <manufacturerURL>https://emby.media</manufacturerURL>
    <modelName>Emby Server</modelName>
    <modelNumber>1.0</modelNumber>
    <modelURL>https://emby.media</modelURL>
    <serialNumber>EMBY</serialNumber>
    <UDN>uuid:emby-go</UDN>
    <iconList>
      <icon>
        <mimetype>image/png</mimetype>
        <width>48</width>
        <height>48</height>
        <depth>32</depth>
        <url>/logo.png</url>
      </icon>
    </iconList>
    <serviceList>
      <service>
        <serviceType>urn:schemas-upnp-org:service:ConnectionManager:1</serviceType>
        <serviceId>ConnectionManager</serviceId>
        <controlURL>/upnp/control/ConnectionManager</controlURL>
        <eventSubURL>/upnp/event/ConnectionManager</eventSubURL>
      </service>
      <service>
        <serviceType>urn:schemas-upnp-org:service:ContentDirectory:1</serviceType>
        <serviceId>ContentDirectory</serviceId>
        <controlURL>/upnp/control/ContentDirectory</controlURL>
        <eventSubURL>/upnp/event/ContentDirectory</eventSubURL>
      </service>
    </serviceList>
  </device>
</root>`
}
