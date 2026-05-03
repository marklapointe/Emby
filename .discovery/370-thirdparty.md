# Component: ThirdParty

**Path:** \`ThirdParty/\`
**Type:** Directory | Third-Party Libraries
**Language:** C++, C#
**Maps to:** \`.discovery/370-thirdparty.md\`

## Description

ThirdParty contains native libraries and third-party dependencies bundled with Emby Server.

## Decomposition

### 7zip (Archive Extraction)

#### File
`7zip/7za.exe` — 7-Zip command-line archive extraction tool. Used for extracting compressed media containers and archives.

### Emby Server Plugins (DLL)

| File | Purpose |
|------|---------|
| `emby/Emby.Server.Connect.dll` | Emby Connect sync service |
| `emby/Emby.Server.MediaEncoding.dll` | Media encoding service |
| `emby/Emby.Server.Sync.dll` | Media synchronization |

### TagLib.Portable (Audio Tagging)

`taglib/TagLib.Portable.dll` — Cross-platform audio metadata (ID3, Vorbis Comments, etc.) reading and writing library.

## Contents

- `7zip/` — 1 files
  - `7za.exe`
- `emby/` — 3 files
  - `Emby.Server.Connect.dll`
  - `Emby.Server.MediaEncoding.dll`
  - `Emby.Server.Sync.dll`
- `taglib/` — 1 files
  - `TagLib.Portable.dll`
