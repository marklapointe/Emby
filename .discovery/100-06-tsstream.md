# Component: TSStream.cs

**Path:** `BDInfo/TSStream.cs`
**Type:** File | Multiple Classes and Enums
**Language:** C#
**Maps to:** `.discovery/100-06-tsstream.md`
**Parent:** `.discovery/100-bdinfo.md`

## Structure

```
TSStream.cs
├── [namespace] BDInfo
│   ├── [enum] public enum TSStreamType : byte
│   │   ├── MPEG1_VIDEO = 0x01, MPEG2_VIDEO = 0x02
│   │   ├── AVC_VIDEO = 0x1b, MVC_VIDEO = 0x20, VC1_VIDEO = 0xea
│   │   ├── MPEG1_AUDIO = 0x03, MPEG2_AUDIO = 0x04
│   │   ├── LPCM_AUDIO = 0x80
│   │   ├── AC3_AUDIO = 0x81, AC3_PLUS_AUDIO = 0x84, AC3_PLUS_SECONDARY_AUDIO = 0xA1
│   │   ├── AC3_TRUE_HD_AUDIO = 0x83
│   │   ├── DTS_AUDIO = 0x82, DTS_HD_AUDIO = 0x85, DTS_HD_SECONDARY_AUDIO = 0xA2
│   │   ├── DTS_HD_MASTER_AUDIO = 0x86
│   │   ├── PRESENTATION_GRAPHICS = 0x90, INTERACTIVE_GRAPHICS = 0x91
│   │   └── SUBTITLE = 0x92
│   ├── [enum] public enum TSVideoFormat : byte
│   │   ├── VIDEOFORMAT_480i, VIDEOFORMAT_576i, VIDEOFORMAT_480p
│   │   ├── VIDEOFORMAT_1080i, VIDEOFORMAT_720p, VIDEOFORMAT_1080p
│   │   └── VIDEOFORMAT_576p
│   ├── [enum] public enum TSFrameRate : byte
│   │   ├── FRAMERATE_23_976, FRAMERATE_24, FRAMERATE_25
│   │   ├── FRAMERATE_29_97, FRAMERATE_50, FRAMERATE_59_94
│   │   └── Unknown = 0
│   ├── [enum] public enum TSChannelLayout : byte
│   │   ├── CHANNELLAYOUT_MONO, CHANNELLAYOUT_STEREO
│   │   ├── CHANNELLAYOUT_MULTI, CHANNELLAYOUT_COMBO
│   │   └── Unknown = 0
│   ├── [enum] public enum TSSampleRate : byte
│   │   ├── SAMPLERATE_48_KHZ, SAMPLERATE_96_KHZ, SAMPLERATE_192_KHZ
│   │   ├── SAMPLERATE_44_1_KHZ, SAMPLERATE_88_2_KHZ, SAMPLERATE_176_4_KHZ
│   │   └── Unknown = 0
│   ├── [enum] public enum TSAspectRatio
│   │   ├── ASPECT_4_3, ASPECT_16_9, ASPECT_2_21
│   │   └── Unknown = 0
│   ├── [class] public class TSDescriptor
│   │   ├── [field] public byte DescriptorTag
│   │   ├── [field] public byte DescriptorData
│   │   └── [constructor] TSDescriptor(byte tag, byte data)
│   ├── [class] public class TSStream
│   │   ├── [field] public ushort PID
│   │   ├── [field] public TSStreamType StreamType
│   │   ├── [field] public bool IsVBR
│   │   ├── [field] public bool IsInitialized
│   │   ├── [field] public List<TSDescriptor> Descriptors
│   │   ├── [field] public ulong BitRate
│   │   ├── [field] public ulong ActiveBitRate
│   │   ├── [field] public bool IsHidden
│   │   ├── [field] public bool IsInitialized
│   │   ├── [field] public bool HasMultipleAngles
│   │   ├── [field] public bool IsPastel
│   │   ├── [field] public string LanguageCode
│   │   ├── [field] public string Description
│   │   ├── [field] public string CodecShortName
│   │   ├── [field] public string CodecLongName
│   │   ├── [field] public string FrameRateDescription
│   │   ├── [field] public string VideoFormatDescription
│   │   ├── [field] public string AspectRatioDescription
│   │   ├── [field] public string ChannelDescription
│   │   ├── [field] public string SampleRateDescription
│   │   ├── [field] public string BitDepthDescription
│   │   ├── [field] public string DelayDescription
│   │   ├── [field] public string VideoResolution
│   │   ├── [field] public string AudioModeDescription
│   │   ├── [constructor] TSStream()
│   │   └── [method] public override string ToString()
│   │       └── Returns Description
│   ├── [class] public class TSVideoStream : TSStream
│   │   ├── [field] public TSVideoFormat VideoFormat
│   │   ├── [field] public TSFrameRate FrameRate
│   │   ├── [field] public TSAspectRatio AspectRatio
│   │   ├── [field] public int Width
│   │   ├── [field] public int Height
│   │   ├── [field] public bool IsInterlaced
│   │   ├── [field] public bool Is3D
│   │   ├── [constructor] TSVideoStream()
│   │   └── [method] public override string ToString()
│   │       └── Returns formatted video description (codec, resolution, framerate)
│   ├── [enum] public enum TSAudioMode
│   │   ├── Unknown, DualMono, Stereo, MultiChannel
│   │   └── Unknown = 0
│   ├── [class] public class TSAudioStream : TSStream
│   │   ├── [field] public TSChannelLayout ChannelLayout
│   │   ├── [field] public TSSampleRate SampleRate
│   │   ├── [field] public int BitDepth
│   │   ├── [field] public int AudioMode
│   │   ├── [field] public int CoreAudioMode
│   │   ├── [field] public int CoreSampleRate
│   │   ├── [field] public int CoreBitDepth
│   │   ├── [field] public int CoreChannelCount
│   │   ├── [field] public int CoreFlags
│   │   ├── [field] public bool HasDialogNormalization
│   │   ├── [field] public int DialogNormalization
│   │   ├── [field] public bool IsLossless
│   │   ├── [field] public double Delay
│   │   ├── [field] public double CoreDelay
│   │   ├── [field] public bool IsVBR
│   │   ├── [constructor] TSAudioStream()
│   │   └── [method] public override string ToString()
│   │       └── Returns formatted audio description (codec, channels, sample rate)
│   ├── [class] public class TSGraphicsStream : TSStream
│   │   ├── [field] public int Width
│   │   ├── [field] public int Height
│   │   ├── [constructor] TSGraphicsStream()
│   │   └── [method] public override string ToString()
│   │       └── Returns formatted graphics description
│   └── [class] public class TSTextStream : TSStream
│       ├── [field] public int Width
│       ├── [field] public int Height
│       ├── [constructor] TSTextStream()
│       └── [method] public override string ToString()
│           └── Returns formatted text/subtitle description
```

## Description

`TSStream.cs` defines the core type hierarchy for Blu-ray transport stream elements. It includes enums for stream types (video/audio/graphics codecs), video formats, frame rates, channel layouts, sample rates, and aspect ratios. The base `TSStream` class holds common properties (PID, bitrate, language, descriptors), with subclasses `TSVideoStream`, `TSAudioStream`, `TSGraphicsStream`, and `TSTextStream` adding type-specific fields.

## Data Flow

```
MPLS/CLPI parsing ──► TSStream subclass instances created
    ├── TSVideoStream ──► Resolution, format, framerate, aspect ratio
    ├── TSAudioStream ──► Channels, sample rate, bit depth, delay
    ├── TSGraphicsStream ──► Width, height
    └── TSTextStream ──► Width, height

TSStream.ToString() ──► Human-readable description strings
    └── Used by UI and logging throughout BDInfo
```
