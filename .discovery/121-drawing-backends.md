# Component: Drawing Backends (ImageMagick, .NET, Skia)

**Path:** `Emby.Drawing.ImageMagick/`, `Emby.Drawing.Net/`, `Emby.Drawing.Skia/`
**Type:** Directory | Library Group
**Language:** C#
**Maps to:** `.discovery/121-drawing-backends.md`

## Description

Three interchangeable image encoding backends implementing `IImageEncoder`. Only one is active at runtime. ImageMagick is the most feature-rich, .NET GDI+ is the fallback, and Skia is the cross-platform high-performance option.

## Structure

### Emby.Drawing.ImageMagick (7 C# files)

```
Emby.Drawing.ImageMagick/
├── ImageMagickEncoder.cs        # Main encoder
│   └── [class] ImageMagickEncoder : IImageEncoder
│       ├── [constructor] ImageMagickEncoder(ILogger, IFileSystem, IApplicationPaths)
│       ├── [property] string[] SupportedInputFormats
│       │   └── JPEG, PNG, GIF, BMP, WebP, TIFF, TGA, SVG, etc.
│       ├── [method] public void EncodeImage(string inputPath, ImageFormat outputFormat, int quality, string outputPath, int width, int height, bool cropWhitespace)
│       │   ├── Uses ImageMagick.NET (MagickImage)
│       │   ├── Resizes with Lanczos or Box filter
│       │   ├── Crops whitespace if requested
│       │   └── Saves to output format
│       ├── [method] public void CreateImageCollage(ImageCollageOptions options)
│       │   └── Creates photo collage using StripCollageBuilder
│       ├── [method] public void CreateSplashscreen(ImageCollageOptions options)
│       │   └── Creates splash screen collage
│       ├── [method] public void DrawPlayedIndicator(string path, int width, int height)
│       │   └── Delegates to PlayedIndicatorDrawer
│       ├── [method] public void DrawUnplayedCountIndicator(string path, int width, int height, int count)
│       │   └── Delegates to UnplayedCountIndicator
│       ├── [method] public void DrawPercentPlayed(string path, double percentPlayed)
│       │   └── Delegates to PercentPlayedDrawer
│       └── [method] public ImageSize GetImageSize(string path)
│           └── Reads image dimensions via ImageMagick
├── ImageHelpers.cs              # ImageMagick-specific helpers
│   └── [class] ImageHelpers
│       └── [method] public static void CropWhitespace(MagickImage image)
│           └── Removes transparent/white borders
├── StripCollageBuilder.cs       # Photo strip collage builder
│   └── [class] StripCollageBuilder
│       └── [method] public void BuildCollage(...)
│           ├── Arranges images in horizontal strip
│           ├── Resizes each to uniform height
│           └── Composites into single image
├── PlayedIndicatorDrawer.cs     # "Played" checkmark overlay
│   └── [class] PlayedIndicatorDrawer
│       └── Draws green checkmark in corner
├── UnplayedCountIndicator.cs    # Unwatched count badge
│   └── [class] UnplayedCountIndicator
│       └── Draws circular badge with number
└── PercentPlayedDrawer.cs       # Progress bar overlay
    └── [class] PercentPlayedDrawer
        └── Draws horizontal progress bar at bottom
```

### Emby.Drawing.Net (8 C# files)

```
Emby.Drawing.Net/
├── GDIImageEncoder.cs           # .NET GDI+ encoder
│   └── [class] GDIImageEncoder : IImageEncoder
│       ├── [method] public void EncodeImage(string inputPath, ImageFormat outputFormat, int quality, string outputPath, int width, int height, bool cropWhitespace)
│       │   ├── Uses System.Drawing (GDI+)
│       │   ├── Resizes with HighQualityBicubic
│       │   └── Saves to output format
│       ├── [method] public void CreateImageCollage(ImageCollageOptions options)
│       │   └── Not supported (throws NotImplementedException)
│       ├── [method] public void DrawPlayedIndicator(string path, int width, int height)
│       │   └── Delegates to PlayedIndicatorDrawer
│       ├── [method] public void DrawUnplayedCountIndicator(string path, int width, int height, int count)
│       │   └── Delegates to UnplayedCountIndicator
│       ├── [method] public void DrawPercentPlayed(string path, double percentPlayed)
│       │   └── Delegates to PercentPlayedDrawer
│       └── [method] public ImageSize GetImageSize(string path)
│           └── Uses System.Drawing.Image.FromFile
├── ImageExtensions.cs           # GDI+ image extensions
│   └── [class] ImageExtensions
│       └── Extension methods for System.Drawing.Image
├── DynamicImageHelpers.cs       # Dynamic image helpers
│   └── [class] DynamicImageHelpers
│       └── Helper methods for dynamic image generation
├── ImageHelpers.cs              # GDI+ image helpers
│   └── [class] ImageHelpers
│       └── [method] public static void CropWhitespace(Image image)
│           └── Removes whitespace borders via pixel scanning
├── StripCollageBuilder.cs       # Not implemented
│   └── [class] StripCollageBuilder
│       └── Throws NotImplementedException
├── PlayedIndicatorDrawer.cs     # GDI+ checkmark overlay
├── UnplayedCountIndicator.cs    # GDI+ count badge
└── PercentPlayedDrawer.cs       # GDI+ progress bar
```

### Emby.Drawing.Skia (6 C# files)

```
Emby.Drawing.Skia/
├── SkiaEncoder.cs               # SkiaSharp encoder
│   └── [class] SkiaEncoder : IImageEncoder
│       ├── [constructor] SkiaEncoder(ILogger, IFileSystem, IApplicationPaths)
│       ├── [method] public void EncodeImage(string inputPath, ImageFormat outputFormat, int quality, string outputPath, int width, int height, bool cropWhitespace)
│       │   ├── Uses SkiaSharp (SKBitmap, SKCanvas)
│       │   ├── Resizes with Lanczos filter
│       │   ├── Crops whitespace if requested
│       │   └── Saves to output format
│       ├── [method] public void CreateImageCollage(ImageCollageOptions options)
│       │   └── Creates collage using StripCollageBuilder
│       ├── [method] public void CreateSplashscreen(ImageCollageOptions options)
│       │   └── Creates splash screen
│       ├── [method] public void DrawPlayedIndicator(string path, int width, int height)
│       │   └── Delegates to PlayedIndicatorDrawer
│       ├── [method] public void DrawUnplayedCountIndicator(string path, int width, int height, int count)
│       │   └── Delegates to UnplayedCountIndicator
│       ├── [method] public void DrawPercentPlayed(string path, double percentPlayed)
│       │   └── Delegates to PercentPlayedDrawer
│       └── [method] public ImageSize GetImageSize(string path)
│           └── Reads image header via SkiaSharp
├── StripCollageBuilder.cs       # Skia photo strip builder
│   └── [class] StripCollageBuilder
│       └── [method] public void BuildCollage(...)
│           ├── Uses SKCanvas for compositing
│           └── Arranges images in strip layout
├── PlayedIndicatorDrawer.cs     # Skia checkmark overlay
├── UnplayedCountIndicator.cs    # Skia count badge
└── PercentPlayedDrawer.cs       # Skia progress bar
```

## Backend Comparison

| Feature | ImageMagick | .NET GDI+ | Skia |
|---------|-------------|-----------|------|
| Image formats | Most comprehensive | Limited | Good |
| Collage creation | ✅ | ❌ | ✅ |
| Splash screen | ✅ | ❌ | ✅ |
| Performance | Good | Moderate | Best |
| Cross-platform | Good | Windows only | Excellent |
| Dependencies | ImageMagick native | System.Drawing | SkiaSharp native |

## Data Flow

```mermaid
graph TD
    A[ImageProcessor.ProcessImage] --&gt; B{Active Backend?}
    B --&gt;|ImageMagick| C[ImageMagickEncoder.EncodeImage]
    B --&gt;|.NET| D[GDIImageEncoder.EncodeImage]
    B --&gt;|Skia| E[SkiaEncoder.EncodeImage]
    C --&gt; F[MagickImage resize/save]
    D --&gt; G[System.Drawing resize/save]
    E --&gt; H[SKBitmap resize/save]
    F --&gt; I[Output file]
    G --&gt; I
    H --&gt; I
```

## Side Effects

- ImageMagick: Requires native ImageMagick libraries (Magick.NET-Q8)
- .NET GDI+: Requires Windows GDI+ or libgdiplus on Linux
- Skia: Requires native SkiaSharp libraries
- All backends: Read source images, write processed images to disk
