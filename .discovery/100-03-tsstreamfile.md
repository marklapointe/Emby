# Component: TSStreamFile.cs

**Path:** `BDInfo/TSStreamFile.cs`
**Type:** File | Multiple Classes
**Language:** C#
**Maps to:** `.discovery/100-03-tsstreamfile.md`
**Parent:** `.discovery/100-bdinfo.md`
**Dependencies:**
- `MediaBrowser.Model.IO` (external) — IFileSystem, FileSystemMetadata
- `TSStream` → `.discovery/100-06-tsstream.md` — Stream base class
- `TSStreamBuffer` → `.discovery/100-10-tsstreambuffer.md` — Byte buffer
- `TSCodec*` → `.discovery/100-11-codecs.md` — Codec scanners
- `TSPlaylistFile` → `.discovery/100-02-tsplaylistfile.md` — Playlist references

## Decomposition

#### Imports
```csharp
#undef DEBUG
using System;
using System.Collections.Generic;
using System.IO;
using MediaBrowser.Model.IO;
```

#### Namespace
```csharp
namespace BDInfo
```

#### Classes
- `TSStreamState` — Tracks per-stream parsing state during transport stream demux
- `TSPacketParser` — Tracks MPEG-TS packet-level parsing state
- `TSStreamDiagnostics` — Holds diagnostic metrics per stream
- `TSStreamFile` — Main transport stream scanner

#### TSStreamState Fields
```csharp
public ulong TransferCount
public string StreamTag
public ulong TotalPackets, WindowPackets
public ulong TotalBytes, WindowBytes
public long PeakTransferLength, PeakTransferRate
public double TransferMarker, TransferInterval
public TSStreamBuffer StreamBuffer
public uint Parse
public bool TransferState
public int TransferLength, PacketLength
public byte PacketLengthParse, PacketParse
public byte PTSParse
public ulong PTS, PTSTemp, PTSLast, PTSPrev, PTSDiff, PTSCount, PTSTransfer
public byte DTSParse
public ulong DTSTemp, DTSPrev
public byte PESHeaderLength, PESHeaderFlags, PESHeaderIndex
public byte[] PESHeader
```

#### TSPacketParser Fields
```csharp
public bool SyncState
public byte TimeCodeParse, PacketLength, HeaderParse
public uint TimeCode
public byte TransportErrorIndicator, PayloadUnitStartIndicator
public byte TransportPriority, TransportScramblingControl
public ushort PID
public byte AdaptionFieldControl, AdaptionFieldParse, AdaptionFieldLength
public bool AdaptionFieldState
public ushort PCRPID
public byte PCRParse
public ulong PCR, PCRTemp, PCRDiff, PCRCount, PCRTransfer
public ulong PTSFirst, PTSLast, PTSDiff
public byte[] PAT
public bool PATSectionStart, PATTransferState
```

#### TSStreamDiagnostics Fields
```csharp
public ulong Bytes
public ulong Packets
public double Seconds
public double BitRate
public double TransferRate
```

#### TSStreamFile Fields
```csharp
public FileSystemMetadata FileInfo
public string Name
public TSInterleavedFile InterleavedFile
public Dictionary<ushort, TSStream> Streams
public bool IsInitialized
public bool Is3D
private readonly IFileSystem _fileSystem
```

#### TSStreamFile Methods
```csharp
public TSStreamFile(FileSystemMetadata fileInfo, IFileSystem fileSystem)
public void Scan(List<TSPlaylistFile> playlists, bool isFullScan)
private void UpdateStreamBitrates(List<TSPlaylistFile> playlists, bool isFullScan)
private void UpdateStreamBitrate(TSStream stream, TSStreamDiagnostics diag)
```

## Structure

```
TSStreamFile.cs
├── [namespace] BDInfo
│   ├── [class] public class TSStreamState
│   │   ├── [field] public ulong TransferCount
│   │   ├── [field] public string StreamTag
│   │   ├── [field] public ulong TotalPackets, WindowPackets
│   │   ├── [field] public ulong TotalBytes, WindowBytes
│   │   ├── [field] public long PeakTransferLength, PeakTransferRate
│   │   ├── [field] public double TransferMarker, TransferInterval
│   │   ├── [field] public TSStreamBuffer StreamBuffer
│   │   ├── [field] public uint Parse
│   │   ├── [field] public bool TransferState
│   │   ├── [field] public int TransferLength, PacketLength
│   │   ├── [field] public byte PacketLengthParse, PacketParse
│   │   ├── [field] public byte PTSParse
│   │   ├── [field] public ulong PTS, PTSTemp, PTSLast, PTSPrev, PTSDiff, PTSCount, PTSTransfer
│   │   ├── [field] public byte DTSParse
│   │   ├── [field] public ulong DTSTemp, DTSPrev
│   │   ├── [field] public byte PESHeaderLength, PESHeaderFlags, PESHeaderIndex
│   │   └── [field] public byte[] PESHeader
│   │   └── Tracks per-stream parsing state during transport stream demux
│   ├── [class] public class TSPacketParser
│   │   ├── [field] public bool SyncState
│   │   ├── [field] public byte TimeCodeParse, PacketLength, HeaderParse
│   │   ├── [field] public uint TimeCode
│   │   ├── [field] public byte TransportErrorIndicator, PayloadUnitStartIndicator
│   │   ├── [field] public byte TransportPriority, TransportScramblingControl
│   │   ├── [field] public ushort PID
│   │   ├── [field] public byte AdaptionFieldControl, AdaptionFieldParse, AdaptionFieldLength
│   │   ├── [field] public bool AdaptionFieldState
│   │   ├── [field] public ushort PCRPID
│   │   ├── [field] public byte PCRParse
│   │   └── [field] public ulong PCR, PCRTemp, PCRDiff, PCRCount, PCRTransfer
│   │   └── Tracks MPEG-TS packet-level parsing state
│   ├── [class] public class TSStreamDiagnostics
│   │   ├── [field] public ulong Bytes
│   │   ├── [field] public ulong Packets
│   │   ├── [field] public double Seconds
│   │   ├── [field] public double BitRate
│   │   └── [field] public double TransferRate
│   │   └── Holds diagnostic metrics per stream
│   └── [class] public class TSStreamFile
│       ├── [field] public FileSystemMetadata FileInfo
│       ├── [field] public string Name
│       ├── [field] public TSInterleavedFile InterleavedFile
│       ├── [field] public Dictionary<ushort, TSStream> Streams
│       ├── [field] public bool IsInitialized
│       ├── [field] public bool Is3D
│       ├── [field] private readonly IFileSystem _fileSystem
│       ├── [constructor] TSStreamFile(FileSystemMetadata fileInfo, IFileSystem fileSystem)
│       │   └── Stores file metadata and file system reference
│       ├── [method] private void UpdateStreamBitrates(List<TSPlaylistFile> playlists, bool isFullScan)
│       │   ├── Calculates per-stream bitrates from packet counts
│       │   ├── Handles VBR (Variable Bit Rate) streams
│       │   └── Updates TSStream.BitRate and ActiveBitRate
│       ├── [method] private void UpdateStreamBitrate(TSStream stream, TSStreamDiagnostics diag)
│       │   └── Updates bitrate for a single stream from diagnostics
│       └── [method] public void Scan(List<TSPlaylistFile> playlists, bool isFullScan)
│           ├── Opens .m2ts file via IFileSystem
│           ├── Initializes TSStreamState per PID
│           ├── Reads file in chunks (transport stream packets)
│           ├── Parses MPEG-TS packet headers (TSPacketParser)
│           ├── Demuxes PES packets per PID
│           ├── Accumulates stream data in TSStreamBuffer
│           ├── Routes to codec scanners based on stream type
│           │   ├── TSCodecAC3.Scan() for AC-3/E-AC-3
│           │   ├── TSCodecAVC.Scan() for H.264/AVC
│           │   ├── TSCodecDTS.Scan() for DTS
│           │   ├── TSCodecDTSHD.Scan() for DTS-HD MA
│           │   ├── TSCodecLPCM.Scan() for LPCM
│           │   ├── TSCodecMPEG2.Scan() for MPEG-2
│           │   ├── TSCodecMVC.Scan() for MVC
│           │   ├── TSCodecTrueHD.Scan() for Dolby TrueHD
│           │   └── TSCodecVC1.Scan() for VC-1
│           ├── Calculates bitrates via UpdateStreamBitrates()
│           ├── Populates TSStreamDiagnostics
│           └── Sets IsInitialized = true
```

## Description

`TSStreamFile` is the core transport stream scanner for BDInfo. It reads Blu-ray `.m2ts` (MPEG-2 Transport Stream) files, demuxes packets by PID, routes them to codec-specific scanners, and calculates bitrates. It uses `TSStreamState` to track per-stream parsing state, `TSPacketParser` for MPEG-TS packet header parsing, and `TSStreamDiagnostics` for metrics collection.

## Data Flow

```
.m2ts file ──► Scan()
    ├── Open file ──► Read transport stream packets (188/192 bytes)
    ├── Parse TS header ──► TSPacketParser
    │   ├── PID extraction
    │   ├── Adaptation field handling
    │   └── PCR timing extraction
    ├── Demux PES packets ──► TSStreamState per PID
    │   ├── PTS/DTS extraction
    │   └── Payload accumulation ──► TSStreamBuffer
    ├── Route to codec scanner ──► Based on stream type from CLPI
    │   └── Populates TSStream properties
    └── Calculate bitrates ──► UpdateStreamBitrates()
        └── TSStreamDiagnostics
```

## Side Effects

- Reads .m2ts files via IFileSystem (potentially large files)
- Creates TSStreamState, TSPacketParser, TSStreamDiagnostics instances
- Modifies TSStream properties via codec scanners
- Disposes file streams in finally block
