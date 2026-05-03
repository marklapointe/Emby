# Component: MediaBrowser.Providers — Books

**Path:** `MediaBrowser.Providers/Books/`
**Type:** Directory | Sub-module
**Language:** C#
**Maps to:** `.discovery/326-mediabrowser-providers-books.md`
**Parent:** `.discovery/320-mediabrowser-providers.md`

## Description

Metadata providers for books and audiobooks. Handles metadata extraction and organization for book content.

## Files

### Root Book Files (2 files)

- `BookMetadataService.cs` — MediaBrowser.Providers/Books/BookMetadataService.cs
- `AudioBookMetadataService.cs` — MediaBrowser.Providers/Books/AudioBookMetadataService.cs

## Key Classes

| Class | File | Purpose |
|-------|------|---------|
| `BookMetadataService` | `BookMetadataService.cs` | Book metadata orchestration |
| `AudioBookMetadataService` | `AudioBookMetadataService.cs` | Audiobook metadata orchestration |

## Dependencies

- **MediaBrowser.Controller** — Base entity types
- **MediaBrowser.Model** — API models
- **HttpClient** — External API calls
