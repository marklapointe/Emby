# Component: MediaBrowser.Providers.Users

**Path:** `MediaBrowser.Providers/Users/`
**Type:** Directory | Sub-Module
**Language:** C#
**Maps to:** `.discovery/346-mediabrowser-providers-users.md`

## Description

User metadata services. Handles metadata and profile management for user entities.

## Directory Structure

```
MediaBrowser.Providers/Users/
└── UserMetadataService.cs
```

## Files

| File | Description |
|------|-------------|
| `UserMetadataService.cs` | User metadata service |

## Decomposition

### UserMetadataService.cs

#### Classes
`UserMetadataService` (public class : IMetadataService)

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `Fetch(MetadataSearchOptions, CancellationToken)` | `Task<bool>` | Fetch user metadata |
| `Save(BaseItem, CancellationToken)` | `Task` | Save user metadata |

## Architecture

```mermaid
graph TB
    A[User Providers] --> B[UserMetadataService]
    B --> C[User Entity]
```

## Dependencies

- MediaBrowser.Controller.Entities — Entity types
- MediaBrowser.Controller.Providers — Provider interfaces

## Statistics

| Metric | Value |
|--------|-------|
| C# Files | 1 |
| LOC | ~1,100 |
| Public Classes | 1 |
