# Migration Checklist

Use this file to migrate legacy documentation to the Obsidian-compatible standard.

## Filename Standard

- [ ] Rename docs files to canonical uppercase names from `[[.opencode/rules/DOCS_FILE_GLOSSARY.md]]`

## Link Standard

- [ ] Replace internal Markdown links with wikilinks `[[...]]`
- [ ] Keep external URLs as standard Markdown links

## Graph Connectivity

- [ ] Add `## Related` section to each major doc artifact
- [ ] Ensure each doc has at least one inbound path and two outbound links (when applicable)

## Hub Notes

- [ ] Ensure `[[HUB-DOCS.md]]` exists and links all major sections
- [ ] Ensure `[[requirements/HUB-REQUIREMENTS.md]]` exists
- [ ] Ensure `[[sprint/HUB-SPRINTS.md]]` exists

## Feature/Sprint Hubs

- [ ] Ensure each `docs/requirements/<feature>/HUB-<FEATURE>.md` links feature artifacts
- [ ] Ensure each `docs/sprint/Sprint-XX/HUB-SPRINT-XX.md` links sprint artifacts

## Validation

- [ ] Validate no broken wikilinks in docs
- [ ] Validate heading anchors used in links exist

## Related

- [[HUB-DOCS.md]]
- [[GLOSSARY.md]]
- [[CONTEXT.md]]
