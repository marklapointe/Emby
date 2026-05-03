# Component: MediaBrowser.XbmcMetadata

**Path:** `MediaBrowser.XbmcMetadata/`
**Type:** Directory | Module
**Language:** C#
**Maps to:** `.discovery/270-mediabrowser-xbmcmetadata.md`

## Description

MediaBrowser.XbmcMetadata provides compatibility with XBMC/Kodi metadata formats. It reads and writes NFO files, which are XML-based metadata files used by Kodi to store media information alongside media files.

## Structure

```
MediaBrowser.XbmcMetadata/
├── MediaBrowser.XbmcMetadata.csproj
├── Nfo/                         # NFO file parsers
│   ├── BaseNfoParser.cs         # Base NFO parser
│   ├── MovieNfoParser.cs        # Movie NFO parser
│   ├── EpisodeNfoParser.cs      # TV episode NFO parser
│   └── ...                      
└── Properties/                  # Assembly info
```

## Key Classes

| Class | File | Purpose |
|-------|------|---------|
| `BaseNfoParser` | `Nfo/` | Base XML parsing logic |
| `MovieNfoParser` | `Nfo/` | Parses movie.nfo files |
| `EpisodeNfoParser` | `Nfo/` | Parses episode NFOs |

## Dependencies

- `MediaBrowser.Controller` — Metadata interfaces
- `MediaBrowser.Model` — Media types

## Side Effects

- Reads NFO XML files from media directories
- Writes NFO files when saving metadata locally

## Reference

- Kodi NFO format: `https://kodi.wiki/view/NFO_files`
