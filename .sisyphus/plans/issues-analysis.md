# Emby-Go ORM Migration Plan

## Background

The current implementation uses raw SQL strings for database operations, which has proven problematic especially with `:memory:` SQLite databases and multi-statement SQL execution. All database interactions should go through GORM ORM.

## Completed Work

### 1. Media Handler Streaming (HIGH)
**File**: `emby-go/internal/api/handlers/media.go`
- **Issue**: `GetStream` and `GetAudioStream` returned hardcoded "stream_data" and "audio_data"
- **Fix**: Now use `http.ServeFile` to serve actual media files
- **Status**: ✅ Fixed

### 2. Transcoding Handler (HIGH)
**File**: `emby-go/internal/api/handlers/transcoding.go`
- **Issue**: `GetTranscodingProfiles` and `GetActiveTranscodes` returned empty data
- **Fix**: Now call `h.transcodingSvc.GetTranscodingProfiles()` and `h.transcodingSvc.GetActiveStreamCount()`
- **Status**: ✅ Fixed

### 3. Dead Code Removal (LOW)
**File**: `emby-go/internal/api/middleware/auth.go`
- **Issue**: Unused `AuthMiddleware`, `NewAuthMiddleware`, `Handle`, `isPublicEndpoint`
- **Fix**: Removed unused code, kept only context key constants
- **Status**: ✅ Fixed

### 4. GORM Dependencies Added
- Added `gorm.io/gorm` v1.31.1
- Added `gorm.io/driver/sqlite` v1.6.0
- Added `github.com/jinzhu/inflection` v1.0.0
- Added `github.com/jinzhu/now` v1.1.5

## ORM Migration Tasks

### Phase 1: Create GORM Models
- [ ] Create `internal/model/gorm_models.go` with GORM-annotated models:
  - `Item` - GORM model for Items table
  - `MediaSource` - GORM model for MediaSources table
  - `User` - GORM model for Users table
  - `UserItem` - GORM model for UserItems table
  - `Session` - GORM model for Sessions table

### Phase 2: Refactor Database Manager
- [ ] Update `internal/database/database.go`:
  - Change from `*sql.DB` to `*gorm.DB`
  - Update `NewManager` to use GORM driver
  - Update `Close` to close GORM session
  - Add `GetDB()` method returning `*gorm.DB`

### Phase 3: Refactor Base Repository
- [ ] Update `internal/repository/base.go`:
  - Change from `*sql.DB` to `*gorm.DB`
  - Replace raw SQL methods with GORM methods
  - Update `WithTransaction` to use GORM transactions

### Phase 4: Migrate Item Repository
- [ ] Update `internal/repository/item.go`:
  - Replace raw SQL in `CreateSchema` with `db.AutoMigrate()`
  - Replace raw SQL queries with GORM methods
  - Update `GetAllItems`, `GetItemsByParent`, `InsertItem`, etc.

### Phase 5: Migrate User Repository
- [ ] Update `internal/repository/user.go`:
  - Replace raw SQL with GORM methods
  - Update authentication queries to use GORM

### Phase 6: Verify & Test
- [ ] Run `go build ./...` to verify compilation
- [ ] Run `go test ./...` to verify all tests pass
- [ ] Test with `:memory:` database specifically

## Database Schema (for reference)

```
Items (
  Id TEXT PRIMARY KEY,
  Name TEXT NOT NULL,
  Overview TEXT,
  Tagline TEXT,
  IndexNumber INTEGER,
  ParentIndex INTEGER,
  CommunityRating REAL,
  RunTimeTicks INTEGER,
  ProductionYear INTEGER,
  OfficialRating TEXT,
  ContentType TEXT,
  MediaType TEXT,
  Genres TEXT,
  Studios TEXT,
  SeasonNumber INTEGER,
  EpisodeNumber INTEGER,
  Album TEXT,
  Artists TEXT,
  ExtraType TEXT,
  ChannelNumber INTEGER,
  StartDate TEXT,
  EndDate TEXT,
  IsLive INTEGER,
  IsSeries INTEGER,
  IsMovie INTEGER,
  IsNews INTEGER,
  IsSports INTEGER,
  IsKids INTEGER,
  IsPremiere INTEGER,
  LocationType TEXT,
  Path TEXT,
  PrimaryImageURL TEXT,
  BackdropImageURL TEXT,
  ParentID TEXT,
  Width INTEGER,
  Height INTEGER,
  Video3DFormat TEXT,
  PostLiveFeedTime INTEGER,
  LiveMediaSourceID TEXT,
  StartTimeTicks INTEGER,
  EndTimeTicks INTEGER,
  RemoteImageURL TEXT,
  LocalTrailerCount INTEGER,
  LockedFields TEXT,
  LockData INTEGER,
  Disabled INTEGER,
  EnableMediaSourceDisplay INTEGER,
  ExtraIds TEXT,
  CreatedDate TEXT,
  ModifiedDate TEXT
)

MediaSources (
  Id TEXT PRIMARY KEY,
  ItemId TEXT NOT NULL REFERENCES Items(Id),
  Name TEXT,
  Type TEXT,
  Container TEXT,
  Size INTEGER,
  Path TEXT,
  Protocol TEXT,
  Encoder INTEGER,
  VideoCodec TEXT,
  AudioCodec TEXT,
  Format TEXT,
  Width INTEGER,
  Height INTEGER,
  RefFrames INTEGER,
  VideoFramerate TEXT,
  VideoBitRate INTEGER,
  AudioBitRate INTEGER,
  AudioChannels INTEGER,
  AudioSampleRate TEXT,
  DefaultAudioStreamIndex INTEGER,
  SupportsTranscoding INTEGER,
  SupportsDirectStream INTEGER,
  SupportsDirectPlay INTEGER,
  IsRemote INTEGER
)

Users (
  Id TEXT PRIMARY KEY,
  Name TEXT NOT NULL,
  Username TEXT,
  EmailAddress TEXT,
  LoginUsername TEXT,
  LoginPassword TEXT,
  InvalidLoginAttemptCount INTEGER,
  LastLoginDate TEXT,
  LastActivityDate TEXT,
  AuthenticationProviderID TEXT,
  PrimaryImageTag TEXT,
  Policy TEXT
)

UserItems (
  Id TEXT PRIMARY KEY,
  UserId TEXT NOT NULL REFERENCES Users(Id),
  ItemID TEXT NOT NULL REFERENCES Items(Id),
  PlaybackPositionTicks INTEGER,
  PlayCount INTEGER,
  IsFavorite INTEGER,
  Liked INTEGER,
  LastPlayedDate TEXT,
  Played INTEGER,
  Rating INTEGER
)

Sessions (
  Id TEXT PRIMARY KEY,
  Client TEXT,
  DeviceName TEXT,
  DisplayName TEXT,
  Endpoint TEXT,
  LocalAddress TEXT,
  RemoteAddress TEXT,
  MachineId TEXT,
  LastActivityTime TEXT,
  LastPlaybackTime TEXT,
  PlaybackPositionTicks INTEGER,
  PlayMethod TEXT,
  SupportsMediaControl INTEGER,
  SupportsPersistentIdentification INTEGER,
  SupportsSync INTEGER,
  IsInActiveSession INTEGER,
  IsTerminal INTEGER,
  StartTimeTicks INTEGER
)
```

## Notes

- GORM AutoMigrate handles table creation idempotently
- GORM uses `gorm.Model` which adds ID, CreatedAt, UpdatedAt, DeletedAt fields
- For SQLite with GORM, use `gorm.Config` with `DisableForeignKeyConstraintWhenMigrating` when needed
- The existing `model/item.go`, `model/user.go`, `model/session.go` contain the API-facing structs - GORM models may need separate structs or can reuse these with GORM tags