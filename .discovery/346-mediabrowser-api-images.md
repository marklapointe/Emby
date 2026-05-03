# MediaBrowser.Api - Images Subdirectory

**Module:** MediaBrowser.Api/Images
**Language:** C#
**Maps to:** `.discovery/346-mediabrowser-api-images.md`

## Decomposition

### ImageService.cs (Main Image Service)

#### Imports
```csharp
using MediaBrowser.Controller.Drawing;
using MediaBrowser.Controller.Entities;
using MediaBrowser.Controller.Net;
using MediaBrowser.Model.Drawing;
using MediaBrowser.Model.Services;
using System.Threading.Tasks;
```

#### Classes
`ImageService` (public class : IRequiresRequest)

#### Key Methods
```csharp
Task GetImage(ImageRequest request)
```

### ImageByNameService.cs (Named Image Service)

#### Classes
`ImageByNameService` (public class)

### ImageRequest.cs (Request Model)

#### Classes
`ImageRequest` (public class)

### RemoteImageService.cs (Remote Image Service)

#### Classes
`RemoteImageService` (public class)

## File Listing

```
Images/
├── ImageService.cs       - Main image service
├── ImageByNameService.cs - Named image service
├── ImageRequest.cs      - Request model
└── RemoteImageService.cs - Remote image service
```

## Description

Images subdirectory contains image-related API services for retrieving and processing media artwork.

## Dependencies

- **MediaBrowser.Controller.Drawing** - Drawing interfaces
- **MediaBrowser.Model.Drawing** - Drawing models
- **MediaBrowser.Model.Services** - Service interfaces
