# Component: Emby.Server.Implementations — I/O Utilities

**Path:** \`Emby.Server.Implementations/IO/\`
**Type:** Directory | Module Group
**Language:** C#
**Maps to:** \`.discovery/166-emby-server-impl-io.md\`

## Description

I/O utilities for file system operations, shortcuts, and managed file system wrappers.

## Files

### IO/

- `ExtendedFileSystemInfo.cs` — Emby.Server.Implementations/IO/ExtendedFileSystemInfo.cs
- `FileRefresher.cs` — Emby.Server.Implementations/IO/FileRefresher.cs
- `IsoManager.cs` — Emby.Server.Implementations/IO/IsoManager.cs
- `LibraryMonitor.cs` — Emby.Server.Implementations/IO/LibraryMonitor.cs
- `ManagedFileSystem.cs` — Emby.Server.Implementations/IO/ManagedFileSystem.cs
- `MbLinkShortcutHandler.cs` — Emby.Server.Implementations/IO/MbLinkShortcutHandler.cs
- `SharpCifsFileSystem.cs` — Emby.Server.Implementations/IO/SharpCifsFileSystem.cs
- `StreamHelper.cs` — Emby.Server.Implementations/IO/StreamHelper.cs
- `ThrottledStream.cs` — Emby.Server.Implementations/IO/ThrottledStream.cs

## SharpCifs Sub-Module

Embedded SMB/CIFS client library. See `.discovery/169-emby-server-impl-sharpcifs.md` for full file listing.

## Decomposition

### ManagedFileSystem.cs (File System Abstraction)

#### Imports
```csharp
using MediaBrowser.Model.IO;
using System;
using System.IO;
using System.Threading;
```

#### Classes
`ManagedFileSystem` (public class : IFileSystem)

#### Key Properties
| Property | Type | Description |
|----------|------|-------------|
| `ApplicationPaths` | `IApplicationPaths` | Path configuration |

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `GetFileInfo(string)` | `FileInfo` | Get file metadata |
| `GetDirectoryInfo(string)` | `DirectoryInfo` | Get directory metadata |
| `DeleteFile(string)` | `void` | Delete file |
| `DeleteDirectory(string, bool)` | `void` | Delete directory |
| `CopyFile(string, string, bool)` | `void` | Copy file |

### LibraryMonitor.cs (File System Watcher)

#### Classes
`LibraryMonitor` (public class)

#### Key Properties
| Property | Type | Description |
|----------|------|-------------|
| `EnableFolderWatching` | `bool` | Enable/disable watching |

#### Key Events
| Event | Description |
|-------|-------------|
| `FileSystemChanged` | File/directory changed |
| `FilesAdded` | New files detected |
| `FilesRemoved` | Files deleted |

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `StartWatching(string, string, bool, FileSystemEventHandler)` | `void` | Start watching |
| `StopWatching(string)` | `void` | Stop watching |

### IsoManager.cs (ISO Mount Management)

#### Classes
`IsoManager` (public class)

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `MountIso(string)` | `string` | Mount ISO file |
| `UnmountIso(string)` | `void` | Unmount ISO |
| `IsIso(string)` | `bool` | Check if ISO file |

### FileRefresher.cs (Directory Watcher)

#### Classes
`FileRefresher` (public class : IDisposable)

#### Key Properties
| Property | Type | Description |
|----------|------|-------------|
| `EnableRaisingEvents` | `bool` | Start/stop watching |
| `Path` | `string` | Watched path |

#### Key Events
| Event | Description |
|-------|-------------|
| `Created` | File created |
| `Changed` | File modified |
| `Deleted` | File deleted |
| `Renamed` | File renamed |

### ThrottledStream.cs (Rate-Limited Stream)

#### Classes
`ThrottledStream` (public class : Stream)

#### Key Properties
| Property | Type | Description |
|----------|------|-------------|
| `Throttle` | `long` | Bytes per second limit |
