# Component: MediaBrowser.Providers — Music

**Path:** `MediaBrowser.Providers/Music/`
**Type:** Directory | Sub-module
**Language:** C#
**Maps to:** `.discovery/323-mediabrowser-providers-music.md`
**Parent:** `.discovery/320-mediabrowser-providers.md`

## Description

Metadata providers for music albums, artists, and tracks. Supports external metadata from MusicBrainz and fanart.tv.

## Files

### Root Music Files (9 files)

- `AlbumMetadataService.cs` — MediaBrowser.Providers/Music/AlbumMetadataService.cs
- `ArtistMetadataService.cs` — MediaBrowser.Providers/Music/ArtistMetadataService.cs
- `MusicAlbumImageProvider.cs` — MediaBrowser.Providers/Music/MusicAlbumImageProvider.cs
- `MusicArtistImageProvider.cs` — MediaBrowser.Providers/Music/MusicArtistImageProvider.cs
- `MusicAlbumProvider.cs` — MediaBrowser.Providers/Music/MusicAlbumProvider.cs
- `MusicArtistProvider.cs` — MediaBrowser.Providers/Music/MusicArtistProvider.cs
- `MusicMetadataSearchExecutor.cs` — MediaBrowser.Providers/Music/MusicMetadataSearchExecutor.cs
- `MusicFanartProvider.cs` — MediaBrowser.Providers/Music/MusicFanartProvider.cs
- `MusicBrainzAlbumProvider.cs` — MediaBrowser.Providers/Music/MusicBrainzAlbumProvider.cs

### FanartMusic/ (4 files)

- `FanartAlbumProvider.cs` — MediaBrowser.Providers/Music/FanartMusic/FanartAlbumProvider.cs
- `FanartArtistProvider.cs` — MediaBrowser.Providers/Music/FanartMusic/FanartArtistProvider.cs
- `FanartMusicHelper.cs` — MediaBrowser.Providers/Music/FanartMusic/FanartMusicHelper.cs
- `FanartMusicImageProvider.cs` — MediaBrowser.Providers/Music/FanartMusic/FanartMusicImageProvider.cs

### MusicBrainz/ (3 files)

- `MusicBrainzAlbumProvider.cs` — MediaBrowser.Providers/Music/MusicBrainz/MusicBrainzAlbumProvider.cs
- `MusicBrainzArtistProvider.cs` — MediaBrowser.Providers/Music/MusicBrainz/MusicBrainzArtistProvider.cs
- `MusicBrainzSearchProvider.cs` — MediaBrowser.Providers/Music/MusicBrainz/MusicBrainzSearchProvider.cs

## Key Classes

| Class | File | Purpose |
|-------|------|---------|
| `AlbumMetadataService` | `AlbumMetadataService.cs` | Album metadata orchestration |
| `ArtistMetadataService` | `ArtistMetadataService.cs` | Artist metadata orchestration |
| `MusicAlbumProvider` | `MusicAlbumProvider.cs` | Album metadata provider |
| `MusicArtistProvider` | `MusicArtistProvider.cs` | Artist metadata provider |
| `MusicBrainzAlbumProvider` | `MusicBrainzAlbumProvider.cs` | MusicBrainz API integration |
| `MusicFanartProvider` | `MusicFanartProvider.cs` | Fanart.tv images |

## External APIs

| Provider | API | Description |
|----------|-----|-------------|
| MusicBrainz | musicbrainz.org | Open music metadata database |
| Fanart.tv | fanart.tv | Music artwork and images |

## Dependencies

- **MediaBrowser.Controller** — Base entity types
- **MediaBrowser.Model** — API models
- **HttpClient** — External API calls
