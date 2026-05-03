# Component: Emby.Server.Implementations — UserViews

**Path:** `Emby.Server.Implementations/UserViews/`
**Type:** Directory | Sub-module
**Language:** C#
**Maps to:** `.discovery/179-emby-server-impl-userviews.md`
**Parent:** `.discovery/160-emby-server-impl.md`

## Description

User view management — creates personalized content views (Latest, Next Up, etc.)
for each user based on their library access and preferences.

## Structure

```
UserViews/
├── UserViewManager.cs            # [class] UserViewManager
│   ├── Creates user-specific views
│   ├── Latest items view
│   ├── Next Up (TV) view
│   └── Resume items view
└── *View*.cs                     # Supporting view classes
```

## Key Classes

| Class | File | Purpose |
|-------|------|---------|
| `UserViewManager` | `UserViewManager.cs` | User view creation |
