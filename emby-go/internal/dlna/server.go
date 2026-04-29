package dlna

import (
	"fmt"
	"net"

	"go.uber.org/zap"
)

// Server represents the DLNA/SSDP server.
type Server struct {
	addr     string
	port     int
	logger   *zap.Logger
	listener net.PacketConn
	alive    bool
}

// NewServer creates a new DLNA server.
func NewServer(port int, logger *zap.Logger) *Server {
	return &Server{
		addr:   "239.255.255.250",
		port:   port,
		logger: logger,
	}
}

// Start begins listening for SSDP messages.
func (s *Server) Start() error {
	listener, err := net.ListenMulticastUDP("udp4", nil, &net.UDPAddr{
		IP:   net.ParseIP(s.addr),
		Port: 1900,
	})
	if err != nil {
		return fmt.Errorf("listen multicast UDP: %w", err)
	}

	s.listener = listener
	s.alive = true

	s.logger.Info("DLNA server started",
		zap.String("address", s.addr),
		zap.Int("port", s.port),
	)

	go s.listenLoop()
	return nil
}

// Stop shuts down the DLNA server.
func (s *Server) Stop() {
	s.alive = false
	if s.listener != nil {
		s.listener.Close()
	}
}

// listenLoop handles incoming SSDP messages.
func (s *Server) listenLoop() {
	buf := make([]byte, 65536)
	for s.alive {
		n, remote, err := s.listener.ReadFrom(buf)
		if err != nil {
			if s.alive {
				s.logger.Error("read error", zap.Error(err))
			}
			continue
		}

		msg := string(buf[:n])
		s.logger.Debug("received SSDP message",
			zap.String("from", remote.String()),
			zap.String("message", msg[:min(n, 100)]),
		)

		// Parse and respond to SSDP messages
		s.handleSSDPMessage(msg, remote)
	}
}

// handleSSDPMessage processes incoming SSDP messages.
func (s *Server) handleSSDPMessage(msg string, remote net.Addr) {
	// Check for M-SEARCH requests
	if !contains(msg, "M-SEARCH") {
		return
	}

	// Parse ST (Search Target) header
	st := extractHeader(msg, "ST:")
	if st == "" {
		return
	}

	// Respond to relevant search targets
	switch {
	case contains(st, "upnp:rootdevice"), contains(st, "ssdp:all"):
		s.sendRootDeviceResponse(remote)
	case contains(st, "uuid:"), contains(st, "upnp:device"):
		s.sendDeviceResponse(remote)
	}
}

// sendRootDeviceResponse sends a root device response.
func (s *Server) sendRootDeviceResponse(remote net.Addr) {
	response := fmt.Sprintf(
		"HTTP/1.1 200 OK\r\n"+
			"CACHE-CONTROL: max-age=1800\r\n"+
			"EXT:\r\n"+
			"LOCATION: http://127.0.0.1:%d/descriptor.xml\r\n"+
			"SERVER: Emby Go Server/1.0 UPnP/1.1\r\n"+
			"ST: upnp:rootdevice\r\n"+
			"USN: uuid:emby-go::upnp:rootdevice\r\n"+
			"\r\n",
		s.port,
	)
	s.sendResponse(response, remote)
}

// sendDeviceResponse sends a device response.
func (s *Server) sendDeviceResponse(remote net.Addr) {
	response := fmt.Sprintf(
		"HTTP/1.1 200 OK\r\n"+
			"CACHE-CONTROL: max-age=1800\r\n"+
			"EXT:\r\n"+
			"LOCATION: http://127.0.0.1:%d/descriptor.xml\r\n"+
			"SERVER: Emby Go Server/1.0 UPnP/1.1\r\n"+
			"ST: uuid:emby-go\r\n"+
			"USN: uuid:emby-go\r\n"+
			"\r\n",
		s.port,
	)
	s.sendResponse(response, remote)
}

// sendResponse sends an SSDP response.
func (s *Server) sendResponse(response string, remote net.Addr) {
	conn, err := net.DialUDP("udp4", nil, remote.(*net.UDPAddr))
	if err != nil {
		s.logger.Error("dial UDP", zap.Error(err))
		return
	}
	defer conn.Close()

	_, err = conn.Write([]byte(response))
	if err != nil {
		s.logger.Error("write response", zap.Error(err))
	}
}

// sendAlive broadcasts that the server is alive.
func (s *Server) sendAlive() {
	aliveMsg := fmt.Sprintf(
		"NOTIFY * HTTP/1.1\r\n"+
			"HOST: 239.255.255.250:1900\r\n"+
			"NT: uuid:emby-go\r\n"+
			"NTS: ssdp:alive\r\n"+
			"LOCATION: http://127.0.0.1:%d/descriptor.xml\r\n"+
			"SERVER: Emby Go Server/1.0 UPnP/1.1\r\n"+
			"Cache-Control: max-age=1800\r\n"+
			"\r\n",
		s.port,
	)

	conn, err := net.DialUDP("udp4", nil, &net.UDPAddr{
		IP:   net.ParseIP("239.255.255.250"),
		Port: 1900,
	})
	if err != nil {
		return
	}
	defer conn.Close()

	conn.Write([]byte(aliveMsg))
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

// Helper functions
func contains(s, substr string) bool {
	return len(s) > 0 && len(substr) > 0 && (len(s) >= len(substr)) && (s == substr || len(s) > len(substr) && (s[:len(substr)] == substr || s[len(s)-len(substr):] == substr || findSubstring(s, substr)))
}

func findSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

func extractHeader(msg, header string) string {
	for _, line := range splitLines(msg) {
		if len(line) >= len(header) && line[:len(header)] == header {
			return line[len(header):]
		}
	}
	return ""
}

func splitLines(s string) []string {
	var lines []string
	start := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			lines = append(lines, s[start:i])
			start = i + 1
		}
	}
	if start < len(s) {
		lines = append(lines, s[start:])
	}
	return lines
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
