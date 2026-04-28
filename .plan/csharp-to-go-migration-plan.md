# Emby Server: C#/.NET to Go Migration Plan

## 1. Executive Summary

This document outlines a comprehensive, phased approach to migrating the Emby Server application from C#/.NET (Mono) to Go (Golang). The goal is to create a native Go implementation that maintains compatibility with existing Emby clients while improving performance, reducing resource consumption, and simplifying deployment.

**Primary Recommendation:** Incremental, component-by-component migration with parallel operation during transition period.

**Key Finding:** The current implementation uses **SQLite** for data storage, NOT MongoDB. This simplifies the migration as SQLite has excellent Go support via `modernc.org/sqlite` (pure Go) or `mattn/go-sqlite3` (CGO).

---

## 2. Current Architecture Analysis

### 2.1 Technology Stack

| Component | Current Technology | Role |
|-----------|-------------------|------|
| Runtime | Mono/.NET Framework | Application runtime |
| Language | C# | Primary implementation language |
| Database | SQLite | Media library, user data, settings storage |
| HTTP Server | Custom (SocketHttpListener) | REST API server |
| Media Processing | FFmpeg | Transcoding, metadata extraction |
| Image Processing | ImageMagick/Skia | Thumbnail generation, image manipulation |
| Protocols | DLNA, HTTP, WebSocket | Client communication |

### 2.2 Project Structure

```
Emby/
├── MediaBrowser.ServerApplication/      # Main application entry point
├── MediaBrowser.Server.Mono/            # Mono-specific launcher
├── Emby.Server.Implementations/         # Core server logic
│   ├── Data/                           # SQLite repositories
│   ├── Library/                        # Media library management
│   ├── HttpServer/                     # HTTP server implementation
│   ├── Session/                        # Session management
│   ├── Security/                       # Authentication/authorization
│   ├── ScheduledTasks/                 # Background tasks
│   ├── LiveTv/                         # Live TV functionality
│   ├── MediaEncoder/                   # FFmpeg integration
│   └── ...
├── MediaBrowser.Api/                    # REST API endpoints
├── MediaBrowser.Providers/              # Metadata providers
├── MediaBrowser.LocalMetadata/          # Local metadata parsing
├── Emby.Dlna/                           # DLNA support (UPnP removed for security)
├── Emby.Drawing/                        # Image abstraction
├── Emby.Drawing.ImageMagick/            # ImageMagick implementation
├── Emby.Drawing.Skia/                   # Skia implementation
├── MediaBrowser.WebDashboard/           # Web UI
└── ThirdParty/                          # Third-party libraries
```

### 2.3 Key Components Identified

#### 2.3.1 Data Layer
- **BaseSqliteRepository.cs** - Base repository for SQLite operations
- **SqliteItemRepository.cs** - Main media item storage and retrieval
- **AuthenticationRepository.cs** - User authentication data
- **Finding:** All data persistence uses SQLite, NOT MongoDB

#### 2.3.2 HTTP Server
- **HttpListenerHost.cs** - HTTP request handling
- **SocketHttpListener/** - Custom HTTP listener implementation
- **WebSocketConnection.cs** - WebSocket support
- **AuthorizationContext.cs** - Request authentication
- **SessionContext.cs** - Session management

#### 2.3.3 API Layer
- **BaseApiService.cs** - Base API service class
- **ApiEntryPoint.cs** - API initialization
- Multiple service classes (LibraryService, SessionsService, etc.)

#### 2.3.4 Media Processing
- **LibraryManager.cs** - Media library scanning and management
- **FFMpegLoader.cs** - FFmpeg integration
- **EncodingManager.cs** - Transcoding management
- **MediaSourceManager.cs** - Media source handling

#### 2.3.5 Networking
- **SocketFactory.cs** - Socket creation
- **UdpSocket.cs** - UDP communication (DLNA, discovery)
- **IWebSocket.cs** - WebSocket interface

### 2.4 Architecture Patterns

1. **Repository Pattern** - Data access abstracted through repositories
2. **Service Layer** - Business logic in service classes
3. **Dependency Injection** - Component wiring via ApplicationHost
4. **Event-driven** - Library changes, sessions via events
5. **Plugin Architecture** - Extensible via plugins

---

## 3. Proposed Go Architecture

### 3.1 High-Level Design

```
┌─────────────────────────────────────────────────────────────┐
│                      Emby Go Server                          │
├─────────────────────────────────────────────────────────────┤
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────────────┐  │
│  │   HTTP/1.1  │  │  WebSocket  │  │   DLNA (local)      │  │
│  │   Server    │  │   Server    │  │   Server            │  │
│  └──────┬──────┘  └──────┬──────┘  └──────────┬──────────┘  │
│         │                │                     │             │
│         └────────────────┼─────────────────────┘             │
│                          │                                   │
│  ┌───────────────────────▼───────────────────────────────┐  │
│  │              API Router (chi/gin)                      │  │
│  └───────────────────────┬───────────────────────────────┘  │
│                          │                                   │
│  ┌───────────────────────▼───────────────────────────────┐  │
│  │              Service Layer                             │  │
│  │  ┌──────────┐ ┌──────────┐ ┌──────────┐ ┌──────────┐  │  │
│  │  │ Library  │ │ Session  │ │  User    │ │  Media   │  │  │
│  │  │ Service  │ │ Service  │ │ Service  │ │ Service  │  │  │
│  │  └──────────┘ └──────────┘ └──────────┘ └──────────┘  │  │
│  └───────────────────────┬───────────────────────────────┘  │
│                          │                                   │
│  ┌───────────────────────▼───────────────────────────────┐  │
│  │              Data Access Layer                         │  │
│  │  ┌──────────────────────────────────────────────────┐ │  │
│  │  │           SQLite (modernc.org/sqlite)            │ │  │
│  │  └──────────────────────────────────────────────────┘ │  │
│  └───────────────────────────────────────────────────────┘  │
│                          │                                   │
│  ┌───────────────────────▼───────────────────────────────┐  │
│  │              External Integrations                     │  │
│  │  ┌──────────┐ ┌──────────┐ ┌────────────────────────┐ │  │
│  │  │  FFmpeg  │ │  Providers│ │  Image Processing     │ │  │
│  │  │ (exec)   │ │ (HTTP)   │ │  (govips/magick)      │ │  │
│  │  └──────────┘ └──────────┘ └────────────────────────┘ │  │
│  └───────────────────────────────────────────────────────┘  │
└─────────────────────────────────────────────────────────────┘
```

### 3.2 Go Module Structure

```
emby-go/
├── cmd/
│   └── emby-server/
│       └── main.go              # Application entry point
├── internal/
│   ├── server/
│   │   ├── http.go              # HTTP server setup
│   │   ├── websocket.go         # WebSocket handling
│   │   └── dlna.go              # DLNA server (UPnP removed)
│   ├── api/
│   │   ├── router.go            # API routing
│   │   ├── middleware/
│   │   │   ├── auth.go          # Authentication middleware
│   │   │   └── session.go       # Session middleware
│   │   └── handlers/
│   │       ├── library.go       # Library API endpoints
│   │       ├── sessions.go      # Sessions API endpoints
│   │       ├── users.go         # Users API endpoints
│   │       └── ...
│   ├── service/
│   │   ├── library/
│   │   │   ├── manager.go       # Library management
│   │   │   ├── scanner.go       # Media scanning
│   │   │   └── notifier.go      # Change notifications
│   │   ├── session/
│   │   │   ├── manager.go       # Session management
│   │   │   └── websocket.go     # WebSocket sessions
│   │   ├── user/
│   │   │   ├── manager.go       # User management
│   │   │   └── auth.go          # Authentication
│   │   ├── media/
│   │   │   ├── encoder.go       # FFmpeg transcoding
│   │   │   ├── source.go        # Media source handling
│   │   │   └── metadata.go      # Metadata management
│   │   └── scheduled/
│   │       └── tasks.go         # Background tasks
│   ├── repository/
│   │   ├── base.go              # Base repository
│   │   ├── item.go              # Item repository
│   │   ├── user.go              # User repository
│   │   └── auth.go              # Auth repository
│   ├── model/
│   │   ├── item.go              # Media item models
│   │   ├── user.go              # User models
│   │   ├── session.go           # Session models
│   │   └── ...
│   ├── provider/
│   │   ├── metadata/
│   │   │   ├── provider.go      # Metadata provider interface
│   │   │   ├── local.go         # Local metadata
│   │   │   └── remote.go        # Remote metadata providers
│   │   └── images/
│   │       ├── provider.go      # Image provider interface
│   │       └── processor.go     # Image processing
│   ├── dlna/
│   │   ├── server.go            # DLNA server (local network only)
│   │   └── xml/
│   │       └── descriptors.go   # DLNA XML descriptors
│   ├── licensing/
│   │   ├── manager.go           # Emby Premiere license management
│   │   ├── features.go          # Premium feature gating
│   │   └── cache.go             # License status caching
│   └── util/
│       ├── fs/                  # File system utilities
│       ├── hash/                # Hashing utilities
│       └── mime/                # MIME type detection
├── pkg/
│   └── emby/                    # Public API (if needed)
├── web/
│   └── dashboard/               # Web UI (existing, may stay as-is)
├── migrations/
│   └── sqlite/                  # Database migrations
├── configs/
│   └── default.yaml             # Default configuration
├── go.mod
├── go.sum
└── Makefile
```

### 3.3 Key Technology Choices

| Component | Go Technology | Rationale |
|-----------|--------------|-----------|
| HTTP Server | `net/http` + `chi` | Standard library + lightweight router |
| WebSocket | `github.com/gorilla/websocket` | Mature, well-tested |
| SQLite | `modernc.org/sqlite` | Pure Go, no CGO, good performance |
| JSON | `encoding/json` + `jsoniter` | Standard + optional speed boost |
| Logging | `go.uber.org/zap` | High performance, structured |
| Configuration | `gopkg.in/yaml.v3` | YAML support |
| Image Processing | `github.com/davidbyttow/govips` | libvips bindings, fast |
| FFmpeg | `os/exec` | Direct FFmpeg binary execution |
| DLNA | Custom (SSDP local only) | UPnP excluded for security |
| Licensing | Custom + HTTP client | mb3admin.com validation |
| Testing | `testing` + `testify` | Standard + assertions |

---

## 4. Migration Phases

### Phase 1: Foundation and Infrastructure

**Duration:** 4-6 weeks

**Goal:** Establish Go project structure, build system, and basic infrastructure

#### 4.1.1 Project Setup

**Note:** Docker configuration will be deferred to Phase 12 (Documentation and Deployment).

**Tasks:**
1. Initialize Go module with proper structure
2. Create Makefile for build, test, run
3. Set up CI/CD pipeline (GitHub Actions)
4. Configure logging, configuration management

**Deliverables:**
- Working Go project skeleton
- Build produces binary
- Basic configuration loading
- Logging infrastructure

#### 4.1.2 Database Layer

**Note**: Database migration is NOT required. The existing SQLite database files from the C# implementation will be used directly.

**Tasks:**
1. Analyze existing SQLite schema from C# code (for reference only)
2. Implement base repository pattern
3. Create connection pool management
4. Implement transaction support

**Files:**
```go
// internal/repository/base.go
type BaseRepository struct {
    db *sql.DB
}

func (r *BaseRepository) Query(query string, args ...interface{}) (*sql.Rows, error)
func (r *BaseRepository) Exec(query string, args ...interface{}) (sql.Result, error)
func (r *BaseRepository) WithTransaction(fn func(*sql.Tx) error) error
```

**Deliverables:**
- Base repository implementation
- Connection pooling configured
- Existing SQLite databases work without migration

#### 4.1.3 Configuration System

**Tasks:**
1. Define configuration structure (YAML)
2. Implement configuration loading
3. Support environment variable overrides
4. Create default configuration file

**Files:**
```go
// internal/config/config.go
type Config struct {
    Server   ServerConfig   `yaml:"server"`
    Database DatabaseConfig `yaml:"database"`
    Library  LibraryConfig  `yaml:"library"`
    // ...
}
```

**Deliverables:**
- Configuration system
- Default config file
- Documentation

### Phase 2: Core HTTP Server and API Framework

**Duration:** 6-8 weeks

**Goal:** Implement HTTP server, routing, and API framework

#### 4.2.1 HTTP Server

**Note:** TLS/SSL will be handled by nginx and certbot in production. The Go application will run behind nginx as a reverse proxy.

**Tasks:**
1. Set up `net/http` server with `chi` router
2. Implement middleware chain (logging, recovery, CORS)
3. Implement request/response logging
4. Add graceful shutdown

**Files:**
```go
// internal/server/http.go
type HTTPServer struct {
    config *config.Config
    router *chi.Mux
    server *http.Server
}

func (s *HTTPServer) Start() error
func (s *HTTPServer) Shutdown(ctx context.Context) error
```

**Deliverables:**
- HTTP server running
- Middleware chain
- Graceful shutdown
- nginx reverse proxy configuration (no TLS in Go)

#### 4.2.2 API Router

**Tasks:**
1. Map existing C# API routes to Go handlers
2. Create route registration system
3. Implement request binding/validation
4. Create response helpers

**Files:**
```go
// internal/api/router.go
func NewRouter(handlers ...HandlerRegistrar) *chi.Mux {
    r := chi.NewRouter()
    
    // Middleware
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)
    r.Use(middleware.CORS)
    
    // Register handlers
    for _, register := range handlers {
        register(r)
    }
    
    return r
}
```

**Deliverables:**
- API routing framework
- Request/response helpers
- Validation framework

#### 4.2.3 Authentication Middleware

**Tasks:**
1. Analyze C# authentication flow (AuthorizationContext.cs)
2. Implement API key authentication
3. Implement session-based authentication
4. Create permission checking

**Files:**
```go
// internal/api/middleware/auth.go
func AuthMiddleware(authService *service.AuthService) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            // Extract token from header
            // Validate token
            // Set user context
            next.ServeHTTP(w, r)
        })
    }
}
```

**Deliverables:**
- Authentication middleware
- API key support
- Session token support

### Phase 3: Data Layer Migration

**Duration:** 6-8 weeks

**Goal:** Migrate all data repositories from C# to Go

#### 4.3.1 Item Repository

**Tasks:**
1. Analyze SqliteItemRepository.cs thoroughly
2. Create Go models for media items
3. Implement CRUD operations
4. Implement search/query methods
5. Migrate all SQL queries

**Key Queries to Migrate:**
- Item retrieval by ID
- Item listing with filters
- Item search
- Item metadata updates
- Item ancestry/path queries

**Files:**
```go
// internal/repository/item.go
type ItemRepository struct {
    *BaseRepository
}

func (r *ItemRepository) GetByID(id string) (*model.Item, error)
func (r *ItemRepository) GetList(query ItemQuery) ([]*model.Item, error)
func (r *ItemRepository) Insert(item *model.Item) error
func (r *ItemRepository) Update(item *model.Item) error
func (r *ItemRepository) Delete(id string) error
```

**Deliverables:**
- Item repository complete
- All queries migrated
- Tests for repository

#### 4.3.2 User Repository

**Tasks:**
1. Analyze user data structures
2. Create user models
3. Implement user CRUD
4. Implement password hashing (match C# algorithm)
5. Migrate user preferences

**Deliverables:**
- User repository complete
- Password compatibility with C# version

#### 4.3.3 Authentication Repository

**Tasks:**
1. Migrate authentication data storage
2. Implement token storage
3. Implement session storage
4. Ensure compatibility with existing clients

**Deliverables:**
- Auth repository complete
- Existing sessions remain valid

### Phase 4: Service Layer Implementation

**Duration:** 8-10 weeks

**Goal:** Implement all business logic services

#### 4.4.1 Library Service

**Tasks:**
1. Migrate LibraryManager.cs logic
2. Implement media scanning
3. Implement library monitoring (file system watching)
4. Implement metadata refresh
5. Implement library change notifications

**Files:**
```go
// internal/service/library/manager.go
type LibraryManager struct {
    itemRepo    *repository.ItemRepository
    providerMgr *provider.Manager
    fsWatcher   *fsnotify.Watcher
}

func (m *LibraryManager) ScanLibrary(path string) error
func (m *LibraryManager) GetItem(id string) (*model.Item, error)
func (m *LibraryManager) Search(query string) ([]*model.Item, error)
```

**Deliverables:**
- Library scanning works
- File monitoring works
- Metadata refresh works

#### 4.4.2 Session Service

**Tasks:**
1. Migrate SessionManager.cs logic
2. Implement session creation/deletion
3. Implement session activity tracking
4. Implement WebSocket session notifications
5. Implement playback reporting

**Deliverables:**
- Session management works
- WebSocket notifications work
- Playback reporting works

#### 4.4.3 User Service

**Tasks:**
1. Migrate UserManager.cs logic
2. Implement user CRUD
3. Implement authentication
4. Implement user preferences
5. Implement parental controls

**Deliverables:**
- User management works
- Authentication works
- Preferences work

#### 4.4.4 Media Encoding Service

**Tasks:**
1. Migrate FFMpegLoader.cs and EncodingManager.cs
2. Implement FFmpeg process management
3. Implement transcoding profiles
4. Implement bitrate control
5. Implement subtitle burning

**Files:**
```go
// internal/service/media/encoder.go
type Encoder struct {
    ffmpegPath string
    config     *EncodingConfig
}

func (e *Encoder) Transcode(req *TranscodeRequest) (io.ReadCloser, error)
func (e *Encoder) GetTranscodeOptions(source *model.MediaSource) *TranscodeOptions
```

**Deliverables:**
- FFmpeg integration works
- Transcoding profiles work
- Live transcoding works

#### 4.4.5 Stream Pooling and Resource Sharing

**Goal:** Implement intelligent stream management to reduce resource consumption by allowing multiple users to share the same content stream.

**Design Principles:**
1. **Live TV Streams:** Always shared automatically (no user choice needed) - all viewers watch the same live feed
2. **Recorded Media (MKV, MPG, etc.):** Present user with option before playback:
   - "Share Content" - Join existing stream if available (starts from current position)
   - "Start From Beginning" - Create new independent stream
3. **Managed Threading:** Each unique content stream runs in a managed goroutine with proper lifecycle management

**Architecture:**
```go
// internal/service/media/stream_manager.go
type StreamManager struct {
    mu           sync.RWMutex
    activeStreams map[string]*ActiveStream  // key: content ID + transcode profile
    streamPool   *sync.Pool                // Reusable stream buffers
}

type ActiveStream struct {
    ContentID      string
    TranscodeProfile string
    Position       time.Duration
    Viewers        map[string]*Viewer  // key: session ID
    Source         io.ReadCloser
    mu             sync.RWMutex
    createdAt      time.Time
    lastAccessed   time.Time
}

type Viewer struct {
    SessionID   string
    UserID      string
    Position    time.Duration
    ConnectedAt time.Time
}
```

**Key Features:**
1. **Stream Identification:**
   - Unique stream key = ContentID + TranscodeProfile (bitrate, codec, resolution)
   - Live TV: ContentID includes channel ID + timestamp window
   
2. **Stream Sharing Logic:**
   ```go
   func (m *StreamManager) GetOrCreateStream(contentID, profile, sessionID string) (*ActiveStream, error) {
       m.mu.Lock()
       defer m.mu.Unlock()
       
       key := fmt.Sprintf("%s:%s", contentID, profile)
       
       // Check for existing stream
       if stream, exists := m.activeStreams[key]; exists {
           // For live TV: always share
           // For recorded media: only share if user chose "Share Content"
           stream.AddViewer(sessionID)
           return stream, nil
       }
       
       // Create new stream
       stream := NewActiveStream(contentID, profile)
       m.activeStreams[key] = stream
       return stream, nil
   }
   ```

3. **User Prompt for Recorded Media:**
   ```go
   // internal/api/handlers/playback.go
   type PlaybackStartRequest struct {
       ContentID       string `json:"contentId"`
       ShareContent    bool   `json:"shareContent"`    // User's choice
       StartFromBeginning bool `json:"startFromBeginning"`
   }
   
   func (h *PlaybackHandler) StartPlayback(w http.ResponseWriter, r *http.Request) {
       var req PlaybackStartRequest
       // ... parse request ...
       
       if !req.ShareContent {
           // Force new stream
           stream = streamManager.CreateNewStream(...)
       } else {
           // Try to join existing stream
           stream = streamManager.GetOrCreateStream(...)
       }
   }
   ```

4. **Stream Lifecycle Management:**
   - Automatic cleanup when last viewer disconnects
   - Periodic health checks on active streams
   - Resource limits (max concurrent unique streams)
   - Graceful stream termination

5. **Position Tracking:**
   - Each viewer maintains independent position
   - Shared stream continues from original source
   - Late joiners can seek to their preferred position (buffer permitting)

**Files:**
```go
// internal/service/media/stream_manager.go    - Stream pooling logic
// internal/service/media/active_stream.go    - Active stream representation
// internal/api/handlers/playback.go          - Playback start with sharing option
// internal/model/playback_request.go         - Playback request with share flag
```

**API Changes:**
```json
// POST /Sessions/Playing
{
  "ItemId": "abc123",
  "ShareContent": true,  // NEW: User opts to share existing stream
  "MediaSourceId": "xyz789",
  "PlaySessionId": "session-uuid"
}
```

**Deliverables:**
- Stream manager with pooling logic
- User prompt in playback UI (web/mobile/TV clients)
- Live TV streams always shared automatically
- Recorded media respects user choice
- Resource monitoring and limits
- Automatic cleanup of abandoned streams
- Reduced CPU/memory usage during concurrent playback

**Performance Benefits:**
- **Live TV:** 1 FFmpeg process per channel instead of per viewer
- **Popular Content:** Single transcode stream serves multiple viewers
- **Resource Savings:** ~60-80% reduction for concurrent viewers of same content
- **Memory:** Shared buffers reduce overall memory footprint

**Edge Cases to Handle:**
1. Viewer joins mid-stream → buffer management for catch-up
2. Network hiccups → reconnection to same stream
3. Stream source fails → notify all viewers, offer restart
4. Maximum viewers per stream → limit to prevent overload
5. Different quality requests → separate streams per transcode profile

### Phase 5: API Endpoints Implementation

**Duration:** 8-10 weeks

**Goal:** Implement all REST API endpoints

#### 4.5.1 Library API

**Endpoints:**
- `GET /Items` - List items
- `GET /Items/{id}` - Get item by ID
- `GET /Items/{id}/File` - Get item file
- `POST /Items/{id}/Playback` - Report playback
- `GET /Library/MediaFolders` - Get media folders
- `POST /Library/Sections/{id}/Refresh` - Refresh section

**Files:**
```go
// internal/api/handlers/library.go
type LibraryHandler struct {
    librarySvc *service.LibraryService
}

func (h *LibraryHandler) RegisterRoutes(r chi.Router) {
    r.Get("/Items", h.GetItems)
    r.Get("/Items/{id}", h.GetItem)
    r.Get("/Items/{id}/File", h.GetItemFile)
    // ...
}
```

**Deliverables:**
- All library endpoints work
- Compatible with existing clients

#### 4.5.2 Session API

**Endpoints:**
- `GET /Sessions` - List sessions
- `DELETE /Sessions/{id}` - Delete session
- `POST /Sessions/Playing` - Report playback start
- `POST /Sessions/Playing/Stopped` - Report playback stop
- `POST /Sessions/Playing/Progress` - Report playback progress

**Deliverables:**
- All session endpoints work
- Playback reporting works

#### 4.5.3 User API

**Endpoints:**
- `POST /Users/Authenticate` - Authenticate user
- `GET /Users/{id}` - Get user
- `GET /Users` - List users
- `POST /Users/New` - Create user
- `PUT /Users/{id}` - Update user

**Deliverables:**
- All user endpoints work
- Authentication works

#### 4.5.4 Additional APIs

**Categories:**
- Images API (thumbnails, posters, etc.)
- Videos API
- TV Shows API
- Movies API
- Live TV API
- Search API
- Configuration API
- System API
- Scheduled Tasks API
- Subtitles API
- Playlists API
- Channels API

**Deliverables:**
- All API endpoints implemented
- Compatible with existing clients

### Phase 6: DLNA Support (UPnP Removed)

**Duration:** 3-4 weeks

**Goal:** Implement DLNA server for media discovery (UPnP functionality removed for security reasons)

**Security Note:** UPnP (Universal Plug and Play) has been intentionally excluded from this migration due to well-documented security concerns:
- UPnP has historically been vulnerable to remote exploitation
- SSDP amplification attacks can be used for DDoS
- Automatic port forwarding poses security risks in multi-tenant environments
- Modern networks increasingly block UPnP traffic by default

DLNA browsing and streaming functionality will be preserved for local network media discovery, but UPnP's automatic port forwarding and remote access features will not be implemented.

#### 4.6.1 DLNA Media Server

**Tasks:**
1. Implement ContentDirectory service
2. Implement ConnectionManager service
3. Create DIDL-Lite XML responses
4. Implement browse/search operations
5. Implement protocol info
6. Implement SSDP discovery (local network only)

**Files:**
```go
// internal/dlna/server.go
type DLNAService struct {
    server *HTTPServer
    port   int
}

func (s *DLNAService) Start() error
func (s *DLNAService) Stop() error
```

**Deliverables:**
- DLNA clients can browse library on local network
- Media streaming works via DLNA
- No UPnP remote access functionality

### Phase 6B: Emby Premiere / Supporter Features

**Duration:** 2-3 weeks

**Goal:** Preserve and migrate Emby Premiere licensing and premium features functionality

**Background:** Emby Premiere (also known as "Supporter" or "MBSupporter") is the licensing system that unlocks premium features. This includes:
- Supporter Key validation and management
- License file handling (MBLicenseFile)
- App Store purchase registration
- Premium feature gating (DVR, hardware transcoding, etc.)
- Registration status caching and validation

**Current Implementation Analysis:**
- **PluginSecurityManager.cs** - Main license validation logic
- **MBLicenseFile** - License file storage and retrieval
- **SupporterKey** - User-provided license key
- **IsSupporter()** - Check if user has active license
- **UpdateSupporterKey()** - Update license key
- **RegisterAppStoreSale()** - Handle app store purchases
- **GetRegistrationStatus()** - Validate license with mb3admin.com

#### 4.6B.1 License Management System

**Tasks:**
1. Create license file reader/writer (compatible with existing .mblicense files)
2. Implement SupporterKey storage and retrieval
3. Implement license validation against mb3admin.com API
4. Implement registration status caching (14-day cache)
5. Handle app store purchase registration flow

**Files:**
```go
// internal/service/licensing/manager.go
type LicenseManager struct {
    httpClient   *http.Client
    licensePath  string
    cache        *LicenseCache
    systemID     string
    appVersion   string
}

func (m *LicenseManager) LoadLicense() error
func (m *LicenseManager) GetSupporterKey() string
func (m *LicenseManager) UpdateSupporterKey(key string) error
func (m *LicenseManager) IsSupporter() (bool, error)
func (m *LicenseManager) ValidateWithServer(ctx context.Context) error
func (m *LicenseManager) RegisterAppStoreSale(params string) error
```

**API Endpoints to Preserve:**
- `POST /Plugins/SecurityInfo` - Get/update supporter status
- `POST /Plugins/SecurityInfo/Update` - Update supporter key
- `POST /Plugins/Registration/AppStore` - Register app store purchase

**Deliverables:**
- Existing license keys work without migration
- License validation works (online and cached)
- App store purchases can be registered
- Premium features can be gated by license status

#### 4.6B.2 Premium Feature Gating

**Tasks:**
1. Identify all premium features in current codebase:
   - Hardware acceleration transcoding
   - Live TV and DVR functionality
   - Offline media downloads
   - Premium metadata providers
   - Advanced codec support
2. Create feature flag system based on license status
3. Implement graceful degradation for non-supporters

**Files:**
```go
// internal/service/licensing/features.go
type FeatureGate struct {
    licenseMgr *LicenseManager
}

func (f *FeatureGate) IsEnabled(feature string) (bool, error)
func (f *FeatureGate) RequireSupporter(feature string) error
```

**Premium Features to Gate:**
| Feature | Current Implementation | Go Implementation |
|---------|----------------------|-------------------|
| Hardware Transcoding | FFMpegLoader.cs | internal/service/media/encoder.go |
| Live TV / DVR | LiveTvManager.cs | internal/service/livetv/manager.go |
| Offline Downloads | MediaBrowser.Api | internal/api/handlers/downloads.go |
| Premium Metadata | MediaBrowser.Providers | internal/provider/metadata/premium.go |

**Deliverables:**
- All premium features properly gated
- Clear error messages for non-supporters
- Upgrade prompts in API responses

#### 4.6B.3 License File Compatibility

**Tasks:**
1. Analyze existing .mblicense file format
2. Implement compatible reader/writer in Go
3. Support existing license migration without re-issue
4. Handle license expiration and renewal

**Current License File Location:**
- `%APPDATA%/Emby-Server/mblicense.txt` (Windows)
- `~/.config/emby-server/mblicense.txt` (Linux/FreeBSD)
- License contains: RegKey, expiration date, last checked timestamp

**Deliverables:**
- Existing licenses work without user action
- License file format preserved
- Graceful handling of expired licenses

### Phase 7: WebSocket and Real-time Features

**Duration:** 4-6 weeks

**Goal:** Implement WebSocket server for real-time updates

#### 4.7.1 WebSocket Server

**Tasks:**
1. Set up WebSocket endpoint
2. Implement connection management
3. Implement message types (from C# IWebSocket)
4. Implement authentication for WebSocket

**Files:**
```go
// internal/server/websocket.go
type WebSocketServer struct {
    upgrader gorilla.Websocket.Upgrader
    clients  map[*WebSocketClient]bool
}

func (s *WebSocketServer) HandleConnection(w http.ResponseWriter, r *http.Request)
func (s *WebSocketServer) Broadcast(message *WebSocketMessage)
```

**Deliverables:**
- WebSocket connections work
- Clients receive real-time updates

#### 4.7.2 Message Types

**Messages to Implement:**
- Library changed notifications
- Session updates
- User data changes
- Scheduled task updates
- Server shutdown notifications

**Deliverables:**
- All message types implemented
- Clients receive correct updates

### Phase 8: Image Processing

**Duration:** 4-6 weeks

**Goal:** Implement image processing and thumbnail generation

#### 4.8.1 Image Processing Backend

**Tasks:**
1. Integrate govips (libvips bindings)
2. Implement image resizing
3. Implement format conversion
4. Implement quality control
5. Implement caching

**Files:**
```go
// internal/provider/images/processor.go
type ImageProcessor struct {
    cache *imagecache.Cache
}

func (p *ImageProcessor) Resize(src io.Reader, width, height int) (io.Reader, error)
func (p *ImageProcessor) GenerateThumbnail(src io.Reader) (io.Reader, error)
```

**Deliverables:**
- Image resizing works
- Thumbnails generated
- Caching works

#### 4.8.2 Image API

**Tasks:**
1. Implement `/Items/{id}/Images/{type}` endpoint
2. Implement `/Items/{id}/PrimaryImage` endpoint
3. Implement image caching headers
4. Implement dynamic image processing

**Deliverables:**
- All image endpoints work
- Images served efficiently

### Phase 9: Metadata Providers

**Duration:** 6-8 weeks

**Goal:** Implement metadata providers (local and remote)

#### 4.9.1 Local Metadata

**Tasks:**
1. Parse NFO files
2. Extract embedded metadata from media files
3. Read image files from disk
4. Implement folder structure parsing

**Deliverables:**
- Local metadata extraction works
- NFO files parsed
- Embedded metadata read

#### 4.9.2 Remote Metadata

**Tasks:**
1. Implement provider interface
2. Create HTTP client for metadata services
3. Implement caching
4. Implement rate limiting

**Providers to Consider:**
- The Movie Database (TMDb)
- The TV Database (TVDb)
- The Open Movie Database (OMDb)
- MusicBrainz
- TheAudioDB

**Deliverables:**
- Remote metadata fetching works
- Caching works
- Rate limiting works

### Phase 10: Scheduled Tasks and Background Jobs

**Duration:** 3-4 weeks

**Goal:** Implement scheduled task system

#### 4.10.1 Task Scheduler

**Tasks:**
1. Implement task scheduler
2. Create task interface
3. Implement task persistence
4. Implement task execution

**Files:**
```go
// internal/service/scheduled/tasks.go
type TaskManager struct {
    scheduler *gocron.Scheduler
    tasks     map[string]Task
}

func (m *TaskManager) Start() error
func (m *TaskManager) RegisterTask(task Task) error
func (m *TaskManager) ExecuteTask(taskID string) error
```

**Deliverables:**
- Task scheduler works
- Tasks execute on schedule

#### 4.10.2 Built-in Tasks

**Tasks to Implement:**
- Library scan
- Metadata refresh
- Thumbnail generation
- Log cleanup
- Session cleanup

**Deliverables:**
- All built-in tasks implemented
- Tasks run on schedule

### Phase 11: Testing and Quality Assurance

**Duration:** 6-8 weeks (overlaps with other phases)

**Goal:** Ensure code quality and compatibility

#### 4.11.1 Unit Tests

**Tasks:**
1. Write unit tests for repositories
2. Write unit tests for services
3. Write unit tests for handlers
4. Achieve >80% code coverage

**Deliverables:**
- Comprehensive unit test suite
- CI runs tests on every commit

#### 4.11.2 Integration Tests

**Tasks:**
1. Create test database fixtures
2. Write API integration tests
3. Test end-to-end workflows
4. Test client compatibility

**Deliverables:**
- Integration test suite
- API compatibility verified

#### 4.11.3 Performance Testing

**Tasks:**
1. Benchmark critical paths
2. Load test API endpoints
3. Test with large libraries
4. Profile memory usage

**Deliverables:**
- Performance benchmarks
- Performance meets requirements

### Phase 12: Documentation and Deployment

**Duration:** 4-6 weeks

**Goal:** Prepare for production deployment

#### 4.12.1 Documentation

**Tasks:**
1. Write API documentation
2. Write deployment guide
3. Write configuration guide

**Deliverables:**
- Complete documentation
- No migration guide needed (existing databases work directly)

#### 4.12.2 Deployment

**Note:** Docker configuration will be handled in this final phase, not Phase 1.

**Tasks:**
1. Create Docker images
2. Create installation packages
3. Test on multiple platforms (Linux, FreeBSD, macOS, Windows)

**Deliverables:**
- Production-ready packages
- Docker images
- Installation guides
- No upgrade scripts needed at this time

---

## 5. Data Migration Strategy

### 5.1 SQLite Compatibility

**Good News:** The current C# implementation uses SQLite, which is fully compatible with Go.

**Approach:**
1. Existing SQLite database files can be used directly
2. No data migration needed for schema
3. Go SQLite drivers support all SQLite features used by C#

### 5.2 Schema Verification

**Tasks:**
1. Extract schema from existing C# code
2. Verify all tables, indexes, triggers
3. Document schema in Go code
4. Create migration system for future changes

### 5.3 Data Integrity

**Tasks:**
1. Verify all SQL queries produce same results
2. Test with existing database files
3. Ensure no data corruption
4. Implement backup before any migration

---

## 6. Client Compatibility

### 6.1 API Compatibility

**Strategy:** Maintain 100% API compatibility with existing clients

**Approach:**
1. Document all existing API endpoints from C# code
2. Implement exact same request/response formats
3. Maintain same authentication mechanisms
4. Test with existing clients (web, mobile, TV apps)

### 6.2 Version Detection

**Tasks:**
1. Implement API versioning
2. Support legacy API versions during transition
3. Plan deprecation strategy for old endpoints

### 6.3 Testing with Clients

**Clients to Test:**
- Emby Web Dashboard
- Emby mobile apps (iOS, Android)
- Emby TV apps (Roku, Apple TV, Android TV, etc.)
- Third-party clients (Jellyfin, Kodi, etc.)

---

## 7. Performance Goals

### 7.1 Targets

| Metric | Current (C#) | Target (Go) | Improvement |
|--------|--------------|-------------|-------------|
| Memory Usage | ~500MB | ~200MB | 60% reduction |
| Startup Time | ~10s | ~2s | 80% reduction |
| API Latency (p95) | ~50ms | ~20ms | 60% reduction |
| Concurrent Streams | ~50 | ~200 | 4x increase |
| Library Scan | ~5min | ~2min | 60% reduction |

### 7.2 Optimization Strategies

1. **Use pure Go SQLite driver** - No CGO overhead
2. **Efficient connection pooling** - Reuse database connections
3. **Async I/O** - Leverage Go's goroutines
4. **Caching** - Cache frequently accessed data
5. **Streaming** - Stream large responses
6. **Compression** - Compress responses where appropriate

---

## 8. Risks and Mitigations

| Risk | Impact | Probability | Mitigation |
|------|--------|-------------|------------|
| API incompatibility | High | Medium | Extensive testing with clients, maintain compatibility layer |
| Data corruption | Critical | Low | Thorough testing, backup before migration, rollback plan |
| Performance regression | Medium | Low | Continuous benchmarking, profiling |
| Missing features | Medium | Medium | Feature parity checklist, user feedback |
| FFmpeg integration issues | Medium | Medium | Extensive testing with various media formats |
| DLNA compatibility | Low | Medium | Test with multiple DLNA clients |
| Team learning curve | Low | High | Documentation, code reviews, pair programming |

---

## 9. Testing Strategy

### 9.1 Test Pyramid

```
        /\
       /  \      E2E Tests (10%)
      /----\    
     /      \    Integration Tests (30%)
    /--------\  
   /          \  Unit Tests (60%)
  /------------\
```

### 9.2 Unit Tests

**Scope:**
- Repository methods
- Service business logic
- Handler request/response processing
- Utility functions

**Tools:**
- `testing` package
- `testify` for assertions
- `gomock` for mocking

### 9.3 Integration Tests

**Scope:**
- API endpoints
- Database operations
- Service interactions
- WebSocket connections

**Tools:**
- `httptest` package
- Test containers for database
- WebSocket test clients

### 9.4 End-to-End Tests

**Scope:**
- Complete user workflows
- Client compatibility
- Real-world scenarios

**Tools:**
- Playwright for web UI
- Real Emby clients
- Automated test scripts

### 9.5 Test Harnesses

**Critical Requirement:** All testing must be performed in isolation. Never test on production systems or with production data.

**Test Environments:**
1. **Unit Tests:** Run on every commit, no external dependencies
2. **Integration Tests:** Run in CI with test containers
3. **E2E Tests:** Run on dedicated test infrastructure
4. **Performance Tests:** Run on isolated performance test environment

---

## 10. Implementation Details

### 10.1 File Changes Summary

| Component | C# Files | Go Files | Status |
|-----------|----------|----------|--------|
| HTTP Server | SocketHttpListener/*.cs | internal/server/http.go | Not Started |
| WebSocket | HttpServer/WebSocketConnection.cs | internal/server/websocket.go | Not Started |
| Item Repository | Data/SqliteItemRepository.cs | internal/repository/item.go | Not Started |
| User Repository | Library/UserManager.cs | internal/repository/user.go | Not Started |
| Auth Repository | Security/AuthenticationRepository.cs | internal/repository/auth.go | Not Started |
| Library Service | Library/LibraryManager.cs | internal/service/library/manager.go | Not Started |
| Session Service | Session/SessionManager.cs | internal/service/session/manager.go | Not Started |
| Media Encoder | FFMpegLoader.cs, EncodingManager.cs | internal/service/media/encoder.go | Not Started |
| API Handlers | MediaBrowser.Api/*.cs | internal/api/handlers/*.go | Not Started |
| DLNA Server | Emby.Dlna/*.cs | internal/dlna/*.go | Not Started |
| Image Processing | Emby.Drawing/*.cs | internal/provider/images/*.go | Not Started |
| Metadata Providers | MediaBrowser.Providers/*.cs | internal/provider/metadata/*.go | Not Started |

### 10.2 Configuration Example

```yaml
# config.yaml
server:
  host: "0.0.0.0"
  port: 8096
  https_port: 8920
  certificate_path: ""
  enable_https: false

database:
  path: "/data/emby.db"
  max_connections: 100

library:
  scan_interval: "24h"
  monitor_filesystem: true
  paths:
    - "/media/movies"
    - "/media/tvshows"
    - "/media/music"

transcoding:
  ffmpeg_path: "/usr/bin/ffmpeg"
  temp_path: "/tmp/transcode"
  throttle: false

logging:
  level: "info"
  path: "/var/log/emby"
```

### 10.3 Build and Run

```bash
# Build
make build

# Run
make run

# Test
make test

# Docker
docker build -t emby-go .
docker run -p 8096:8096 -v /data:/data emby-go
```

---

## 11. Testing Strategy and Isolated Test Harnesses

**Critical Requirement:** All testing must be performed in isolation. The test harnesses must never use production data or affect production systems.

### 11.1 Testing Philosophy

| Level | Environment | Purpose |
|-------|-------------|---------|
| Unit Tests | Local development | Test logic in isolation |
| Integration Tests | CI/CD with test containers | Test component interactions |
| E2E Tests | Dedicated test environment | Test complete workflows |
| Performance Tests | Isolated performance lab | Measure performance metrics |
| Compatibility Tests | Client test lab | Verify client compatibility |

### 11.2 Unit Test Harness

**File:** `internal/repository/item_test.go`

```go
package repository

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestItemRepository_GetByID(t *testing.T) {
    // Create in-memory test database
    db := createTestDB(t)
    repo := NewItemRepository(db)
    
    // Insert test data
    expected := &model.Item{
        ID:   "test-123",
        Name: "Test Movie",
        Type: "Movie",
    }
    err := repo.Insert(expected)
    assert.NoError(t, err)
    
    // Retrieve and verify
    actual, err := repo.GetByID("test-123")
    assert.NoError(t, err)
    assert.Equal(t, expected.ID, actual.ID)
    assert.Equal(t, expected.Name, actual.Name)
}

func createTestDB(t *testing.T) *sql.DB {
    db, err := sql.Open("sqlite", ":memory:")
    require.NoError(t, err)
    
    // Create schema
    _, err = db.Exec(schemaSQL)
    require.NoError(t, err)
    
    return db
}
```

**Run:**
```bash
go test ./internal/repository/...
```

### 11.3 Integration Test Harness

**File:** `tests/integration/api_test.go`

```go
package integration

import (
    "testing"
    "net/http/httptest"
    "github.com/stretchr/testify/suite"
)

type APISuite struct {
    suite.Suite
    server *httptest.Server
    client *http.Client
}

func (s *APISuite) SetupSuite() {
    // Create test server with real database
    config := loadTestConfig()
    app := createApp(config)
    s.server = httptest.NewServer(app)
    s.client = s.server.Client()
}

func (s *APISuite) TearDownSuite() {
    s.server.Close()
}

func (s *APISuite) TestGetItems() {
    resp, err := s.client.Get(s.server.URL + "/Items")
    s.Require().NoError(err)
    s.Equal(200, resp.StatusCode)
    
    var items []model.Item
    err = json.NewDecoder(resp.Body).Decode(&items)
    s.Require().NoError(err)
    s.NotEmpty(items)
}

func TestAPISuite(t *testing.T) {
    suite.Run(t, new(APISuite))
}
```

**Run:**
```bash
go test ./tests/integration/...
```

### 11.4 E2E Test Harness

**File:** `tests/e2e/library_test.go`

```go
package e2e

import (
    "testing"
    "time"
    "github.com/playwright-community/playwright-go"
)

func TestLibraryScan(t *testing.T) {
    pw, err := playwright.Run()
    if err != nil {
        t.Fatalf("could not start playwright: %v", err)
    }
    defer pw.Stop()
    
    browser, err := pw.Chromium.Launch()
    if err != nil {
        t.Fatalf("could not launch browser: %v", err)
    }
    defer browser.Close()
    
    page, err := browser.NewPage()
    if err != nil {
        t.Fatalf("could not create page: %v", err)
    }
    
    // Navigate to Emby dashboard
    err = page.Goto("http://localhost:8096")
    if err != nil {
        t.Fatalf("could not goto: %v", err)
    }
    
    // Login
    err = page.Fill("#txtUserName", "admin")
    if err != nil {
        t.Fatalf("could not fill username: %v", err)
    }
    err = page.Fill("#txtManualPassword", "admin")
    if err != nil {
        t.Fatalf("could not fill password: %v", err)
    }
    err = page.Click("#btnSignIn")
    if err != nil {
        t.Fatalf("could not click sign in: %v", err)
    }
    
    // Wait for dashboard
    time.Sleep(2 * time.Second)
    
    // Navigate to library
    err = page.Click("text=Movies")
    if err != nil {
        t.Fatalf("could not click movies: %v", err)
    }
    
    // Verify movies are displayed
    err = page.WaitForSelector(".movieItem")
    if err != nil {
        t.Fatalf("could not wait for movies: %v", err)
    }
    
    t.Log("E2E test passed: Library scan and display works")
}
```

**Run:**
```bash
go test ./tests/e2e/...
```

### 11.5 Performance Test Harness

**File:** `tests/performance/api_bench_test.go`

```go
package performance

import (
    "testing"
    "net/http"
    "net/http/httptest"
)

func BenchmarkGetItem(b *testing.B) {
    // Setup
    config := loadTestConfig()
    app := createApp(config)
    server := httptest.NewServer(app)
    defer server.Close()
    
    client := &http.Client{}
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        resp, err := client.Get(server.URL + "/Items/test-123")
        if err != nil {
            b.Fatal(err)
        }
        resp.Body.Close()
    }
}

func BenchmarkLibraryScan(b *testing.B) {
    // Setup with large test library
    config := loadTestConfig()
    config.Library.Paths = []string{"/test/media/large"}
    app := createApp(config)
    server := httptest.NewServer(app)
    defer server.Close()
    
    client := &http.Client{}
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        resp, err := client.Post(server.URL + "/Library/Sections/1/Refresh", "", nil)
        if err != nil {
            b.Fatal(err)
        }
        resp.Body.Close()
    }
}
```

**Run:**
```bash
go test -bench=. ./tests/performance/...
```

### 11.6 Test Summary

| Test Type | Location | Environment | Production Impact |
|-----------|----------|-------------|-------------------|
| Unit Tests | `internal/*/*_test.go` | Local/CI | **None** |
| Integration Tests | `tests/integration/` | CI with containers | **None** |
| E2E Tests | `tests/e2e/` | Dedicated test env | **None** |
| Performance Tests | `tests/performance/` | Isolated lab | **None** |
| Compatibility Tests | `tests/compatibility/` | Client test lab | **None** |

**Safety Rules:**
1. Never run tests against production database
2. Always use test fixtures and mock data
3. CI tests run in isolated containers
4. Performance tests run on dedicated hardware
5. Client compatibility tests use test clients only

---

## 12. TODO — Step-by-Step Implementation Tracker

This section is the master checklist for implementing the Go migration. Each task includes:
- **Status:** `NOT STARTED` | `IN PROGRESS` | `COMPLETED`
- **Owner:** Who is working on it
- **Start Date:** When work began
- **End Date:** When work finished
- **Dependencies:** What must be done first
- **Files Modified:** What files are touched
- **Notes:** Any blockers, decisions, or context

### Phase 1: Foundation and Infrastructure

| # | Task | Status | Owner | Start | End | Dependencies | Files | Notes |
|---|------|--------|-------|-------|-----|--------------|-------|-------|
| 1.1 | Initialize Go module structure | NOT STARTED | | | | | `go.mod`, `Makefile` | Create project skeleton |
| 1.2 | Create build system (Makefile) | NOT STARTED | | | | 1.1 | `Makefile` | build, test, run targets |
| 1.3 | Set up CI/CD pipeline | NOT STARTED | | | | 1.2 | `.github/workflows/ci.yml` | GitHub Actions |
| 1.4 | Implement logging infrastructure | NOT STARTED | | | | 1.1 | `internal/util/log.go` | Zap logger setup |
| 1.5 | Implement configuration system | NOT STARTED | | | | 1.1 | `internal/config/config.go` | YAML config loading |
| 1.6 | Create Dockerfile | NOT STARTED | | | | 1.2 | `Dockerfile` | Multi-stage build |
| 1.7 | Analyze existing SQLite schema | NOT STARTED | | | | | All `*.cs` files | Extract schema from C# code (reference only) |

### Phase 2: Core HTTP Server and API Framework

| # | Task | Status | Owner | Start | End | Dependencies | Files | Notes |
|---|------|--------|-------|-------|-----|--------------|-------|-------|
| 2.1 | Set up HTTP server with chi router | NOT STARTED | | | | 1.5 | `internal/server/http.go` | Basic server |
| 2.2 | Implement middleware chain | NOT STARTED | | | | 2.1 | `internal/api/middleware/` | Logger, recovery, CORS |
| 2.3 | Implement graceful shutdown | NOT STARTED | | | | 2.1 | `internal/server/http.go` | Signal handling |
| 2.4 | Create API router framework | NOT STARTED | | | | 2.1 | `internal/api/router.go` | Route registration |
| 2.5 | Create nginx reverse proxy config | NOT STARTED | | | | 2.1 | `deploy/nginx.conf` | SSL termination |
| 2.6 | Implement request binding | NOT STARTED | | | | 2.5 | `internal/api/binding.go` | JSON to struct |
| 2.7 | Analyze C# authentication flow | NOT STARTED | | | | | `AuthorizationContext.cs` | Understand auth flow |
| 2.8 | Implement auth middleware | NOT STARTED | | | | 2.7 | `internal/api/middleware/auth.go` | API key + session |

### Phase 3: Data Layer Migration

| # | Task | Status | Owner | Start | End | Dependencies | Files | Notes |
|---|------|--------|-------|-------|-----|--------------|-------|-------|
| 3.1 | Create Go models for items | NOT STARTED | | | | 1.7 | `internal/model/item.go` | Match C# structures |
| 3.2 | Implement base repository | NOT STARTED | | | | | `internal/repository/base.go` | DB connection pool |
| 3.3 | Migrate item repository | NOT STARTED | | | | 3.1, 3.2 | `internal/repository/item.go` | `SqliteItemRepository.cs` |
| 3.4 | Create user models | NOT STARTED | | | | 1.7 | `internal/model/user.go` | Match C# structures |
| 3.5 | Migrate user repository | NOT STARTED | | | | 3.4 | `internal/repository/user.go` | `UserManager.cs` |
| 3.6 | Migrate auth repository | NOT STARTED | | | | 3.5 | `internal/repository/auth.go` | `AuthenticationRepository.cs` |
| 3.7 | Verify password hashing compatibility | NOT STARTED | | | | 3.6 | `internal/repository/auth.go` | Match C# algorithm |
| 3.8 | Test with existing database | NOT STARTED | | | | 3.3-3.7 | | Ensure compatibility |

### Phase 4: Service Layer Implementation

| # | Task | Status | Owner | Start | End | Dependencies | Files | Notes |
|---|------|--------|-------|-------|-----|--------------|-------|-------|
| 4.1 | Implement library manager | NOT STARTED | | | | 3.3 | `internal/service/library/manager.go` | `LibraryManager.cs` |
| 4.2 | Implement media scanner | NOT STARTED | | | | 4.1 | `internal/service/library/scanner.go` | File system scanning |
| 4.3 | Implement file system watcher | NOT STARTED | | | | 4.1 | `internal/service/library/notifier.go` | fsnotify integration |
| 4.4 | Implement session manager | NOT STARTED | | | | 3.6 | `internal/service/session/manager.go` | `SessionManager.cs` |
| 4.5 | Implement user manager | NOT STARTED | | | | 3.5 | `internal/service/user/manager.go` | `UserManager.cs` |
| 4.6 | Implement media encoder | NOT STARTED | | | | | `internal/service/media/encoder.go` | `FFMpegLoader.cs`, `EncodingManager.cs` |
| 4.7 | Implement FFmpeg process management | NOT STARTED | | | | 4.6 | `internal/service/media/ffmpeg.go` | Process control |
| 4.8 | Implement transcoding profiles | NOT STARTED | | | | 4.6 | `internal/service/media/profiles.go` | Quality profiles |
| 4.9 | Implement stream pooling manager | NOT STARTED | | | | | `internal/service/media/stream_manager.go` | Stream sharing architecture |
| 4.10 | Implement active stream tracking | NOT STARTED | | | | 4.9 | `internal/service/media/active_stream.go` | Viewer management |
| 4.11 | Implement stream sharing logic (Live TV) | NOT STARTED | | | | 4.9 | `internal/service/media/stream_manager.go` | Auto-share live streams |
| 4.12 | Implement user prompt for recorded media | NOT STARTED | | | | 4.9 | `internal/api/handlers/playback.go` | "Share Content" vs "Start From Beginning" |
| 4.13 | Implement stream lifecycle management | NOT STARTED | | | | 4.9 | `internal/service/media/stream_manager.go` | Cleanup, health checks |
| 4.14 | Implement resource limits and monitoring | NOT STARTED | | | | 4.9 | `internal/service/media/stream_manager.go` | Max concurrent streams |

### Phase 5: API Endpoints Implementation

| # | Task | Status | Owner | Start | End | Dependencies | Files | Notes |
|---|------|--------|-------|-------|-----|--------------|-------|-------|
| 5.1 | Map all C# API routes | NOT STARTED | | | | | Documentation | Create route mapping |
| 5.2 | Implement library API handlers | NOT STARTED | | | | 4.1 | `internal/api/handlers/library.go` | `LibraryService.cs` |
| 5.3 | Implement session API handlers | NOT STARTED | | | | 4.4 | `internal/api/handlers/sessions.go` | `SessionsService.cs` |
| 5.4 | Implement user API handlers | NOT STARTED | | | | 4.5 | `internal/api/handlers/users.go` | `UserService.cs` |
| 5.5 | Implement images API handlers | NOT STARTED | | | | | `internal/api/handlers/images.go` | `ImageService.cs` |
| 5.6 | Implement videos API handlers | NOT STARTED | | | | | `internal/api/handlers/videos.go` | `VideosService.cs` |
| 5.7 | Implement TV shows API handlers | NOT STARTED | | | | | `internal/api/handlers/tvshows.go` | `TvShowsService.cs` |
| 5.8 | Implement movies API handlers | NOT STARTED | | | | | `internal/api/handlers/movies.go` | `MoviesService.cs` |
| 5.9 | Implement Live TV API handlers | NOT STARTED | | | | | `internal/api/handlers/livetv.go` | `LiveTvService.cs` |
| 5.10 | Implement search API handlers | NOT STARTED | | | | | `internal/api/handlers/search.go` | `SearchService.cs` |
| 5.11 | Implement configuration API handlers | NOT STARTED | | | | | `internal/api/handlers/config.go` | `ConfigurationService.cs` |
| 5.12 | Implement system API handlers | NOT STARTED | | | | | `internal/api/handlers/system.go` | `SystemService.cs` |
| 5.13 | Implement scheduled tasks API handlers | NOT STARTED | | | | | `internal/api/handlers/tasks.go` | `ScheduledTaskService.cs` |
| 5.14 | Implement subtitles API handlers | NOT STARTED | | | | | `internal/api/handlers/subtitles.go` | `SubtitleService.cs` |
| 5.15 | Test API compatibility with clients | NOT STARTED | | | | 5.2-5.14 | | Verify all clients work |

### Phase 6: DLNA Support (UPnP Removed for Security)

| # | Task | Status | Owner | Start | End | Dependencies | Files | Notes |
|---|------|--------|-------|-------|-----|--------------|-------|-------|
| 6.1 | Implement SSDP discovery (local only) | NOT STARTED | | | | | `internal/dlna/server.go` | Local network discovery |
| 6.2 | Create device description XML | NOT STARTED | | | | 6.1 | `internal/dlna/xml/device.go` | DLNA device descriptor |
| 6.3 | Implement ContentDirectory service | NOT STARTED | | | | 6.1 | `internal/dlna/content.go` | Browse/search |
| 6.4 | Implement ConnectionManager service | NOT STARTED | | | | 6.1 | `internal/dlna/connection.go` | Protocol info |
| 6.5 | Create DIDL-Lite XML responses | NOT STARTED | | | | 6.3 | `internal/dlna/xml/didl.go` | Media metadata |
| 6.6 | Test with DLNA clients | NOT STARTED | | | | 6.3-6.5 | | Verify local network compatibility |
| 6.7 | Document UPnP security decision | NOT STARTED | | | | | `docs/upnp-removal.md` | Security rationale |

### Phase 6B: Emby Premiere / Supporter Features

| # | Task | Status | Owner | Start | End | Dependencies | Files | Notes |
|---|------|--------|-------|-------|-----|--------------|-------|-------|
| 6B.1 | Analyze existing .mblicense format | NOT STARTED | | | | | `Emby.Server.Implementations/Security/` | License file format |
| 6B.2 | Implement license file reader/writer | NOT STARTED | | | | 6B.1 | `internal/service/licensing/manager.go` | Compatible with existing |
| 6B.3 | Implement SupporterKey storage | NOT STARTED | | | | 6B.2 | `internal/service/licensing/manager.go` | Key management |
| 6B.4 | Implement license validation API | NOT STARTED | | | | | `internal/service/licensing/manager.go` | mb3admin.com integration |
| 6B.5 | Implement registration status caching | NOT STARTED | | | | 6B.4 | `internal/service/licensing/cache.go` | 14-day cache |
| 6B.6 | Implement app store purchase flow | NOT STARTED | | | | 6B.4 | `internal/service/licensing/manager.go` | App Store registration |
| 6B.7 | Identify premium features | NOT STARTED | | | | | Documentation | Feature inventory |
| 6B.8 | Implement feature gating system | NOT STARTED | | | | 6B.7 | `internal/service/licensing/features.go` | License-based gating |
| 6B.9 | Gate hardware transcoding | NOT STARTED | | | | 6B.8 | `internal/service/media/encoder.go` | Premiere feature |
| 6B.10 | Gate Live TV / DVR | NOT STARTED | | | | 6B.8 | `internal/service/livetv/manager.go` | Premiere feature |
| 6B.11 | Preserve API endpoints | NOT STARTED | | | | 6B.2 | `internal/api/handlers/licensing.go` | /Plugins/SecurityInfo |
| 6B.12 | Test license migration | NOT STARTED | | | | 6B.2-6B.11 | | Existing licenses work |

### Phase 7: WebSocket and Real-time Features

| # | Task | Status | Owner | Start | End | Dependencies | Files | Notes |
|---|------|--------|-------|-------|-----|--------------|-------|-------|
| 7.1 | Set up WebSocket server | NOT STARTED | | | | 2.1 | `internal/server/websocket.go` | Gorilla WebSocket |
| 7.2 | Implement connection management | NOT STARTED | | | | 7.1 | `internal/server/websocket.go` | Client tracking |
| 7.3 | Analyze C# WebSocket messages | NOT STARTED | | | | | `IWebSocket.cs` | Message types |
| 7.4 | Implement message types | NOT STARTED | | | | 7.3 | `internal/server/messages.go` | All message types |
| 7.5 | Implement WebSocket authentication | NOT STARTED | | | | 7.1 | `internal/server/websocket.go` | Token validation |
| 7.6 | Implement broadcast system | NOT STARTED | | | | 7.2 | `internal/server/websocket.go` | Send to all clients |
| 7.7 | Test with existing clients | NOT STARTED | | | | 7.4-7.6 | | Verify real-time updates |

### Phase 8: Image Processing

| # | Task | Status | Owner | Start | End | Dependencies | Files | Notes |
|---|------|--------|-------|-------|-----|--------------|-------|-------|
| 8.1 | Integrate govips library | NOT STARTED | | | | | `internal/provider/images/processor.go` | libvips bindings |
| 8.2 | Implement image resizing | NOT STARTED | | | | 8.1 | `internal/provider/images/processor.go` | Resize operations |
| 8.3 | Implement format conversion | NOT STARTED | | | | 8.1 | `internal/provider/images/processor.go` | JPEG, PNG, WebP |
| 8.4 | Implement image caching | NOT STARTED | | | | 8.2 | `internal/provider/images/cache.go` | Memory + disk cache |
| 8.5 | Implement image API endpoints | NOT STARTED | | | | 8.2 | `internal/api/handlers/images.go` | Serve processed images |
| 8.6 | Test image quality and performance | NOT STARTED | | | | 8.5 | | Verify quality |

### Phase 9: Metadata Providers

| # | Task | Status | Owner | Start | End | Dependencies | Files | Notes |
|---|------|--------|-------|-------|-----|--------------|-------|-------|
| 9.1 | Create provider interface | NOT STARTED | | | | | `internal/provider/metadata/provider.go` | Provider contract |
| 9.2 | Implement local metadata parser | NOT STARTED | | | | 9.1 | `internal/provider/metadata/local.go` | NFO files |
| 9.3 | Implement embedded metadata extractor | NOT STARTED | | | | 9.1 | `internal/provider/metadata/embedded.go` | Media file metadata |
| 9.4 | Implement TMDb provider | NOT STARTED | | | | 9.1 | `internal/provider/metadata/tmdb.go` | The Movie Database |
| 9.5 | Implement TVDb provider | NOT STARTED | | | | 9.1 | `internal/provider/metadata/tvdb.go` | The TV Database |
| 9.6 | Implement MusicBrainz provider | NOT STARTED | | | | 9.1 | `internal/provider/metadata/musicbrainz.go` | Music metadata |
| 9.7 | Implement metadata caching | NOT STARTED | | | | 9.4-9.6 | `internal/provider/metadata/cache.go` | Reduce API calls |
| 9.8 | Implement rate limiting | NOT STARTED | | | | 9.4-9.6 | `internal/provider/metadata/limiter.go` | Respect API limits |

### Phase 10: Scheduled Tasks and Background Jobs

| # | Task | Status | Owner | Start | End | Dependencies | Files | Notes |
|---|------|--------|-------|-------|-----|--------------|-------|-------|
| 10.1 | Implement task scheduler | NOT STARTED | | | | | `internal/service/scheduled/tasks.go` | gocron integration |
| 10.2 | Create task interface | NOT STARTED | | | | 10.1 | `internal/service/scheduled/task.go` | Task contract |
| 10.3 | Implement task persistence | NOT STARTED | | | | 10.1 | `internal/service/scheduled/storage.go` | Store task state |
| 10.4 | Implement library scan task | NOT STARTED | | | | 10.2 | `internal/service/scheduled/scan.go` | Periodic scan |
| 10.5 | Implement metadata refresh task | NOT STARTED | | | | 10.2 | `internal/service/scheduled/refresh.go` | Update metadata |
| 10.6 | Implement thumbnail generation task | NOT STARTED | | | | 10.2 | `internal/service/scheduled/thumbnails.go` | Generate thumbnails |
| 10.7 | Implement log cleanup task | NOT STARTED | | | | 10.2 | `internal/service/scheduled/cleanup.go` | Old log removal |
| 10.8 | Test task scheduling | NOT STARTED | | | | 10.4-10.7 | | Verify schedules |

### Phase 11: Testing and Quality Assurance

| # | Task | Status | Owner | Start | End | Dependencies | Files | Notes |
|---|------|--------|-------|-------|-----|--------------|-------|-------|
| 11.1 | Write repository unit tests | NOT STARTED | | | | 3.3-3.7 | `internal/repository/*_test.go` | Test all repos |
| 11.2 | Write service unit tests | NOT STARTED | | | | 4.1-4.8 | `internal/service/*/*_test.go` | Test all services |
| 11.3 | Write handler unit tests | NOT STARTED | | | | 5.2-5.14 | `internal/api/handlers/*_test.go` | Test all handlers |
| 11.4 | Create integration test suite | NOT STARTED | | | | 5.2-5.14 | `tests/integration/` | API integration |
| 11.5 | Create E2E test suite | NOT STARTED | | | | 5.2-5.14 | `tests/e2e/` | Playwright tests |
| 11.6 | Create performance test suite | NOT STARTED | | | | | `tests/performance/` | Benchmarks |
| 11.7 | Achieve >80% code coverage | NOT STARTED | | | | 11.1-11.3 | | Coverage goal |
| 11.8 | Test with all client types | NOT STARTED | | | | 5.15 | | Web, mobile, TV |

### Phase 12: Documentation and Deployment

| # | Task | Status | Owner | Start | End | Dependencies | Files | Notes |
|---|------|--------|-------|-------|-----|--------------|-------|-------|
| 12.1 | Write API documentation | NOT STARTED | | | | 5.1 | `docs/api.md` | OpenAPI spec |
| 12.2 | Write deployment guide | NOT STARTED | | | | | `docs/deployment.md` | Installation |
| 12.3 | Write configuration guide | NOT STARTED | | | | 1.5 | `docs/configuration.md` | Config options |
| 12.4 | Write migration guide | NOT STARTED | | | | | `docs/migration.md` | C# to Go migration |
| 12.5 | Create Docker images | NOT STARTED | | | | 1.6 | `Dockerfile` | Multi-arch builds |
| 12.6 | Create installation packages | NOT STARTED | | | | 1.2 | `packaging/` | DEB, RPM, etc. |
| 12.7 | Test on multiple platforms | NOT STARTED | | | | 12.6 | | Linux, FreeBSD, macOS, Windows |
| 12.8 | Create upgrade scripts | NOT STARTED | | | | 12.6 | `scripts/upgrade.sh` | Migration scripts |

---

## 13. Future Enhancements

1. **GraphQL API** - Alternative to REST for flexible queries
2. **gRPC internal services** - Microservices architecture option
3. **Plugin system in Go** - Native Go plugin support
4. **Hardware acceleration** - Better GPU transcoding support
5. **AV1 encoding** - Next-gen codec support
6. **Cloud integration** - S3, Google Drive, OneDrive support
7. **AI/ML features** - Auto-tagging, content recommendations
8. **Multi-server clustering** - Horizontal scaling

---

## 14. Conclusion

The migration of Emby Server from C#/.NET to Go is a significant undertaking that offers substantial benefits:

**Benefits:**
- **Performance:** Go's native compilation and efficient runtime provide better performance
- **Resource Efficiency:** Lower memory footprint and CPU usage
- **Simplified Deployment:** Single binary, no runtime dependencies
- **Better Concurrency:** Go's goroutines for handling concurrent requests
- **Easier Maintenance:** Simpler codebase, easier to understand and modify

**Key Success Factors:**
1. **Maintain API Compatibility:** Existing clients must continue to work without modification
2. **Data Integrity:** Existing SQLite databases must work without corruption
3. **Incremental Approach:** Migrate component by component, test thoroughly
4. **Comprehensive Testing:** Extensive unit, integration, and E2E testing
5. **Documentation:** Clear documentation for users and developers

**Risk Mitigation:**
- Parallel operation during transition (run C# and Go versions side-by-side)
- Rollback plan if issues arise
- Extensive testing before production deployment
- Community feedback throughout development

This plan provides a roadmap for a successful migration while minimizing risks and ensuring continuity for existing users.

---

**Document Version:** 1.0  
**Last Updated:** 2026-04-27  
**Author:** Junie (JetBrains AI Assistant)  
**Based on:** FreeBSD PPPoE multithreaded plan format
