# Component: Emby.Server.Implementations — Browser

**Path:** `Emby.Server.Implementations/Browser/`
**Type:** Directory | Module
**Language:** C#
**Maps to:** `.discovery/225-emby-server-impl-browser.md`

## Description

Web browser launcher integration. Opens media items in the system's default web browser.

## Files

- `BrowserLauncher.cs` — Emby.Server.Implementations/Browser/BrowserLauncher.cs

## Decomposition

### BrowserLauncher.cs (Browser Launcher)

#### Imports
```csharp
using MediaBrowser.Controller.Entities;
using System;
using System.Diagnostics;
```

#### Classes
`BrowserLauncher` (public static class)

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `Launch(BaseItem)` | `void` | Open item in browser |
| `LaunchUrl(string)` | `void` | Open URL in browser |

## Data Flow

```mermaid
graph LR
    A[BaseItem] --> B[BrowserLauncher]
    B --> C[Process.Start]
    C --> D[System Browser]
```

## Dependencies

- `System.Diagnostics` — Process management
- `MediaBrowser.Controller.Entities` — Media item types

## Statistics

| Metric | Value |
|--------|-------|
| Files | 1 |
| Classes | 1 |
| LOC | ~30 |
