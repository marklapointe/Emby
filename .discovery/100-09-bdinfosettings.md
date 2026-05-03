# Component: BDInfoSettings.cs

**Path:** `BDInfo/BDInfoSettings.cs`
**Type:** File | Static Settings
**Language:** C#
**Maps to:** `.discovery/100-09-bdinfosettings.md`
**Parent:** `.discovery/100-bdinfo.md`

## Structure

```
BDInfoSettings.cs
├── [namespace] BDInfo
│   └── [class] public static class BDInfoSettings
│       ├── [field] public static bool EnableSSIF = true
│       │   └── Enables Stereoscopic Interleaved File processing
│       ├── [field] public static bool EnableFilterLoopingPlaylists = true
│       │   └── Filters out playlists that loop (e.g., menu backgrounds)
│       ├── [field] public static bool EnableFilterShortPlaylists = true
│       │   └── Filters out playlists shorter than threshold
│       ├── [field] public static int FilterShortPlaylistsValue = 20
│       │   └── Minimum playlist duration in seconds (default: 20s)
│       ├── [field] public static bool KeepStreamCount = true
│       │   └── Maintains consistent stream count across clips
│       ├── [field] public static bool ShowSourceStreamInfo = false
│       │   └── Shows raw stream info from source files
│       └── [field] public static bool GenerateStreamDiagnostics = false
│           └── Generates detailed stream diagnostic output
```

## Description

`BDInfoSettings` is a static configuration class controlling BDInfo behavior. Settings include SSIF (3D interleaved) support, playlist filtering (looping/short playlists), and diagnostic output options. All fields are public static with sensible defaults.
