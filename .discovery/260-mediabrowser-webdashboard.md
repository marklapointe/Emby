# Component: MediaBrowser.WebDashboard

**Path:** `MediaBrowser.WebDashboard/`
**Type:** Directory | Module
**Language:** C# / HTML / JavaScript
**Maps to:** `.discovery/260-mediabrowser-webdashboard.md`

## Description

MediaBrowser.WebDashboard provides the web-based administrative dashboard for Emby Server. It serves the HTML/CSS/JavaScript frontend that allows users to configure the server, manage libraries, view statistics, and control playback. The dashboard is served as embedded resources and static files.

## Structure

```
MediaBrowser.WebDashboard/
├── MediaBrowser.WebDashboard.csproj
├── Api/                         # Dashboard API endpoints
├── App/                         # Dashboard application files
│   ├── index.html               # Main dashboard page
│   ├── scripts/                 # JavaScript files
│   ├── css/                     # Stylesheets
│   └── ...                      # UI components
└── Properties/                  # Assembly info
```

## Key Components

| Component | Path | Purpose |
|-----------|------|---------|
| `DashboardApp` | `App/` | Web UI frontend |
| `DashboardApi` | `Api/` | Dashboard-specific API |

## Dependencies

- `MediaBrowser.Api` — Backend API
- `MediaBrowser.Controller` — Server interfaces

## Side Effects

- Serves static web assets
- Provides server configuration UI

## Reference

- Dashboard accessed at `http://server:8096/web`
