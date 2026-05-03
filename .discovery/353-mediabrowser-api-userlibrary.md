# MediaBrowser.Api - UserLibrary Subdirectory

**Module:** MediaBrowser.Api/UserLibrary
**Language:** C#
**Maps to:** `.discovery/353-mediabrowser-api-userlibrary.md`

## Decomposition

### UserLibraryService.cs (User Library Service)

#### Imports
```csharp
using MediaBrowser.Controller.Library;
using MediaBrowser.Controller.Net;
using MediaBrowser.Model.Services;
using System.Threading.Tasks;
```

#### Classes
`UserLibraryService` (public class : IRequiresRequest)

### ArtistsService.cs (Artists Service)

#### Classes
`ArtistsService` (public class)

### BaseItemsByNameService.cs (Base Items By Name)

#### Classes
`BaseItemsByNameService` (public class)

### GenresService.cs / MusicGenresService.cs / GameGenresService.cs (Genre Services)

#### Classes
Various genre service classes

### ItemsService.cs (Items Service)

#### Classes
`ItemsService` (public class)

### PersonsService.cs / StudiosService.cs (Metadata Services)

#### Classes
Various metadata service classes

### UserViewsService.cs / YearsService.cs (View Services)

#### Classes
Various view service classes

### BaseItemsRequest.cs (Request Model)

#### Classes
`BaseItemsRequest` (public class)

## File Listing

```
UserLibrary/
├── UserLibraryService.cs         - Main user library service
├── ArtistsService.cs            - Artists service
├── BaseItemsByNameService.cs    - Base items by name
├── GenresService.cs             - Genres service
├── MusicGenresService.cs        - Music genres service
├── GameGenresService.cs         - Game genres service
├── ItemsService.cs              - Items service
├── PersonsService.cs            - Persons service
├── StudiosService.cs            - Studios service
├── UserViewsService.cs          - User views service
├── YearsService.cs              - Years service
└── BaseItemsRequest.cs          - Request model
```

## Description

UserLibrary subdirectory contains comprehensive library browsing services for artists, genres, items, persons, studios, and user views.

## Dependencies

- **MediaBrowser.Controller.Library** - Library interfaces
- **MediaBrowser.Model.Services** - Service interfaces
