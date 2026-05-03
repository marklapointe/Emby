# Component: Project Artifacts & Configuration

**Path:** \`./\`
**Type:** Configuration | Artifacts
**Maps to:** \`.discovery/950-project-artifacts.md\`

## Description

Project-level configuration files, IDE settings, planning documents, and build artifacts.

## Files

### Git Configuration

- \`.gitignore\` — Git ignore rules for build outputs, IDE files, packages, etc.

### Visual Studio

- \`.vs/\` — Visual Studio 2017 settings and caches
  - \`MediaBrowser/v15/Server/sqlite3\` — VS server database

### Planning Documents

- \`.plan/\` — Migration planning documents
  - \`csharp-to-go-migration-plan.md\` — C# to Go migration strategy
- \`.plan.d/\` — Detailed planning documents
  - \`000.0-Migration-TOC.md\` — Migration table of contents
  - \`000.1-Agent-Workflow.md\` — Agent workflow specification
  - \`001.0-Overview.md\` — Migration overview
  - \`001.1-Implementation-Phases.md\` — Implementation phases
  - \`001.2-CSharp-Analysis.md\` — C# codebase analysis
  - \`002.0-API-Specification.md\` — API specification
  - \`003.0-Implementation-Details.md\` — Implementation details
  - \`004.0-Testing.md\` — Testing strategy
  - \`005.0-Risks-TODO.md\` — Risks and TODO items

### Build Artifacts

- \`.output.txt\` — Build/test output log
