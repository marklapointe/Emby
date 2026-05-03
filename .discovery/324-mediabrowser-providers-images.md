# Component: MediaBrowser.Providers — Images

**Path:** `MediaBrowser.Providers/Images/`
**Type:** Directory | Sub-module
**Language:** C#
**Maps to:** `.discovery/324-mediabrowser-providers-images.md`
**Parent:** `.discovery/320-mediabrowser-providers.md`

## Description

Image metadata providers for fetching posters, backdrops, and other artwork
from external sources.

## Structure

```
Images/
├── *ImageProvider.cs             # Image providers
│   ├── TmdbImageProvider.cs      # TMDB images
│   ├── FanArtImageProvider.cs    # FanArt.tv images
│   └── *ImageProvider.cs         # Other image sources
└── *ImageHelper.cs               # Image helpers
```

## Key Classes

| Class | File | Purpose |
|-------|------|---------|
| `TmdbImageProvider` | `TmdbImageProvider.cs` | TMDB artwork |
| `FanArtImageProvider` | `FanArtImageProvider.cs` | FanArt.tv artwork |
