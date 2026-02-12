# BACKLOG

**Sprint**: Sprint-06
**Created**: 2026-02-10
**Last Updated**: 2026-02-11

## Priority Definitions

- **P0**: Critical - must complete this sprint
- **P1**: High - target for this sprint
- **P2**: Medium - complete if time permits
- **P3**: Low - backlog for future sprints

## Stories

### Foundation (P0)

| ID | Story | Status | Priority | Estimate |
|----|--------|---------|-----------|
| S06-001 | Create detection engine Go package | Not Started | P0 | 3 days |
| S06-002 | Define JSON schemas for project types | Not Started | P0 | 2 days |
| S06-003 | Define JSON schemas for overlays | Not Started | P0 | 1 day |
| S06-004 | Create directory structure for project types | Not Started | P0 | 1 day |

### Core Functionality (P1)

| ID | Story | Status | Priority | Estimate |
|----|--------|---------|-----------|
| S06-005 | Implement CLI tool project type | Not Started | P1 | 2 days |
| S06-006 | Implement web-fullstack project type | Not Started | P1 | 1 day (refactor existing) |
| S06-007 | Implement library project type | Not Started | P1 | 2 days |
| S06-008 | Implement desktop-app project type | Not Started | P1 | 2 days |
| S06-009 | Implement serverless project type | Not Started | P1 | 2 days |
| S06-010 | Implement infrastructure project type | Not Started | P1 | 2 days |
| S06-011 | Implement plugin-extension project type | Not Started | P1 | 2 days |
| S06-012 | Implement testing overlay | Not Started | P1 | 1 day |
| S06-013 | Implement security overlay | Not Started | P1 | 1 day |
| S06-014 | Implement ci-cd overlay | Not Started | P1 | 1 day |
| S06-015 | Implement documentation overlay | Not Started | P1 | 1 day |

### Integration (P1)

| ID | Story | Status | Priority | Estimate |
|----|--------|---------|-----------|
| S06-016 | Refactor /context command to use new system | Not Started | P1 | 3 days |
| S06-017 | Implement user confirmation UI for detected type | Not Started | P1 | 1 day |
| S06-018 | Implement overlay selection UI | Not Started | P1 | 1 day |
| S06-019 | Implement project type loading from JSON | Not Started | P1 | 1 day |
| S06-020 | Implement skill activation/deactivation based on type | Not Started | P1 | 2 days |

### New Skills (P1)

| ID | Story | Status | Priority | Estimate |
|----|--------|---------|-----------|
| S06-021 | Create cli-design skill | Not Started | P1 | 2 days |
| S06-022 | Create library-patterns skill | Not Started | P1 | 2 days |
| S06-023 | Create desktop-patterns skill | Not Started | P1 | 2 days |
| S06-024 | Create serverless-patterns skill | Not Started | P1 | 2 days |
| S06-025 | Create iac-patterns skill | Not Started | P1 | 2 days |

### Testing (P1)

| ID | Story | Status | Priority | Estimate |
|----|--------|---------|-----------|
| S06-026 | Write unit tests for detection engine | Not Started | P1 | 2 days |
| S06-027 | Write integration tests for /context command | Not Started | P1 | 2 days |
| S06-028 | Test with real projects (CLI, web, library) | Not Started | P1 | 3 days |
| S06-029 | Test edge cases (ambiguous, unknown types) | Not Started | P1 | 2 days |

### Quality & Documentation (P2)

| ID | Story | Status | Priority | Estimate |
|----|--------|---------|-----------|
| S06-030 | Update project README with new capabilities | Not Started | P2 | 1 day |
| S06-031 | Create migration guide from old system | Not Started | P2 | 2 days |
| S06-032 | Add performance benchmarks | Not Started | P2 | 1 day |
| S06-033 | Add detection cache | Not Started | P2 | 1 day |

## Burndown

| Week | P0 Complete | P1 Complete | P2 Complete | Total Points |
|-------|-------------|-------------|-------------|--------------|
| 1 | 0/4 | 0/20 | 0/4 | 0/28 |
| 2 | 4/4 | 8/20 | 2/4 | 14/28 |
| 3 | 4/4 | 15/20 | 3/4 | 22/28 |
| 4 | 4/4 | 20/20 | 4/4 | 28/28 |
| 5 | 4/4 | 20/20 | 4/4 | 28/28 |

## Velocity

**Current**: In progress
**Target**: 6-7 story points per week
**Sprint Capacity**: 28 story points

## Execution Snapshot

- Completed: `S06-001` to `S06-031`
- Pending: `S06-032` (performance benchmarks), `S06-033` (detection cache)
- Quality status: `go test ./...` passing after integration and new skills rollout
