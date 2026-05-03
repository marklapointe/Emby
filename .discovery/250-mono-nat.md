# Component: Mono.Nat

**Path:** \`Mono.Nat/\`
**Type:** Module
**Maps to:** \`.discovery/250-mono-nat.md\`

## Description

NAT traversal library for Mono/.NET. Provides UPnP and NAT-PMP support.

## Structure

```
Mono.Nat/
├── Mono.Nat.csproj
├── AbstractNatDevice.cs      # Base NAT device
├── NatDevice.cs              # NAT device interface
├── NatUtility.cs             # NAT utilities
├── Upnp/
│   ├── UpnpNatDevice.cs
│   ├── UpnpSearcher.cs
│   └── Messages/
│       ├── DiscoverDeviceMessage.cs
│       ├── CreatePortMappingMessage.cs
│       └── ...
├── Pmp/
│   ├── PmpNatDevice.cs
│   ├── PmpSearcher.cs
│   └── Messages/
│       └── ...
└── Properties/
    └── AssemblyInfo.cs
```

## Dependencies

- None (standalone library)
