---
trigger: always_on
priority: P0
applies_to: [orchestrator, all-agents, all-commands, all-skills]
---

# Single-Level Hub Linking Policy

Enforce layered navigation and avoid cross-level backlinks in documentation.

## Overview

This policy establishes a hierarchical linking structure for all documentation to maintain a clean knowledge graph. Links should flow in a single direction through the hierarchy without skipping levels.

## Hierarchy Layers

| Layer | Examples | Links To |
|-------|----------|----------|
| **Layer 0 (Root)** | `docs/HUB-DOCS.md`, `docs/CONTEXT.md` | Layer 1 hubs only |
| **Layer 1 (Section Hubs)** | `docs/requirements/HUB-REQUIREMENTS.md`, `docs/sprint/HUB-SPRINTS.md` | Layer 0, Layer 2, sibling Layer 1 |
| **Layer 2 (Feature/Sprint Hubs)** | `docs/requirements/<feature>/HUB-<FEATURE>.md`, `docs/sprint/Sprint-XX/HUB-SPRINT-XX.md` | Layer 1, Layer 3, local peers |
| **Layer 3 (Artifacts)** | `docs/sprint/Sprint-XX/TASKS.md`, `docs/requirements/<feature>/PLAN.md` | Layer 2 only |

## Linking Rules

### Rule 1: No Level Skipping

Links MUST NOT skip layers:

```
Allowed:   HUB-DOCS -> HUB-SPRINTS -> HUB-SPRINT-20 -> TASKS.md
Disallowed: HUB-DOCS -> HUB-SPRINT-20 (skips HUB-SPRINTS)
Disallowed: TASKS.md -> HUB-SPRINTS (skips HUB-SPRINT-20)
```

### Rule 2: Direction of Flow

- **Downward flow** (parent to child): Always allowed within same branch
- **Upward flow** (child to parent): Must go to immediate parent only
- **Cross-layer flow**: Only allowed between sibling hubs at same layer

### Rule 3: Hub-to-Hub Navigation

```
HUB-DOCS.md
  |-- docs/*.md (root docs)
  |-- docs/requirements/HUB-REQUIREMENTS.md (Layer 1)
  |-- docs/sprint/HUB-SPRINTS.md (Layer 1)
  |-- docs/adr/HUB-ADR.md (Layer 1)

requirements/HUB-REQUIREMENTS.md
  |-- HUB-DOCS.md (Layer 0 - back reference)
  |-- requirements/<feature>/HUB-<FEATURE>.md (Layer 2)
  |-- sprint/HUB-SPRINTS.md (sibling Layer 1)

sprint/Sprint-20/TASKS.md
  |-- HUB-SPRINT-20.md (Layer 2 - immediate parent)
  |-- HUB-SPRINT-20.md#BACKLOG (local peer reference)
```

### Rule 4: Related Section Format

Use only immediate-level references in `## Related` sections:

```markdown
## Related

- [[HUB-<LOCAL-SCOPE>.md]]        <!-- Immediate parent hub -->
- [[<LOCAL-PEER>.md]]              <!-- Sibling at same level -->
```

**Do NOT use:**
```markdown
## Related (AVOID)

- [[HUB-DOCS.md]]                  <!-- From Layer 3 to Layer 0 -->
- [[HUB-SPRINT-20.md]]             <!-- From Layer 3 to Layer 2 when parent is Layer 2 -->
```

## Examples

### Feature Documentation

```
docs/requirements/auth-system/
  |-- HUB-AUTH-SYSTEM.md        (Layer 2)
  |-- PROBLEM_STATEMENT.md      (Layer 3)
  |-- USER_STORIES.md           (Layer 3)
  |-- ACCEPTANCE_CRITERIA.md    (Layer 3)
  |-- PLAN.md                   (Layer 3)
  |-- RISKS.md                  (Layer 3)

HUB-AUTH-SYSTEM.md links to:
  - [[requirements/HUB-REQUIREMENTS.md]]     (Layer 1 - OK)
  - [[PROBLEM_STATEMENT.md]]                (Layer 3 peer - OK)
  
PROBLEM_STATEMENT.md links to:
  - [[HUB-AUTH-SYSTEM.md]]                  (Layer 2 parent - OK)
  - [[USER_STORIES.md]]                     (Layer 3 peer - OK)
  - [[requirements/HUB-REQUIREMENTS.md]]    (SKIP - NOT ALLOWED)
```

### Sprint Documentation

```
docs/sprint/Sprint-05/
  |-- HUB-SPRINT-05.md          (Layer 2)
  |-- SPRINT_GOAL.md            (Layer 3)
  |-- BACKLOG.md                (Layer 3)
  |-- TASKS.md                  (Layer 3)
  |-- RISK_REGISTER.md          (Layer 3)

HUB-SPRINT-05.md links to:
  - [[sprint/HUB-SPRINTS.md]]  (Layer 1 - OK)
  - [[TASKS.md]]                (Layer 3 peer - OK)
  
TASKS.md links to:
  - [[HUB-SPRINT-05.md]]       (Layer 2 parent - OK)
  - [[BACKLOG.md]]              (Layer 3 peer - OK)
  - [[sprint/HUB-SPRINTS.md]]   (SKIP - NOT ALLOWED)
```

## Enforcement

This policy is enforced through:

1. **Documentation Review**: Links are validated during code review
2. **Obsidian Link Lint**: Use `obsidian-link-lint` tool to detect violations
3. **Graph Visualization**: Check backlink panel in Obsidian for unexpected connections

### Lint Command

```bash
obsidian-link-lint --root docs
```

This will report:
- Level skipping violations
- Missing parent hub references
- Broken wikilinks

## Exceptions

Exceptions require approval from the documentation lead:

1. **Cross-cutting concerns**: Links that genuinely span layers (e.g., GLOSSARY references)
2. **Legacy documentation**: Pre-existing docs may have non-compliant links
3. **Generated content**: Automated docs may temporarily violate until regenerated

## Related

- [[OBSIDIAN_LINKING.md|Obsidian Documentation Protocol]]
- [[DOCS_FILE_GLOSSARY.md|Docs File Glossary]]
