# Component: Emby.Photos

**Path:** `Emby.Photos/`
**Type:** Directory | Module
**Language:** C#
**Maps to:** `.discovery/150-emby-photos.md`

## Description

Emby.Photos provides photo album and image library management. It handles photo metadata extraction (EXIF), album organization, and image-specific features like slideshows and photo viewing.

## Structure

```
Emby.Photos/
├── Emby.Photos.csproj           # Project file
├── PhotoProvider.cs             # Photo metadata provider
└── Properties/                  # Assembly info
```

## Key Classes

| Class | File | Purpose |
|-------|------|---------|
| `PhotoProvider` | `PhotoProvider.cs` | Extracts photo metadata |

## Dependencies

- `MediaBrowser.Controller` — Media interfaces
- `MediaBrowser.Model` — Photo types

## Side Effects

- Reads EXIF data from image files
- Generates photo thumbnails

## Reference

- `IImageProvider` interface in `MediaBrowser.Controller`
