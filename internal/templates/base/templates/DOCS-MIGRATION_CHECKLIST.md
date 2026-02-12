# Migration Checklist

Use this file to migrate legacy documentation to the Obsidian-compatible standard.

## Filename Standard

- [ ] Rename docs files to canonical uppercase names from `[[.opencode/rules/DOCS_FILE_GLOSSARY.md]]`

## Link Standard

- [ ] Replace internal Markdown links with wikilinks `[[docs/...]]`
- [ ] Keep external URLs as standard Markdown links

## Graph Connectivity

- [ ] Add `## Related` section to each major doc artifact
- [ ] Ensure each doc has at least one inbound path and two outbound links (when applicable)

## Hub Notes

- [ ] Ensure `[[docs/README.md]]` exists and links all major sections
- [ ] Ensure `[[docs/requirements/README.md]]` exists
- [ ] Ensure `[[docs/sprint/README.md]]` exists

## Feature/Sprint Hubs

- [ ] Ensure each `docs/requirements/<feature>/README.md` links feature artifacts
- [ ] Ensure each `docs/sprint/Sprint-XX/README.md` links sprint artifacts

## Validation

- [ ] Validate no broken wikilinks in docs
- [ ] Validate heading anchors used in links exist

## Related

- [[docs/README.md]]
- [[docs/GLOSSARY.md]]
- [[docs/CONTEXT.md]]
