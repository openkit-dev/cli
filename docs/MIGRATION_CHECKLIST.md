# Docs Migration Checklist

## Scope

Migration of `docs/` to OpenKit Obsidian-compatible standard.

## Progress

- [x] Audited current `docs/` structure and naming.
- [x] Re-audited `docs/` for OpenKit Obsidian compliance (2026-02-11).
- [x] Renamed non-canonical files to uppercase canonical names where applicable.
- [x] Created mandatory hubs:
  - [x] `docs/HUB-DOCS.md`
  - [x] `docs/requirements/HUB-REQUIREMENTS.md`
  - [x] `docs/sprint/HUB-SPRINTS.md`
  - [x] `docs/requirements/<feature>/HUB-<FEATURE>.md` for all features
  - [x] `docs/sprint/Sprint-XX/HUB-SPRINT-XX.md` for all Sprint-01..06
- [x] Added `docs/adr/HUB-ADR.md` hub.
- [x] Added `docs/content-protocol/README.md` hub.
- [x] Added `docs/agent-compat/agents/README.md` hub.
- [x] Normalized `## Related` sections across root docs, requirements, sprint artifacts, ADR, and protocol docs.
- [x] Converted key internal path references to Obsidian wikilinks.
- [x] Ran Obsidian link lint and fixed broken links.
- [x] Applied mandatory hub rename map from `README.md` to canonical `HUB-*.md` files.
- [x] Updated internal wikilinks that referenced renamed hub files.
- [x] Removed stray nested legacy file `docs/docs/requirements/semantic-memory/README.md`.

## Rename Map

- `docs/requirements/cursor-codex-targets/analysis.md` -> `docs/requirements/cursor-codex-targets/ANALYSIS.md`
- `docs/agent-compat/agents/claude.md` -> `docs/agent-compat/agents/CLAUDE.md`
- `docs/agent-compat/agents/codex.md` -> `docs/agent-compat/agents/CODEX.md`
- `docs/agent-compat/agents/cursor.md` -> `docs/agent-compat/agents/CURSOR.md`
- `docs/agent-compat/agents/gemini.md` -> `docs/agent-compat/agents/GEMINI.md`
- `docs/agent-compat/agents/opencode.md` -> `docs/agent-compat/agents/OPENCODE.md`
- `docs/README.md` -> `docs/HUB-DOCS.md`
- `docs/requirements/README.md` -> `docs/requirements/HUB-REQUIREMENTS.md`
- `docs/sprint/README.md` -> `docs/sprint/HUB-SPRINTS.md`
- `docs/adr/README.md` -> `docs/adr/HUB-ADR.md`
- `docs/requirements/cursor-codex-targets/README.md` -> `docs/requirements/cursor-codex-targets/HUB-CURSOR-CODEX-TARGETS.md`
- `docs/requirements/gemini-sync/README.md` -> `docs/requirements/gemini-sync/HUB-GEMINI-SYNC.md`
- `docs/requirements/opencode-scripts/README.md` -> `docs/requirements/opencode-scripts/HUB-OPENCODE-SCRIPTS.md`
- `docs/requirements/openkit-sync/README.md` -> `docs/requirements/openkit-sync/HUB-OPENKIT-SYNC.md`
- `docs/requirements/remove-blueprints-references/README.md` -> `docs/requirements/remove-blueprints-references/HUB-REMOVE-BLUEPRINTS-REFERENCES.md`
- `docs/requirements/security-scan-hardening/README.md` -> `docs/requirements/security-scan-hardening/HUB-SECURITY-SCAN-HARDENING.md`
- `docs/requirements/semantic-memory/README.md` -> `docs/requirements/semantic-memory/HUB-SEMANTIC-MEMORY.md`
- `docs/sprint/Sprint-01/README.md` -> `docs/sprint/Sprint-01/HUB-SPRINT-01.md`
- `docs/sprint/Sprint-02/README.md` -> `docs/sprint/Sprint-02/HUB-SPRINT-02.md`
- `docs/sprint/Sprint-03/README.md` -> `docs/sprint/Sprint-03/HUB-SPRINT-03.md`
- `docs/sprint/Sprint-04/README.md` -> `docs/sprint/Sprint-04/HUB-SPRINT-04.md`
- `docs/sprint/Sprint-05/README.md` -> `docs/sprint/Sprint-05/HUB-SPRINT-05.md`
- `docs/sprint/Sprint-06/README.md` -> `docs/sprint/Sprint-06/HUB-SPRINT-06.md`

## Risks / Notes

- Legacy text references in code blocks/backticks remain in historical notes when they are not navigational links.
- Some files include historical paths under `cli/...`; these were preserved as context unless they were clear internal doc references.
- No content semantics were rewritten; migration remained structural (hubs, links, and file placement).

## Related

- [[docs/HUB-DOCS.md]]
- [[docs/requirements/HUB-REQUIREMENTS.md]]
- [[docs/sprint/HUB-SPRINTS.md]]
- [[.opencode/rules/OBSIDIAN_LINKING.md]]
- [[.opencode/rules/DOCS_FILE_GLOSSARY.md]]
