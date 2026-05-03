# Component: Emby.Server.Implementations — Media & Channels

**Path:** \`Emby.Server.Implementations/\`
**Type:** Directory | Module Group
**Language:** C#
**Maps to:** \`.discovery/163-emby-server-impl-media.md\`

## Description

Media handling, channel plugins, and Live TV support components.

## Files

### Channels/

- `ChannelDynamicMediaSourceProvider.cs` — Emby.Server.Implementations/Channels/ChannelDynamicMediaSourceProvider.cs
- `ChannelImageProvider.cs` — Emby.Server.Implementations/Channels/ChannelImageProvider.cs
- `ChannelManager.cs` — Emby.Server.Implementations/Channels/ChannelManager.cs
- `ChannelPostScanTask.cs` — Emby.Server.Implementations/Channels/ChannelPostScanTask.cs
- `RefreshChannelsScheduledTask.cs` — Emby.Server.Implementations/Channels/RefreshChannelsScheduledTask.cs

### LiveTv/

- `DirectRecorder.cs` — Emby.Server.Implementations/LiveTv/EmbyTV/DirectRecorder.cs
- `EmbyTV.cs` — Emby.Server.Implementations/LiveTv/EmbyTV/EmbyTV.cs
- `EncodedRecorder.cs` — Emby.Server.Implementations/LiveTv/EmbyTV/EncodedRecorder.cs
- `EntryPoint.cs` — Emby.Server.Implementations/LiveTv/EmbyTV/EntryPoint.cs
- `IRecorder.cs` — Emby.Server.Implementations/LiveTv/EmbyTV/IRecorder.cs
- `ItemDataProvider.cs` — Emby.Server.Implementations/LiveTv/EmbyTV/ItemDataProvider.cs
- `RecordingHelper.cs` — Emby.Server.Implementations/LiveTv/EmbyTV/RecordingHelper.cs
- `SeriesTimerManager.cs` — Emby.Server.Implementations/LiveTv/EmbyTV/SeriesTimerManager.cs
- `TimerManager.cs` — Emby.Server.Implementations/LiveTv/EmbyTV/TimerManager.cs
- `SchedulesDirect.cs` — Emby.Server.Implementations/LiveTv/Listings/SchedulesDirect.cs
- `LiveTvConfigurationFactory.cs` — Emby.Server.Implementations/LiveTv/LiveTvConfigurationFactory.cs
- `LiveTvDtoService.cs` — Emby.Server.Implementations/LiveTv/LiveTvDtoService.cs
- `LiveTvManager.cs` — Emby.Server.Implementations/LiveTv/LiveTvManager.cs
- `LiveTvMediaSourceProvider.cs` — Emby.Server.Implementations/LiveTv/LiveTvMediaSourceProvider.cs
- `RefreshChannelsScheduledTask.cs` — Emby.Server.Implementations/LiveTv/RefreshChannelsScheduledTask.cs
- `BaseTunerHost.cs` — Emby.Server.Implementations/LiveTv/TunerHosts/BaseTunerHost.cs
- `HdHomerunHost.cs` — Emby.Server.Implementations/LiveTv/TunerHosts/HdHomerun/HdHomerunHost.cs
- `HdHomerunManager.cs` — Emby.Server.Implementations/LiveTv/TunerHosts/HdHomerun/HdHomerunManager.cs
- `HdHomerunUdpStream.cs` — Emby.Server.Implementations/LiveTv/TunerHosts/HdHomerun/HdHomerunUdpStream.cs
- `LiveStream.cs` — Emby.Server.Implementations/LiveTv/TunerHosts/LiveStream.cs
- `M3uParser.cs` — Emby.Server.Implementations/LiveTv/TunerHosts/M3uParser.cs
- `M3UTunerHost.cs` — Emby.Server.Implementations/LiveTv/TunerHosts/M3UTunerHost.cs
- `SharedHttpStream.cs` — Emby.Server.Implementations/LiveTv/TunerHosts/SharedHttpStream.cs

## Decomposition

### ChannelManager.cs (Channel Management)

#### Imports
```csharp
using MediaBrowser.Controller.Channels;
using MediaBrowser.Controller.Entities;
using MediaBrowser.Model.Channels;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
```

#### Classes
`ChannelManager` (public class : IChannelManager)

#### Key Properties
| Property | Type | Description |
|----------|------|-------------|
| `Channels` | `IEnumerable<Channel>` | All channels |

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `GetChannels(ChannelQuery)` | `Task<ChannelBag>` | Get channels |
| `GetChannelFeatures(string)` | `Task<ChannelFeatures>` | Channel capabilities |
| `GetChannelItems(string, ChannelItemQuery)` | `Task<ChannelItemResult>` | Channel content |

### LiveTvManager.cs (Live TV Core)

#### Classes
`LiveTvManager` (public class : ILiveTvManager)

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `GetServices()` | `Task<IEnumerable<LiveTvServiceInfo>>` | TV services |
| `GetChannels(LiveTvChannelQuery)` | `Task< IEnumerable<BaseItem>` | Get channels |
| `GetPrograms(LiveTvProgramQuery)` | `Task< IEnumerable<BaseItem>` | Get programs |
| `GetRecordingGroups(RecordingGroupQuery)` | `Task<IEnumerable<BaseItem>>` | Recording groups |

### EmbyTV.cs (Emby TV Backend)

#### Classes
`EmbyTV` (public class : IEmbyTV)

#### Key Properties
| Property | Type | Description |
|----------|------|-------------|
| `Recorder` | `IRecorder` | Current recorder |
| `IsRecording` | `bool` | Recording in progress |

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `StartRecording(SeriesTimerInfo)` | `Task<string>` | Start recording |
| `StopRecording(string)` | `Task` | Stop recording |
| `GetLiveStream(MediaSourceInfo)` | `Task<MediaSourceInfo>` | Get live stream |

### SchedulesDirect.cs (TV Guide Provider)

#### Classes
`SchedulesDirect` (public class)

#### Key Properties
| Property | Type | Description |
|----------|------|-------------|
| `Token` | `string` | Auth token |

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `GetScheduleForChannel(string, DateTime)` | `Task<Schedule>` | Get schedule |
| `GetPrograms(IEnumerable<string>)` | `Task<IEnumerable<Program>>` | Get program details |

### HdHomerunManager.cs (HDHomerun Tuner Control)

#### Classes
`HdHomerunManager` (public class)

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `DiscoverDevices()` | `Task<IEnumerable<HdHomerunDevice>>` | Find tuners |
| `GetDevice(string)` | `HdHomerunDevice` | Get device |
| `Tune(string, uint)` | `Task` | Tune to channel |

