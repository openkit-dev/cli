# Custom Commands - OpenKit

> Slash commands for the OpenCode TUI
> OpenKit integration with structured commands

---

## Available Commands

When you type `/` in the OpenCode TUI, you will see:

| Command | Description | Usage |
|---------|-------------|-------|
| `/engineer` | Universal orchestrator for complex tasks | `/engineer build full e-commerce` |
| `/specify` | Create feature specification | `/specify add user profiles` |
| `/clarify` | Clarify requirements in spec | `/clarify` |
| `/plan` | Create structured task plans | `/plan create auth system` |
| `/tasks` | Generate executable tasks | `/tasks` |
| `/analyze` | Cross-artifact consistency | `/analyze` |
| `/checklist` | Spec/plan quality checklist | `/checklist` |
| `/impl` | Execute implementation from a plan | `/impl from docs/sprint/Sprint-17/TASKS.md` |
| `/test` | Run tests and checks | `/test all` or `/test security` |
| `/debug` | 4-phase systematic debugging | `/debug login not working` |
| `/ui-ux` | Design systems and UX audits | `/ui-ux create design system` |
| `/deploy` | Safe deploy with verification | `/deploy staging` or `/deploy production` |
| `/context` | Generate repo context pack | `/context` |
| `/doc` | Write or update docs | `/doc update api index` |
| `/preview` | Manage local preview server | `/preview start` |
| `/status` | Show project status | `/status` |
| `/create` | Create a new application | `/create blog site` |
| `/brainstorm` | Explore ideas and options | `/brainstorm auth system` |

---

## How to Use

### 1. With Arguments (Direct)
```bash
/plan create JWT auth system
/engineer build e-commerce with Stripe checkout
```

### 2. Without Arguments (Interactive)
```bash
/plan
→ (via question tool) "What would you like to plan?"
→ You describe it
→ System executes
```

### 3. Full Flow
```bash
# 1. Specify
/specify add dark mode feature
→ Creates docs/requirements/<feature>/
→ STOP: "Proceed to clarify or plan?"

# 2. Clarify
/clarify
→ Resolves ambiguities in spec
→ STOP: "Proceed to plan?"

# 3. Plan
/plan add dark mode feature
→ Creates plan + auxiliary artifacts when applicable
→ STOP: "Review and confirm?"

# 4. Tasks
/tasks
→ Generates docs/sprint/Sprint-XX/TASKS.md
→ STOP: "Proceed to implement?"

# 5. Implement (after approval)
/impl from docs/sprint/Sprint-XX/TASKS.md
→ Executes plan tasks
→ STOP between phases (P0→P1→P2→P3)

# 6. Analyze + Checklist
/analyze
→ Validates spec/plan/tasks consistency
/checklist
→ Confirms readiness and gaps

# 7. Test
/test all
→ Runs all checks
→ Reports results

# 8. Deploy (if approved)
/deploy production
→ Pre-deploy checklist
→ STOP: "Confirm production deploy?"
→ Executes deploy
→ Post-deploy verification
```

---

## Hybrid Behavior

All commands follow a hybrid behavior:

1. **Show:** Load and explain the appropriate command protocol
2. **Ask (via question tool):** Mandatory STOP points for approval
3. **Execute:** Only after user confirmation

## Obsidian Documentation Standard

All command-generated docs must use Obsidian-compatible internal links:

- Prefer wikilinks for internal references: `[[HUB-DOCS.md]]`
- Keep external references as regular Markdown links
- Add `## Related` sections to connect artifacts across `docs/`
- Follow canonical docs filenames from `.opencode/rules/DOCS_FILE_GLOSSARY.md`

---

## STOP Points

The system always pauses for approval at critical moments:

### `/plan`
> "Plan recorded in `docs/requirements/` and `docs/sprint/Sprint-XX/`. Review and confirm to proceed?"

### `/impl`
- Between each phase (P0, P1, P2, P3)
- At the end: "Run final verification?"

### `/deploy`
- Production requires explicit confirmation
- After deploy: "Confirm success?"

---

## File Structure

```
.opencode/
└── commands/
    ├── analyze.md     # Spec/plan/tasks analysis
    ├── brainstorm.md  # Idea exploration
    ├── checklist.md   # Spec/plan checklist
    ├── clarify.md     # Spec clarification
    ├── context.md     # Context pack generation
    ├── create.md      # New app creation
    ├── doc.md         # Documentation updates
    ├── engineer.md    # Universal orchestrator
    ├── impl.md        # Implementation command
    ├── plan.md        # Planning command
    ├── preview.md     # Preview management
    ├── specify.md     # Feature specification
    ├── status.md      # Status reporting
    ├── tasks.md       # Task generation
    ├── test.md        # Testing command
    ├── ui-ux.md       # Design command
    └── deploy.md      # Deploy command
```

```
.opencode/
├── agents/            # Specialist agents
├── skills/            # Knowledge modules
└── scripts/           # Python validation scripts
```

---

## Command Features

### Placeholders
- `$ARGUMENTS` - All user arguments
- `$1`, `$2`, `$3` - Positional arguments

### Shell Integration
- `!command` - Inject shell output
- Example: `!git log --oneline -10`

### File References
- `@file` - Include file content
- Example: `@src/components/Button.tsx`

---

## Configuration

### Local (This Project)
Files in `.opencode/commands/` are available only in this project.

### Global (All Projects)
Copy to `~/.config/opencode/commands/` to use everywhere:

```bash
cp -r .opencode/commands/* ~/.config/opencode/commands/
```

---

## Usage Examples

### Example 1: Simple Feature
```bash
# Plan
/plan add dark mode toggle
→ Updates docs artifacts

# Approve and implement
/impl from docs/sprint/Sprint-XX/TASKS.md
→ Executes tasks

# Test
/test coverage
→ 95% coverage, all passing
```

### Example 2: Complex System
```bash
# Orchestrate everything
/engineer build e-commerce with Stripe
→ Phase 1: Planning (project-planner)
 → STOP: "Proceed?"
→ Phase 2: Implementation (multiple agents)
→ STOP: "Verify?"
→ Phase 3: Verification
→ Project complete!
```

### Example 3: Debug
```bash
/debug users can't login after last update
→ Phase 1: Symptom Analysis
→ Phase 2: Information Gathering
→ Phase 3: Hypothesis Testing
→ Phase 4: Resolution
→ Bug fixed! 
```

---

## Tips

1. **Use `/engineer` for complex tasks** - It coordinates multiple agents
2. **Use `/plan` first** - Always plan before implementing
3. **Do not skip STOPs** - Manual approvals ensure quality
4. **Keep plans small** - If > 10 tasks, split into parts
5. **Run `/test` often** - Catch issues early

---

## 🆘 Troubleshooting

### "Command does not appear after typing /"
→ Restart the OpenCode TUI

### "Command not found"
→ Verify that `.opencode/commands/` exists

### "Python scripts do not run"
→ Ensure Python is installed and dependencies are satisfied

---

## OpenKit Integration

These commands reuse:
-  Specialist prompts (`.opencode/prompts/`)
-  Skills (`.opencode/skills/`)
-  Python validation scripts (`.opencode/scripts/`)
-  Original rules and protocols

They are only adapted to the OpenCode command format.

---

## Additional Documentation

- **OpenKit Architecture:** `.opencode/ARCHITECTURE.md`

---

## Ready to use!

Open the OpenCode TUI and type `/` to see all available commands. 
