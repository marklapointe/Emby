# Component: Emby.Server.Implementations — Updates

**Path:** `Emby.Server.Implementations/Updates/`
**Type:** Directory | Module
**Language:** C#
**Maps to:** `.discovery/218-emby-server-impl-updates.md`

## Description

Package and system update management. Handles Emby Server updates and plugin installations.

## Files

- `InstallationManager.cs` — Emby.Server.Implementations/Updates/InstallationManager.cs

## Decomposition

### InstallationManager.cs (Installation Manager)

#### Imports
```csharp
using MediaBrowser.Common.Updates;
using MediaBrowser.Model.Configuration;
using MediaBrowser.Model.Events;
using MediaBrowser.Model.Updates;
using System;
using System.Collections.Generic;
using System.IO;
using System.Threading;
using System.Threading.Tasks;
```

#### Classes
`InstallationManager` (public class : IInstallationManager)

#### Key Properties
| Property | Type | Description |
|----------|------|-------------|
| `Installations` | `IEnumerable<InstallationInfo>` | Active installations |

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `InstallPackage(PackageVersionInfo, CancellationToken)` | `Task` | Install package |
| `UninstallPackage(string)` | `Task` | Uninstall package |
| `GetInstalledPackages(bool)` | `IEnumerable<PackageVersionInfo>` | Get installed |
| `GetAvailablePackages()` | `Task<IEnumerable<PackageInfo>>` | Check for updates |
| `ReportInstallationProgress(PackageVersionInfo, double, string)` | `void` | Report progress |

#### Key Events
| Event | Description |
|-------|-------------|
| `PackageInstalling` | Package install started |
| `PackageInstallationFailed` | Install failed |
| `PackageInstalled` | Install completed |
| `PackageUninstalled` | Uninstall completed |

## Data Flow

```mermaid
graph LR
    A[Update Check] --> B[InstallationManager]
    B --> C[Package Repository]
    C --> D[Download]
    D --> E[Install]
    E --> F[File System]
```

## Dependencies

- `MediaBrowser.Model.Updates` — Update models
- `System.IO` — File operations
- `MediaBrowser.Common.Updates` — Update interfaces

## Statistics

| Metric | Value |
|--------|-------|
| Files | 1 |
| Classes | 1 |
| LOC | ~450 |
