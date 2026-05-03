# Component: TSStreamBuffer.cs

**Path:** `BDInfo/TSStreamBuffer.cs`
**Type:** File | Class
**Language:** C#
**Maps to:** `.discovery/100-10-tsstreambuffer.md`
**Parent:** `.discovery/100-bdinfo.md`

## Structure

```
TSStreamBuffer.cs
├── [namespace] BDInfo
│   └── [class] public class TSStreamBuffer
│       ├── [field] private byte[] _buffer
│       ├── [field] private int _length
│       ├── [field] private int _position
│       ├── [constructor] TSStreamBuffer()
│       ├── [method] public void Write(byte[] data, int offset, int count)
│       │   └── Appends data to internal buffer, resizing if needed
│       ├── [method] public byte[] Read(int count)
│       │   └── Reads specified bytes from current position
│       ├── [method] public void Clear()
│       │   └── Resets buffer and position
│       ├── [property] public int Length
│       │   └── Returns current data length
│       └── [property] public int Position
│           └── Gets/sets current read position
```

## Description

`TSStreamBuffer` is a simple byte buffer for accumulating and reading transport stream data during Blu-ray scanning. It provides write, read, and clear operations with automatic buffer resizing.
