# C# to Go Migration Guide

## Overview

This document describes the migration of Emby Server from C# (.NET) to Go. The migration preserves the core functionality while leveraging Go's simplicity, performance, and deployment model.

## Architecture Changes

### Runtime Environment

| Aspect | C# (.NET) | Go |
|--------|-----------|-----|
| Runtime | .NET Runtime | Native binary |
| Memory Management | Garbage Collector | Garbage Collector |
| Concurrency | Tasks/Async-Await | Goroutines/Channels |
| HTTP Server | Kestrel | net/http + Chi |
| Database | SQLite (Microsoft.Data.Sqlite) | SQLite (mattn/go-sqlite3) |
| Configuration | JSON/YAML | YAML + env vars |
| Logging | Serilog | Zap |

### Key Differences

1. **No Dependency Injection Framework**: Go uses simple struct composition instead of DI containers.
2. **No ORM**: Repository pattern with raw SQL queries instead of Entity Framework.
3. **No Reflection-Based Serialization**: Manual JSON marshaling/unmarshaling.
4. **No Async/Await**: Goroutines replace async/await for concurrency.
5. **No NuGet Packages**: Go modules replace NuGet packages.

## Migration Steps

### Phase 1: Foundation

1. Initialize Go module with `go mod init`
2. Create project structure following Go conventions
3. Set up build system with Makefile
4. Configure logging with Zap
5. Implement configuration system

### Phase 2: HTTP Server

1. Set up Chi router for HTTP routing
2. Implement middleware chain
3. Add graceful shutdown support
4. Create API router framework

### Phase 3: Data Layer

1. Create Go models matching C# structures
2. Implement base repository with raw SQL
3. Migrate item repository
4. Create user models
5. Migrate user repository

### Phase 4: Services

1. Implement library manager
2. Implement media scanner
3. Implement file system watcher (fsnotify)
4. Implement session manager
5. Implement user manager
6. Implement media encoder (FFmpeg)
7. Implement transcoding profiles

### Phase 5: API Handlers

1. Map all C# API routes
2. Implement library API handlers
3. Implement session API handlers
4. Implement user API handlers
5. Implement images API handlers
6. Implement videos API handlers
7. Implement TV shows API handlers
8. Implement movies API handlers
9. Implement Live TV API handlers
10. Implement search API handlers
11. Implement configuration API handlers
12. Implement system API handlers
13. Implement scheduled tasks API handlers
14. Implement subtitles API handlers

### Phase 6: DLNA Support

1. Implement SSDP discovery
2. Create device description XML
3. Implement ContentDirectory service
4. Implement ConnectionManager service
5. Create DIDL-Lite XML responses

### Phase 7: Plugin System

1. Define plugin interface
2. Implement plugin loader
3. Create plugin manager

### Phase 8: Authentication

1. Analyze C# authentication flow
2. Implement auth middleware
3. Implement password hashing compatibility

### Phase 9: Metadata Providers

1. Implement metadata fetcher
2. Create image downloader
3. Implement metadata cache

### Phase 10: Notifications

1. Implement notification manager
2. Create notification providers

### Phase 11: Testing

1. Write repository unit tests
2. Write service unit tests
3. Write handler unit tests
4. Create integration test suite
5. Create E2E test suite
6. Create performance test suite
7. Achieve >80% code coverage

### Phase 12: Documentation and Deployment

1. Write API documentation
2. Write deployment guide
3. Write configuration guide
4. Write migration guide
5. Create Docker images
6. Create installation packages

## Code Examples

### C# vs Go: Repository Pattern

**C#:**
```csharp
public class SqliteItemRepository : IItemRepository
{
    private readonly SqliteConnection _connection;
    
    public async Task<Item> GetItemAsync(string id)
    {
        var cmd = _connection.CreateCommand();
        cmd.CommandText = "SELECT * FROM Items WHERE Id = @id";
        cmd.Parameters.AddWithValue("@id", id);
        // ...
    }
}
```

**Go:**
```go
type ItemRepository struct {
    db *sql.DB
}

func (r *ItemRepository) GetItem(id string) (map[string]interface{}, error) {
    query := "SELECT * FROM Items WHERE Id = ?"
    row := r.db.QueryRow(query, id)
    // ...
}
```

### C# vs Go: HTTP Handler

**C#:**
```csharp
[HttpGet("Library/Root")]
public IActionResult GetLibraryRoot()
{
    return Ok(new { Name = "Media Library" });
}
```

**Go:**
```go
func (h *LibraryHandler) GetLibraryRoot(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]interface{}{
        "Name": "Media Library",
    })
}
```

### C# vs Go: Configuration

**C#:**
```csharp
var config = new ConfigurationBuilder()
    .AddJsonFile("appsettings.json")
    .AddEnvironmentVariables()
    .Build();
```

**Go:**
```go
cfg, err := config.LoadConfig("")
if err != nil {
    log.Fatal(err)
}
```

## Deployment

### Docker

```bash
docker build -t emby-server .
docker run -p 8096:8096 emby-server
```

### Native Binary

```bash
make build
./bin/emby-server
```

## Testing

### Unit Tests

```bash
go test ./internal/repository/...
go test ./internal/service/...
```

### Integration Tests

```bash
go test ./tests/integration/...
```

### Performance Tests

```bash
go test -bench=. ./tests/performance/...
```

## Known Limitations

1. **No EF Core**: Raw SQL queries instead of ORM
2. **No DI Container**: Manual dependency injection
3. **No Reflection**: Manual JSON marshaling
4. **No Async/Await**: Goroutines for concurrency
5. **No NuGet**: Go modules for dependencies

## Future Work

1. Implement full authentication flow
2. Add more metadata providers
3. Implement plugin system
4. Add more DLNA features
5. Improve transcoding profiles
6. Add more API endpoints
7. Implement Live TV features
8. Add mobile app support

## Conclusion

This migration preserves the core functionality of Emby Server while leveraging Go's simplicity and performance. The Go version is easier to deploy, has a smaller footprint, and is more maintainable than the C# version.
