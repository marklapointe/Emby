# Component: Emby.Server.Implementations — Networking

**Path:** `Emby.Server.Implementations/Networking/`
**Type:** Directory | Module
**Language:** C#
**Maps to:** `.discovery/209a-emby-server-impl-networking.md`

## Description

Advanced networking utilities including IP network address calculations and subnet management.

## Directory Structure

```
Networking/
├── IPNetwork/
│   ├── BigIntegerExt.cs
│   ├── IPAddressCollection.cs
│   ├── IPNetworkCollection.cs
│   ├── IPNetwork.cs
│   └── LICENSE.txt
└── (root files)
```

## Files

### Root Networking Files

- (none)

### IPNetwork/ (IP Network Utilities)

- `BigIntegerExt.cs` — Emby.Server.Implementations/Networking/IPNetwork/BigIntegerExt.cs
- `IPAddressCollection.cs` — Emby.Server.Implementations/Networking/IPNetwork/IPAddressCollection.cs
- `IPNetworkCollection.cs` — Emby.Server.Implementations/Networking/IPNetwork/IPNetworkCollection.cs
- `IPNetwork.cs` — Emby.Server.Implementations/Networking/IPNetwork/IPNetwork.cs

## Decomposition

### IPNetwork.cs (IP Network Calculator)

#### Imports
```csharp
using System;
using System.Collections;
using System.Collections.Generic;
using System.Net;
using System.Numerics;
```

#### Classes
`IPNetwork` (public class)

#### Key Properties
| Property | Type | Description |
|----------|------|-------------|
| `Network` | `IPAddress` | Network address |
| `Netmask` | `IPAddress` | Subnet mask |
| `Cidr` | `int` | CIDR notation |
| `NetworkAddress` | `IPAddress` | Start of network |
| `Broadcast` | `IPAddress` | End of network |
| `Count` | `BigInteger` | Usable hosts |

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `Parse(string)` | `static IPNetwork` | Parse CIDR string |
| `Contains(IPAddress)` | `bool` | Is IP in network |
| `Subnet(int)` | `IPNetwork[]` | Create subnets |
| `Supernet(IPNetwork)` | `static IPNetwork` | Combine networks |

### IPNetworkCollection.cs (IP Network Collection)

#### Classes
`IPNetworkCollection` (public class : ICollection<IPNetwork>)

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `Contains(IPAddress)` | `bool` | Check if IP matches any |
| `Match(IPAddress)` | `IPNetwork` | Find matching network |

## Data Flow

```mermaid
graph LR
    A[IP Address] --> B[IPNetwork]
    B --> C[Subnet Check]
    C --> D[Match/No Match]
```

## Dependencies

- `System.Net` — IP address types
- `System.Numerics` — BigInteger support

## Statistics

| Metric | Value |
|--------|-------|
| Files | 4 |
| Classes | 3 |
| LOC | ~5,500 (IPNetwork.cs is ~1800 lines) |
