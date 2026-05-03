# Component: Emby.Dlna — Server

**Path:** `Emby.Dlna/Server/`
**Type:** Directory | Sub-module
**Language:** C#
**Maps to:** `.discovery/332-emby-dlna-server.md`
**Parent:** `.discovery/330-emby-dlna.md`

## Description

DLNA server implementation. Handles UPnP/DLNA protocol communication,
content directory browsing, and media streaming.

## Structure

```
Server/
├── DlnaServer.cs                 # [class] DlnaServer
│   ├── Main DLNA server
│   ├── Handles UPnP announcements
│   └── Content directory service
├── ContentDirectory.cs           # [class] ContentDirectory
│   └── Browsable content tree
├── ConnectionManager.cs          # [class] ConnectionManager
│   └── Connection management
└── *Handler.cs                   # Protocol handlers
```

## Key Classes

| Class | File | Purpose |
|-------|------|---------|
| `DlnaServer` | `DlnaServer.cs` | DLNA server core |
| `ContentDirectory` | `ContentDirectory.cs` | Content browsing |
| `ConnectionManager` | `ConnectionManager.cs` | Connection management |
