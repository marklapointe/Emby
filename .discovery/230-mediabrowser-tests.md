# Component: MediaBrowser.Tests

**Path:** \`MediaBrowser.Tests/\`
**Type:** Module
**Maps to:** \`.discovery/230-mediabrowser-tests.md\`

## Description

Unit and integration tests for Emby Server.

## Structure

```
MediaBrowser.Tests/
├── MediaBrowser.Tests.csproj
├── Common/                   # Common test utilities
├── Controller/                 # Controller tests
├── Dlna/                       # DLNA tests
├── Drawing/                    # Drawing tests
├── IO/                         # I/O tests
├── Library/                    # Library tests
├── LiveTv/                     # LiveTV tests
├── MediaInfo/                  # Media info tests
├── Providers/                  # Provider tests
├── Serialization/              # Serialization tests
└── packages.config
```

## Test Framework

- NUnit or MSTest (verify in packages.config)
