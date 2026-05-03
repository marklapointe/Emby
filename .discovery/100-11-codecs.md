# Component: BDInfo Codec Scanners

**Path:** `BDInfo/TSCodec*.cs`
**Type:** File Group | Static Codec Parsers
**Language:** C#
**Maps to:** `.discovery/100-11-codecs.md`
**Parent:** `.discovery/100-bdinfo.md`
**Dependencies:**
- `TSStream` → `.discovery/100-06-tsstream.md` — Stream instances populated by scanners
- `TSStreamBuffer` → `.discovery/100-10-tsstreambuffer.md` — Byte buffer for stream data

## Structure

```
BDInfo Codec Scanners (9 files)
├── TSCodecAC3.cs
│   └── [class] public static class TSCodecAC3
│       ├── [field] private static byte[] eac3_blocks = { 1, 2, 3, 6 }
│       └── [method] public static void Scan(TSAudioStream stream, TSStreamBuffer buffer)
│           ├── Parses AC-3/E-AC-3 syncframe header
│           ├── Extracts bitrate, channel count, sample rate, frame size
│           ├── Handles E-AC-3 block structure
│           └── Populates TSAudioStream properties
├── TSCodecAVC.cs
│   └── [class] public static class TSCodecAVC
│       └── [method] public static void Scan(TSVideoStream stream, TSStreamBuffer buffer)
│           ├── Parses H.264/AVC NAL units (SPS, PPS)
│           ├── Extracts resolution, profile, level, frame rate
│           ├── Handles SEI messages for timing
│           └── Populates TSVideoStream properties
├── TSCodecDTS.cs
│   └── [class] public static class TSCodecDTS
│       ├── [field] private static int[] dca_sample_rates
│       ├── [field] private static int[] dca_bit_rates
│       ├── [field] private static int[] dca_channels
│       ├── [field] private static int[] dca_bits_per_sample
│       └── [method] public static void Scan(TSAudioStream stream, TSStreamBuffer buffer)
│           ├── Parses DTS syncword and frame header
│           ├── Extracts sample rate, bitrate, channels, bit depth
│           └── Populates TSAudioStream properties
├── TSCodecDTSHD.cs
│   └── [class] public static class TSCodecDTSHD
│       ├── [field] private static int[] SampleRates
│       └── [method] public static void Scan(TSAudioStream stream, TSStreamBuffer buffer)
│           ├── Parses DTS-HD extension headers
│           ├── Extracts master audio properties (lossless)
│           ├── Handles core + extension structure
│           └── Populates TSAudioStream properties
├── TSCodecLPCM.cs
│   └── [class] public static class TSCodecLPCM
│       └── [method] public static void Scan(TSAudioStream stream, TSStreamBuffer buffer)
│           ├── Parses Blu-ray LPCM audio header
│           ├── Extracts sample rate, bit depth, channel count
│           └── Populates TSAudioStream properties
├── TSCodecMPEG2.cs
│   └── [class] public static class TSCodecMPEG2
│       └── [method] public static void Scan(TSVideoStream stream, TSStreamBuffer buffer)
│           ├── Parses MPEG-2 video sequence header
│           ├── Extracts resolution, aspect ratio, frame rate, bitrate
│           ├── Handles progressive/interlaced detection
│           └── Populates TSVideoStream properties
├── TSCodecMVC.cs
│   └── [class] public static class TSCodecMVC
│       └── [method] public static void Scan(TSVideoStream stream, TSStreamBuffer buffer)
│           ├── Parses MVC (Multiview Video Coding) NAL units
│           ├── Extracts 3D video properties
│           └── Populates TSVideoStream properties
├── TSCodecTrueHD.cs
│   └── [class] public static class TSCodecTrueHD
│       └── [method] public static void Scan(TSAudioStream stream, TSStreamBuffer buffer)
│           ├── Parses Dolby TrueHD syncframe
│           ├── Extracts sample rate, channels, bit depth (lossless)
│           ├── Handles AC-3 interleaved core
│           └── Populates TSAudioStream properties
└── TSCodecVC1.cs
    └── [class] public static class TSCodecVC1
        └── [method] public static void Scan(TSVideoStream stream, TSStreamBuffer buffer)
            ├── Parses VC-1 sequence header
            ├── Extracts resolution, profile, level, frame rate
            └── Populates TSVideoStream properties
```

## Description

The BDInfo codec scanner modules are static utility classes that parse raw transport stream data to extract codec-specific metadata. Each scanner reads from a `TSStreamBuffer`, parses the codec headers (syncframes, sequence headers, NAL units), and populates the corresponding `TSStream` subclass properties (resolution, bitrate, channels, sample rate, etc.).

## Data Flow

```
TSStreamFile.Scan() ──► TSStreamBuffer ──► Codec Scanner
    ├── TSCodecAC3 ──► TSAudioStream (Dolby Digital)
    ├── TSCodecAVC ──► TSVideoStream (H.264/AVC)
    ├── TSCodecDTS ──► TSAudioStream (DTS)
    ├── TSCodecDTSHD ──► TSAudioStream (DTS-HD MA)
    ├── TSCodecLPCM ──► TSAudioStream (LPCM)
    ├── TSCodecMPEG2 ──► TSVideoStream (MPEG-2)
    ├── TSCodecMVC ──► TSVideoStream (3D MVC)
    ├── TSCodecTrueHD ──► TSAudioStream (Dolby TrueHD)
    └── TSCodecVC1 ──► TSVideoStream (VC-1)
```

## Side Effects

- Reads from TSStreamBuffer (does not modify)
- Modifies TSStream subclass properties in-place
- No file I/O (buffer already loaded)
