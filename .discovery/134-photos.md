# Component: Emby.Photos

**Path:** `Emby.Photos/`
**Type:** Directory | Plugin
**Language:** C#
**Maps to:** `.discovery/134-photos.md`

## Description

Emby.Photos is a lightweight plugin that provides metadata extraction for photo/image files. It reads EXIF data (date taken, camera model, GPS coordinates, orientation) and IPTC metadata (keywords, description) from image files to populate Photo item properties.

## Structure

```
Emby.Photos/
├── Emby.Photos.csproj
├── Properties/
│   └── AssemblyInfo.cs            # Assembly metadata
└── PhotoProvider.cs               # Photo metadata provider
    └── [class] PhotoProvider : ICustomMetadataProvider<Photo>, IForcedProvider, IHasItemChangeMonitor
        ├── [method] public Task<MetadataResult<Photo>> GetMetadata(ItemLookupInfo info, IDirectoryService directoryService, CancellationToken cancellationToken)
        │   ├── Opens image file via IFileSystem
        │   ├── Reads EXIF metadata (if available)
        │   │   ├── DateTaken → Photo.DateCreated
        │   │   ├── CameraModel → Photo.CameraModel
        │   │   ├── Orientation → Photo.Orientation
        │   │   ├── GPSLatitude → Photo.Latitude
        │   │   ├── GPSLongitude → Photo.Longitude
        │   │   └── GPSAltitude → Photo.Altitude
        │   ├── Reads IPTC metadata (if available)
        │   │   ├── Keywords → Photo.Tags
        │   │   └── Description → Photo.Overview
        │   ├── Sets Photo.Width and Photo.Height from image dimensions
        │   └── Returns populated MetadataResult<Photo>
        ├── [method] public bool HasChanged(BaseItem item, IDirectoryService directoryService)
        │   └── Returns true if image file modification time changed
        └── [property] public string Name
            └── Returns "Photos" (provider name)
```

## Supported Image Formats

| Format | Extension | EXIF | IPTC |
|--------|-----------|------|------|
| JPEG | `.jpg`, `.jpeg` | ✅ | ✅ |
| TIFF | `.tiff`, `.tif` | ✅ | ✅ |
| PNG | `.png` | ⚠️ Limited | ❌ |
| BMP | `.bmp` | ❌ | ❌ |
| GIF | `.gif` | ❌ | ❌ |
| RAW | `.cr2`, `.nef`, `.arw` | ✅ | ✅ |

## EXIF Tags Mapped

| EXIF Tag | Photo Property | Description |
|----------|---------------|-------------|
| DateTimeOriginal | DateCreated | When photo was taken |
| Make + Model | CameraModel | Camera manufacturer/model |
| Orientation | Orientation | Rotation (1=normal, 6=90° CW, etc.) |
| GPSLatitude | Latitude | GPS latitude |
| GPSLongitude | Longitude | GPS longitude |
| GPSAltitude | Altitude | GPS altitude |
| ImageWidth | Width | Image width in pixels |
| ImageLength | Height | Image height in pixels |

## Data Flow

```mermaid
graph TD
    A[Library scan] --&gt; B[PhotoProvider.GetMetadata]
    B --&gt; C[Open image file]
    C --&gt; D{EXIF available?}
    D --&gt;|Yes| E[Read EXIF tags]
    E --&gt; F[Map to Photo properties]
    D --&gt;|No| G[Skip EXIF]
    H{IPTC available?} --&gt;|Yes| I[Read IPTC tags]
    I --&gt; J[Map to Photo properties]
    H --&gt;|No| K[Skip IPTC]
    F --&gt; L[Return MetadataResult]
    J --&gt; L
    G --&gt; L
    K --&gt; L
```

## Side Effects

- Reads image files via IFileSystem
- Parses EXIF/IPTC metadata from binary image data
- No external network calls
- No file writes
