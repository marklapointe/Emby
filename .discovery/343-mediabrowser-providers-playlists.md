# MediaBrowser.Providers - Playlists Module

**Module:** MediaBrowser.Providers/Playlists
**Language:** C#
**Maps to:** `.discovery/343-mediabrowser-providers-playlists.md`

## Decomposition

### PlaylistMetadataService.cs (Playlist Metadata Service)

#### Classes
`PlaylistMetadataService` (public class : IMetadataService)

### PlaylistItemsProvider.cs (Playlist Items Provider)

#### Classes
`PlaylistItemsProvider` (public class : IItemVideoResolver)

## File Listing

```
Playlists/
├── PlaylistMetadataService.cs - Playlist metadata service
└── PlaylistItemsProvider.cs  - Playlist items resolver
```

## Description

Playlists module provides metadata services and item resolvers for playlist items.

## Dependencies

- **MediaBrowser.Controller.Providers** - Provider interfaces
