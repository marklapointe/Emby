# Component: Emby.Server.Implementations — Text Encoding & Localization

**Path:** \`Emby.Server.Implementations/\`
**Type:** Directory | Module Group
**Language:** C#
**Maps to:** \`.discovery/167-emby-server-impl-encoding.md\`

## Description

Text encoding detection, localization, and internationalization support including embedded UniversalDetector and NLangDetect libraries.

## Files

### Localization/

- `LocalizationManager.cs` — Emby.Server.Implementations/Localization/LocalizationManager.cs
- `TextLocalizer.cs` — Emby.Server.Implementations/Localization/TextLocalizer.cs

### TextEncoding/

- `Detector.cs` — Emby.Server.Implementations/TextEncoding/NLangDetect/Detector.cs
- `DetectorFactory.cs` — Emby.Server.Implementations/TextEncoding/NLangDetect/DetectorFactory.cs
- `ErrorCode.cs` — Emby.Server.Implementations/TextEncoding/NLangDetect/ErrorCode.cs
- `CharExtensions.cs` — Emby.Server.Implementations/TextEncoding/NLangDetect/Extensions/CharExtensions.cs
- `RandomExtensions.cs` — Emby.Server.Implementations/TextEncoding/NLangDetect/Extensions/RandomExtensions.cs
- `StringExtensions.cs` — Emby.Server.Implementations/TextEncoding/NLangDetect/Extensions/StringExtensions.cs
- `UnicodeBlock.cs` — Emby.Server.Implementations/TextEncoding/NLangDetect/Extensions/UnicodeBlock.cs
- `GenProfile.cs` — Emby.Server.Implementations/TextEncoding/NLangDetect/GenProfile.cs
- `InternalException.cs` — Emby.Server.Implementations/TextEncoding/NLangDetect/InternalException.cs
- `Language.cs` — Emby.Server.Implementations/TextEncoding/NLangDetect/Language.cs
- `LanguageDetector.cs` — Emby.Server.Implementations/TextEncoding/NLangDetect/LanguageDetector.cs
- `NLangDetectException.cs` — Emby.Server.Implementations/TextEncoding/NLangDetect/NLangDetectException.cs
- `ProbVector.cs` — Emby.Server.Implementations/TextEncoding/NLangDetect/ProbVector.cs
- `LangProfile.cs` — Emby.Server.Implementations/TextEncoding/NLangDetect/Utils/LangProfile.cs
- `Messages.cs` — Emby.Server.Implementations/TextEncoding/NLangDetect/Utils/Messages.cs
- `NGram.cs` — Emby.Server.Implementations/TextEncoding/NLangDetect/Utils/NGram.cs
- `TagExtractor.cs` — Emby.Server.Implementations/TextEncoding/NLangDetect/Utils/TagExtractor.cs
- `TextEncoding.cs` — Emby.Server.Implementations/TextEncoding/TextEncoding.cs
- `TextEncodingDetect.cs` — Emby.Server.Implementations/TextEncoding/TextEncodingDetect.cs
- `CharsetDetector.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/CharsetDetector.cs
- `Big5Prober.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/Big5Prober.cs
- `BitPackage.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/BitPackage.cs
- `CharDistributionAnalyser.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/CharDistributionAnalyser.cs
- `CharsetProber.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/CharsetProber.cs
- `Charsets.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/Charsets.cs
- `CodingStateMachine.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/CodingStateMachine.cs
- `EscCharsetProber.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/EscCharsetProber.cs
- `EscSM.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/EscSM.cs
- `EUCJPProber.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/EUCJPProber.cs
- `EUCKRProber.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/EUCKRProber.cs
- `EUCTWProber.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/EUCTWProber.cs
- `GB18030Prober.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/GB18030Prober.cs
- `HebrewProber.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/HebrewProber.cs
- `JapaneseContextAnalyser.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/JapaneseContextAnalyser.cs
- `LangBulgarianModel.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/LangBulgarianModel.cs
- `LangCyrillicModel.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/LangCyrillicModel.cs
- `LangGreekModel.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/LangGreekModel.cs
- `LangHebrewModel.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/LangHebrewModel.cs
- `LangHungarianModel.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/LangHungarianModel.cs
- `LangThaiModel.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/LangThaiModel.cs
- `Latin1Prober.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/Latin1Prober.cs
- `MBCSGroupProber.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/MBCSGroupProber.cs
- `MBCSSM.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/MBCSSM.cs
- `SBCharsetProber.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/SBCharsetProber.cs
- `SBCSGroupProber.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/SBCSGroupProber.cs
- `SequenceModel.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/SequenceModel.cs
- `SJISProber.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/SJISProber.cs
- `SMModel.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/SMModel.cs
- `UniversalDetector.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/UniversalDetector.cs
- `UTF8Prober.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/UTF8Prober.cs
- `DetectionConfidence.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/DetectionConfidence.cs
- `ICharsetDetector.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/ICharsetDetector.cs

### TextEncoding/NLangDetect/

- `Detector.cs` — Emby.Server.Implementations/TextEncoding/NLangDetect/Detector.cs
- `DetectorFactory.cs` — Emby.Server.Implementations/TextEncoding/NLangDetect/DetectorFactory.cs
- `ErrorCode.cs` — Emby.Server.Implementations/TextEncoding/NLangDetect/ErrorCode.cs
- `CharExtensions.cs` — Emby.Server.Implementations/TextEncoding/NLangDetect/Extensions/CharExtensions.cs
- `RandomExtensions.cs` — Emby.Server.Implementations/TextEncoding/NLangDetect/Extensions/RandomExtensions.cs
- `StringExtensions.cs` — Emby.Server.Implementations/TextEncoding/NLangDetect/Extensions/StringExtensions.cs
- `UnicodeBlock.cs` — Emby.Server.Implementations/TextEncoding/NLangDetect/Extensions/UnicodeBlock.cs
- `GenProfile.cs` — Emby.Server.Implementations/TextEncoding/NLangDetect/GenProfile.cs
- `InternalException.cs` — Emby.Server.Implementations/TextEncoding/NLangDetect/InternalException.cs
- `Language.cs` — Emby.Server.Implementations/TextEncoding/NLangDetect/Language.cs
- `LanguageDetector.cs` — Emby.Server.Implementations/TextEncoding/NLangDetect/LanguageDetector.cs
- `NLangDetectException.cs` — Emby.Server.Implementations/TextEncoding/NLangDetect/NLangDetectException.cs
- `ProbVector.cs` — Emby.Server.Implementations/TextEncoding/NLangDetect/ProbVector.cs
- `LangProfile.cs` — Emby.Server.Implementations/TextEncoding/NLangDetect/Utils/LangProfile.cs
- `Messages.cs` — Emby.Server.Implementations/TextEncoding/NLangDetect/Utils/Messages.cs
- `NGram.cs` — Emby.Server.Implementations/TextEncoding/NLangDetect/Utils/NGram.cs
- `TagExtractor.cs` — Emby.Server.Implementations/TextEncoding/NLangDetect/Utils/TagExtractor.cs

### TextEncoding/NLangDetect/Extensions/

- `CharExtensions.cs` — Emby.Server.Implementations/TextEncoding/NLangDetect/Extensions/CharExtensions.cs
- `RandomExtensions.cs` — Emby.Server.Implementations/TextEncoding/NLangDetect/Extensions/RandomExtensions.cs
- `StringExtensions.cs` — Emby.Server.Implementations/TextEncoding/NLangDetect/Extensions/StringExtensions.cs
- `UnicodeBlock.cs` — Emby.Server.Implementations/TextEncoding/NLangDetect/Extensions/UnicodeBlock.cs

### TextEncoding/NLangDetect/Profiles/


### TextEncoding/NLangDetect/Utils/

- `LangProfile.cs` — Emby.Server.Implementations/TextEncoding/NLangDetect/Utils/LangProfile.cs
- `Messages.cs` — Emby.Server.Implementations/TextEncoding/NLangDetect/Utils/Messages.cs
- `NGram.cs` — Emby.Server.Implementations/TextEncoding/NLangDetect/Utils/NGram.cs
- `TagExtractor.cs` — Emby.Server.Implementations/TextEncoding/NLangDetect/Utils/TagExtractor.cs

### TextEncoding/UniversalDetector/

- `CharsetDetector.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/CharsetDetector.cs
- `Big5Prober.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/Big5Prober.cs
- `BitPackage.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/BitPackage.cs
- `CharDistributionAnalyser.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/CharDistributionAnalyser.cs
- `CharsetProber.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/CharsetProber.cs
- `Charsets.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/Charsets.cs
- `CodingStateMachine.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/CodingStateMachine.cs
- `EscCharsetProber.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/EscCharsetProber.cs
- `EscSM.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/EscSM.cs
- `EUCJPProber.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/EUCJPProber.cs
- `EUCKRProber.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/EUCKRProber.cs
- `EUCTWProber.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/EUCTWProber.cs
- `GB18030Prober.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/GB18030Prober.cs
- `HebrewProber.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/HebrewProber.cs
- `JapaneseContextAnalyser.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/JapaneseContextAnalyser.cs
- `LangBulgarianModel.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/LangBulgarianModel.cs
- `LangCyrillicModel.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/LangCyrillicModel.cs
- `LangGreekModel.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/LangGreekModel.cs
- `LangHebrewModel.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/LangHebrewModel.cs
- `LangHungarianModel.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/LangHungarianModel.cs
- `LangThaiModel.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/LangThaiModel.cs
- `Latin1Prober.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/Latin1Prober.cs
- `MBCSGroupProber.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/MBCSGroupProber.cs
- `MBCSSM.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/MBCSSM.cs
- `SBCharsetProber.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/SBCharsetProber.cs
- `SBCSGroupProber.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/SBCSGroupProber.cs
- `SequenceModel.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/SequenceModel.cs
- `SJISProber.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/SJISProber.cs
- `SMModel.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/SMModel.cs
- `UniversalDetector.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/UniversalDetector.cs
- `UTF8Prober.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/UTF8Prober.cs
- `DetectionConfidence.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/DetectionConfidence.cs
- `ICharsetDetector.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/ICharsetDetector.cs

### TextEncoding/UniversalDetector/Core/

- `Big5Prober.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/Big5Prober.cs
- `BitPackage.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/BitPackage.cs
- `CharDistributionAnalyser.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/CharDistributionAnalyser.cs
- `CharsetProber.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/CharsetProber.cs
- `Charsets.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/Charsets.cs
- `CodingStateMachine.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/CodingStateMachine.cs
- `EscCharsetProber.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/EscCharsetProber.cs
- `EscSM.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/EscSM.cs
- `EUCJPProber.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/EUCJPProber.cs
- `EUCKRProber.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/EUCKRProber.cs
- `EUCTWProber.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/EUCTWProber.cs
- `GB18030Prober.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/GB18030Prober.cs
- `HebrewProber.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/HebrewProber.cs
- `JapaneseContextAnalyser.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/JapaneseContextAnalyser.cs
- `LangBulgarianModel.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/LangBulgarianModel.cs
- `LangCyrillicModel.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/LangCyrillicModel.cs
- `LangGreekModel.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/LangGreekModel.cs
- `LangHebrewModel.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/LangHebrewModel.cs
- `LangHungarianModel.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/LangHungarianModel.cs
- `LangThaiModel.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/LangThaiModel.cs
- `Latin1Prober.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/Latin1Prober.cs
- `MBCSGroupProber.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/MBCSGroupProber.cs
- `MBCSSM.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/MBCSSM.cs
- `SBCharsetProber.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/SBCharsetProber.cs
- `SBCSGroupProber.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/SBCSGroupProber.cs
- `SequenceModel.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/SequenceModel.cs
- `SJISProber.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/SJISProber.cs
- `SMModel.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/SMModel.cs
- `UniversalDetector.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/UniversalDetector.cs
- `UTF8Prober.cs` — Emby.Server.Implementations/TextEncoding/UniversalDetector/Core/UTF8Prober.cs

## Decomposition

### LocalizationManager.cs (Localization Manager)

#### Imports
```csharp
using MediaBrowser.Model.Globalization;
using System;
using System.Collections.Generic;
using System.Globalization;
```

#### Classes
`LocalizationManager` (public class : ILocalizationManager)

#### Key Properties
| Property | Type | Description |
|----------|------|-------------|
| `CurrentCulture` | `CultureInfo` | Current culture |

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `GetLocalizedString(string)` | `string` | Get localized string |
| `GetCultures()` | `IEnumerable<CultureDto>` | Get available cultures |

### TextEncodingDetect.cs (Encoding Detection)

#### Classes
`TextEncodingDetect` (public static class)

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `DetectEncoding(byte[])` | `Encoding` | Detect encoding from bytes |
| `DetectEncoding(string)` | `Encoding` | Detect encoding from string |

### UniversalDetector.cs (Charset Detection - Mozilla)

#### Classes
`UniversalDetector` (public class)

#### Key Properties
| Property | Type | Description |
|----------|------|-------------|
| `DetectedCharset` | `string` | Detected charset name |
| `DetectedLanguage` | `string` | Detected language |

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `DoIt(byte[], int, bool)` | `void` | Process data |
| `Done()` | `void` | Finish detection |

### LanguageDetector.cs (Language Detection - NLangDetect)

#### Classes
`LanguageDetector` (public class)

#### Key Properties
| Property | Type | Description |
|----------|------|-------------|
| `Probability` | `double` | Detection confidence |

#### Key Methods
| Method | Return | Description |
|--------|--------|-------------|
| `Detect(string)` | `string` | Detect language |
| `DetectLangs(string)` | `IEnumerable<Language>` | Get language probabilities |

