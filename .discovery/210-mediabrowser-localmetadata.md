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

## Decomposition

### Images/LocalImageProvider.cs (Local Image Provider)

#### Imports
```csharp
using MediaBrowser.Controller.Entities;
using MediaBrowser.Controller.Providers;
using MediaBrowser.Model.Entities;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
```

#### Classes
`LocalImageProvider` (public class : ILocalImageProvider)

#### Key Properties
| Property | Type | Description |
|----------|------|-------------|
| `SupportedImages` | `IEnumerable<LocalImageProvider.ImageType>` | Image types supported |

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `GetImages(BaseItem, IDirectoryService)` | `Task<IEnumerable<LocalImage>>` | Get local images |
| `HasImage(BaseItem, ImageType)` | `bool` | Check for image |

### Audio/LocalAudioMetadata.cs (Audio Metadata Reader)

#### Classes
`LocalAudioMetadata` (public class)

#### Key Properties
| Property | Type | Description |
|----------|------|-------------|
| `Album` | `string` | Album name |
| `Artist` | `string` | Artist name |
| `Genres` | `IEnumerable<string>` | Genre tags |

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `ReadFromFile(string)` | `LocalAudioMetadata` | Read from file |
| `ReadFromTag(TagLib.Tag)` | `void` | Parse tag |

### Videos/LocalVideoMetadata.cs (Video Metadata Reader)

#### Classes
`LocalVideoMetadata` (public class)

#### Key Properties
| Property | Type | Description |
|----------|------|-------------|
| `Format` | `string` | Video format |
| `Duration` | `TimeSpan` | Video duration |
| `Resolution` | `string` | Video resolution |

## Reference

- `ILocalMetadataProvider` interface in `MediaBrowser.Controller`
