# MediaBrowser.Api - Movies Subdirectory

**Module:** MediaBrowser.Api/Movies
**Language:** C#
**Maps to:** `.discovery/349-mediabrowser-api-movies.md`

## Decomposition

### MoviesService.cs (Movies API Service)

#### Imports
```csharp
using MediaBrowser.Controller.Entities;
using MediaBrowser.Controller.Library;
using MediaBrowser.Controller.Net;
using MediaBrowser.Model.Services;
using System.Threading.Tasks;
```

#### Classes
`MoviesService` (public class : IRequiresRequest)

### CollectionService.cs (Collection Service)

#### Classes
`CollectionService` (public class)

### TrailersService.cs (Trailers Service)

#### Classes
`TrailersService` (public class)

## File Listing

```
Movies/
├── MoviesService.cs     - Movies API service
├── CollectionService.cs - Collection management service
└── TrailersService.cs  - Trailers service
```

## Description

Movies subdirectory contains movie-related API services including collection and trailer management.

## Dependencies

- **MediaBrowser.Controller.Entities** - Movie entities
- **MediaBrowser.Controller.Library** - Library interfaces
- **MediaBrowser.Model.Services** - Service interfaces
