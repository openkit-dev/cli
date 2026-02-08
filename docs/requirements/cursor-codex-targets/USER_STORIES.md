# User Stories: Cursor + Codex Targets

## Cursor Target

### US-CUR-01: Install Cursor Rules
**As a** developer using Cursor
**I want** OpenKit to install modular rules in `.cursor/rules/`
**So that** I get consistent guidance without a monolithic `.cursorrules` file

**Acceptance Criteria:**
- `openkit cursor sync` creates `.cursor/rules/openkit.mdc`
- Rules reference skills under `.cursor/skills/` (if installed)
- `.cursorrules` still generated for legacy compatibility

### US-CUR-02: Cursor Skills Reference
**As a** developer using Cursor
**I want** OpenKit to install skills in `.cursor/skills/`
**So that** I can reference them in my Cursor conversations

**Acceptance Criteria:**
- `openkit cursor sync` copies skills to `.cursor/skills/<name>/SKILL.md`
- Rules file mentions available skills

### US-CUR-03: Cursor Doctor Diagnostics
**As a** developer using Cursor
**I want** `openkit cursor doctor` to check my setup
**So that** I know if my configuration is healthy

**Acceptance Criteria:**
- Doctor checks for `.cursorrules` OR `.cursor/rules/`
- Reports managed vs unmanaged state
- Suggests running sync if files are missing

---

## Codex Target

### US-CDX-01: Rich AGENTS.md
**As a** developer using OpenAI Codex CLI
**I want** OpenKit to install a comprehensive AGENTS.md
**So that** Codex understands my project conventions

**Acceptance Criteria:**
- `openkit codex sync` creates `AGENTS.md` with:
  - Project expectations section
  - Skills reference section
  - SDD workflow summary
- File follows Codex discovery conventions

### US-CDX-02: Codex Rules File
**As a** developer using Codex CLI
**I want** OpenKit to install sandbox rules in `.codex/rules/`
**So that** common development commands are pre-approved

**Acceptance Criteria:**
- `openkit codex sync` creates `.codex/rules/openkit.rules`
- Rules allow common safe commands (npm, git status, etc.)
- Uses Starlark prefix_rule() syntax

### US-CDX-03: Codex Skills Installation
**As a** developer using Codex CLI
**I want** OpenKit to install skills in `.agents/skills/`
**So that** Codex can reference them during work

**Acceptance Criteria:**
- Skills copied to `.agents/skills/<name>/SKILL.md`
- AGENTS.md references the skills path

### US-CDX-04: Codex Doctor Diagnostics
**As a** developer using Codex CLI
**I want** `openkit codex doctor` to check my setup
**So that** I know if AGENTS.md and rules are healthy

**Acceptance Criteria:**
- Doctor checks for `AGENTS.md`
- Doctor checks for `.codex/rules/` if rules installed
- Reports managed file count and drift status
