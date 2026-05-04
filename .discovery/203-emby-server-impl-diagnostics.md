# Component: Emby.Server.Implementations — Diagnostics

**Path:** `Emby.Server.Implementations/Diagnostics/`
**Type:** Directory | Module
**Language:** C#
**Maps to:** `.discovery/203-emby-server-impl-diagnostics.md`

## Description

Diagnostics and process management utilities. Provides process discovery and monitoring capabilities.

## Files

- `CommonProcess.cs` — Emby.Server.Implementations/Diagnostics/CommonProcess.cs
- `ProcessFactory.cs` — Emby.Server.Implementations/Diagnostics/ProcessFactory.cs

## Decomposition

### CommonProcess.cs (Common Process Wrapper)

#### Imports
```csharp
using System;
using System.Diagnostics;
using System.Threading.Tasks;
```

#### Classes
`CommonProcess` (public class)

#### Key Properties
| Property | Type | Description |
|----------|------|-------------|
| `ProcessName` | `string` | Name of process |
| `Id` | `int` | Process ID |
| `StartTime` | `DateTime` | When started |
| `HasExited` | `bool` | Is process running |

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `Kill()` | `void` | Terminate process |
| `WaitForExit(int)` | `bool` | Wait for exit |
| `Start(ProcessStartInfo)` | `static CommonProcess` | Start new process |

### ProcessFactory.cs (Process Factory)

#### Classes
`ProcessFactory` (public interface)

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `CreateProcess(ProcessStartInfo)` | `IProcess` | Create new process |

## Data Flow

```mermaid
graph LR
    A[Diagnostics Request] --> B[ProcessFactory]
    B --> C[CommonProcess]
    C --> D[System Process]
```

## Dependencies

- `System.Diagnostics` — Process management
- `System.Threading.Tasks` — Async operations

## Statistics

| Metric | Value |
|--------|-------|
| Files | 2 |
| Classes | 2 |
| LOC | ~115 |
