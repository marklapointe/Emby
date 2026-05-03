# Component: DvdLib

**Path:** `DvdLib/`
**Type:** Directory | Library
**Language:** C#
**Maps to:** `.discovery/110-dvdlib.md`

## Description

DvdLib is a DVD Video IFO (InFOrmation) file parsing library. It reads DVD-Video disc structures (VIDEO_TS.IFO, VTS_*.IFO) to extract metadata about titles, chapters, audio/video attributes, program chains, and cell playback information. Used by Emby for DVD media analysis.

## Structure

```
DvdLib/
├── DvdLib.csproj              # Project file
├── packages.config            # NuGet packages
├── BigEndianBinaryReader.cs   # Big-endian binary reader
│   └── [class] BigEndianBinaryReader : BinaryReader
│       └── Reads multi-byte values in big-endian order (required for DVD IFO format)
├── Ifo/
│   ├── Dvd.cs                 # Main DVD parser entry point
│   │   └── [class] Dvd
│   │       ├── Constructor: Dvd(string path, IFileSystem fileSystem)
│   │       ├── Reads VIDEO_TS.IFO (VMG - Video Manager)
│   │       ├── Reads VTS_*.IFO (VTS - Video Title Set)
│   │       ├── Parses TT_SRPT (Title Search Pointer Table)
│   │       └── Populates Titles list
│   ├── Title.cs               # DVD title
│   │   └── [class] Title
│   │       ├── ParseTT_SRPT() — parses title search pointer
│   │       ├── ParseVTS() — parses video title set
│   │       └── Contains ProgramChains, Chapters
│   ├── ProgramChain.cs        # Program Chain (PGC)
│   │   └── [class] ProgramChain
│   │       ├── Parse() — reads PGC from IFO
│   │       ├── Contains Cells, Programs
│   │       └── Playback mode flags
│   ├── Cell.cs                # Cell (playback unit)
│   │   └── [class] Cell
│   │       ├── CellPlaybackInfo
│   │       └── CellPositionInfo
│   ├── CellPlaybackInfo.cs    # Cell playback metadata
│   │   └── [class] CellPlaybackInfo
│   │       ├── BlockMode, BlockType enums
│   │       ├── PlaybackMode enum
│   │       └── Start/end times, interleaving flags
│   ├── CellPositionInfo.cs    # Cell position metadata
│   │   └── [class] CellPositionInfo
│   │       └── VOB ID, cell ID, start/end sectors
│   ├── Chapter.cs             # Chapter
│   │   └── [class] Chapter
│   │       └── Start/end cell indices
│   ├── Program.cs             # Program (within PGC)
│   │   └── [class] Program
│   │       └── Cell indices, playback time
│   ├── AudioAttributes.cs     # Audio stream attributes
│   │   └── [class] AudioAttributes
│   │       ├── AudioCodec enum (AC3, MPEG1, MPEG2, LPCM, DTS, SDDS)
│   │       ├── ApplicationMode enum
│   │       └── Language, channels, sample rate, bit rate
│   ├── VideoAttributes.cs     # Video stream attributes
│   │   └── [class] VideoAttributes
│   │       ├── VideoCodec enum (MPEG1, MPEG2)
│   │       ├── VideoFormat enum (NTSC, PAL)
│   │       ├── AspectRatio enum (4:3, 16:9)
│   │       ├── FilmMode enum
│   │       └── Resolution, frame rate, bit rate
│   ├── DvdTime.cs             # DVD time representation
│   │   └── [class] DvdTime
│   │       ├── Parses BCD-encoded DVD time format
│   │       └── Converts to TimeSpan
│   ├── PgcCommandTable.cs     # Program Chain Command Table
│   │   └── [class] ProgramChainCommandTable
│   │       └── [class] VirtualMachineCommand
│   │           └── DVD virtual machine bytecode commands
│   └── UserOperation.cs       # User operation flags
│       └── [enum] UserOperation
│           └── Playback control flags ( prohibited operations)
└── Properties/
    └── AssemblyInfo.cs
```

## Data Flow

```mermaid
graph TD
    A[Dvd constructor] --&gt; B[Locate VIDEO_TS.IFO]
    B --&gt; C[Parse VMG header]
    C --&gt; D[Read TT_SRPT]
    D --&gt; E[Create Title instances]
    E --&gt; F[For each VTS]
    F --&gt; G[Parse VTS_*.IFO]
    G --&gt; H[ProgramChain.Parse]
    H --&gt; I[Cell, Chapter, Program]
    I --&gt; J[AudioAttributes, VideoAttributes]
```

## Key Classes

| Class | File | Purpose |
|-------|------|---------|
| `Dvd` | `Ifo/Dvd.cs` | Main entry point for DVD parsing |
| `Title` | `Ifo/Title.cs` | DVD title with program chains |
| `ProgramChain` | `Ifo/ProgramChain.cs` | Playback sequence (PGC) |
| `Cell` | `Ifo/Cell.cs` | Smallest playback unit |
| `Chapter` | `Ifo/Chapter.cs` | Logical chapter markers |
| `AudioAttributes` | `Ifo/AudioAttributes.cs` | Audio stream metadata |
| `VideoAttributes` | `Ifo/VideoAttributes.cs` | Video stream metadata |

## Decomposition

### BigEndianBinaryReader.cs (Binary Reader)

#### Imports
```csharp
using System;
using System.IO;
using System.Text;
```

#### Classes
`BigEndianBinaryReader` (public class : BinaryReader)

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `ReadInt16()` | `Int16` | Read big-endian Int16 |
| `ReadInt32()` | `Int32` | Read big-endian Int32 |
| `ReadUInt16()` | `UInt16` | Read big-endian UInt16 |
| `ReadUInt32()` | `UInt32` | Read big-endian UInt32 |
| `ReadSingle()` | `Single` | Read big-endian float |
| `ReadString(int)` | `string` | Read fixed-length string |

### Ifo/Dvd.cs (Main Parser)

#### Imports
```csharp
using MediaBrowser.Model.IO;
using System;
using System.Collections.Generic;
using System.IO;
```

#### Classes
`Dvd` (public class)

#### Key Properties
| Property | Type | Description |
|----------|------|-------------|
| `Titles` | `IEnumerable<DvdTitle>` | All parsed titles |

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `Open(string, IFileSystem)` | `DvdInfo` | Parse DVD directory |

### Ifo/Title.cs (Title Parser)

#### Classes
`DvdTitle` (public class)

#### Key Properties
| Property | Type | Description |
|----------|------|-------------|
| `TitleNumber` | `int` | Title index |
| `Duration` | `TimeSpan` | Total playback time |
| `Chapters` | `List<DvdChapter>` | Chapter list |
| `VideoStream` | `DvdVideoStream` | Video track info |
| `AudioStreams` | `List<DvdAudioStream>` | Audio tracks |

### Ifo/ProgramChain.cs (Playback Sequence)

#### Classes
`DvdProgramChain` (public class)

#### Key Properties
| Property | Type | Description |
|----------|------|-------------|
| `Cells` | `List<DvdCell>` | Playback cells |
| `Programs` | `List<DvdProgram>` | Program entries |
| `IsStill` | `bool` | Still frame flag |
| `StillTime` | `int` | Still duration |

## Side Effects

- Reads IFO/BUP files via IFileSystem
- Uses BigEndianBinaryReader for DVD binary format
- No write operations (read-only parsing)
