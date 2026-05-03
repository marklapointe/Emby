# Component: DvdLib

**Path:** `DvdLib/`
**Type:** Directory | Library
**Language:** C#
**Maps to:** `.discovery/110-dvdlib.md`

## Description

DvdLib is a DVD IFO (Information File) parsing library. It reads DVD structure files to extract metadata about titles, chapters, audio tracks, subtitles, and navigation commands. Used by Emby for DVD media analysis and metadata extraction.

## Structure

```
DvdLib/
├── DvdLib.csproj                # Project file
├── BigEndianBinaryReader.cs     # Big-endian binary reader for DVD structures
├── Ifo/                         # IFO parsing components
│   ├── Dvd.cs                   # Main DVD parser → [class] Dvd
│   ├── Title.cs                 # DVD title
│   ├── Chapter.cs               # Chapter within a title
│   ├── Program.cs               # Program (cell group)
│   ├── ProgramChain.cs          # Program chain (PGC)
│   ├── Cell.cs                  # Individual cell
│   ├── CellPlaybackInfo.cs      # Cell playback information
│   ├── CellPositionInfo.cs      # Cell position information
│   ├── VideoAttributes.cs       # Video stream attributes
│   ├── AudioAttributes.cs       # Audio stream attributes
│   ├── UserOperation.cs         # User operation permissions
│   ├── DvdTime.cs               # DVD time representation
│   └── PgcCommandTable.cs       # Program chain command table
└── Properties/                  # Assembly info
```

## Key Classes

| Class | File | Purpose |
|-------|------|---------|
| `Dvd` | `Ifo/Dvd.cs` | Main entry point for DVD parsing |
| `Title` | `Ifo/Title.cs` | Represents a DVD title |
| `Chapter` | `Ifo/Chapter.cs` | Represents a chapter |
| `ProgramChain` | `Ifo/ProgramChain.cs` | Navigation program chain |
| `VideoAttributes` | `Ifo/VideoAttributes.cs` | Video codec/resolution info |
| `AudioAttributes` | `Ifo/AudioAttributes.cs` | Audio codec/language info |

## Data Flow

```mermaid
graph LR
    A[Dvd constructor] --> B[Read VIDEO_TS.IFO]
    B --> C[Parse VMGI]
    C --> D[Extract title count]
    D --> E[For each title:]
    E --> F[Read VTS_XX_0.IFO]
    F --> G[Parse VTSI]
    G --> H[Extract chapters, programs, cells]
    H --> I[Extract video/audio attributes]
```

## Dependencies

- `BigEndianBinaryReader` — Custom binary reader for DVD big-endian data

## Side Effects

- Reads DVD IFO files from filesystem
- Parses binary DVD structures

## Reference

- DVD-Video specification (ISO/IEC 13818-1)
