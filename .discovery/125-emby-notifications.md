# Component: Emby.Notifications

**Path:** `Emby.Notifications/`
**Type:** Directory | Plugin
**Language:** C#
**Maps to:** `.discovery/125-emby-notifications.md`

## Description

Emby.Notifications is a server plugin that manages user notifications. It provides a notification manager, API endpoints for reading/marking notifications, and integration points for notification providers (email, push, webhooks, etc.).

## Structure

```
Emby.Notifications/
├── Emby.Notifications.csproj
├── Notifications.cs               # Plugin entry point
│   └── [class] Notifications : IServerEntryPoint
│       ├── [method] public void Run()
│       │   └── Initializes notification system on server startup
│       └── [method] public void Dispose()
│           └── Cleans up notification resources
├── NotificationManager.cs           # Core notification manager
│   └── [class] NotificationManager : INotificationManager
│       ├── [field] List<INotificationService> _notificationServices
│       ├── [constructor] NotificationManager(ILogger, IServerConfigurationManager, IJsonSerializer, IApplicationPaths, IFileSystem, IUserManager, ILibraryManager, ISessionManager, ILocalizationManager, IDeviceManager)
│       ├── [method] public void AddParts(IEnumerable<INotificationService> services)
│       │   └── Registers notification provider plugins
│       ├── [method] public Task SendNotification(NotificationRequest request, CancellationToken cancellationToken)
│       │   ├── Validates request
│       │   ├── Finds matching notification services
│       │   └── Sends to all registered providers
│       ├── [method] public NotificationOption[] GetNotificationOptions()
│       │   └── Returns available notification configuration options
│       └── [method] public NameIdPair[] GetNotificationServices()
│           └── Returns list of registered notification service names
├── CoreNotificationTypes.cs         # Built-in notification type definitions
│   └── [class] CoreNotificationTypes : INotificationTypeFactory
│       └── [method] public NotificationTypeInfo[] GetNotificationTypes()
│           └── Returns core types: PlaybackStart, PlaybackStop, VideoPlayback, AudioPlayback, etc.
├── NotificationConfigurationFactory.cs # Configuration UI factory
│   └── [class] NotificationConfigurationFactory : IConfigurationFactory
│       └── [method] public IEnumerable<ConfigurationInfo> GetConfigurations()
│           └── Provides notification settings page in dashboard
└── Api/
    └── NotificationsService.cs      # REST API endpoints
        └── [class] NotificationsService : IService
            ├── [class] GetNotifications : IReturn<NotificationResult>
            │   └── GET /Notifications — list user notifications
            ├── [class] Notification
            │   └── Notification DTO (Id, UserId, Date, Name, Description, Url, Level, IsRead)
            ├── [class] NotificationResult
            │   └── Paginated result (Items[], TotalRecordCount)
            ├── [class] NotificationsSummary
            │   └── Summary DTO (UnreadCount)
            ├── [class] GetNotificationsSummary : IReturn<NotificationsSummary>
            │   └── GET /Notifications/Summary — unread count
            ├── [class] GetNotificationTypes : IReturn<List<NotificationTypeInfo>>
            │   └── GET /Notifications/Types — available notification types
            ├── [class] GetNotificationServices : IReturn<List<NameIdPair>>
            │   └── GET /Notifications/Services — registered services
            ├── [class] AddAdminNotification : IReturnVoid
            │   └── POST /Notifications/Admin — send admin notification
            ├── [class] MarkRead : IReturnVoid
            │   └── POST /Notifications/{Id}/Read — mark as read
            └── [class] MarkUnread : IReturnVoid
                └── POST /Notifications/{Id}/Unread — mark as unread
```

## Data Flow

```mermaid
graph TD
    A[Event: PlaybackStart/Stop] --&gt; B[NotificationManager.SendNotification]
    B --&gt; C[Find matching services]
    C --&gt; D[EmailService]
    C --&gt; E[PushService]
    C --&gt; F[WebhookService]
    D --&gt; G[User receives notification]
    E --&gt; G
    F --&gt; G
    H[Client] --&gt; I[GET /Notifications]
    I --&gt; J[NotificationsService]
    J --&gt; K[Return notification list]
```

## API Endpoints

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/Notifications` | GET | List notifications |
| `/Notifications/Summary` | GET | Unread count summary |
| `/Notifications/Types` | GET | Available notification types |
| `/Notifications/Services` | GET | Registered notification services |
| `/Notifications/Admin` | POST | Send admin notification |
| `/Notifications/{Id}/Read` | POST | Mark notification as read |
| `/Notifications/{Id}/Unread` | POST | Mark notification as unread |

## Side Effects

- Reads/Writes notification data to server data store
- Sends notifications via external services (email, push)
- Integrates with IUserManager, ISessionManager, ILibraryManager
