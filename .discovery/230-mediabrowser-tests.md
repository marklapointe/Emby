# Component: MediaBrowser.Tests

**Path:** `MediaBrowser.Tests/`
**Type:** Directory | Test Project
**Language:** C#
**Maps to:** `.discovery/230-mediabrowser-tests.md`

## Description

Unit tests for Emby Server components including M3U playlist parsing, subtitle conversion, and consistency checking.

## Files

### Root Files (4 files)

- `M3uParserTest.cs` — MediaBrowser.Tests/M3uParserTest.cs
- `MediaBrowser.Tests.csproj` — MediaBrowser.Tests/MediaBrowser.Tests.csproj
- `Properties/AssemblyInfo.cs` — MediaBrowser.Tests/Properties/AssemblyInfo.cs
- `app.config` — MediaBrowser.Tests/app.config
- `packages.config` — MediaBrowser.Tests/packages.config

### ConsistencyTests/ (9 files)

- `StringUsageReporter.cs` — MediaBrowser.Tests/ConsistencyTests/StringUsageReporter.cs
- `Resources/SampleTransformed.htm` — MediaBrowser.Tests/ConsistencyTests/Resources/SampleTransformed.htm
- `Resources/StringCheck.xslt` — MediaBrowser.Tests/ConsistencyTests/Resources/StringCheck.xslt
- `Resources/StringCheckSample.xml` — MediaBrowser.Tests/ConsistencyTests/Resources/StringCheckSample.xml
- `TextIndexing/IndexBuilder.cs` — MediaBrowser.Tests/ConsistencyTests/TextIndexing/IndexBuilder.cs
- `TextIndexing/WordIndex.cs` — MediaBrowser.Tests/ConsistencyTests/TextIndexing/WordIndex.cs
- `TextIndexing/WordOccurrence.cs` — MediaBrowser.Tests/ConsistencyTests/TextIndexing/WordOccurrence.cs
- `TextIndexing/WordOccurrences.cs` — MediaBrowser.Tests/ConsistencyTests/TextIndexing/WordOccurrences.cs

### MediaEncoding/Subtitles/ (6 files)

- `SrtParserTests.cs` — MediaBrowser.Tests/MediaEncoding/Subtitles/SrtParserTests.cs
- `VttWriterTest.cs` — MediaBrowser.Tests/MediaEncoding/Subtitles/VttWriterTest.cs
- `TestSubtitles/data.ass` — MediaBrowser.Tests/MediaEncoding/Subtitles/TestSubtitles/data.ass
- `TestSubtitles/data2.ass` — MediaBrowser.Tests/MediaEncoding/Subtitles/TestSubtitles/data2.ass
- `TestSubtitles/expected.vtt` — MediaBrowser.Tests/MediaEncoding/Subtitles/TestSubtitles/expected.vtt
- `TestSubtitles/unit.srt` — MediaBrowser.Tests/MediaEncoding/Subtitles/TestSubtitles/unit.srt

## Test Coverage

| Module | Description |
|--------|-------------|
| M3uParserTest | Playlist file parsing |
| SrtParserTests | SubRip subtitle parsing |
| VttWriterTest | WebVTT conversion |
| StringUsageReporter | Localization consistency |
| TextIndexing | Search index building |

## Test Framework

- xUnit or similar (check packages.config)

## Dependencies

- MediaBrowser.Controller
- MediaBrowser.Model
