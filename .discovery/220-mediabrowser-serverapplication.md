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

## Dependencies

- Emby.Server.Implementations
- MediaBrowser.Api
- MediaBrowser.WebDashboard
