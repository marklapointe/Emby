# Component: Emby.Server.Implementations — Browser

**Path:** `Emby.Server.Implementations/Browser/`
**Type:** Directory | Sub-module
**Language:** C#
**Maps to:** `.discovery/183-browser.md`

## Description

Browser integration for launching external browsers. Handles URL opening and browser detection.

## Files

- `BrowserLauncher.cs` — Emby.Server.Implementations/Browser/BrowserLauncher.cs

## Architecture

```mermaid
graph LR
    A[Web UI] --> B[BrowserLauncher]
    B --> C[External Browser]
```

## Dependencies

- `System.Diagnostics` — Process management
