# Component: MediaBrowser.ServerApplication

**Path:** `MediaBrowser.ServerApplication/`
**Type:** Directory | Executable
**Language:** C#
**Maps to:** `.discovery/137-server-application.md`

## Description

MediaBrowser.ServerApplication is the Windows entry point for Emby Server. It provides the main executable that bootstraps the server on Windows with full desktop integration: system tray icon, Windows service support, auto-startup, self-update, power management, and native networking. It also includes a WebSocket server implementation (SocketSharp) for real-time client communication.

## Structure

```
MediaBrowser.ServerApplication/
├── MediaBrowser.ServerApplication.csproj
├── Properties/
│   ├── AssemblyInfo.cs            # Assembly metadata
│   └── Resources.Designer.cs      # Windows resources
├── MainStartup.cs                 # Main entry point
│   └── [class] MainStartup
│       ├── [method] public static void Main(string[] args)
│       │   ├── Parses command-line arguments
│       │   ├── Creates WindowsAppHost instance
│       │   ├── Calls appHost.RunStartupTasks()
│       │   └── Blocks on Console.ReadLine()
│       └── [method] private static void SetUnixSocketPermissions(...)
│           └── No-op on Windows
├── WindowsAppHost.cs              # Windows-specific application host
│   └── [class] WindowsAppHost : ApplicationHost
│       ├── [method] public override void RunStartupTasks()
│       │   ├── Calls base.RunStartupTasks()
│       │   ├── Creates ServerNotifyIcon (system tray)
│       │   ├── Sets up auto-start registry keys
│       │   └── Initializes Windows-specific features
│       ├── [method] public override string GetType()
│       │   └── Returns "Windows" (platform identifier)
│       ├── [method] public override bool CanSelfRestart
│       │   └── Returns true (supports self-restart)
│       ├── [method] public override bool CanSelfUpdate
│       │   └── Returns true (supports self-update)
│       ├── [method] public override bool SupportsRunningAsService
│       │   └── Returns true (Windows service mode)
│       ├── [method] public override bool SupportsLibraryMonitor
│       │   └── Returns true (file system watcher)
│       ├── [method] public override bool SupportsAutoRunAtStartup
│       │   └── Returns true (registry auto-start)
│       ├── [method] public override bool SupportsRunningAsTrayIcon
│       │   └── Returns true (system tray icon)
│       ├── [method] public override bool SupportsHttps
│       │   └── Returns true (HTTPS via certificates)
│       ├── [method] public override void Shutdown()
│       │   └── Calls Application.Exit()
│       ├── [method] public override void Restart()
│       │   └── Spawns new process and exits current
│       ├── [method] public override void ConfigureAutoRun(bool autorun)
│       │   ├── Sets HKCU\Software\Microsoft\Windows\CurrentVersion\Run registry key
│       │   └── Adds/removes Emby from Windows startup
│       └── [method] public override void UpdateApplication(...)
│           └── Downloads and applies update package
├── ApplicationPathHelper.cs       # Windows path resolution
│   └── [class] ApplicationPathHelper
│       ├── [method] public static string GetProgramDataPath()
│       │   ├── Returns %ProgramData%\Emby-Server\
│       │   └── Standard Windows application data path
│       └── [method] public static string GetProgramSystemPath()
│           └── Returns application installation directory
├── ImageEncoderHelper.cs          # Windows image encoder selection
│   └── [class] ImageEncoderHelper
│       ├── [method] public static IImageEncoder GetImageEncoder(ILogger logger)
│       │   ├── Tries Emby.Drawing.Skia (SkiaSharp)
│       │   ├── Tries Emby.Drawing.ImageMagick (ImageMagick)
│       │   ├── Tries Emby.Drawing.Net (System.Drawing)
│       │   └── Returns first available encoder or null
│       └── [method] public static bool IsSkiaAvailable()
│           └── Checks if SkiaSharp native library is present
├── BackgroundService.cs           # Windows service wrapper
│   └── [class] BackgroundService : ServiceBase
│       ├── [method] protected override void OnStart(string[] args)
│       │   └── Calls MainStartup.Main(args)
│       └── [method] protected override void OnStop()
│           └── Calls appHost.Shutdown()
├── BackgroundServiceInstaller.cs  # Windows service installer
│   └── [class] BackgroundServiceInstaller : Installer
│       └── [method] public BackgroundServiceInstaller()
│           └── Configures ServiceProcessInstaller and ServiceInstaller
├── ServerNotifyIcon.cs            # System tray icon
│   └── [class] ServerNotifyIcon : IDisposable
│       ├── [method] public void Show()
│       │   ├── Creates NotifyIcon in system tray
│       │   ├── Adds context menu (Open Dashboard, Restart, Exit)
│       │   └── Shows balloon tip on startup
│       ├── [method] public void Hide()
│       │   └── Hides NotifyIcon
│       └── [method] public void Dispose()
│           └── Disposes NotifyIcon resources
├── Updates/
│   └── ApplicationUpdater.cs      # Self-update mechanism
│       └── [class] ApplicationUpdater
│           ├── [method] public void UpdateApplication(...)
│           │   ├── Downloads update package from Emby update server
│           │   ├── Verifies package signature
│           │   ├── Extracts to temp directory
│           │   └── Restarts application with new version
│           └── [method] public bool IsUpdateAvailable(...)
│               └── Compares current version with latest online
├── Splash/
│   ├── SplashForm.cs              # Startup splash screen
│   │   └── Windows Forms splash form showing startup progress
│   └── SplashForm.Designer.cs     # Designer-generated code
├── Native/
│   ├── PowerManagement.cs         # Windows power management
│   │   └── [class] PowerManagement : IPowerManagement
│   │       ├── [method] public void PreventSystemStandby()
│   │       │   └── Calls SetThreadExecutionState(ES_SYSTEM_REQUIRED)
│   │       ├── [method] public void AllowSystemStandby()
│   │       │   └── Calls SetThreadExecutionState(ES_CONTINUOUS)
│   │       └── [method] public void Shutdown()
│   │           └── Calls ExitWindowsEx(EWX_SHUTDOWN)
│   ├── Standby.cs                 # Standby prevention
│   │   └── P/Invoke to SetThreadExecutionState
│   ├── ServerAuthorization.cs     # Windows authorization
│   │   └── ACL management for server directories
│   ├── LnkShortcutHandler.cs      # Windows shortcut (.lnk) parser
│   │   ├── [class] LnkShortcutHandler : IShortcutHandler
│   │   │   ├── [method] public string ResolveShortcut(string path)
│   │   │   │   └── Reads .lnk file target path via IShellLinkW COM
│   │   │   └── [method] public bool IsShortcut(string path)
│   │   │       └── Returns true for .lnk files
│   │   ├── [interface] IShellLinkW
│   │   └── [class] ShellLink
│   │       └── COM wrapper for IShellLinkW
│   └── LoopUtil.cs                # Windows loopback utility
│       ├── [class] LoopUtil
│       │   └── [method] public static void RegisterAppContainerLoopback(...)
│       │       └── Configures Windows firewall loopback for UWP apps
│       └── [class] AppContainer
│           └── Windows AppContainer SID wrapper
├── Networking/
│   ├── NetworkManager.cs          # Windows network manager
│   │   └── [class] NetworkManager : Emby.Server.Implementations.Networking.NetworkManager
│   │       ├── [method] public override string GetInternalIp()
│   │       │   └── Returns primary local IP address
│   │       └── [method] public override bool IsInLocalNetwork(string remoteIp)
│   │           └── Checks if IP is in local subnet
│   ├── NetworkShares.cs           # Windows network shares
│   │   ├── [class] Share
│   │   │   ├── Properties: Name, Path, Description, Server
│   │   │   └── Represents a Windows SMB share
│   │   └── [class] ShareCollection : ReadOnlyCollectionBase
│   │       └── [method] public static ShareCollection GetShares(...)
│   │           └── Enumerates Windows network shares via NetShareEnum
│   └── NativeMethods.cs           # Windows networking P/Invoke
│       └── P/Invoke declarations for NetShareEnum, NetApiBufferFree
└── SocketSharp/                   # WebSocket server (SocketSharp)
    ├── WebSocketSharpListener.cs  # WebSocket HTTP listener
    │   └── [class] WebSocketSharpListener : IHttpListener
    │       ├── [method] public void Start(...)
    │       │   └── Starts WebSocketSharp HTTP server
    │       ├── [method] public void Stop()
    │       │   └── Stops HTTP server
    │       └── [method] public Task<IHttpContext> GetContextAsync()
    │           └── Returns WebSocketSharpRequest/Response context
    ├── SharpWebSocket.cs          # WebSocket connection
    │   └── [class] SharpWebSocket : IWebSocket
    │       ├── [method] public void Send(string message)
    │       │   └── Sends text message via WebSocket
    │       ├── [method] public void Send(byte[] message)
    │       │   └── Sends binary message via WebSocket
    │       └── [method] public void Close()
    │           └── Closes WebSocket connection
    ├── WebSocketSharpRequest.cs   # WebSocket HTTP request wrapper
    │   └── [class] WebSocketSharpRequest : IHttpRequest
    │       ├── Wraps WebSocketSharp HttpListenerRequest
    │       └── [class] HttpFile : IHttpFile
    │           └── Represents uploaded file in multipart request
    ├── WebSocketSharpResponse.cs  # WebSocket HTTP response wrapper
    │   └── [class] WebSocketSharpResponse : IHttpResponse
    │       └── Wraps WebSocketSharp HttpListenerResponse
    └── RequestMono.cs             # Mono compatibility for requests
        └── [class] Element
            └── Internal request element parsing
```

## Platform-Specific Behavior

| Feature | Windows | Mono/Linux |
|---------|---------|------------|
| Self-restart | ✅ | ❌ |
| Self-update | ✅ | ❌ |
| Tray icon | ✅ | ❌ |
| Auto-start | ✅ (registry) | ❌ |
| Service mode | ✅ | ✅ |
| Library monitor | ✅ | ✅ |
| HTTPS | ✅ | ✅ |
| Splash screen | ✅ | ❌ |
| Windows shortcuts | ✅ | ❌ |
| Network shares | ✅ | ❌ |

## Data Flow

```mermaid
graph TD
    A[Program.Main] --&gt; B[Create WindowsAppHost]
    B --&gt; C[RunStartupTasks]
    C --&gt; D[Create ServerNotifyIcon]
    C --&gt; E[Setup auto-start registry]
    C --&gt; F[Select image encoder]
    D --&gt; G[System tray menu]
    G --&gt; H[Open Dashboard]
    G --&gt; I[Restart Server]
    G --&gt; J[Exit]
    F --&gt; K{Skia?}
    K --&gt;|Yes| L[Use SkiaSharp]
    K --&gt;|No| M{ImageMagick?}
    M --&gt;|Yes| N[Use ImageMagick]
    M --&gt;|No| O{System.Drawing?}
    O --&gt;|Yes| P[Use System.Drawing]
    O --&gt;|No| Q[No image encoder]
```

## Side Effects

- Reads registry (auto-start configuration)
- Writes registry (auto-start keys)
- Creates system tray icon (Windows Forms)
- Creates splash screen (Windows Forms)
- Downloads update packages from Emby update server
- Calls Windows P/Invoke APIs (power management, networking, shortcuts)
- Loads native libraries (SkiaSharp, ImageMagick)
- Manages Windows firewall rules (loopback)
- No external network calls (except update check/download)
