# Component: MediaBrowser.ServerApplication

**Path:** `MediaBrowser.ServerApplication/`
**Type:** Directory | Executable
**Language:** C#
**Maps to:** `.discovery/130-server-application.md`

## Description

MediaBrowser.ServerApplication is the Windows desktop entry point for Emby Server. It provides a Windows Forms application with system tray icon, background Windows service support, auto-updater, WebSocket server, native Windows networking, and power management.

## Structure

```
MediaBrowser.ServerApplication/
├── MediaBrowser.ServerApplication.csproj
├── MainStartup.cs                 # Windows main entry point
│   └── [class] MainStartup
│       ├── [method] public static void Main(string[] args)
│       │   ├── Parses command-line arguments
│       │   ├── Creates WindowsAppHost instance
│       │   ├── Runs startup tasks
│       │   └── Shows splash form and system tray icon
│       └── [method] private static void Shutdown()
│           └── Graceful shutdown handler
├── WindowsAppHost.cs              # Windows application host
│   └── [class] WindowsAppHost : ApplicationHost
│       ├── [constructor] WindowsAppHost(IServerApplicationPaths applicationPaths, ILogManager logManager, IFileSystem fileSystem, IPowerManagement powerManagement, IImageEncoder imageEncoder, ISystemEvents systemEvents, INetworkManager networkManager, IConfigurationManager configurationManager, ILocalizationManager localizationManager)
│       ├── [method] public override bool IsRunningAsService
│       │   └── Returns true if running as Windows service
│       ├── [method] public override bool CanSelfRestart
│       │   └── Returns true (Windows can self-restart)
│       ├── [method] public override bool CanSelfUpdate
│       │   └── Returns true (Windows can self-update)
│       ├── [method] public override void Restart()
│       │   └── Restarts the application process
│       ├── [method] public override void SelfUpdate()
│       │   └── Downloads and applies update package
│       ├── [method] public override void Shutdown()
│       │   └── Calls base.Shutdown() and exits process
│       └── [method] public override void ShowSplashScreen()
│           └── Shows Windows splash form
├── ApplicationPathHelper.cs       # Windows application paths
│   └── [class] ApplicationPathHelper
│       └── [method] public static ServerApplicationPaths GetApplicationPaths(string applicationPath, string programDataPath)
│           └── Returns %ProgramData%\Emby-Server paths
├── ImageEncoderHelper.cs          # Windows image encoder selection
│   └── [class] ImageEncoderHelper
│       └── [method] public static IImageEncoder GetImageEncoder(ILogger logger, IFileSystem fileSystem, IApplicationPaths appPaths)
│           ├── Tries to load ImageMagick encoder first
│           ├── Falls back to .NET GDI+ encoder
│           └── Returns IImageEncoder instance
├── BackgroundService.cs           # Windows service wrapper
│   └── [class] BackgroundService : ServiceBase
│       ├── [method] protected override void OnStart(string[] args)
│       │   └── Starts Emby server in service mode
│       └── [method] protected override void OnStop()
│           └── Stops Emby server gracefully
├── BackgroundServiceInstaller.cs  # Windows service installer
│   └── [class] BackgroundServiceInstaller : Installer
│       └── Installs/uninstalls Emby as Windows service
├── ServerNotifyIcon.cs            # System tray icon
│   └── [class] ServerNotifyIcon : IDisposable
│       ├── Shows Emby icon in Windows system tray
│       ├── Provides context menu (Open Dashboard, Restart, Shutdown)
│       └── Handles balloon notifications
├── Splash/
│   ├── SplashForm.cs               # Splash screen form
│   │   └── [class] SplashForm : Form
│   │       └── Shows startup progress and version info
│   └── SplashForm.Designer.cs    # Designer-generated code
├── Updates/
│   └── ApplicationUpdater.cs      # Auto-updater
│       └── [class] ApplicationUpdater
│           ├── [method] public void UpdateApplication(ILogger logger, string updatePackagePath, string applicationPath)
│           │   ├── Extracts update package
│           │   ├── Backs up current installation
│           │   ├── Applies update files
│           │   └── Restarts application
│           └── [method] public void CheckForUpdate(ILogger logger, string updateUrl)
│               └── Downloads update package if available
├── Native/
│   ├── PowerManagement.cs          # Windows power management
│   │   └── [class] PowerManagement : IPowerManagement
│   │       ├── [method] public void PreventSystemStandby()
│   │       │   └── Calls SetThreadExecutionState(ES_SYSTEM_REQUIRED)
│   │       ├── [method] public void AllowSystemStandby()
│   │       │   └── Clears execution state
│   │       └── [method] public void ShutdownSystem()
│   │           └── Calls ExitWindowsEx(EWX_SHUTDOWN)
│   ├── Standby.cs                  # Standby prevention helper
│   │   └── Prevents Windows from entering standby during playback
│   ├── ServerAuthorization.cs    # Windows authorization
│   │   └── Manages Windows user/group permissions
│   ├── LoopUtil.cs                 # Loopback utility
│   │   └── [class] LoopUtil
│   │       └── Manages loopback addresses for UWP apps
│   │       └── [class] AppContainer
│   │           └── Windows app container SID management
│   └── LnkShortcutHandler.cs     # Windows shortcut (.lnk) handler
│       └── [class] LnkShortcutHandler : IShortcutHandler
│           ├── [method] public string ResolveShortcut(string shortcutPath)
│           │   └── Resolves .lnk file target path via IShellLinkW COM
│           └── [interface] IShellLinkW
│               └── Windows Shell Link COM interface
│           └── [class] ShellLink
│               └── COM wrapper for IShellLinkW
├── Networking/
│   ├── NetworkManager.cs           # Windows network manager
│   │   └── [class] NetworkManager : Emby.Server.Implementations.Networking.NetworkManager
│   │       └── Windows-specific network operations
│   ├── NetworkShares.cs            # Windows network shares
│   │   ├── [class] Share
│   │   │   └── Represents a Windows network share (Name, Path, Description)
│   │   └── [class] ShareCollection : ReadOnlyCollectionBase
│   │       └── Collection of network shares
│   └── NativeMethods.cs            # Windows native API P/Invoke
│       └── P/Invoke declarations for netapi32.dll, mpr.dll
└── SocketSharp/
    ├── WebSocketSharpListener.cs   # WebSocket server listener
    │   └── [class] WebSocketSharpListener : IHttpListener
    │       ├── [method] public void Start(IEnumerable<string> urlPrefixes)
    │       │   └── Starts WebSocketSharp HTTP listener
    │       └── [method] public void Stop()
    │           └── Stops listener
    ├── WebSocketSharpRequest.cs    # WebSocket request wrapper
    │   └── [class] WebSocketSharpRequest : IHttpRequest
    │       └── Wraps WebSocketSharp HttpListenerRequest
    ├── WebSocketSharpResponse.cs   # WebSocket response wrapper
    │   └── [class] WebSocketSharpResponse : IHttpResponse
    │       └── Wraps WebSocketSharp HttpListenerResponse
    ├── SharpWebSocket.cs           # WebSocket connection
    │   └── [class] SharpWebSocket : IWebSocket
    │       ├── [method] public void Send(string message)
    │       │   └── Sends message via WebSocket
    │       └── [method] public void Close()
    │           └── Closes WebSocket connection
    └── RequestMono.cs            # Mono request compatibility
        └── [class] Element
            └── HTTP request element for Mono compatibility
```

## Data Flow

```mermaid
graph TD
    A[Emby.ServerApplication.exe] --&gt; B[MainStartup.Main]
    B --&gt; C[ApplicationPathHelper.GetApplicationPaths]
    C --&gt; D[WindowsAppHost]
    D --&gt; E[ImageEncoderHelper.GetImageEncoder]
    E --&gt; F{ImageMagick available?}
    F --&gt;|Yes| G[ImageMagickEncoder]
    F --&gt;|No| H[ImageEncoder]
    D --&gt; I[ShowSplashScreen]
    I --&gt; J[SplashForm]
    D --&gt; K[ServerNotifyIcon]
    D --&gt; L[RunStartupTasks]
    L --&gt; M[Server starts]
    N[WebSocketSharpListener] --&gt; O[WebSocket connections]
```

## Platform Differences

| Feature | Windows (.NET) | Mono/Linux |
|---------|---------------|------------|
| Service mode | Yes (BackgroundService) | No |
| Self-restart | Yes | No |
| Self-update | Yes (ApplicationUpdater) | No |
| Standby prevention | Yes (PowerManagement) | No |
| System tray icon | Yes (ServerNotifyIcon) | No |
| Splash screen | Yes (SplashForm) | No |
| WebSocket server | Yes (WebSocketSharp) | No |
| Network shares | Yes (NetworkShares) | No |

## Side Effects

- Creates %ProgramData%\Emby-Server directory
- Installs Windows service (optional)
- Shows system tray icon
- Prevents Windows standby during playback
- Downloads and applies auto-updates
- Creates/restores .lnk shortcuts
