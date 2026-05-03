# Component: SocketHttpListener

**Path:** \`SocketHttpListener/\`
**Type:** Directory | Module
**Language:** C#
**Maps to:** \`.discovery/350-sockethttplistener.md\`

## Description

SocketHttpListener is a custom HTTP listener implementation for Emby Server. It provides a cross-platform HTTP server built on top of .NET sockets. Contains 61 C# files.

## Sub-Modules

- `Net/` — 39 C# files
- `Net/WebSockets/` — 6 C# files
- `Primitives/` — 1 C# files
- `Properties/` — 1 C# files

## Root Files

- `ByteOrder.cs` — SocketHttpListener/ByteOrder.cs
- `CloseEventArgs.cs` — SocketHttpListener/CloseEventArgs.cs
- `CloseStatusCode.cs` — SocketHttpListener/CloseStatusCode.cs
- `CompressionMethod.cs` — SocketHttpListener/CompressionMethod.cs
- `ErrorEventArgs.cs` — SocketHttpListener/ErrorEventArgs.cs
- `Ext.cs` — SocketHttpListener/Ext.cs
- `Fin.cs` — SocketHttpListener/Fin.cs
- `HttpBase.cs` — SocketHttpListener/HttpBase.cs
- `HttpResponse.cs` — SocketHttpListener/HttpResponse.cs
- `Mask.cs` — SocketHttpListener/Mask.cs
- `MessageEventArgs.cs` — SocketHttpListener/MessageEventArgs.cs
- `Opcode.cs` — SocketHttpListener/Opcode.cs
- `PayloadData.cs` — SocketHttpListener/PayloadData.cs
- `Rsv.cs` — SocketHttpListener/Rsv.cs
- `SocketStream.cs` — SocketHttpListener/SocketStream.cs
- `WebSocket.cs` — SocketHttpListener/WebSocket.cs
- `WebSocketException.cs` — SocketHttpListener/WebSocketException.cs
- `WebSocketFrame.cs` — SocketHttpListener/WebSocketFrame.cs

## Project Files

- `packages.config` — SocketHttpListener/packages.config
- `SocketHttpListener.csproj` — SocketHttpListener/SocketHttpListener.csproj

## Decomposition

### WebSocket.cs (WebSocket Implementation)

#### Imports
```csharp
using System;
using System.IO;
using System.Net;
using System.Net.Sockets;
using System.Text;
using System.Threading.Tasks;
```

#### Classes
`WebSocket` (public class)

#### Key Properties
| Property | Type | Description |
|----------|------|-------------|
| `State` | `WebSocketState` | Connection state |
| `Url` | `Uri` | WebSocket URL |

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `Connect()` | `Task` | Connect to server |
| `Close()` | `Task` | Close connection |
| `Send(string)` | `Task` | Send text message |
| `Send(byte[])` | `Task` | Send binary message |
| `Receive()` | `Task<MessageEventArgs>` | Receive message |

### HttpResponse.cs (HTTP Response Builder)

#### Classes
`HttpResponse` (public class : HttpBase)

#### Key Properties
| Property | Type | Description |
|----------|------|-------------|
| `StatusCode` | `int` | HTTP status code |
| `Headers` | `NameValueCollection` | Response headers |
| `Body` | `byte[]` | Response body |

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `WriteTo(Stream)` | `Task` | Write to stream |
| `SetHeader(string, string)` | `void` | Set header |

### WebSocketFrame.cs (Frame Parser)

#### Classes
`WebSocketFrame` (public static class)

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `CreateFrame(Opcode, byte[], bool)` | `byte[]` | Create frame |
| `ReadFrame(Stream)` | `WebSocketFrame` | Read frame |
| `GetOpcode(byte)` | `Opcode` | Get opcode |
| `IsFinalFrame(byte)` | `bool` | Check final flag |

### SocketStream.cs (Socket I/O)

#### Classes
`SocketStream` (public class : Stream)

#### Key Properties
| Property | Type | Description |
|----------|------|-------------|
| `Socket` | `Socket` | Underlying socket |
| `IsConnected` | `bool` | Connection status |

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `ReceiveAsync(byte[], int, int)` | `Task<int>` | Receive data |
| `SendAsync(byte[], int, int)` | `Task` | Send data |
