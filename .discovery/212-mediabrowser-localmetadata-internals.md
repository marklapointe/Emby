# MediaBrowser.LocalMetadata - Internals

**Module:** MediaBrowser.LocalMetadata
**Language:** C#
**Maps to:** `.discovery/212-mediabrowser-localmetadata-internals.md`

## Decomposition

### Providers/ (Local Metadata Providers)

#### Key Classes
`LocalMetadataProvider` (public class : ILocalMetadataProvider)
`LocalSeriesProvider` (public class : ILocalMetadataProvider)
`LocalMovieProvider` (public class : ILocalMetadataProvider)

#### Key Methods
```csharp
Task<LocalMetadataResult<T>> GetMetadata(...)
bool Supports(MetadataFields[] fields)
```

### Savers/ (Metadata Writers)

#### Key Classes
`XmlMetadataSaver` (public class : IMetadataSaver)
`JsonMetadataSaver` (public class : IMetadataSaver)

#### Key Methods
```csharp
Task SaveAsync(BaseItem item, CancellationToken cancellationToken)
string GetSavePath(BaseItem item)
```

### Parsers/ (Metadata Parsers)

#### Key Classes
`XmlParser` (public class)
`NfoParser` (public class)
`MovieNfoParser` (public class : NfoParser)
`SeriesNfoParser` (public class : NfoParser)

#### Key Methods
```csharp
T Parse(string path)
T Parse(Stream stream)
```

### Images/ (Local Image Provider)

#### Key Classes
`LocalImageProvider` (public class : ILocalImageProvider)

#### Key Methods
```csharp
Task<LocalImageResult> GetImages(...)
bool Supports(BaseItem item)
```

## File Listing

```
MediaBrowser.LocalMetadata/
├── Providers/        - Local metadata providers
│   ├── LocalMetadataProvider.cs
│   ├── LocalSeriesProvider.cs
│   ├── LocalMovieProvider.cs
│   └── LocalAlbumProvider.cs
├── Savers/           - Metadata writers
│   ├── XmlMetadataSaver.cs
│   ├── JsonMetadataSaver.cs
│   └── NfoMetadataSaver.cs
├── Parsers/          - Metadata parsers
│   ├── XmlParser.cs
│   ├── NfoParser.cs
│   ├── MovieNfoParser.cs
│   └── SeriesNfoParser.cs
└── Images/           - Local images
    └── LocalImageProvider.cs
```

## Description

LocalMetadata reads and writes metadata from/to local files. It supports NFO, XML, and JSON formats for storing media information alongside media files. Parsers extract metadata from these files, providers expose it to the system, and savers write changes back.

## Dependencies

- **MediaBrowser.Controller.Providers** - Provider interfaces
- **MediaBrowser.Model.Serialization** - JSON/XML

## Statistics

- **Files:** 15+
- **Lines:** ~2,000+
- **Classes:** 12+
