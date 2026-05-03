# Component: Emby.Server.Implementations — HTTP Server & Services

**Path:** \`Emby.Server.Implementations/\`
**Type:** Directory | Module Group
**Language:** C#
**Maps to:** \`.discovery/164-emby-server-impl-http.md\`

## Description

HTTP server implementation, services, and entry points.

## Files

### HttpServer/

- `FileWriter.cs` — Emby.Server.Implementations/HttpServer/FileWriter.cs
- `HttpListenerHost.cs` — Emby.Server.Implementations/HttpServer/HttpListenerHost.cs
- `HttpResultFactory.cs` — Emby.Server.Implementations/HttpServer/HttpResultFactory.cs
- `IHttpListener.cs` — Emby.Server.Implementations/HttpServer/IHttpListener.cs
- `LoggerUtils.cs` — Emby.Server.Implementations/HttpServer/LoggerUtils.cs
- `RangeRequestWriter.cs` — Emby.Server.Implementations/HttpServer/RangeRequestWriter.cs
- `ResponseFilter.cs` — Emby.Server.Implementations/HttpServer/ResponseFilter.cs
- `AuthorizationContext.cs` — Emby.Server.Implementations/HttpServer/Security/AuthorizationContext.cs
- `AuthService.cs` — Emby.Server.Implementations/HttpServer/Security/AuthService.cs
- `SessionContext.cs` — Emby.Server.Implementations/HttpServer/Security/SessionContext.cs
- `StreamWriter.cs` — Emby.Server.Implementations/HttpServer/StreamWriter.cs
- `WebSocketConnection.cs` — Emby.Server.Implementations/HttpServer/WebSocketConnection.cs

### EntryPoints/

- `AutomaticRestartEntryPoint.cs` — Emby.Server.Implementations/EntryPoints/AutomaticRestartEntryPoint.cs
- `ExternalPortForwarding.cs` — Emby.Server.Implementations/EntryPoints/ExternalPortForwarding.cs
- `KeepServerAwake.cs` — Emby.Server.Implementations/EntryPoints/KeepServerAwake.cs
- `LibraryChangedNotifier.cs` — Emby.Server.Implementations/EntryPoints/LibraryChangedNotifier.cs
- `RecordingNotifier.cs` — Emby.Server.Implementations/EntryPoints/RecordingNotifier.cs
- `RefreshUsersMetadata.cs` — Emby.Server.Implementations/EntryPoints/RefreshUsersMetadata.cs
- `ServerEventNotifier.cs` — Emby.Server.Implementations/EntryPoints/ServerEventNotifier.cs
- `StartupWizard.cs` — Emby.Server.Implementations/EntryPoints/StartupWizard.cs
- `SystemEvents.cs` — Emby.Server.Implementations/EntryPoints/SystemEvents.cs
- `UdpServerEntryPoint.cs` — Emby.Server.Implementations/EntryPoints/UdpServerEntryPoint.cs
- `UsageEntryPoint.cs` — Emby.Server.Implementations/EntryPoints/UsageEntryPoint.cs
- `UsageReporter.cs` — Emby.Server.Implementations/EntryPoints/UsageReporter.cs
- `UserDataChangeNotifier.cs` — Emby.Server.Implementations/EntryPoints/UserDataChangeNotifier.cs

