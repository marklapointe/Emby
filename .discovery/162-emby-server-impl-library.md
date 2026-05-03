# Component: Emby.Server.Implementations — Library Management

**Path:** \`Emby.Server.Implementations/\`
**Type:** Directory | Module Group
**Language:** C#
**Maps to:** \`.discovery/162-emby-server-impl-library.md\`

## Description

Library management components including media library core, item resolvers, validators, collections, playlists, user views, and sorting.

## Files

### Library/

- `CoreResolutionIgnoreRule.cs` — Emby.Server.Implementations/Library/CoreResolutionIgnoreRule.cs
- `DefaultAuthenticationProvider.cs` — Emby.Server.Implementations/Library/DefaultAuthenticationProvider.cs
- `ExclusiveLiveStream.cs` — Emby.Server.Implementations/Library/ExclusiveLiveStream.cs
- `LibraryManager.cs` — Emby.Server.Implementations/Library/LibraryManager.cs
- `LiveStreamHelper.cs` — Emby.Server.Implementations/Library/LiveStreamHelper.cs
- `MediaSourceManager.cs` — Emby.Server.Implementations/Library/MediaSourceManager.cs
- `MediaStreamSelector.cs` — Emby.Server.Implementations/Library/MediaStreamSelector.cs
- `MusicManager.cs` — Emby.Server.Implementations/Library/MusicManager.cs
- `PathExtensions.cs` — Emby.Server.Implementations/Library/PathExtensions.cs
- `ResolverHelper.cs` — Emby.Server.Implementations/Library/ResolverHelper.cs
- `AudioResolver.cs` — Emby.Server.Implementations/Library/Resolvers/Audio/AudioResolver.cs
- `MusicAlbumResolver.cs` — Emby.Server.Implementations/Library/Resolvers/Audio/MusicAlbumResolver.cs
- `MusicArtistResolver.cs` — Emby.Server.Implementations/Library/Resolvers/Audio/MusicArtistResolver.cs
- `BaseVideoResolver.cs` — Emby.Server.Implementations/Library/Resolvers/BaseVideoResolver.cs
- `BookResolver.cs` — Emby.Server.Implementations/Library/Resolvers/Books/BookResolver.cs
- `FolderResolver.cs` — Emby.Server.Implementations/Library/Resolvers/FolderResolver.cs
- `ItemResolver.cs` — Emby.Server.Implementations/Library/Resolvers/ItemResolver.cs
- `BoxSetResolver.cs` — Emby.Server.Implementations/Library/Resolvers/Movies/BoxSetResolver.cs
- `MovieResolver.cs` — Emby.Server.Implementations/Library/Resolvers/Movies/MovieResolver.cs
- `PhotoAlbumResolver.cs` — Emby.Server.Implementations/Library/Resolvers/PhotoAlbumResolver.cs
- `PhotoResolver.cs` — Emby.Server.Implementations/Library/Resolvers/PhotoResolver.cs
- `PlaylistResolver.cs` — Emby.Server.Implementations/Library/Resolvers/PlaylistResolver.cs
- `SpecialFolderResolver.cs` — Emby.Server.Implementations/Library/Resolvers/SpecialFolderResolver.cs
- `EpisodeResolver.cs` — Emby.Server.Implementations/Library/Resolvers/TV/EpisodeResolver.cs
- `SeasonResolver.cs` — Emby.Server.Implementations/Library/Resolvers/TV/SeasonResolver.cs
- `SeriesResolver.cs` — Emby.Server.Implementations/Library/Resolvers/TV/SeriesResolver.cs
- `VideoResolver.cs` — Emby.Server.Implementations/Library/Resolvers/VideoResolver.cs
- `SearchEngine.cs` — Emby.Server.Implementations/Library/SearchEngine.cs
- `UserDataManager.cs` — Emby.Server.Implementations/Library/UserDataManager.cs
- `UserManager.cs` — Emby.Server.Implementations/Library/UserManager.cs
- `UserViewManager.cs` — Emby.Server.Implementations/Library/UserViewManager.cs
- `ArtistsPostScanTask.cs` — Emby.Server.Implementations/Library/Validators/ArtistsPostScanTask.cs
- `ArtistsValidator.cs` — Emby.Server.Implementations/Library/Validators/ArtistsValidator.cs
- `GameGenresPostScanTask.cs` — Emby.Server.Implementations/Library/Validators/GameGenresPostScanTask.cs
- `GameGenresValidator.cs` — Emby.Server.Implementations/Library/Validators/GameGenresValidator.cs
- `GenresPostScanTask.cs` — Emby.Server.Implementations/Library/Validators/GenresPostScanTask.cs
- `GenresValidator.cs` — Emby.Server.Implementations/Library/Validators/GenresValidator.cs
- `MusicGenresPostScanTask.cs` — Emby.Server.Implementations/Library/Validators/MusicGenresPostScanTask.cs
- `MusicGenresValidator.cs` — Emby.Server.Implementations/Library/Validators/MusicGenresValidator.cs
- `PeopleValidator.cs` — Emby.Server.Implementations/Library/Validators/PeopleValidator.cs
- `StudiosPostScanTask.cs` — Emby.Server.Implementations/Library/Validators/StudiosPostScanTask.cs
- `StudiosValidator.cs` — Emby.Server.Implementations/Library/Validators/StudiosValidator.cs

### Collections/

- `CollectionImageProvider.cs` — Emby.Server.Implementations/Collections/CollectionImageProvider.cs
- `CollectionManager.cs` — Emby.Server.Implementations/Collections/CollectionManager.cs

### Playlists/

- `ManualPlaylistsFolder.cs` — Emby.Server.Implementations/Playlists/ManualPlaylistsFolder.cs
- `PlaylistImageProvider.cs` — Emby.Server.Implementations/Playlists/PlaylistImageProvider.cs
- `PlaylistManager.cs` — Emby.Server.Implementations/Playlists/PlaylistManager.cs

### UserViews/

- `CollectionFolderImageProvider.cs` — Emby.Server.Implementations/UserViews/CollectionFolderImageProvider.cs
- `DynamicImageProvider.cs` — Emby.Server.Implementations/UserViews/DynamicImageProvider.cs
- `FolderImageProvider.cs` — Emby.Server.Implementations/UserViews/FolderImageProvider.cs

