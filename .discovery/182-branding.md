# Component: Emby.Server.Implementations — Branding

**Path:** `Emby.Server.Implementations/Branding/`
**Type:** Directory | Sub-module
**Language:** C#
**Maps to:** `.discovery/182-branding.md`

## Description

Server branding and customization configuration. Manages custom logos, colors, and branding elements.

## Files

- `BrandingConfigurationFactory.cs` — Emby.Server.Implementations/Branding/BrandingConfigurationFactory.cs

## Architecture

```mermaid
graph LR
    A[Configuration] --> B[BrandingConfigurationFactory]
    B --> C[Branding Options]
```

## Dependencies

- `MediaBrowser.Controller` — Configuration interfaces
