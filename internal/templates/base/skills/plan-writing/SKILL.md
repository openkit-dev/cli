---
name: plan-writing
description: Structured task planning with clear breakdowns, dependencies, and verification criteria. Use when implementing features, refactoring, or any multi-step work.
allowed-tools: Read, Glob, Grep
---

# Plan Writing

> Source: obra/superpowers

## Overview
This skill provides a framework for breaking down work into clear, actionable tasks with verification criteria.

## Task Breakdown Principles

### 1. Small, Focused Tasks
- Each task should take 2-5 minutes
- One clear outcome per task
- Independently verifiable

### 2. Clear Verification
- How do you know it's done?
- What can you check/test?
- What's expected output?

### 3. Logical Ordering
- Dependencies identified
- Parallel work where possible
- Critical path highlighted
 - **Phase 3: Verification is always LAST**

### 4. Planning Lives in memory/

### 5. Question Tool Usage (MANDATORY)

When planning requires user input or decisions:
- Use `question` tool for all multi-option scenarios
- Provide clear options with descriptions
- For tech stack decisions, structure choices properly

**Example:**
```javascript
question({
  questions: [{
      question: "What tech stack?",
      header: "Stack",
      options: [
        { label: "React + FastAPI", description: "Full-stack modern" },
        { label: "Next.js + Node", description: "All-in-one" }
      ]
    }]
})
```

See `.opencode/rules/MASTER.md` for complete Question Tool Protocol.
- Planning artifacts live under `memory/requirements/` and `memory/sprint/`
- **Discovery** outputs go to `memory/requirements/<feature>/`
- **Sprint planning** outputs go to `memory/sprint/Sprint-XX/`
- If `memory/` does not exist, create it before writing

## Planning Principles (NOT Templates!)

>  **NO fixed templates. Each plan is UNIQUE to the task.**

### Principle 1: Keep It SHORT

|  Wrong |  Right |
|----------|----------|
| 50 tasks with sub-sub-tasks | 5-10 clear tasks max |
| Every micro-step listed | Only actionable items |
| Verbose descriptions | One-line per task |

> **Rule:** If plan is longer than 1 page, it's too long. Simplify.

---

### Principle 2: Be SPECIFIC, Not Generic

|  Wrong |  Right |
|----------|----------|
| "Set up project" | "Run `npx create-next-app`" |
| "Add authentication" | "Install next-auth, create `/api/auth/[...nextauth].ts`" |
| "Style the UI" | "Add Tailwind classes to `Header.tsx`" |

> **Rule:** Each task should have a clear, verifiable outcome.

---

### Principle 3: Dynamic Content Based on Project Type

**For NEW PROJECT:**
- What tech stack? (decide first)
- What's the MVP? (minimal features)
- What's the file structure?

**For FEATURE ADDITION:**
- Which files are affected?
- What dependencies needed?
- How to verify it works?

**For BUG FIX:**
- What's the root cause?
- What file/line to change?
- How to test the fix?

---

### Principle 4: Scripts Are Project-Specific

>  **DO NOT copy-paste script commands. Choose based on project type.**

| Project Type | Relevant Scripts |
|--------------|------------------|
| Frontend/React | `ux_audit.py`, `accessibility_checker.py` |
| Backend/API | `api_validator.py`, `security_scan.py` |
| Mobile | `mobile_audit.py` |
| Database | `schema_validator.py` |
| Full-stack | Mix of above based on what you touched |

**Wrong:** Adding all scripts to every plan
**Right:** Only scripts relevant to THIS task

---

### Principle 5: Verification is Simple

|  Wrong |  Right |
|----------|----------|
| "Verify the component works correctly" | "Run `npm run dev`, click button, see toast" |
| "Test the API" | "curl localhost:3000/api/users returns 200" |
| "Check styles" | "Open browser, verify dark mode toggle works" |

---

## Plan Structure (Flexible, Not Fixed!)

```
# [Task Name]

## Goal
One sentence: What are we building/fixing?

## Tasks
- [ ] Task 1: [Specific action] → Verify: [How to check]
- [ ] Task 2: [Specific action] → Verify: [How to check]
- [ ] Task 3: [Specific action] → Verify: [How to check]

## Done When
- [ ] [Main success criteria]
```

> **That's it.** No phases, no sub-sections unless truly needed.
> Keep it minimal. Add complexity only when required.

## Notes
[Any important considerations]
```

---

## Best Practices (Quick Reference)

1. **Start with goal** - What are we building/fixing?
2. **Max 10 tasks** - If more, break into multiple plans
3. **Each task verifiable** - Clear "done" criteria
4. **Docs-first** - Requirements in `memory/requirements/`, sprint tasks in `memory/sprint/`
5. **Update as you go** - Mark `[x]` when complete in sprint Tasks

---

## When to Use

- New project from scratch
- Adding a feature
- Fixing a bug (if complex)
- Refactoring multiple files
