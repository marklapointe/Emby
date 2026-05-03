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

## Dependencies

- SkiaSharp (NuGet)
- Emby.Drawing (base interfaces)
