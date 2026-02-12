# RISK REGISTER

**Sprint**: Sprint-06
**Created**: 2026-02-10
**Last Updated**: 2026-02-10

## Risk Assessment Matrix

| ID | Risk | Impact | Probability | Risk Score | Mitigation Strategy | Owner | Status |
|----|-------|--------------|--------------|-------------------|--------|--------|
| R06-001 | False positive project type detection | High | Medium | High | Confidence thresholds, user confirmation required, evidence display | TBD | Open |
| R06-002 | Complex detection logic leads to bugs | High | Medium | High | Clear heuristics, extensive unit tests, integration tests with real projects | TBD | Open |
| R06-003 | Breaking changes to existing /context command | High | Medium | High | Backward compatibility mode, gradual migration, deprecation warnings | TBD | Open |
| R06-004 | Performance overhead from project scanning | Medium | High | Medium | Caching mechanism, lazy loading of type configs, async scanning | TBD | Open |
| R06-005 | Inconsistent project type definitions | Medium | High | High | JSON schema validation, strict type checking, automated linting of type configs | TBD | Open |
| R06-006 | Skills not properly disabled for irrelevant types | Medium | Medium | Medium | Clear skill exclusion rules, testing with each project type, skill metadata | TBD | Open |
| R06-007 | Overlap between project types causes ambiguity | High | Medium | High | Clear priority rules, conflict detection, user prompts for resolution | TBD | Open |
| R06-008 | New skills lack depth compared to existing ones | Medium | Medium | Medium | Leverage existing skill patterns, peer review, user feedback | TBD | Open |
| R06-009 | Schema changes break backward compatibility | High | Low | Medium | Versioned schemas, migration support, deprecation warnings | TBD | Open |
| R06-010 | Edge cases not handled (monorepo, nested structures) | Medium | Medium | Medium | Explicit edge case handling, testing with complex projects | TBD | Open |

## Detailed Risk Analysis

### R06-001: False Positive Project Type Detection

**Description**: Detection engine incorrectly identifies project type (e.g., identifies CLI tool as web app).

**Impact**: High - Wrong documentation and skills loaded, confusion for user

**Probability**: Medium - Heuristics may fail on custom structures

**Mitigation**:
1. Implement confidence scoring (0-100%)
2. Require user confirmation for confidence < 80%
3. Display evidence (files/patterns that matched)
4. Allow user to manually override
5. Conflict detection (patterns that shouldn't coexist)

**Owner**: TBD (backend-specialist)

**Trigger**: Detection implementation (Week 1)

---

### R06-002: Complex Detection Logic Bugs

**Description**: Heuristic-based detection becomes too complex, leading to hard-to-debug issues.

**Impact**: High - Detection fails completely, poor user experience

**Probability**: Medium - Complex systems inherently have more bugs

**Mitigation**:
1. Keep heuristics simple and explicit
2. Extensive unit test coverage (target 90%+)
3. Integration tests with real projects (10+ types)
4. Logging of detection decisions
5. Easy fallback to "custom" type

**Owner**: TBD (backend-specialist, test-engineer)

**Trigger**: Detection implementation (Week 1-2)

---

### R06-003: Breaking Changes to /context Command

**Description**: Refactoring /context command breaks existing user workflows or output format.

**Impact**: High - Breaking change for existing users, migration friction

**Probability**: Medium - Significant refactoring has risk

**Mitigation**:
1. Maintain backward compatibility mode flag
2. Keep existing output format where possible
3. Add deprecation warnings for old behavior
4. Gradual migration with clear documentation
5. Test with existing projects (Sprint 01-05)

**Owner**: TBD (backend-specialist, test-engineer)

**Trigger**: Context refactor (Week 2)

---

### R06-004: Performance Overhead

**Description**: Project scanning on every /context command becomes too slow for large projects.

**Impact**: Medium - Slow command execution, poor UX

**Probability**: High - Scanning many files is inherently expensive

**Mitigation**:
1. Implement detection cache with TTL
2. Lazy load project type configs
3. Async scanning with timeout
4. Limit scan depth (max 3 levels deep)
5. Cache per project directory

**Owner**: TBD (backend-specialist, performance-optimizer)

**Trigger**: Detection implementation (Week 2)

---

### R06-005: Inconsistent Type Definitions

**Description**: Different project types define incompatible or overlapping structures.

**Impact**: Medium - Confusion, maintenance burden, bugs in loading

**Probability**: High - Manual JSON definitions prone to inconsistency

**Mitigation**:
1. JSON schema validation for all type configs
2. Strict type checking in Go code
3. Linting rules for type definitions
4. Automated testing of all types on load
5. Clear documentation of required fields

**Owner**: TBD (backend-specialist)

**Trigger**: Type implementation (Week 1-3)

---

### R06-006: Skills Not Properly Disabled

**Description**: Inappropriate skills remain active for project types where they don't apply.

**Impact**: Medium - Irrelevant suggestions, confusion, noise

**Probability**: Medium - Skill activation logic complex

**Mitigation**:
1. Clear skill exclusion rules in type config
2. Testing with each project type to verify exclusions
3. Skill metadata (compatible project types)
4. Automated verification of active skills
5. User-visible skill activation status

**Owner**: TBD (backend-specialist, test-engineer)

**Trigger**: Skill integration (Week 3)

---

### R06-007: Overlap Between Types Causes Ambiguity

**Description**: Projects match multiple project types with similar scores, unclear which to select.

**Impact**: High - Wrong type selection, poor detection UX

**Probability**: Medium - Custom and hybrid structures exist

**Mitigation**:
1. Clear priority rules (required patterns > suggested)
2. Conflict detection (mutually exclusive patterns)
3. Minimum score threshold for detection
4. User prompt for manual resolution when close
5. Evidence display for each candidate type

**Owner**: TBD (backend-specialist)

**Trigger**: Detection implementation (Week 1-2)

---

### R06-008: New Skills Lack Depth

**Description**: Newly created skills (cli-design, library-patterns, etc.) don't have as much depth as existing ones.

**Impact**: Medium - Lower value for non-web projects, inconsistent quality

**Probability**: Medium - New skills need iteration

**Mitigation**:
1. Study existing frontend/backend skills for patterns
2. Peer review during skill creation
3. User feedback collection after release
4. Iterative improvement in follow-up sprints
5. Leverage common patterns across skills

**Owner**: TBD (project-planner, frontend-specialist for mentorship)

**Trigger**: Skill creation (Week 3)

---

### R06-009: Schema Changes Break Compatibility

**Description**: JSON schema changes for types/overlays break existing custom definitions.

**Impact**: High - Breaking change for users with custom types, migration required

**Probability**: Low - Schemas typically stable after v1

**Mitigation**:
1. Versioned schemas (v1, v2, etc.)
2. Migration support for old schemas
3. Deprecation warnings before breaking
4. Backward compatibility support for at least 2 versions
5. Clear migration documentation

**Owner**: TBD (backend-specialist)

**Trigger**: Schema definition (Week 1)

---

### R06-010: Edge Cases Not Handled

**Description**: Complex structures like monorepos, nested projects not properly detected.

**Impact**: Medium - Detection fails for real-world complex projects

**Probability**: Medium - Edge cases inherently harder

**Mitigation**:
1. Explicit edge case handling (monorepo type)
2. Testing with complex projects (monorepo, nested, custom)
3. Fallback to "custom" type when unclear
4. User prompts for manual type selection
5. Documented handling of known edge cases

**Owner**: TBD (backend-specialist, test-engineer)

**Trigger**: Detection implementation (Week 1-3)

---

## Risk Monitoring

### Weekly Review

During sprint standups, review:
1. Which risks were encountered this week?
2. New risks identified?
3. Mitigation strategies effective?
4. Risk scores updated?

### Sprint Retrospective

At sprint end:
1. Which risks materialized?
2. Which mitigations worked?
3. Which didn't work?
4. Lessons learned for next sprint?

### Risk Score Calculation

**Risk Score** = Impact × Probability

| Score | Severity | Action Required |
|--------|-----------|----------------|
| 0-2.5 | Low | Monitor |
| 3-6.5 | Medium | Plan mitigation |
| 7-12.5 | High | Mitigate now |

## Current Summary

- **Total Risks**: 10
- **High Risk**: 3 (R06-001, R06-003, R06-007)
- **Medium Risk**: 6 (R06-002, R06-004, R06-005, R06-006, R06-008, R06-010)
- **Low Risk**: 1 (R06-009)
- **Open Risks**: 10
- **Closed Risks**: 0
- **Mean Risk Score**: 5.35 (Medium)

## Related

- [[docs/sprint/Sprint-06/HUB-SPRINT-06.md]]
- [[docs/sprint/Sprint-06/SPRINT_GOAL.md]]
- [[docs/sprint/Sprint-06/BACKLOG.md]]
- [[docs/requirements/security-scan-hardening/RISKS.md]]
