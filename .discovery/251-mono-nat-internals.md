# Component: Mono.Nat — Internals

**Path:** `Mono.Nat/`
**Type:** Module
**Language:** C#
**Maps to:** `.discovery/251-mono-nat-internals.md`
**Parent:** `.discovery/250-mono-nat.md`

## Description

Network Address Translation (NAT) traversal library. Provides UPnP and PMP
port mapping functionality for opening ports through routers.

## Structure

```
Mono.Nat/
├── NatUtility.cs                 # [class] NatUtility
│   ├── Entry point for NAT discovery
│   ├── Discovers NAT devices
│   └── Creates port mappings
├── AbstractNatDevice.cs          # [class] AbstractNatDevice
│   └── Base NAT device class
├── INatDevice.cs                 # [interface] INatDevice
│   └── NAT device interface
├── Mapping.cs                    # [class] Mapping
│   └── Port mapping definition
├── NatProtocol.cs                # [enum] NatProtocol
│   └── UPnP / PMP protocols
├── Upnp/
│   ├── UpnpNatDevice.cs          # [class] UpnpNatDevice
│   │   └── UPnP NAT device
│   ├── UpnpSearcher.cs           # [class] UpnpSearcher
│   │   └── UPnP device discovery
│   └── *Upnp*.cs                 # UPnP helpers
└── Pmp/
    ├── PmpNatDevice.cs           # [class] PmpNatDevice
    │   └── PMP NAT device
    ├── PmpSearcher.cs            # [class] PmpSearcher
    │   └── PMP device discovery
    └── *Pmp*.cs                  # PMP helpers
```

## Key Classes

| Class | File | Purpose |
|-------|------|---------|
| `NatUtility` | `NatUtility.cs` | NAT entry point |
| `UpnpNatDevice` | `Upnp/UpnpNatDevice.cs` | UPnP device |
| `PmpNatDevice` | `Pmp/PmpNatDevice.cs` | PMP device |
| `Mapping` | `Mapping.cs` | Port mapping |

## NAT Traversal Flow

```mermaid
graph TD
    A[NatUtility.DiscoverDeviceAsync] --> B{Protocol?}
    B --> C[UpnpSearcher]
    B --> D[PmpSearcher]
    C --> E[UpnpNatDevice]
    D --> F[PmpNatDevice]
    E --> G[CreatePortMap]
    F --> G
```

## Decomposition

### NatUtility.cs (Entry Point)

#### Imports
```csharp
using System;
using System.Collections.Generic;
using System.Net;
using System.Threading.Tasks;
```

#### Classes
`NatUtility` (public static class)

#### Key Events
| Event | Description |
|-------|-------------|
| `DeviceFound` | NAT device discovered |
| `DeviceLost` | NAT device unavailable |

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `DiscoverDevicesAsync()` | `Task<IEnumerable<INatDevice>>` | Find NAT devices |
| `AddProtocol(Protocol)` | `void` | Add search protocol |
| `RemoveProtocol(Protocol)` | `void` | Remove search protocol |

### Mapping.cs (Port Mapping)

#### Classes
`Mapping` (public class)

#### Key Properties
| Property | Type | Description |
|----------|------|-------------|
| `Protocol` | `NatProtocol` | UPnP or PMP |
| `PrivatePort` | `int` | Internal port |
| `PublicPort` | `int` | External port |
| `Description` | `string` | Mapping description |
| `Lifetime` | `int` | Lease duration |

### UpnpNatDevice.cs (UPnP Device)

#### Classes
`UpnpNatDevice` (public class : AbstractNatDevice)

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `CreatePortMap(Mapping)` | `Task` | Create port mapping |
| `DeletePortMap(Mapping)` | `Task` | Remove port mapping |
| `GetExternalIP()` | `Task<IPAddress>` | Get WAN IP |

### PmpNatDevice.cs (PMP Device)

#### Classes
`PmpNatDevice` (public class : AbstractNatDevice)

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `CreatePortMap(Mapping)` | `Task` | Create port mapping |
| `DeletePortMap(Mapping)` | `Task` | Remove port mapping |

## Usage Example

```csharp
// Discover NAT devices
var devices = await NatUtility.DiscoverDevicesAsync();

// Create port mapping
var mapping = new Mapping(NatProtocol.Udp, 8080, 8080)
{
    Description = "Emby Server"
};
await device.CreatePortMap(mapping);
```
