# Component: Emby.Notifications

**Path:** `Emby.Notifications/`
**Type:** Directory | Module
**Language:** C#
**Maps to:** `.discovery/140-emby-notifications.md`

## Description

Emby.Notifications provides the notification infrastructure for Emby Server. It supports sending alerts via multiple channels (email, webhooks, mobile push) and allows plugins to register custom notification services.

## Structure

```
Emby.Notifications/
├── Emby.Notifications.csproj                    # Project file
├── NotificationManager.cs                        # Main notification manager
├── CoreNotificationTypes.cs                      # Core notification types
├── NotificationConfigurationFactory.cs            # Configuration factory
├── Notifications.cs                              # Notifications integration
├── Api/
│   └── NotificationsService.cs                   # REST API endpoints
└── Properties/
    └── AssemblyInfo.cs                          # Assembly metadata
```

## Files

| File | Purpose |
|------|---------|
| `NotificationManager.cs` | Main notification manager - routes notifications to services |
| `CoreNotificationTypes.cs` | Defines all built-in notification types (plugin, system, playback, user) |
| `NotificationConfigurationFactory.cs` | Configuration store for notification settings |
| `Notifications.cs` | Entry point for library update notifications |
| `Api/NotificationsService.cs` | REST API for notifications (GET/POST endpoints) |
| `Properties/AssemblyInfo.cs` | Assembly metadata |

## Dependencies

- `MediaBrowser.Controller` — Notification interfaces
- `MediaBrowser.Model` — Notification types

## Decomposition

### NotificationManager.cs (Main Notification Manager)

#### Imports
```csharp
using MediaBrowser.Common.Implementations;
using MediaBrowser.Controller.Notifications;
using MediaBrowser.Model.Logging;
using MediaBrowser.Model.Notifications;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
```

#### Classes
`NotificationManager` (public class : INotificationManager)

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `QueueNotification(NotificationRequest)` | `Task` | Queue a notification for delivery |
| `GetNotificationTypes()` | `IEnumerable<NotificationTypeDescription>` | Get available notification types |
| `GetNotificationServices()` | `IEnumerable<NotificationServiceInfo>` | Get registered services |

### CoreNotificationTypes.cs (Core Notification Types)

#### Imports
```csharp
using MediaBrowser.Controller;
using MediaBrowser.Controller.Notifications;
using MediaBrowser.Model.Notifications;
using System;
using System.Collections.Generic;
using System.Linq;
using MediaBrowser.Model.Globalization;
```

#### Classes
`CoreNotificationTypes` (public class : INotificationTypeFactory)

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `GetNotificationTypes()` | `IEnumerable<NotificationTypeInfo>` | Returns all core notification types |

#### Notification Types Defined
- `ApplicationUpdateInstalled` / `ApplicationUpdateAvailable`
- `InstallationFailed`
- `PluginInstalled` / `PluginError` / `PluginUninstalled` / `PluginUpdateInstalled`
- `ServerRestartRequired`
- `TaskFailed`
- `NewLibraryContent`
- `AudioPlayback` / `VideoPlayback` / `GamePlayback` (and Stopped variants)
- `CameraImageUploaded`
- `UserLockedOut`

### NotificationConfigurationFactory.cs (Configuration)

#### Imports
```csharp
using MediaBrowser.Common.Configuration;
using MediaBrowser.Model.Notifications;
using System.Collections.Generic;
```

#### Classes
`NotificationConfigurationFactory` (public class : IConfigurationFactory)

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `GetConfigurations()` | `IEnumerable<ConfigurationStore>` | Returns NotificationOptions config store |

### Api/NotificationsService.cs (REST API)

#### Imports
```csharp
using MediaBrowser.Controller.Library;
using MediaBrowser.Controller.Net;
using MediaBrowser.Controller.Notifications;
using MediaBrowser.Model.Notifications;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading;
using System.Threading.Tasks;
using MediaBrowser.Model.Services;
using MediaBrowser.Model.Dto;
```

#### Classes
`NotificationsService` (public class : IService)

#### API Endpoints
| Route | Method | Description |
|-------|--------|-------------|
| `/Notifications/{UserId}` | GET | Gets notifications for user |
| `/Notifications/{UserId}/Summary` | GET | Gets notification summary |
| `/Notifications/Types` | GET | Gets notification types |
| `/Notifications/Services` | GET | Gets notification services |
| `/Notifications/Admin` | POST | Sends admin notification |
| `/Notifications/{UserId}/Read` | POST | Marks notifications as read |
| `/Notifications/{UserId}/Unread` | POST | Marks notifications as unread |

## Side Effects

- Sends emails
- Makes webhook HTTP requests
- Queues notifications for delivery
- Updates notification read/unread status

## Reference

- `INotificationService` interface in `MediaBrowser.Controller`
