# Component: Emby.Server.Implementations — FFMpeg

**Path:** `Emby.Server.Implementations/FFMpeg/`
**Type:** Directory | Sub-module
**Language:** C#
**Maps to:** `.discovery/176-emby-server-impl-ffmpeg.md`
**Parent:** `.discovery/160-emby-server-impl.md`

## Description

FFMpeg integration for media transcoding, thumbnail extraction, and media analysis. This module manages FFMpeg processes, configures transcoding parameters, and handles media conversion.

## Files

| File | Purpose |
|------|---------|
| `FFMpegInfo.cs` | FFMpeg version and capabilities info |
| `FFMpegInstallInfo.cs` | FFMpeg installation status |
| `FFMpegLoader.cs` | FFMpeg discovery and loading |

## Structure

```
FFMpeg/
├── FFMpegManager.cs              # [class] FFMpegManager
│   ├── Manages FFMpeg processes
│   ├── Configures transcoding parameters
│   └── Handles media conversion
├── FFMpegEncoder.cs            # [class] FFMpegEncoder
│   └── Media encoding/transcoding
├── FFMpegImageExtractor.cs       # [class] FFMpegImageExtractor
│   └── Thumbnail/screenshot extraction
├── FFMpegInfo.cs                 # [class] FFMpegInfo
├── FFMpegInstallInfo.cs          # [class] FFMpegInstallInfo
└── FFMpegLoader.cs               # [class] FFMpegLoader
```

## Key Classes

| Class | File | Purpose |
|-------|------|---------|
| `FFMpegManager` | `FFMpegManager.cs` | FFMpeg process management |
| `FFMpegEncoder` | `FFMpegEncoder.cs` | Media transcoding |
| `FFMpegImageExtractor` | `FFMpegImageExtractor.cs` | Image extraction |
| `FFMpegInfo` | `FFMpegInfo.cs` | FFMpeg version/capabilities |
| `FFMpegInstallInfo` | `FFMpegInstallInfo.cs` | Installation status |
| `FFMpegLoader` | `FFMpegLoader.cs` | FFMpeg discovery |

## Decomposition

### FFMpegLoader.cs (FFMpeg Discovery)

#### Imports
```csharp
using MediaBrowser.Common.Configuration;
using MediaBrowser.Common.IO;
using MediaBrowser.Controller.Configuration;
using MediaBrowser.Model.IO;
using MediaBrowser.Model.Logging;
using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.IO;
using System.Linq;
using System.Threading.Tasks;
```

#### Classes
`FFMpegLoader` (public class)

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `GetFilesystemMetadata()` | `IEnumerable<string>` | Get FFMpeg executable paths |
| `GetProcessMetadata()` | `IEnumerable<string>` | Get FFMpeg process paths |

### FFMpegInfo.cs (FFMpeg Info)

#### Classes
`FFMpegInfo` (public class)

#### Key Properties
| Property | Type | Description |
|----------|------|-------------|
| `Version` | `string` | FFMpeg version |
| `Path` | `string` | FFMpeg executable path |

### FFMpegInstallInfo.cs (Install Info)

#### Classes
`FFMpegInstallInfo` (public class)

#### Key Properties
| Property | Type | Description |
|----------|------|-------------|
| `Path` | `string` | Installation path |
| `Version` | `string` | Version info |
