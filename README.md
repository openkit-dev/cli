# OpenKit CLI

> Universal Spec-Driven Development toolkit for AI coding agents.

Configure a multi-agent AI development environment with specialized agents, 33+ domain skills, and 18 development commands.

## What is OpenKit?

OpenKit is a **CLI toolkit** that configures **Spec-Driven Development** environments for multiple AI coding agents:

- **Multi-Agent Support**: OpenCode, Claude Code, Cursor, Gemini CLI, Codex, Windsurf
- **33+ Domain Skills**: Frontend, backend, security, testing, architecture
- **18 Commands**: Slash commands for orchestrated workflows
- **Safe-by-Default Sync**: Managed state tracking with conflict detection
- **Cross-Platform**: Single binary, runs on macOS, Linux, Windows
- **No Dependencies**: No runtime required, no npm packages

## How It Works

1. **Install OpenKit**: Download CLI binary or use install script
2. **Sync for Your Agent**: `openkit <agent> sync` installs agent-specific configuration
3. **Development**: Use your AI agent with OpenKit commands and skills
4. **Upgrade**: `openkit <agent> upgrade` safely updates configuration

## Installation

### macOS / Linux / WSL

```bash
curl -fsSL https://raw.githubusercontent.com/openkit-devtools/openkit/main/scripts/install.sh | bash
```

### Windows (PowerShell)

```powershell
irm https://raw.githubusercontent.com/openkit-devtools/openkit/main/scripts/install.ps1 | iex
```

### Manual Download

Download the binary for your platform from the [latest release](https://github.com/openkit-devtools/openkit/releases/latest):

- **macOS (Intel):** `cli_Darwin_x86_64.tar.gz`
- **macOS (Apple Silicon):** `cli_Darwin_arm64.tar.gz`
- **Linux (x64):** `cli_Linux_x86_64.tar.gz`
- **Linux (ARM64):** `cli_Linux_arm64.tar.gz`
- **Windows:** `cli_Windows_x86_64.zip`

Note: current releases use `openkit_*` filenames.

Extract and move to your PATH:
```bash
tar -xzf cli_*.tar.gz
sudo mv openkit /usr/local/bin/
```

### From Source

```bash
go install github.com/openkit-devtools/openkit/cmd/openkit@latest
```

## Quick Start

### Option 1: New Project

```bash
# Create new project with OpenKit
openkit init my-app --ai opencode

# Navigate to project
cd my-app

# Start developing with OpenCode
opencode
```

### Option 2: Existing Project

```bash
# Navigate to your project
cd your-project

# Sync OpenKit for your agent
openkit opencode sync

# Start developing
opencode
```

### Option 3: Check Available Agents

```bash
# See which agents are installed on your system
openkit check
```

## CLI Commands

### Project Management

| Command | Description |
|---------|-------------|
| `openkit init <name>` | Initialize a new project with SDD templates |
| `openkit check` | Check system requirements and installed agents |
| `openkit version` | Print version information |

**Init Flags:**
- `--ai <agent>` - AI agent to configure (opencode, claude, cursor, gemini, codex)
- `--here` - Initialize in current directory
- `--force` - Overwrite existing files
- `--no-git` - Skip git initialization

### Agent-Specific Commands

Each agent has dedicated commands for configuration management:

```bash
openkit <agent> sync      # Install/update OpenKit configuration
openkit <agent> upgrade   # Upgrade to latest version
openkit <agent> doctor    # Check configuration health
```

**Sync Flags:**
- `--dry-run` - Preview changes without writing
- `--overwrite` - Overwrite unmanaged or drifted files
- `--prune` - Remove managed files no longer in the plan

## Supported Agents

### OpenCode

[OpenCode](https://github.com/stackblitz-labs/opencode) - Terminal-based AI coding agent

**Installation:**
```bash
npm i -g @opencode/cli
```

**What OpenKit Installs:**
- `opencode.json` - Agent configuration
- `.opencode/commands/` - 18 slash commands
- `.opencode/prompts/` - Specialized agent prompts
- `.opencode/rules/` - Master ruleset
- `.opencode/skills/` - 33+ domain skills
- `.opencode/scripts/` - Verification scripts

**Usage:**
```bash
openkit opencode sync
opencode  # Start OpenCode in your project
```

---

### Claude Code

[Claude Code](https://docs.anthropic.com/claude/docs/claude-code) - Official Anthropic AI agent

**What OpenKit Installs:**
- `.claude/CLAUDE.md` - Main instruction file
- `.claude/settings.json` - Project settings
- `.claude/rules/` - Universal rules
- `.claude/skills/` - Domain skills
- `.claude/agents/` - Specialized prompts

**Usage:**
```bash
openkit claude sync
# Use Claude Code extension in your IDE
```

---

### Cursor

[Cursor](https://cursor.sh) - AI-first code editor

**What OpenKit Installs:**
- `.cursorrules` - Project rules (legacy format)
- `.cursor/rules/openkit.mdc` - Modular rules with frontmatter
- `.cursor/skills/` - Domain skills

**Usage:**
```bash
openkit cursor sync
# Open project in Cursor IDE
```

---

### Gemini CLI

[Gemini CLI](https://ai.google.dev/gemini-api/docs/cli) - Google's AI coding agent

**What OpenKit Installs:**
- `GEMINI.md` - Main instruction file
- `.gemini/settings.json` - Agent settings
- `.gemini/commands/openkit/*.toml` - 18 TOML commands
- `.gemini/rules/` - Universal rules
- `.gemini/skills/` - Domain skills

**Usage:**
```bash
openkit gemini sync
gemini  # Start Gemini CLI in your project
```

---

### Codex CLI

[Codex CLI](https://github.com/openai/codex) - OpenAI's terminal coding agent

**What OpenKit Installs:**
- `AGENTS.md` - Comprehensive agent configuration
- `.codex/rules/openkit.rules` - Starlark command policies
- `.agents/skills/` - Domain skills

**Usage:**
```bash
openkit codex sync
codex  # Start Codex CLI in your project
```

---

### Status Summary

| Agent | Status | Files Installed |
|-------|--------|-----------------|
| OpenCode | ✅ Supported | 150+ files |
| Claude Code | ✅ Supported | 145+ files |
| Cursor | ✅ Supported | 147+ files |
| Gemini CLI | ✅ Supported | 171+ files |
| Codex CLI | ✅ Supported | 147+ files |
| Windsurf | 🚧 Planned | - |

## Development Workflow

OpenKit provides 18 commands for Spec-Driven Development:

### Core SDD Commands

| Command | Purpose |
|---------|---------|
| `/specify` | Create feature specification |
| `/clarify` | Resolve spec ambiguities |
| `/plan` | Create implementation plan |
| `/tasks` | Generate executable tasks |
| `/impl` | Execute implementation |
| `/test` | Run tests and quality checks |

### Specialized Commands

| Command | Purpose |
|---------|---------|
| `/engineer` | Orchestrate multi-domain tasks |
| `/debug` | Systematic root cause analysis |
| `/ui-ux` | Design system generation |
| `/deploy` | Safe deployment procedures |
| `/doc` | Documentation generation |
| `/context` | Generate codebase context |
| `/brainstorm` | Explore approaches |
| `/status` | View progress |
| `/preview` | Manage dev environment |
| `/analyze` | Validate consistency |
| `/checklist` | Readiness checks |
| `/create` | Bootstrap new apps |

### Standard Workflow (SDD)

Use this for most feature work:

```text
/specify → /clarify → /plan → /tasks → /impl → /test
```

**Key Properties:**
- Planning recorded in `docs/requirements/<feature>/`
- Tasks tracked in `docs/sprint/Sprint-XX/`
- STOP points require explicit approval

**Example:**
```bash
# In your AI agent (e.g., opencode)
/specify add user authentication
# Agent creates docs/requirements/user-auth/

/clarify
# Resolve ambiguities about OAuth vs JWT

/plan add user authentication
# Agent creates PLAN.md with phases

/tasks
# Agent generates docs/sprint/Sprint-XX/TASKS.md

/impl from docs/sprint/Sprint-XX/TASKS.md
# Agent executes tasks one by one

/test
# Run verification suite
```

### Orchestrated Workflow (/engineer)

Use when task spans multiple domains:

```text
/engineer <mission>
  Phase 1: Planning (project-planner) → STOP
  Phase 2: Implementation (parallel specialists) → STOP
  Phase X: Verification (scripts)
```

**Example:**
```bash
/engineer build secure e-commerce checkout with Stripe integration

# Phase 1: Planning
# - Creates specs, risks, plan
# - STOP for approval

# Phase 2: Implementation
# - database-architect: schema
# - backend-specialist: API
# - frontend-specialist: UI
# - security-auditor: PCI compliance
# - STOP for review

# Phase X: Verification
# - Runs security scan
# - Runs tests
# - Validates deployment
```

### Verification & Quality

OpenKit includes verification scripts (OpenCode target):

```bash
# Lint and type check
npm run lint && npx tsc --noEmit

# Security scan
python .opencode/scripts/security_scan.py .

# UX audit
python .opencode/scripts/ux_audit.py .

# Full suite (requires running server)
python .opencode/scripts/verify_all.py . --url http://localhost:3000

# E2E tests (requires server)
python .opencode/scripts/playwright_runner.py http://localhost:3000
```

## Domain Skills

OpenKit includes 33+ modular knowledge domains:

### Frontend & Design
- `frontend-design` - UI/UX engine with 50+ styles and 97 palettes
- `nextjs-react-expert` - React performance (Vercel best practices)
- `tailwind-patterns` - Tailwind v4 utilities
- `mobile-design` - iOS/Android patterns

### Backend & Data
- `python-patterns` - FastAPI, Pydantic, async/await
- `database-design` - Schema optimization, Alembic
- `api-patterns` - RESTful design, error handling

### Quality & Security
- `webapp-testing` - Playwright E2E automation
- `vulnerability-scanner` - Security auditing
- `clean-code` - Universal coding standards
- `testing-patterns` - Unit/integration/E2E strategies

### Architecture & Planning
- `architecture` - Decision-making framework
- `plan-writing` - Structured task planning
- `brainstorming` - Socratic questioning

### Operational
- `deployment-procedures` - Production deployment
- `server-management` - Process management
- `performance-profiling` - Optimization techniques

[See all 33+ skills →](docs/SKILLS.md)

## Managed State & Safety

OpenKit tracks all installed files in `.openkit/managed.json`:

```json
{
  "schema_version": 1,
  "agents": {
    "opencode": {
      "pack": {
        "id": "embedded",
        "version": "0.1.0"
      },
      "files": {
        "opencode.json": {
          "installed_sha256": "abc123...",
          "mode": "copy"
        }
      }
    }
  }
}
```

**Safety Features:**
- **Conflict Detection**: Warns about unmanaged files before overwriting
- **Drift Detection**: Detects manual changes to managed files
- **Backup**: Creates timestamped backups before overwriting
- **Idempotent**: Running sync twice produces no changes
- **Prune**: Safe removal of orphaned files with `--prune`

**Doctor Command:**
```bash
openkit opencode doctor

# Output:
# [OK] opencode.json
# [OK] .opencode/
# 
# Managed files: 150
# Drifted:       0
# Missing:       0
# Pack:          embedded@0.1.0
```

## Agent-Specific Guides

Each agent has different configuration formats:

### OpenCode
- Uses `opencode.json` for agent/tool configuration
- Markdown-based commands in `.opencode/commands/`
- See [docs/agent-compat/agents/opencode.md](docs/agent-compat/agents/opencode.md)

### Claude Code
- Uses `.claude/CLAUDE.md` as entrypoint
- Settings in `.claude/settings.json` (do NOT manage `settings.local.json`)
- See [docs/agent-compat/agents/claude.md](docs/agent-compat/agents/claude.md)

### Cursor
- Uses `.cursorrules` (legacy) + `.cursor/rules/*.mdc` (modern)
- Modular rules have YAML frontmatter
- See [docs/agent-compat/agents/cursor.md](docs/agent-compat/agents/cursor.md)

### Gemini CLI
- Uses `GEMINI.md` as entrypoint
- Commands are TOML files in `.gemini/commands/openkit/*.toml`
- May require repo to be trusted
- See [docs/agent-compat/agents/gemini.md](docs/agent-compat/agents/gemini.md)

### Codex CLI
- Uses `AGENTS.md` as entrypoint (max 32KB)
- Rules are Starlark in `.codex/rules/*.rules`
- Supports hierarchical discovery (project → user global)
- See [docs/agent-compat/agents/codex.md](docs/agent-compat/agents/codex.md)

## Upgrade & Migration

### Safe Upgrade

```bash
# Preview changes
openkit opencode upgrade --dry-run

# Apply upgrade (skip conflicts by default)
openkit opencode upgrade

# Force overwrite conflicts
openkit opencode upgrade --overwrite
```

### Uninstall

```bash
# Remove all managed files for an agent
openkit opencode uninstall

# Preview what would be removed
openkit opencode uninstall --dry-run
```

### Migration Between Agents

```bash
# Sync for new agent (safe, no conflicts)
openkit claude sync

# Remove old agent files
openkit opencode uninstall
```

## Contributing

OpenKit CLI is open source:

```bash
# Clone the repository
git clone https://github.com/openkit-devtools/openkit.git
cd openkit/cli

# Build
go build -o openkit ./cmd/openkit

# Run tests
go test ./...

# Build for all platforms
goreleaser build --snapshot --clean
```

## Documentation

- **[Agent Compatibility](docs/agent-compat/)** - Per-agent configuration guides
- **[Content Protocol](docs/content-protocol/)** - Canonical artifact mapping
- **[Architecture Decision Records](docs/adr/)** - Design decisions
- **[Requirements](docs/requirements/)** - Feature specifications
- **[Sprint Planning](docs/sprint/)** - Development sprints

## License

MIT
