# MediaBrowser.Providers - Manager Module

**Module:** MediaBrowser.Providers/Manager
**Language:** C#
**Maps to:** `.discovery/338-mediabrowser-providers-manager.md`

## Decomposition

### ProviderManager.cs (Main Provider Manager - Core Component)

#### Imports
```csharp
using MediaBrowser.Controller.Configuration;
using MediaBrowser.Controller.Entities;
using MediaBrowser.Controller.Library;
using MediaBrowser.Controller.Providers;
using MediaBrowser.Model.Configuration;
using MediaBrowser.Model.Entities;
using MediaBrowser.Model.IO;
using MediaBrowser.Model.Logging;
using MediaBrowser.Model.Providers;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading;
using System.Threading.Tasks;
```

#### Classes
`ProviderManager` (public class : IProviderManager, IServerEntryPoint)

#### Key Properties
```csharp
IEnumerable<IMetadataService> MetadataServices { get; }
IEnumerable<IMetadataProvider> MetadataProviders { get; }
IEnumerable<IRemoteImageProvider> RemoteImageProviders { get; }
```

#### Key Methods
```csharp
Task<ItemUpdateType> RefreshSingleItem(...)
Task<HttpResponseInfo> GetRemoteImageAsync(...)
Task SaveImageAsync(BaseItem item, string source, ...)
void SaveImage(BaseItem item, string source, ...)
```

### MetadataService.cs (Base Metadata Service)

#### Classes
`MetadataService` (public abstract class : IMetadataService)

### ItemImageProvider.cs (Item Image Provider)

#### Classes
`ItemImageProvider` (public class : IItemImageProvider)

### ImageSaver.cs (Image Persistence)

#### Classes
`ImageSaver` (public static class)

### ProviderUtils.cs (Provider Utilities)

#### Classes
`ProviderUtils` (public static class)

### Priority Queue Classes

- `GenericPriorityQueue<T, TPriority>` - Generic priority queue
- `SimplePriorityQueue<T, TPriority>` - Simple priority queue
- `IPriorityQueue<T, TPriority>` / `IFixedSizePriorityQueue<T, TPriority>` - Interfaces
- `GenericPriorityQueueNode<T, TPriority>` - Node wrapper

## Architecture

```mermaid
graph TD
    IProviderManager["IProviderManager<br/>(Interface)"]
    ProviderManager["ProviderManager<br/>(Implementation)"]
    
    ProviderManager -->|uses| MetadataService["MetadataService<br/>(Abstract)"]
    ProviderManager -->|uses| ItemImageProvider["ItemImageProvider"]
    ProviderManager -->|uses| ImageSaver["ImageSaver"]
    ProviderManager -->|uses| GenericPriorityQueue["GenericPriorityQueue"]
```

## File Listing

```
Manager/
├── ProviderManager.cs           - Main provider manager (core)
├── MetadataService.cs          - Base metadata service
├── ItemImageProvider.cs         - Item image provider
├── ImageSaver.cs                - Image persistence
├── ProviderUtils.cs             - Provider utilities
├── GenericPriorityQueue.cs      - Generic priority queue
├── SimplePriorityQueue.cs       - Simple priority queue
├── IPriorityQueue.cs            - Queue interface
├── IFixedSizePriorityQueue.cs   - Fixed size queue interface
└── GenericPriorityQueueNode.cs  - Queue node
```

## Description

Manager module is the core of the MediaBrowser providers system. ProviderManager coordinates all metadata services, metadata providers, and image providers. MetadataService provides base functionality for all metadata services. Priority queues manage provider execution order.

## Dependencies

- **MediaBrowser.Controller.Providers** - Provider interfaces
- **MediaBrowser.Controller.Entities** - BaseItem entity
- **MediaBrowser.Controller.Configuration** - Configuration interfaces
- **MediaBrowser.Model.Providers** - Provider models

## Statistics

- **Files:** 10
- **Lines:** ~3,000
- **Classes:** 10
