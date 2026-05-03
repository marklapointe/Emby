# Codebase Discovery — Table of Contents

**Project:** Emby Server
**Generated:** 2026-05-03
**Root:** MediaBrowser.sln
**Total files:** 4507
**Mapped files:** 94 (2.1% - INCOMPLETE)
**Coverage:** ⚠️ IN PROGRESS - Adding decomposition to modules

---

## 📋 Master File Tracking Table

> **CRITICAL:** EVERY file in the project MUST appear in this table.
> NO file is "not interesting" — every file has a purpose.
> 
> **STATUS:** This table is being progressively updated. Currently at 94/4507 files.

### Source Code Files (C#)

| File | Mapped | Discovery Doc | Where Used | Purpose |
|------|--------|---------------|------------|---------|
| `BDInfo/BDROM.cs` | ✅ | [100-01-bdrom.md](./100-01-bdrom.md) | DvdLib (reference) | Main BD-ROM disc parser + **decomposition** |
| `BDInfo/TSPlaylistFile.cs` | ✅ | [100-02-tsplaylistfile.md](./100-02-tsplaylistfile.md) | BDROM (imports) | Playlist file parser + **decomposition** |
| `BDInfo/TSStreamFile.cs` | ✅ | [100-03-tsstreamfile.md](./100-03-tsstreamfile.md) | BDROM (imports) | Transport stream file reader |
| `BDInfo/TSStreamClipFile.cs` | ✅ | [100-04-tsstreamclipfile.md](./100-04-tsstreamclipfile.md) | BDROM (imports) | Stream clip file parser |
| `BDInfo/LanguageCodes.cs` | ✅ | [100-05-languagecodes.md](./100-05-languagecodes.md) | BDROM (static) | ISO 639-2 language codes |
| `BDInfo/TSStream.cs` | ✅ | [100-06-tsstream.md](./100-06-tsstream.md) | TSStreamFile (imports) | Transport stream definition |
| `BDInfo/TSStreamClip.cs` | ✅ | [100-07-tsstreamclip.md](./100-07-tsstreamclip.md) | TSStreamClipFile (imports) | Stream clip definition |
| `BDInfo/TSInterleavedFile.cs` | ✅ | [100-08-tsinterleavedfile.md](./100-08-tsinterleavedfile.md) | BDROM (imports) | Interleaved stream file |
| `BDInfo/BDInfoSettings.cs` | ✅ | [100-09-bdinfosettings.md](./100-09-bdinfosettings.md) | BDROM (static) | Configuration settings |
| `BDInfo/TSStreamBuffer.cs` | ✅ | [100-10-tsstreambuffer.md](./100-10-tsstreambuffer.md) | TSStream (imports) | Stream buffer implementation |
| `BDInfo/TSCodec*.cs` | ✅ | [100-11-codecs.md](./100-11-codecs.md) | TSStream (codecs) | Audio/video codec handlers |
| `BDInfo/BDInfo.csproj` | ✅ | [100-bdinfo.md](./100-bdinfo.md) | Solution (project) | BDInfo project file |
| `BDInfo/packages.config` | ✅ | [100-bdinfo.md](./100-bdinfo.md) | BDInfo (deps) | NuGet package references |
| `BDInfo/Properties/AssemblyInfo.cs` | ✅ | [101-bdinfo-internals.md](./101-bdinfo-internals.md) | Build (auto) | Assembly metadata |
| `DvdLib/DvdLib.csproj` | ✅ | [110-dvdlib.md](./110-dvdlib.md) | Solution (project) | DvdLib project file |
| `DvdLib/DvdReader.cs` | ✅ | [111-dvdliv-internals.md](./111-dvdliv-internals.md) | Emby.Server (imports) | DVD disc reader |
| `DvdLib/IfoCodec*.cs` | ✅ | [111-dvdliv-internals.md](./111-dvdliv-internals.md) | DvdReader (codecs) | IFO parsing |
| `Emby.Drawing/Emby.Drawing.csproj` | ✅ | [120-emby-drawing.md](./120-emby-drawing.md) | Solution (project) | Drawing project |
| `Emby.Drawing.ImageMagick/*.cs` | ✅ | [121-emby-drawing-imagemagick.md](./121-emby-drawing-imagemagick.md) | Emby.Drawing (backend) | ImageMagick backend |
| `Emby.Drawing.Net/*.cs` | ✅ | [122-emby-drawing-net.md](./122-emby-drawing-net.md) | Emby.Drawing (backend) | .NET drawing backend |
| `Emby.Drawing.Skia/*.cs` | ✅ | [123-emby-drawing-skia.md](./123-emby-drawing-skia.md) | Emby.Drawing (backend) | SkiaSharp backend |
| `Emby.Notifications/*.cs` | ✅ | [131-emby-notifications-internals.md](./131-emby-notifications-internals.md) | Server (notifications) | Notification system |
| `Emby.Server.Implementations/**/*.cs` | ✅ With Decomposition | Multiple docs | Server core | Core implementation + **decomposition** |
| `Emby.Dlna/**/*.cs` | ✅ With Decomposition | [330-emby-dlna.md](./330-emby-dlna.md) | DLNA | DLNA/UPnP + **decomposition** |
| `MediaBrowser.Api/**/*.cs` | ✅ With Decomposition | [340-mediabrowser-api.md](./340-mediabrowser-api.md) | HTTP API | REST endpoints + **decomposition** |
| `MediaBrowser.Providers/**/*.cs` | ✅ With Decomposition | [320-mediabrowser-providers.md](./320-mediabrowser-providers.md) | Metadata | Content providers + **decomposition** |
| `emby-go/**/*.go` | ⚠️ Partial | [360-emby-go.md](./360-emby-go.md) | Emby.Server (remote) | Go service client |
| `MediaBrowser.WebDashboard/**/*.{js,html,css}` | ⚠️ Partial | [261-266](./260-mediabrowser-webdashboard.md) | HTTP (static) | Web UI |
| *(remaining 4413 files)* | ❌ | TBD | - | Pending mapping |

### Configuration & Project Files

| File | Mapped | Purpose |
|------|--------|---------|
| `MediaBrowser.sln` | ✅ | Visual Studio solution file |
| `SharedVersion.cs` | ✅ | Shared version constants |
| `README.md` | ✅ | Project documentation |
| `CONTRIBUTORS.md` | ✅ | Contributors list |
| `LICENSE.md` | ✅ | License file |

---

## Coverage Statistics

| Category | Total | Mapped | With Decomposition |
|----------|-------|--------|-------------------|
| C# Source Files | 1019 | ~50 | ~100 |

### Modules With Decomposition (Updated 2026-05-03)

| Module | Status | Discovery Doc |
|--------|--------|--------------|
| BDInfo | ✅ Decomposition | [100-01-bdrom.md](./100-01-bdrom.md) |
| DvdLib | ✅ Decomposition | [111-dvdliv-internals.md](./111-dvdliv-internals.md) |
| Emby.Server.Implementations | ✅ Decomposition | [161-emby-server-impl-core.md](./161-emby-server-impl-core.md) |
| Emby.Server.Implementations/LiveTV | ✅ Decomposition | [189-livetv-full.md](./189-livetv-full.md) |
| Emby.Dlna | ✅ Decomposition | [330-emby-dlna.md](./330-emby-dlna.md) |
| MediaBrowser.Api | ✅ Decomposition | [340-mediabrowser-api.md](./340-mediabrowser-api.md) |
| MediaBrowser.Providers | ✅ Decomposition | [320-mediabrowser-providers.md](./320-mediabrowser-providers.md) |
| MediaBrowser.WebDashboard | ✅ Decomposition | [260-mediabrowser-webdashboard.md](./260-mediabrowser-webdashboard.md) |
| emby-go | ✅ Decomposition | [360-emby-go.md](./360-emby-go.md) |

---

## Gaps Identified

### Critical (Must Fix)
1. **1019 C# files** - Only ~50 documented (5%)
2. **72 Go files** - Only ~5 documented (7%)
3. **349 JS/TS files** - Only ~10 documented (3%)
4. **1081 XML files** - Not systematically documented

### Required by Codebase-Mapper Skill
1. ⚠️ Master File Tracking Table - In progress
2. ❌ 100% coverage not achieved
3. ✅ Source file decomposition - ADDING decomposition to modules
4. ⚠️ Cross-reference tracking - In progress
5. ⚠️ Entry points - Identified

---

## Next Steps (Priority Order)

1. Add decomposition to remaining modules (BDInfo, DvdLib, etc.)
2. Complete Master File Tracking Table with ALL 4507 files
3. Build cross-reference dependency graph
4. Identify all entry points
5. Verify 100% coverage

---

## Project Structure

```
Emby/
├── BDInfo/                          → .discovery/100-bdinfo.md, 101-bdinfo-internals.md
├── DvdLib/                          → .discovery/110-dvdlib.md, 111-dvdliv-internals.md
├── Emby.Dlna/                       → .discovery/330-emby-dlna.md
├── Emby.Drawing/                    → .discovery/120-emby-drawing.md
│   ├── Emby.Drawing.ImageMagick/    → .discovery/121-emby-drawing-imagemagick.md
│   ├── Emby.Drawing.Net/            → .discovery/122-emby-drawing-net.md
│   └── Emby.Drawing.Skia/          → .discovery/123-emby-drawing-skia.md
├── Emby.Notifications/              → .discovery/140-emby-notifications.md, 131-emby-notifications-internals.md
├── Emby.Photos/                     → .discovery/150-emby-photos.md
├── Emby.Server.Implementations/     → .discovery/160-emby-server-impl.md
│   ├── Core Infrastructure          → .discovery/161-emby-server-impl-core.md
│   ├── Library Management           → .discovery/162-emby-server-impl-library.md
│   ├── Media & Channels            → .discovery/163-emby-server-impl-media.md
│   ├── HTTP Server & Services       → .discovery/164-emby-server-impl-http.md
│   ├── Scheduled Tasks              → .discovery/165-emby-server-impl-tasks.md
│   ├── I/O Utilities               → .discovery/166-emby-server-impl-io.md
│   ├── Text Encoding & Localization → .discovery/167-emby-server-impl-encoding.md
│   ├── Security & Users             → .discovery/168-emby-server-impl-security.md
│   ├── SharpCifs (Embedded)        → .discovery/169-emby-server-impl-sharpcifs.md
│   ├── LiveTV                       → .discovery/170-emby-server-impl-livetv.md, 189-livetv-full.md
│   ├── Data & Persistence           → .discovery/171-emby-server-impl-data.md
│   ├── Resolvers                   → .discovery/172-emby-server-impl-resolvers.md
│   ├── DTO Service                 → .discovery/173-emby-server-impl-dto.md, 185-dto.md
│   ├── Images                      → .discovery/174-emby-server-impl-images.md
│   ├── HTTP Client                 → .discovery/175-emby-server-impl-httpclient.md
│   ├── FFMpeg                      → .discovery/176-emby-server-impl-ffmpeg.md, 187-ffmpeg.md
│   ├── Networking                  → .discovery/178-emby-server-impl-networking.md
│   ├── Activity                    → .discovery/180-activity.md
│   ├── Archiving                   → .discovery/181-archiving.md
│   ├── Branding                    → .discovery/182-branding.md
│   ├── Browser                     → .discovery/183-browser.md
│   ├── Devices                     → .discovery/184-devices.md
│   ├── EntryPoints                 → .discovery/186-entrypoints.md
│   ├── Localization                → .discovery/190-localization.md
│   ├── TV                          → .discovery/191-tv.md
│   ├── UDP                         → .discovery/192-udp.md
│   └── UserViews                   → .discovery/179-emby-server-impl-userviews.md, 193-userviews.md
├── MediaBrowser.Api/                → .discovery/340-mediabrowser-api.md
│   ├── Controllers                 → .discovery/341-mediabrowser-api-controllers.md
│   ├── Models                      → .discovery/342-mediabrowser-api-models.md
│   └── Services                    → .discovery/343-mediabrowser-api-services.md
├── MediaBrowser.LocalMetadata/      → .discovery/210-mediabrowser-localmetadata.md
├── MediaBrowser.Providers/          → .discovery/320-mediabrowser-providers.md
│   ├── Movies                      → .discovery/321-mediabrowser-providers-movies.md
│   ├── TV                          → .discovery/322-mediabrowser-providers-tv.md
│   ├── Music                       → .discovery/323-mediabrowser-providers-music.md
│   ├── Images                      → .discovery/324-mediabrowser-providers-images.md
│   ├── People                      → .discovery/325-mediabrowser-providers-people.md
│   ├── Books                       → .discovery/326-mediabrowser-providers-books.md
│   ├── Subtitles                   → .discovery/328-mediabrowser-providers-subtitles.md
│   ├── Videos                      → .discovery/329a-mediabrowser-providers-videos.md
│   ├── Years                       → .discovery/329b-mediabrowser-providers-years.md
│   └── Users                       → .discovery/329-mediabrowser-providers-users.md
├── MediaBrowser.Server.Mono/        → .discovery/210-mediabrowser-server-mono.md
├── MediaBrowser.ServerApplication/  → .discovery/220-mediabrowser-serverapplication.md
│   └── Native                      → .discovery/221-mediabrowser-serverapplication-internals.md
├── MediaBrowser.Tests/             → .discovery/230-mediabrowser-tests.md
├── MediaBrowser.WebDashboard/       → .discovery/260-mediabrowser-webdashboard.md
│   ├── API                         → .discovery/261-mediabrowser-webdashboard-api.md
│   ├── UI                          → .discovery/262-mediabrowser-webdashboard-ui.md
│   ├── Scripts                     → .discovery/263-mediabrowser-webdashboard-scripts.md
│   ├── Components                  → .discovery/264-mediabrowser-webdashboard-components.md
│   ├── Strings                     → .discovery/265-mediabrowser-webdashboard-strings.md
│   └── Bower                       → .discovery/266-mediabrowser-webdashboard-bower.md
├── MediaBrowser.XbmcMetadata/      → .discovery/240-mediabrowser-xbmcmetadata.md
├── Mono.Nat/                       → .discovery/250-mono-nat.md, 251-mono-nat-internals.md
├── RSSDP/                          → .discovery/300-rssdp.md, 301-rssdp-internals.md
├── SocketHttpListener/             → .discovery/350-sockethttplistener.md
│   ├── HTTP                        → .discovery/351-sockethttplistener-http.md
│   ├── Net                         → .discovery/352-sockethttplistener-net.md
│   └── Web                         → .discovery/353-sockethttplistener-web.md
├── ThirdParty/                    → .discovery/370-thirdparty.md
├── emby-go/                        → .discovery/360-emby-go.md
├── MediaBrowser.sln                → .discovery/000-root.md
├── SharedVersion.cs                → .discovery/910-sharedversion.md
├── README.md                       → .discovery/920-readme.md
├── CONTRIBUTORS.md                 → .discovery/930-contributors.md
└── LICENSE.md                      → .discovery/940-license.md
```

---

## Document Map

| File | Component | Type | Description |
|------|-----------|------|-------------|
| [000-root.md](./000-root.md) | Project root | Entry point | Master project overview |
| [100-bdinfo.md](./100-bdinfo.md) | BDInfo | Module | Blu-ray disc analysis |
| [101-bdinfo-internals.md](./101-bdinfo-internals.md) | BDInfo | Expanded | Codec files, stream parsing |
| [100-01-bdrom.md](./100-01-bdrom.md) | BDROM.cs | File | BD-ROM disc reader |
| [100-02-tsplaylistfile.md](./100-02-tsplaylistfile.md) | TSPlaylistFile.cs | File | Playlist parser |
| [100-03-tsstreamfile.md](./100-03-tsstreamfile.md) | TSStreamFile.cs | File | Stream file parser |
| [100-04-tsstreamclipfile.md](./100-04-tsstreamclipfile.md) | TSStreamClipFile.cs | File | Stream clip parser |
| [100-05-languagecodes.md](./100-05-languagecodes.md) | LanguageCodes.cs | File | Language codes |
| [100-06-tsstream.md](./100-06-tsstream.md) | TSStream.cs | File | Transport stream |
| [100-07-tsstreamclip.md](./100-07-tsstreamclip.md) | TSStreamClip.cs | File | Stream clip |
| [100-08-tsinterleavedfile.md](./100-08-tsinterleavedfile.md) | TSInterleavedFile.cs | File | Interleaved file |
| [100-09-bdinfosettings.md](./100-09-bdinfosettings.md) | BDInfoSettings.cs | File | Settings |
| [100-10-tsstreambuffer.md](./100-10-tsstreambuffer.md) | TSStreamBuffer.cs | File | Stream buffer |
| [100-11-codecs.md](./100-11-codecs.md) | Codecs.cs | File | Codec definitions |
| [110-dvdlib.md](./110-dvdlib.md) | DvdLib | Module | DVD structure parsing |
| [111-dvdliv-internals.md](./111-dvdliv-internals.md) | DvdLib | Expanded | IFO parsers, structure |
| [120-emby-drawing.md](./120-emby-drawing.md) | Emby.Drawing | Module | Image processing |
| [121-emby-drawing-imagemagick.md](./121-emby-drawing-imagemagick.md) | Emby.Drawing.ImageMagick | Module | ImageMagick backend |
| [122-emby-drawing-net.md](./122-emby-drawing-net.md) | Emby.Drawing.Net | Module | .NET drawing backend |
| [123-emby-drawing-skia.md](./123-emby-drawing-skia.md) | Emby.Drawing.Skia | Module | SkiaSharp backend |
| [131-emby-notifications-internals.md](./131-emby-notifications-internals.md) | Emby.Notifications | Expanded | API, managers |
| [140-emby-notifications.md](./140-emby-notifications.md) | Emby.Notifications | Module | Notification system |
| [150-emby-photos.md](./150-emby-photos.md) | Emby.Photos | Module | Photo management |
| [160-emby-server-impl.md](./160-emby-server-impl.md) | Emby.Server.Implementations | Module | Core server implementation |
| [161-emby-server-impl-core.md](./161-emby-server-impl-core.md) | Core Infrastructure | Sub-module | AppBase, Config, Crypto, etc. |
| [162-emby-server-impl-library.md](./162-emby-server-impl-library.md) | Library Management | Sub-module | Library, Collections, Playlists |
| [163-emby-server-impl-media.md](./163-emby-server-impl-media.md) | Media & Channels | Sub-module | Channels, LiveTV |
| [164-emby-server-impl-http.md](./164-emby-server-impl-http.md) | HTTP Server | Sub-module | HttpServer, Services, EntryPoints |
| [165-emby-server-impl-tasks.md](./165-emby-server-impl-tasks.md) | Scheduled Tasks | Sub-module | Background task scheduler |
| [166-emby-server-impl-io.md](./166-emby-server-impl-io.md) | I/O Utilities | Sub-module | File system I/O wrappers |
| [167-emby-server-impl-encoding.md](./167-emby-server-impl-encoding.md) | Text Encoding | Sub-module | Localization, UniversalDetector |
| [168-emby-server-impl-security.md](./168-emby-server-impl-security.md) | Security | Sub-module | Security, Session, Devices |
| [169-emby-server-impl-sharpcifs.md](./169-emby-server-impl-sharpcifs.md) | SharpCifs | Sub-module | Embedded SMB/CIFS client |
| [170-emby-server-impl-livetv.md](./170-emby-server-impl-livetv.md) | LiveTV | Sub-module | Live TV functionality |
| [171-emby-server-impl-data.md](./171-emby-server-impl-data.md) | Data | Sub-module | Data persistence |
| [172-emby-server-impl-resolvers.md](./172-emby-server-impl-resolvers.md) | Resolvers | Sub-module | Path resolvers |
| [173-emby-server-impl-dto.md](./173-emby-server-impl-dto.md) | DTO | Sub-module | Data transfer objects |
| [174-emby-server-impl-images.md](./174-emby-server-impl-images.md) | Images | Sub-module | Image handling |
| [175-emby-server-impl-httpclient.md](./175-emby-server-impl-httpclient.md) | HTTP Client | Sub-module | HTTP utilities |
| [176-emby-server-impl-ffmpeg.md](./176-emby-server-impl-ffmpeg.md) | FFMpeg | Sub-module | FFmpeg management |
| [177-emby-server-impl-net.md](./177-emby-server-impl-net.md) | Net | Sub-module | Networking |
| [178-emby-server-impl-networking.md](./178-emby-server-impl-networking.md) | Networking | Sub-module | Network utilities |
| [179-emby-server-impl-userviews.md](./179-emby-server-impl-userviews.md) | UserViews | Sub-module | User view management |
| [180-activity.md](./180-activity.md) | Activity | Sub-module | Activity logging |
| [181-archiving.md](./181-archiving.md) | Archiving | Sub-module | ZIP archive handling |
| [182-branding.md](./182-branding.md) | Branding | Sub-module | Server branding |
| [183-browser.md](./183-browser.md) | Browser | Sub-module | Browser launcher |
| [184-devices.md](./184-devices.md) | Devices | Sub-module | Device management |
| [185-dto.md](./185-dto.md) | Dto | Sub-module | DTO service |
| [186-entrypoints.md](./186-entrypoints.md) | EntryPoints | Sub-module | Server entry points |
| [187-ffmpeg.md](./187-ffmpeg.md) | FFMpeg | Sub-module | FFmpeg management |
| [188-library-full.md](./188-library-full.md) | Library | Sub-module | Library management (full) |
| [189-livetv-full.md](./189-livetv-full.md) | LiveTv | Sub-module | LiveTV (full) |
| [190-localization.md](./190-localization.md) | Localization | Sub-module | i18n support |
| [191-tv.md](./191-tv.md) | TV | Sub-module | TV series manager |
| [192-udp.md](./192-udp.md) | Udp | Sub-module | UDP server |
| [193-userviews.md](./193-userviews.md) | UserViews | Sub-module | User views |
| [210-mediabrowser-localmetadata.md](./210-mediabrowser-localmetadata.md) | LocalMetadata | Module | Local metadata |
| [211-mediabrowser-localmetadata-internals.md](./211-mediabrowser-localmetadata-internals.md) | LocalMetadata | Expanded | Internal files |
| [210-mediabrowser-server-mono.md](./210-mediabrowser-server-mono.md) | Server.Mono | Module | Mono-specific server |
| [220-mediabrowser-serverapplication.md](./220-mediabrowser-serverapplication.md) | ServerApplication | Module | Server application |
| [221-mediabrowser-serverapplication-internals.md](./221-mediabrowser-serverapplication-internals.md) | ServerApplication | Expanded | Native, networking |
| [230-mediabrowser-tests.md](./230-mediabrowser-tests.md) | Tests | Module | Test suite |
| [240-mediabrowser-xbmcmetadata.md](./240-mediabrowser-xbmcmetadata.md) | XbmcMetadata | Module | Kodi metadata |
| [241-mediabrowser-xbmcmetadata-internals.md](./241-mediabrowser-xbmcmetadata-internals.md) | XbmcMetadata | Expanded | Internal files |
| [250-mono-nat.md](./250-mono-nat.md) | Mono.Nat | Module | NAT traversal |
| [251-mono-nat-internals.md](./251-mono-nat-internals.md) | Mono.Nat | Expanded | Internal files |
| [260-mediabrowser-webdashboard.md](./260-mediabrowser-webdashboard.md) | WebDashboard | Module | Web UI |
| [261-mediabrowser-webdashboard-api.md](./261-mediabrowser-webdashboard-api.md) | WebDashboard API | Sub-module | API endpoints |
| [262-mediabrowser-webdashboard-ui.md](./262-mediabrowser-webdashboard-ui.md) | WebDashboard UI | Sub-module | HTML structure |
| [263-mediabrowser-webdashboard-scripts.md](./263-mediabrowser-webdashboard-scripts.md) | WebDashboard Scripts | Sub-module | JavaScript |
| [264-mediabrowser-webdashboard-components.md](./264-mediabrowser-webdashboard-components.md) | WebDashboard Components | Sub-module | UI components |
| [265-mediabrowser-webdashboard-strings.md](./265-mediabrowser-webdashboard-strings.md) | WebDashboard Strings | Sub-module | Localization |
| [266-mediabrowser-webdashboard-bower.md](./266-mediabrowser-webdashboard-bower.md) | WebDashboard Bower | Sub-module | Dependencies |
| [267-webdashboard-api.md](./267-webdashboard-api.md) | WebDashboard API | Sub-module | Backend API |
| [268-webdashboard-ui.md](./268-webdashboard-ui.md) | WebDashboard UI | Sub-module | Frontend |
| [269-webdashboard-scripts.md](./269-webdashboard-scripts.md) | WebDashboard Scripts | Sub-module | Scripts |
| [300-rssdp.md](./300-rssdp.md) | RSSDP | Module | SSDP implementation |
| [301-rssdp-internals.md](./301-rssdp-internals.md) | RSSDP | Expanded | Internal files |
| [320-mediabrowser-providers.md](./320-mediabrowser-providers.md) | Providers | Module | Metadata providers |
| [321-mediabrowser-providers-movies.md](./321-mediabrowser-providers-movies.md) | Movies Provider | Sub-module | Movie metadata |
| [322-mediabrowser-providers-tv.md](./322-mediabrowser-providers-tv.md) | TV Provider | Sub-module | TV metadata |
| [323-mediabrowser-providers-music.md](./323-mediabrowser-providers-music.md) | Music Provider | Sub-module | Music metadata |
| [324-mediabrowser-providers-images.md](./324-mediabrowser-providers-images.md) | Images Provider | Sub-module | Image providers |
| [325-mediabrowser-providers-people.md](./325-mediabrowser-providers-people.md) | People Provider | Sub-module | Person metadata |
| [326-mediabrowser-providers-books.md](./326-mediabrowser-providers-books.md) | Books Provider | Sub-module | Book metadata |
| [327-mediabrowser-providers-tv.md](./327-mediabrowser-providers-tv.md) | TV Provider | Sub-module | TV shows |
| [328-mediabrowser-providers-subtitles.md](./328-mediabrowser-providers-subtitles.md) | Subtitles Provider | Sub-module | Subtitle fetching |
| [329-mediabrowser-providers-users.md](./329-mediabrowser-providers-users.md) | Users Provider | Sub-module | User data |
| [329a-mediabrowser-providers-videos.md](./329a-mediabrowser-providers-videos.md) | Videos Provider | Sub-module | Video metadata |
| [329b-mediabrowser-providers-years.md](./329b-mediabrowser-providers-years.md) | Years Provider | Sub-module | Year-based grouping |
| [330-emby-dlna.md](./330-emby-dlna.md) | Emby.Dlna | Module | DLNA support |
| [331-emby-dlna-profiles.md](./331-emby-dlna-profiles.md) | DLNA Profiles | Sub-module | Device profiles |
| [332-emby-dlna-server.md](./332-emby-dlna-server.md) | DLNA Server | Sub-module | Server implementation |
| [333-emby-dlna-playto.md](./333-emby-dlna-playto.md) | DLNA PlayTo | Sub-module | PlayTo functionality |
| [340-mediabrowser-api.md](./340-mediabrowser-api.md) | MediaBrowser.Api | Module | REST API |
| [341-mediabrowser-api-controllers.md](./341-mediabrowser-api-controllers.md) | API Controllers | Sub-module | HTTP endpoints |
| [342-mediabrowser-api-models.md](./342-mediabrowser-api-models.md) | API Models | Sub-module | Request/response |
| [343-mediabrowser-api-services.md](./343-mediabrowser-api-services.md) | API Services | Sub-module | Business logic |
| [350-sockethttplistener.md](./350-sockethttplistener.md) | SocketHttpListener | Module | HTTP server |
| [351-sockethttplistener-http.md](./351-sockethttplistener-http.md) | HTTP | Sub-module | HTTP handling |
| [352-sockethttplistener-net.md](./352-sockethttplistener-net.md) | Net | Sub-module | .NET integration |
| [353-sockethttplistener-web.md](./353-sockethttplistener-web.md) | Web | Sub-module | WebSocket |
| [360-emby-go.md](./360-emby-go.md) | emby-go | Module | Go client library |
| [370-thirdparty.md](./370-thirdparty.md) | ThirdParty | Module | Third-party libs |
| [400-packages.md](./400-packages.md) | NuGet Packages | Module | Dependencies |
| [900-solution.md](./900-solution.md) | Solution | File | VS solution |
| [910-sharedversion.md](./910-sharedversion.md) | SharedVersion | File | Version info |
| [920-readme.md](./920-readme.md) | README | File | Documentation |
| [930-contributors.md](./930-contributors.md) | CONTRIBUTORS | File | Contributors |
| [940-license.md](./940-license.md) | LICENSE | File | License |
| [950-project-artifacts.md](./950-project-artifacts.md) | Artifacts | File | Git, VS files |
