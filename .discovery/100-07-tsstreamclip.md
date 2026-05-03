# Component: TSStreamClip.cs

**Path:** `BDInfo/TSStreamClip.cs`
**Type:** File | Class
**Language:** C#
**Maps to:** `.discovery/100-07-tsstreamclip.md`
**Parent:** `.discovery/100-bdinfo.md`
**Dependencies:**
- `TSStreamFile` → `.discovery/100-03-tsstreamfile.md`
- `TSStreamClipFile` → `.discovery/100-04-tsstreamclipfile.md`
- `TSStream` → `.discovery/100-06-tsstream.md`
- `BDInfoSettings` → `.discovery/100-09-bdinfosettings.md`

## Structure

```
TSStreamClip.cs
├── [namespace] BDInfo
│   └── [class] public class TSStreamClip
│       ├── [field] public int AngleIndex = 0
│       ├── [field] public string Name
│       ├── [field] public double TimeIn
│       ├── [field] public double TimeOut
│       ├── [field] public double RelativeTimeIn
│       ├── [field] public double RelativeTimeOut
│       ├── [field] public double Length
│       ├── [field] public ulong FileSize = 0
│       ├── [field] public ulong InterleavedFileSize = 0
│       ├── [field] public ulong PayloadBytes = 0
│       ├── [field] public ulong PacketCount = 0
│       ├── [field] public double PacketSeconds = 0
│       ├── [field] public List<double> Chapters
│       ├── [field] public TSStreamFile StreamFile = null
│       ├── [field] public TSStreamClipFile StreamClipFile = null
│       ├── [constructor] TSStreamClip(TSStreamFile streamFile, TSStreamClipFile streamClipFile)
│       │   ├── Sets Name from streamFile.Name
│       │   ├── Sets FileSize from streamFile.FileInfo.Length
│       │   ├── Sets InterleavedFileSize if interleaved file exists
│       │   └── Stores references to StreamFile and StreamClipFile
│       ├── [property] public string DisplayName
│       │   └── Returns interleaved file name if BDInfoSettings.EnableSSIF, else Name
│       ├── [property] public ulong PacketSize
│       │   └── Returns PacketCount * 192 (BDAV packet size)
│       ├── [property] public ulong PacketBitRate
│       │   └── Calculates bitrate from PacketSize / PacketSeconds
│       └── [method] public bool IsCompatible(TSStreamClip clip)
│           └── Compares stream types by PID across clips
```

## Description

`TSStreamClip` represents a single clip entry within a Blu-ray playlist (MPLS). It holds timing information (TimeIn/TimeOut), file references (TSStreamFile, TSStreamClipFile), and calculated properties like packet size and bitrate. The `IsCompatible` method checks if two clips have matching stream types by PID, used for seamless branching detection.
