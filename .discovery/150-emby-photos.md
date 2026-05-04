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
└── Properties/
    └── AssemblyInfo.cs          # Assembly metadata
```

## Files

| File | Purpose |
|------|---------|
| `PhotoProvider.cs` | Photo metadata provider - extracts EXIF data |
| `Properties/AssemblyInfo.cs` | Assembly metadata |

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

## Decomposition

### PhotoProvider.cs (Photo Metadata Provider)

#### Imports
```csharp
using MediaBrowser.Controller.Entities;
using MediaBrowser.Controller.Providers;
using MediaBrowser.Model.Entities;
using MediaBrowser.Model.IO;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
```

#### Classes
`PhotoProvider` (public class : ILocalMetadataProvider)

#### Key Properties
| Property | Type | Description |
|----------|------|-------------|
| `SupportsEmbeddedMetadata` | `bool` | Supports EXIF extraction |
| `SupportedItemTypes` | `Type[]` | Photo item types |

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `GetMetadata(ItemMetadataOptions, BaseItem, IDirectoryMonitor, ILogger, CancellationToken)` | `Task<LocalMetadataResult<Photo>` | Extract photo metadata |
| `GetSaveFilePaths(BaseItem, IDirectoryService)` | `IEnumerable<string>` | Save file locations |

## Reference

- `IImageProvider` interface in `MediaBrowser.Controller`
