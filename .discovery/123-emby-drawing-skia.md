# Component: Emby.Drawing.Skia

**Path:** \`Emby.Drawing.Skia/\`
**Type:** Module
**Maps to:** \`.discovery/123-emby-drawing-skia.md\`

## Description

SkiaSharp-based image processing backend for Emby Server.

## Structure

```
Emby.Drawing.Skia/
├── Emby.Drawing.Skia.csproj
├── ImageEncoder.cs           # Image encoding
├── ImageProcessor.cs         # Image processing
├── Image.cs                  # Image wrapper
├── SkiaCodec.cs              # Codec wrapper
├── SkiaEncoder.cs            # Skia encoder
├── SkiaHelper.cs             # Helper utilities
├── packages.config           # NuGet dependencies
└── Properties/
    └── AssemblyInfo.cs
```

## Decomposition

### ImageEncoder.cs (Skia Image Encoder)

#### Imports
```csharp
using MediaBrowser.Controller.Drawing;
using MediaBrowser.Model.Drawing;
using SkiaSharp;
using System;
using System.IO;
using System.Threading.Tasks;
```

#### Classes
`ImageEncoder` (public class : IImageEncoder)

#### Key Properties
| Property | Type | Description |
|----------|------|-------------|
| `Name` | `string` | "SkiaSharp" |
| `SupportedInputFormats` | `string[]` | JPEG, PNG, WebP, GIF, BMP |

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `EncodeImage(ImageInfo, string, bool)` | `Task<string>` | Encode to file |
| `EncodeImage(ImageInfo, Stream, bool)` | `Task` | Encode to stream |

### ImageProcessor.cs (Skia Image Processor)

#### Classes
`ImageProcessor` (public class)

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `ProcessImage(ImageProcessingOptions)` | `Task<Tuple<string,string,DateTime>>` | Process image |
| `CropImage(SKBitmap, int, int, int, int)` | `SKBitmap` | Crop image |
| `ResizeImage(SKBitmap, int, int)` | `SKBitmap` | Resize image |

### SkiaCodec.cs (Skia Codec Wrapper)

#### Classes
`SkiaCodec` (public static class)

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `CreateFromPath(string)` | `SKCodec` | Load codec from file |
| `CreateFromStream(Stream)` | `SKCodec` | Load codec from stream |
| `GetFrameCount(SKCodec)` | `int` | Get frame count |

## Dependencies

- SkiaSharp (NuGet)
- Emby.Drawing (base interfaces)
