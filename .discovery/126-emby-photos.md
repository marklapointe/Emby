# Component: Emby.Photos

**Path:** `Emby.Photos/`
**Type:** Directory | Plugin
**Language:** C#
**Maps to:** `.discovery/126-emby-photos.md`

## Description

Emby.Photos is a metadata provider plugin for photo/image items. It extracts EXIF metadata (camera model, date taken, GPS coordinates, orientation) from image files and provides it to the Emby library system.

## Structure

```
Emby.Photos/
├── Emby.Photos.csproj
├── PhotoProvider.cs               # Photo metadata provider
│   └── [class] PhotoProvider : ICustomMetadataProvider<Photo>, IForcedProvider, IHasItemChangeMonitor
│       ├── [constructor] PhotoProvider(ILogger, IFileSystem)
│       ├── [method] public Task<MetadataResult<Photo>> GetMetadata(ItemLookupInfo info, CancellationToken cancellationToken)
│       │   ├── Opens image file via IFileSystem
│       │   ├── Reads EXIF metadata (if available)
│       │   ├── Extracts:
│       │   │   ├── Camera model
│       │   │   ├── Date taken (DateTimeOriginal)
│       │   │   ├── GPS coordinates (latitude/longitude)
│       │   │   ├── Orientation (rotation)
│       │   │   ├── Aperture, shutter speed, ISO
│       │   │   └── Image dimensions
│       │   ├── Sets Photo.ProductionYear from date taken
│       │   ├── Sets Photo.PremiereDate from date taken
│       │   └── Returns populated Photo item
│       ├── [method] public bool HasChanged(BaseItem item, IDirectoryService directoryService)
│       │   └── Returns true if file modification time changed
│       └── [method] public string Name
│           └── Returns "Emby Photos"
└── Properties/
    └── AssemblyInfo.cs
```

## Data Flow

```mermaid
graph TD
    A[Library scan] --&gt; B[PhotoProvider.GetMetadata]
    B --&gt; C[Open image file]
    C --&gt; D[Read EXIF tags]
    D --&gt; E[Extract camera, date, GPS]
    E --&gt; F[Populate Photo item]
    F --&gt; G[Return to library manager]
```

## EXIF Tags Extracted

| Tag | Property | Description |
|-----|----------|-------------|
| DateTimeOriginal | PremiereDate | When photo was taken |
| Make + Model | Overview | Camera manufacturer/model |
| GPSLatitude/GPSLongitude | Overview | GPS coordinates |
| Orientation | Overview | Image rotation |
| FNumber | Overview | Aperture |
| ExposureTime | Overview | Shutter speed |
| ISOSpeedRatings | Overview | ISO sensitivity |

## Side Effects

- Reads image files via IFileSystem
- Parses EXIF metadata (read-only)
- No external network calls
