# Rules Directory

> Central repository of mandatory protocols for the OpenCode agent system.

---

## Overview

This directory contains **system-wide rules** that enforce consistent behavior across:
- **Orchestrator** (`.opencode/prompts/orchestrator.md`)
- **All specialist prompts** (`.opencode/prompts/*.md`)
- **All commands** (`.opencode/commands/*.md`)
- **All skills** (`.opencode/skills/*/SKILL.md`)

---

## Rule Files

### 1. **MASTER.md** (Priority: P0)
**Purpose:** Single source of truth for all mandatory rules.

**Contains:**
- Tool usage reference (delegates to TOOL_USAGE.md)
- Context & language protocols
- Question tool protocol
- TodoList protocol
- Model selection guidelines
- Agent & skill protocol
- Project structure standards
- Sprint documentation requirements
- Final checklist & scripts
- Best practices

**Applies to:** All agents, commands, skills

**Trigger:** `always_on`

---

### 2. **TOOL_USAGE.md** (Priority: P0)
**Purpose:** Definitive guide for correct tool selection and usage patterns.

**Contains:**
- Tool selection decision tree
- File discovery & search patterns
- File operation rules (read → edit/write)
- Code execution & shell patterns
- User interaction protocols
- Task tracking lifecycle
- Knowledge loading patterns
- Common mistakes & fixes
- Validation checklist

**Applies to:** All agents, commands, skills

**Trigger:** `always_on`

**Key Enforcements:**
-  Use `glob` for file patterns, NOT `bash find`
-  Use `grep` for content search, NOT `bash grep`
-  Use `read` for file reading, NOT `bash cat`
-  Use `edit` for existing files, `write` for new files
-  Use `question` tool for multiple options, NOT inline questions
-  Use `todolist` tools for complex tasks (3+ steps), NOT inline tracking
-  Use `workdir` parameter, NOT `cd && command`

---

### 3. **OBSIDIAN_LINKING.md** (Priority: P0)
**Purpose:** Enforce Obsidian-compatible documentation with wikilinks and graph navigation.

**Contains:**
- Wikilink standards for internal docs
- Heading and block reference patterns
- Required `## Related` section conventions
- Graph connectivity requirements (inbound/outbound links)

**Applies to:** All agents, commands, skills

**Trigger:** `always_on`

---

### 4. **DOCS_FILE_GLOSSARY.md** (Priority: P0)
**Purpose:** Define canonical documentation filenames for all projects.

**Contains:**
- Global naming convention (artifact files in `TUDO_MAIUSCULO.md` and hub files in `HUB-<RESOURCE>.md`)
- Required and optional docs artifacts by context
- Sprint/requirements/ADR/runbook filename standards

**Applies to:** All agents, commands, skills

**Trigger:** `always_on`

---

## Rule Hierarchy

```
MASTER.md (Root)
    ↓
    ├─ TOOL_USAGE.md (Tool patterns)
    │
    ├─ OBSIDIAN_LINKING.md (Documentation graph patterns)
    │
    ├─ DOCS_FILE_GLOSSARY.md (Canonical doc filenames)
    │
    ├─ Agent Prompts (.opencode/prompts/*.md)
    │   └─ Reference MASTER.md + TOOL_USAGE.md
    │
    ├─ Command Files (.opencode/commands/*.md)
    │   └─ Reference MASTER.md + TOOL_USAGE.md
    │
    └─ Skill Files (.opencode/skills/*/SKILL.md)
        └─ Reference MASTER.md + TOOL_USAGE.md
```

---

## How Rules Are Applied

### 1. **Automatic Loading**
- All rules with `trigger: always_on` are loaded at system startup
- `MASTER.md` is ALWAYS active
- `TOOL_USAGE.md` is ALWAYS active

### 2. **Agent Invocation**
When an agent is invoked:
1. System loads `MASTER.md` (automatic)
2. System loads `TOOL_USAGE.md` (automatic)
3. Agent reads its own prompt (`.opencode/prompts/<agent>.md`)
4. Agent loads skills from frontmatter

### 3. **Command Execution**
When a command is executed:
1. System loads `MASTER.md` (automatic)
2. System loads `TOOL_USAGE.md` (automatic)
3. Command reads its own file (`.opencode/commands/<command>.md`)
4. Command follows protocols from rules

### 4. **Skill Loading**
When a skill is loaded:
1. System loads `MASTER.md` (automatic)
2. System loads `TOOL_USAGE.md` (automatic)
3. Skill instructions are read from `SKILL.md`
4. Skill follows protocols from rules

---

## Enforcement

### Priority Levels
- **P0 (Critical):** Must be followed. Violations break the system.
- **P1 (High):** Should be followed. Violations cause inconsistency.
- **P2 (Medium):** Recommended. Violations reduce quality.
- **P3 (Low):** Nice to have. Violations are acceptable.

### Conflict Resolution
1. `MASTER.md` supersedes all other rules
2. `TOOL_USAGE.md` supersedes agent/command/skill instructions for tool usage
3. Agent-specific rules override generic patterns (within their domain)
4. User instructions override all rules (when explicit)

---

## Validation

### For Agents
Before any action, verify:
- [ ] Have I read `MASTER.md`? (automatic)
- [ ] Have I read `TOOL_USAGE.md`? (automatic)
- [ ] Have I read `OBSIDIAN_LINKING.md`? (automatic)
- [ ] Have I read `DOCS_FILE_GLOSSARY.md`? (automatic)
- [ ] Am I using the correct tool? (check decision tree)
- [ ] Have I loaded required skills? (check frontmatter)
- [ ] Am I following todolist protocol? (for complex tasks)
- [ ] Am I following question tool protocol? (for multiple options)

### For Commands
Before execution, verify:
- [ ] Have I read `MASTER.md`? (automatic)
- [ ] Have I read `TOOL_USAGE.md`? (automatic)
- [ ] Have I read `OBSIDIAN_LINKING.md`? (automatic)
- [ ] Have I read `DOCS_FILE_GLOSSARY.md`? (automatic)
- [ ] Am I using the correct tool? (check decision tree)
- [ ] Am I following user interaction patterns? (question/todolist)

### For Skills
Before applying, verify:
- [ ] Have I read `MASTER.md`? (automatic)
- [ ] Have I read `TOOL_USAGE.md`? (automatic)
- [ ] Have I read `OBSIDIAN_LINKING.md`? (automatic)
- [ ] Have I read `DOCS_FILE_GLOSSARY.md`? (automatic)
- [ ] Am I using the correct tool? (check decision tree)
- [ ] Am I announcing skill usage? (`Applying knowledge of @[skill-name]...`)

---

## Quick Reference

### Tool Configuration

**Tools are configured in `opencode.json`**, NOT in prompt markdown files.

Example orchestrator configuration:
```json
{
  "agent": {
    "orchestrator": {
      "tools": {
        "read": true,
        "grep": true,
        "glob": true,
        "list": true,
        "bash": true,
        "edit": true,
        "write": true,
        "patch": true,
        "webfetch": true,
        "skill": true,
        "todowrite": true,
        "todoread": true,
        "question": true
      }
    }
  }
}
```

### Tool Selection
| Need | Use | NOT |
|------|-----|-----|
| Find files by name | `glob` | `bash find` |
| Find files by content | `grep` | `bash grep` |
| List directory | `list` | `bash ls` |
| Read file | `read` | `bash cat` |
| Edit existing file | `edit` | `write` (unless rewriting) |
| Create new file | `write` | `bash echo >` |
| Apply patch | `patch` | `bash patch` |
| Execute command | `bash` | Manual execution |
| Change directory | `workdir` param | `cd &&` |
| Ask multiple options | `question` tool | Inline text |
| Track complex tasks | `todolist` tools | Inline text |

### Common Patterns
```javascript
//  File discovery
glob("**/*.tsx")
grep("useState", "src/**/*.tsx")

//  File operations
read("file.ts")
edit("file.ts", "old", "new")

//  User interaction
question({
  questions: [{
    question: "Choose option",
    header: "Title",
    options: [...]
  }]
})

//  Task tracking
todoread()
todowrite({ todos: [...] })

//  Command execution
bash("pytest tests", workdir="/backend")
```

---

## Maintenance

### Adding New Rules
1. Create new `.md` file in `.opencode/rules/`
2. Add frontmatter with `trigger` and `priority`
3. Update `MASTER.md` to reference it (if needed)
4. Update this README with the new rule

### Updating Existing Rules
1. Edit the rule file directly
2. Update version/date in frontmatter (if present)
3. Announce changes to all agents (via system update)

### Deprecating Rules
1. Mark rule as `deprecated: true` in frontmatter
2. Add `superseded_by: <new-rule>.md` reference
3. Keep file for 2 versions, then delete

---

## Related Documentation

- **Agent Prompts**: `.opencode/prompts/`
- **Command System**: `.opencode/commands/README.md`
- **Skill System**: `memory/SKILLS.md`
- **Obsidian Protocol**: `.opencode/rules/OBSIDIAN_LINKING.md`
- **Docs File Glossary**: `.opencode/rules/DOCS_FILE_GLOSSARY.md`
- **Architecture**: `.opencode/ARCHITECTURE.md`

---

**Last Updated:** 2026-02-04  
**Maintained By:** OpenCode System
