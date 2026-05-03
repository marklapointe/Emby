# Component: MediaBrowser.Providers — Movies

**Path:** `MediaBrowser.Providers/Movies/`
**Type:** Directory | Sub-module
**Language:** C#
**Maps to:** `.discovery/321-mediabrowser-providers-movies.md`
**Parent:** `.discovery/320-mediabrowser-providers.md`

## Description

Metadata providers for movie content. Fetches movie information from external
sources like TMDB, OMDB, and local metadata files.

## Structure

```
Movies/
├── MovieMetadataService.cs       # [class] MovieMetadataService
│   └── Orchestrates movie metadata fetching
├── *MovieProvider.cs             # Various movie providers
│   ├── TmdbMovieProvider.cs      # TMDB provider
│   ├── OmdbMovieProvider.cs      # OMDB provider
│   └── LocalMovieProvider.cs     # Local metadata
└── *Helper.cs                    # Provider helpers
```

## Key Classes

| Class | File | Purpose |
|-------|------|---------|
| `MovieMetadataService` | `MovieMetadataService.cs` | Movie metadata orchestration |
| `TmdbMovieProvider` | `TmdbMovieProvider.cs` | TMDB API integration |
| `OmdbMovieProvider` | `OmdbMovieProvider.cs` | OMDB API integration |
