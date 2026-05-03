# Component: MediaBrowser.XbmcMetadata

**Path:** \`MediaBrowser.XbmcMetadata/\`
**Type:** Module
**Maps to:** \`.discovery/240-mediabrowser-xbmcmetadata.md\`

## Description

XBMC/Kodi metadata provider for Emby Server.

## Structure

```
MediaBrowser.XbmcMetadata/
├── MediaBrowser.XbmcMetadata.csproj
├── Configuration/
│   └── XbmcMetadataOptions.cs
├── Parsers/
│   ├── BaseNfoParser.cs
│   ├── EpisodeNfoParser.cs
│   ├── MovieNfoParser.cs
│   ├── SeasonNfoParser.cs
│   └── SeriesNfoParser.cs
├── Savers/
│   ├── BaseNfoSaver.cs
│   ├── EpisodeNfoSaver.cs
│   ├── MovieNfoSaver.cs
│   ├── SeasonNfoSaver.cs
│   └── SeriesNfoSaver.cs
├── XbmcMetadataProvider.cs
├── XbmcMetadataSaver.cs
└── packages.config
```

## Dependencies

- MediaBrowser.Providers
- MediaBrowser.Model
