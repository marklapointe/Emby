# MediaBrowser.Api - Subtitles Subdirectory

**Module:** MediaBrowser.Api/Subtitles
**Language:** C#
**Maps to:** `.discovery/355-mediabrowser-api-subtitles.md`

## Decomposition

### SubtitleService.cs (Subtitles API Service)

#### Imports
```csharp
using MediaBrowser.Controller.Net;
using MediaBrowser.Model.Services;
using System.Threading.Tasks;
```

#### Classes
`SubtitleService` (public class : IRequiresRequest)

## File Listing

```
Subtitles/
└── SubtitleService.cs - Subtitles API service
```

## Description

Subtitles subdirectory contains subtitle-related API services for downloading and streaming subtitles.

## Dependencies

- **MediaBrowser.Controller.Net** - Networking interfaces
- **MediaBrowser.Model.Services** - Service interfaces
