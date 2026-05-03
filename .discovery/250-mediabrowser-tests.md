# Component: MediaBrowser.Tests

**Path:** `MediaBrowser.Tests/`
**Type:** Directory | Test Project
**Language:** C#
**Maps to:** `.discovery/250-mediabrowser-tests.md`

## Description

MediaBrowser.Tests contains unit and integration tests for Emby Server. It uses a testing framework (likely NUnit or MSTest) to validate core server functionality, API endpoints, metadata parsing, and media handling.

## Structure

```
MediaBrowser.Tests/
├── MediaBrowser.Tests.csproj    # Test project file
├── ...                          # Test classes
└── Properties/                  # Assembly info
```

## Key Test Areas

| Area | Description |
|------|-------------|
| Library | Media library management tests |
| Providers | Metadata provider tests |
| API | REST API endpoint tests |
| Drawing | Image processing tests |
| DLNA | DLNA protocol tests |

## Dependencies

- `MediaBrowser.Api` — API under test
- `Emby.Server.Implementations` — Server under test
- Test framework (NUnit/MSTest)

## Reference

- Run tests via Visual Studio Test Explorer or `dotnet test`
