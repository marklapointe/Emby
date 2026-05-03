# Component: MediaBrowser.Providers — People

**Path:** `MediaBrowser.Providers/People/`
**Type:** Directory | Sub-module
**Language:** C#
**Maps to:** `.discovery/325-mediabrowser-providers-people.md`
**Parent:** `.discovery/320-mediabrowser-providers.md`

## Description

People (actors, directors, etc.) metadata providers. Fetches person information from TheMovieDb and TheTVDB.

## Files

### Root People Files (4 files)

- `PersonMetadataService.cs` — MediaBrowser.Providers/People/PersonMetadataService.cs
- `MovieDbPersonProvider.cs` — MediaBrowser.Providers/People/MovieDbPersonProvider.cs
- `MovieDbPersonImageProvider.cs` — MediaBrowser.Providers/People/MovieDbPersonImageProvider.cs
- `TvdbPersonImageProvider.cs` — MediaBrowser.Providers/People/TvdbPersonImageProvider.cs

## Key Classes

| Class | File | Purpose |
|-------|------|---------|
| `PersonMetadataService` | `PersonMetadataService.cs` | Person metadata orchestration |
| `MovieDbPersonProvider` | `MovieDbPersonProvider.cs` | TMDB person API integration |
| `MovieDbPersonImageProvider` | `MovieDbPersonImageProvider.cs` | TMDB person images |
| `TvdbPersonImageProvider` | `TvdbPersonImageProvider.cs` | TVDB person images |

## External APIs

| Provider | API | Description |
|----------|-----|-------------|
| TheMovieDb | tmdb.org | Movie/TV database with person info |
| TheTVDB | thetvdb.com | TV series database with person info |

## Dependencies

- **MediaBrowser.Controller** — Base entity types
- **MediaBrowser.Model** — API models
- **HttpClient** — External API calls
