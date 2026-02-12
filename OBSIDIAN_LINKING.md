---
trigger: always_on
priority: P0
applies_to: [orchestrator, all-agents, all-commands, all-skills]
---

# OBSIDIAN DOCUMENTATION PROTOCOL

Mandatory documentation standard for OpenKit using Obsidian-compatible linking.

## Purpose

All project documentation must be written as a connected knowledge graph.
Use Obsidian-compatible links so docs can be navigated, indexed, and reused as long-term memory.

## Required Link Standard

1. Prefer wikilinks for internal Markdown references.
   - `[[CONTEXT.md]]`
   - `[[requirements/<feature>/PLAN.md]]`
   - `[[requirements/<feature>/PLAN.md#Success Criteria]]`
2. Use aliases when display text should be shorter.
   - `[[GLOSSARY.md|Glossary]]`
3. Use heading links for precise references.
   - `[[SECURITY.md#Threats]]`
4. Use block references only when exact paragraph traceability is needed.
   - `[[CONTEXT.md#^evidence-cli-entry]]`
5. Use standard Markdown links only for external URLs.
   - `[Obsidian Help](https://help.obsidian.md/)`

## Graph Requirements

For every new or updated doc in `docs/`:

1. Add at least one inbound path from an index or parent doc.
2. Add at least two outbound wikilinks to related docs when relevant.
3. Keep references stable and explicit (full path inside the vault).

## Link Hierarchy Policy (Mandatory)

Name: **Single-Level Hub Linking Policy** (Layered Hub-and-Spoke)

Goal: enforce layered navigation and avoid cross-level backlinks.

### Rule

1. `HUB-DOCS.md` links only to:
   - root docs (`docs/*.md`)
   - first-level hubs (`docs/<section>/HUB-*.md`)
2. First-level hubs (for example `requirements/HUB-REQUIREMENTS.md`, `sprint/HUB-SPRINTS.md`) link only to:
   - sibling first-level hubs when relevant
   - direct child hubs in their own section
3. Deep docs (for example `docs/sprint/Sprint-20/TASKS.md`) link only to:
   - their immediate local hub (for example `[[HUB-SPRINT-20.md]]`)
   - local peer docs when relevant
4. No level skipping:
   - child docs MUST NOT link directly to grandparent hubs
   - parent hubs SHOULD NOT link directly to grandchild docs

### Allowed vs Disallowed

- Allowed: `HUB-DOCS -> HUB-SPRINTS -> HUB-SPRINT-20 -> TASKS.md`
- Disallowed: `HUB-DOCS -> HUB-SPRINT-20`
- Disallowed: `TASKS.md -> HUB-SPRINTS`

## Hub Notes (Mandatory)

Use index notes to improve graph navigation and context retrieval:

- `HUB-DOCS.md` as global documentation hub.
- `requirements/HUB-REQUIREMENTS.md` as requirements hub.
- `sprint/HUB-SPRINTS.md` as sprint hub.
- `requirements/<feature>/HUB-<FEATURE>.md` as feature hub.
- `sprint/Sprint-XX/HUB-SPRINT-XX.md` as sprint execution hub.

Each hub note MUST include:

1. `## Context` describing scope and intended usage.
2. `## Status Overview` summarizing the current state (phase, completion, blockers, readiness, or risk).

## Required Related Section

Whenever applicable, include only immediate-level navigation links:

```markdown
## Related

- [[HUB-<LOCAL-SCOPE>.md]]
- [[<LOCAL-PEER>.md]]
```

For feature docs, use feature-local references only:

```markdown
## Related

- [[requirements/<feature>/PROBLEM_STATEMENT.md]]
- [[requirements/<feature>/USER_STORIES.md]]
- [[requirements/<feature>/ACCEPTANCE_CRITERIA.md]]
- [[requirements/<feature>/PLAN.md]]
```

## Obsidian Compatibility Notes

- Keep filenames stable after publication.
- Prefer kebab-case folders and canonical artifact names used by OpenKit.
- Avoid broken anchors: heading text in links must match target headings.
- For documentation files, follow canonical uppercase naming from `[[.opencode/rules/DOCS_FILE_GLOSSARY.md]]`.
- Preserve hub adjacency: links should stay within immediate parent/child layers.

## Enforcement

This protocol is mandatory for rule files, command docs, skill docs, and sprint or requirement artifacts in `docs/`.

If a generated or updated doc does not include Obsidian-compatible internal links, the task is incomplete.

## References

- [Obsidian Help](https://help.obsidian.md/)
- [Internal links](https://help.obsidian.md/links)
- [Obsidian Flavored Markdown](https://help.obsidian.md/obsidian-flavored-markdown)
