# Component: MediaBrowser.LocalMetadata — Expanded

**Path:** `MediaBrowser.LocalMetadata/`
**Type:** Directory | Module
**Language:** C#
**Maps to:** `.discovery/211-mediabrowser-localmetadata-internals.md`

## Description

Local metadata providers for reading and writing XML metadata from media files and folders. Parses NFO files and other local metadata formats.

## Files

### Root Files

- `BaseXmlProvider.cs` — MediaBrowser.LocalMetadata/BaseXmlProvider.cs

### Images/ (5 files)

- `CollectionFolderImageProvider.cs` — MediaBrowser.LocalMetadata/Images/CollectionFolderImageProvider.cs
- `EpisodeLocalImageProvider.cs` — MediaBrowser.LocalMetadata/Images/EpisodeLocalImageProvider.cs
- `InternalMetadataFolderImageProvider.cs` — MediaBrowser.LocalMetadata/Images/InternalMetadataFolderImageProvider.cs
- `LocalImageProvider.cs` — MediaBrowser.LocalMetadata/Images/LocalImageProvider.cs

### Parsers/ (6 files)

- `BaseItemXmlParser.cs` — MediaBrowser.LocalMetadata/Parsers/BaseItemXmlParser.cs
- `BoxSetXmlParser.cs` — MediaBrowser.LocalMetadata/Parsers/BoxSetXmlParser.cs
- `GameSystemXmlParser.cs` — MediaBrowser.LocalMetadata/Parsers/GameSystemXmlParser.cs
- `GameXmlParser.cs` — MediaBrowser.LocalMetadata/Parsers/GameXmlParser.cs
- `PlaylistXmlParser.cs` — MediaBrowser.LocalMetadata/Parsers/PlaylistXmlParser.cs

### Providers/ (4 files)

- `BoxSetXmlProvider.cs` — MediaBrowser.LocalMetadata/Providers/BoxSetXmlProvider.cs
- `GameSystemXmlProvider.cs` — MediaBrowser.LocalMetadata/Providers/GameSystemXmlProvider.cs
- `GameXmlProvider.cs` — MediaBrowser.LocalMetadata/Providers/GameXmlProvider.cs
- `PlaylistXmlProvider.cs` — MediaBrowser.LocalMetadata/Providers/PlaylistXmlProvider.cs

### Savers/ (7 files)

- `BaseXmlSaver.cs` — MediaBrowser.LocalMetadata/Savers/BaseXmlSaver.cs
- `BoxSetXmlSaver.cs` — MediaBrowser.LocalMetadata/Savers/BoxSetXmlSaver.cs
- `GameSystemXmlSaver.cs` — MediaBrowser.LocalMetadata/Savers/GameSystemXmlSaver.cs
- `GameXmlSaver.cs` — MediaBrowser.LocalMetadata/Savers/GameXmlSaver.cs
- `PersonXmlSaver.cs` — MediaBrowser.LocalMetadata/Savers/PersonXmlSaver.cs
- `PlaylistXmlSaver.cs` — MediaBrowser.LocalMetadata/Savers/PlaylistXmlSaver.cs

### Properties/ (1 file)

- `AssemblyInfo.cs` — MediaBrowser.LocalMetadata/Properties/AssemblyInfo.cs

## Data Flow

```mermaid
graph LR
    A[Media Files] --> B[LocalImageProvider]
    A --> C[XmlParser]
    B --> D[Images]
    C --> E[Metadata]
    E --> F[BaseXmlSaver]
    F --> G[XML Files]
```

## Supported Formats

| Format | Parser | Saver |
|--------|--------|-------|
| NFO (generic) | BaseItemXmlParser | BaseXmlSaver |
| Box Set | BoxSetXmlParser | BoxSetXmlSaver |
| Game | GameXmlParser | GameXmlSaver |
| Game System | GameSystemXmlParser | GameSystemXmlSaver |
| Playlist | PlaylistXmlParser | PlaylistXmlSaver |
| Person | - | PersonXmlSaver |

## Dependencies

- `MediaBrowser.Controller` — Base entity types
- `MediaBrowser.Model` — API models
