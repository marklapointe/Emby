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
├── Emby.Notifications.csproj    # Project file
├── Core/                        # Core notification logic
│   ├── NotificationManager.cs   # Manager → [class] NotificationManager
│   └── ...                      
├── Services/                    # Notification service implementations
└── Properties/                  # Assembly info
```

## Key Classes

| Class | File | Purpose |
|-------|------|---------|
| `NotificationManager` | `Core/` | Routes notifications to services |

## Dependencies

- `MediaBrowser.Controller` — Notification interfaces
- `MediaBrowser.Model` — Notification types

## Decomposition

### Core/NotificationManager.cs (Notification Management)

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

## Side Effects

- Sends emails
- Makes webhook HTTP requests
- Queues notifications for delivery

## Reference

- `INotificationService` interface in `MediaBrowser.Controller`
