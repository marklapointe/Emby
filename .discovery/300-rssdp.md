# Component: RSSDP

**Path:** \`RSSDP/\`
**Type:** Module
**Maps to:** \`.discovery/300-rssdp.md\`

## Description

Really Simple Service Discovery Protocol (SSDP) implementation for .NET.

## Structure

```
RSSDP/
├── RSSDP.csproj
├── Infrastructure/
│   ├── ISsdpCommunicationsServer.cs
│   ├── ISsdpDeviceLocator.cs
│   ├── ISsdpDevicePublisher.cs
│   ├── SsdpCommunicationsServer.cs
│   ├── SsdpDeviceLocator.cs
│   └── SsdpDevicePublisher.cs
├── SsdpDevice.cs             # SSDP device base
├── SsdpEmbeddedDevice.cs     # Embedded device
├── SsdpRootDevice.cs         # Root device
├── SsdpService.cs            # SSDP service
├── SsdpConstants.cs          # Protocol constants
└── Properties/
    └── AssemblyInfo.cs
```

## Dependencies

- None (standalone library)
