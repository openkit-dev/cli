---
name: documentation-templates
description: Documentation templates and structure guidelines. README, API docs, code comments, and AI-friendly documentation.
allowed-tools: Read, Glob, Grep
---

# Documentation Templates

> Templates and structure guidelines for common documentation types.

---

## 0. Obsidian Compatibility (Mandatory)

When writing docs in `docs/`, use Obsidian-compatible internal references:

- Prefer wikilinks for internal docs:
  - `[[HUB-DOCS.md]]`
  - `[[requirements/<feature>/PLAN.md]]`
- Use standard Markdown links for external URLs.
- Add `## Related` section to connect artifacts.
- For hub docs (`HUB-*.md`), add:
  - `## Context` (what the hub covers)
  - `## Status Overview` (current phase, completion, blockers)

```markdown
## Related

- [[HUB-DOCS.md]]
- [[GLOSSARY.md]]
```

---

## 1. README Structure

### Essential Sections (Priority Order)

| Section | Purpose |
|---------|---------|
| **Title + One-liner** | What is this? |
| **Quick Start** | Running in <5 min |
| **Features** | What can I do? |
| **Configuration** | How to customize |
| **API Reference** | Link to detailed docs |
| **Contributing** | How to help |
| **License** | Legal |

### README Template

```markdown
# Project Name

Brief one-line description.

## Quick Start

[Minimum steps to run]

## Features

- Feature 1
- Feature 2

## Configuration

| Variable | Description | Default |
|----------|-------------|---------|
| PORT | Server port | 3000 |

## Documentation

- [API Reference](./docs/API.md)
- [Architecture](./docs/ARCHITECTURE.md)

## License

MIT
```

---

## 2. API Documentation Structure

### Per-Endpoint Template

```markdown
## GET /users/:id

Get a user by ID.

**Parameters:**
| Name | Type | Required | Description |
|------|------|----------|-------------|
| id | string | Yes | User ID |

**Response:**
- 200: User object
- 404: User not found

**Example:**
[Request and response example]
```

---

## 3. Code Comment Guidelines

### JSDoc/TSDoc Template

```typescript
/**
 * Brief description of what the function does.
 * 
 * @param paramName - Description of parameter
 * @returns Description of return value
 * @throws ErrorType - When this error occurs
 * 
 * @example
 * const result = functionName(input);
 */
```

### When to Comment

|  Comment |  Don't Comment |
|-----------|-----------------|
| Why (business logic) | What (obvious) |
| Complex algorithms | Every line |
| Non-obvious behavior | Self-explanatory code |
| API contracts | Implementation details |

---

## 4. Changelog Template (Keep a Changelog)

```markdown
# Changelog

## [Unreleased]
### Added
- New feature

## [1.0.0] - 2025-01-01
### Added
- Initial release
### Changed
- Updated dependency
### Fixed
- Bug fix
```

---

## 5. Architecture Decision Record (ADR)

```markdown
# ADR-001: [Title]

## Status
Accepted / Deprecated / Superseded

## Context
Why are we making this decision?

## Decision
What did we decide?

## Consequences
What are the trade-offs?
```

---

## 6. AI-Friendly Documentation (2025)

### llms.txt Template

For AI crawlers and agents:

```markdown
# Project Name
> One-line objective.

## Core Files
- [src/index.ts]: Main entry
- [src/api/]: API routes
- [docs/]: Documentation

## Key Concepts
- Concept 1: Brief explanation
- Concept 2: Brief explanation
```

### MCP-Ready Documentation

For RAG indexing:
- Clear H1-H3 hierarchy
- JSON/YAML examples for data structures
- Mermaid diagrams for flows
- Self-contained sections

---

## 7. Planning Artifacts in docs/

### Requirements (docs/requirements/<feature>/)

```markdown
# Problem Statement — [Feature]

## The Problem
[Context and user pain]

## The Objective
[Measurable objective]

## Scope

### In-Scope
- [ ] Item

### Out-of-Scope
- [ ] Item

## Success Criteria (KPIs)
- [ ] Indicator
```

```markdown
# User Stories — [Feature]

| ID | Persona | I want... (Action) | So that... (Value) | Priority |
|----|---------|--------------------|-------------------|----------|
```

```markdown
# Acceptance Criteria (Gherkin)

## US-01: [Title]
```gherkin
Scenario: [Nome]
  Given ...
  When ...
  Then ...
```
```

```markdown
# Data Contracts — [Feature]

## Entities
| Field | Type | Required | Description |
|-------|------|----------|-------------|
```

```markdown
# Risks — [Feature]

| Risk | Probability | Impact | Mitigation |
|------|-------------|--------|------------|
```

### Sprint Planning (docs/sprint/Sprint-XX/)

```markdown
# Sprint XX Goal

## Objective
[Sprint objective]

## Key Deliverables
1. [Deliverable]
```

```markdown
# Sprint XX Backlog

| ID | Story | Epic | Priority | Points | Dependencies | Owner | Status |
|----|-------|------|----------|--------|--------------|-------|--------|
```

```markdown
## SXX-001: [Title]
Owner: [owner]

- [ ] Task
```

## 8. Structure Principles

| Principle | Why |
|-----------|-----|
| **Scannable** | Headers, lists, tables |
| **Examples first** | Show, don't just tell |
| **Progressive detail** | Simple → Complex |
| **Up to date** | Outdated = misleading |

---

> **Remember:** Templates are starting points. Adapt to your project's needs.
