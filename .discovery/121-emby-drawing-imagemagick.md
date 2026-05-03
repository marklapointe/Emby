# Component: Emby.Drawing.ImageMagick

**Path:** \`Emby.Drawing.ImageMagick/\`
**Type:** Module
**Maps to:** \`.discovery/121-emby-drawing-imagemagick.md\`

## Description

ImageMagick-based image processing backend for Emby Server.

## Structure

```
Emby.Drawing.ImageMagick/
├── Emby.Drawing.ImageMagick.csproj
├── ImageMagickEncoder.cs     # Image encoding
├── ImageMagickImage.cs       # Image wrapper
├── ImageMagickImageProcessor.cs # Image processing
├── ImageMagickSharp.dll.config # Native library config
├── ImageMagickSharp.dll      # Native wrapper
├── ImageMagickSharp.dll.mdb  # Debug symbols
├── ImageMagickSharp.XML       # XML documentation
└── packages.config           # NuGet dependencies
```

## Decomposition

### ImageMagickEncoder.cs (Main Encoder)

#### Imports
```csharp
using ImageMagick;
using MediaBrowser.Controller.Drawing;
using MediaBrowser.Model.Drawing;
using System;
using System.IO;
using System.Threading.Tasks;
```

#### Classes
`ImageMagickEncoder` (public class : IImageEncoder)

#### Key Properties
| Property | Type | Description |
|----------|------|-------------|
| `Name` | `string` | "ImageMagick" |
| `SupportedInputFormats` | `string[]` | ImageMagick formats |

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `EncodeImage(ImageInfo, string, bool)` | `Task<string>` | Encode to file |
| `EncodeImage(ImageInfo, Stream, bool)` | `Task` | Encode to stream |
| `ExtractVideoFrame(string, TimeSpan)` | `Task<string>` | Extract frame |

### ImageMagickImageProcessor.cs (Image Processing)

#### Classes
`ImageMagickImageProcessor` (public class)

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `ProcessImage(ImageProcessingOptions)` | `Task<Tuple<string,string,DateTime>>` | Process and cache |

## Dependencies

- ImageMagickSharp (NuGet)
- Emby.Drawing (base interfaces)
