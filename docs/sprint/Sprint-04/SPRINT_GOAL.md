# Sprint 04 Goal

**Objective:** Enhance Cursor and Codex targets to provide rich, agent-specific configurations.

## Scope

1. **Cursor Target:**
   - Generate `.cursorrules` with meaningful content
   - Generate `.cursor/rules/openkit.mdc` for modular rules
   - Copy skills to `.cursor/skills/`

2. **Codex Target:**
   - Generate comprehensive `AGENTS.md`
   - Generate `.codex/rules/openkit.rules` with Starlark syntax
   - Maintain skills in `.agents/skills/`

3. **Doctor Enhancements:**
   - Improve checks for Cursor files
   - Improve checks for Codex files

4. **Testing:**
   - Unit tests for content generation
   - Integration tests for sync commands

## Success Metrics

- All sync commands produce expected files
- All tests pass (`go test ./...`)
- Doctor commands provide useful diagnostics
- Documentation updated

## Out of Scope

- Windsurf target (future sprint)
- CI/CD pipeline (future sprint)
- Init command improvements (future sprint)
