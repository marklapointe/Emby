# Component: MediaBrowser.LocalMetadata

**Path:** `MediaBrowser.LocalMetadata/`
**Type:** Directory | Library
**Language:** C#
**Maps to:** `.discovery/127-local-metadata.md`

## Description

MediaBrowser.LocalMetadata provides XML-based local metadata providers and savers for Emby. It reads and writes `.xml` metadata files alongside media files (e.g., `movie.xml`, `tvshow.xml`) and discovers local images (posters, backdrops, etc.). Supports BoxSets, Games, GameSystems, and Playlists.

## Structure

```
MediaBrowser.LocalMetadata/
├── MediaBrowser.LocalMetadata.csproj
├── BaseXmlProvider.cs             # Abstract base for XML metadata providers
│   └── [class] BaseXmlProvider<T> : ILocalMetadataProvider<T>, IHasItemChangeMonitor, IHasOrder
│       ├── [method] public Task<MetadataResult<T>> GetMetadata(ItemLookupInfo info, CancellationToken cancellationToken)
│       │   ├── Looks for `{itemName}.xml` alongside media file
│       │   ├── If found: parses XML via corresponding XmlParser
│       │   ├── If not found: returns empty result
│       │   └── Returns populated metadata result
│       ├── [method] public bool HasChanged(BaseItem item, IDirectoryService directoryService)
│       │   └── Returns true if XML file modification time changed
│       └── [method] public int Order
│           └── Returns 0 (highest priority)
├── Parsers/
│   ├── BaseItemXmlParser.cs       # Base XML parser
│   │   └── [class] BaseItemXmlParser<T>
│   │       ├── [method] public void Fetch(BaseItem item, string metadataFile, CancellationToken cancellationToken)
│   │       │   ├── Loads XML document
│   │       │   ├── Parses common fields (Title, Overview, Year, Rating, Genres, Tags, etc.)
│   │       │   ├── Parses actors/people with roles
│   │       │   ├── Parses images (poster, backdrop, logo, etc.)
│   │       │   └── Populates item properties
│   │       ├── [method] protected virtual void FetchDataFromXmlNode(...)
│   │       │   └── Override point for type-specific parsing
│   │       └── [method] protected void AddActor(BaseItem item, string name, string role, string type, string imagePath)
│   │           └── Adds person to item People list
│   ├── BoxSetXmlParser.cs         # BoxSet XML parser
│   │   └── [class] BoxSetXmlParser : BaseItemXmlParser<BoxSet>
│   │       └── Parses BoxSet-specific fields (DisplayOrder, etc.)
│   ├── GameSystemXmlParser.cs     # GameSystem XML parser
│   │   └── [class] GameSystemXmlParser : BaseItemXmlParser<GameSystem>
│   │       └── Parses GameSystem-specific fields
│   ├── GameXmlParser.cs           # Game XML parser
│   │   └── [class] GameXmlParser : BaseItemXmlParser<Game>
│   │       └── Parses Game-specific fields (Players, GameSystemId, etc.)
│   └── PlaylistXmlParser.cs       # Playlist XML parser
│       └── [class] PlaylistXmlParser : BaseItemXmlParser<Playlist>
│           └── Parses Playlist-specific fields (Shares, PlaylistMediaType, etc.)
├── Providers/
│   ├── BoxSetXmlProvider.cs       # BoxSet XML metadata provider
│   │   └── [class] BoxSetXmlProvider : BaseXmlProvider<BoxSet>
│   │       └── Looks for `boxset.xml` or `{foldername}.xml`
│   ├── GameSystemXmlProvider.cs   # GameSystem XML metadata provider
│   │   └── [class] GameSystemXmlProvider : BaseXmlProvider<GameSystem>
│   ├── GameXmlProvider.cs         # Game XML metadata provider
│   │   └── [class] GameXmlProvider : BaseXmlProvider<Game>
│   └── PlaylistXmlProvider.cs     # Playlist XML metadata provider
│       └── [class] PlaylistXmlProvider : BaseXmlProvider<Playlist>
├── Savers/
│   ├── BaseXmlSaver.cs            # Abstract base for XML metadata savers
│   │   └── [class] BaseXmlSaver : IMetadataFileSaver
│   │       ├── [method] public void Save(BaseItem item, CancellationToken cancellationToken)
│   │       │   ├── Generates XML from item properties
│   │       │   ├── Writes to `{itemName}.xml` or `metadata.xml`
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
│   ├── PersonXmlSaver.cs          # Person XML saver (commented out)
│   │   └── // [class] PersonXmlSaver : BaseXmlSaver
│   └── PlaylistXmlSaver.cs        # Playlist XML saver
│       └── [class] PlaylistXmlSaver : BaseXmlSaver
└── Images/
    ├── LocalImageProvider.cs      # Local image discovery
    │   └── [class] LocalImageProvider : ILocalImageFileProvider, IHasOrder
    │       ├── [method] public List<LocalImageInfo> GetImages(BaseItem item, IDirectoryService directoryService)
    │       │   ├── Scans item directory for image files
    │       │   ├── Recognizes naming patterns:
    │       │   │   ├── poster.jpg/png → Primary image
    │       │   │   ├── backdrop.jpg/png → Backdrop
    │       │   │   ├── logo.png → Logo
    │       │   │   ├── banner.jpg → Banner
    │       │   │   ├── thumb.jpg → Thumb
    │       │   │   ├── clearart.png → Art
    │       │   │   ├── disc.png → Disc
    │       │   │   └── folder.jpg → Primary
    │       │   └── Returns list of discovered images
    │       └── [method] public int Order
    │           └── Returns 0 (highest priority)
    ├── CollectionFolderImageProvider.cs # Collection folder images
    │   └── [class] CollectionFolderLocalImageProvider : ILocalImageFileProvider, IHasOrder
    ├── EpisodeLocalImageProvider.cs   # Episode image discovery
    │   └── [class] EpisodeLocalLocalImageProvider : ILocalImageFileProvider, IHasOrder
    │       └── Discovers episode-specific images (screenshots, etc.)
    └── InternalMetadataFolderImageProvider.cs # Internal metadata folder images
        └── [class] InternalMetadataFolderImageProvider : ILocalImageFileProvider, IHasOrder
            └── Discovers images from internal metadata folder
```

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
    I[LocalImageProvider] --&gt; J[Scan directory]
    J --&gt; K[Find poster/backdrop/etc.]
    K --&gt; L[Return image list]
```

## Supported Image Naming Patterns

| Filename | Image Type | Description |
|----------|-----------|-------------|
| `poster.jpg/png` | Primary | Main poster/cover |
| `backdrop.jpg/png` | Backdrop | Background image |
| `logo.png` | Logo | Channel/show logo |
| `banner.jpg` | Banner | Wide banner image |
| `thumb.jpg` | Thumb | Thumbnail |
| `clearart.png` | Art | Clear art overlay |
| `disc.png` | Disc | Disc/CD image |
| `folder.jpg` | Primary | Folder thumbnail |

## Side Effects

- Reads XML metadata files via IFileSystem
- Reads local image files
- Writes XML metadata files (savers)
- No external network calls
