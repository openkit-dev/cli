# OpenKit Commands

> 7 Commands for the complete development workflow

---

## Available Commands

| # | Command | Description | Usage |
|---|---------|-------------|-------|
| 1 | `/discover` | Analyze project context (mandatory) | `/discover` |
| 2 | `/specify` | Specification + Planning + Tasks | `/specify add user auth` |
| 3 | `/create` | Implementation from plan | `/create from docs/sprint/Sprint-1/TASKS.md` |
| 4 | `/verify` | Quality verification (tests, lint, security) | `/verify all` |
| 5 | `/orchestrate` | Universal orchestrator for complex tasks | `/orchestrate build e-commerce` |
| 6 | `/debug` | Systematic debugging | `/debug login not working` |
| 7 | `/deploy` | Safe deployment | `/deploy staging` |

---

## Standard Workflow

```
/discover → /specify → /create → /verify → /deploy
```

For complex tasks, use `/orchestrate` to automate the entire flow.

---

## How to Use

### With Arguments
```bash
/specify add user authentication
/create from docs/sprint/Sprint-1/TASKS.md
/verify all
```

### Without Arguments (Interactive)
```bash
/specify
→ (via question tool) "Describe the feature"
→ System executes
```

---

## Workflow

### 1. Discover
```bash
/discover
→ Analyzes project structure
→ Creates context documentation
→ Identifies risks
```

### 2. Specify
```bash
/specify add user authentication
→ Creates specification (Problem Statement, User Stories, Acceptance Criteria)
→ Creates plan (PLAN.md, SPRINT_GOAL.md, BACKLOG.md)
→ Creates task breakdown (TASKS.md)
```

### 3. Create
```bash
/create from docs/sprint/Sprint-1/TASKS.md
→ Executes P0: Foundation (DB + Security)
→ Executes P1: Backend
→ Executes P2: UI/UX
→ Executes P3: Polish
```

### 4. Verify
```bash
/verify all
→ Runs lint + type check
→ Runs security scan
→ Runs unit tests
→ Runs UX audit
→ Runs performance checks
```

### 5. Deploy
```bash
/deploy staging
/deploy production
→ Prepares deployment
→ Executes deploy
→ Verifies post-deploy
```

---

## Command Mapping

| Old Command | New Command | Status |
|-------------|------------|--------|
| `/context` | `/discover` | Replaced |
| `/specify` | `/specify` | Expanded |
| `/clarify` | `/specify` | Absorbed |
| `/plan` | `/specify` | Absorbed |
| `/tasks` | `/specify` | Absorbed |
| `/impl` | `/create` | Replaced |
| `/test` | `/verify` | Replaced |
| `/checklist` | `/verify` | Absorbed |
| `/analyze` | `/verify` | Absorbed |
| `/engineer` | `/orchestrate` | Replaced |
| `/debug` | `/debug` | Maintained |
| `/deploy` | `/deploy` | Maintained |
| `/ui-ux` | `/orchestrate` | Absorbed |
| `/doc` | `/orchestrate` | Absorbed |
| `/status` | `/verify` | Absorbed |
| `/preview` | `/deploy` | Absorbed |
| `/create` | `/orchestrate` | Absorbed |
| `/brainstorm` | `/discover` | Absorbed |

---

## STOP Points

Each command has mandatory STOP points:

| After | Prompt |
|-------|--------|
| `/discover` | "Proceed to /specify?" |
| `/specify` | "Proceed to /create?" |
| `/create` (P0) | "Proceed to P1?" |
| `/create` (P1) | "Proceed to P2?" |
| `/create` (P2) | "Proceed to P3?" |
| `/create` (P3) | "Run /verify?" |
| `/verify` | "Proceed to /deploy?" |
| `/deploy` | "Confirm success?" |

---

## Usage Examples

### Example 1: New Feature
```bash
/discover
/specify add dark mode
/create from docs/sprint/Sprint-1/TASKS.md
/verify
/deploy staging
```

### Example 2: Bug Fix
```bash
/debug login not working after update
/verify
```

### Example 3: Complex Project
```bash
/orchestrate build e-commerce with Stripe
→ Automates entire workflow
→ Coordinates multiple agents
→ Runs verification
```

---

## Tips

1. **Always start with `/discover`** - Required before `/specify`
2. **Use `/orchestrate` for complex tasks** - Automates everything
3. **Use `/specify` for planned features** - Full specification flow
4. **Never skip STOP points** - Manual approval ensures quality
5. **Run `/verify` before deploy** - Ensures code quality

---

## Troubleshooting

### "Command not found"
→ Verify commands exist in `.opencode/commands/`

### "Python scripts do not run"
→ Ensure Python is installed and dependencies are satisfied

---

## Ready to use!

Open the OpenCode TUI and type `/` to see all available commands.
