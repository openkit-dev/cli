---
trigger: always_on
---

# MASTER RULESET — OPENCODE AGENT SYSTEM

> Single source of mandatory rules for all agents, skills, and commands.

---

## Tool Usage (MANDATORY)

**All agents, commands, and skills MUST follow:** `.opencode/rules/TOOL_USAGE.md`

This rule defines:
- Correct tool selection (glob vs grep vs read vs bash)
- File operation patterns (read -> edit/write)
- User interaction (question tool vs inline)
- Task tracking (todolist tools vs inline)
- Common mistakes and fixes

**Key principles:**
- Use `glob` for file patterns, NOT `bash find`
- Use `grep` for content search, NOT `bash grep`
- Use `read` for file reading, NOT `bash cat`
- Use `edit` for existing files, `write` for new files
- Use `question` tool for multiple options, NOT inline questions
- Use `todolist` tools for complex tasks (3+ steps), NOT inline tracking
- Use `workdir` parameter, NOT `cd && command`

---

## Obsidian Documentation Protocol (MANDATORY)

**All agents, commands, and skills MUST follow:** `.opencode/rules/OBSIDIAN_LINKING.md`
**Single-Level Hub Linking Policy:** `.opencode/rules/SINGLE_LEVEL_HUB_LINKS.md`
**Canonical docs filenames MUST follow:** `.opencode/rules/DOCS_FILE_GLOSSARY.md`

Key principles:
- Use wikilinks for internal documentation references (for example, `[[CONTEXT.md]]`)
- Build documentation as a graph with inbound and outbound links
- Add a `## Related` section in docs artifacts when applicable
- Keep links stable and explicit to improve retrieval and long-term memory
- Use canonical uppercase docs filenames from the glossary

---

## Context & Language

- Respond in the user's language; keep code/comments in English.
- Follow `@[skills/clean-code]`: SRP, DRY, KISS, functions <= 20 lines, max 3 args, use guard clauses.
- Before changing any file, understand dependencies via `CODEBASE.md` or local references.
- Philosophy `Read → Understand → Apply`: identify objective, principles, and differences from generic output.

## Response Style (MANDATORY)

- Be objective, technical, consistent, and direct.
- Avoid embellishment, filler, and marketing language.
- Prefer short, information-dense sentences.

## Emoji Policy (MANDATORY)

- Do not use emojis in assistant responses.
- Do not introduce emojis into documentation.
- Exception: only use emojis if the user explicitly requests them.

---

## Question Tool Protocol (MANDATORY)

### When to Use `question` Tool

**MUST use `question` tool (never inline questions) when:**

| Situation | Example |
|-----------|----------|
| User must choose from **multiple options** | "Which framework: FastAPI or Django?" |
| User needs to make a **decision** | "Deploy to staging or production?" |
| **Ambiguous instructions** need clarification | "What authentication method: JWT or OAuth?" |
| **STOP points** requiring explicit approval | "Proceed with implementation?" |
| **Multiple choices** with trade-offs | "Option A or Option B for caching?" |

Additional enforcement:
- If you present 2+ next-step options (e.g., "I can do A or B"), you MUST use the `question` tool.
- Do not present multiple options as plain text and wait for a free-form reply.

### When Inline Questions Are OK

Single yes/no confirmations where context is clear:
- "Delete this file?" (after showing what will be deleted)
- "Continue with the plan?" (after presenting it)
- "What would you like to do?" (has multiple options) -> MUST use `question`

### Question Format

Use `question` tool with proper structure:

```javascript
question({
  questions: [{
    question: "Clear, concise question text",
    header: "Short title (max 30 chars)",
    options: [
      { label: "Option 1", description: "Why choose this" },
      { label: "Option 2", description: "Alternative reason" }
    ],
    multiple: false // or true for multi-select
  }]
})
```

**This is a MANDATORY pattern. Agents, skills, and commands MUST follow it.**

---

## TodoList Protocol (MANDATORY)

**Complete protocol:** `.opencode/rules/TODOLIST_PROTOCOL.md`

### When to Use `todowrite`/`todoread` Tools

**MUST use todolist tools (never inline tracking) when:**

| Situation | Example |
|-----------|----------|
| **Multi-step tasks** (3+ steps) | "Build e-commerce checkout with 5+ components" |
| **Complex planning phases** | Sprint planning with 5+ stories |
| **Orchestration mode** | Coordinating 3+ agents (backend, frontend, security) |
| **Tracking execution** | Marking progress across agent invocations |
| **Parallel workflows** | Managing multiple simultaneous tasks with dependencies |
| **STOP points with multiple outcomes** | "P0 complete. Proceed to P1, P2, or stop?" |

CRITICAL: After creating todolist, MUST use `question` tool to request user approval before execution.

### When Inline Tracking Is OK

Single-task operations where progress is obvious:
- "Installing dependencies..." (immediate single step)
- "Running tests..." (one command, output shows result)
- "Building feature with 5 steps" -> MUST use todolist

### TodoList Lifecycle

**1. Create:**
- Use `todowrite` at START of complex tasks (3+ steps)
- Include: task id, description (NO emojis), status, priority

**2. Request Approval (MANDATORY):**
- Use `question` tool to ask user approval
- Wait for user response before executing

**3. Update:**
- Mark tasks as `in_progress` when starting
- Mark tasks as `completed` when done (NO emojis in content)
- Always `todoread` before updating to check current state

**4. Read:**
- Use `todoread` to check current state before any action
- Verify no duplicate or conflicting tasks exist

### TodoList Format

Use `todowrite` tool with proper structure:

```javascript
todowrite({
  todos: [
    {
      id: "task-1",
      content: "Brief task description (max 50 chars)",
      status: "pending",
      priority: "high"
    },
    {
      id: "task-2",
      content: "Another task description",
      status: "in_progress",
      priority: "medium"
    }
  ]
})
```

**Status values:** `pending`, `in_progress`, `completed`, `cancelled`  
**Priority values:** `high`, `medium`, `low`

CRITICAL RULES:
- NO emojis in task content
- MUST use `question` tool for approval after creating todolist
- Only ONE task `in_progress` at a time

**See complete protocol:** `.opencode/rules/TODOLIST_PROTOCOL.md`

**This is a MANDATORY pattern. Agents, skills, and commands MUST follow it.**

---

## Runtime Model Choice

Model selection is handled by runtime/user environment. Agents and prompts must not emit model recommendation suggestions during execution.

---

## Agent & Skill Protocol

1. Detect the domain (frontend, backend, mobile, security, etc.).
2. Read the corresponding prompt file (`.opencode/prompts/<agent>.md`).
3. Load all skills listed in the frontmatter (`SKILL.md`).
4. Announce: `Applying knowledge of @[agent-name]...` before any specialized output.
5. Follow `@[skills/intelligent-routing]` and keep the checklist:
   - Correct agent? → If not, re-evaluate.
   - Agent file read? → Open and review.
   - Skills loaded? → Read each relevant `SKILL.md`.
   - Announcement made? → Inform the user.
6. For any documentation generation/update, enforce Obsidian and filename protocols:
   - `.opencode/rules/OBSIDIAN_LINKING.md`
   - `.opencode/rules/DOCS_FILE_GLOSSARY.md`

### Socratic Gate
Always ask clarifying questions before execution when there is significant ambiguity. Use `@[skills/brainstorming]` for discovery.

---

## Project Structure & Standards

- **frontend/**: Use stack defined in project (check existing code or `docs/requirements/`). Follow `frontend-specialist` rules.
- **backend/**: Use stack defined in project (check existing code or `docs/requirements/`). Follow `backend-specialist` rules.
- **docs/**: Source of truth for discovery, planning, and documentation. Create `docs/` if missing.
- **Stack Selection**: For new projects, use `@[skills/stack-selection]` to determine appropriate technologies with user input.
- Record cross-scope impacts in `docs/ACTION_ITEMS.md` (when it exists).
- Branches `feat/<area>-<slug>`; commits `feat|fix|docs|test|refactor|perf|chore|revert`.
- No `console.log`/`print`. Use official loggers.

---

## SDD Gate (Mandatory)

Before any `/create` execution, the following artifacts MUST exist:

**From `/specify` (Specification):**
- `docs/requirements/<feature>/PROBLEM_STATEMENT.md`
- `docs/requirements/<feature>/USER_STORIES.md`
- `docs/requirements/<feature>/ACCEPTANCE_CRITERIA.md`
- `docs/requirements/<feature>/RISKS.md`

**From `/specify` (Planning - included in specify):**
- `docs/requirements/<feature>/PLAN.md`
- `docs/sprint/Sprint-XX/SPRINT_GOAL.md`
- `docs/sprint/Sprint-XX/BACKLOG.md`
- `docs/sprint/Sprint-XX/RISK_REGISTER.md`

**From `/specify` (Task Breakdown - included in specify):**
- `docs/sprint/Sprint-XX/TASKS.md`

If any are missing, STOP and direct the user to run `/specify` first.

**Command Flow:**
```
/discover (MANDATORY) → /specify → /create → /verify → /deploy
```

---

## Standard Phase Workflow

All agents and commands MUST use this standardized phase naming:

| Phase | Name | Command(s) | Output | Code? |
|-------|------|------------|--------|-------|
| **0** | Discovery | `/discover` (MANDATORY) | `docs/CONTEXT.md`, decisions | NO |
| **1** | Specification | `/specify` | `docs/requirements/<feature>/`, `docs/sprint/Sprint-XX/` | NO |
| **2** | Implementation | `/create` | Working code | YES |
| **3** | Verification | `/verify` | Verified project | Scripts |
| **4** | Deployment | `/deploy` | Deployed application | - |

**STOP Points:** After Phase 1, Phase 2 (each P0-P3), and Phase 3 (user approval required)

---

## Discovery Gate (Phase 0 - Mandatory)

Before starting any feature work, execute discovery:

| Command | When to Use | Output |
|---------|-------------|--------|
| `/discover` | **ALWAYS** (mandatory) | `docs/CONTEXT.md`, `docs/SECURITY.md`, `docs/QUALITY_GATES.md` |

**Discovery Gate Rules:**
1. `/discover` is ALWAYS required before `/specify`
2. If `/discover` was not run, STOP and direct user to run it first

---

## Sprint Documentation & Artifacts

1. **Identification**: Locate `docs/sprint/Sprint-XX/` and review `HUB-SPRINT-XX.md`, `SPRINT_GOAL.md`, `BACKLOG.md`, `TASKS.md`.
2. **Sprint Selection**: Ask the user whether to use the latest sprint or create a new one.
   - If no sprint exists, create `Sprint-01`.
   - If creating a new sprint, use the next sequential number.
3. **Specification** (`/specify`): Create requirements in `docs/requirements/<feature>/` and update `HUB-SPRINT-XX.md`, `SPRINT_GOAL.md`, `BACKLOG.md`, `RISK_REGISTER.md`, `TASKS.md`.
4. **Execution** (`/create`): Mark progress in `TASKS.md`; update story status in `BACKLOG.md`.
5. **Completion**: Mark tasks as `[x]`, register changes in `docs/CHANGELOG.md` when requested.
6. **Templates**: Use `@[skills/documentation-templates]` for requirements, sprints, and reports.

**Command Ownership:**
| Artifact | Created By |
|----------|------------|
| docs/HUB-DOCS.md | `/discover` or `/specify` |
| docs/requirements/HUB-REQUIREMENTS.md | `/discover` or `/specify` |
| docs/requirements/<feature>/HUB-<FEATURE>.md | `/specify` |
| docs/sprint/HUB-SPRINTS.md | `/discover` or `/specify` |
| docs/sprint/Sprint-XX/HUB-SPRINT-XX.md | `/specify` |
| SPRINT_GOAL.md | `/specify` |
| BACKLOG.md | `/specify` |
| RISK_REGISTER.md | `/specify` |
| TASKS.md | `/specify` |

CRITICAL: Ending a task without syncing `docs/sprint/` is a protocol violation.

---

## Final Checklist & Scripts

1. Execution order: **Security → Lint → Schema → Tests → UX → SEO → Lighthouse/E2E**.
2. Main commands:
   - `python .opencode/scripts/checklist.py .`
   - `python .opencode/scripts/checklist.py . --url <URL>` (full pre-deploy)
3. Skill scripts (run based on scope):
   - Security: `python .opencode/skills/vulnerability-scanner/scripts/security_scan.py .`
   - Lint/Types: `python .opencode/skills/lint-and-validate/scripts/lint_runner.py .`
   - Tests: `python .opencode/skills/testing-patterns/scripts/test_runner.py .`
   - UX/Accessibility: `python .opencode/skills/frontend-design/scripts/ux_audit.py .`
   - SEO: `python .opencode/skills/seo-fundamentals/scripts/seo_checker.py .`
   - Mobile/Performance/etc.: see `skills/<skill>/scripts/`.

A task only ends when `checklist.py` succeeds. If it fails, resolve critical blockers (Security/Lint) before proceeding.

---

## Additional Best Practices

- References to commands (`/discover`, `/specify`, `/create`, `/verify`, `/orchestrate`, `/debug`, `/deploy`) must follow `.opencode/commands/`.
- When generating new artifacts, record times using timezone UTC-3 as per `AGENTS.md`.
- Update `docs/.context` and `docs/ACTION_ITEMS.md` when there is multi-repo impact.
- Always validate sensitive data: never expose secrets, tokens, or `.env`.

This file supersedes all previous rules (`CORE`, `ROUTING`, `CHECKLIST`, `SPRINTS`).

---

## References

- **Glossary:** `docs/GLOSSARY.md` - Standard terminology definitions
- **TodoList Protocol:** `.opencode/rules/TODOLIST_PROTOCOL.md` - Complete workflow
- **Tool Usage:** `.opencode/rules/TOOL_USAGE.md` - Correct tool selection
- **Obsidian Linking:** `.opencode/rules/OBSIDIAN_LINKING.md` - Documentation graph protocol
- **Single-Level Hub Links:** `.opencode/rules/SINGLE_LEVEL_HUB_LINKS.md` - Hub hierarchy policy
- **Docs File Glossary:** `.opencode/rules/DOCS_FILE_GLOSSARY.md` - Canonical documentation filenames
- **Conflict Analysis:** `docs/CONFLICT_ANALYSIS_REPORT.md` - Resolution history
