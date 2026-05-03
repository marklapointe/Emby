# Component: emby-go

**Path:** `emby-go/`
**Type:** Directory | Go Application
**Language:** Go
**Maps to:** `.discovery/140-emby-go.md`

## Description

emby-go is a Go-based reimplementation of Emby Server. It provides a modern HTTP API server using Go's standard library and Chi router, with services for library management, media streaming, user authentication, session management, transcoding, and DLNA. It uses SQLite for persistence and Zap for structured logging.

## Structure

```
emby-go/
├── go.mod                         # Go module definition
├── cmd/
│   └── emby-server/
│       └── main.go                # Application entry point
│           └── [func] main()
│               ├── Loads configuration from file
│               ├── Initializes Zap logger
│               ├── Opens SQLite database
│               ├── Creates HTTPServer
│               ├── Registers all API routes
│               ├── Starts HTTP server
│               └── Blocks on OS signals for graceful shutdown
├── internal/
│   ├── api/
│   │   ├── router.go              # Chi HTTP router
│   │   │   └── [type] Router struct
│   │   │       ├── [func] NewRouter(cfg, logger, dbManager) *Router
│   │   │       ├── [func] RegisterAll() — registers all route groups
│   │   │       ├── [func] registerLibraryRoutes()
│   │   │       ├── [func] registerSessionRoutes()
│   │   │       ├── [func] registerUserRoutes()
│   │   │       ├── [func] registerDeviceRoutes()
│   │   │       ├── [func] registerImageRoutes()
│   │   │       ├── [func] registerMediaRoutes()
│   │   │       ├── [func] registerNotificationRoutes()
│   │   │       ├── [func] registerScheduledTaskRoutes()
│   │   │       ├── [func] registerTranscodingRoutes()
│   │   │       └── [func] registerRoute(method, path, handler)
│   │   ├── handlers/              # HTTP request handlers (30+ files)
│   │   │   ├── activity.go          # Activity log handlers
│   │   │   ├── branding.go          # Server branding handlers
│   │   │   ├── channel.go           # Channel handlers
│   │   │   ├── config.go            # Configuration handlers
│   │   │   ├── device.go            # Device management handlers
│   │   │   ├── displayprefs.go      # Display preferences handlers
│   │   │   ├── environment.go       # Environment info handlers
│   │   │   ├── filter.go            # Content filter handlers
│   │   │   ├── games.go             # Game library handlers
│   │   │   ├── image.go             # Image processing handlers
│   │   │   ├── library.go           # Media library handlers
│   │   │   ├── livetv.go            # Live TV handlers
│   │   │   ├── localization.go      # Localization handlers
│   │   │   ├── media.go             # Media item handlers
│   │   │   ├── movies.go            # Movie library handlers
│   │   │   ├── notification.go      # Notification handlers
│   │   │   ├── package.go           # Package/plugin handlers
│   │   │   ├── playback.go          # Playback session handlers
│   │   │   ├── playlist.go          # Playlist handlers
│   │   │   ├── scheduledtask.go     # Scheduled task handlers
│   │   │   ├── search.go            # Search handlers
│   │   │   ├── session.go           # Session management handlers
│   │   │   ├── startup.go           # Startup configuration handlers
│   │   │   ├── system.go            # System info handlers
│   │   │   ├── transcoding.go       # Transcoding handlers
│   │   │   ├── tvshows.go           # TV show library handlers
│   │   │   └── user.go              # User management handlers
│   │   └── middleware/
│   │       ├── auth.go              # Authentication middleware
│   │       │   └── [func] AuthMiddleware(tokenService) — validates JWT tokens
│   │       └── middleware.go        # Common middleware
│   │           └── [func] LoggingMiddleware(logger) — request logging
│   │           └── [func] CORSMiddleware() — CORS headers
│   │           └── [func] RecoveryMiddleware() — panic recovery
│   ├── config/
│   │   ├── config.go                # Configuration management
│   │   │   └── [type] Config struct
│   │   │       ├── ServerConfig (port, bind address, TLS)
│   │   │       ├── DatabaseConfig (SQLite path, max connections)
│   │   │       ├── LibraryConfig (scan intervals, paths)
│   │   │       └── LoggingConfig (level, format, output)
│   │   │   └── [func] DefaultConfig() — returns default configuration
│   │   │   └── [func] LoadConfig(path) — loads from JSON/YAML
│   │   │   └── [func] SaveConfig(path) — persists to file
│   │   └── config_test.go           # Config tests
│   ├── database/
│   │   └── database.go              # SQLite database manager
│   │       └── [type] Manager struct
│   │           ├── [func] NewManager(config) — opens SQLite connection
│   │           ├── [func] Migrate() — runs schema migrations
│   │           ├── [func] Close() — closes connection
│   │           └── [func] Ping() — health check
│   ├── model/
│   │   ├── item.go                  # Media item models
│   │   │   └── [type] Item struct
│   │   │       ├── ID, Name, Type, Path, Metadata
│   │   │       ├── ParentID, IndexNumber, ProductionYear
│   │   │       └── CreatedAt, UpdatedAt
│   │   ├── session.go               # Session models
│   │   │   └── [type] Session struct
│   │   │       ├── ID, UserID, DeviceID, ClientName
│   │   │       ├── LastActivityTime, IsActive
│   │   │       └── PlaybackInfo, Capabilities
│   │   ├── stream.go                # Stream models
│   │   │   └── [type] StreamInfo struct
│   │   │       ├── ItemID, MediaSourceID
│   │   │       ├── VideoCodec, AudioCodec, Container
│   │   │       └── Bitrate, Width, Height
│   │   ├── user.go                  # User models
│   │   │   └── [type] User struct
│   │   │       ├── ID, Name, Email, PasswordHash
│   │   │       ├── IsAdmin, IsHidden, IsDisabled
│   │   │       └── Configuration, Policy
│   │   └── model_test.go            # Model tests
│   ├── repository/
│   │   ├── base.go                  # Base repository
│   │   │   └── [type] Repository<T> interface
│   │   │       ├── Get(id) (T, error)
│   │   │       ├── GetAll() ([]T, error)
│   │   │       ├── Create(item) error
│   │   │       ├── Update(item) error
│   │   │       └── Delete(id) error
│   │   ├── item.go                  # Item repository
│   │   │   └── [type] ItemRepository struct
│   │   │       ├── [func] GetByPath(path) — lookup by file path
│   │   │       ├── [func] GetChildren(parentID) — child items
│   │   │       ├── [func] Search(query) — full-text search
│   │   │       └── [func] GetByType(itemType) — filter by type
│   │   └── item_test.go             # Item repository tests
│   ├── service/
│   │   ├── auth/
│   │   │   └── auth.go              # Authentication service
│   │   │       └── [func] Authenticate(username, password) — validates credentials
│   │   │       └── [func] GenerateToken(user) — creates JWT
│   │   │       └── [func] ValidateToken(token) — verifies JWT signature
│   │   ├── device/
│   │   │   └── device.go            # Device management service
│   │   ├── image/
│   │   │   ├── image.go             # Image service
│   │   │   └── processor.go         # Image processing (resize, format conversion)
│   │   ├── library/
│   │   │   ├── library.go           # Library service
│   │   │   ├── scanner.go           # Library scanner
│   │   │   │   └── [func] ScanLibrary(path) — discovers media files
│   │   │   │   └── [func] RefreshMetadata(item) — fetches metadata
│   │   │   └── notifier.go          # Library change notifier
│   │   ├── media/
│   │   │   ├── media.go             # Media streaming service
│   │   │   └── stream_manager.go    # Stream session manager
│   │   ├── metadata/
│   │   │   ├── metadata.go          # Metadata service
│   │   │   ├── fetcher.go           # Metadata fetcher (OMDb, TMDb, TVDb)
│   │   │   └── limiter.go           # Rate limiter for API calls
│   │   ├── notification/
│   │   │   └── manager.go           # Notification manager
│   │   ├── scheduled/
│   │   │   └── tasks.go             # Scheduled task runner
│   │   ├── session/
│   │   │   ├── session.go           # Session service
│   │   │   ├── websocket.go         # WebSocket session handler
│   │   │   └── session_test.go      # Session tests
│   │   ├── transcoding/
│   │   │   └── transcoding.go       # Transcoding service (FFmpeg wrapper)
│   │   └── user/
│   │       ├── user.go              # User service
│   │       └── user_test.go         # User service tests
│   ├── server/
│   │   ├── http.go                  # HTTP server
│   │   │   └── [type] HTTPServer struct
│   │   │       ├── [func] Start() — binds to configured port
│   │   │       ├── [func] Shutdown(ctx) — graceful shutdown
│   │   │       ├── [func] Router() — returns Chi mux
│   │   │       └── [func] GetConfig(), GetLogger()
│   │   └── ws/
│   │       └── websocket.go         # WebSocket server
│   │           └── [func] HandleWebSocket(w, r) — upgrades HTTP to WS
│   ├── logging/
│   │   └── logging.go               # Zap logger setup
│   │       └── [func] NewLogger(config) — creates structured logger
│   ├── dlna/
│   │   ├── server.go                # DLNA server
│   │   └── xml/
│   │       └── descriptors.go       # DLNA XML descriptors
│   └── plugin/
│       └── manager.go               # Plugin manager
└── tests/
    ├── e2e/
    │   └── e2e_test.go              # End-to-end tests
    ├── integration/
    │   └── integration_test.go      # Integration tests
    └── performance/
        └── benchmark_test.go        # Performance benchmarks
```

## API Route Groups

| Group | Handler File | Routes |
|-------|-------------|--------|
| Library | library.go | GET/POST /libraries, /libraries/{id}/items |
| Session | session.go | GET/POST /sessions, /sessions/{id}/playing |
| User | user.go | GET/POST /users, /users/{id} |
| Device | device.go | GET/POST /devices |
| Image | image.go | GET /images/{id} |
| Media | media.go | GET /media/{id}/stream |
| Notification | notification.go | GET/POST /notifications |
| Scheduled Task | scheduledtask.go | GET/POST /scheduledtasks |
| Transcoding | transcoding.go | GET /transcoding/{id}/progress |
| Movies | movies.go | GET /movies |
| TV Shows | tvshows.go | GET /tvshows |
| Live TV | livetv.go | GET /livetv |
| Playback | playback.go | POST /playback/start, /playback/stop |
| Search | search.go | GET /search |
| System | system.go | GET /system/info |

## Technology Stack

| Component | Technology |
|-----------|------------|
| HTTP Router | Chi (go-chi/chi) |
| Database | SQLite (mattn/go-sqlite3) |
| Logging | Zap (uber-go/zap) |
| Authentication | JWT |
| WebSocket | gorilla/websocket |
| Image Processing | Go image packages |
| Transcoding | FFmpeg (external process) |

## Data Flow

```mermaid
graph TD
    A[Client Request] --&gt; B[Chi Router]
    B --&gt; C[Middleware]
    C --&gt; D{Auth required?}
    D --&gt;|Yes| E[AuthMiddleware]
    E --&gt; F[Validate JWT]
    D --&gt;|No| G[Handler]
    F --&gt; G
    G --&gt; H[Service Layer]
    H --&gt; I[Repository]
    I --&gt; J[SQLite]
    H --&gt; K[External API]
    H --&gt; L[FFmpeg]
    H --&gt; M[Response]
```

## Side Effects

- Reads/writes SQLite database file
- Reads configuration from JSON/YAML
- Writes log files (Zap)
- Spawns FFmpeg processes for transcoding
- Calls external metadata APIs (OMDb, TMDb, TVDb)
- Opens HTTP server on configured port
- Manages WebSocket connections
- DLNA UDP multicast announcements
