# Component: Emby.Server.Implementations — Images

**Path:** `Emby.Server.Implementations/Images/`
**Type:** Directory | Module
**Language:** C#
**Maps to:** `.discovery/206-emby-server-impl-images.md`

## Description

Dynamic image generation and management. Provides base image providers for posters, backdrops, and thumbnails.

## Files

- `BaseDynamicImageProvider.cs` — Emby.Server.Implementations/Images/BaseDynamicImageProvider.cs

## Decomposition

### BaseDynamicImageProvider.cs (Base Dynamic Image Provider)

#### Imports
```csharp
using MediaBrowser.Controller.Drawing;
using MediaBrowser.Controller.Entities;
using MediaBrowser.Model.Drawing;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
```

#### Classes
`BaseDynamicImageProvider<T>` (public abstract class : IDynamicImageProvider) where T : BaseItem

#### Key Properties
| Property | Type | Description |
|----------|------|-------------|
| `Name` | `string` | Provider name |
| `SupportedImages` | `ImageType[]` | Supported image types |

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `GetImages(T, IEnumerable<ImageType>)` | `Task<IEnumerable<ItemImageInfo>>` | Get images |
| `CreateImage(Object, ItemImageInfo, IImageGenerator)` | `Task` | Create image |
| `GetHash(T, ImageType)` | `string` | Get image hash |

## Data Flow

```mermaid
graph LR
    A[BaseItem] --> B[BaseDynamicImageProvider]
    B --> C[ItemImageInfo]
    C --> D[ImageGenerator]
    D --> E[Image File]
```

## Dependencies

- `MediaBrowser.Controller.Drawing` — Drawing interfaces
- `MediaBrowser.Controller.Entities` — Base item types

## Statistics

| Metric | Value |
|--------|-------|
| Files | 1 |
| Classes | 1 |
| LOC | ~100 |
