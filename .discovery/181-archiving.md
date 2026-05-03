# Component: Emby.Server.Implementations — Archiving

**Path:** `Emby.Server.Implementations/Archiving/`
**Type:** Directory | Sub-module
**Language:** C#
**Maps to:** `.discovery/181-archiving.md`

## Description

ZIP archive handling utilities for packaging and extracting media collections.

## Files

- `ZipClient.cs` — Emby.Server.Implementations/Archiving/ZipClient.cs

## Architecture

```mermaid
graph LR
    A[Media Files] --> B[ZipClient]
    B --> C[ZIP Archive]
    C --> D[Extracted Files]
```

## Dependencies

- `System.IO.Compression` — ZIP support
