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

## Dependencies

- ImageMagickSharp (NuGet)
- Emby.Drawing (base interfaces)
