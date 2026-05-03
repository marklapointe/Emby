# Component: Emby.Server.Implementations — Localization

**Path:** `Emby.Server.Implementations/Localization/`
**Type:** Directory | Sub-module
**Language:** C#
**Maps to:** `.discovery/190-localization.md`

## Description

Internationalization and localization support. Manages translations and localized strings.

## Files

- `LocalizationManager.cs` — Emby.Server.Implementations/Localization/LocalizationManager.cs
- `TextLocalizer.cs` — Emby.Server.Implementations/Localization/TextLocalizer.cs

## Architecture

```mermaid
graph LR
    A[User Request] --> B[LocalizationManager]
    B --> C[TextLocalizer]
    C --> D[Translation Files]
```

## Key Classes

| Class | Responsibility |
|-------|----------------|
| `LocalizationManager` | Manages available languages |
| `TextLocalizer` | String translation |

## Dependencies

- `MediaBrowser.Model` — Localization models
