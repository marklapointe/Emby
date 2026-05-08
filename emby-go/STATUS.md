# Emby Go Server - Current Status

## Last Updated
Thu May 07 2026 (evening)

## Goal
Full API compatibility between Go server and Emby web client, including the startup wizard flow.

## Constraints & Preferences
- Both lowercase and PascalCase routes must work
- Match Emby server version format (4.8.1.0)
- User wants proactive implementation of all endpoints rather than waiting for wizard to reveal them one-by-one

## What Was Just Completed

### Bug Fix - isAjaxRequest in static.go
- Fixed `isAjaxRequest` to only return true for actual API requests (application/json), not generic requests
- Previously: `Accept: */*` (curl default) was treated as AJAX, preventing wizard redirect
- Now: Only `Accept: application/json` or `X-Requested-With: XMLHttpRequest` triggers AJAX behavior
- **Verification**: `/web/index.html` now returns 302 redirect to wizard when wizard not completed

### Bug Fix - Artists and AlbumArtists Endpoints
- Added `GetArtists` handler in `filter.go` - handles `/Artists` endpoint
- Added `GetAlbumArtists` handler in `filter.go` - handles `/AlbumArtists` endpoint
- Registered routes in `router.go` via `registerFilterRoutes`
- Added special case mappings in `toTitleCase`:
  - `artists` → `Artists`
  - `albumartists` → `AlbumArtists`
- **Verification**: Both endpoints return `[]` (empty) instead of 404

### Bug Fix - toTitleCase Compound Words
- Fixed `toTitleCase` in `router.go` - PascalCase routing now works for compound words
- Added missing special case mappings:
  - `mediafolders` → `MediaFolders`
  - `virtualfolders` → `VirtualFolders`
  - `availableoptions` → `AvailableOptions`
  - `defaultdirectorybrowser` → `DefaultDirectoryBrowser`
  - `directorycontents` → `DirectoryContents`
- **Issue**: `/Library/MediaFolders` returned 404 because "mediafolders" → "Mediafolders" instead of "MediaFolders"
- **Verification**: All `TestToTitleCase_*` tests pass, server responds correctly to PascalCase routes

### New Services Created
- `internal/service/dlna/dlna.go` - DLNA service (stub)
- `internal/service/sync/sync.go` - Sync service (stub)

### New Handlers Created
- `internal/api/handlers/dlna.go` - DLNA endpoints (/Dlna/Profiles, /Dlna/ProfileInfos, /Dlna/Profiles/Default)
- `internal/api/handlers/sync.go` - Sync endpoints (/Sync/Jobs, /Sync/Jobs/{id}, /Sync/Jobs/{id}/Items/{itemId})
- `internal/api/handlers/plugin.go` - Plugin endpoints (/Plugins, /Plugins/{id}, /Plugins/{id}/Configuration, /Plugins/SecurityInfo, /Plugins/Released)
- `internal/api/handlers/collection.go` - Collection endpoints (/Collections, /Collections/{id}/Items)
- `internal/api/handlers/auth.go` - Auth endpoints (/Auth/Providers)

### Extended Handlers
- `livetv.go` - Added: GetRecommendedPrograms, GetSeriesTimers, GetTimerProviders, GetTunerHosts, GetTunerHost, CreateTunerHost, DeleteTunerHost, GetTunerHostTypes, GetListingProviders, CreateListingProvider, GetDefaultListingProvider, GetSchedulesDirectCountries, CreateChannelMapping, GetChannelMappingOptions
- `filter.go` - Added: GetCultures, GetCountries, GetMusicGenres, GetArtists, GetAlbumArtists

### Router Updates
- Added DLNA, Sync, Plugin, Collection, and Auth route registration in `router.go`
- Fixed `toTitleCase` to handle Emby-specific PascalCase paths (e.g., LiveTv, RecommendedPrograms, SchedulesDirect, etc.)
- Added special case mappings for: livetv, dlna, scheduledtasks, recommendedprograms, seriestimers, timerproviders, tunerhosts, listingproviders, schedulesdirect, channelmappings, channelmappingoptions, profileinfos, musicgenres

### Tests Created/Updated
- `internal/service/dlna/dlna_test.go`
- `internal/service/sync/sync_test.go`
- `internal/api/handlers/dlna_test.go`
- `internal/api/handlers/sync_test.go`
- `internal/api/handlers/plugin_test.go`
- `internal/api/handlers/auth_test.go`
- `internal/api/handlers/collection_test.go`
- `internal/api/handlers/filter_test.go`
- `internal/api/handlers/livetv_test.go`
- `internal/api/router_test.go`
- `tests/integration/integration_test.go` - Extended with more endpoint tests

### Build Status
- All tests pass
- Build succeeds
- `go vet` shows pre-existing issues (not from new code)

## Context From Previous Sessions

### Key Fixes Applied Earlier
1. Removed `EnsureDefaultUser()` call from `GetPublicUsers` - wizard shows user now
2. Fixed `isCoreHtml` in `static.go` - handles empty path after `http.StripPrefix`
3. Added `application/x-www-form-urlencoded` support to chi middleware
4. Added form-encoded parsing to `PostStartupConfig` and `PostUser` handlers
5. Fixed `toTitleCase` - properly TitleCases ALL segments (not just after first)
6. Added explicit route `router.Get("/Library/VirtualFolders", ...)` for PascalCase
7. Added `GetAvailableOptions` handler for `/Libraries/AvailableOptions`
8. Added `GetDefaultDirectoryBrowser` for `/Environment/DefaultDirectoryBrowser`
9. Added `GetDirectoryContents` for `/Environment/DirectoryContents`
10. Fixed `SystemHandler` - reads `StartupWizardCompleted` from database config, not hardcoded `false`

### Critical Context
- Found 91 `ApiClient` call patterns in web client - most have been addressed
- User preference: "reducing later is better" - implement all now vs waiting for errors
- External clients may cause additional issues but will address as discovered

## Relevant Files
- `emby-go/internal/server/static.go`: isCoreHtml, isAjaxRequest, wizard redirect logic
- `emby-go/internal/api/handlers/startup.go`: PostStartupConfig, PostUser with form-encoded support
- `emby-go/internal/api/handlers/system.go`: GetPublicSystemInfo - reads from DB
- `emby-go/internal/api/handlers/environment.go`: DirectoryContents, ParentPath handlers
- `emby-go/internal/api/handlers/library.go`: GetVirtualFolders, GetAvailableOptions
- `emby-go/internal/api/router.go`: Route registration, toTitleCase fix
- `emby-go/cmd/emby-server/main.go`: AllowContentType middleware

## Next Steps
1. Test the wizard flow end-to-end with a browser (web client loads but can't fully test without browser)
2. Address any additional API compatibility issues discovered during browser testing
3. Consider running the emby-server binary and testing with actual web client

## Verified Working Endpoints
- `/emby/System/Info/Public` - Returns correct wizard state
- `/emby/Startup/First` - Returns `IsFirstRun: true` correctly
- `/emby/Library/MediaFolders` - Now works with PascalCase (FIXED)
- `/emby/Library/VirtualFolders` - Works with PascalCase
- `/emby/Startup/Configuration` - Returns culture/country settings
- `/emby/Localization/Options` - Returns language list
- `/emby/Branding/Configuration` - Returns branding settings
- `/emby/Artists` - Returns `[]` (Artists endpoint implemented)
- `/emby/AlbumArtists` - Returns `[]` (AlbumArtists endpoint implemented)
- `/web/index.html` - Returns 302 redirect to wizard (FIXED)

## How to Test
```bash
cd /home/mlapointe/git/Emby/emby-go
go build -o emby-server cmd/emby-server/main.go
./emby-server &
# Test endpoints
curl http://localhost:8096/emby/System/Info/Public
curl http://localhost:8096/emby/Startup/First
curl http://localhost:8096/emby/LiveTv/Channels
```
