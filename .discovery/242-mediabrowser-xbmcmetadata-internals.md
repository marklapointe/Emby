# MediaBrowser.XbmcMetadata - Internals

**Module:** MediaBrowser.XbmcMetadata
**Language:** C#
**Maps to:** `.discovery/242-mediabrowser-xbmcmetadata-internals.md`

## Decomposition

### Providers/ (XBMCMetadata Providers)

#### Key Classes
`XbmcMetadataProvider` (public class : ILocalMetadataProvider)
`XbmcMovieProvider` (public class : IRemoteMetadataProvider)
`XbmcSeriesProvider` (public class : IRemoteMetadataProvider)

#### Key Methods
```csharp
Task<LocalMetadataResult<T>> GetMetadata(...)
Task<RemoteMetadataResult<T>> GetRemoteMetadata(...)
```

### Parsers/ (NFO Parsers)

#### Key Classes
`XbmcNfoParser` (public class)
`XbmcMovieNfoParser` (public class : XbmcNfoParser)
`XbmcSeriesNfoParser` (public class : XbmcNfoParser)

#### Key Methods
```csharp
T ParseFromFile(string path)
T ParseFromUrl(string url)
```

### Savers/ (Metadata Writers)

#### Key Classes
`XbmcMetadataSaver` (public class : IMetadataSaver)
`XbmcMovieSaver` (public class : XbmcMetadataSaver)
`XbmcSeriesSaver` (public class : XbmcMetadataSaver)

#### Key Methods
```csharp
Task SaveAsync(BaseItem item, CancellationToken cancellationToken)
string GetSavePath(BaseItem item)
```

### Configuration/ (Settings)

#### Key Classes
`XbmcMetadataOptions` (public class)

## File Listing

```
MediaBrowser.XbmcMetadata/
├── Providers/
│   ├── XbmcMetadataProvider.cs
│   ├── XbmcMovieProvider.cs
│   └── XbmcSeriesProvider.cs
├── Parsers/
│   ├── XbmcNfoParser.cs
│   ├── XbmcMovieNfoParser.cs
│   └── XbmcSeriesNfoParser.cs
├── Savers/
│   ├── XbmcMetadataSaver.cs
│   ├── XbmcMovieSaver.cs
│   └── XbmcSeriesSaver.cs
├── Configuration/
│   └── XbmcMetadataOptions.cs
└── Properties/
    └── AssemblyInfo.cs
```

## Description

XbmcMetadata provides XBMC/Kodi-compatible metadata storage for Emby. It reads and writes NFO files in the XBMC format, allowing Emby libraries to be shared with Kodi. Includes both local (NFO file) and remote (Kodi library) provider implementations.

## Dependencies

- **MediaBrowser.Controller.Providers** - Provider interfaces
- **MediaBrowser.Model.Serialization** - JSON/XML

## Statistics

- **Files:** 12
- **Lines:** ~2,000+
- **Classes:** 10
