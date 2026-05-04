# Component: emby-go

**Path:** \`emby-go/\`
**Type:** Directory | Module
**Language:** Go
**Maps to:** \`.discovery/360-emby-go.md\`

## Decomposition

### main.go (Entry Point)

#### Package
\`package main\`

#### Imports
```go
import (
    "context"
    "log"
    "os"
    "os/signal"
    "syscall"
)
```

#### Key Functions
```go
func main()
func run(ctx context.Context) error
```

### internal/server/server.go (HTTP Server)

#### Package
\`package server\`

#### Imports
```go
import (
    "context"
    "net/http"
    "github.com/MediaBrowser/emby-go/internal/api"
    "github.com/MediaBrowser/emby-go/internal/config"
)
```

#### Key Functions
```go
func NewServer(cfg *config.Config) *Server
func (s *Server) Start(ctx context.Context) error
func (s *Server) Shutdown(ctx context.Context) error
```

### internal/api/api.go (API Router)

#### Package
\`package api\`

#### Key Types
```go
type Router struct {
    mux *http.ServeMux
    handlers map[string]http.Handler
}
```

#### Key Functions
```go
func NewRouter() *Router
func (r *Router) Handle(pattern string, handler http.Handler)
func (r *Router) ServeHTTP(w http.ResponseWriter, r *http.Request)
```

### internal/service/library/library.go (Library Service)

#### Package
\`package library\`

#### Key Types
```go
type LibraryService struct {
    repo Repository
}
```

#### Key Functions
```go
func NewLibraryService(repo Repository) *LibraryService
func (s *LibraryService) GetItems(ctx context.Context, query *Query) (*Result, error)
func (s *LibraryService) GetItem(ctx context.Context, id string) (*Item, error)
```

### internal/model/*.go (Data Models)

#### Key Types
```go
type Item struct {
    ID          string
    Name        string
    Type        string
    Path        string
    MediaSources []MediaSource
}

type MediaSource struct {
    ID       string
    Path     string
    Duration int
}
```

## Description

emby-go contains Go language bindings and utilities for Emby Server. It provides Go packages for interacting with the Emby API and server functionality. Contains 95 Go files.

## Directories

- `bin/` — Build output binaries
- `cmd/` — Command-line tools
- `cmd/emby-server/` — Emby server command
- `configs/` — Configuration files
- `docs/` — Documentation files
- `internal/` — Internal packages (68 Go files)
- `migrations/` — Database migrations
- `packaging/` — Packaging scripts
- `pkg/` — Public packages
- `tests/` — Test files

## Internal Packages

- `internal/api/` — 30 Go files
- `internal/api/handlers/` — 27 Go files
- `internal/api/middleware/` — 2 Go files
- `internal/config/` — 2 Go files
- `internal/database/` — 1 Go files
- `internal/dlna/` — 2 Go files
- `internal/dlna/xml/` — 1 Go files
- `internal/licensing/` — License management
- `internal/logging/` — 1 Go files
- `internal/model/` — 5 Go files
- `internal/plugin/` — 1 Go files
- `internal/provider/` — Media provider services
- `internal/provider/images/` — Image provider
- `internal/provider/metadata/` — Metadata provider
- `internal/repository/` — 3 Go files
- `internal/server/` — 2 Go files
- `internal/server/ws/` — 1 Go files
- `internal/service/` — 21 Go files
- `internal/service/auth/` — 1 Go files
- `internal/service/device/` — 1 Go files
- `internal/service/image/` — 2 Go files
- `internal/service/library/` — 4 Go files
- `internal/service/media/` — 2 Go files
- `internal/service/metadata/` — 3 Go files
- `internal/service/notification/` — 1 Go files
- `internal/service/scheduled/` — 1 Go files
- `internal/service/session/` — 3 Go files
- `internal/service/transcoding/` — 1 Go files
- `internal/service/user/` — 2 Go files
- `internal/util/` — Utility packages
- `internal/util/fs/` — Filesystem utilities
- `internal/util/hash/` — Hash utilities
- `internal/util/mime/` — MIME type utilities

## Test Directories

- `tests/` — 3 Go files
- `tests/e2e/` — 1 Go files
- `tests/integration/` — 1 Go files
- `tests/performance/` — 1 Go files

## Root Files


## Project Files

- `go.mod` — emby-go/go.mod
- `go.sum` — emby-go/go.sum
- `Makefile` — emby-go/Makefile
