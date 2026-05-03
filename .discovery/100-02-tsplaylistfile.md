# Component: TSPlaylistFile.cs

**Path:** `BDInfo/TSPlaylistFile.cs`
**Type:** File | Class
**Language:** C#
**Maps to:** `.discovery/100-02-tsplaylistfile.md`
**Parent:** `.discovery/100-bdinfo.md`
**Dependencies:**
- `MediaBrowser.Model.IO` (external) — File system abstractions
- `MediaBrowser.Model.Text` (external) — Text encoding
- `BDROM` → `.discovery/100-01-bdrom.md` — Parent BDROM instance
- `TSStream` → `.discovery/100-06-tsstream.md` — Stream base class
- `TSStreamClip` → `.discovery/100-07-tsstreamclip.md` — Stream clip
- `TSVideoStream`, `TSAudioStream`, `TSTextStream`, `TSGraphicsStream` — Stream subclasses

## Decomposition

#### Imports
```csharp
#define DEBUG
using System;
using System.Collections.Generic;
using System.IO;
using System.Text;
using MediaBrowser.Model.IO;
using MediaBrowser.Model.Text;
```

#### Namespace
```csharp
namespace BDInfo
```

#### Classes
`TSPlaylistFile` (public class)

#### Fields
```csharp
private readonly IFileSystem _fileSystem
private readonly ITextEncoding _textEncoding
private FileSystemMetadata FileInfo
public string FileType
public bool IsInitialized
public string Name
public BDROM BDROM
public bool HasHiddenTracks
public bool HasLoops
public bool IsCustom
public List<double> Chapters
public Dictionary<ushort, TSStream> Streams
public Dictionary<ushort, TSStream> PlaylistStreams
public List<TSStreamClip> StreamClips
public List<Dictionary<ushort, TSStream>> AngleStreams
public List<Dictionary<double, TSStreamClip>> AngleClips
public int AngleCount
public List<TSStream> SortedStreams
public List<TSVideoStream> VideoStreams
public List<TSAudioStream> AudioStreams
public List<TSTextStream> TextStreams
public List<TSGraphicsStream> GraphicsStreams
```

#### Properties
```csharp
public ulong InterleavedFileSize
public ulong FileSize
public double TotalLength
public double TotalAngleLength
public ulong TotalSize
public ulong TotalAngleSize
public ulong TotalBitRate
public ulong TotalAngleBitRate
public bool IsValid
```

#### Constructors
```csharp
TSPlaylistFile(BDROM bdrom, FileSystemMetadata fileInfo, IFileSystem fileSystem, ITextEncoding textEncoding)
TSPlaylistFile(BDROM bdrom, string name, List<TSStreamClip> clips, IFileSystem fileSystem, ITextEncoding textEncoding)
```

#### Methods
```csharp
public void Scan(Dictionary<string, TSStreamFile> streamFiles, Dictionary<string, TSStreamClipFile> streamClipFiles)
public void Initialize()
protected TSStream CreatePlaylistStream(byte[] data, ref int pos)
private void LoadStreamClips()
public void ClearBitrates()
```

## Structure

```
TSPlaylistFile.cs
├── [using] System, System.Collections.Generic, System.IO, System.Text
├── [using] MediaBrowser.Model.IO
│   └── IFileSystem, FileSystemMetadata
├── [using] MediaBrowser.Model.Text
│   └── ITextEncoding
├── [namespace] BDInfo
│   └── [class] public class TSPlaylistFile
│       ├── [field] private readonly IFileSystem _fileSystem
│       ├── [field] private readonly ITextEncoding _textEncoding
│       ├── [field] private FileSystemMetadata FileInfo
│       ├── [field] public string FileType
│       ├── [field] public bool IsInitialized
│       ├── [field] public string Name
│       ├── [field] public BDROM BDROM
│       ├── [field] public bool HasHiddenTracks
│       ├── [field] public bool HasLoops
│       ├── [field] public bool IsCustom
│       ├── [field] public List<double> Chapters
│       ├── [field] public Dictionary<ushort, TSStream> Streams
│       ├── [field] public Dictionary<ushort, TSStream> PlaylistStreams
│       ├── [field] public List<TSStreamClip> StreamClips
│       ├── [field] public List<Dictionary<ushort, TSStream>> AngleStreams
│       ├── [field] public List<Dictionary<double, TSStreamClip>> AngleClips
│       ├── [field] public int AngleCount
│       ├── [field] public List<TSStream> SortedStreams
│       ├── [field] public List<TSVideoStream> VideoStreams
│       ├── [field] public List<TSAudioStream> AudioStreams
│       ├── [field] public List<TSTextStream> TextStreams
│       ├── [field] public List<TSGraphicsStream> GraphicsStreams
│       ├── [constructor] TSPlaylistFile(BDROM bdrom, FileSystemMetadata fileInfo, IFileSystem fileSystem, ITextEncoding textEncoding)
│       │   └── Standard constructor from file system
│       ├── [constructor] TSPlaylistFile(BDROM bdrom, string name, List<TSStreamClip> clips, IFileSystem fileSystem, ITextEncoding textEncoding)
│       │   └── Custom playlist constructor (IsCustom = true)
│       │   └── Clones stream clips with relative timing
│       ├── [property] public override string ToString()
│       │   └── Returns Name
│       ├── [property] public ulong InterleavedFileSize
│       │   └── Sums interleaved file sizes across clips
│       ├── [property] public ulong FileSize
│       │   └── Sums stream file sizes across clips
│       ├── [property] public double TotalLength
│       │   └── Sums clip lengths (TimeOut - TimeIn)
│       ├── [property] public double TotalAngleLength
│       │   └── Sums angle clip lengths
│       ├── [property] public ulong TotalSize
│       │   └── Returns FileSize + InterleavedFileSize
│       ├── [property] public ulong TotalAngleSize
│       │   └── Sums angle stream sizes
│       ├── [property] public ulong TotalBitRate
│       │   └── Calculates bitrate from TotalSize / TotalLength
│       ├── [property] public ulong TotalAngleBitRate
│       │   └── Calculates angle bitrate
│       ├── [method] public void Scan(Dictionary<string, TSStreamFile> streamFiles, Dictionary<string, TSStreamClipFile> streamClipFiles)
│       │   ├── Reads MPLS file bytes via _fileSystem
│       │   ├── Parses MPLS header (type indicator, playlist count)
│       │   ├── Parses playlist items (stream clips with IN/OUT timestamps)
│       │   ├── Parses sub-path entries (PiP, secondary video)
│       │   ├── Parses chapter markers
│       │   ├── Calls LoadStreamClips() to resolve stream metadata
│       │   └── Sets IsInitialized = true
│       ├── [method] public void Initialize()
│       │   ├── Populates VideoStreams, AudioStreams, TextStreams, GraphicsStreams
│       │   ├── Populates SortedStreams (ordered by stream type/pid)
│       │   ├── Detects hidden tracks (streams in clips but not playlist)
│       │   └── Calculates bitrates per stream
│       ├── [method] protected TSStream CreatePlaylistStream(byte[] data, ref int pos)
│       │   ├── Reads stream type, PID, attributes from MPLS byte array
│       │   ├── Creates appropriate TSStream subclass (video/audio/text/graphics)
│       │   └── Returns populated TSStream instance
│       ├── [method] private void LoadStreamClips()
│       │   ├── Matches clip names to TSStreamClipFile entries
│       │   ├── Creates TSStreamClip instances with timing
│       │   ├── Resolves stream references from clip files
│       │   └── Handles multi-angle streams
│       ├── [method] public void ClearBitrates()
│       │   └── Resets bitrate calculations on all streams
│       └── [method] public bool IsValid
│           └── Returns true if file type is "MPLS" and has stream clips
```

## Description

`TSPlaylistFile` parses Blu-ray MPLS (Movie PlayLiSt) files, which define the playback order of transport stream clips. It reads the binary MPLS format, extracts chapter markers, stream clip sequences, and multi-angle information. It also calculates aggregate properties like total duration, file size, and bitrate across all clips in the playlist.

## Data Flow

```
MPLS file bytes ──► Scan()
    ├── Parse header ──► FileType, playlist count
    ├── Parse playlist items ──► StreamClips (with IN/OUT timestamps)
    ├── Parse sub-paths ──► Secondary streams
    ├── Parse chapters ──► Chapters list
    └── LoadStreamClips() ──► Resolves TSStream metadata

Initialize()
    ├── Streams.Values ──► VideoStreams, AudioStreams, TextStreams, GraphicsStreams
    ├── SortedStreams ──► Ordered by type/PID
    └── Bitrate calculations ──► Per-stream and total bitrates
```

## Side Effects

- Reads MPLS files via IFileSystem
- Creates TSStream, TSStreamClip, and subclass instances
- Modifies stream bitrate properties during Initialize()
