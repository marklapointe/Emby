# Component: Emby.Drawing

**Path:** `Emby.Drawing/`
**Type:** Directory | Library
**Language:** C#
**Maps to:** `.discovery/120-emby-drawing.md`

## Description

Emby.Drawing is the image processing library for Emby Server. It provides image resizing, format conversion, enhancement, and caching. It delegates actual encoding to backend implementations (ImageMagick, Skia, or .NET built-in) via the `IImageEncoder` interface.

## Structure

```
Emby.Drawing/
├── Emby.Drawing.csproj        # Project file
├── ImageProcessor.cs          # Main image processing engine
│   └── [class] ImageProcessor : IImageProcessor, IDisposable
│       ├── [field] IImageEnhancer[] ImageEnhancers
│       ├── [field] IImageEncoder _imageEncoder
│       ├── [field] IFileSystem _fileSystem
│       ├── [field] IServerApplicationPaths _appPaths
│       ├── [field] Func<ILibraryManager> _libraryManager
│       ├── [field] Func<IMediaEncoder> _mediaEncoder
│       ├── [constructor] ImageProcessor(ILogger, IServerApplicationPaths, IFileSystem, IJsonSerializer, Func<ILibraryManager>, Func<IMediaEncoder>)
│       ├── [property] IImageEncoder ImageEncoder
│       │   └── Gets/sets the active image encoder backend
│       ├── [property] string[] SupportedInputFormats
│       │   └── Delegates to IImageEncoder.SupportedInputFormats
│       ├── [property] bool SupportsImageCollageCreation
│       │   └── Delegates to IImageEncoder.SupportsImageCollageCreation
│       ├── [method] public void AddParts(IEnumerable<IImageEnhancer> enhancers)
│       │   └── Registers image enhancer plugins
│       ├── [method] public async Task ProcessImage(ImageProcessingOptions options, Stream toStream)
│       │   └── Processes image and writes to output stream
│       ├── [method] public async Task<Tuple<string, string, DateTime>> ProcessImage(ImageProcessingOptions options)
│       │   ├── Validates input path
│       │   ├── Determines output format (PNG/JPEG/WebP)
│       │   ├── Checks cache for existing processed image
│       │   ├── Applies image enhancers (if any)
│       │   ├── Resizes image via IImageEncoder
│       │   ├── Adds played/unwatched indicators
│       │   ├── Caches result
│       │   └── Returns (cachePath, mimeType, dateModified)
│       ├── [method] public ImageFormat[] GetSupportedImageOutputFormats()
│       │   └── Returns PNG, JPEG, WebP, GIF, BMP
│       ├── [method] public bool SupportsTransparency(string path)
│       │   └── Returns true for .png, .webp, .gif
│       ├── [method] private ImageFormat GetOutputFormat(ImageFormat[] clientSupportedFormats, bool requiresTransparency)
│       │   └── Selects best output format based on client support and transparency needs
│       ├── [method] private string GetCacheFilePath(...)
│       │   ├── Generates deterministic cache path from options hash
│       │   └── Includes version string for cache invalidation
│       ├── [method] public ImageSize GetImageSize(BaseItem item, ItemImageInfo info)
│       │   └── Returns image dimensions (width x height)
│       ├── [method] public ImageSize GetImageSize(string path)
│       │   └── Reads image header for fast size detection
│       ├── [method] public string GetImageCacheTag(BaseItem item, ItemImageInfo image)
│       │   └── Generates cache invalidation tag from item + enhancer state
│       ├── [method] public async Task<string> GetEnhancedImage(BaseItem item, ImageType imageType, int imageIndex)
│       │   └── Applies registered image enhancers to item image
│       ├── [method] private async Task ExecuteImageEnhancers(...)
│       │   └── Runs enhancer pipeline (e.g., cover art, thumbnails)
│       └── [method] public string GetCachePath(string path, string uniqueName, string fileExtension)
│           └── Returns cache directory path for processed images
├── Common/
│   └── ImageHeader.cs           # Image header parser
│       └── [class] ImageHeader
│           ├── [method] public static ImageSize GetDimensions(string path, IFileSystem fileSystem, ILogger logger, bool allowSlowMethod)
│           │   ├── Reads image file headers (JPEG, PNG, GIF, BMP, WebP)
│           │   ├── Extracts width/height without full decode
│           │   └── Falls back to slow method if header parsing fails
│           ├── [method] private static ImageSize? GetDimensionsInternal(...)
│           │   └── Format-specific header parsers
│           └── [method] private static ImageSize GetDimensionsSlow(string path, IImageEncoder encoder)
│               └── Uses IImageEncoder to decode and measure
└── NullImageEncoder.cs          # Fallback no-op encoder
    └── [class] NullImageEncoder : IImageEncoder
        └── No-op implementation when no image backend is available
```

## Data Flow

```mermaid
graph TD
    A[Client requests image] --&gt; B[ProcessImage]
    B --&gt; C[Check cache]
    C --&gt; D{Cache hit?}
    D --&gt;|Yes| E[Return cached image]
    D --&gt;|No| F[Get original image]
    F --&gt; G[Apply enhancers]
    G --&gt; H[Resize via IImageEncoder]
    H --&gt; I[Add indicators]
    I --&gt; J[Save to cache]
    J --&gt; E
```

## Key Interfaces

| Interface | Implemented By | Purpose |
|-----------|---------------|---------|
| `IImageProcessor` | `ImageProcessor` | Main image processing API |
| `IImageEncoder` | `ImageMagickEncoder`, `SkiaEncoder`, `NetEncoder`, `NullImageEncoder` | Backend encoding |
| `IImageEnhancer` | Various plugins | Image enhancement (covers, etc.) |

## Side Effects

- Reads original image files via IFileSystem
- Writes processed images to cache directory
- Invokes IImageEncoder backend (may spawn external processes)
- Executes image enhancer plugins
