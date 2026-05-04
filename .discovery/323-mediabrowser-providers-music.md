# Component: MediaBrowser.Providers — Music

**Path:** `MediaBrowser.Providers/Music/`
**Type:** Directory | Sub-module
**Language:** C#
**Maps to:** `.discovery/323-mediabrowser-providers-music.md`
**Parent:** `.discovery/320-mediabrowser-providers.md`

## Description

Metadata providers for music albums, artists, and tracks. Supports external metadata from MusicBrainz and fanart.tv.

## Files

### Root Music Files (18 files)

- `AlbumMetadataService.cs` — MediaBrowser.Providers/Music/AlbumMetadataService.cs
- `ArtistMetadataService.cs` — MediaBrowser.Providers/Music/ArtistMetadataService.cs
- `AudioDbAlbumImageProvider.cs` — MediaBrowser.Providers/Music/AudioDbAlbumImageProvider.cs
- `AudioDbAlbumProvider.cs` — MediaBrowser.Providers/Music/AudioDbAlbumProvider.cs
- `AudioDbArtistImageProvider.cs` — MediaBrowser.Providers/Music/AudioDbArtistImageProvider.cs
- `AudioDbArtistProvider.cs` — MediaBrowser.Providers/Music/AudioDbArtistProvider.cs
- `AudioDbExternalIds.cs` — MediaBrowser.Providers/Music/AudioDbExternalIds.cs
- `AudioMetadataService.cs` — MediaBrowser.Providers/Music/AudioMetadataService.cs
- `FanartAlbumProvider.cs` — MediaBrowser.Providers/Music/FanartMusic/FanartAlbumProvider.cs
- `FanartArtistProvider.cs` — MediaBrowser.Providers/Music/FanartMusic/FanartArtistProvider.cs
- `FanArtAlbumProvider.cs` — MediaBrowser.Providers/Music/FanArtAlbumProvider.cs
- `FanArtArtistProvider.cs` — MediaBrowser.Providers/Music/FanArtArtistProvider.cs
- `FanartMusicHelper.cs` — MediaBrowser.Providers/Music/FanartMusic/FanartMusicHelper.cs
- `FanartMusicImageProvider.cs` — MediaBrowser.Providers/Music/FanartMusic/FanartMusicImageProvider.cs
- `MovieDbMusicVideoProvider.cs` — MediaBrowser.Providers/Music/MovieDbMusicVideoProvider.cs
- `MusicAlbumImageProvider.cs` — MediaBrowser.Providers/Music/MusicAlbumImageProvider.cs
- `MusicArtistImageProvider.cs` — MediaBrowser.Providers/Music/MusicArtistImageProvider.cs
- `MusicAlbumProvider.cs` — MediaBrowser.Providers/Music/MusicAlbumProvider.cs
- `MusicArtistProvider.cs` — MediaBrowser.Providers/Music/MusicArtistProvider.cs
- `MusicExternalIds.cs` — MediaBrowser.Providers/Music/MusicExternalIds.cs
- `MusicFanartProvider.cs` — MediaBrowser.Providers/Music/MusicFanartProvider.cs
- `MusicMetadataSearchExecutor.cs` — MediaBrowser.Providers/Music/MusicMetadataSearchExecutor.cs
- `MusicVideoMetadataService.cs` — MediaBrowser.Providers/Music/MusicVideoMetadataService.cs
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
