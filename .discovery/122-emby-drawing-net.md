# Component: Emby.Drawing.Net

**Path:** \`Emby.Drawing.Net/\`
**Type:** Module
**Maps to:** \`.discovery/122-emby-drawing-net.md\`

## Description

.NET Framework drawing backend for Emby Server (System.Drawing-based).

## Structure

```
Emby.Drawing.Net/
├── Emby.Drawing.Net.csproj
├── ImageEncoder.cs           # Image encoding
├── ImageProcessor.cs         # Image processing
├── Image.cs                  # Image wrapper
├── WebpEncoder.cs            # WebP encoding
├── packages.config           # NuGet dependencies
└── Properties/
    └── AssemblyInfo.cs
```

## Decomposition

### ImageEncoder.cs (.NET Image Encoder)

#### Imports
```csharp
using MediaBrowser.Controller.Drawing;
using MediaBrowser.Model.Drawing;
using System;
using System.Drawing;
using System.Drawing.Imaging;
using System.IO;
using System.Threading.Tasks;
```

#### Classes
`ImageEncoder` (public class : IImageEncoder)

#### Key Properties
| Property | Type | Description |
|----------|------|-------------|
| `Name` | `string` | "Default" |
| `SupportedInputFormats` | `string[]` | BMP, GIF, JPG, PNG, TIFF |

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `EncodeImage(ImageInfo, string, bool)` | `Task<string>` | Encode to file |
| `EncodeImage(ImageInfo, Stream, bool)` | `Task` | Encode to stream |

### ImageProcessor.cs (.NET Image Processor)

#### Classes
`ImageProcessor` (public class)

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `ProcessImage(ImageProcessingOptions)` | `Task<Tuple<string,string,DateTime>>` | Process image |

### WebpEncoder.cs (WebP Encoder)

#### Classes
`WebpEncoder` (public class)

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `EncodeToWebp(Bitmap, Stream)` | `void` | Encode bitmap to WebP |

## Dependencies

- System.Drawing (built-in)
- Emby.Drawing (base interfaces)
