# Component: TSInterleavedFile.cs

**Path:** `BDInfo/TSInterleavedFile.cs`
**Type:** File | Class
**Language:** C#
**Maps to:** `.discovery/100-08-tsinterleavedfile.md`
**Parent:** `.discovery/100-bdinfo.md`
**Dependencies:**
- `MediaBrowser.Model.IO` (external) — FileSystemMetadata

## Structure

```
TSInterleavedFile.cs
├── [namespace] BDInfo
│   └── [class] public class TSInterleavedFile
│       ├── [field] public string Name
│       ├── [field] public FileSystemMetadata FileInfo
│       └── [constructor] TSInterleavedFile(FileSystemMetadata fileInfo)
│           └── Stores file metadata reference
```

## Description

`TSInterleavedFile` is a thin wrapper around `FileSystemMetadata` for 3D Blu-ray SSIF (Stereoscopic Interleaved File) entries. These files contain interleaved left/right eye frames for stereoscopic playback.
