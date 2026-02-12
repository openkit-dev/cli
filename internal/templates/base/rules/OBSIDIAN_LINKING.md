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

## Hub Notes (Mandatory)

Use index notes to improve graph navigation and context retrieval:

- `HUB-DOCS.md` as global documentation hub.
- `requirements/HUB-REQUIREMENTS.md` as requirements hub.
- `sprint/HUB-SPRINTS.md` as sprint hub.
- `requirements/<feature>/HUB-<FEATURE>.md` as feature hub.
- `sprint/Sprint-XX/HUB-SPRINT-XX.md` as sprint execution hub.

## Required Related Section

Whenever applicable, include:

```markdown
## Related

- [[HUB-DOCS.md]]
- [[GLOSSARY.md]]
```

For feature docs, prefer feature-local references:

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

## Enforcement

This protocol is mandatory for rule files, command docs, skill docs, and sprint or requirement artifacts in `docs/`.

If a generated or updated doc does not include Obsidian-compatible internal links, the task is incomplete.

## References

- [Obsidian Help](https://help.obsidian.md/)
- [Internal links](https://help.obsidian.md/links)
- [Obsidian Flavored Markdown](https://help.obsidian.md/obsidian-flavored-markdown)
