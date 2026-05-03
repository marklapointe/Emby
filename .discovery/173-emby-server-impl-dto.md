# Component: Emby.Server.Implementations — DTO

**Path:** `Emby.Server.Implementations/Dto/`
**Type:** Directory | Sub-module
**Language:** C#
**Maps to:** `.discovery/173-emby-server-impl-dto.md`
**Parent:** `.discovery/160-emby-server-impl.md`

## Description

Data Transfer Object (DTO) creation and mapping. Converts internal entity models
to API-friendly DTOs for client consumption.

## Structure

```
Dto/
├── DtoService.cs                 # [class] DtoService → IDtoService
│   ├── Main DTO creation service
│   ├── Maps entities to DTOs
│   └── Handles recursive DTO expansion
└── *Dto.cs                       # Various DTO helper classes
```

## Key Classes

| Class | File | Purpose |
|-------|------|---------|
| `DtoService` | `DtoService.cs` | Central DTO creation service |

## Dependencies

- `MediaBrowser.Controller.Dto` — DTO interfaces
- `MediaBrowser.Controller.Entities` — Entity models
- `MediaBrowser.Model.Dto` — DTO model definitions
