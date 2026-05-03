# Component: MediaBrowser.Server.Mono

**Path:** \`MediaBrowser.Server.Mono/\`
**Type:** Module
**Maps to:** \`.discovery/210-mediabrowser-server-mono.md\`

## Description

Mono runtime support for Emby Server. Provides platform-specific implementations for Linux/macOS.

## Structure

```
MediaBrowser.Server.Mono/
├── MediaBrowser.Server.Mono.csproj
├── MonoApp.cs                # Mono application host
├── MonoServerApplicationPaths.cs # Mono-specific paths
├── Native/                     # Native interop
│   ├── PosixFileSystem.cs
│   └── Syscall.cs
├── Security/                   # Mono security
│   └── MonoPrincipal.cs
└── packages.config
```

## Dependencies

- Emby.Server.Implementations
- MediaBrowser.Model
