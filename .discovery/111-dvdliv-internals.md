# Component: DvdLib — Expanded

**Path:** `DvdLib/`
**Type:** Directory | Library
**Language:** C#
**Maps to:** `.discovery/111-dvdliv-internals.md`

## Decomposition

### BigEndianBinaryReader.cs (Binary Reader)

#### Imports
```csharp
using System;
using System.IO;
using System.Text;
```

#### Classes
\`BigEndianBinaryReader\` (public class : BinaryReader)

#### Key Methods
```csharp
public override Int16 ReadInt16()
public override Int32 ReadInt32()
public override UInt16 ReadUInt16()
public override UInt32 ReadUInt32()
public override Single ReadSingle()
public string ReadString(int length)
```

### Dvd.cs (Main Entry Point)

#### Imports
```csharp
using MediaBrowser.Model.IO;
using System;
using System.Collections.Generic;
using System.IO;
```

#### Classes
\`Dvd\` (public class)

#### Key Methods
```csharp
public static DvdInfo Open(string path, IFileSystem fileSystem)
public IEnumerable<DvdTitle> Titles { get; }
```

### Title.cs (Title Definition)

#### Classes
\`DvdTitle\` (public class)

#### Key Properties
```csharp
public int TitleNumber
public TimeSpan Duration
public List<DvdChapter> Chapters
public DvdVideoStream VideoStream
public List<DvdAudioStream> AudioStreams
```

### ProgramChain.cs (Program Chain)

#### Classes
\`DvdProgramChain\` (public class)

#### Key Properties
```csharp
public List<DvdCell> Cells
public List<DvdProgram> Programs
public bool IsStill
public int StillTime
```

## Description

DVD structure analysis library. Parses DVD IFO files to extract title, chapter, and stream information. Uses big-endian binary reading for DVD format compatibility.

## Files

### Root Files (2 files)

- `BigEndianBinaryReader.cs` — DvdLib/BigEndianBinaryReader.cs
- `DvdLib.csproj` — DvdLib/DvdLib.csproj

### Ifo/ (12 files)

- `AudioAttributes.cs` — DvdLib/Ifo/AudioAttributes.cs
- `Cell.cs` — DvdLib/Ifo/Cell.cs
- `CellPlaybackInfo.cs` — DvdLib/Ifo/CellPlaybackInfo.cs
- `CellPositionInfo.cs` — DvdLib/Ifo/CellPositionInfo.cs
- `Chapter.cs` — DvdLib/Ifo/Chapter.cs
- `Dvd.cs` — DvdLib/Ifo/Dvd.cs
- `DvdTime.cs` — DvdLib/Ifo/DvdTime.cs
- `PgcCommandTable.cs` — DvdLib/Ifo/PgcCommandTable.cs
- `Program.cs` — DvdLib/Ifo/Program.cs
- `ProgramChain.cs` — DvdLib/Ifo/ProgramChain.cs
- `Title.cs` — DvdLib/Ifo/Title.cs
- `UserOperation.cs` — DvdLib/Ifo/UserOperation.cs
- `VideoAttributes.cs` — DvdLib/Ifo/VideoAttributes.cs

### Properties/ (1 file)

- `Properties/AssemblyInfo.cs` — DvdLib/Properties/AssemblyInfo.cs

### Config Files (2 files)

- `packages.config` — DvdLib/packages.config

## DVD Structure

```mermaid
graph TD
    A[DVD Disc] --> B[IFO Files]
    B --> C[VideoTitleSet]
    C --> D[Title]
    D --> E[ProgramChain]
    E --> F[Cell]
    E --> G[Program]
    D --> H[Chapter]
    C --> I[AudioAttributes]
    C --> J[VideoAttributes]
```

## Key Classes

| Class | Responsibility |
|-------|----------------|
| `BigEndianBinaryReader` | Reads big-endian binary data |
| `Dvd` | Main DVD parsing entry point |
| `Title` | Represents a DVD title |
| `ProgramChain` | Program sequence information |
| `Cell` | Playback cell data |
| `Chapter` | Chapter navigation |

## Dependencies

- Standard .NET libraries
