# Emby.Server.Implementations - Server Core Migration Plan

## Overview

This document covers the migration of the Emby.Server.Implementations core components from C# to Go.

**Discovery Document:** `.discovery/160-emby-server-impl.md`, `.discovery/161-emby-server-impl-core.md`  
**Priority:** HIGH  
**Status:** PARTIAL  

---

## 1. Component Summary

### 1.1 Module Structure

| Subdirectory | Files | Status | Priority |
|--------------|-------|--------|----------|
| AppBase | 15 | Partial | High |
| Activity | 5 | Complete | Medium |
| Archiving | 3 | Not Started | Low |
| Branding | 3 | Complete | Medium |
| Browser | 2 | Not Started | Low |
| Channels | 25+ | Partial | Medium |
| Collections | 3 | Complete | Medium |
| Configuration | 10+ | Complete | High |
| Cryptography | 5 | Complete | High |
| Data | 50+ | Complete | High |
| Diagnostics | 5 | Not Started | Low |
| dto | 10+ | Partial | Medium |
| Encoding | 30+ | Partial | High |
| EntryPoints | 10+ | Complete | Medium |
| EnvironmentInfo | 5 | Complete | Low |
| HttpClient | 5 | Complete | Medium |
| HttpServer | 50+ | Partial | High |
| Images | 20+ | Complete | Medium |
| IO | 20+ | Complete | High |
| Library | 60+ | Partial | High |
| LiveTv | 40+ | Not Started | High |
| Localization | 15+ | Complete | Medium |
| Logging | 10+ | Complete | High |
| MediaEncoder | 15+ | Complete | High |
| Net | 10+ | Complete | Medium |
| Networking | 15+ | Partial | Medium |
| News | 3 | Not Started | Low |
| Notifications | 5 | Complete | Medium |
| Photos | 5 | Complete | Low |
| Playlists | 5 | Complete | Medium |
| Reflection | 3 | Not Started | Low |
| ScheduledTasks | 20+ | Complete | High |
| Security | 20+ | Complete | High |
| Serialization | 10+ | Complete | Medium |
| Services | 15+ | Complete | High |
| Session | 30+ | Complete | High |
| SharpCifs | 10+ | Not Started | Low |
| Sorting | 5 | Complete | Low |
| TextEncoding | 10+ | Complete | Medium |
| Threading | 5 | Not Started | Low |
| TV | 10+ | Not Started | Medium |
| Updates | 15+ | Complete | High |
| Udp | 5 | Not Started | Low |
| UserViews | 10+ | Complete | Medium |
| Xml | 5 | Complete | Low |

---

## 2. Discovery to Implementation Mapping

### 2.1 Core Components

| Discovery Doc | C# Class | Go Package | Status |
|--------------|----------|-----------|--------|
| `161-emby-server-impl-core.md` | ApplicationHost | `internal/server/` | ✓ |
| `161-emby-server-impl-core.md` | BaseConfigurationManager | `internal/config/` | ✓ |
| `161-emby-server-impl-core.md` | BaseApplicationPaths | `internal/config/` | ✓ |
| `161-emby-server-impl-core.md` | LibraryManager | `internal/service/library/` | ✓ |
| `161-emby-server-impl-core.md` | SessionManager | `internal/service/session/` | ✓ |
| `161-emby-server-impl-core.md` | UserManager | `internal/service/user/` | ✓ |
| `161-emby-server-impl-core.md` | DeviceManager | `internal/service/device/` | ✓ |
| `161-emby-server-impl-core.md` | ProviderManager | `internal/service/metadata/` | ⚠️ |

### 2.2 HTTP Server Components

| Discovery Doc | C# Class | Go Package | Status |
|--------------|----------|-----------|--------|
| `164-emby-server-impl-http.md` | HttpListenerHost | `internal/server/` | ✓ |
| `164-emby-server-impl-http.md` | WebSocketConnection | `internal/server/ws/` | ✓ |
| `164-emby-server-impl-http.md` | AuthorizationContext | `internal/service/auth/` | ✓ |
| `164-emby-server-impl-http.md` | SessionContext | `internal/service/session/` | ✓ |
| `164-emby-server-impl-http.md` | ApiEntryPoint | `internal/api/` | ✓ |

### 2.3 Library Components

| Discovery Doc | C# Class | Go Package | Status |
|--------------|----------|-----------|--------|
| `162-emby-server-impl-library.md` | LibraryManager | `internal/service/library/` | ✓ |
| `162-emby-server-impl-library.md` | UserManager | `internal/service/user/` | ✓ |
| `162-emby-server-impl-library.md` | ItemResolver | `internal/service/library/` | ⚠️ |
| `162-emby-server-impl-library.md` | MediaSourceManager | `internal/service/media/` | ✓ |
| `162-emby-server-impl-library.md` | CollectionManager | `internal/service/library/` | ✓ |
| `162-emby-server-impl-library.md` | PlaylistManager | `internal/service/library/` | ✓ |

---

## 3. Key Implementation Details

### 3.1 ApplicationHost (Main Entry Point)

**C# Source:** `Emby.Server.Implementations/ApplicationHost.cs`

```csharp
// Key interfaces to implement
public interface IServerApplicationHost {
    string Name { get; }
    string SystemId { get; }
    Version Version { get; }
    string GetApiKey();
    string GetFriendlyVersion();
    Task StartAsync();
    Task StopAsync();
}
```

**Go Implementation:** `emby-go/cmd/emby-server/main.go`

```go
func main() {
    cfg := config.Load()
    srv := server.NewServer(cfg)
    
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()
    
    if err := srv.Start(ctx); err != nil {
        log.Fatal(err)
    }
}
```

### 3.2 LibraryManager

**C# Source:** `Emby.Server.Implementations/Library/LibraryManager.cs`

| Method | Description | Go Status |
|--------|-------------|-----------|
| `GetItemInfo()` | Get item metadata | ✓ |
| `ValidateMediaLibrary()` | Validate library | ✓ |
| `GetMediaFolders()` | Get media folders | ✓ |
| `AddVirtualFolder()` | Add library | ✓ |
| `RemoveVirtualFolder()` | Remove library | ✓ |
| `CreateItem()` | Create new item | ✓ |
| `DeleteItem()` | Delete item | ✓ |
| `GetItem()` | Get item by ID | ✓ |
| `GetItems()` | Get items by query | ✓ |

### 3.3 SessionManager

**C# Source:** `Emby.Server.Implementations/Session/SessionManager.cs`

| Property | Type | Description |
|----------|------|-------------|
| `Sessions` | `IEnumerable<SessionInfo>` | Active sessions |

**Go Implementation:** `emby-go/internal/service/session/session.go`

---

## 4. Missing Components

### 4.1 Live TV Module

**Discovery:** `.discovery/170-emby-server-impl-livetv.md`

| Component | C# Files | Status |
|-----------|----------|--------|
| EmbyTV | 10+ | ✗ |
| SchedulesDirect | 5+ | ✗ |
| HDHomerunManager | 3+ | ✗ |
| MediaBrowser.Model.LiveTv | 20+ | ✗ |

### 4.2 SharpCifs (SMB Client)

**Discovery:** `.discovery/169-emby-server-impl-sharpcifs.md`

| Component | Status |
|-----------|--------|
| SmbFile | ✗ |
| SmbSession | ✗ |
| NtlmAuth | ✗ |

### 4.3 Additional Missing

| Component | Discovery | Status |
|-----------|-----------|--------|
| TextEncoding/UniversalDetector | `216-emby-server-impl-textencoding.md` | ✗ |
| ActivityManager | `180-activity.md` | ✓ |
| LocalizationManager | `228-emby-server-impl-localization.md` | ✗ |

---

## 5. Migration Tasks

### 5.1 Priority 1 (Critical)

| # | Task | Status | Files |
|---|------|--------|-------|
| 1.1 | Complete API layer coverage | ⚠️ | `internal/api/handlers/` |
| 1.2 | Implement SocketHttpListener equivalent | ✗ | `internal/server/` |
| 1.3 | Verify session management parity | ✓ | `internal/service/session/` |
| 1.4 | Verify library management parity | ⚠️ | `internal/service/library/` |

### 5.2 Priority 2 (High)

| # | Task | Status | Files |
|---|------|--------|-------|
| 2.1 | Implement Live TV support | ✗ | `internal/service/livetv/` |
| 2.2 | Implement channel management | ✗ | `internal/service/channel/` |
| 2.3 | Implement DLNA full stack | ⚠️ | `internal/dlna/` |
| 2.4 | Implement notifications full stack | ⚠️ | `internal/service/notification/` |

### 5.3 Priority 3 (Medium)

| # | Task | Status | Files |
|---|------|--------|-------|
| 3.1 | Implement SharpCifs/SMB client | ✗ | `internal/network/smb/` |
| 3.2 | Implement text encoding detection | ✗ | `internal/text/` |
| 3.3 | Implement localization | ✗ | `internal/i18n/` |
| 3.4 | Implement UPnP device discovery | ✗ | `internal/discovery/` |

### 5.4 Priority 4 (Low)

| # | Task | Status | Files |
|---|------|--------|-------|
| 4.1 | Implement BDInfo parser | ✗ | `internal/media/bdinfo/` |
| 4.2 | Implement DVD parser | ✗ | `internal/media/dvd/` |
| 4.3 | Implement Mono.Nat | ✗ | `internal/nat/` |

---

## 6. API Coverage

### 6.1 Implemented Endpoints

| Handler | Endpoints | Status |
|---------|-----------|--------|
| `activity.go` | 2 | ✓ |
| `branding.go` | 3 | ✓ |
| `channel.go` | 8 | ✓ |
| `config.go` | 5 | ✓ |
| `device.go` | 5 | ✓ |
| `displayprefs.go` | 3 | ⚠️ |
| `environment.go` | 4 | ✓ |
| `filter.go` | 2 | ✓ |
| `games.go` | 3 | ⚠️ |
| `image.go` | 8 | ✓ |
| `library.go` | 15 | ✓ |
| `livetv.go` | 10 | ⚠️ |
| `localization.go` | 5 | ✓ |
| `media.go` | 8 | ✓ |
| `movies.go` | 8 | ✓ |
| `notification.go` | 6 | ✓ |
| `package.go` | 4 | ⚠️ |
| `playback.go` | 5 | ⚠️ |
| `playlist.go` | 6 | ✓ |
| `scheduledtask.go` | 4 | ✓ |
| `search.go` | 3 | ✓ |
| `session.go` | 10 | ✓ |
| `startup.go` | 3 | ✓ |
| `system.go` | 8 | ✓ |
| `transcoding.go` | 5 | ✓ |
| `tvshows.go` | 8 | ✓ |
| `user.go` | 12 | ✓ |

### 6.2 Missing Endpoints

| Handler | Endpoints | Notes |
|---------|-----------|-------|
| Collection endpoints | 5 | Not implemented |
| UserViews endpoints | 4 | Not implemented |
| Sync endpoints | 3 | Not implemented |
| HdHomerun endpoints | 2 | Not implemented |
| SchedulesDirect endpoints | 4 | Not implemented |

---

## 7. Verification Checklist

- [x] ApplicationHost equivalent implemented
- [x] Configuration management implemented
- [x] Session management implemented
- [x] User management implemented
- [x] Library management implemented
- [x] Media source management implemented
- [x] Collection management implemented
- [x] Playlist management implemented
- [x] HTTP server implemented
- [x] WebSocket support implemented
- [x] API router implemented
- [x] Auth middleware implemented
- [ ] Live TV support
- [ ] Channel management
- [ ] DLNA full stack
- [ ] Notifications full stack
- [ ] SharpCifs/SMB client
- [ ] Text encoding detection
- [ ] Localization

---

## Appendix: Related Documents

- [Master Migration Plan](./000-migration-master-plan.md)
- [Discovery: Server Core](./.discovery/161-emby-server-impl-core.md)
- [Discovery: Library](./.discovery/162-emby-server-impl-library.md)
- [Discovery: HTTP](./.discovery/164-emby-server-impl-http.md)
- [Discovery: LiveTV](./.discovery/170-emby-server-impl-livetv.md)
- [Discovery: Tasks](./.discovery/165-emby-server-impl-tasks.md)

---

**Document Version:** 1.0  
**Last Updated:** 2026-05-04  
**Status:** In Progress
