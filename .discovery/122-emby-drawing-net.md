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

## Dependencies

- System.Drawing (built-in)
- Emby.Drawing (base interfaces)
