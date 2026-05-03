# Component: Emby.Server.Implementations — Scheduled Tasks

**Path:** \`Emby.Server.Implementations/ScheduledTasks/\`
**Type:** Directory | Module Group
**Language:** C#
**Maps to:** \`.discovery/165-emby-server-impl-tasks.md\`

## Description

Background task scheduler and built-in task implementations.

## Files

### ScheduledTasks/Tasks/

- `DeleteCacheFileTask.cs` — Emby.Server.Implementations/ScheduledTasks/Tasks/DeleteCacheFileTask.cs
- `DeleteLogFileTask.cs` — Emby.Server.Implementations/ScheduledTasks/Tasks/DeleteLogFileTask.cs
- `ReloadLoggerFileTask.cs` — Emby.Server.Implementations/ScheduledTasks/Tasks/ReloadLoggerFileTask.cs

## Decomposition

### IScheduledTask.cs (Task Interface)

#### Imports
```csharp
using MediaBrowser.Model.Tasks;
using System;
using System.Threading;
```

#### Classes
`IScheduledTask` (public interface)

#### Key Properties
| Property | Type | Description |
|----------|------|-------------|
| `Key` | `string` | Unique task key |
| `Name` | `string` | Display name |
| `Description` | `string` | Task description |
| `Category` | `string` | Task category |
| `IsHidden` | `bool` | Hide from UI |

### ScheduledTaskManager.cs (Task Scheduler)

#### Classes
`ScheduledTaskManager` (public class : IScheduledTaskManager)

#### Key Properties
| Property | Type | Description |
|----------|------|-------------|
| `ScheduledTasks` | `IEnumerable<IScheduledTask>` | All registered tasks |
| `RunningTasks` | `IEnumerable<TaskExecutionSummary>` | Currently running |

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `Execute<T>(string, bool)` | `Task` | Execute task by type |
| `Cancel(string)` | `void` | Cancel running task |
| `GetScheduledTasks(Type)` | `IEnumerable<IScheduledTask>` | Get tasks by type |

### DeleteCacheFileTask.cs (Cache Cleanup)

#### Classes
`DeleteCacheFileTask` (public class : IScheduledTask)

#### Key Properties
| Property | Type | Description |
|----------|------|-------------|
| `Name` | `string` | "Delete Cache Files" |
| `Description` | `string` | "Deletes cached thumbnail and media images" |

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `GetProgress(IProgress<double>)` | `Task` | Execute cleanup |

### DeleteLogFileTask.cs (Log Cleanup)

#### Classes
`DeleteLogFileTask` (public class : IScheduledTask)

#### Key Properties
| Property | Type | Description |
|----------|------|-------------|
| `Name` | `string` | "Delete Log Files" |
| `Description` | `string` | "Deletes log files older than 3 days" |

