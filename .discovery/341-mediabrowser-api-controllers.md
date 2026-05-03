# Component: MediaBrowser.Api — Controllers

**Path:** `MediaBrowser.Api/`
**Type:** Directory | Module
**Language:** C#
**Maps to:** `.discovery/341-mediabrowser-api-controllers.md`
**Parent:** `.discovery/340-mediabrowser-api.md`

## Description

REST API controllers for Emby Server. Each controller handles a specific domain
of the API (movies, TV, music, users, etc.).

## Structure

```
MediaBrowser.Api/
├── Movies/
│   └── MoviesService.cs          # [class] MoviesService
├── TvShows/
│   └── TvShowsService.cs         # [class] TvShowsService
├── Music/
│   └── MusicService.cs           # [class] MusicService
├── UserLibrary/
│   └── *Service.cs               # User library services
├── Playback/
│   └── *Service.cs               # Playback services
├── Images/
│   └── ImageService.cs           # [class] ImageService
├── LiveTv/
│   └── LiveTvService.cs          # [class] LiveTvService
├── Channels/
│   └── ChannelService.cs         # [class] ChannelService
├── Configuration/
│   └── ConfigurationService.cs   # [class] ConfigurationService
├── Devices/
│   └── DeviceService.cs          # [class] DeviceService
├── Dlna/
│   └── DlnaService.cs            # [class] DlnaService
├── Notifications/
│   └── NotificationService.cs    # [class] NotificationService
├── Sessions/
│   └── SessionService.cs         # [class] SessionService
├── Subtitles/
│   └── SubtitleService.cs        # [class] SubtitleService
├── Sync/
│   └── SyncService.cs            # [class] SyncService
├── UserActivity/
│   └── UserActivityService.cs    # [class] UserActivityService
└── *Service.cs                   # Other API services
```

## Key Classes

| Class | File | Purpose |
|-------|------|---------|
| `MoviesService` | `Movies/MoviesService.cs` | Movie API endpoints |
| `TvShowsService` | `TvShows/TvShowsService.cs` | TV API endpoints |
| `MusicService` | `Music/MusicService.cs` | Music API endpoints |
| `ImageService` | `Images/ImageService.cs` | Image API endpoints |
| `LiveTvService` | `LiveTv/LiveTvService.cs` | LiveTV API endpoints |
| `SessionService` | `Sessions/SessionService.cs` | Session API endpoints |
