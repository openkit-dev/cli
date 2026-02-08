# Changelog

All notable changes to OpenKit CLI will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.2.0] - 2026-02-08

### Added

#### Cursor Target Enhancements
- Enhanced `.cursorrules` with comprehensive SDD workflow guidance
- Added `.cursor/rules/openkit.mdc` - modular rules with YAML frontmatter
- Added `.cursor/skills/**` - complete skills library (145+ files)
- Improved `openkit cursor doctor` with checks for `.cursor/rules/` and `.cursor/skills/`

#### Codex Target Enhancements
- Added comprehensive `AGENTS.md` (< 32KB limit) with complete SDD guide
- Added `.codex/rules/openkit.rules` - Starlark command policies for safe automation
- Added `.agents/skills/**` - complete skills library (145+ files)
- Improved `openkit codex doctor` with checks for `.codex/rules/` and `.agents/skills/`

#### Testing
- Added 22 unit tests for Cursor and Codex content generators
- Added 2 integration tests for Cursor and Codex sync commands
- All tests passing with 100% success rate

#### Documentation
- Expanded README from 163 to 532 lines with comprehensive multi-agent guide
- Added complete agent-specific setup guides for 5 agents
- Added workflow examples with practical commands
- Added managed state & safety documentation
- Added upgrade & migration procedures
- Added per-agent documentation in `docs/agent-compat/agents/`

### Changed
- Cursor target now generates 147 files (was ~10)
- Codex target now generates 147 files (was ~10)
- Doctor commands provide more detailed diagnostics

### Fixed
- Corrected archive filename in install scripts (`cli_` instead of `openkit_`)
- Updated installation instructions with correct GitHub URLs

## [0.1.0] - 2026-02-07

### Added
- Initial release
- Multi-agent support: OpenCode, Claude Code, Cursor, Gemini CLI, Codex
- Safe-by-default sync engine with managed state tracking
- Conflict detection and drift detection
- `openkit init` command for project scaffolding
- `openkit check` command for system verification
- `openkit <agent> sync` commands for configuration installation
- `openkit <agent> doctor` commands for health checks
- `openkit <agent> upgrade` commands for safe updates
- Gemini target with TOML command generation
- Complete skills library (33+ domain skills)
- Embedded templates for all agents
- Cross-platform support (macOS, Linux, Windows)

[0.2.0]: https://github.com/openkit-devtools/openkit/compare/v0.1.0...v0.2.0
[0.1.0]: https://github.com/openkit-devtools/openkit/releases/tag/v0.1.0
