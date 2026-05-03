# Component: MediaBrowser.ServerApplication

**Path:** \`MediaBrowser.ServerApplication/\`
**Type:** Module
**Maps to:** \`.discovery/220-mediabrowser-serverapplication.md\`

## Description

Windows desktop server application entry point and UI.

## Structure

```
MediaBrowser.ServerApplication/
├── MediaBrowser.ServerApplication.csproj
├── MainStartup.cs            # Application entry point
├── ServerApplicationPaths.cs # Application paths
├── NativeApp.cs              # Native app wrapper
├── SystemEvents.cs           # System event handling
├── WindowsApp.cs             # Windows-specific app
├── SplashForm.cs             # Splash screen
├── Program.cs                # Main program
├── Properties/
│   └── AssemblyInfo.cs
└── packages.config
```

## Decomposition

### Program.cs (Main Entry Point)

#### Imports
```csharp
using MediaBrowser.Server.Implementations;
using System;
using System.Threading.Tasks;
```

#### Classes
`Program` (internal static class)

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `Main(string[])` | `int` | Application entry point |
| `RunServerStartupOptions(ServerStartupOptions)` | `Task` | Run server |

### MainStartup.cs (Application Startup)

#### Classes
`MainStartup` (public static class)

#### Key Properties
| Property | Type | Description |
|----------|------|-------------|
| `StartupOptions` | `StartupOptions` | Server startup options |

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `Configure(StartupOptions)` | `void` | Configure application |
| `Run(StartupOptions)` | `Task` | Run application |

### ServerApplicationPaths.cs (Path Configuration)

#### Classes
`ServerApplicationPaths` (public class : BaseApplicationPaths)

#### Key Properties
| Property | Type | Description |
|----------|------|-------------|
| `ProgramDataPath` | `string` | Program data directory |
| `WebSocketPort` | `int` | WebSocket port |
| `EnableUPnP` | `bool` | UPnP enable flag |

## Dependencies

- Emby.Server.Implementations
- MediaBrowser.Api
- MediaBrowser.WebDashboard
