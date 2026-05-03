# MediaBrowser.Api - Devices Subdirectory

**Module:** MediaBrowser.Api/Devices
**Language:** C#
**Maps to:** `.discovery/345-mediabrowser-api-devices.md`

## Decomposition

### DeviceService.cs (Device Management Service)

#### Imports
```csharp
using MediaBrowser.Controller.Devices;
using MediaBrowser.Controller.Net;
using MediaBrowser.Model.Services;
using System.Threading.Tasks;
```

#### Classes
`DeviceService` (public class : IRequiresRequest)

## File Listing

```
Devices/
└── DeviceService.cs - Device management API service
```

## Description

Devices subdirectory contains device management API services for registering and querying client devices.

## Dependencies

- **MediaBrowser.Controller.Devices** - Device interfaces
- **MediaBrowser.Model.Services** - Service interfaces
