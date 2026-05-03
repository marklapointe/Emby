# MediaBrowser.Api - Session Subdirectory

**Module:** MediaBrowser.Api/Session
**Language:** C#
**Maps to:** `.discovery/351-mediabrowser-api-session.md`

## Decomposition

### SessionsService.cs (Sessions API Service)

#### Imports
```csharp
using MediaBrowser.Controller.Net;
using MediaBrowser.Model.Services;
using System.Threading.Tasks;
```

#### Classes
`SessionsService` (public class : IRequiresRequest)

### SessionInfoWebSocketListener.cs (Session WebSocket Listener)

#### Classes
`SessionInfoWebSocketListener` (public class)

## File Listing

```
Session/
├── SessionsService.cs            - Sessions API service
└── SessionInfoWebSocketListener.cs - Session WebSocket events
```

## Description

Session subdirectory contains session management API services and WebSocket listeners for real-time session updates.

## Dependencies

- **MediaBrowser.Controller.Net** - Networking interfaces
- **MediaBrowser.Model.Services** - Service interfaces
