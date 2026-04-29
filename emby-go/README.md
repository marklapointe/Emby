# Emby Server (Go)

A native Go implementation of the Emby Server media server, migrated from the original C#/.NET codebase.

## Overview

Emby Server is a personal media server with both server and client components. This Go implementation provides:

- **Native performance** - No Mono/.NET runtime required
- **Lower memory footprint** - Go's efficient memory management
- **Simplified deployment** - Single binary, no external dependencies
- **Cross-platform** - Linux, macOS, Windows support
- **Full API compatibility** - Works with existing Emby clients

## Project Structure

```
emby-go/
├── cmd/emby-server/       # Application entry point
├── internal/
│   ├── api/               # HTTP API layer
│   │   ├── handlers/      # Request handlers
│   │   ├── middleware/    # HTTP middleware
│   │   └── router.go      # API routing
│   ├── config/            # Configuration management
│   ├── database/          # Database connection management
│   ├── dlna/              # DLNA/SSDP support
│   ├── logging/           # Structured logging
│   ├── model/             # Data models
│   ├── plugin/            # Plugin management
│   ├── repository/        # Data access layer
│   ├── server/            # HTTP server
│   └── service/           # Business logic
│       ├── device/        # Device management
│       ├── image/         # Image processing
│       ├── library/       # Library scanning
│       ├── media/         # Media handling
│       ├── metadata/      # Metadata providers
│       ├── notification/  # Notifications
│       ├── scheduled/     # Scheduled tasks
│       ├── session/       # Session management
│       ├── transcoding/   # FFmpeg integration
│       └── user/          # User management
├── configs/               # Configuration files
└── Makefile               # Build commands
```

## Quick Start

### Prerequisites

- Go 1.21+
- FFmpeg (for transcoding)
- SQLite (bundled via Go module)

### Building

```bash
# Clone the repository
git clone https://github.com/emby/emby-server.git
cd emby-server/emby-go

# Build the server
make build

# Run tests
make test

# Run the server
make run
```

### Configuration

The server can be configured via:

1. **YAML config file** (`configs/default.yaml` or `config.yaml`)
2. **Environment variables** (overrides config file values)

#### Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `EMBY_SERVER_HOST` | Server bind address | `0.0.0.0` |
| `EMBY_SERVER_PORT` | Server port | `8096` |
| `EMBY_DATABASE_PATH` | SQLite database path | `data/emby-server.db` |
| `EMBY_LOG_LEVEL` | Log level (debug, info, warn, error) | `info` |

#### Config File Example

```yaml
server:
  host: "0.0.0.0"
  port: 8096
  max_header_bytes: 1048576
  read_timeout: 30
  write_timeout: 30

database:
  path: "data/emby-server.db"
  max_open_conns: 25
  max_idle_conns: 10
  conn_max_lifetime: 300
  enable_wal: true
  pragma_journal_mode: "WAL"

library:
  scan_interval_minutes: 24
  enable_auto_deep_scan: true
  content_types:
    - "Video"
    - "Music"
    - "Photos"
    - "Books"
  ignore_paths:
    - ".cache"
    - "tmp"

logging:
  level: "info"
  format: "json"

stream_pooling:
  enabled: true
  max_concurrent_streams: 50
  idle_timeout: "5m"
  health_check_interval: "30s"
  metrics_enabled: true
```

## API Endpoints

### Library Endpoints

| Method | Path | Description |
|--------|------|-------------|
| GET | `/Library/Root` | Get library root |
| GET | `/Library/Items` | Get library items |
| GET | `/Library/MediaFolders` | Get media folders |
| POST | `/Library/MediaFolders` | Create media folder |
| GET | `/Library/MediaFolders/{id}` | Get media folder |
| DELETE | `/Library/MediaFolders/{id}` | Delete media folder |
| GET | `/Library/MediaFolders/{id}/Items` | Get folder items |
| GET | `/Items/{id}` | Get item details |
| GET | `/Items/{id}/Images/{type}` | Get item image |
| GET | `/Items/{id}/Stream` | Stream item |
| GET | `/Items/{id}/Subtitles` | Get subtitles |
| GET | `/Items/{id}/Subtitles/{index}/Stream` | Get subtitle stream |

### Session Endpoints

| Method | Path | Description |
|--------|------|-------------|
| GET | `/Sessions` | Get all sessions |
| GET | `/Sessions/{id}` | Get session details |
| POST | `/Sessions/{id}/Playing` | Start playback |
| POST | `/Sessions/{id}/Playing/Progress` | Update playback progress |
| POST | `/Sessions/{id}/Playing/Stopped` | Stop playback |
| POST | `/Sessions/{id}/Volume` | Set volume |
| POST | `/Sessions/{id}/Pause` | Pause playback |
| POST | `/Sessions/{id}/Unpause` | Resume playback |
| POST | `/Sessions/{id}/ToggleFullscreen` | Toggle fullscreen |
| POST | `/Sessions/{id}/GoTo` | Navigate to item |
| POST | `/Sessions/{id}/SendKey` | Send key press |
| POST | `/Sessions/{id}/SendText` | Send text input |
| DELETE | `/Sessions/{id}` | Close session |

### User Endpoints

| Method | Path | Description |
|--------|------|-------------|
| GET | `/Users` | Get all users |
| GET | `/Users/Public` | Get public users |
| POST | `/Users/Login` | Login user |
| POST | `/Users/Logout` | Logout user |
| GET | `/Users/{id}` | Get user details |
| PUT | `/Users/{id}` | Update user |
| DELETE | `/Users/{id}` | Delete user |
| POST | `/Users/{id}/Password` | Change password |
| GET | `/Users/{id}/Images/{type}` | Get user image |
| GET | `/Users/{id}/Configuration` | Get user config |
| PUT | `/Users/{id}/Configuration` | Update user config |
| GET | `/Users/{id}/Policy` | Get user policy |
| PUT | `/Users/{id}/Policy` | Update user policy |

### Device Endpoints

| Method | Path | Description |
|--------|------|-------------|
| GET | `/Devices` | Get all devices |
| GET | `/Devices/{id}` | Get device details |
| PUT | `/Devices/{id}` | Update device |
| DELETE | `/Devices/{id}` | Delete device |
| GET | `/Devices/{id}/Icon` | Get device icon |

### DLNA Endpoints

| Method | Path | Description |
|--------|------|-------------|
| GET | `/upnp/desc/uuid:emby-go/EmbyServer/device.xml` | Device descriptor |
| POST | `/upnp/control/ConnectionManager` | ConnectionManager actions |
| POST | `/upnp/control/ContentDirectory` | ContentDirectory actions |
| POST | `/upnp/event/ConnectionManager` | ConnectionManager events |
| POST | `/upnp/event/ContentDirectory` | ContentDirectory events |

## Development

### Running Tests

```bash
# Run all tests
make test

# Run specific package tests
go test ./internal/service/user/...

# Run with coverage
go test -cover ./...

# Run with race detector
go test -race ./...
```

### Code Style

- Follow [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- Use `gofmt` for formatting
- Use `golint` for linting
- Write tests for all new code

### Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests for new functionality
5. Ensure all tests pass
6. Submit a pull request

## Architecture

### Data Flow

```
Client Request → HTTP Server → Router → Middleware → Handler → Service → Repository → Database
```

### Key Components

1. **HTTP Server** - Handles incoming HTTP requests with chi router
2. **Service Layer** - Business logic (user, library, media, session management)
3. **Repository Layer** - Data access with SQLite
4. **DLNA Server** - UPnP/DLNA device discovery and control
5. **Transcoding Service** - FFmpeg integration for media transcoding
6. **Plugin System** - Extensible plugin architecture

### Stream Pooling

The Go implementation includes stream pooling for efficient resource usage:

- **Live TV**: One FFmpeg process per channel, shared by all viewers
- **Recorded Media**: Optional sharing based on user preference
- **Automatic Cleanup**: Streams closed when last viewer disconnects

## Deployment

### Docker

```bash
docker run -d \
  --name emby-server \
  -p 8096:8096 \
  -v /path/to/config:/config \
  -v /path/to/media:/media \
  emby/emby-server:latest
```

### Systemd

```ini
[Unit]
Description=Emby Server
After=network.target

[Service]
Type=simple
User=emby
ExecStart=/usr/local/bin/emby-server
Restart=on-failure
RestartSec=10

[Install]
WantedBy=multi-user.target
```

### nginx Reverse Proxy

```nginx
server {
    listen 443 ssl;
    server_name emby.example.com;

    ssl_certificate /etc/letsencrypt/live/emby.example.com/fullchain.pem;
    ssl_certificate_key /etc/letsencrypt/live/emby.example.com/privkey.pem;

    location / {
        proxy_pass http://127.0.0.1:8096;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
        
        # WebSocket support
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection "upgrade";
    }
}
```

## License

This project is licensed under the same license as the original Emby Server. See [LICENSE.md](../../LICENSE.md) for details.

## Acknowledgments

- Original Emby Server C# implementation
- [chi](https://github.com/go-chi/chi) - Lightweight HTTP router
- [modernc.org/sqlite](https://github.com/mattn/go-sqlite3) - SQLite driver
- [gorilla/websocket](https://github.com/gorilla/websocket) - WebSocket support
- [zap](https://github.com/uber-go/zap) - Structured logging
