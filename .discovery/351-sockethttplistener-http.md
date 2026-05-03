# Component: SocketHttpListener — HTTP

**Path:** `SocketHttpListener/`
**Type:** Module
**Language:** C#
**Maps to:** `.discovery/351-sockethttplistener-http.md`
**Parent:** `.discovery/350-sockethttplistener.md`

## Description

Custom HTTP listener implementation. Provides HTTP server functionality
without relying on Windows HTTP.sys.

## Structure

```
SocketHttpListener/
├── HttpListener.cs               # [class] HttpListener
│   ├── Main HTTP listener
│   ├── Accepts incoming connections
│   └── Routes requests to handlers
├── HttpListenerContext.cs        # [class] HttpListenerContext
│   └── Request/response context
├── HttpListenerRequest.cs        # [class] HttpListenerRequest
│   └── HTTP request parsing
├── HttpListenerResponse.cs       # [class] HttpListenerResponse
│   └── HTTP response construction
├── HttpWebSocket.cs              # [class] HttpWebSocket
│   └── WebSocket support
└── *Http*.cs                     # Supporting HTTP classes
```

## Key Classes

| Class | File | Purpose |
|-------|------|---------|
| `HttpListener` | `HttpListener.cs` | HTTP server core |
| `HttpListenerContext` | `HttpListenerContext.cs` | Request context |
| `HttpListenerRequest` | `HttpListenerRequest.cs` | Request parsing |
| `HttpListenerResponse` | `HttpListenerResponse.cs` | Response building |
| `HttpWebSocket` | `HttpWebSocket.cs` | WebSocket support |
