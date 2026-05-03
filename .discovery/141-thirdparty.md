# Component: ThirdParty

**Path:** `ThirdParty/`
**Type:** Directory | Binary Dependencies
**Maps to:** `.discovery/141-thirdparty.md`

## Description

ThirdParty contains external binary dependencies shipped with Emby Server. These are pre-compiled DLLs and executables that are not built from source in this repository but are required at runtime.

## Structure

```
ThirdParty/
├── 7zip/
│   └── 7za.exe                    # 7-Zip command-line archiver
│       └── [BINARY] Windows executable for archive extraction
│       └── Used by: Emby.Server.Implementations.Archiving
├── emby/
│   ├── Emby.Server.Connect.dll    # Emby Connect service
│   │   └── [BINARY] Emby cloud connectivity (remote access, user linking)
│   │   └── Used by: MediaBrowser.Api.ConnectService
│   ├── Emby.Server.MediaEncoding.dll # Media encoding extensions
│   │   └── [BINARY] Advanced media encoding features
│   │   └── Used by: Emby.Server.Implementations.MediaEncoder
│   └── Emby.Server.Sync.dll       # Server sync service
│       └── [BINARY] Content synchronization between servers
│       └── Used by: Emby.Server.Implementations.SyncManager
└── taglib/
    └── TagLib.Portable.dll        # Audio metadata library
        └── [BINARY] Cross-platform audio tag reading (ID3, Vorbis, APE)
        └── Used by: Emby.Server.Implementations.AudioMetadataProvider
```

## Dependencies

| File | Type | Purpose | Consumers |
|------|------|---------|-----------|
| 7za.exe | Executable | Archive extraction | Archiving service |
| Emby.Server.Connect.dll | .NET Assembly | Cloud connectivity | Connect API |
| Emby.Server.MediaEncoding.dll | .NET Assembly | Advanced encoding | Media encoder |
| Emby.Server.Sync.dll | .NET Assembly | Server sync | Sync manager |
| TagLib.Portable.dll | .NET Assembly | Audio metadata | Audio metadata provider |

## Side Effects

- Loaded at runtime via Assembly.Load or P/Invoke
- No source code in this repository
- Updated via Emby update mechanism
