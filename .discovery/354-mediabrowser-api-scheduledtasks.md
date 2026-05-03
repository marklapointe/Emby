# MediaBrowser.Api - ScheduledTasks Subdirectory

**Module:** MediaBrowser.Api/ScheduledTasks
**Language:** C#
**Maps to:** `.discovery/354-mediabrowser-api-scheduledtasks.md`

## Decomposition

### ScheduledTaskService.cs (Scheduled Tasks Service)

#### Imports
```csharp
using MediaBrowser.Controller.Net;
using MediaBrowser.Controller.ScheduledTasks;
using MediaBrowser.Model.Services;
using System.Threading.Tasks;
```

#### Classes
`ScheduledTaskService` (public class : IRequiresRequest)

### ScheduledTasksWebSocketListener.cs (Tasks WebSocket Listener)

#### Classes
`ScheduledTasksWebSocketListener` (public class)

## File Listing

```
ScheduledTasks/
├── ScheduledTaskService.cs          - Scheduled tasks API service
└── ScheduledTasksWebSocketListener.cs - Tasks WebSocket events
```

## Description

ScheduledTasks subdirectory contains scheduled task management API services and WebSocket listeners.

## Dependencies

- **MediaBrowser.Controller.ScheduledTasks** - Task interfaces
- **MediaBrowser.Model.Services** - Service interfaces
