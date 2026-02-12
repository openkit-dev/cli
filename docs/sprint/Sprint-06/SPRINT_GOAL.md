# SPRINT GOAL

**Sprint**: Sprint-06
**Title**: Context-Aware OpenKit Refactoration
**Created**: 2026-02-10
**Status**: In Progress
**Duration**: 4-5 weeks

## Objective

Transform OpenKit from a web-focused (backend + frontend) system to a **context-aware, project-type-adaptive** system that intelligently generates relevant documentation and skills based on the actual project structure.

## Problem Statement

Current OpenKit assumes all projects follow the web application pattern with separate `backend/` and `frontend/` directories. This causes issues for:

- **CLI tools** (like OpenKit itself): No backend/frontend structure
- **Libraries/SDKs**: Different architecture, focus on public API
- **Desktop applications**: Hybrid structure, not web standard
- **Serverless functions**: Isolated functions, no backend/frontend split
- **Infrastructure/DevOps**: Terraform, Docker, K8s configurations
- **Plugins/Extensions**: Browser extensions, IDE plugins

This results in:
- Irrelevant documentation (`BACKEND.md`, `FRONTEND.md`) being generated for non-web projects
- Inappropriate skills being loaded for project types where they don't apply
- Confusion and noise in project documentation

## Solution: Hybrid System (Type-Based + Adaptive Overlays)

Implement **Option C** from brainstorming: A hybrid system combining explicit project types with adaptive overlays.

### Key Components

1. **Project Type Registry**
   - 6+ project types: `cli-tool`, `web-fullstack`, `library`, `desktop-app`, `serverless`, `infrastructure`, `plugin`
   - Each type defines: detection rules, base context docs, relevant skills, excluded skills
   - Explicit JSON schemas for configuration

2. **Overlay System**
   - 6+ overlays: `testing`, `security`, `ci-cd`, `documentation`, `performance`, `i18n`
   - Overlays add context docs and skills conditionally
   - Composable with any project type
   - Reusable across types

3. **Detection Engine**
   - Heuristic-based project type detection
   - File pattern matching (required vs suggested vs conflicting)
   - Confidence scoring
   - Hybrid detection (auto-detect + user confirmation)

4. **Adaptive Workflows**
   - Modify `/context` to use project type and overlays
   - Modify `/specify`, `/plan`, `/tasks` to be type-aware
   - Keep skills existing, only deactivate when inappropriate

## Success Criteria

- [ ] `/context` detects project type correctly (CLI, web, library, etc.)
- [ ] `/context` generates only relevant documentation for detected type
- [ ] Skills are loaded dynamically based on project type
- [ ] User can confirm/override detected type
- [ ] User can select overlays interactively
- [ ] No documentation generated for irrelevant domains (e.g., no `FRONTEND.md` for CLI tools)
- [ ] Backward compatible with existing web-fullstack projects
- [ ] New project types fully documented and tested

## Scope

### In Scope

1. **Detection System** (Week 1)
   - Create `internal/detection/` package
   - Implement project type detection with heuristics
   - Implement overlay detection
   - Confidence scoring
   - Conflict detection

2. **Project Type Configuration** (Week 1-2)
   - Create `.opencode/project-types/` directory structure
   - Define JSON schemas for project types and overlays
   - Create 6+ initial project types
   - Create 4+ initial overlays

3. **Context Generation Refactor** (Week 2)
   - Modify `internal/cli/context.go` to use new system
   - Implement project type loading
   - Implement overlay application
   - Update templates for new project types

4. **New Skills** (Week 3)
   - Create `cli-design` skill
   - Create `library-patterns` skill
   - Create `desktop-patterns` skill
   - Create `serverless-patterns` skill
   - Create `iac-patterns` skill

5. **Testing** (Week 4)
   - Write unit tests for detection engine
   - Write integration tests for context generation
   - Test with 6+ project types (CLI, web, library, etc.)
   - Test edge cases (monorepo, ambiguous detection)

6. **Documentation** (Week 4-5)
   - Update project documentation
   - Create migration guide
   - Update README with new capabilities

### Out of Scope

- Complete rewrite of existing skills (keep as-is)
- New agent creation (use existing specialists)
- UI redesign for OpenKit CLI (use existing CLI patterns)
- Breaking changes to existing `/context` command output format (extend, don't replace)

## Dependencies

- Go 1.25.7 (existing)
- Existing Cobra CLI framework
- Existing skill system architecture
- `jsonschema` library for schema validation (new dependency)

## Risks

| Risk | Impact | Mitigation |
|-------|---------|------------|
| False positive detection | High | Confidence thresholds, user confirmation required |
| Complex detection logic | Medium | Clear heuristics, extensive unit tests |
| Breaking existing workflows | High | Backward compatibility mode for web-fullstack |
| Performance overhead | Medium | Caching, lazy loading of types |
| Inconsistent type definitions | Medium | JSON schema validation, strict type checking |

## Timeline

| Week | Phase | Deliverables |
|-------|--------|--------------|
| 1 | Foundation | Detection engine, 3 project types (CLI, web, library), 2 overlays |
| 2 | Core | 3 more project types (desktop, serverless, infra), context refactor |
| 3 | Skills | 5 new skills, overlays system complete |
| 4 | Testing | Full test suite, edge case handling |
| 5 | Polish | Documentation, migration guide, release prep |

## Definition of Done

- All 6+ project types implemented and tested
- All 4+ overlays implemented and tested
- `/context` command uses new system
- All new skills implemented
- Test coverage >= 80% for new code
- Documentation updated
- Migration guide created
- Backward compatibility verified with existing projects

## Related

- [[docs/sprint/Sprint-06/HUB-SPRINT-06.md]]
- [[docs/sprint/Sprint-06/BACKLOG.md]]
- [[docs/sprint/Sprint-06/TASKS.md]]
- [[docs/requirements/remove-blueprints-references/HUB-REMOVE-BLUEPRINTS-REFERENCES.md]]
