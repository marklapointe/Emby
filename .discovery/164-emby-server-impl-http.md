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

## Decomposition

### HttpListenerHost.cs (HTTP Server Host)

#### Imports
```csharp
using MediaBrowser.Model.Net;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
```

#### Classes
`HttpListenerHost` (public class : IHttpListener)

#### Key Properties
| Property | Type | Description |
|----------|------|-------------|
| `TotalRequests` | `long` | Request counter |
| `ActiveConnections` | `int` | Current connections |

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `Start()` | `Task` | Start HTTP listener |
| `Stop()` | `void` | Stop listener |
| `OnRequestReceived(Request)` | `Task` | Handle request |

### HttpResultFactory.cs (HTTP Response Factory)

#### Classes
`HttpResultFactory` (public class : IHttpResultFactory)

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `GetResult(Request)` | `HttpResult` | Create HTTP result |
| `GetStaticResult(Request, string)` | `Task<HttpResult>` | Static file response |

### WebSocketConnection.cs (WebSocket Handler)

#### Classes
`WebSocketConnection` (public class : IWebSocketConnection)

#### Key Events
| Event | Description |
|-------|-------------|
| `Receive` | Message received |
| `Send` | Message sent |

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `SendAsync(byte[])` | `Task` | Send binary message |
| `SendAsync(string)` | `Task` | Send text message |
| `Close()` | `Task` | Close connection |

### AuthService.cs (Authentication Service)

#### Classes
`AuthService` (public class : IService)

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `GetAuthenticatingUri(Request)` | `Task<string>` | Get auth URL |
| `Authenticate(Request)` | `Task<AuthResult>` | Authenticate request |

### ExternalPortForwarding.cs (UPnP Port Forwarding)

#### Classes
`ExternalPortForwarding` (public class)

#### Key Properties
| Property | Type | Description |
|----------|------|-------------|
| `IsMapped` | `bool` | Port forwarded |

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `MapPorts()` | `Task` | Create port mapping |
| `UnmapPorts()` | `Task` | Remove port mapping |

