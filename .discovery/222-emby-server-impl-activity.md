# Component: Emby.Server.Implementations — Activity

**Path:** `Emby.Server.Implementations/Activity/`
**Type:** Directory | Module
**Language:** C#
**Maps to:** `.discovery/222-emby-server-impl-activity.md`

## Description

Activity logging and tracking. Records user activities like playback, library changes, and system events.

## Files

- `ActivityLogEntryPoint.cs` — Emby.Server.Implementations/Activity/ActivityLogEntryPoint.cs
- `ActivityManager.cs` — Emby.Server.Implementations/Activity/ActivityManager.cs
- `ActivityRepository.cs` — Emby.Server.Implementations/Activity/ActivityRepository.cs

## Decomposition

### ActivityManager.cs (Activity Manager)

#### Imports
```csharp
using MediaBrowser.Controller.Activity;
using MediaBrowser.Model.Activity;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
```

#### Classes
`ActivityManager` (public class : IActivityManager)

#### Key Properties
| Property | Type | Description |
|----------|------|-------------|
| `ActivityLogEntries` | `IEnumerable<ActivityLogEntry>` | Log entries |

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `GetActivityLog(DateTime?, DateTime?)` | `Task<IEnumerable<ActivityLogEntry>>` | Get log |
| `LogEntry(ActivityLogEntry)` | `Task` | Log activity |
| `GetItemLogEntries(Guid)` | `Task<IEnumerable<ActivityLogEntry>>` | Get item activity |

### ActivityRepository.cs (Activity Repository)

#### Classes
`ActivityRepository` (public class)

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `Create(ActivityLogEntry)` | `Task` | Create log entry |
| `Get(DateTime?, DateTime?)` | `IEnumerable<ActivityLogEntry>` | Get entries |

## Data Flow

```mermaid
graph LR
    A[User Action] --> B[ActivityLogEntryPoint]
    B --> C[ActivityManager]
    C --> D[ActivityRepository]
    D --> E[Database]
```

## Dependencies

- `MediaBrowser.Controller.Activity` — Activity interfaces
- `MediaBrowser.Model.Activity` — Activity models

## Statistics

| Metric | Value |
|--------|-------|
| Files | 3 |
| Classes | 3 |
| LOC | ~150 |
