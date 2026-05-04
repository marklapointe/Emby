# Component: Emby.Server.Implementations — Services

**Path:** `Emby.Server.Implementations/Services/`
**Type:** Directory | Module
**Language:** C#
**Maps to:** `.discovery/214-emby-server-impl-services.md`

## Description

Service layer and dependency injection container. Manages service registration and resolution.

## Files

- `ServiceController.cs` — Emby.Server.Implementations/Services/ServiceController.cs
- `ServiceHost.cs` — Emby.Server.Implementations/Services/ServiceHost.cs
- `ServiceRegistration.cs` — Emby.Server.Implementations/Services/ServiceRegistration.cs

### Subdirectories (6 files each)

- `HttpListenerServiceHost/` — 2 files
- `RestSharp/` — 4 files

## Decomposition

### ServiceHost.cs (Service Host)

#### Imports
```csharp
using MediaBrowser.Common.Net;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
```

#### Classes
`ServiceHost` (public class : IServiceHost, IDisposable)

#### Key Properties
| Property | Type | Description |
|----------|------|-------------|
| `Services` | `IDictionary<Type, object>` | Registered services |

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `Register<TInterface, TImplementation>()` | `void` | Register service |
| `Resolve<T>()` | `T` | Get service instance |
| `GetAllInstances<T>()` | `IEnumerable<T>` | Get all implementations |

### ServiceController.cs (Service Controller)

#### Classes
`ServiceController` (public class : IServiceController)

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `ExecuteRequest(ServiceRequest)` | `Task<ServiceResult>` | Execute service request |
| `RegisterService(object)` | `void` | Register service instance |

## Data Flow

```mermaid
graph LR
    A[Request] --> B[ServiceController]
    B --> C[ServiceHost]
    C --> D[Service Resolution]
    D --> E[Service Implementation]
```

## Dependencies

- `MediaBrowser.Common` — Common interfaces
- `MediaBrowser.Common.Net` — Network interfaces

## Statistics

| Metric | Value |
|--------|-------|
| Files | 9 |
| Classes | 5 |
| LOC | ~400 |
