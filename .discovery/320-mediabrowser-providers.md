# Component: MediaBrowser.Providers

**Path:** \`MediaBrowser.Providers/\`
**Type:** Directory | Module
**Language:** C#
**Maps to:** \`.discovery/320-mediabrowser-providers.md\`

## Decomposition

### ProviderManager.cs (Central Manager)

#### Imports
```csharp
using MediaBrowser.Controller.Entities;
using MediaBrowser.Controller.Providers;
using MediaBrowser.Model.Entities;
using System;
using System.Collections.Generic;
using System.Threading;
using System.Threading.Tasks;
```

#### Classes
\`ProviderManager\` (public class : IProviderManager)

#### Key Methods
```csharp
Task<IEnumerable<RemoteSearchResult>> GetRemoteSearchResults<T>(T searchInfo, CancellationToken cancellationToken)
Task<MetadataRefreshResult> RefreshFullItem(BaseItem item, MetadataRefreshOptions options, CancellationToken cancellationToken)
void RegisterProvider(IProviderService provider)
```

### ItemMusicProvider.cs (Music Metadata)

#### Classes
\`ItemMusicProvider\` (public class : IMetadataProvider)

#### Key Methods
```csharp
Task<MetadataResult<Track>> GetMetadata(MusicProviderInfo id, CancellationToken cancellationToken)
```

### OmdbItemProvider.cs (OMDb Integration)

#### Imports
```csharp
using MediaBrowser.Controller.Entities;
using MediaBrowser.Model.Providers;
using System.Net.Http;
```

#### Classes
\`OmdbItemProvider\` (public class : IRemoteMetadataProvider)

#### Key Methods
```csharp
Task<RemoteMetadataResult<Movie>> GetMetadata(MovieInfo info, CancellationToken cancellationToken)
```

### TrailerFileProvider.cs (Trailer Extraction)

#### Classes
\`TrailerFileProvider\` (public class : ILocalMetadataProvider)

#### Key Methods
```csharp
Task<LocalMetadataResult<Trailer>> GetMetadata(LibraryOptions options, BaseItem item, CancellationToken cancellationToken)
```

## Description

MediaBrowser.Providers contains metadata providers for all media types. It fetches and aggregates metadata from external sources (OMDb, MusicBrainz, etc.) and local files. Contains 101 C# files across 20+ provider categories.

## Provider Categories

- `Books/` — 2 C# files
- `BoxSets/` — 3 C# files
- `Channels/` — 1 C# files
- `Chapters/` — 1 C# files
- `Folders/` — 3 C# files
- `GameGenres/` — 1 C# files
- `Games/` — 2 C# files
- `Genres/` — 1 C# files
- `LiveTv/` — 1 C# files
- `Manager/` — 10 C# files
- `MediaInfo/` — 8 C# files
- `Movies/` — 9 C# files
- `Music/` — 16 C# files
- `MusicGenres/` — 1 C# files
- `Omdb/` — 3 C# files
- `People/` — 4 C# files
- `Photos/` — 2 C# files
- `Playlists/` — 2 C# files
- `Properties/` — 1 C# files
- `Studios/` — 2 C# files

## Root Files


## Project Files

- `app.config` — MediaBrowser.Providers/app.config
- `MediaBrowser.Providers.csproj` — MediaBrowser.Providers/MediaBrowser.Providers.csproj
- `packages.config` — MediaBrowser.Providers/packages.config
