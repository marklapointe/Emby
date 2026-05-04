# Component: Emby.Server.Implementations — Reflection

**Path:** `Emby.Server.Implementations/Reflection/`
**Type:** Directory | Module
**Language:** C#
**Maps to:** `.discovery/212-emby-server-impl-reflection.md`

## Description

Assembly metadata and reflection utilities. Contains only assembly attribute definitions.

## Files

- `AssemblyInfo.cs` — Emby.Server.Implementations/Reflection/AssemblyInfo.cs

## Content

### AssemblyInfo.cs

Contains standard .NET assembly attributes:

```csharp
using System.Reflection;

[assembly: AssemblyTitle("Emby.Server.Implementations")]
[assembly: AssemblyDescription("Server implementation library")]
[assembly: AssemblyVersion("4.0.0.0")]
[assembly: AssemblyFileVersion("4.0.0.0")]
```

## Dependencies

- `System.Reflection` — Assembly metadata

## Statistics

| Metric | Value |
|--------|-------|
| Files | 1 |
| Classes | 0 |
| LOC | ~10 |
