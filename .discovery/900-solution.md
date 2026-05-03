# Component: MediaBrowser.sln

**Path:** `MediaBrowser.sln`
**Type:** File | Solution Configuration
**Language:** Visual Studio Solution
**Maps to:** `.discovery/900-solution.md`

## Description

MediaBrowser.sln is the Visual Studio solution file that defines the Emby Server project structure. It contains 20 C# projects with multiple build configurations (Debug, Release, Release Mono, Signed) and platform targets (Any CPU, x86, x64).

## Structure

```
MediaBrowser.sln
├── Solution Items
│   └── Performance profiles (.psess)
├── .nuget configuration
├── Projects (20 total):
│   ├── MediaBrowser.Api
│   ├── MediaBrowser.WebDashboard
│   ├── MediaBrowser.Tests
│   ├── MediaBrowser.Providers
│   ├── MediaBrowser.ServerApplication
│   ├── MediaBrowser.XbmcMetadata
│   ├── MediaBrowser.LocalMetadata
│   ├── MediaBrowser.Server.Mono
│   ├── Emby.Drawing
│   ├── Emby.Photos
│   ├── DvdLib
│   ├── BDInfo
│   ├── Emby.Server.Implementations
│   ├── RSSDP
│   ├── Emby.Dlna
│   ├── Emby.Drawing.ImageMagick
│   ├── Mono.Nat
│   ├── SocketHttpListener
│   └── Emby.Drawing.Skia
│   └── Emby.Notifications
├── Configurations:
│   ├── Debug
│   ├── Release
│   ├── Release Mono
│   └── Signed
└── Platforms: Any CPU, x86, x64, Win32, Mixed Platforms
```

## Build Configurations

| Configuration | Purpose |
|-------------|---------|
| Debug | Development with symbols |
| Release | Production optimized |
| Release Mono | Mono/Linux optimized |
| Signed | Strong-named assemblies |

## Dependencies

- Visual Studio 2017+ or MSBuild
- .NET Framework 4.6.1+ or Mono

## Side Effects

- Defines project build order
- Specifies solution-wide properties

## Reference

- Visual Studio solution format version 12.00
