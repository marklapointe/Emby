# Component: MediaBrowser.WebDashboard

**Path:** \`MediaBrowser.WebDashboard/\`
**Type:** Directory | Module
**Language:** C#, JavaScript, HTML, CSS
**Maps to:** \`.discovery/260-mediabrowser-webdashboard.md\`

## Decomposition

### DashboardService.cs (Main API Endpoint)

#### Imports
```csharp
using MediaBrowser.Model.Services;
using System.Threading.Tasks;
```

#### Classes
\`DashboardService\` (public class : IService)

#### Key Methods
```csharp
Task<object> Get(DashboardInfo request)
```

### WebSocketListener.cs (WebSocket Handler)

#### Classes
\`WebSocketListener\` (public class)

#### Key Methods
```csharp
void SendProgress(object progress)
void SendRestartNotification()
```

### index.html (Main Entry Point)

#### Structure
```html
<!DOCTYPE html>
<html>
<head>
    <script src="dashboard.js"></script>
</head>
<body>
    <div id="dashboard-app"></div>
</body>
</html>
```

### main.js (Application Bootstrap)

#### Key Functions
```javascript
function initializeDashboard()
function loadLibrary()
function setupWebSocket()
```

### apphost.js (Emby Client Host)

#### Classes/Functions
```javascript
class AppHost extends Events
    connect()
    getApiClient()
    getCurrentUser()
```

### globalize.js (i18n Support)

#### Key Functions
```javascript
function t(key, params)
function cultureInfo()
```

## Description

MediaBrowser.WebDashboard provides the web-based management interface for Emby Server. It includes a C# backend API and a rich HTML5/JavaScript frontend. The dashboard allows users to configure server settings, manage libraries, view activity, and control playback.

## Sub-Modules

| Sub-Module | Document | Description |
|------------|----------|-------------|
| API & Backend | [261-mediabrowser-webdashboard-api.md](./261-mediabrowser-webdashboard-api.md) | C# API services, entry point, project files |
| UI Structure | [262-mediabrowser-webdashboard-ui.md](./262-mediabrowser-webdashboard-ui.md) | HTML pages, CSS, dashboard views |
| Scripts | [263-mediabrowser-webdashboard-scripts.md](./263-mediabrowser-webdashboard-scripts.md) | JavaScript modules (91 files) |
| Components | [264-mediabrowser-webdashboard-components.md](./264-mediabrowser-webdashboard-components.md) | Reusable web components |
| Localization | [265-mediabrowser-webdashboard-strings.md](./265-mediabrowser-webdashboard-strings.md) | i18n string files (46 languages) |
| Bower Dependencies | [266-mediabrowser-webdashboard-bower.md](./266-mediabrowser-webdashboard-bower.md) | Third-party frontend libraries |

## Statistics

- Total files: 670
- Project files (non-bower): 308
- Bower dependencies: 362
- C# backend files: 5
- JavaScript files: 91
- CSS files: 15
- HTML files: ~20
- Localization files: 46
