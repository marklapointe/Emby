# Component: Emby.Server.Implementations — Security & User Management

**Path:** \`Emby.Server.Implementations/\`
**Type:** Directory | Module Group
**Language:** C#
**Maps to:** \`.discovery/168-emby-server-impl-security.md\`

## Description

Security, authentication, authorization, session management, device management, DTOs, activity tracking, and related user management components.

## Files

### Security/

- `AuthenticationRepository.cs` — Emby.Server.Implementations/Security/AuthenticationRepository.cs
- `EncryptionManager.cs` — Emby.Server.Implementations/Security/EncryptionManager.cs
- `MBLicenseFile.cs` — Emby.Server.Implementations/Security/MBLicenseFile.cs
- `PluginSecurityManager.cs` — Emby.Server.Implementations/Security/PluginSecurityManager.cs
- `RegRecord.cs` — Emby.Server.Implementations/Security/RegRecord.cs

### Session/

- `FirebaseSessionController.cs` — Emby.Server.Implementations/Session/FirebaseSessionController.cs
- `HttpSessionController.cs` — Emby.Server.Implementations/Session/HttpSessionController.cs
- `SessionManager.cs` — Emby.Server.Implementations/Session/SessionManager.cs
- `SessionWebSocketListener.cs` — Emby.Server.Implementations/Session/SessionWebSocketListener.cs
- `WebSocketController.cs` — Emby.Server.Implementations/Session/WebSocketController.cs

### Devices/

- `DeviceId.cs` — Emby.Server.Implementations/Devices/DeviceId.cs
- `DeviceManager.cs` — Emby.Server.Implementations/Devices/DeviceManager.cs

### Dto/

- `DtoService.cs` — Emby.Server.Implementations/Dto/DtoService.cs

### Activity/

- `ActivityLogEntryPoint.cs` — Emby.Server.Implementations/Activity/ActivityLogEntryPoint.cs
- `ActivityManager.cs` — Emby.Server.Implementations/Activity/ActivityManager.cs
- `ActivityRepository.cs` — Emby.Server.Implementations/Activity/ActivityRepository.cs

### Browser/

- `BrowserLauncher.cs` — Emby.Server.Implementations/Browser/BrowserLauncher.cs

### Branding/

- `BrandingConfigurationFactory.cs` — Emby.Server.Implementations/Branding/BrandingConfigurationFactory.cs

### Archiving/

- `ZipClient.cs` — Emby.Server.Implementations/Archiving/ZipClient.cs

## Decomposition

### SessionManager.cs (Session Management)

#### Imports
```csharp
using MediaBrowser.Controller.Net;
using MediaBrowser.Controller.Session;
using MediaBrowser.Model.Session;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
```

#### Classes
`SessionManager` (public class : ISessionManager)

#### Key Properties
| Property | Type | Description |
|----------|------|-------------|
| `Sessions` | `IEnumerable<SessionInfo>` | Active sessions |

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `CreateSession(string, string)` | `Task<SessionInfo>` | Create new session |
| `GetSession(Guid)` | `SessionInfo` | Get session by ID |
| `LogSessionActivity(Guid, Guid, string, string)` | `Task` | Log activity |
| `SendMessageToUser(Guid, string, string)` | `Task` | Send message |

### EncryptionManager.cs (Cryptographic Operations)

#### Classes
`EncryptionManager` (public class : IEncryptionManager)

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `ComputeMD5(byte[])` | `string` | MD5 hash |
| `ComputeSHA1(byte[])` | `string` | SHA1 hash |
| `ComputeHmac(string, string)` | `string` | HMAC |

### DeviceManager.cs (Device Management)

#### Classes
`DeviceManager` (public class : IDeviceManager)

#### Key Properties
| Property | Type | Description |
|----------|------|-------------|
| `Devices` | `IEnumerable<DeviceInfo>` | Registered devices |

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `GetDevice(string)` | `DeviceInfo` | Get device by ID |
| `RegisterDevice(DeviceInfo)` | `Task` | Register new device |

### ActivityManager.cs (Activity Logging)

#### Classes
`ActivityManager` (public class : IActivityManager)

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `GetActivityLog(DateTime?, DateTime?)` | `IEnumerable<ActivityLogEntry>` | Get log entries |
| `LogEntry(ActivityLogEntry)` | `Task` | Log activity |

### DtoService.cs (DTO Transformation)

#### Classes
`DtoService` (public class : IDtoService)

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `GetBaseItemDto(BaseItem, DtoOptions)` | `BaseItemDto` | Transform to DTO |
| `GetItemDtos(IEnumerable<BaseItem>, DtoOptions)` | `IEnumerable<BaseItemDto>` | Batch transform |

### ZipClient.cs (Archive Operations)

#### Classes
`ZipClient` (public interface)

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `Extract(string, string)` | `Task` | Extract archive |
| `Compress(string, string[])` | `Task` | Create archive |

