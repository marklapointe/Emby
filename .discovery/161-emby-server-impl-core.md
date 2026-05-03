# Component: Emby.Server.Implementations — Core Infrastructure

**Path:** \`Emby.Server.Implementations/\`
**Type:** Directory | Module Group
**Language:** C#
**Maps to:** \`.discovery/161-emby-server-impl-core.md\`

## Description

Core infrastructure components of Emby.Server.Implementations covering application base, configuration, cryptography, data persistence, environment, diagnostics, logging, networking, reflection, serialization, security, threading, and XML handling.

## Files

### AppBase/

- `BaseApplicationPaths.cs` — Emby.Server.Implementations/AppBase/BaseApplicationPaths.cs
- `BaseConfigurationManager.cs` — Emby.Server.Implementations/AppBase/BaseConfigurationManager.cs
- `ConfigurationHelper.cs` — Emby.Server.Implementations/AppBase/ConfigurationHelper.cs

### Configuration/

- `ServerConfigurationManager.cs` — Emby.Server.Implementations/Configuration/ServerConfigurationManager.cs

### Cryptography/

- `CryptographyProvider.cs` — Emby.Server.Implementations/Cryptography/CryptographyProvider.cs

### Data/

- `BaseSqliteRepository.cs` — Emby.Server.Implementations/Data/BaseSqliteRepository.cs
- `CleanDatabaseScheduledTask.cs` — Emby.Server.Implementations/Data/CleanDatabaseScheduledTask.cs
- `ManagedConnection.cs` — Emby.Server.Implementations/Data/ManagedConnection.cs
- `SqliteDisplayPreferencesRepository.cs` — Emby.Server.Implementations/Data/SqliteDisplayPreferencesRepository.cs
- `SqliteExtensions.cs` — Emby.Server.Implementations/Data/SqliteExtensions.cs
- `SqliteItemRepository.cs` — Emby.Server.Implementations/Data/SqliteItemRepository.cs
- `SqliteUserDataRepository.cs` — Emby.Server.Implementations/Data/SqliteUserDataRepository.cs
- `SqliteUserRepository.cs` — Emby.Server.Implementations/Data/SqliteUserRepository.cs
- `TypeMapper.cs` — Emby.Server.Implementations/Data/TypeMapper.cs

### Diagnostics/

- `CommonProcess.cs` — Emby.Server.Implementations/Diagnostics/CommonProcess.cs
- `ProcessFactory.cs` — Emby.Server.Implementations/Diagnostics/ProcessFactory.cs

### EnvironmentInfo/

- `EnvironmentInfo.cs` — Emby.Server.Implementations/EnvironmentInfo/EnvironmentInfo.cs

### HttpClientManager/

- `HttpClientInfo.cs` — Emby.Server.Implementations/HttpClientManager/HttpClientInfo.cs
- `HttpClientManager.cs` — Emby.Server.Implementations/HttpClientManager/HttpClientManager.cs

### Images/

- `BaseDynamicImageProvider.cs` — Emby.Server.Implementations/Images/BaseDynamicImageProvider.cs

### Logging/

- `ConsoleLogger.cs` — Emby.Server.Implementations/Logging/ConsoleLogger.cs
- `SimpleLogManager.cs` — Emby.Server.Implementations/Logging/SimpleLogManager.cs
- `UnhandledExceptionWriter.cs` — Emby.Server.Implementations/Logging/UnhandledExceptionWriter.cs

### MediaEncoder/

- `EncodingManager.cs` — Emby.Server.Implementations/MediaEncoder/EncodingManager.cs

### Net/

- `DisposableManagedObjectBase.cs` — Emby.Server.Implementations/Net/DisposableManagedObjectBase.cs
- `IWebSocket.cs` — Emby.Server.Implementations/Net/IWebSocket.cs
- `SocketFactory.cs` — Emby.Server.Implementations/Net/SocketFactory.cs
- `UdpSocket.cs` — Emby.Server.Implementations/Net/UdpSocket.cs
- `WebSocketConnectEventArgs.cs` — Emby.Server.Implementations/Net/WebSocketConnectEventArgs.cs

### Networking/

- `BigIntegerExt.cs` — Emby.Server.Implementations/Networking/IPNetwork/BigIntegerExt.cs
- `IPAddressCollection.cs` — Emby.Server.Implementations/Networking/IPNetwork/IPAddressCollection.cs
- `IPNetworkCollection.cs` — Emby.Server.Implementations/Networking/IPNetwork/IPNetworkCollection.cs
- `IPNetwork.cs` — Emby.Server.Implementations/Networking/IPNetwork/IPNetwork.cs
- `NetworkManager.cs` — Emby.Server.Implementations/Networking/NetworkManager.cs

### News/

- `NewsEntryPoint.cs` — Emby.Server.Implementations/News/NewsEntryPoint.cs
- `NewsService.cs` — Emby.Server.Implementations/News/NewsService.cs

### Reflection/

- `AssemblyInfo.cs` — Emby.Server.Implementations/Reflection/AssemblyInfo.cs

### ScheduledTasks/

- `ChapterImagesTask.cs` — Emby.Server.Implementations/ScheduledTasks/ChapterImagesTask.cs
- `DailyTrigger.cs` — Emby.Server.Implementations/ScheduledTasks/DailyTrigger.cs
- `IntervalTrigger.cs` — Emby.Server.Implementations/ScheduledTasks/IntervalTrigger.cs
- `PeopleValidationTask.cs` — Emby.Server.Implementations/ScheduledTasks/PeopleValidationTask.cs
- `PluginUpdateTask.cs` — Emby.Server.Implementations/ScheduledTasks/PluginUpdateTask.cs
- `RefreshMediaLibraryTask.cs` — Emby.Server.Implementations/ScheduledTasks/RefreshMediaLibraryTask.cs
- `ScheduledTaskWorker.cs` — Emby.Server.Implementations/ScheduledTasks/ScheduledTaskWorker.cs
- `StartupTrigger.cs` — Emby.Server.Implementations/ScheduledTasks/StartupTrigger.cs
- `SystemEventTrigger.cs` — Emby.Server.Implementations/ScheduledTasks/SystemEventTrigger.cs
- `SystemUpdateTask.cs` — Emby.Server.Implementations/ScheduledTasks/SystemUpdateTask.cs
- `TaskManager.cs` — Emby.Server.Implementations/ScheduledTasks/TaskManager.cs
- `DeleteCacheFileTask.cs` — Emby.Server.Implementations/ScheduledTasks/Tasks/DeleteCacheFileTask.cs
- `DeleteLogFileTask.cs` — Emby.Server.Implementations/ScheduledTasks/Tasks/DeleteLogFileTask.cs
- `ReloadLoggerFileTask.cs` — Emby.Server.Implementations/ScheduledTasks/Tasks/ReloadLoggerFileTask.cs
- `WeeklyTrigger.cs` — Emby.Server.Implementations/ScheduledTasks/WeeklyTrigger.cs

### Serialization/

- `JsonSerializer.cs` — Emby.Server.Implementations/Serialization/JsonSerializer.cs
- `XmlSerializer.cs` — Emby.Server.Implementations/Serialization/XmlSerializer.cs

### Services/

- `HttpResult.cs` — Emby.Server.Implementations/Services/HttpResult.cs
- `RequestHelper.cs` — Emby.Server.Implementations/Services/RequestHelper.cs
- `ResponseHelper.cs` — Emby.Server.Implementations/Services/ResponseHelper.cs
- `ServiceController.cs` — Emby.Server.Implementations/Services/ServiceController.cs
- `ServiceExec.cs` — Emby.Server.Implementations/Services/ServiceExec.cs
- `ServiceHandler.cs` — Emby.Server.Implementations/Services/ServiceHandler.cs
- `ServiceMethod.cs` — Emby.Server.Implementations/Services/ServiceMethod.cs
- `ServicePath.cs` — Emby.Server.Implementations/Services/ServicePath.cs
- `StringMapTypeDeserializer.cs` — Emby.Server.Implementations/Services/StringMapTypeDeserializer.cs
- `UrlExtensions.cs` — Emby.Server.Implementations/Services/UrlExtensions.cs

### Session/

- `FirebaseSessionController.cs` — Emby.Server.Implementations/Session/FirebaseSessionController.cs
- `HttpSessionController.cs` — Emby.Server.Implementations/Session/HttpSessionController.cs
- `SessionManager.cs` — Emby.Server.Implementations/Session/SessionManager.cs
- `SessionWebSocketListener.cs` — Emby.Server.Implementations/Session/SessionWebSocketListener.cs
- `WebSocketController.cs` — Emby.Server.Implementations/Session/WebSocketController.cs

### Sorting/

- `AiredEpisodeOrderComparer.cs` — Emby.Server.Implementations/Sorting/AiredEpisodeOrderComparer.cs
- `AlbumArtistComparer.cs` — Emby.Server.Implementations/Sorting/AlbumArtistComparer.cs
- `AlbumComparer.cs` — Emby.Server.Implementations/Sorting/AlbumComparer.cs
- `AlphanumComparator.cs` — Emby.Server.Implementations/Sorting/AlphanumComparator.cs
- `ArtistComparer.cs` — Emby.Server.Implementations/Sorting/ArtistComparer.cs
- `CommunityRatingComparer.cs` — Emby.Server.Implementations/Sorting/CommunityRatingComparer.cs
- `CriticRatingComparer.cs` — Emby.Server.Implementations/Sorting/CriticRatingComparer.cs
- `DateCreatedComparer.cs` — Emby.Server.Implementations/Sorting/DateCreatedComparer.cs
- `DateLastMediaAddedComparer.cs` — Emby.Server.Implementations/Sorting/DateLastMediaAddedComparer.cs
- `DatePlayedComparer.cs` — Emby.Server.Implementations/Sorting/DatePlayedComparer.cs
- `GameSystemComparer.cs` — Emby.Server.Implementations/Sorting/GameSystemComparer.cs
- `IsFavoriteOrLikeComparer.cs` — Emby.Server.Implementations/Sorting/IsFavoriteOrLikeComparer.cs
- `IsFolderComparer.cs` — Emby.Server.Implementations/Sorting/IsFolderComparer.cs
- `IsPlayedComparer.cs` — Emby.Server.Implementations/Sorting/IsPlayedComparer.cs
- `IsUnplayedComparer.cs` — Emby.Server.Implementations/Sorting/IsUnplayedComparer.cs
- `NameComparer.cs` — Emby.Server.Implementations/Sorting/NameComparer.cs
- `OfficialRatingComparer.cs` — Emby.Server.Implementations/Sorting/OfficialRatingComparer.cs
- `PlayCountComparer.cs` — Emby.Server.Implementations/Sorting/PlayCountComparer.cs
- `PlayersComparer.cs` — Emby.Server.Implementations/Sorting/PlayersComparer.cs
- `PremiereDateComparer.cs` — Emby.Server.Implementations/Sorting/PremiereDateComparer.cs
- `ProductionYearComparer.cs` — Emby.Server.Implementations/Sorting/ProductionYearComparer.cs
- `RandomComparer.cs` — Emby.Server.Implementations/Sorting/RandomComparer.cs
- `RuntimeComparer.cs` — Emby.Server.Implementations/Sorting/RuntimeComparer.cs
- `SeriesSortNameComparer.cs` — Emby.Server.Implementations/Sorting/SeriesSortNameComparer.cs
- `SortNameComparer.cs` — Emby.Server.Implementations/Sorting/SortNameComparer.cs
- `StartDateComparer.cs` — Emby.Server.Implementations/Sorting/StartDateComparer.cs
- `StudioComparer.cs` — Emby.Server.Implementations/Sorting/StudioComparer.cs

### Threading/

- `CommonTimer.cs` — Emby.Server.Implementations/Threading/CommonTimer.cs
- `TimerFactory.cs` — Emby.Server.Implementations/Threading/TimerFactory.cs

### Updates/

- `InstallationManager.cs` — Emby.Server.Implementations/Updates/InstallationManager.cs

### Xml/

- `XmlReaderSettingsFactory.cs` — Emby.Server.Implementations/Xml/XmlReaderSettingsFactory.cs

### Root files

- `ApplicationHost.cs` — Emby.Server.Implementations/ApplicationHost.cs
- `ResourceFileManager.cs` — Emby.Server.Implementations/ResourceFileManager.cs
- `ServerApplicationPaths.cs` — Emby.Server.Implementations/ServerApplicationPaths.cs
- `StartupOptions.cs` — Emby.Server.Implementations/StartupOptions.cs
- `SystemEvents.cs` — Emby.Server.Implementations/SystemEvents.cs
