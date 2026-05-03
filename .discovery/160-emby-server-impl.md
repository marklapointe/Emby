# Component: Emby.Server.Implementations

**Path:** \`Emby.Server.Implementations/\`
**Type:** Directory | Module
**Language:** C#
**Maps to:** \`.discovery/160-emby-server-impl.md\`

## Description

Emby.Server.Implementations is the core server module containing the bulk of Emby's business logic. It implements the \`MediaBrowser.Controller\` interfaces with concrete classes for library management, user management, media encoding, HTTP server, session management, scheduled tasks, and more. This is the largest and most critical module in the solution with 514 C# source files.

## Sub-Modules

| Sub-Module | Document | Description |
|------------|----------|-------------|
| Core Infrastructure | [161-emby-server-impl-core.md](./161-emby-server-impl-core.md) | AppBase, Configuration, Cryptography, Data, Diagnostics, Environment, Logging, Networking, Reflection, Serialization, Security, Threading, XML |
| Library Management | [162-emby-server-impl-library.md](./162-emby-server-impl-library.md) | Library, Collections, Playlists, UserViews, Sorting |
| Media & Channels | [163-emby-server-impl-media.md](./163-emby-server-impl-media.md) | Channels, LiveTV |
| HTTP Server & Services | [164-emby-server-impl-http.md](./164-emby-server-impl-http.md) | HttpServer, HttpClientManager, Services, EntryPoints |
| Scheduled Tasks | [165-emby-server-impl-tasks.md](./165-emby-server-impl-tasks.md) | Background task scheduler and built-in tasks |
| I/O Utilities | [166-emby-server-impl-io.md](./166-emby-server-impl-io.md) | File system I/O wrappers |
| Text Encoding & Localization | [167-emby-server-impl-encoding.md](./167-emby-server-impl-encoding.md) | Localization, text encoding, UniversalDetector, NLangDetect |
| Security & Users | [168-emby-server-impl-security.md](./168-emby-server-impl-security.md) | Security, Session, Devices, DTOs, Activity, Browser, Branding |
| SharpCifs (Embedded) | [169-emby-server-impl-sharpcifs.md](./169-emby-server-impl-sharpcifs.md) | Embedded SMB/CIFS client library |

## Root Files

- `ApplicationHost.cs` — Emby.Server.Implementations/ApplicationHost.cs
- `ResourceFileManager.cs` — Emby.Server.Implementations/ResourceFileManager.cs
- `ServerApplicationPaths.cs` — Emby.Server.Implementations/ServerApplicationPaths.cs
- `StartupOptions.cs` — Emby.Server.Implementations/StartupOptions.cs
- `SystemEvents.cs` — Emby.Server.Implementations/SystemEvents.cs
