# Component: Emby.Dlna — Profiles

**Path:** `Emby.Dlna/Profiles/`
**Type:** Directory | Sub-module
**Language:** C#
**Maps to:** `.discovery/331-emby-dlna-profiles.md`
**Parent:** `.discovery/330-emby-dlna.md`

## Description

DLNA device profiles define the capabilities of specific devices (TVs, consoles, etc.)
and how media should be transcoded for them.

## Structure

```
Profiles/
├── DefaultProfile.cs             # [class] DefaultProfile
│   └── Base profile with default settings
├── *Profile.cs                   # Device-specific profiles
│   ├── SamsungProfile.cs         # Samsung TVs
│   ├── SonyProfile.cs            # Sony devices
│   ├── LgProfile.cs              # LG TVs
│   └── XboxProfile.cs            # Xbox consoles
└── ProfileHelper.cs              # [class] ProfileHelper
```

## Key Classes

| Class | File | Purpose |
|-------|------|---------|
| `DefaultProfile` | `DefaultProfile.cs` | Base DLNA profile |
| Various | `*Profile.cs` | Device-specific profiles |
