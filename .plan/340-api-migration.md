# MediaBrowser.Api - API Migration Plan

## Overview

This document covers the migration of the MediaBrowser.Api REST API endpoints from C# to Go.

**Discovery Document:** `.discovery/340-mediabrowser-api.md`  
**Priority:** HIGH  
**Status:** PARTIAL (~80% complete)  

---

## 1. API Overview

### 1.1 API Structure

| Component | Files | Status | Coverage |
|-----------|-------|--------|----------|
| Services | 61 | ✓ | ~90% |
| Controllers | 25+ | ✓ | ~85% |
| Models | 40+ | ✓ | ~80% |
| Attributes | 10+ | ✓ | ~100% |

### 1.2 Service Categories

| Category | Services | Implemented | Coverage |
|----------|----------|------------|----------|
| Library | 15 | 12 | 80% |
| Session | 10 | 10 | 100% |
| User | 12 | 12 | 100% |
| Media | 8 | 8 | 100% |
| System | 8 | 8 | 100% |
| LiveTV | 15 | 5 | 33% |
| Channels | 8 | 6 | 75% |
| Images | 8 | 8 | 100% |
| Playback | 5 | 2 | 40% |
| Notifications | 6 | 6 | 100% |
| Playlists | 6 | 6 | 100% |
| ScheduledTasks | 4 | 4 | 100% |
| Branding | 3 | 3 | 100% |
| Localization | 5 | 5 | 100% |
| Configuration | 5 | 5 | 100% |

---

## 2. Discovery to Implementation Mapping

### 2.1 Root API Files

| Discovery Doc | C# File | Go Handler | Status |
|--------------|---------|-----------|--------|
| `340-mediabrowser-api.md` | BaseApiService.cs | `internal/api/router.go` | ✓ |
| `340-mediabrowser-api.md` | ApiEntryPoint.cs | `internal/api/handlers/system.go` | ✓ |
| `340-mediabrowser-api.md` | IServiceSelector.cs | N/A | ✓ |

### 2.2 Service Implementations

| C# Service | Go Handler | Status |
|-------------|-----------|--------|
| LibraryService | `library.go` | ✓ |
| SessionsService | `session.go` | ✓ |
| UserService | `user.go` | ✓ |
| ItemUpdateService | `library.go` | ✓ |
| ChannelService | `channel.go` | ✓ |
| ImageService | `image.go` | ✓ |
| SystemService | `system.go` | ✓ |
| BrandingService | `branding.go` | ✓ |
| LocalizationService | `localization.go` | ✓ |
| ConfigurationService | `config.go` | ✓ |
| NotificationService | `notification.go` | ✓ |
| PlaylistService | `playlist.go` | ✓ |
| ScheduledTaskService | `scheduledtask.go` | ✓ |
| MediaService | `media.go` | ✓ |
| StartupService | `startup.go` | ✓ |
| ActivityService | `activity.go` | ✓ |
| DisplayPreferencesService | `displayprefs.go` | ⚠️ |
| GamesService | `games.go` | ⚠️ |
| LiveTvService | `livetv.go` | ⚠️ |
| PackageService | `package.go` | ⚠️ |
| SyncService | — | ✗ |
| CollectionService | — | ✗ |

---

## 3. Implemented API Endpoints

### 3.1 Library Endpoints

| Method | Path | Handler | Status |
|--------|------|---------|--------|
| GET | `/Library/VirtualFolders` | `library.go` | ✓ |
| POST | `/Library/VirtualFolders` | `library.go` | ✓ |
| DELETE | `/Library/VirtualFolders/{name}` | `library.go` | ✓ |
| GET | `/Items` | `library.go` | ✓ |
| GET | `/Items/{id}` | `library.go` | ✓ |
| POST | `/Items/{id}` | `library.go` | ✓ |
| DELETE | `/Items/{id}` | `library.go` | ✓ |
| GET | `/Shows/{id}` | `tvshows.go` | ✓ |
| GET | `/Shows/{id}/Episodes` | `tvshows.go` | ✓ |
| GET | `/Movies/{id}` | `movies.go` | ✓ |
| GET | `/Music/{id}` | `library.go` | ✓ |
| GET | `/Genres` | `library.go` | ✓ |
| GET | `/Studios` | `library.go` | ✓ |

### 3.2 Session Endpoints

| Method | Path | Handler | Status |
|--------|------|---------|--------|
| GET | `/Sessions` | `session.go` | ✓ |
| POST | `/Sessions/Logout` | `session.go` | ✓ |
| POST | `/Sessions/{id}/Logout` | `session.go` | ✓ |
| POST | `/Sessions/{id}/Ping` | `session.go` | ✓ |
| POST | `/Sessions/{id}/Command` | `session.go` | ✓ |
| GET | `/Sessions/{id}` | `session.go` | ✓ |
| POST | `/Sessions/{id}/Playing` | `playback.go` | ⚠️ |
| POST | `/Sessions/{id}/Playing/{command}` | `playback.go` | ⚠️ |

### 3.3 User Endpoints

| Method | Path | Handler | Status |
|--------|------|---------|--------|
| GET | `/Users` | `user.go` | ✓ |
| GET | `/Users/{id}` | `user.go` | ✓ |
| POST | `/Users` | `user.go` | ✓ |
| PUT | `/Users/{id}` | `user.go` | ✓ |
| DELETE | `/Users/{id}` | `user.go` | ✓ |
| GET | `/Users/{id}/Views` | `user.go` | ✓ |
| GET | `/Users/{id}/Items` | `user.go` | ✓ |
| POST | `/Users/{id}/Authenticate` | `user.go` | ✓ |
| GET | `/Users/{id}/Policy` | `user.go` | ✓ |
| PUT | `/Users/{id}/Policy` | `user.go` | ✓ |
| GET | `/Users/{id}/Password` | `user.go` | ✓ |
| POST | `/Users/{id}/Password` | `user.go` | ✓ |

### 3.4 System Endpoints

| Method | Path | Handler | Status |
|--------|------|---------|--------|
| GET | `/System/Info` | `system.go` | ✓ |
| GET | `/System/Info/Public` | `system.go` | ✓ |
| GET | `/System/Ping` | `system.go` | ✓ |
| GET | `/System/Logs` | `system.go` | ✓ |
| GET | `/System/Logs/{name}` | `system.go` | ✓ |
| DELETE | `/System/Logs` | `system.go` | ✓ |
| GET | `/System/Configuration` | `config.go` | ✓ |
| POST | `/System/Configuration` | `config.go` | ✓ |

### 3.5 Image Endpoints

| Method | Path | Handler | Status |
|--------|------|---------|--------|
| GET | `/Items/{id}/Images` | `image.go` | ✓ |
| GET | `/Items/{id}/Images/{type}` | `image.go` | ✓ |
| GET | `/Items/{id}/Images/{type}/{index}` | `image.go` | ✓ |
| POST | `/Items/{id}/Images/{type}` | `image.go` | ✓ |
| DELETE | `/Items/{id}/Images/{type}/{index}` | `image.go` | ✓ |
| GET | `/Images/{name}` | `image.go` | ✓ |
| GET | `/Images/Genres/{name}` | `image.go` | ✓ |
| GET | `/Images/Studios/{name}` | `image.go` | ✓ |

---

## 4. Missing API Endpoints

### 4.1 LiveTV Endpoints

| Method | Path | Status | Notes |
|--------|------|--------|-------|
| GET | `/LiveTv/Channels` | ⚠️ | Basic |
| GET | `/LiveTv/Channels/{id}` | ⚠️ | Basic |
| GET | `/LiveTv/Programs` | ✗ | Not implemented |
| GET | `/LiveTv/Programs/{id}` | ✗ | Not implemented |
| GET | `/LiveTv/Recordings` | ✗ | Not implemented |
| GET | `/LiveTv/Recordings/{id}` | ✗ | Not implemented |
| DELETE | `/LiveTv/Recordings/{id}` | ✗ | Not implemented |
| POST | `/LiveTv/Recordings/{id}/Timer` | ✗ | Not implemented |
| GET | `/LiveTv/Timers` | ✗ | Not implemented |
| POST | `/LiveTv/Timers` | ✗ | Not implemented |
| GET | `/LiveTv/SchedulesDirect/Status` | ✗ | Not implemented |
| GET | `/LiveTv/Lineups` | ✗ | Not implemented |
| GET | `/LiveTv/Lineups/{id}` | ✗ | Not implemented |
| GET | `/LiveTv/Hdhr/Channels` | ✗ | Not implemented |
| GET | `/LiveTv/Hdhr/Channels/{id}` | ✗ | Not implemented |

### 4.2 Sync Endpoints

| Method | Path | Status |
|--------|------|--------|
| GET | `/Sync` | ✗ |
| POST | `/Sync` | ✗ |
| GET | `/Sync/{id}` | ✗ |
| DELETE | `/Sync/{id}` | ✗ |
| POST | `/Sync/{id}/Items` | ✗ |
| POST | `/Sync/{id}/Cancel` | ✗ |

### 4.3 Collection Endpoints

| Method | Path | Status |
|--------|------|--------|
| GET | `/Collections` | ✗ |
| POST | `/Collections` | ✗ |
| GET | `/Collections/{id}` | ✗ |
| PUT | `/Collections/{id}` | ✗ |
| DELETE | `/Collections/{id}` | ✗ |
| POST | `/Collections/{id}/Items` | ✗ |
| DELETE | `/Collections/{id}/Items` | ✗ |

### 4.4 UserViews Endpoints

| Method | Path | Status |
|--------|------|--------|
| GET | `/UserViews` | ✗ |
| POST | `/UserViews` | ✗ |
| GET | `/UserViews/{id}` | ✗ |
| DELETE | `/UserViews/{id}` | ✗ |

### 4.5 Playback Endpoints

| Method | Path | Status |
|--------|------|--------|
| GET | `/Playback/{id}/Stream` | ⚠️ | Partial |
| GET | `/Playback/{id}/Subtitles/{index}` | ⚠️ | Partial |
| GET | `/Playback/{id}/Chapters` | ⚠️ | Partial |
| POST | `/Playback/{id}/Progress` | ⚠️ | Partial |
| GET | `/Playback/{id}/AudioStreams` | ⚠️ | Partial |
| GET | `/Playback/{id}/SubtitleStreams` | ⚠️ | Partial |

---

## 5. API Models

### 5.1 Core Models

| C# Model | Go Struct | Status |
|----------|-----------|--------|
| BaseItemDto | `internal/model/item.go` | ✓ |
| BaseItemDtoQueryResult | `internal/model/item.go` | ✓ |
| SessionInfo | `internal/model/session.go` | ✓ |
| UserDto | `internal/model/user.go` | ✓ |
| UserPolicy | `internal/model/user.go` | ✓ |
| MediaSourceInfo | `internal/model/stream.go` | ✓ |
| TranscodingInfo | `internal/model/stream.go` | ✓ |

### 5.2 Request/Response Models

| C# Model | Go Struct | Status |
|----------|-----------|--------|
| AuthenticateRequest | `internal/model/user.go` | ✓ |
| AuthenticateResult | `internal/model/user.go` | ✓ |
| LibraryOptions | `internal/model/item.go` | ✓ |
| UpdateUserPolicy | `internal/model/user.go` | ✓ |
| ItemUpdateInfo | `internal/model/item.go` | ✓ |

---

## 6. API Authentication

### 6.1 Authentication Methods

| Method | Status | Implementation |
|--------|--------|----------------|
| API Key (X-Emby-Token) | ✓ | `internal/api/middleware/auth.go` |
| User/Password | ✓ | `internal/service/auth/` |
| Device ID | ✓ | `internal/service/device/` |
| Token refresh | ✓ | `internal/service/auth/` |

### 6.2 Middleware Chain

```go
// Current implementation
func MiddlewareChain(next http.Handler) http.Handler {
    return middleware.Recovery(
        middleware.Logger(
            middleware.CORS(
                middleware.Auth(next)
            )
        )
    )
}
```

---

## 7. Migration Tasks

### 7.1 Priority 1 (Critical)

| # | Task | Status | Files |
|---|------|--------|-------|
| 1.1 | Complete LiveTV API | ✗ | `internal/api/handlers/livetv.go` |
| 1.2 | Complete Sync API | ✗ | `internal/api/handlers/sync.go` |
| 1.3 | Complete Collection API | ✗ | `internal/api/handlers/collection.go` |
| 1.4 | Complete UserViews API | ✗ | `internal/api/handlers/userviews.go` |

### 7.2 Priority 2 (High)

| # | Task | Status | Files |
|---|------|--------|-------|
| 2.1 | Complete Playback API | ⚠️ | `internal/api/handlers/playback.go` |
| 2.2 | Add missing Games endpoints | ⚠️ | `internal/api/handlers/games.go` |
| 2.3 | Add missing Package endpoints | ⚠️ | `internal/api/handlers/package.go` |
| 2.4 | Add missing DisplayPrefs endpoints | ⚠️ | `internal/api/handlers/displayprefs.go` |

### 7.3 Priority 3 (Medium)

| # | Task | Status | Files |
|---|------|--------|-------|
| 3.1 | Add HdHomerun endpoints | ✗ | `internal/api/handlers/hdhomerun.go` |
| 3.2 | Add SchedulesDirect endpoints | ✗ | `internal/api/handlers/schedulesdirect.go` |
| 3.3 | Verify all query parameter handling | ⚠️ | All handlers |
| 3.4 | Add API versioning support | ✗ | `internal/api/router.go` |

---

## 8. Verification Checklist

### 8.1 API Compatibility

- [ ] All 150+ C# API endpoints mapped
- [ ] All request/response models implemented
- [ ] All authentication methods supported
- [ ] All query parameters handled
- [ ] All error codes returned correctly
- [ ] All media types (JSON, XML) supported

### 8.2 Endpoint Coverage

| Category | Total | Implemented | Percentage |
|----------|-------|-------------|------------|
| Library | 15 | 12 | 80% |
| Session | 10 | 10 | 100% |
| User | 12 | 12 | 100% |
| Media | 8 | 8 | 100% |
| System | 8 | 8 | 100% |
| LiveTV | 15 | 5 | 33% |
| Channels | 8 | 6 | 75% |
| Images | 8 | 8 | 100% |
| Playback | 5 | 2 | 40% |
| Notifications | 6 | 6 | 100% |
| Playlists | 6 | 6 | 100% |
| ScheduledTasks | 4 | 4 | 100% |
| **Total** | **~150** | **~120** | **~80%** |

---

## Appendix: Related Documents

- [Master Migration Plan](./000-migration-master-plan.md)
- [Discovery: API](./.discovery/340-mediabrowser-api.md)
- [Discovery: Services](./.discovery/343-mediabrowser-api-services.md)
- [Discovery: Controllers](./.discovery/344-mediabrowser-api-controllers.md)
- [Go Implementation](./.discovery/360-emby-go.md)

---

**Document Version:** 1.0  
**Last Updated:** 2026-05-04  
**Status:** In Progress - ~80% Complete
