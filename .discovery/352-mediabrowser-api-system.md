# MediaBrowser.Api - System Subdirectory

**Module:** MediaBrowser.Api/System
**Language:** C#
**Maps to:** `.discovery/352-mediabrowser-api-system.md`

## Decomposition

### SystemService.cs (System API Service)

#### Imports
```csharp
using MediaBrowser.Controller.Net;
using MediaBrowser.Model.Services;
using System.Threading.Tasks;
```

#### Classes
`SystemService` (public class : IRequiresRequest)

### ActivityLogService.cs (Activity Log Service)

#### Classes
`ActivityLogService` (public class)

### ActivityLogWebSocketListener.cs (Activity Log WebSocket)

#### Classes
`ActivityLogWebSocketListener` (public class)

## File Listing

```
System/
├── SystemService.cs               - System API service
├── ActivityLogService.cs          - Activity log service
└── ActivityLogWebSocketListener.cs - Activity log WebSocket
```

## Description

System subdirectory contains system-level API services for server health, activity logging, and system configuration.

## Dependencies

- **MediaBrowser.Controller.Net** - Networking interfaces
- **MediaBrowser.Model.Services** - Service interfaces
