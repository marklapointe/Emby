# MediaBrowser.Api - Music Subdirectory

**Module:** MediaBrowser.Api/Music
**Language:** C#
**Maps to:** `.discovery/350-mediabrowser-api-music.md`

## Decomposition

### AlbumsService.cs (Albums API Service)

#### Imports
```csharp
using MediaBrowser.Controller.Net;
using MediaBrowser.Model.Services;
using System.Threading.Tasks;
```

#### Classes
`AlbumsService` (public class : IRequiresRequest)

### InstantMixService.cs (Instant Mix Service)

#### Classes
`InstantMixService` (public class)

## File Listing

```
Music/
├── AlbumsService.cs     - Albums API service
└── InstantMixService.cs - Instant mix generation service
```

## Description

Music subdirectory contains music-related API services for album browsing and instant mix generation.

## Dependencies

- **MediaBrowser.Controller.Net** - Networking interfaces
- **MediaBrowser.Model.Services** - Service interfaces
