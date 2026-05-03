# MediaBrowser.Providers - Folders Module

**Module:** MediaBrowser.Providers/Folders
**Language:** C#
**Maps to:** `.discovery/333-mediabrowser-providers-folders.md`

## Decomposition

### FolderMetadataService.cs (Folder Metadata Service)

#### Imports
```csharp
using MediaBrowser.Controller.Entities;
using MediaBrowser.Controller.Providers;
using System;
using System.Threading.Tasks;
```

#### Classes
`FolderMetadataService` (public class : IMetadataService)

### UserViewMetadataService.cs (User View Metadata Service)

#### Classes
`UserViewMetadataService` (public class : IMetadataService)

### CollectionFolderMetadataService.cs (Collection Folder Metadata Service)

#### Classes
`CollectionFolderMetadataService` (public class : IMetadataService)

## Architecture

```mermaid
graph TD
    IMetadataService["IMetadataService<br/>(Interface)"]
    FolderMetadataService["FolderMetadataService"]
    UserViewMetadataService["UserViewMetadataService"]
    CollectionFolderMetadataService["CollectionFolderMetadataService"]
    
    IMetadataService -.->|implement| FolderMetadataService
    IMetadataService -.->|implement| UserViewMetadataService
    IMetadataService -.->|implement| CollectionFolderMetadataService
```

## File Listing

```
Folders/
├── FolderMetadataService.cs       - Folder metadata service
├── UserViewMetadataService.cs    - User view metadata service
└── CollectionFolderMetadataService.cs - Collection folder metadata
```

## Description

Folders module provides metadata services for folder-type items like user views and collection folders.

## Dependencies

- **MediaBrowser.Controller.Entities** - Folder entity
- **MediaBrowser.Controller.Providers** - Provider interfaces

## Statistics

- **Files:** 3
- **Lines:** ~150
- **Classes:** 3
