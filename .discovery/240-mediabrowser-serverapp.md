# Component: MediaBrowser.ServerApplication

**Path:** `MediaBrowser.ServerApplication/`
**Type:** Directory | Application
**Language:** C#
**Maps to:** `.discovery/240-mediabrowser-serverapp.md`

## Description

MediaBrowser.ServerApplication is the Windows entry point for Emby Server. It provides a Windows Forms or console application that bootstraps the server on .NET Framework, handling Windows-specific concerns like service installation, tray icon, and registry configuration.

## Structure

```
MediaBrowser.ServerApplication/
├── MediaBrowser.ServerApplication.csproj
├── Program.cs                   # Windows entry point → [class] Program
├── MainForm.cs                  # Windows Forms UI (optional)
├── ServerApplicationPaths.cs    # Windows path configuration
└── Properties/                  # Assembly info
```

## Key Classes

| Class | File | Purpose |
|-------|------|---------|
| `Program` | `Program.cs` | Main entry point for Windows |
| `MainForm` | `MainForm.cs` | Windows UI (if present) |

## Data Flow

```mermaid
graph LR
    A[Program.Main] --> B[ServerApplicationPaths]
    B --> C[ApplicationHost]
    C --> D[Emby.Server.Implementations]
    D --> E[Server Running]
```

## Dependencies

- `Emby.Server.Implementations` — Core server
- `MediaBrowser.Controller` — Controller interfaces

## Side Effects

- Starts server process on Windows
- May show system tray icon
- Writes to Windows event log

## Reference

- .NET Framework runtime
