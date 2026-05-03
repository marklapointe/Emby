# Component: SocketHttpListener

**Path:** `SocketHttpListener/`
**Type:** Directory | Library
**Language:** C#
**Maps to:** `.discovery/330-sockethttplistener.md`

## Description

SocketHttpListener is a custom HTTP server implementation built on .NET sockets. It provides cross-platform HTTP request handling for Emby Server, replacing the need for IIS or other external web servers. Supports HTTP/1.1, WebSocket upgrades, and SSL/TLS.

## Structure

```
SocketHttpListener/
├── SocketHttpListener.csproj    # Project file
├── Net/                         # HTTP protocol implementation
│   ├── HttpListener.cs          # Main listener → [class] HttpListener
│   ├── HttpListenerContext.cs   # Request/response context
│   ├── HttpListenerRequest.cs   # HTTP request parsing
│   ├── HttpListenerResponse.cs  # HTTP response generation
│   └── ...                      # Cookie, authentication, SSL
├── Primitives/                  # Low-level socket primitives
│   ├── SocketFactory.cs         # Socket creation
│   └── ...                      
└── Properties/                  # Assembly info
```

## Key Classes

| Class | File | Purpose |
|-------|------|---------|
| `HttpListener` | `Net/HttpListener.cs` | Main HTTP server |
| `HttpListenerContext` | `Net/HttpListenerContext.cs` | Request/response pair |
| `HttpListenerRequest` | `Net/HttpListenerRequest.cs` | Parses HTTP requests |
| `HttpListenerResponse` | `Net/HttpListenerResponse.cs` | Builds HTTP responses |

## Data Flow

```mermaid
graph LR
    A[Client Request] --> B[TCP Socket]
    B --> C[HttpListener]
    C --> D[HttpListenerRequest]
    D --> E[Route Resolution]
    E --> F[MediaBrowser.Api]
    F --> G[HttpListenerResponse]
    G --> H[TCP Socket]
    H --> I[Client]
```

## Dependencies

- `System.Net.Sockets` — TCP socket layer
- `MediaBrowser.Controller` — Request routing

## Side Effects

- Binds to TCP ports (default 8096)
- Accepts incoming connections
- Manages SSL certificates
- Handles WebSocket upgrades

## Reference

- HTTP/1.1 specification (RFC 7230)
