# OpenKit CLI

> Universal Spec-Driven Development toolkit for AI coding agents.

Configure a multi-agent AI development environment with specialized agents, 33+ domain skills, and 7 development commands.

## What is OpenKit?

OpenKit is a **CLI toolkit** that configures **Spec-Driven Development** environments for multiple AI coding agents:

- **Multi-Agent Support**: OpenCode, Claude Code, Cursor, Gemini CLI, Codex, Windsurf
- **33+ Domain Skills**: Frontend, backend, security, testing, architecture
- **7 Commands**: Slash commands for the complete development workflow
- **Obsidian-Native Docs Graph**: Canonical `HUB-*.md` structure with wikilinks and connected planning artifacts
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

- **macOS (Intel):** `openkit_Darwin_x86_64.tar.gz`
- **macOS (Apple Silicon):** `openkit_Darwin_arm64.tar.gz`
- **Linux (x64):** `openkit_Linux_x86_64.tar.gz`
- **Linux (ARM64):** `openkit_Linux_arm64.tar.gz`
- **Windows:** `openkit_Windows_x86_64.zip`

Extract and move to your PATH:
```bash
tar -xzf openkit_*.tar.gz
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

### Obsidian Vault Setup (Recommended)

OpenKit documentation is designed for Obsidian-compatible wikilinks and hub navigation.

- Open your vault at `./docs` (the `docs` folder itself)
- Do not open the repository root as vault if your docs links are normalized for `docs` root
- Start from `HUB-DOCS.md`, then navigate to `requirements/HUB-REQUIREMENTS.md` and `sprint/HUB-SPRINTS.md`

This avoids path mismatches like accidental `docs/docs/...` notes when creating links from Obsidian.

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

### Memory Kernel Commands

Manage docs-first project memory using the new runtime bridge:

```bash
openkit memory init                 # Initialize memory contracts and directories
openkit memory doctor --json --write
openkit memory capture --session-id s01 --summary "Sprint work" --action create
openkit memory review --json
```

Notes:
- `openkit memory` delegates to the Rust runtime (`openkit-rs`).
- Runtime resolution order: explicit `OPENKIT_MEMORY_RUNTIME_PATH`, sidecar binary next to `openkit`, `openkit-rs` on PATH, then cargo fallback in `rust-cli/`.
- You can force runtime mode with `OPENKIT_MEMORY_RUNTIME=binary|cargo`.
- Release installers attempt to install `openkit-rs` sidecar when an official asset exists for your platform.

## Supported Agents

### OpenCode

[OpenCode](https://github.com/stackblitz-labs/opencode) - Terminal-based AI coding agent

**Installation:**
```bash
npm i -g @opencode/cli
```

**What OpenKit Installs:**
- `opencode.json` - Agent configuration
- `.opencode/commands/` - 7 slash commands
- `.opencode/prompts/` - Specialized agent prompts
- `.opencode/rules/` - Master ruleset
- `.opencode/skills/` - 33+ domain skills
- `.opencode/scripts/` - Verification scripts

**Usage:**
```bash
openkit opencode sync
opencode  # Start OpenCode in your project
```

**Permanent Memory (Docs-First):**

OpenKit is migrating from the legacy semantic memory plugin to a docs-first permanent memory model.

- Durable context in `docs/` (requirements, sprint, ADR, related links)
- Operational state in `.openkit/ops/`
- Memory contracts in `.openkit/memory/`

See `docs/MEMORY_LEGACY_MIGRATION.md` and `docs/DEPRECATIONS.md` for migration details.

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
- `.gemini/commands/openkit/*.toml` - 7 TOML commands
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

## 7 Development Commands

OpenKit provides **7 commands** for the complete development workflow:

| # | Command | Purpose |
|---|---------|---------|
| 1 | `/discover` | Analyze project context (mandatory) |
| 2 | `/specify` | Specification + Planning + Tasks |
| 3 | `/create` | Implementation from plan |
| 4 | `/verify` | Quality verification (tests, lint, security) |
| 5 | `/orchestrate` | Universal orchestrator for complex tasks |
| 6 | `/debug` | Systematic debugging |
| 7 | `/deploy` | Safe deployment |

---

## Standard Workflow

```
/discover → /specify → /create → /verify → /deploy
```

---

## Usage Examples

### Example 1: New Feature

```bash
# In OpenCode

# 1. Discover project context
/discover
# → Analyzes structure, identifies risks, creates CONTEXT.md

# 2. Specify and plan the feature
/specify add user authentication
# → Creates PROBLEM_STATEMENT.md, USER_STORIES.md, ACCEPTANCE_CRITERIA.md
# → Creates PLAN.md, SPRINT_GOAL.md, BACKLOG.md
# → Creates TASKS.md with task breakdown

# 3. Implement
/create from docs/sprint/Sprint-1/TASKS.md
# → Executes P0: Foundation (DB + Security)
# → Executes P1: Backend
# → Executes P2: UI/UX
# → Executes P3: Polish

# 4. Verify
/verify all
# → Runs lint + type check
# → Runs security scan
# → Runs unit tests
# → Runs UX audit
# → Runs performance checks

# 5. Deploy
/deploy staging
# → Prepares deployment
# → Executes deploy
# → Verifies post-deploy
```

### Example 2: Bug Fix

```bash
# Debug the issue
/debug login not working after update
# → Phase 1: Symptom Analysis
# → Phase 2: Information Gathering
# → Phase 3: Hypothesis Testing
# → Phase 4: Resolution

# Verify the fix
/verify quick
```

### Example 3: Complex Project

```bash
# Use orchestrator for complex tasks
/orchestrate build e-commerce with Stripe checkout
# → Automatically runs /discover
# → Automatically runs /specify
# → Coordinates multiple agents (database-architect, backend-specialist, frontend-specialist, security-auditor)
# → Runs /verify
# → Reports results
```

---

## Command Details

### /discover
Analyzes project structure and generates context documentation.
- Creates: `docs/CONTEXT.md`, `docs/SECURITY.md`, `docs/QUALITY_GATES.md`
- **Mandatory** before `/specify`

### /specify
Complete specification + planning + task breakdown.
- Creates: Requirements, Plan, Sprint goals, Task breakdown
- Unifies: old commands (/specify, /clarify, /plan, /tasks)

### /create
Executes implementation from specification.
- Reads: `docs/sprint/Sprint-XX/TASKS.md`
- Executes: P0→P1→P2→P3 phases with STOP points

### /verify
Quality verification suite.
- Runs: Lint, Security scan, Tests, UX audit, Performance
- Unifies: old commands (/test, /checklist, /analyze)

### /orchestrate
Universal orchestrator for complex multi-domain tasks.
- Coordinates: Multiple specialist agents
- Automates: Entire workflow

### /debug
4-phase systematic debugging protocol.
- Phase 1: Symptom Analysis
- Phase 2: Information Gathering
- Phase 3: Hypothesis Testing
- Phase 4: Resolution

### /deploy
Safe deployment with verification.
- Environments: Staging, Production, Preview
- Includes: Pre-deploy checklist, health checks, rollback

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

### CLI Self-Upgrade

```bash
# Check latest version and artifact URLs (no install)
openkit upgrade --check

# (alias)
openkit upgrade --dry-run

# Download + verify checksums + install
openkit upgrade
```

Note: `openkit upgrade --dry-run` checks the CLI binary update; `openkit <agent> upgrade --dry-run` previews config changes.

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

### Uninstall CLI Binary

```bash
# Interactive (asks for confirmation)
curl -fsSL https://raw.githubusercontent.com/openkit-devtools/openkit/main/scripts/uninstall.sh | bash

# Non-interactive
curl -fsSL https://raw.githubusercontent.com/openkit-devtools/openkit/main/scripts/uninstall.sh | bash -s -- -y
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
