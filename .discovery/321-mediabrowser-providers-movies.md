# Component: MediaBrowser.Providers — Movies

**Path:** `MediaBrowser.Providers/Movies/`
**Type:** Directory | Sub-module
**Language:** C#
**Maps to:** `.discovery/321-mediabrowser-providers-movies.md`
**Parent:** `.discovery/320-mediabrowser-providers.md`

## Description

Metadata providers for movie content. Fetches movie information from external sources like TMDB, OMDB, Fanart.tv, and local metadata files.

## Files

| File | Purpose |
|------|---------|
| `FanartMovieImageProvider.cs` | Fanart.tv movie images |
| `GenericMovieDbInfo.cs` | Generic TMDB info helper |
| `MovieDbImageProvider.cs` | TMDB image provider |
| `MovieDbProvider.cs` | Base TMDB provider |
| `MovieDbSearch.cs` | TMDB search functionality |
| `MovieDbTrailerProvider.cs` | TMDB trailer provider |
| `MovieExternalIds.cs` | External ID mapping |
| `MovieMetadataService.cs` | Movie metadata orchestration |
| `TmdbSettings.cs` | TMDB configuration |

## Structure

```
Movies/
├── MovieMetadataService.cs       # [class] MovieMetadataService
│   └── Orchestrates movie metadata fetching
├── MovieDbProvider.cs            # [class] MovieDbProvider
│   └── Base TMDB provider
├── MovieDbSearch.cs              # [class] MovieDbSearch
│   └── TMDB search
├── MovieDbImageProvider.cs       # [class] MovieDbImageProvider
│   └── TMDB images
├── MovieDbTrailerProvider.cs     # [class] MovieDbTrailerProvider
│   └── TMDB trailers
├── FanartMovieImageProvider.cs   # [class] FanartMovieImageProvider
│   └── Fanart.tv images
├── MovieExternalIds.cs           # [class] MovieExternalIds
│   └── External ID mapping
├── GenericMovieDbInfo.cs        # [class] GenericMovieDbInfo
│   └── TMDB info helper
└── TmdbSettings.cs              # [class] TmdbSettings
    └── TMDB configuration
```

## Key Classes

| Class | File | Purpose |
|-------|------|---------|
| `MovieMetadataService` | `MovieMetadataService.cs` | Movie metadata orchestration |
| `MovieDbProvider` | `MovieDbProvider.cs` | TMDB API integration |
| `MovieDbSearch` | `MovieDbSearch.cs` | TMDB search |
| `FanartMovieImageProvider` | `FanartMovieImageProvider.cs` | Fanart.tv images |
