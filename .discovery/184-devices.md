# Component: Emby.Server.Implementations — Devices

**Path:** `Emby.Server.Implementations/Devices/`
**Type:** Directory | Sub-module
**Language:** C#
**Maps to:** `.discovery/184-devices.md`

## Description

Device identification and management. Tracks client devices connecting to the server.

## Files

- `DeviceId.cs` — Emby.Server.Implementations/Devices/DeviceId.cs
- `DeviceManager.cs` — Emby.Server.Implementations/Devices/DeviceManager.cs

## Architecture

```mermaid
graph LR
    A[Client Connection] --> B[DeviceId]
    B --> C[DeviceManager]
    C --> D[SQLite Database]
```

## Key Classes

| Class | Responsibility |
|-------|----------------|
| `DeviceId` | Device identification |
| `DeviceManager` | Device registration and tracking |

## Dependencies

- `MediaBrowser.Controller` — Device interfaces
