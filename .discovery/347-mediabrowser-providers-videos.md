# Component: MediaBrowser.Providers.Videos

**Path:** `MediaBrowser.Providers/Videos/`
**Type:** Directory | Sub-Module
**Language:** C#
**Maps to:** `.discovery/347-mediabrowser-providers-videos.md`

## Description

Video metadata services. Handles metadata and artwork for generic video entities.

## Directory Structure

```
MediaBrowser.Providers/Videos/
└── VideoMetadataService.cs
```

## Files

| File | Description |
|------|-------------|
| `VideoMetadataService.cs` | Video metadata service |

## Decomposition

### VideoMetadataService.cs

#### Classes
`VideoMetadataService` (public class : IMetadataService)

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `Fetch(MetadataSearchOptions, CancellationToken)` | `Task<bool>` | Fetch video metadata |
| `Save(BaseItem, CancellationToken)` | `Task` | Save video metadata |

## Architecture

```mermaid
graph TB
    A[Video Providers] --> B[VideoMetadataService]
    B --> C[Video Entity]
```

## Dependencies

- MediaBrowser.Controller.Entities — Entity types
- MediaBrowser.Controller.Providers — Provider interfaces

## Statistics

| Metric | Value |
|--------|-------|
| C# Files | 1 |
| LOC | ~1,300 |
| Public Classes | 1 |
