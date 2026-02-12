# TASKS

**Sprint**: Sprint-06
**Title**: Context-Aware OpenKit Refactoration
**Created**: 2026-02-10
**Status**: Planning
**Total Tasks**: 33

## Task Breakdown

### P0: Foundation (Week 1)

#### Task 1: Create Detection Engine Go Package

**ID**: S06-T001
**Story**: S06-001
**Priority**: P0
**Assigned**: backend-specialist
**Estimate**: 3 days

**INPUT**:
- Existing project structure (internal/detection/ may not exist)
- Brainstorm Option C architecture document
- Go 1.25.7
- Existing pattern matching from internal/targets/

**OUTPUT**:
- `internal/detection/detector.go` - Main detection logic
- `internal/detection/types.go` - Type definitions (ProjectType, Overlay, DetectionResult)
- `internal/detection/heuristics.go` - Heuristic rules
- `internal/detection/cache.go` - Caching mechanism
- Unit tests in `internal/detection/*_test.go`

**VERIFY**:
- [ ] detector.go exports Detect() function
- [ ] types.go defines ProjectType, Overlay, DetectionResult structs
- [ ] Heuristics.go has 6+ heuristic rules
- [ ] Unit tests cover all detection paths
- [ ] Cache implementation reduces redundant scans
- [ ] go test ./internal/detection/ passes

---

#### Task 2: Define JSON Schemas for Project Types

**ID**: S06-T002
**Story**: S06-002
**Priority**: P0
**Assigned**: database-architect
**Estimate**: 2 days

**INPUT**:
- Brainstorm Option C project type definitions
- Existing opencode.json schema format
- JSON Schema draft standards

**OUTPUT**:
- `.opencode/schemas/project-type.v1.json` - Schema definition
- `.opencode/schemas/overlay.v1.json` - Schema definition
- Documentation of schema fields and validation rules
- Schema validation tests

**VERIFY**:
- [ ] project-type.v1.json defines required fields (id, detection, base-context)
- [ ] overlay.v1.json defines required fields (id, condition, adds)
- [ ] Schemas validate with jsonschema library
- [ ] Test cases cover valid and invalid configs
- [ ] Documentation explains each field

---

#### Task 3: Define JSON Schemas for Overlays

**ID**: S06-T003
**Story**: S06-003
**Priority**: P0
**Assigned**: database-architect
**Estimate**: 1 day

**INPUT**:
- Brainstorm Option C overlay definitions
- JSON Schema from Task S06-T002
- Example overlays (testing, security, ci-cd)

**OUTPUT**:
- Updated overlay.v1.json with complex conditions (any, all, has-file)
- Example overlay configurations for each type
- Condition matching tests

**VERIFY**:
- [ ] overlay.v1.json supports complex conditions
- [ ] Test overlays with various conditions match correctly
- [ ] Example configs are valid against schema
- [ ] Condition logic is documented

---

#### Task 4: Create Directory Structure for Project Types

**ID**: S06-T004
**Story**: S06-004
**Priority**: P0
**Assigned**: backend-specialist
**Estimate**: 1 day

**INPUT**:
- Brainstorm Option C directory structure
- Existing .opencode/templates/ structure
- Go embed pattern from internal/templates/

**OUTPUT**:
- `.opencode/project-types/` directory created
- `.opencode/project-types/base.json` - Common fields
- `.opencode/overlays/` directory created
- Directory structure documentation
- Go embed code updated to include new directories

**VERIFY**:
- [ ] .opencode/project-types/ exists with subdirectories
- [ ] .opencode/overlays/ exists
- [ ] internal/templates/embed.go includes new paths
- [ ] go build succeeds with new embedded files
- [ ] Directory structure matches brainstorm design

---

### P1: Core Functionality (Week 1-2)

#### Task 5: Implement CLI Tool Project Type

**ID**: S06-T005
**Story**: S06-005
**Priority**: P1
**Assigned**: backend-specialist
**Estimate**: 2 days

**INPUT**:
- Brainstorm CLI tool detection patterns
- OpenKit project structure (cmd/, go.mod)
- JSON schema from Task S06-T002

**OUTPUT**:
- `.opencode/project-types/cli-tool.json` - Full type definition
- Documentation of CLI-specific context docs
- CLI-specific skill exclusions
- Detection tests for CLI tools

**VERIFY**:
- [ ] cli-tool.json is valid against schema
- [ ] Detection matches cmd/ + go.mod patterns
- [ ] Context docs exclude FRONTEND.md, API.md
- [ ] Skills exclude frontend-design, mobile-design
- [ ] Tests verify detection of OpenKit itself

---

#### Task 6: Implement Web-Fullstack Project Type

**ID**: S06-T006
**Story**: S06-006
**Priority**: P1
**Assigned**: backend-specialist
**Estimate**: 1 day

**INPUT**:
- Existing web project handling
- Brainstorm web-fullstack definition
- Current backend/frontend docs

**OUTPUT**:
- `.opencode/project-types/web-fullstack.json` - Refactored type
- Backward compatible with existing behavior
- Documentation of web-specific context

**VERIFY**:
- [ ] web-fullstack.json generates BACKEND.md and FRONTEND.md
- [ ] Detection matches frontend/ + backend/ patterns
- [ ] Skills include frontend-specialist, backend-specialist
- [ ] Existing web projects still work correctly
- [ ] Backward compatibility tests pass

---

#### Task 7: Implement Library Project Type

**ID**: S06-T007
**Story**: S06-007
**Priority**: P1
**Assigned**: backend-specialist
**Estimate**: 2 days

**INPUT**:
- Brainstorm library structure (Cargo.toml, src/lib, main: "index")
- Library project patterns
- JSON schema

**OUTPUT**:
- `.opencode/project-types/library.json` - Full type definition
- Library-specific context docs (PUBLIC_API.md, VERSIONING.md)
- Library-specific skills (api-design, semantic-versioning)
- Detection tests

**VERIFY**:
- [ ] library.json is valid against schema
- [ ] Detection matches Cargo.toml, go.mod, or package.json with library markers
- [ ] Context docs include PUBLIC_API.md
- [ ] Skills include api-design skill
- [ ] Tests with real library projects pass

---

#### Task 8: Implement Desktop-App Project Type

**ID**: S06-T008
**Story**: S06-008
**Priority**: P1
**Assigned**: backend-specialist
**Estimate**: 2 days

**INPUT**:
- Brainstorm desktop app structure (electron/, tauri.conf.json)
- Desktop-specific patterns
- JSON schema

**OUTPUT**:
- `.opencode/project-types/desktop-app.json` - Full type definition
- Desktop-specific context docs (DESKTOP_ARCHITECTURE.md, UPDATE_MECHANISM.md)
- Desktop-specific skills (desktop-patterns)
- Detection tests

**VERIFY**:
- [ ] desktop-app.json is valid against schema
- [ ] Detection matches electron/ or tauri.conf.json patterns
- [ ] Context docs include DESKTOP_ARCHITECTURE.md
- [ ] Skills include desktop-patterns skill
- [ ] Tests with Electron/Tauri projects pass

---

#### Task 9: Implement Serverless Project Type

**ID**: S06-T009
**Story**: S06-009
**Priority**: P1
**Assigned**: backend-specialist
**Estimate**: 2 days

**INPUT**:
- Brainstorm serverless structure (functions/, netlify/, vercel/)
- Serverless-specific patterns
- JSON schema

**OUTPUT**:
- `.opencode/project-types/serverless.json` - Full type definition
- Serverless-specific context docs (FUNCTIONS.md, COLD_BOOT.md)
- Serverless-specific skills (serverless-patterns)
- Detection tests

**VERIFY**:
- [ ] serverless.json is valid against schema
- [ ] Detection matches functions/ or serverless config files
- [ ] Context docs include FUNCTIONS.md
- [ ] Skills include serverless-patterns skill
- [ ] Tests with Vercel/Netlify functions pass

---

#### Task 10: Implement Infrastructure Project Type

**ID**: S06-T010
**Story**: S06-010
**Priority**: P1
**Assigned**: backend-specialist
**Estimate**: 2 days

**INPUT**:
- Brainstorm infrastructure structure (*.tf, docker-compose.yml, k8s/)
- Infrastructure-specific patterns
- JSON schema

**OUTPUT**:
- `.opencode/project-types/infrastructure.json` - Full type definition
- Infrastructure-specific context docs (IAC_STRUCTURE.md, DEPLOYMENT.md)
- Infrastructure-specific skills (iac-patterns)
- Detection tests

**VERIFY**:
- [ ] infrastructure.json is valid against schema
- [ ] Detection matches *.tf, docker-compose.yml, k8s/ patterns
- [ ] Context docs include IAC_STRUCTURE.md
- [ ] Skills include iac-patterns skill
- [ ] Tests with Terraform/K8s projects pass

---

#### Task 11: Implement Plugin-Extension Project Type

**ID**: S06-T011
**Story**: S06-011
**Priority**: P1
**Assigned**: backend-specialist
**Estimate**: 2 days

**INPUT**:
- Brainstorm plugin structure (manifest.json, extension.ts)
- Plugin-specific patterns
- JSON schema

**OUTPUT**:
- `.opencode/project-types/plugin-extension.json` - Full type definition
- Plugin-specific context docs (EXTENSION_API.md, HOST_API.md, SANDBOX.md)
- Plugin-specific skills (no new skills, use existing)
- Detection tests

**VERIFY**:
- [ ] plugin-extension.json is valid against schema
- [ ] Detection matches manifest.json or extension.ts patterns
- [ ] Context docs include EXTENSION_API.md
- [ ] Tests with Chrome/VSCode extensions pass

---

#### Task 12: Implement Testing Overlay

**ID**: S06-T012
**Story**: S06-012
**Priority**: P1
**Assigned**: backend-specialist
**Estimate**: 1 day

**INPUT**:
- Brainstorm testing overlay definition
- JSON overlay schema from Task S06-T003
- Testing patterns (_test.go, .test.ts, tests/)

**OUTPUT**:
- `.opencode/overlays/testing-overlay.json` - Full overlay definition
- Testing context docs (TESTING.md, TEST_COVERAGE.md)
- Testing skills (testing-patterns, tdd-workflow, webapp-testing)
- Overlay tests

**VERIFY**:
- [ ] testing-overlay.json is valid against schema
- [ ] Condition matches test file patterns
- [ ] Adds TESTING.md and TEST_COVERAGE.md
- [ ] Adds testing-patterns, tdd-workflow, webapp-testing skills
- [ ] Tests detect and activate overlay correctly

---

#### Task 13: Implement Security Overlay

**ID**: S06-T013
**Story**: S06-013
**Priority**: P1
**Assigned**: security-auditor
**Estimate**: 1 day

**INPUT**:
- Brainstorm security overlay definition
- Security patterns (auth, secrets, .env files)
- JSON overlay schema

**OUTPUT**:
- `.opencode/overlays/security-overlay.json` - Full overlay definition
- Security context docs (SECURITY_AUDIT.md)
- Security skills (penetration-testing, vulnerability-scanner)
- Security condition logic

**VERIFY**:
- [ ] security-overlay.json is valid against schema
- [ ] Condition matches auth or secret usage patterns
- [ ] Adds SECURITY_AUDIT.md
- [ ] Adds penetration-testing, vulnerability-scanner skills
- [ ] Tests verify overlay activation

---

#### Task 14: Implement CI-CD Overlay

**ID**: S06-T014
**Story**: S06-014
**Priority**: P1
**Assigned**: devops-engineer
**Estimate**: 1 day

**INPUT**:
- Brainstorm ci-cd overlay definition
- CI patterns (.github/workflows/, .gitlab-ci.yml)
- JSON overlay schema

**OUTPUT**:
- `.opencode/overlays/ci-cd-overlay.json` - Full overlay definition
- CI/CD context docs (CI_PIPELINE.md)
- DevOps skills (devops-engineer)
- CI condition logic

**VERIFY**:
- [ ] ci-cd-overlay.json is valid against schema
- [ ] Condition matches GitHub Actions or GitLab CI patterns
- [ ] Adds CI_PIPELINE.md
- [ ] Adds devops-engineer skill
- [ ] Tests verify overlay activation

---

#### Task 15: Implement Documentation Overlay

**ID**: S06-T015
**Story**: S06-015
**Priority**: P1
**Assigned**: documentation-writer
**Estimate**: 1 day

**INPUT**:
- Brainstorm documentation overlay definition
- Documentation patterns (docs/, README.md, extensive comments)
- JSON overlay schema

**OUTPUT**:
- `.opencode/overlays/documentation-overlay.json` - Full overlay definition
- Documentation context docs (DOCS_GUIDE.md)
- Documentation skills (documentation-writer)
- Documentation condition logic

**VERIFY**:
- [ ] documentation-overlay.json is valid against schema
- [ ] Condition matches docs/ directory or README.md size
- [ ] Adds DOCS_GUIDE.md
- [ ] Adds documentation-writer skill
- [ ] Tests verify overlay activation

---

### P1: Integration (Week 2)

#### Task 16: Refactor /context Command to Use New System

**ID**: S06-T016
**Story**: S06-016
**Priority**: P1
**Assigned**: backend-specialist
**Estimate**: 3 days

**INPUT**:
- Existing internal/cli/context.go
- Detection engine from Task S06-T001
- Project types and overlays from Tasks S06-T005 through S06-T015

**OUTPUT**:
- Refactored internal/cli/context.go
- New context generation logic using project types
- Overlay application logic
- Backward compatibility mode (if needed)
- Updated CLI help text

**VERIFY**:
- [ ] /context command detects project type
- [ ] /context generates type-specific docs only
- [ ] /context suggests overlays
- [ ] /context prompts for user confirmation
- [ ] Existing web projects still generate BACKEND.md and FRONTEND.md
- [ ] Integration tests pass

---

#### Task 17: Implement User Confirmation UI for Detected Type

**ID**: S06-T017
**Story**: S06-017
**Priority**: P1
**Assigned**: backend-specialist
**Estimate**: 1 day

**INPUT**:
- Detection result from Task S06-T001
- Brainstorm confirmation UI requirements
- Existing UI patterns in internal/cli/

**OUTPUT**:
- User-friendly confirmation message
- Evidence display (files/patterns matched)
- Option to override detected type
- Option to select different type
- Option to proceed without overlays

**VERIFY**:
- [ ] Confirmation shows detected type clearly
- [ ] Evidence (matched files) is displayed
- [ ] User can override with different type
- [ ] User can select/deselect overlays
- [ ] UI tests verify all paths

---

#### Task 18: Implement Overlay Selection UI

**ID**: S06-T018
**Story**: S06-018
**Priority**: P1
**Assigned**: backend-specialist
**Estimate**: 1 day

**INPUT**:
- Available overlays from Tasks S06-T012 through S06-T015
- Brainstorm overlay selection requirements
- Existing selection patterns

**OUTPUT**:
- Multi-select UI for overlays
- Each overlay shows description and what it adds
- Defaults based on project detection
- Option to skip all overlays

**VERIFY**:
- [ ] UI shows all applicable overlays
- [ ] Each overlay shows what docs/skills it adds
- [ ] Defaults are reasonable
- [ ] User can skip or select specific overlays
- [ ] Integration tests verify selection

---

#### Task 19: Implement Project Type Loading from JSON

**ID**: S06-T019
**Story**: S06-019
**Priority**: P1
**Assigned**: backend-specialist
**Estimate**: 1 day

**INPUT**:
- JSON schemas from Tasks S06-T002, S06-T003
- Project type configs from Tasks S06-T005 through S06-T011
- Go jsonschema library (or similar)

**OUTPUT**:
- JSON loading/unmarshaling code
- Schema validation
- Error handling for invalid configs
- Type conversion to Go structs

**VERIFY**:
- [ ] JSON files load successfully
- [ ] Schema validation catches invalid configs
- [ ] Error messages are clear
- [ ] Invalid configs don't crash /context
- [ ] Unit tests cover invalid config cases

---

#### Task 20: Implement Skill Activation/Deactivation Based on Type

**ID**: S06-T020
**Story**: S06-020
**Priority**: P1
**Assigned**: backend-specialist
**Estimate**: 2 days

**INPUT**:
- Skill system from internal/agents/registry.go
- Project type configs with skill inclusions/exclusions
- Existing skill loading logic

**OUTPUT**:
- Skill filtering logic based on project type
- Overlay skill addition logic
- Skill activation status tracking
- Warning when skills are excluded
- Documentation of active skills

**VERIFY**:
- [ ] CLI tools don't load frontend-design skill
- [ ] Libraries don't load mobile-developer skill
- [ ] Web projects load frontend and backend skills
- [ ] Overlay skills are added when selected
- [ ] User can see which skills are active
- [ ] Tests verify skill activation for each type

---

### P1: New Skills (Week 3)

#### Task 21: Create CLI Design Skill

**ID**: S06-T021
**Story**: S06-021
**Priority**: P1
**Assigned**: frontend-specialist (mentorship), backend-specialist (implementation)
**Estimate**: 2 days

**INPUT**:
- Existing skill format (.opencode/skills/*/
- CLI best practices from brainstorm
- Cobra command patterns

**OUTPUT**:
- `.opencode/skills/cli-design/SKILL.md` - Complete skill
- CLI design patterns (command structure, flags, help)
- CLI UX guidelines (output formatting, error messages)
- Examples from real CLI tools

**VERIFY**:
- [ ] SKILL.md follows skill format
- [ ] Covers command design patterns
- [ ] Covers CLI UX best practices
- [ ] Includes examples
- [ ] Review by frontend-specialist for patterns

---

#### Task 22: Create Library Patterns Skill

**ID**: S06-T022
**Story**: S06-022
**Priority**: P1
**Assigned**: backend-specialist
**Estimate**: 2 days

**INPUT**:
- Existing skill format
- Library design patterns (API stability, versioning, semver)
- Examples from popular libraries

**OUTPUT**:
- `.opencode/skills/library-patterns/SKILL.md` - Complete skill
- API design patterns (public vs internal, deprecation)
- Versioning strategies (semantic versioning, changelog)
- Breaking change guidelines

**VERIFY**:
- [ ] SKILL.md follows skill format
- [ ] Covers API design for libraries
- [ ] Covers semantic versioning
- [ ] Includes examples
- [ ] Matches library project requirements

---

#### Task 23: Create Desktop Patterns Skill

**ID**: S06-T023
**Story**: S06-023
**Priority**: P1
**Assigned**: mobile-developer (mentorship), backend-specialist (implementation)
**Estimate**: 2 days

**INPUT**:
- Existing skill format
- Desktop app patterns (Electron, Tauri, native)
- IPC patterns, auto-update mechanisms

**OUTPUT**:
- `.opencode/skills/desktop-patterns/SKILL.md` - Complete skill
- Desktop architecture patterns
- IPC communication patterns
- Auto-update strategies
- Native integration guidelines

**VERIFY**:
- [ ] SKILL.md follows skill format
- [ ] Covers desktop-specific concerns
- [ ] Includes IPC and update patterns
- [ ] Includes examples
- [ ] Review by mobile-developer for patterns

---

#### Task 24: Create Serverless Patterns Skill

**ID**: S06-T024
**Story**: S06-024
**Priority**: P1
**Assigned**: backend-specialist
**Estimate**: 2 days

**INPUT**:
- Existing skill format
- Serverless patterns (Vercel, Netlify, AWS Lambda)
- Cold start optimization, state management

**OUTPUT**:
- `.opencode/skills/serverless-patterns/SKILL.md` - Complete skill
- Serverless function design patterns
- State management strategies (Redis, DynamoDB)
- Cold start optimization
- Environment variable handling

**VERIFY**:
- [ ] SKILL.md follows skill format
- [ ] Covers serverless-specific concerns
- [ ] Includes cold start and state patterns
- [ ] Includes examples
- [ ] Matches serverless project requirements

---

#### Task 25: Create IaC Patterns Skill

**ID**: S06-T025
**Story**: S06-025
**Priority**: P1
**Assigned**: devops-engineer
**Estimate**: 2 days

**INPUT**:
- Existing skill format
- IaC patterns (Terraform, Docker, K8s)
- Infrastructure as code best practices

**OUTPUT**:
- `.opencode/skills/iac-patterns/SKILL.md` - Complete skill
- Terraform module patterns
- Docker containerization best practices
- Kubernetes deployment patterns
- Environment configuration strategies

**VERIFY**:
- [ ] SKILL.md follows skill format
- [ ] Covers IaC-specific concerns
- [ ] Includes Terraform, Docker, K8s patterns
- [ ] Includes examples
- [ ] Matches infrastructure project requirements

---

### P1: Testing (Week 4)

#### Task 26: Write Unit Tests for Detection Engine

**ID**: S06-T026
**Story**: S06-026
**Priority**: P1
**Assigned**: test-engineer
**Estimate**: 2 days

**INPUT**:
- Detection engine code from Task S06-T001
- Test coverage requirements (target 90%+)
- Test case scenarios

**OUTPUT**:
- `internal/detection/detector_test.go` - Unit tests
- `internal/detection/heuristics_test.go` - Unit tests
- `internal/detection/cache_test.go` - Unit tests
- Test coverage report

**VERIFY**:
- [ ] Unit tests cover all detection functions
- [ ] Tests cover 6+ project type scenarios
- [ ] Tests cover ambiguous cases
- [ ] Tests cover conflict detection
- [ ] Tests cover caching logic
- [ ] Coverage >= 85%

---

#### Task 27: Write Integration Tests for /context Command

**ID**: S06-T027
**Story**: S06-027
**Priority**: P1
**Assigned**: test-engineer
**Estimate**: 2 days

**INPUT**:
- Refactored /context command from Task S06-T016
- Real project structures for testing
- Integration test patterns

**OUTPUT**:
- `internal/cli/context_integration_test.go` - Integration tests
- Test fixture projects (CLI, web, library, desktop, etc.)
- End-to-end test scenarios

**VERIFY**:
- [ ] Tests verify type detection
- [ ] Tests verify doc generation
- [ ] Tests verify skill activation
- [ ] Tests verify overlay application
- [ ] Tests verify user confirmation flow
- [ ] All 6+ project types tested

---

#### Task 28: Test with Real Projects (CLI, Web, Library)

**ID**: S06-T028
**Story**: S06-028
**Priority**: P1
**Assigned**: test-engineer
**Estimate**: 3 days

**INPUT**:
- Real-world projects (OpenKit itself, web apps, libraries)
- Test scenarios for each type
- Existing test patterns

**OUTPUT**:
- Test results for OpenKit (CLI tool)
- Test results for web-fullstack project
- Test results for library project
- Test results for other types (desktop, serverless, infra)
- Bug reports and fixes

**VERIFY**:
- [ ] OpenKit project detected as CLI tool
- [ ] Correct docs generated for each project type
- [ ] Skills activated correctly
- [ ] No irrelevant docs generated
- [ ] No false positives in detection
- [ ] All reported bugs fixed

---

#### Task 29: Test Edge Cases (Ambiguous, Unknown Types)

**ID**: S06-T029
**Story**: S06-029
**Priority**: P1
**Assigned**: test-engineer
**Estimate**: 2 days

**INPUT**:
- Edge cases from risk register (R06-002, R06-007, R06-010)
- Complex project structures (monorepo, nested)
- Ambiguous detection scenarios

**OUTPUT**:
- Test results for monorepo detection
- Test results for ambiguous projects
- Test results for unknown/custom projects
- Edge case handling documentation

**VERIFY**:
- [ ] Monorepos handled correctly
- [ ] Ambiguous projects trigger user confirmation
- [ ] Unknown projects fallback to custom type
- [ ] No crashes on edge cases
- [ ] Clear error messages
- [ ] Edge case documentation complete

---

### P2: Quality & Documentation (Week 4-5)

#### Task 30: Update Project README with New Capabilities

**ID**: S06-T030
**Story**: S06-030
**Priority**: P2
**Assigned**: documentation-writer
**Estimate**: 1 day

**INPUT**:
- New context-aware system documentation
- Existing README.md
- Usage examples for new features

**OUTPUT**:
- Updated README.md
- Project type detection section
- Supported project types list
- New skills documentation
- Usage examples

**VERIFY**:
- [ ] README explains context-aware system
- [ ] All 6+ project types listed
- [ ] New skills mentioned
- [ ] Examples show detection and confirmation
- [ ] README is clear and accurate

---

#### Task 31: Create Migration Guide from Old System

**ID**: S06-T031
**Story**: S06-031
**Priority**: P2
**Assigned**: documentation-writer
**Estimate**: 2 days

**INPUT**:
- Old system documentation (web-focused)
- New system documentation (context-aware)
- Migration considerations

**OUTPUT**:
- `docs/MIGRATION_GUIDE.md` - Complete migration guide
- Before/after comparisons
- Migration steps
- Backward compatibility notes
- FAQ section

**VERIFY**:
- [ ] Guide explains changes from old to new
- [ ] Migration steps are clear
- [ ] Backward compatibility is documented
- [ ] Common questions answered
- [ ] Guide is tested by reviewer

---

#### Task 32: Add Performance Benchmarks

**ID**: S06-T032
**Story**: S06-032
**Priority**: P2
**Assigned**: performance-optimizer
**Estimate**: 1 day

**INPUT**:
- Detection engine code
- Caching implementation
- Benchmarking patterns

**OUTPUT**:
- Benchmark tests for detection
- Benchmark tests for context generation
- Performance report
- Optimization recommendations (if needed)

**VERIFY**:
- [ ] Benchmarks measure detection time
- [ ] Benchmarks measure doc generation time
- [ ] Performance is acceptable (< 2s for medium projects)
- [ ] Caching effectiveness is measured
- [ ] Report includes before/after comparison

---

#### Task 33: Add Detection Cache

**ID**: S06-T033
**Story**: S06-033
**Priority**: P2
**Assigned**: performance-optimizer
**Estimate**: 1 day

**INPUT**:
- Detection code from Task S06-T001
- Cache requirements (per project, TTL)
- Existing caching patterns in Go

**OUTPUT**:
- Cache implementation in internal/detection/cache.go
- Cache invalidation logic
- Performance tests
- Cache documentation

**VERIFY**:
- [ ] Cache reduces redundant scans
- [ ] Cache invalidates on file changes
- [ ] Cache has configurable TTL
- [ ] Cache doesn't cause stale results
- [ ] Performance tests show improvement

---

## Task Dependencies

```
S06-T001 (Detection Engine)
  ├─ S06-T002, S06-T003 (Schemas)
  ├─ S06-T004 (Directory Structure)
  ├─ S06-T005 through S06-T015 (Project Types & Overlays)
  │
  ├─ S06-T016 (/context Refactor)
  │   ├─ S06-T017 (Confirmation UI)
  │   ├─ S06-T018 (Overlay Selection)
  │   ├─ S06-T019 (Type Loading)
  │   └─ S06-T020 (Skill Activation)
  │
  ├─ S06-T021 through S06-T025 (New Skills)
  │
  └─ S06-T026 through S06-T029 (Testing)
      └─ S06-T030 through S06-T033 (Quality)
```

## Progress Tracking

| Phase | Tasks | Complete | Progress |
|-------|--------|----------|----------|
| **P0: Foundation** | 4 | 4/4 | 100% |
| **P1: Core** | 11 | 11/11 | 100% |
| **P1: Integration** | 5 | 5/5 | 100% |
| **P1: New Skills** | 5 | 5/5 | 100% |
| **P1: Testing** | 4 | 4/4 | 100% |
| **P2: Quality** | 4 | 2/4 | 50% |
| **TOTAL** | 33 | 31/33 | 94% |

## Acceptance Criteria

### Sprint Success

- [x] All 6+ project types implemented and documented
- [x] All 4+ overlays implemented and documented
- [x] /context command uses new system
- [x] Detection works correctly for CLI, web, library, desktop, serverless, infra
- [x] User confirmation flow implemented and tested
- [x] All 5 new skills created
- [x] Test coverage >= 80% for new code
- [ ] Performance benchmarks acceptable (< 2s detection)
- [x] Documentation complete (README, migration guide)
- [x] Backward compatibility verified
- [x] Zero P0 bugs remaining

### Definition of Done

A task is "Done" when:
- Code is implemented and follows Go best practices
- Unit tests pass (go test ./internal/...)
- Code is reviewed and approved
- Task is marked complete in BACKLOG.md
- Sprint tasks reflect completion
- Documentation is updated (if applicable)

A sprint is "Done" when:
- All P0 tasks are complete
- All P1 tasks are complete
- P2 tasks are complete or deferred with documentation
- Sprint goal is achieved
- Sprint retrospective is completed
- Next sprint is planned (if needed)

## Related

- [[sprint/Sprint-06/HUB-SPRINT-06.md]]
- [[sprint/Sprint-06/SPRINT_GOAL.md]]
- [[sprint/Sprint-06/BACKLOG.md]]
- [[requirements/security-scan-hardening/PLAN.md]]
