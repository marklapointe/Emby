# Component: Emby.Server.Implementations — Sorting

**Path:** `Emby.Server.Implementations/Sorting/`
**Type:** Directory | Module
**Language:** C#
**Maps to:** `.discovery/216-emby-server-impl-sorting.md`

## Description

Custom sort comparers for library items. Provides specialized sorting algorithms for different media types and criteria.

## Files

### Root Sorting Files

- `AlbumArtistComparer.cs` — Emby.Server.Implementations/Sorting/AlbumArtistComparer.cs
- `AlbumComparer.cs` — Emby.Server.Implementations/Sorting/AlbumComparer.cs
- `ArtistComparer.cs` — Emby.Server.Implementations/Sorting/ArtistComparer.cs
- `BaseComparer.cs` — Emby.Server.Implementations/Sorting/BaseComparer.cs
- `BoxSetComparer.cs` — Emby.Server.Implementations/Sorting/BoxSetComparer.cs
- `DirectorComparer.cs` — Emby.Server.Implementations/Sorting/DirectorComparer.cs
- `EpisodeComparer.cs` — Emby.Server.Implementations/Sorting/EpisodeComparer.cs
- `GameComparer.cs` — Emby.Server.Implementations/Sorting/GameComparer.cs
- `GenreComparer.cs` — Emby.Server.Implementations/Sorting/GenreComparer.cs
- `ImplicitPlayerComparer.cs` — Emby.Server.Implementations/Sorting/ImplicitPlayerComparer.cs
- `IsFavoriteOrLikesComparer.cs` — Emby.Server.Implementations/Sorting/IsFavoriteOrLikesComparer.cs
- `IsFolderComparer.cs` — Emby.Server.Implementations/Sorting/IsFolderComparer.cs
- `IsPlayedComparer.cs` — Emby.Server.Implementations/Sorting/IsPlayedComparer.cs
- `LastPlaybackComparer.cs` — Emby.Server.Implementations/Sorting/LastPlaybackComparer.cs
- `MediaComparer.cs` — Emby.Server.Implementations/Sorting/MediaComparer.cs
- `Mp4BayItemInstallerComparer.cs` — Emby.Server.Implementations/Sorting/Mp4BayItemInstallerComparer.cs
- `MusicAlbumReleaseDateComparer.cs` — Emby.Server.Implementations/Sorting/MusicAlbumReleaseDateComparer.cs
- `NameComparer.cs` — Emby.Server.Implementations/Sorting/NameComparer.cs
- `OfficialRatingComparer.cs` — Emby.Server.Implementations/Sorting/OfficialRatingComparer.cs
- `ParentalRatingComparer.cs` — Emby.Server.Implementations/Sorting/ParentalRatingComparer.cs
- `PlayCountComparer.cs` — Emby.Server.Implementations/Sorting/PlayCountComparer.cs
- `PremiereDateComparer.cs` — Emby.Server.Implementations/Sorting/PremiereDateComparer.cs
- `RandomComparer.cs` — Emby.Server.Implementations/Sorting/RandomComparer.cs
- `RuntimeComparer.cs` — Emby.Server.Implementations/Sorting/RuntimeComparer.cs
- `SeasonComparer.cs` — Emby.Server.Implementations/Sorting/SeasonComparer.cs
- `SortHelper.cs` — Emby.Server.Implementations/Sorting/SortHelper.cs

## Decomposition

### BaseComparer.cs (Base Comparer)

#### Imports
```csharp
using MediaBrowser.Controller.Entities;
using MediaBrowser.Model.Entities;
using System;
using System.Collections.Generic;
```

#### Classes
`BaseComparer` (public abstract class : IComparer<BaseItem>)

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `Compare(BaseItem, BaseItem)` | `int` | Compare two items |
| `GetValue(BaseItem)` | `IComparable` | Get sort value |

### NameComparer.cs (Name Comparer)

#### Classes
`NameComparer` (public class : BaseComparer)

Sorts items alphabetically by name.

### PremiereDateComparer.cs (Premiere Date Comparer)

#### Classes
`PremiereDateComparer` (public class : BaseComparer)

Sorts items by premiere/release date.

### RuntimeComparer.cs (Runtime Comparer)

#### Classes
`RuntimeComparer` (public class : BaseComparer)

Sorts items by runtime duration.

## Data Flow

```mermaid
graph LR
    A[Items List] --> B[SortHelper]
    B --> C[Comparers]
    C --> D[Sorted List]
```

## Dependencies

- `MediaBrowser.Controller.Entities` — Base item types
- `System.Collections.Generic` — Comparer interfaces

## Statistics

| Metric | Value |
|--------|-------|
| Files | 26 |
| Classes | 26 |
| LOC | ~500 |
