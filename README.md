# OpenKit CLI

Universal Spec-Driven Development toolkit for AI coding agents.

## Features

- **Multi-Agent Support**: Works with OpenCode, Claude Code, Cursor, Gemini CLI, Codex, and Windsurf
- **SDD Workflow**: Embedded templates, commands, skills, and prompts for Spec-Driven Development
- **Cross-Platform**: Runs on macOS, Linux, and Windows
- **No Dependencies**: Single binary, no runtime required

## Installation

### macOS / Linux / WSL

```bash
curl -fsSL https://openkit.dev/install | bash
```

### Windows (PowerShell)

```powershell
irm https://openkit.dev/install.ps1 | iex
```

### From Source

```bash
go install github.com/openkit-dev/cli/cmd/openkit@latest
```

## Quick Start

```bash
# Check system requirements
openkit check

# Create a new project
openkit init my-app --ai opencode

# Or for Claude Code
openkit init my-app --ai claude

# Initialize in current directory
openkit init --here
```

## Commands

| Command | Description |
|---------|-------------|
| `openkit init <name>` | Initialize a new project with SDD templates |
| `openkit check` | Check system requirements and installed agents |
| `openkit version` | Print version information |

### Init Flags

| Flag | Description |
|------|-------------|
| `--ai <agent>` | AI agent to configure (opencode, claude, cursor, gemini) |
| `--here` | Initialize in current directory |
| `--force` | Overwrite existing files |
| `--no-git` | Skip git initialization |

## Supported Agents

| Agent | Folder | Status |
|-------|--------|--------|
| OpenCode | `.opencode/` | Supported |
| Claude Code | `.claude/` + `CLAUDE.md` | Supported |
| Cursor | `.cursor/` | Supported |
| Gemini CLI | `.gemini/` | Supported |
| Codex CLI | `.codex/` | Planned |
| Windsurf | `.windsurf/` | Planned |

## SDD Workflow

OpenKit implements Spec-Driven Development with these commands:

1. `/specify` - Create feature specification
2. `/clarify` - Resolve spec ambiguities
3. `/plan` - Create implementation plan
4. `/tasks` - Generate executable tasks
5. `/impl` - Execute implementation
6. `/test` - Generate or run tests

## Development

```bash
# Clone the repository
git clone https://github.com/openkit-dev/cli.git
cd cli

# Build
make build

# Run tests
make test

# Build for all platforms
make build-all
```

## License

MIT
