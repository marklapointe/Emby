# Component: Emby.Server.Implementations — EnvironmentInfo

**Path:** `Emby.Server.Implementations/EnvironmentInfo/`
**Type:** Directory | Module
**Language:** C#
**Maps to:** `.discovery/204-emby-server-impl-environmentinfo.md`

## Description

Provides system and environment information for the server. Reports OS details, version info, and runtime environment.

## Files

- `EnvironmentInfo.cs` — Emby.Server.Implementations/EnvironmentInfo/EnvironmentInfo.cs

## Decomposition

### EnvironmentInfo.cs (Environment Information)

#### Imports
```csharp
using MediaBrowser.Model.System;
using System;
using System.Runtime.InteropServices;
using SystemInfo = MediaBrowser.Model.System.SystemInfo;
```

#### Classes
`EnvironmentInfo` (public class : IEnvironmentInfo)

#### Key Properties
| Property | Type | Description |
|----------|------|-------------|
| `OperatingSystem` | `string` | OS name |
| `OperatingSystemVersion` | `string` | OS version |
| `RuntimeIdentifier` | `string` | .NET runtime |
| `HasLowGraphicsMemory` | `bool` | GPU memory status |

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `GetSystemInfo()` | `SystemInfo` | Get full system info |
| `GetEnvironmentVariable(string)` | `string` | Get env var |
| `GetPath(string)` | `string` | Get system path |

## Data Flow

```mermaid
graph LR
    A[System] --> B[EnvironmentInfo]
    B --> C[SystemInfo Model]
    C --> D[API Response]
```

## Dependencies

- `MediaBrowser.Model.System` — System info models
- `System.Runtime.InteropServices` — Platform interop

## Statistics

| Metric | Value |
|--------|-------|
| Files | 1 |
| Classes | 1 |
| LOC | ~80 |
