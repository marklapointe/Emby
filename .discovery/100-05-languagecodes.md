# Component: LanguageCodes.cs

**Path:** `BDInfo/LanguageCodes.cs`
**Type:** File | Static Lookup
**Language:** C#
**Maps to:** `.discovery/100-05-languagecodes.md`
**Parent:** `.discovery/100-bdinfo.md`

## Structure

```
LanguageCodes.cs
├── [namespace] BDInfo
│   └── [class] public static class LanguageCodes
│       ├── [field] private static readonly Dictionary<string, string> _lookup
│       │   └── ISO 639-2/B language code → English name mapping
│       ├── [static constructor] LanguageCodes()
│       │   └── Populates _lookup with 400+ language codes
│       │   └── Examples: "eng" → "English", "fra" → "French", "jpn" → "Japanese"
│       └── [method] public static string GetName(string code)
│           └── Returns English name for ISO 639-2/B code
│           └── Returns code itself if not found
```

## Description

`LanguageCodes` is a static lookup table mapping ISO 639-2/B three-letter language codes (e.g., "eng", "fra", "jpn") to their English names. Used by BDInfo to translate language codes found in Blu-ray stream metadata into human-readable labels.
