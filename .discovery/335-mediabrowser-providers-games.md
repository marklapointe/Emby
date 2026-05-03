# MediaBrowser.Providers - Games Module

**Module:** MediaBrowser.Providers/Games
**Language:** C#
**Maps to:** `.discovery/335-mediabrowser-providers-games.md`

## Decomposition

### GameMetadataService.cs (Game Metadata Service)

#### Imports
```csharp
using MediaBrowser.Controller.Entities;
using MediaBrowser.Controller.Providers;
using System;
using System.Threading.Tasks;
```

#### Classes
`GameMetadataService` (public class : IMetadataService)

### GameSystemMetadataService.cs (Game System Metadata Service)

#### Classes
`GameSystemMetadataService` (public class : IMetadataService)

## Architecture

```mermaid
graph TD
    IMetadataService["IMetadataService<br/>(Interface)"]
    GameMetadataService["GameMetadataService"]
    GameSystemMetadataService["GameSystemMetadataService"]
    
    IMetadataService -.->|implement| GameMetadataService
    IMetadataService -.->|implement| GameSystemMetadataService
```

## File Listing

```
Games/
├── GameMetadataService.cs       - Game metadata service
└── GameSystemMetadataService.cs - Game system metadata service
```

## Description

Games module provides metadata services for game and game system items.

## Dependencies

- **MediaBrowser.Controller.Entities** - Game entities
- **MediaBrowser.Controller.Providers** - Provider interfaces
