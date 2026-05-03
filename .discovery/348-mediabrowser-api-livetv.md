# MediaBrowser.Api - LiveTv Subdirectory

**Module:** MediaBrowser.Api/LiveTv
**Language:** C#
**Maps to:** `.discovery/348-mediabrowser-api-livetv.md`

## Decomposition

### LiveTvService.cs (Live TV Service)

#### Imports
```csharp
using MediaBrowser.Controller.LiveTv;
using MediaBrowser.Controller.Net;
using MediaBrowser.Model.Services;
using System.Threading.Tasks;
```

#### Classes
`LiveTvService` (public class : IRequiresRequest)

### ProgressiveFileCopier.cs (File Copier)

#### Classes
`ProgressiveFileCopier` (public class)

## File Listing

```
LiveTv/
├── LiveTvService.cs       - Live TV API service
└── ProgressiveFileCopier.cs - Progressive file copier for recordings
```

## Description

LiveTv subdirectory contains Live TV API services for managing live TV streams and recordings.

## Dependencies

- **MediaBrowser.Controller.LiveTv** - Live TV interfaces
- **MediaBrowser.Model.Services** - Service interfaces
