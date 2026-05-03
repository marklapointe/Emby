# Component: MediaBrowser.Server.Mono

**Path:** `MediaBrowser.Server.Mono/`
**Type:** Directory | Application
**Language:** C#
**Maps to:** `.discovery/210-mediabrowser-server-mono.md`

## Description

Mono runtime launcher for Emby Server. Provides platform-specific implementations for Linux/macOS environments.

## Files

### Root Files

- `ApplicationPathHelper.cs` — MediaBrowser.Server.Mono/ApplicationPathHelper.cs
- `ImageEncoderHelper.cs` — MediaBrowser.Server.Mono/ImageEncoderHelper.cs
- `ImageMagickSharp.dll.config` — MediaBrowser.Server.Mono/ImageMagickSharp.dll.config
- `MediaBrowser.Server.Mono.csproj` — MediaBrowser.Server.Mono/MediaBrowser.Server.Mono.csproj
- `MonoAppHost.cs` — MediaBrowser.Server.Mono/MonoAppHost.cs
- `Program.cs` — MediaBrowser.Server.Mono/Program.cs
- `Properties/AssemblyInfo.cs` — MediaBrowser.Server.Mono/Properties/AssemblyInfo.cs
- `SQLitePCLRaw.provider.sqlite3.dll.config` — MediaBrowser.Server.Mono/SQLitePCLRaw.provider.sqlite3.dll.config
- `SkiaSharp.dll.config` — MediaBrowser.Server.Mono/SkiaSharp.dll.config
- `app.config` — MediaBrowser.Server.Mono/app.config
- `packages.config` — MediaBrowser.Server.Mono/packages.config

### Native/

- `PowerManagement.cs` — MediaBrowser.Server.Mono/Native/PowerManagement.cs

## Architecture

```mermaid
graph TD
    A[Program.cs] --> B[MonoAppHost]
    B --> C[Emby.Server.Implementations]
    D[Native/PowerManagement] --> B
```

## Key Classes

| Class | Responsibility |
|-------|----------------|
| `MonoAppHost` | Mono-specific application host |
| `Program` | Entry point |
| `ApplicationPathHelper` | Path resolution |
| `ImageEncoderHelper` | Image encoder setup |
| `PowerManagement` | System power control |

## Decomposition

### Program.cs (Entry Point)

#### Imports
```csharp
using MediaBrowser.Server.Implementations;
using System;
using System.Threading.Tasks;
```

#### Classes
`Program` (internal static class)

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `Main(string[])` | `int` | Application entry point |

### MonoAppHost.cs (Mono Application Host)

#### Classes
`MonoAppHost` (public class : ApplicationHost)

#### Key Properties
| Property | Type | Description |
|----------|------|-------------|
| `IsLinux` | `bool` | Running on Linux |
| `IsMac` | `bool` | Running on macOS |

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `ConfigureUdpPorts()` | `void` | Configure UDP listeners |
| `GetImageEncoder()` | `IImageEncoder` | Get image encoder |

### ApplicationPathHelper.cs (Path Resolution)

#### Classes
`ApplicationPathHelper` (public static class)

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `GetBasePath()` | `string` | Get base path |
| `GetDataPath()` | `string` | Get data directory |
| `GetLogPath()` | `string` | Get log directory |

### ImageEncoderHelper.cs (Encoder Setup)

#### Classes
`ImageEncoderHelper` (public static class)

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `GetEncoder()` | `IImageEncoder` | Get configured encoder |
| `IsSkiaAvailable()` | `bool` | Check Skia availability |
| `IsImageMagickAvailable()` | `bool` | Check ImageMagick availability |

## Dependencies

- Emby.Server.Implementations
- MediaBrowser.Model
