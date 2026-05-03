# Component: MediaBrowser.XbmcMetadata

**Path:** `MediaBrowser.XbmcMetadata/`
**Type:** Directory | Library
**Language:** C#
**Maps to:** `.discovery/128-xbmc-metadata.md`

## Description

MediaBrowser.XbmcMetadata provides NFO (XML) metadata providers and savers compatible with Kodi/XBMC format. Reads and writes `.nfo` files alongside media files (e.g., `movie.nfo`, `tvshow.nfo`) for Movies, Series, Episodes, Seasons, MusicAlbums, and MusicArtists.

## Structure

```
MediaBrowser.XbmcMetadata/
├── MediaBrowser.XbmcMetadata.csproj
├── EntryPoint.cs                  # Plugin entry point
│   └── [class] EntryPoint : IServerEntryPoint
│       ├── [method] public void Run()
│       │   └── Registers NFO providers and savers
│       └── [method] public void Dispose()
├── Configuration/
│   └── NfoOptions.cs              # Dashboard configuration
│       └── [class] ConfigurationFactory : IConfigurationFactory
│           └── Provides NFO settings page (enable/disable NFO reading/writing)
├── Parsers/
│   ├── BaseNfoParser.cs           # Base NFO XML parser
│   │   └── [class] BaseNfoParser<T>
│   │       ├── [method] public void Fetch(BaseItem item, string metadataFile, CancellationToken cancellationToken)
│   │       │   ├── Loads NFO XML document
│   │       │   ├── Parses common fields (Title, OriginalTitle, SortTitle, Overview, Tagline, etc.)
│   │       │   ├── Parses actors with roles and images
│   │       │   ├── Parses genres, tags, studios, countries
│   │       │   ├── Parses ratings (IMDb, TMDB, etc.)
│   │       │   ├── Parses images (poster, fanart, banner, clearlogo, etc.)
│   │       │   └── Populates item properties
│   │       ├── [method] protected virtual void FetchDataFromXmlNode(...)
│   │       │   └── Override point for type-specific parsing
│   │       └── [method] protected void AddActor(BaseItem item, string name, string role, string type, string imagePath)
│   ├── EpisodeNfoParser.cs        # Episode NFO parser
│   │   └── [class] EpisodeNfoParser : BaseNfoParser<Episode>
│   │       └── Parses episode-specific fields (EpisodeNumber, SeasonNumber, Aired, etc.)
│   ├── MovieNfoParser.cs          # Movie NFO parser
│   │   └── [class] MovieNfoParser : BaseNfoParser<Movie>
│   │       └── Parses movie-specific fields (Runtime, Budget, Revenue, etc.)
│   ├── SeasonNfoParser.cs         # Season NFO parser
│   │   └── [class] SeasonNfoParser : BaseNfoParser<Season>
│   │       └── Parses season-specific fields (SeasonNumber, Aired, etc.)
│   └── SeriesNfoParser.cs         # Series NFO parser
│       └── [class] SeriesNfoParser : BaseNfoParser<Series>
│           └── Parses series-specific fields (Status, AirDays, AirTime, etc.)
├── Providers/
│   ├── BaseNfoProvider.cs         # Abstract base for NFO providers
│   │   └── [class] BaseNfoProvider<T> : ILocalMetadataProvider<T>, IHasItemChangeMonitor
│   │       ├── [method] public Task<MetadataResult<T>> GetMetadata(ItemLookupInfo info, CancellationToken cancellationToken)
│   │       │   ├── Looks for `{itemName}.nfo` alongside media file
│   │       │   ├── If found: parses NFO via corresponding NfoParser
│   │       │   └── Returns populated metadata result
│   │       └── [method] public bool HasChanged(BaseItem item, IDirectoryService directoryService)
│   │           └── Returns true if NFO file modification time changed
│   ├── BaseVideoNfoProvider.cs    # Base for video NFO providers
│   │   └── [class] BaseVideoNfoProvider<T> : BaseNfoProvider<T>
│   │       └── Adds video-specific metadata handling
│   ├── AlbumNfoProvider.cs        # Album NFO provider
│   │   └── [class] AlbumNfoProvider : BaseNfoProvider<MusicAlbum>
│   ├── ArtistNfoProvider.cs       # Artist NFO provider
│   │   └── [class] ArtistNfoProvider : BaseNfoProvider<MusicArtist>
│   ├── EpisodeNfoProvider.cs      # Episode NFO provider
│   │   └── [class] EpisodeNfoProvider : BaseNfoProvider<Episode>
│   ├── MovieNfoProvider.cs        # Movie NFO provider
│   │   ├── [class] MovieNfoProvider : BaseVideoNfoProvider<Movie>
│   │   ├── [class] MusicVideoNfoProvider : BaseVideoNfoProvider<MusicVideo>
│   │   └── [class] VideoNfoProvider : BaseVideoNfoProvider<Video>
│   ├── SeasonNfoProvider.cs       # Season NFO provider
│   │   └── [class] SeasonNfoProvider : BaseNfoProvider<Season>
│   └── SeriesNfoProvider.cs       # Series NFO provider
│       └── [class] SeriesNfoProvider : BaseNfoProvider<Series>
└── Savers/
    ├── BaseNfoSaver.cs            # Abstract base for NFO savers
    │   └── [class] BaseNfoSaver : IMetadataFileSaver
    │       ├── [method] public void Save(BaseItem item, CancellationToken cancellationToken)
    │       │   ├── Generates NFO XML from item properties
    │       │   ├── Writes to `{itemName}.nfo` alongside media file
    │       │   └── Saves alongside media file
    │       ├── [method] protected abstract List<string> GetTagsUsed(BaseItem item)
    │       │   └── Returns list of NFO tags to write
    │       └── [method] protected virtual void WriteCustomElements(...)
    │           └── Override point for type-specific NFO elements
    ├── AlbumNfoSaver.cs           # Album NFO saver
    │   └── [class] AlbumNfoSaver : BaseNfoSaver
    ├── ArtistNfoSaver.cs          # Artist NFO saver
    │   └── [class] ArtistNfoSaver : BaseNfoSaver
    ├── EpisodeNfoSaver.cs         # Episode NFO saver
    │   └── [class] EpisodeNfoSaver : BaseNfoSaver
    ├── MovieNfoSaver.cs           # Movie NFO saver
    │   └── [class] MovieNfoSaver : BaseNfoSaver
    ├── SeasonNfoSaver.cs          # Season NFO saver
    │   └── [class] SeasonNfoSaver : BaseNfoSaver
    └── SeriesNfoSaver.cs          # Series NFO saver
        └── [class] SeriesNfoSaver : BaseNfoSaver
```

## Data Flow

```mermaid
graph TD
    A[Library scan] --&gt; B[BaseNfoProvider.GetMetadata]
    B --&gt; C[Find {item}.nfo]
    C --&gt; D{NFO exists?}
    D --&gt;|Yes| E[BaseNfoParser.Fetch]
    E --&gt; F[Parse NFO XML fields]
    F --&gt; G[Populate item metadata]
    D --&gt;|No| H[Return empty result]
    I[BaseNfoSaver.Save] --&gt; J[Generate NFO XML]
    J --&gt; K[Write to {item}.nfo]
```

## Supported NFO Types

| Type | Provider | Parser | Saver | NFO Filename |
|------|----------|--------|-------|--------------|
| Movie | MovieNfoProvider | MovieNfoParser | MovieNfoSaver | `movie.nfo` or `{filename}.nfo` |
| Series | SeriesNfoProvider | SeriesNfoParser | SeriesNfoSaver | `tvshow.nfo` |
| Episode | EpisodeNfoProvider | EpisodeNfoParser | EpisodeNfoSaver | `{filename}.nfo` |
| Season | SeasonNfoProvider | SeasonNfoParser | SeasonNfoSaver | `season.nfo` |
| MusicAlbum | AlbumNfoProvider | — | AlbumNfoSaver | `album.nfo` |
| MusicArtist | ArtistNfoProvider | — | ArtistNfoSaver | `artist.nfo` |

## Side Effects

- Reads NFO XML files via IFileSystem
- Writes NFO XML files (savers)
- No external network calls
