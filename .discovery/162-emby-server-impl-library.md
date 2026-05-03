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

## Decomposition

### LibraryManager.cs (Core Library Management)

#### Imports
```csharp
using MediaBrowser.Controller.Channels;
using MediaBrowser.Controller.Collections;
using MediaBrowser.Controller.Configuration;
using MediaBrowser.Controller.Entities;
using MediaBrowser.Controller.IO;
using MediaBrowser.Controller.Library;
using MediaBrowser.Model.Entities;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
```

#### Classes
`LibraryManager` (public class : ILibraryManager)

#### Key Properties
| Property | Type | Description |
|----------|------|-------------|
| `RootFolder` | `Folder` | Library root folder |
| `ItemById` | `ConcurrentDictionary<Guid, BaseItem>` | Items by ID |

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `GetItemInfo(BaseItem)` | `BaseItemInfo` | Get item metadata |
| `ValidateMediaLibrary()` | `Task` | Validate library |
| `GetMediaFolders()` | `IEnumerable<Folder>` | Get media folders |
| `AddVirtualFolder(string, bool)` | `Task` | Add library |
| `RemoveVirtualFolder(string, bool)` | `Task` | Remove library |

### UserManager.cs (User Management)

#### Classes
`UserManager` (public class : IUserManager)

#### Key Properties
| Property | Type | Description |
|----------|------|-------------|
| `Users` | `IEnumerable<User>` | All users |

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `GetUser(Guid)` | `User` | Get user by ID |
| `GetUserByName(string)` | `User` | Get user by name |
| `CreateUser(string)` | `Task<User>` | Create new user |
| `DeleteUser(User)` | `Task` | Delete user |

### ItemResolver.cs (Base Item Resolver)

#### Classes
`ItemResolver` (public abstract class)

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `ResolvePath(LibraryOptions, Folder, string)` | `ResolveResult<Item>` | Resolve item |

### CollectionManager.cs (Collection Management)

#### Classes
`CollectionManager` (public class : ICollectionManager)

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `CreateCollection(CollectionCreationOptions)` | `Task<BoxSet>` | Create collection |
| `AddToCollection(Guid, IEnumerable<Guid>)` | `Task` | Add items |
| `RemoveFromCollection(Guid, IEnumerable<Guid>)` | `Task` | Remove items |

### PlaylistManager.cs (Playlist Management)

#### Classes
`PlaylistManager` (public class)

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `GetPlaylists(Guid?)` | `IEnumerable<Playlist>` | Get playlists |
| `CreatePlaylist(PlaylistCreationOptions)` | `Task<Playlist>` | Create playlist |
| `AddToPlaylist(Guid, IEnumerable<Guid>)` | `Task` | Add items |

