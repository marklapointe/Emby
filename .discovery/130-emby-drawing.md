# Component: Emby.Drawing

**Path:** `Emby.Drawing/`
**Type:** Directory | Module
**Language:** C#
**Maps to:** `.discovery/130-emby-drawing.md`

## Description

Emby.Drawing provides image processing abstractions and a common interface for image encoding/decoding. It defines the `IImageEncoder` interface and `ImageProcessor` class used by the server and other components. Actual image processing is delegated to backend implementations (Skia, ImageMagick, .NET Drawing).

## Structure

```
Emby.Drawing/
├── Emby.Drawing.csproj          # Project file
├── ImageProcessor.cs            # Main image processor → [class] ImageProcessor
│   ├── Processes image resize, crop, format conversion
│   ├── Caches processed images
│   └── Delegates to IImageEncoder backends
├── NullImageEncoder.cs          # No-op image encoder fallback
├── Common/                      # Shared drawing utilities
└── Properties/                  # Assembly info
```

## Key Classes

| Class | File | Purpose |
|-------|------|---------|
| `ImageProcessor` | `ImageProcessor.cs` | Main image processing coordinator |
| `NullImageEncoder` | `NullImageEncoder.cs` | Fallback no-op encoder |

## Data Flow

```mermaid
graph LR
    A[Image Request] --> B[ImageProcessor]
    B --> C[Cache Check]
    C --|miss| --> D[IImageEncoder]
    D --> E[Skia/ImageMagick/.NET]
    E --> F[Processed Image]
    F --> G[Cache Store]
    G --> H[Response]
    C --|hit| --> H
```

## Dependencies

- `MediaBrowser.Controller.Drawing` — Drawing controller interfaces
- `MediaBrowser.Model.Drawing` — Drawing model types

## Backend Implementations

| Backend | Project | Description |
|---------|---------|-------------|
| Skia | `Emby.Drawing.Skia` → `.discovery/133-emby-drawing-skia.md` | Cross-platform 2D graphics |
| ImageMagick | `Emby.Drawing.ImageMagick` → `.discovery/131-emby-drawing-imagemagick.md` | Advanced image processing |
| .NET | `Emby.Drawing.Net` → `.discovery/132-emby-drawing-net.md` | System.Drawing backend |

## Side Effects

- Reads source images from filesystem
- Writes processed images to cache directory
- Loads image encoder plugins dynamically

## Reference

- `IImageEncoder` interface defined in `MediaBrowser.Controller.Drawing`
