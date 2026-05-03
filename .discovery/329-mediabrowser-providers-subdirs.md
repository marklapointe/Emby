# MediaBrowser.Providers - Subdirectories

**Module:** MediaBrowser.Providers
**Language:** C#
**Maps to:** `.discovery/329-mediabrowser-providers-subdirs.md`

## Decomposition

### Manager/ (Provider Manager)

#### Key Classes
`ProviderManager` (public class : IProviderManager)

#### Key Methods
```csharp
Task<MetadataResult<T>> GetMetadata<T>(...)
Task RefreshFullItem(BaseItem item, MetadataRefreshOptions options)
void RegisterProvider(IProviderService provider)
```

### MediaInfo/ (Media Information)

#### Key Classes
`MediaInfoProvider` (public class : IMetadataProvider)

#### Key Methods
```csharp
Task<MetadataResult<T>> GetMetadata(...)
```

### Movies/ (Movie Providers)

#### Key Classes
`MovieMetadataService` (public class)
`MovieProviderService` (public class)

### TV/ (TV Providers)

#### Key Classes
`SeriesMetadataService` (public class)
`EpisodeMetadataService` (public class)

### Music/ (Music Providers)

#### Key Classes
`MusicMetadataService` (public class)
`AlbumProvider` (public class)

### Subtitles/ (Subtitle Providers)

#### Key Classes
`SubtitleProvider` (public class : ISubtitleProvider)

#### Key Methods
```csharp
Task<IEnumerable<SubtitleInfo>> GetSubtitles(...)
Task<SubtitleDownloadResult> DownloadSubtitle(...)
```

### People/ (People Providers)

#### Key Classes
`PersonMetadataService` (public class)

### Games/ (Game Providers)

#### Key Classes
`GameMetadataService` (public class)

### Photos/ (Photo Providers)

#### Key Classes
`PhotoProvider` (public class)

### Playlists/ (Playlist Providers)

#### Key Classes
`PlaylistProvider` (public class)

### Studios/ (Studio Providers)

#### Key Classes
`StudioProvider` (public class : IRemoteMetadataProvider)

### LiveTv/ (Live TV Providers)

#### Key Classes
`LiveTvMetadataService` (public class)

## File Listing

```
MediaBrowser.Providers/
├── Manager/         - Provider manager
├── MediaInfo/      - Media information
├── Movies/         - Movie metadata
├── TV/            - TV metadata
├── Music/         - Music metadata
├── Subtitles/     - Subtitle providers
├── People/        - Person metadata
├── Games/         - Game metadata
├── Photos/        - Photo metadata
├── Playlists/     - Playlist metadata
├── Studios/       - Studio metadata
├── LiveTv/        - Live TV metadata
├── Books/         - Book metadata
├── BoxSets/       - Box set metadata
├── Channels/      - Channel metadata
├── Chapters/      - Chapter metadata
├── Folders/       - Folder metadata
├── Genres/        - Genre metadata
├── MusicGenres/   - Music genre metadata
└── Years/         - Year metadata
```

## Statistics

- **Subdirectories:** 20
- **Total Providers:** 100+
