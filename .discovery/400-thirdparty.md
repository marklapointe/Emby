# Component: ThirdParty

**Path:** `ThirdParty/`
**Type:** Directory | External Libraries
**Language:** C++ / C#
**Maps to:** `.discovery/400-thirdparty.md`

## Description

ThirdParty contains external libraries integrated into Emby. These include 7zip for archive extraction, taglib for audio metadata reading, and the Emby core library. These are compiled or referenced as dependencies.

## Structure

```
ThirdParty/
├── 7zip/                        # 7-Zip compression library
│   └── ...                      # C++ sources for archive support
├── emby/                        # Emby core third-party components
│   └── ...                      
└── taglib/                      # TagLib audio metadata library
    └── ...                      # C++ sources for ID3, Vorbis, etc.
```

## Components

| Library | Path | Purpose |
|---------|------|---------|
| 7zip | `ThirdParty/7zip/` | Archive extraction (zip, rar, 7z) |
| taglib | `ThirdParty/taglib/` | Audio metadata (MP3, FLAC, OGG) |
| emby | `ThirdParty/emby/` | Emby-specific third-party code |

## Side Effects

- Linked at compile time
- Used by `MediaBrowser.Providers` and `Emby.Server.Implementations`

## Reference

- 7-Zip: `https://www.7-zip.org/`
- TagLib: `https://taglib.org/`
