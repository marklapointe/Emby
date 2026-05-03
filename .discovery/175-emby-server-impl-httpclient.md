# Component: Emby.Server.Implementations — HttpClientManager

**Path:** `Emby.Server.Implementations/HttpClientManager/`
**Type:** Directory | Sub-module
**Language:** C#
**Maps to:** `.discovery/175-emby-server-impl-httpclient.md`
**Parent:** `.discovery/160-emby-server-impl.md`

## Description

HTTP client management for outgoing requests. Provides centralized HTTP client
configuration, request/response handling, and network utilities.

## Structure

```
HttpClientManager/
├── HttpClientManager.cs          # [class] HttpClientManager → IHttpClient
│   ├── Manages HttpClient instances
│   ├── Handles request configuration
│   └── Provides download/upload helpers
└── *HttpClient*.cs               # Supporting classes
```

## Key Classes

| Class | File | Purpose |
|-------|------|---------|
| `HttpClientManager` | `HttpClientManager.cs` | Central HTTP client manager |
