# MediaBrowser.Providers - LiveTv Module

**Module:** MediaBrowser.Providers/LiveTv
**Language:** C#
**Maps to:** `.discovery/337-mediabrowser-providers-livetv.md`

## Decomposition

### ProgramMetadataService.cs (Live TV Program Metadata Service)

#### Imports
```csharp
using MediaBrowser.Controller.Entities;
using MediaBrowser.Controller.LiveTv;
using MediaBrowser.Controller.Providers;
using System;
using System.Threading.Tasks;
```

#### Classes
`ProgramMetadataService` (public class : IMetadataService)

## File Listing

```
LiveTv/
└── ProgramMetadataService.cs - Live TV program metadata service
```

## Description

LiveTv module provides metadata services for live TV program items.

## Dependencies

- **MediaBrowser.Controller.Entities** - Program entity
- **MediaBrowser.Controller.LiveTv** - Live TV interfaces
- **MediaBrowser.Controller.Providers** - Provider interfaces
