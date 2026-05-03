# Component: TSStreamClipFile.cs

**Path:** `BDInfo/TSStreamClipFile.cs`
**Type:** File | Class
**Language:** C#
**Maps to:** `.discovery/100-04-tsstreamclipfile.md`
**Parent:** `.discovery/100-bdinfo.md`
**Dependencies:**
- `MediaBrowser.Model.IO` (external) — IFileSystem, FileSystemMetadata
- `MediaBrowser.Model.Text` (external) — ITextEncoding
- `TSStream` → `.discovery/100-06-tsstream.md` — Stream base class

## Structure

```
TSStreamClipFile.cs
├── [using] System, System.Collections.Generic, System.IO
├── [using] MediaBrowser.Model.IO
│   └── IFileSystem, FileSystemMetadata
├── [using] MediaBrowser.Model.Text
│   └── ITextEncoding
├── [namespace] BDInfo
│   └── [class] public class TSStreamClipFile
│       ├── [field] private readonly IFileSystem _fileSystem
│       ├── [field] private readonly ITextEncoding _textEncoding
│       ├── [field] public FileSystemMetadata FileInfo = null
│       ├── [field] public string FileType = null
│       ├── [field] public bool IsValid = false
│       ├── [field] public string Name = null
│       ├── [field] public Dictionary<ushort, TSStream> Streams
│       │   └── Maps PID → TSStream for all streams in this clip
│       ├── [constructor] TSStreamClipFile(FileSystemMetadata fileInfo, IFileSystem fileSystem, ITextEncoding textEncoding)
│       │   └── Stores file metadata and service references
│       └── [method] public void Scan()
│           ├── Opens CLPI file via _fileSystem
│           ├── Reads CLPI header (type indicator)
│           ├── Parses clip info section
│           ├── Parses stream info section (PID, stream type, attributes)
│           ├── Creates TSStream instances via CreatePlaylistStream pattern
│           ├── Populates Streams dictionary
│           └── Sets IsValid = true on success
```

## Description

`TSStreamClipFile` parses Blu-ray CLPI (Clip Information) files, which contain metadata about individual transport stream clips (.m2ts files). It reads the binary CLPI format, extracts stream information (PID, codec type, language, attributes), and populates a dictionary of `TSStream` instances keyed by PID.

## Data Flow

```
CLPI file bytes ──► Scan()
    ├── Parse header ──► FileType validation
    ├── Parse clip info ──► Timing, encoding info
    ├── Parse stream info ──► PID, stream type, attributes
    └── Create TSStream instances ──► Streams dictionary
```

## Side Effects

- Reads CLPI files via IFileSystem
- Creates TSStream and subclass instances
- Disposes file readers in finally block
