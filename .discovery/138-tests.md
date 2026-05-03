# Component: MediaBrowser.Tests

**Path:** `MediaBrowser.Tests/`
**Type:** Directory | Test Suite
**Language:** C#
**Maps to:** `.discovery/138-tests.md`

## Description

MediaBrowser.Tests contains the unit and integration test suite for Emby Server. It uses MSTest (Microsoft.VisualStudio.TestTools) as the test framework. Tests cover M3U playlist parsing, subtitle format conversion (SRT/VTT), and string usage consistency across the codebase.

## Structure

```
MediaBrowser.Tests/
├── MediaBrowser.Tests.csproj
├── Properties/
│   └── AssemblyInfo.cs            # Assembly metadata
├── M3uParserTest.cs               # M3U playlist parser tests
│   └── [class] M3uParserTest
│       ├── [test] public void TestM3uParser()
│       │   ├── Parses sample M3U playlist file
│       │   ├── Verifies channel count
│       │   └── Verifies channel properties (name, URL, logo)
│       ├── [test] public void TestM3uParserWithOptions()
│       │   ├── Parses M3U with extended attributes
│       │   └── Verifies tvg-id, group-title attributes
│       └── [test] public void TestM3uParserEmpty()
│           └── Verifies empty playlist returns empty list
├── MediaEncoding/
│   └── Subtitles/
│       ├── SrtParserTests.cs      # SRT subtitle parser tests
│       │   └── [class] SrtParserTests
│       │       ├── [test] public void TestSrtParse()
│       │       │   ├── Parses sample SRT file
│       │       │   ├── Verifies subtitle count
│       │       │   └── Verifies timing and text content
│       │       ├── [test] public void TestSrtParseWithFormatting()
│       │       │   └── Verifies bold/italic/underline tags preserved
│       │       └── [test] public void TestSrtParseInvalid()
│       │           └── Verifies graceful handling of malformed SRT
│       └── VttWriterTest.cs       # VTT subtitle writer tests
│           └── [class] VttWriterTest
│               ├── [test] public void TestVttWrite()
│               │   ├── Converts SRT to VTT format
│               │   ├── Verifies VTT header (WEBVTT)
│               │   └── Verifies cue timing format
│               └── [test] public void TestVttWriteWithStyles()
│                   └── Verifies CSS style blocks in VTT output
└── ConsistencyTests/
    ├── StringUsageReporter.cs     # String usage analysis
    │   └── [class] StringUsageReporter
    │       ├── [method] public void ReportStringUsage()
    │       │   ├── Scans all source files for string literals
    │       │   ├── Identifies hardcoded strings
    │       │   └── Generates report of string usage patterns
    │       └── [method] public void ReportLocalizationIssues()
    │           └── Flags strings that should be localized
    └── TextIndexing/              # Text indexing tests
        ├── IndexBuilder.cs        # Inverted index builder
        │   └── [class] IndexBuilder
        │       ├── [method] public void BuildIndex()
        │       │   ├── Tokenizes text into words
        │       │   ├── Builds inverted index (word → documents)
        │       │   └── Supports stemming and stop words
        │       └── [method] public List<string> Search(string query)
        │           └── Returns documents matching query terms
        ├── WordIndex.cs           # Word-to-documents index
        │   └── [class] WordIndex : Dictionary<string, WordOccurrences>
        │       └── Maps words to their occurrences across documents
        ├── WordOccurrences.cs     # Occurrence list
        │   └── [class] WordOccurrences : List<WordOccurrence>
        │       └── List of occurrences for a specific word
        └── WordOccurrence.cs      # Single occurrence
            └── [class] WordOccurrence
                ├── Properties: FilePath, LineNumber, Column, Context
                └── Represents one occurrence of a word in source code
```

## Test Framework

| Aspect | Details |
|--------|---------|
| Framework | MSTest (Microsoft.VisualStudio.TestTools) |
| Runner | Visual Studio Test Explorer / vstest.console |
| Coverage | M3U parsing, subtitle conversion, string analysis |

## Test Categories

| Category | Files | Description |
|----------|-------|-------------|
| Playlist Parsing | M3uParserTest.cs | M3U/M3U8 playlist format |
| Subtitle Conversion | SrtParserTests.cs, VttWriterTest.cs | SRT ↔ VTT |
| Code Quality | StringUsageReporter.cs | Hardcoded string detection |
| Text Indexing | IndexBuilder.cs, WordIndex.cs | Inverted index construction |

## Data Flow

```mermaid
graph TD
    A[Test Runner] --&gt; B[M3uParserTest]
    A --&gt; C[SrtParserTests]
    A --&gt; D[VttWriterTest]
    A --&gt; E[StringUsageReporter]
    B --&gt; F[Parse M3U file]
    F --&gt; G[Assert channel properties]
    C --&gt; H[Parse SRT file]
    H --&gt; I[Assert subtitle entries]
    D --&gt; J[Convert SRT to VTT]
    J --&gt; K[Assert VTT format]
    E --&gt; L[Scan source files]
    L --&gt; M[Report string usage]
```

## Side Effects

- Reads test fixture files (sample M3U, SRT files)
- No external network calls
- No file writes (except test output)
