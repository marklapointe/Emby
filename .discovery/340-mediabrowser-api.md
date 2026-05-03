# Component: MediaBrowser.Api

**Path:** \`MediaBrowser.Api/\`
**Type:** Directory | Module
**Language:** C#
**Maps to:** \`.discovery/340-mediabrowser-api.md\`

## Decomposition

### BaseApiService.cs (Base Class)

#### Imports
```csharp
using MediaBrowser.Controller.Dto;
using MediaBrowser.Controller.Entities;
using MediaBrowser.Controller.Library;
using MediaBrowser.Controller.Net;
using MediaBrowser.Controller.Session;
using MediaBrowser.Model.Entities;
using MediaBrowser.Model.Logging;
using MediaBrowser.Model.Services;
using MediaBrowser.Model.Extensions;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
```

#### Classes
\`BaseApiService\` (public class : IService, IRequiresRequest)

#### Properties
```csharp
public ILogger Logger
public IHttpResultFactory ResultFactory
public IRequest Request { get; set; }
```

#### Static Methods
```csharp
public static string[] SplitValue(string value, char delim)
public static Guid[] GetGuids(string value)
```

### ApiEntryPoint.cs (Service Locator)

#### Classes
\`ApiEntryPoint\` (sealed class)

#### Static Properties
```csharp
public static ApiEntryPoint Instance
public ILogger Logger
public IHttpResultFactory ResultFactory
```

### UserService.cs (User Management)

#### Methods
```csharp
public Task<object> Get(UserDto request)
public Task<object> DeleteUser(UserDto request)
public Task<object> UpdateUser(UserDto request)
```

### ItemUpdateService.cs (Item Updates)

#### Methods
```csharp
public Task<object> Post(UpdateUserItem request)
public Task<object> DeleteUserItem(DeleteUserItem request)
```

### ChannelService.cs (Channel Management)

#### Methods
```csharp
public Task<object> Get(ChannelFeatures request)
public Task<object> GetChannels(Channel request)
```

## Description

MediaBrowser.Api implements the REST API endpoints for Emby Server. It uses ServiceStack to expose HTTP endpoints for library, user, playback, and system management. Contains 63 C# files.

## API Areas

- `Devices/` — 1 C# files
- `Images/` — 4 C# files
- `Library/` — 2 C# files
- `LiveTv/` — 2 C# files
- `Movies/` — 3 C# files
- `Music/` — 2 C# files
- `Properties/` — 1 C# files
- `ScheduledTasks/` — 2 C# files
- `Session/` — 2 C# files
- `Subtitles/` — 1 C# files
- `System/` — 3 C# files
- `UserLibrary/` — 12 C# files

## Root Files

- `ApiEntryPoint.cs` — MediaBrowser.Api/ApiEntryPoint.cs
- `BaseApiService.cs` — MediaBrowser.Api/BaseApiService.cs
- `BrandingService.cs` — MediaBrowser.Api/BrandingService.cs
- `ChannelService.cs` — MediaBrowser.Api/ChannelService.cs
- `ConfigurationService.cs` — MediaBrowser.Api/ConfigurationService.cs
- `DisplayPreferencesService.cs` — MediaBrowser.Api/DisplayPreferencesService.cs
- `EnvironmentService.cs` — MediaBrowser.Api/EnvironmentService.cs
- `FilterService.cs` — MediaBrowser.Api/FilterService.cs
- `GamesService.cs` — MediaBrowser.Api/GamesService.cs
- `IHasDtoOptions.cs` — MediaBrowser.Api/IHasDtoOptions.cs
- `IHasItemFields.cs` — MediaBrowser.Api/IHasItemFields.cs
- `ItemLookupService.cs` — MediaBrowser.Api/ItemLookupService.cs
- `ItemRefreshService.cs` — MediaBrowser.Api/ItemRefreshService.cs
- `ItemUpdateService.cs` — MediaBrowser.Api/ItemUpdateService.cs
- `LocalizationService.cs` — MediaBrowser.Api/LocalizationService.cs
- `NewsService.cs` — MediaBrowser.Api/NewsService.cs
- `PackageService.cs` — MediaBrowser.Api/PackageService.cs
- `PlaylistService.cs` — MediaBrowser.Api/PlaylistService.cs
- `PluginService.cs` — MediaBrowser.Api/PluginService.cs
- `SearchService.cs` — MediaBrowser.Api/SearchService.cs
- `SimilarItemsHelper.cs` — MediaBrowser.Api/SimilarItemsHelper.cs
- `StartupWizardService.cs` — MediaBrowser.Api/StartupWizardService.cs
- `SuggestionsService.cs` — MediaBrowser.Api/SuggestionsService.cs
- `TvShowsService.cs` — MediaBrowser.Api/TvShowsService.cs
- `UserService.cs` — MediaBrowser.Api/UserService.cs
- `VideosService.cs` — MediaBrowser.Api/VideosService.cs

## Project Files

- `MediaBrowser.Api.csproj` — MediaBrowser.Api/MediaBrowser.Api.csproj
- `packages.config` — MediaBrowser.Api/packages.config
