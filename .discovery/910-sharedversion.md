# Component: SharedVersion.cs

**Path:** `SharedVersion.cs`
**Type:** File | Configuration
**Language:** C#
**Maps to:** `.discovery/910-sharedversion.md`

## Description

SharedVersion.cs defines the assembly version for all projects in the Emby solution. It uses the `AssemblyVersion` attribute to set a consistent version number (3.5.3.0) across all compiled assemblies.

## Structure

```
SharedVersion.cs
├── [import] System.Reflection
└── [attribute] AssemblyVersion("3.5.3.0")
    └── Applied to all assemblies in solution
```

## Dependencies

- `System.Reflection` — Standard .NET library

## Side Effects

- Sets version metadata on all compiled DLLs
- Used by update checker to determine current version

## Reference

- Included as a linked file in all `.csproj` files
