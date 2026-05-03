# Component: MediaBrowser.Server.Mono

**Path:** `MediaBrowser.Server.Mono/`
**Type:** Directory | Executable
**Language:** C#
**Maps to:** `.discovery/136-server-mono.md`

## Description

MediaBrowser.Server.Mono is the Mono/Linux entry point for Emby Server. It provides the main executable that bootstraps the server on Unix-like platforms (Linux, macOS, FreeBSD) using the Mono runtime. It includes platform-specific path resolution, image encoder selection, and power management integration.

## Structure

```
MediaBrowser.Server.Mono/
├── MediaBrowser.Server.Mono.csproj
├── Properties/
│   └── AssemblyInfo.cs            # Assembly metadata
├── Program.cs                     # Main entry point
│   └── [class] MainClass
│       ├── [method] public static void Main(string[] args)
│       │   ├── Parses command-line arguments
│       │   ├── Sets up Mono-specific configuration
│       │   ├── Creates MonoAppHost instance
│       │   ├── Calls appHost.RunStartupTasks()
│       │   └── Blocks on Console.ReadLine() to keep process alive
│       └── [method] private static void SetUnixSocketPermissions(...)
│           └── Sets file permissions on Unix domain sockets
├── MonoAppHost.cs                 # Mono-specific application host
│   └── [class] MonoAppHost : ApplicationHost
│       ├── [method] public override void RunStartupTasks()
│       │   ├── Calls base.RunStartupTasks()
│       │   ├── Registers Mono-specific services
│       │   └── Initializes platform-specific features
│       ├── [method] public override string GetType()
│       │   └── Returns "Mono" (platform identifier)
│       ├── [method] public override bool CanSelfRestart
│       │   └── Returns false (Mono cannot self-restart)
│       ├── [method] public override bool CanSelfUpdate
│       │   └── Returns false (Mono cannot self-update)
│       ├── [method] public override bool SupportsRunningAsService
│       │   └── Returns true (supports systemd/service mode)
│       ├── [method] public override bool SupportsLibraryMonitor
│       │   └── Returns true (supports file system watching)
│       ├── [method] public override bool SupportsAutoRunAtStartup
│       │   └── Returns false (no auto-start on Mono)
│       ├── [method] public override bool SupportsRunningAsTrayIcon
│       │   └── Returns false (no tray icon on Mono)
│       ├── [method] public override bool SupportsHttps
│       │   └── Returns true (HTTPS supported via certificates)
│       ├── [method] public override void Shutdown()
│       │   └── Calls Environment.Exit(0)
│       ├── [method] public override void Restart()
│       │   └── Throws NotSupportedException (cannot self-restart)
│       └── [method] public override void ConfigureAutoRun(bool autorun)
│           └── No-op (not supported on Mono)
├── ApplicationPathHelper.cs       # Mono path resolution
│   └── [class] ApplicationPathHelper
│       ├── [method] public static string GetProgramDataPath()
│       │   ├── Checks $XDG_DATA_HOME environment variable
│       │   ├── Falls back to ~/.local/share/emby-server/
│       │   └── Returns absolute data directory path
│       ├── [method] public static string GetProgramSystemPath()
│       │   ├── Checks /usr/share/emby-server/ (system-wide)
│       │   ├── Falls back to application directory
│       │   └── Returns system resource path
│       └── [method] public static string GetMonoExecutablePath()
│           └── Returns path to mono executable
├── ImageEncoderHelper.cs          # Mono image encoder selection
│   └── [class] ImageEncoderHelper
│       ├── [method] public static IImageEncoder GetImageEncoder(ILogger logger)
│       │   ├── Tries Emby.Drawing.Skia (SkiaSharp)
│       │   ├── Tries Emby.Drawing.ImageMagick (ImageMagick)
│       │   ├── Tries Emby.Drawing.Net (System.Drawing)
│       │   └── Returns first available encoder or null
│       └── [method] public static bool IsSkiaAvailable()
│           └── Checks if SkiaSharp native library is present
└── Native/
    └── PowerManagement.cs         # Mono power management
        └── [class] PowerManagement : IPowerManagement
            ├── [method] public void PreventSystemStandby()
            │   ├── Uses dbus to inhibit screensaver/standby
            │   └── Calls org.freedesktop.ScreenSaver.Inhibit
            ├── [method] public void AllowSystemStandby()
            │   └── Releases inhibition via dbus
            └── [method] public void Shutdown()
                └── Calls dbus system shutdown
```

## Platform-Specific Behavior

| Feature | Windows | Mono/Linux |
|---------|---------|------------|
| Self-restart | ✅ | ❌ |
| Self-update | ✅ | ❌ |
| Tray icon | ✅ | ❌ |
| Auto-start | ✅ | ❌ |
| Service mode | ✅ | ✅ |
| Library monitor | ✅ | ✅ |
| HTTPS | ✅ | ✅ |

## Data Flow

```mermaid
graph TD
    A[Program.Main] --&gt; B[Parse CLI args]
    B --&gt; C[Create MonoAppHost]
    C --&gt; D[RunStartupTasks]
    D --&gt; E[Select image encoder]
    E --&gt; F{Skia?}
    F --&gt;|Yes| G[Use SkiaSharp]
    F --&gt;|No| H{ImageMagick?}
    H --&gt;|Yes| I[Use ImageMagick]
    H --&gt;|No| J{System.Drawing?}
    J --&gt;|Yes| K[Use System.Drawing]
    J --&gt;|No| L[No image encoder]
    D --&gt; M[Block on Console.ReadLine]
```

## Side Effects

- Reads environment variables ($XDG_DATA_HOME)
- Accesses file system for path resolution
- Calls dbus for power management (Linux)
- Loads native libraries (SkiaSharp, ImageMagick)
- No external network calls
