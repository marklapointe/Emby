# Component: MediaBrowser.XbmcMetadata

**Path:** `MediaBrowser.XbmcMetadata/`
**Type:** Directory | Plugin
**Language:** C#
**Maps to:** `.discovery/135-xbmc-metadata.md`

## Description

MediaBrowser.XbmcMetadata provides XBMC/Kodi-compatible NFO file parsing and generation. NFO files are XML metadata files used by Kodi to store media metadata alongside media files. This plugin reads NFO files to populate Emby item properties and writes NFO files to maintain Kodi compatibility.

## Structure

```
MediaBrowser.XbmcMetadata/
├── MediaBrowser.XbmcMetadata.csproj
├── Properties/
│   └── AssemblyInfo.cs            # Assembly metadata
├── EntryPoint.cs                  # Plugin entry point
│   └── [class] EntryPoint : IServerEntryPoint
│       ├── [method] public void Run(IServerApplicationHost appHost, ILoggerFactory loggerFactory)
│       │   ├── Registers all NFO providers and savers
│       │   └── Registers NfoOptions configuration factory
│       └── [method] public void Dispose()
├── Configuration/
│   └── NfoOptions.cs              # NFO configuration options
│       └── [class] ConfigurationFactory : IConfigurationFactory
│           └── [method] public IEnumerable<ConfigurationInfo> GetConfigurations()
│               └── Returns NFO options (enable/disable NFO reading/writing)
├── Parsers/
│   ├── BaseNfoParser.cs           # Base NFO XML parser
│   │   └── [class] BaseNfoParser<T>
│   │       ├── [method] public void Fetch(BaseItem item, string metadataFile, CancellationToken cancellationToken)
│   │       │   ├── Loads NFO XML document
│   │       │   ├── Parses common fields (title, originaltitle, sorttitle, rating, year, etc.)
│   │       │   ├── Parses actors (name, role, type, thumb)
│   │       │   ├── Parses genres, studios, tags, countries
│   │       │   ├── Parses images (poster, fanart, banner, clearlogo, clearart, landscape, discart)
│   │       │   ├── Parses fileinfo (streamdetails: video, audio, subtitle)
│   │       │   └── Populates item properties
│   │       └── [method] protected virtual void FetchDataFromXmlNode(...)
│   │           └── Override point for type-specific parsing
│   ├── EpisodeNfoParser.cs        # Episode NFO parser
│   │   └── [class] EpisodeNfoParser : BaseNfoParser<Episode>
│   │       └── Parses episode-specific fields (season, episode, aired, displayseason, displayepisode)
│   ├── MovieNfoParser.cs          # Movie NFO parser
│   │   └── [class] MovieNfoParser : BaseNfoParser<Movie>
│   │       └── Parses movie-specific fields (mpaa, imdbid, tmdbid, set, tagline)
│   ├── SeasonNfoParser.cs         # Season NFO parser
│   │   └── [class] SeasonNfoParser : BaseNfoParser<Season>
│   │       └── Parses season-specific fields (seasonnumber, aired)
│   └── SeriesNfoParser.cs         # Series NFO parser
│       └── [class] SeriesNfoParser : BaseNfoParser<Series>
│           └── Parses series-specific fields (status, premiered, studio, episodeguide)
├── Providers/
│   ├── BaseNfoProvider.cs         # Abstract base for NFO providers
│   │   └── [class] BaseNfoProvider<T> : ILocalMetadataProvider<T>, IHasItemChangeMonitor
│   │       ├── [method] public Task<MetadataResult<T>> GetMetadata(ItemLookupInfo info, IDirectoryService directoryService, CancellationToken cancellationToken)
│   │       │   ├── Looks for `{itemName}.nfo` alongside media file
│   │       │   ├── If found: parses NFO via corresponding NfoParser
│   │       │   └── Returns populated metadata result
│   │       └── [method] public bool HasChanged(BaseItem item, IDirectoryService directoryService)
│   │           └── Returns true if NFO file modification time changed
│   ├── BaseVideoNfoProvider.cs    # Base for video NFO providers
│   │   └── [class] BaseVideoNfoProvider<T> : BaseNfoProvider<T>
│   ├── MovieNfoProvider.cs        # Movie NFO provider
│   │   ├── [class] MovieNfoProvider : BaseVideoNfoProvider<Movie>
│   │   ├── [class] MusicVideoNfoProvider : BaseVideoNfoProvider<MusicVideo>
│   │   └── [class] VideoNfoProvider : BaseVideoNfoProvider<Video>
│   ├── EpisodeNfoProvider.cs      # Episode NFO provider
│   │   └── [class] EpisodeNfoProvider : BaseNfoProvider<Episode>
│   ├── SeriesNfoProvider.cs       # Series NFO provider
│   │   └── [class] SeriesNfoProvider : BaseNfoProvider<Series>
│   ├── SeasonNfoProvider.cs       # Season NFO provider
│   │   └── [class] SeasonNfoProvider : BaseNfoProvider<Season>
│   ├── AlbumNfoProvider.cs        # Album NFO provider
│   │   └── [class] AlbumNfoProvider : BaseNfoProvider<MusicAlbum>
│   └── ArtistNfoProvider.cs       # Artist NFO provider
│       └── [class] ArtistNfoProvider : BaseNfoProvider<MusicArtist>
└── Savers/
    ├── BaseNfoSaver.cs            # Abstract base for NFO savers
    │   └── [class] BaseNfoSaver : IMetadataFileSaver
    │       ├── [method] public void Save(BaseItem item, CancellationToken cancellationToken)
    │       │   ├── Generates NFO XML from item properties
    │       │   ├── Writes to `{itemName}.nfo` alongside media file
    │       │   └── Saves alongside media file
    │       ├── [method] protected abstract List<string> GetTagsUsed(BaseItem item)
    │       │   └── Returns list of NFO XML tags to write
    │       └── [method] protected virtual void WriteCustomElements(...)
    │           └── Override point for type-specific XML elements
    ├── MovieNfoSaver.cs           # Movie NFO saver
    │   └── [class] MovieNfoSaver : BaseNfoSaver
    ├── EpisodeNfoSaver.cs         # Episode NFO saver
    │   └── [class] EpisodeNfoSaver : BaseNfoSaver
    ├── SeriesNfoSaver.cs          # Series NFO saver
    │   └── [class] SeriesNfoSaver : BaseNfoSaver
    ├── SeasonNfoSaver.cs          # Season NFO saver
    │   └── [class] SeasonNfoSaver : BaseNfoSaver
    ├── AlbumNfoSaver.cs           # Album NFO saver
    │   └── [class] AlbumNfoSaver : BaseNfoSaver
    └── ArtistNfoSaver.cs          # Artist NFO saver
        └── [class] ArtistNfoSaver : BaseNfoSaver
```

## Supported NFO Types

| Type | Provider | Parser | Saver | NFO Filename |
|------|----------|--------|-------|--------------|
| Movie | MovieNfoProvider | MovieNfoParser | MovieNfoSaver | `{itemName}.nfo` |
| MusicVideo | MusicVideoNfoProvider | MovieNfoParser | MovieNfoSaver | `{itemName}.nfo` |
| Video | VideoNfoProvider | MovieNfoParser | MovieNfoSaver | `{itemName}.nfo` |
| Episode | EpisodeNfoProvider | EpisodeNfoParser | EpisodeNfoSaver | `{itemName}.nfo` |
| Series | SeriesNfoProvider | SeriesNfoParser | SeriesNfoSaver | `tvshow.nfo` |
| Season | SeasonNfoProvider | SeasonNfoParser | SeasonNfoSaver | `season.nfo` |
| MusicAlbum | AlbumNfoProvider | BaseNfoParser | AlbumNfoSaver | `album.nfo` |
| MusicArtist | ArtistNfoProvider | BaseNfoParser | ArtistNfoSaver | `artist.nfo` |

## NFO XML Structure

```xml
<?xml version="1.0" encoding="UTF-8" standalone="yes" ?>
<movie>
  <title>Movie Title</title>
  <originaltitle>Original Title</originaltitle>
  <sorttitle>Sort Title</sorttitle>
  <set>Collection Name</set>
  <rating>8.5</rating>
  <year>2023</year>
  <mpaa>PG-13</mpaa>
  <plot>Movie overview...</plot>
  <tagline>Tagline...</tagline>
  <runtime>120</runtime>
  <thumb aspect="poster" preview="...">...poster.jpg</thumb>
  <fanart>
    <thumb>...backdrop.jpg</thumb>
  </fanart>
  <actor>
    <name>Actor Name</name>
    <role>Character</role>
    <type>Actor</type>
    <thumb>...actor.jpg</thumb>
  </actor>
  <genre>Action</genre>
  <studio>Studio Name</studio>
  <fileinfo>
    <streamdetails>
      <video>
        <codec>h264</codec>
        <aspect>1.78</aspect>
        <width>1920</width>
        <height>1080</height>
      </video>
      <audio>
        <codec>aac</codec>
        <channels>6</channels>
        <lang>en</lang>
      </audio>
    </streamdetails>
  </fileinfo>
</movie>
```

## Data Flow

```mermaid
graph TD
    A[Library scan] --&gt; B[BaseNfoProvider.GetMetadata]
    B --&gt; C[Find {item}.nfo]
    C --&gt; D{NFO exists?}
    D --&gt;|Yes| E[BaseNfoParser.Fetch]
    E --&gt; F[Parse XML fields]
    F --&gt; G[Populate item metadata]
    D --&gt;|No| H[Return empty result]
    I[BaseNfoSaver.Save] --&gt; J[Generate NFO XML]
    J --&gt; K[Write to {item}.nfo]
```

## Side Effects

- Reads NFO XML files via IFileSystem
- Writes NFO XML files (savers)
- No external network calls
