# Component: MediaBrowser.LocalMetadata

**Path:** `MediaBrowser.LocalMetadata/`
**Type:** Directory | Library
**Language:** C#
**Maps to:** `.discovery/132-local-metadata.md`

## Description

MediaBrowser.LocalMetadata provides XML-based local metadata providers and savers for BoxSets, GameSystems, Games, and Playlists. It also provides local image file providers that discover images alongside media files (e.g., `folder.jpg`, `poster.jpg`, `backdrop.jpg`).

## Structure

```
MediaBrowser.LocalMetadata/
├── MediaBrowser.LocalMetadata.csproj
├── Properties/
│   └── AssemblyInfo.cs            # Assembly metadata
├── BaseXmlProvider.cs             # Abstract base for XML providers
│   └── [class] BaseXmlProvider<T> : ILocalMetadataProvider<T>, IHasItemChangeMonitor, IHasOrder
│       ├── [method] public Task<MetadataResult<T>> GetMetadata(ItemLookupInfo info, CancellationToken cancellationToken)
│       │   ├── Looks for `{itemName}.xml` alongside media file
│       │   ├── If found: parses XML via corresponding XmlParser
│       │   └── Returns populated metadata result
│       ├── [method] public bool HasChanged(BaseItem item, IDirectoryService directoryService)
│       │   └── Returns true if XML file modification time changed
│       └── [property] public int Order
│           └── Returns 0 (highest priority)
├── Parsers/
│   ├── BaseItemXmlParser.cs       # Base XML parser
│   │   └── [class] BaseItemXmlParser<T>
│   │       ├── [method] public void Fetch(BaseItem item, string metadataFile, CancellationToken cancellationToken)
│   │       │   ├── Loads XML document
│   │       │   ├── Parses common fields (Title, OriginalTitle, SortTitle, Overview, etc.)
│   │       │   ├── Parses actors, genres, studios, tags
│   │       │   ├── Parses images (poster, backdrop, banner, logo, thumb)
│   │       │   └── Populates item properties
│   │       └── [method] protected virtual void FetchDataFromXmlNode(...)
│   │           └── Override point for type-specific parsing
│   ├── BoxSetXmlParser.cs         # BoxSet XML parser
│   │   └── [class] BoxSetXmlParser : BaseItemXmlParser<BoxSet>
│   ├── GameSystemXmlParser.cs     # GameSystem XML parser
│   │   └── [class] GameSystemXmlParser : BaseItemXmlParser<GameSystem>
│   ├── GameXmlParser.cs           # Game XML parser
│   │   └── [class] GameXmlParser : BaseItemXmlParser<Game>
│   │       └── Parses game-specific fields (GameSystem, Players, etc.)
│   └── PlaylistXmlParser.cs       # Playlist XML parser
│       └── [class] PlaylistXmlParser : BaseItemXmlParser<Playlist>
├── Providers/
│   ├── BoxSetXmlProvider.cs       # BoxSet XML provider
│   │   └── [class] BoxSetXmlProvider : BaseXmlProvider<BoxSet>
│   ├── GameSystemXmlProvider.cs   # GameSystem XML provider
│   │   └── [class] GameSystemXmlProvider : BaseXmlProvider<GameSystem>
│   ├── GameXmlProvider.cs         # Game XML provider
│   │   └── [class] GameXmlProvider : BaseXmlProvider<Game>
│   └── PlaylistXmlProvider.cs     # Playlist XML provider
│       └── [class] PlaylistXmlProvider : BaseXmlProvider<Playlist>
├── Savers/
│   ├── BaseXmlSaver.cs            # Abstract base for XML savers
│   │   └── [class] BaseXmlSaver : IMetadataFileSaver
│   │       ├── [method] public void Save(BaseItem item, CancellationToken cancellationToken)
│   │       │   ├── Generates XML from item properties
│   │       │   ├── Writes to `{itemName}.xml` alongside media file
│   │       │   └── Saves alongside media file
│   │       ├── [method] protected abstract List<string> GetTagsUsed(BaseItem item)
│   │       │   └── Returns list of XML tags to write
│   │       └── [method] protected virtual void WriteCustomElements(...)
│   │           └── Override point for type-specific XML elements
│   ├── BoxSetXmlSaver.cs          # BoxSet XML saver
│   │   └── [class] BoxSetXmlSaver : BaseXmlSaver
│   ├── GameSystemXmlSaver.cs      # GameSystem XML saver
│   │   └── [class] GameSystemXmlSaver : BaseXmlSaver
│   ├── GameXmlSaver.cs            # Game XML saver
│   │   └── [class] GameXmlSaver : BaseXmlSaver
│   ├── PlaylistXmlSaver.cs        # Playlist XML saver
│   │   └── [class] PlaylistXmlSaver : BaseXmlSaver
│   └── PersonXmlSaver.cs          # Person XML saver (commented out)
│       └── //public class PersonXmlSaver : BaseXmlSaver
└── Images/
    ├── LocalImageProvider.cs       # Local image file provider
    │   └── [class] LocalImageProvider : ILocalImageFileProvider, IHasOrder
    │       ├── [method] public Task<IEnumerable<LocalImageInfo>> GetImages(BaseItem item, IDirectoryService directoryService, CancellationToken cancellationToken)
    │       │   ├── Scans item directory for image files
    │       │   ├── Matches filenames to image types (poster, backdrop, banner, logo, thumb, disc, menu)
    │       │   └── Returns list of LocalImageInfo
    │       └── [method] public int GetImagePriority(ImageType type)
    │           └── Returns priority for image type
    ├── EpisodeLocalImageProvider.cs  # Episode local image provider
    │   └── [class] EpisodeLocalLocalImageProvider : ILocalImageFileProvider, IHasOrder
    │       └── Looks for episode-specific images (e.g., `{episodeName}-thumb.jpg`)
    ├── CollectionFolderImageProvider.cs  # Collection folder image provider
    │   └── [class] CollectionFolderLocalImageProvider : ILocalImageFileProvider, IHasOrder
    │       └── Looks for images in collection folder root
    └── InternalMetadataFolderImageProvider.cs  # Internal metadata image provider
        └── [class] InternalMetadataFolderImageProvider : ILocalImageFileProvider, IHasOrder
            └── Looks for images in internal metadata folder
```

## Supported XML Types

| Type | Provider | Parser | Saver | XML Filename |
|------|----------|--------|-------|--------------|
| BoxSet | BoxSetXmlProvider | BoxSetXmlParser | BoxSetXmlSaver | `{itemName}.xml` |
| GameSystem | GameSystemXmlProvider | GameSystemXmlParser | GameSystemXmlSaver | `{itemName}.xml` |
| Game | GameXmlProvider | GameXmlParser | GameXmlSaver | `{itemName}.xml` |
| Playlist | PlaylistXmlProvider | PlaylistXmlParser | PlaylistXmlSaver | `{itemName}.xml` |

## Local Image Filename Patterns

| Image Type | Filenames |
|------------|-----------|
| Primary | `poster.jpg`, `folder.jpg`, `cover.jpg` |
| Backdrop | `backdrop.jpg`, `fanart.jpg` |
| Banner | `banner.jpg` |
| Logo | `logo.jpg`, `clearlogo.jpg` |
| Thumb | `thumb.jpg`, `landscape.jpg` |
| Disc | `disc.jpg`, `cdart.jpg` |
| Menu | `menu.jpg` |

## Data Flow

```mermaid
graph TD
    A[Library scan] --&gt; B[BaseXmlProvider.GetMetadata]
    B --&gt; C[Find {item}.xml]
    C --&gt; D{XML exists?}
    D --&gt;|Yes| E[BaseItemXmlParser.Fetch]
    E --&gt; F[Parse XML fields]
    F --&gt; G[Populate item metadata]
    D --&gt;|No| H[Return empty result]
    I[BaseXmlSaver.Save] --&gt; J[Generate XML]
    J --&gt; K[Write to {item}.xml]
    L[LocalImageProvider.GetImages] --&gt; M[Scan directory]
    M --&gt; N[Match image filenames]
    N --&gt; O[Return LocalImageInfo]
```

## Side Effects

- Reads XML files via IFileSystem
- Writes XML files (savers)
- Scans directories for image files
- No external network calls
