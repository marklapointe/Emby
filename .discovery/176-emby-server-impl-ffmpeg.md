# Component: Emby.Server.Implementations — FFMpeg

**Path:** `Emby.Server.Implementations/FFMpeg/`
**Type:** Directory | Sub-module
**Language:** C#
**Maps to:** `.discovery/176-emby-server-impl-ffmpeg.md`
**Parent:** `.discovery/160-emby-server-impl.md`

## Description

FFMpeg integration for media transcoding, thumbnail extraction, and media analysis.

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
└── *FFMpeg*.cs                   # Supporting classes
```

## Key Classes

| Class | File | Purpose |
|-------|------|---------|
| `FFMpegManager` | `FFMpegManager.cs` | FFMpeg process management |
| `FFMpegEncoder` | `FFMpegEncoder.cs` | Media transcoding |
| `FFMpegImageExtractor` | `FFMpegImageExtractor.cs` | Image extraction |
