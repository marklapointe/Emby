# Component: MediaBrowser.LocalMetadata

**Path:** `MediaBrowser.LocalMetadata/`
**Type:** Directory | Module
**Language:** C#
**Maps to:** `.discovery/210-mediabrowser-localmetadata.md`

## Description

MediaBrowser.LocalMetadata provides metadata extraction from local media files without external API calls. It reads embedded metadata from audio/video files (ID3, EXIF, etc.) and parses local artwork and subtitle files.

## Structure

```
MediaBrowser.LocalMetadata/
├── MediaBrowser.LocalMetadata.csproj
├── Images/                      # Local image metadata
├── Audio/                       # Audio file metadata
├── Videos/                      # Video file metadata
└── Properties/                  # Assembly info
```

## Key Classes

| Class | File | Purpose |
|-------|------|---------|
| `LocalImageProvider` | `Images/` | Reads local artwork |
| `LocalAudioMetadata` | `Audio/` | Reads audio tags |
| `LocalVideoMetadata` | `Videos/` | Reads video metadata |

## Dependencies

- `MediaBrowser.Controller` — Metadata interfaces
- `ThirdParty/taglib` — Audio tag reading

## Side Effects

- Reads media file headers
- Extracts embedded images

## Reference

- `ILocalMetadataProvider` interface in `MediaBrowser.Controller`
