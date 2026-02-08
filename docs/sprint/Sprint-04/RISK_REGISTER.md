# Sprint 04 Risk Register

| ID | Risk | Probability | Impact | Mitigation | Status |
|----|------|-------------|--------|------------|--------|
| R1 | Cursor .mdc format not documented | Medium | Medium | Use simple markdown, test with Cursor | Open |
| R2 | Codex Starlark syntax errors | Low | High | Add syntax validation test | Open |
| R3 | AGENTS.md exceeds 32KB | Low | Medium | Keep content concise, reference skills | Open |
| R4 | Breaking existing installations | Medium | Medium | Safe-by-default, clear conflict messaging | Open |
| R5 | Cursor skills path not standard | Low | Low | Use convention, document in rules | Open |

## Contingency Plans

### R1 Contingency
If `.cursor/rules/*.mdc` doesn't work:
- Fallback to enhanced `.cursorrules` only
- Document limitation in agent-compat docs

### R2 Contingency
If Starlark validation fails:
- Simplify rules to most basic prefix_rule
- Test with actual Codex CLI before release

### R4 Contingency
If users report conflicts:
- Improve doctor messaging
- Add `--force` migration flag
