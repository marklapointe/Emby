# Component: Emby.Dlna — PlayTo

**Path:** `Emby.Dlna/PlayTo/`
**Type:** Directory | Sub-module
**Language:** C#
**Maps to:** `.discovery/333-emby-dlna-playto.md`
**Parent:** `.discovery/330-emby-dlna.md`

## Description

PlayTo functionality for pushing media to DLNA renderers.

## Structure

```
PlayTo/
├── PlayToManager.cs              # [class] PlayToManager
│   └── Manages PlayTo sessions
├── PlayToController.cs           # [class] PlayToController
│   └── Controls remote playback
└── *Device.cs                    # Renderer device classes
```

## Key Classes

| Class | File | Purpose |
|-------|------|---------|
| `PlayToManager` | `PlayToManager.cs` | PlayTo session management |
| `PlayToController` | `PlayToController.cs` | Remote playback control |
