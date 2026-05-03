# Component: MediaBrowser.Server.Mono

**Path:** `MediaBrowser.Server.Mono/`
**Type:** Directory | Executable
**Language:** C#
**Maps to:** `.discovery/129-server-mono.md`

## Description

MediaBrowser.Server.Mono is the Mono/Linux entry point for Emby Server. It provides the main executable that runs on Mono runtime, with platform-specific implementations for application paths, image encoding, and power management.

## Structure

```
MediaBrowser.Server.Mono/
├── MediaBrowser.Server.Mono.csproj
├── Program.cs                     # Main entry point
│   └── [class] MainClass
│       ├── [method] public static void Main(string[] args)
│       │   ├── Parses command-line arguments
│       │   ├── Creates MonoAppHost instance
│       │   ├── Calls appHost.RunStartupTasks()
│       │   └── Blocks on Console.ReadLine() or waits for shutdown signal
│       └── [method] private static void Shutdown()
│           └── Graceful shutdown handler
├── MonoAppHost.cs                 # Mono-specific application host
│   └── [class] MonoAppHost : ApplicationHost
│       ├── [constructor] MonoAppHost(IServerApplicationPaths applicationPaths, ILogManager logManager, IFileSystem fileSystem, IPowerManagement powerManagement, IImageEncoder imageEncoder, ISystemEvents systemEvents, INetworkManager networkManager, IConfigurationManager configurationManager, ILocalizationManager localizationManager)
│       ├── [method] public override bool IsRunningAsService
│       │   └── Returns false (Mono runs as console app)
│       ├── [method] public override bool CanSelfRestart
│       │   └── Returns false (Mono cannot self-restart)
│       ├── [method] public override bool CanSelfUpdate
│       │   └── Returns false (Mono cannot self-update)
│       ├── [method] public override void Shutdown()
│       │   └── Calls base.Shutdown() and exits process
│       ├── [method] public override void Restart()
│       │   └── Throws NotImplementedException
│       └── [method] public override void SelfUpdate()
│           └── Throws NotImplementedException
├── ApplicationPathHelper.cs       # Mono application paths
│   └── [class] ApplicationPathHelper
│       ├── [method] public static ServerApplicationPaths GetApplicationPaths(string applicationPath, string programDataPath)
│       │   ├── Resolves application directory
│       │   ├── Sets up program data path (default: ~/.config/emby-server)
│       │   └── Returns ServerApplicationPaths instance
│       └── [method] public static string GetProgramDataPath()
│           └── Returns ~/.config/emby-server or $XDG_CONFIG_HOME/emby-server
├── ImageEncoderHelper.cs          # Mono image encoder selection
│   └── [class] ImageEncoderHelper
│       ├── [method] public static IImageEncoder GetImageEncoder(ILogger logger, IFileSystem fileSystem, IApplicationPaths appPaths)
│       │   ├── Tries to load ImageMagick encoder first
│       │   ├── Falls back to .NET GDI+ encoder
│       │   └── Returns IImageEncoder instance
│       └── [method] public static bool ImageMagickAvailable()
│           └── Checks if ImageMagick native library is available
└── Native/
    └── PowerManagement.cs         # Mono power management
        └── [class] PowerManagement : IPowerManagement
            ├── [method] public void PreventSystemStandby()
            │   └── No-op on Mono (no standby prevention)
            ├── [method] public void AllowSystemStandby()
            │   └── No-op on Mono
            └── [method] public void ShutdownSystem()
                └── No-op on Mono
```

## Data Flow

```mermaid
graph TD
    A[mono MediaBrowser.Server.Mono.exe] --&gt; B[MainClass.Main]
    B --&gt; C[ApplicationPathHelper.GetApplicationPaths]
    C --&gt; D[MonoAppHost]
    D --&gt; E[ImageEncoderHelper.GetImageEncoder]
    E --&gt; F{ImageMagick available?}
    F --&gt;|Yes| G[ImageMagickEncoder]
    F --&gt;|No| H[ImageEncoder]
    D --&gt; I[appHost.RunStartupTasks]
    I --&gt; J[Server starts]
```

## Platform Differences

| Feature | Windows (.NET) | Mono/Linux |
|---------|---------------|------------|
| Service mode | Yes | No |
| Self-restart | Yes | No |
| Self-update | Yes | No |
| Standby prevention | Yes | No |
| Default data path | %ProgramData% | ~/.config/emby-server |
| ImageMagick | Optional | Preferred |

## Side Effects

- Creates ~/.config/emby-server directory
- Loads ImageMagick native library (if available)
- Blocks on Console.ReadLine() until shutdown
