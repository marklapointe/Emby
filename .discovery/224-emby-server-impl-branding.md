# Component: Emby.Server.Implementations — Branding

**Path:** `Emby.Server.Implementations/Branding/`
**Type:** Directory | Module
**Language:** C#
**Maps to:** `.discovery/224-emby-server-impl-branding.md`

## Description

Server branding and customization. Manages custom branding configuration including logos, welcome text, and theme settings.

## Files

- `BrandingConfigurationFactory.cs` — Emby.Server.Implementations/Branding/BrandingConfigurationFactory.cs

## Decomposition

### BrandingConfigurationFactory.cs (Branding Configuration Factory)

#### Imports
```csharp
using MediaBrowser.Model.Branding;
using System;
using System.IO;
```

#### Classes
`BrandingConfigurationFactory` (public static class)

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `GetBrandingOptions(IApplicationPaths)` | `BrandingOptions` | Get branding config |
| `GetConfiguration(IApplicationPaths)` | `BrandingConfiguration` | Get full config |

## Data Flow

```mermaid
graph LR
    A[App Paths] --> B[BrandingConfigurationFactory]
    B --> C[BrandingConfiguration]
    C --> D[Web UI]
```

## Dependencies

- `MediaBrowser.Model.Branding` — Branding models

## Statistics

| Metric | Value |
|--------|-------|
| Files | 1 |
| Classes | 1 |
| LOC | ~30 |
