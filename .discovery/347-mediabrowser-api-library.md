# MediaBrowser.Api - Library Subdirectory

**Module:** MediaBrowser.Api/Library
**Language:** C#
**Maps to:** `.discovery/347-mediabrowser-api-library.md`

## Decomposition

### LibraryService.cs (Library Management Service)

#### Imports
```csharp
using MediaBrowser.Controller.Library;
using MediaBrowser.Controller.Net;
using MediaBrowser.Model.Services;
using System.Threading.Tasks;
```

#### Classes
`LibraryService` (public class : IRequiresRequest)

### LibraryStructureService.cs (Library Structure Service)

#### Classes
`LibraryStructureService` (public class)

## File Listing

```
Library/
├── LibraryService.cs         - Library management service
└── LibraryStructureService.cs - Library structure service
```

## Description

Library subdirectory contains library management API services for managing media libraries and their structure.

## Dependencies

- **MediaBrowser.Controller.Library** - Library interfaces
- **MediaBrowser.Model.Services** - Service interfaces
