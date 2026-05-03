# Component: MediaBrowser.Tests

**Path:** `MediaBrowser.Tests/`
**Type:** Directory | Test Project
**Language:** C#
**Maps to:** `.discovery/131-tests.md`

## Description

MediaBrowser.Tests is the unit and integration test project for Emby Server. It uses MSTest (Microsoft.VisualStudio.TestTools) framework and covers M3U playlist parsing, subtitle format conversion (SRT, VTT), and string consistency analysis.

## Structure

```
MediaBrowser.Tests/
├── MediaBrowser.Tests.csproj
├── Properties/
│   └── AssemblyInfo.cs            # Assembly metadata
├── M3uParserTest.cs               # M3U playlist parser tests
│   └── [class] M3uParserTest
│       ├── [method] public void TestM3uParser()
│       │   ├── Parses sample M3U playlist file
│       │   ├── Verifies channel count
│       │   └── Verifies channel properties (name, URL, logo)
│       └── [method] public void TestM3uParserWithAttributes()
│           ├── Parses M3U with EXTINF attributes
│           └── Verifies attribute extraction
├── MediaEncoding/
│   └── Subtitles/
│       ├── SrtParserTests.cs      # SRT subtitle parser tests
│       │   └── [class] SrtParserTests
│       │       ├── [method] public void TestSrtParse()
│       │       │   ├── Parses sample SRT file
│       │       │   ├── Verifies subtitle count
│       │       │   └── Verifies timing and text extraction
│       │       └── [method] public void TestSrtParseWithFormatting()
│       │           ├── Parses SRT with HTML formatting tags
│       │           └── Verifies formatting preservation
│       └── VttWriterTest.cs       # VTT subtitle writer tests
│           └── [class] VttWriterTest
│               ├── [method] public void TestVttWrite()
│               │   ├── Writes subtitle items to VTT format
│               │   ├── Verifies WEBVTT header
│               │   └── Verifies cue timing format
│               └── [method] public void TestVttWriteWithStyles()
│                   ├── Writes VTT with CSS styling
│                   └── Verifies style block output
└── ConsistencyTests/
    ├── StringUsageReporter.cs     # String usage analysis
    │   └── [class] StringUsageReporter
    │       ├── [method] public void ReportStringUsage()
    │       │   ├── Scans codebase for string literals
    │       │   ├── Identifies hardcoded strings
    │       │   └── Reports potential localization issues
    │       └── [method] public void ReportDuplicateStrings()
    │           ├── Finds duplicate string literals
    │           └── Suggests extraction to resources
    └── TextIndexing/
        ├── IndexBuilder.cs        # Text index builder
        │   └── [class] IndexBuilder
        │       ├── [method] public void BuildIndex()
        │       │   ├── Tokenizes text content
        │       │   ├── Builds inverted index
        │       │   └── Stores word occurrences
        │       └── [method] public void Search(string query)
        │           └── Returns matching documents
        ├── WordIndex.cs           # Word-to-occurrences index
        │   └── [class] WordIndex : Dictionary<string, WordOccurrences>
        │       └── Maps words to their occurrence lists
        ├── WordOccurrences.cs     # Word occurrence collection
        │   └── [class] WordOccurrences : List<WordOccurrence>
        │       └── Collection of occurrences for a single word
        └── WordOccurrence.cs      # Single word occurrence
            └── [class] WordOccurrence
                ├── [property] public string Word
                ├── [property] public string FilePath
                ├── [property] public int LineNumber
                └── [property] public int ColumnNumber
```

## Test Coverage

| Test Class | Target Component | Test Count | Framework |
|------------|-----------------|------------|-----------|
| M3uParserTest | Emby.Server.Implementations.LiveTv | 2 | MSTest |
| SrtParserTests | Emby.Server.Implementations.MediaEncoding.Subtitles | 2 | MSTest |
| VttWriterTest | Emby.Server.Implementations.MediaEncoding.Subtitles | 2 | MSTest |
| StringUsageReporter | Codebase consistency | 2 | MSTest |
| IndexBuilder | TextIndexing | 2 | MSTest |

## Data Flow

```mermaid
graph TD
    A[Test Method] --&gt; B[Arrange test data]
    B --&gt; C[Act: call target method]
    C --&gt; D[Assert: verify results]
    D --&gt; E[Test passes/fails]
```

## Side Effects

- No external network calls (uses sample/test data)
- No file system writes (uses in-memory streams)
- StringUsageReporter scans source files for analysis
