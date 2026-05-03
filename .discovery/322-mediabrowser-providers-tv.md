# Component: MediaBrowser.Providers — TV

**Path:** `MediaBrowser.Providers/TV/`
**Type:** Directory | Sub-module
**Language:** C#
**Maps to:** `.discovery/322-mediabrowser-providers-tv.md`
**Parent:** `.discovery/320-mediabrowser-providers.md`

## Description

Metadata providers for TV series, seasons, and episodes.

## Structure

```
TV/
├── SeriesMetadataService.cs      # [class] SeriesMetadataService
├── EpisodeMetadataService.cs     # [class] EpisodeMetadataService
├── *SeriesProvider.cs            # Series providers
├── *EpisodeProvider.cs           # Episode providers
└── *Helper.cs                    # TV helpers
```

## Key Classes

| Class | File | Purpose |
|-------|------|---------|
| `SeriesMetadataService` | `SeriesMetadataService.cs` | Series metadata orchestration |
| `EpisodeMetadataService` | `EpisodeMetadataService.cs` | Episode metadata orchestration |
