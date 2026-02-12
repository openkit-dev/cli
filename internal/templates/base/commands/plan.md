---
description: Create structured task plans with INPUT→OUTPUT→VERIFY criteria
subtask: false
---

**PLANNING MODE ACTIVATED**

**User task:** $ARGUMENTS

**If $ARGUMENTS is empty:**
Use the question tool to ask: "What would you like to plan? Describe the task in detail (e.g., 'create a JWT auth system', 'add dark mode', etc.)"

**Mandatory protocol:**

### Phase 0: TodoList Setup (MANDATORY)

Before any planning work:

1. **Read existing todolist:**
   ```javascript
   todoread()
   ```

2. **Create planning todolist (using standard ID schema):**
   ```javascript
   todowrite({
     todos: [
       {
         id: "sdd-plan-01-gate",
         content: "SDD Gate check (verify /specify ran)",
         status: "pending",
         priority: "high"
       },
       {
         id: "sdd-plan-02-stack",
         content: "Select technology stack (if new project)",
         status: "pending",
         priority: "high"
       },
       {
         id: "sdd-plan-03-clarify",
         content: "Socratic gate: clarify requirements",
         status: "pending",
         priority: "high"
       },
       {
         id: "sdd-plan-04-analyze",
         content: "Analyze project type and identify agents",
         status: "pending",
         priority: "high"
       },
       {
         id: "sdd-plan-05-requirements",
         content: "Create PLAN.md in docs/requirements/",
         status: "pending",
         priority: "high"
       },
       {
         id: "sdd-plan-06-sprint",
         content: "Create sprint artifacts (GOAL, BACKLOG, RISKS)",
         status: "pending",
         priority: "high"
       }
     ]
   })
   ```

3. **Update todolist** as you complete each step (mark `in_progress` → `completed`)

### Phase 1: Stack Selection (NEW PROJECTS ONLY)

**Check if this is a new project (no existing code):**

1. **Look for existing stack:**
   - Check for `package.json`, `pyproject.toml`, `go.mod`, `Dockerfile`, etc.
   - If exists → Skip stack selection, maintain consistency

2. **For new projects, use stack-selection skill:**
   ```javascript
   question({
     questions: [{
         question: "Do you have any tech stack preferences?",
         header: "Stack Preferences",
         options: [
           {
             label: "I have specific preferences",
             description: "I'll tell you exactly what tech stack I want"
           },
           {
             label: "Suggest based on requirements",
             description: "Help me choose based on what I'm building"
           }
         ]
       }]
   })
   ```

3. **Document stack decision** in `docs/requirements/<feature>/TECH_STACK.md`:
   - Frontend stack (framework, build tool, state management, UI library, styling)
   - Backend stack (runtime, framework, database, ORM)
   - Infrastructure (containerization, deployment)
   - Rationale for each choice

### Phase 1.5: SDD Gate Check (MANDATORY)

Before creating planning artifacts, verify if `/specify` artifacts exist:

1. **Check for existing specification:**
   - Look for `docs/requirements/<feature>/PROBLEM_STATEMENT.md`
   - Look for `docs/requirements/<feature>/USER_STORIES.md`
   
2. **If specification exists:**
   - Read existing artifacts to understand scope
   - Skip redundant Socratic questions already answered
   - Proceed to Phase 2

3. **If specification is MISSING:**
   - Use the question tool:
   ```javascript
   question({
     questions: [{
       header: "Specification Required",
       question: "No specification found for this feature. The /specify command creates the foundation (Problem Statement, User Stories, Acceptance Criteria). How would you like to proceed?",
       options: [
         { label: "Run /specify first (Recommended)", description: "Create specification before planning" },
         { label: "Create minimal spec inline", description: "I'll ask key questions and create basic spec" }
       ]
     }]
   })
   ```

### Phase 2: Socratic Gate
- Use the question tool to ask clarifying questions if needed:
  1. What is the main goal.
  2. Any technology preferences (if not set in Phase 1).
  3. Any constraints or priorities.

### Phase 2.5: Analysis
- Determine project type: WEB / MOBILE / BACKEND
- Identify required agents
- List core components

### Phase 3: Plan Creation
Ensure `docs/` exists and follow the planning standard:

**Before creating artifacts:**
- Update todolist: Mark "plan-03-analyze" as `completed`
- Mark "plan-04-requirements" as `in_progress`

1. **Discovery/Requirements (MANDATORY)**
    - Create/update `docs/requirements/<feature>/` with:
      `PROBLEM_STATEMENT.md`, `USER_STORIES.md`, `ACCEPTANCE_CRITERIA.md`,
      `DATA_CONTRACTS.md`, `RISKS.md`, `PLAN.md`, `README.md`.

**Required artifacts (SDD Gate):**
- `docs/requirements/<feature>/PLAN.md`
- `docs/requirements/<feature>/PROBLEM_STATEMENT.md`
- `docs/requirements/<feature>/USER_STORIES.md`
- `docs/requirements/<feature>/ACCEPTANCE_CRITERIA.md`
- `docs/requirements/<feature>/RISKS.md`

**Auxiliary artifacts (create when applicable):**
- `docs/requirements/<feature>/RESEARCH.md`
- `docs/requirements/<feature>/QUICKSTART.md`
- `docs/requirements/<feature>/ANALYSIS.md`
- `docs/requirements/<feature>/TECH_STACK.md`
- `docs/requirements/<feature>/CHECKLIST.md`

**When to create auxiliary artifacts:**
- Create `RESEARCH.md` when there are unresolved unknowns or external dependencies.
- Update `DATA_CONTRACTS.md` when the feature introduces or changes data entities.
- Create `QUICKSTART.md` when setup or verification needs explicit steps.
- Create `ANALYSIS.md` for deep cross-artifact checks when needed.
- Create `CHECKLIST.md` for feature-level readiness checks when needed.

**Contracts template usage:**
- Use `.opencode/templates/SDD-Contracts.md` as the base content for `docs/requirements/<feature>/DATA_CONTRACTS.md`.

**Templates (REQUIRED):**
- `.opencode/templates/SDD-ProblemStatement.md`
- `.opencode/templates/SDD-UserStories.md`
- `.opencode/templates/SDD-AcceptanceCriteria.md`
- `.opencode/templates/SDD-Risks.md`
- `.opencode/templates/SDD-Plan.md`
- `.opencode/templates/SDD-RequirementsIndex.md` (for `docs/requirements/README.md`)
- `.opencode/templates/SDD-FeatureIndex.md` (for `docs/requirements/<feature>/README.md`)
- `.opencode/templates/SDD-Research.md` (if applicable)
- `.opencode/templates/SDD-Contracts.md` (if applicable)
- `.opencode/templates/SDD-Quickstart.md` (if applicable)
- `.opencode/templates/TechStack.md` (for `TECH_STACK.md`, if applicable)
- `.opencode/templates/SDD-SprintIndex.md` (for `docs/sprint/Sprint-XX/README.md`)

**After requirements:**
- Update todolist: Mark "plan-04-requirements" as `completed`
- Mark "plan-05-sprint" as `in_progress`

2. **Sprint (MANDATORY)**
    - Use the question tool to ask the user:
      "Do you want to use the latest sprint or create a new sprint?"
    - If no sprint exists, create `docs/sprint/Sprint-01/`.
    - If creating a new sprint, use the next sequential number.
    - Update `README.md`, `SPRINT_GOAL.md`, `BACKLOG.md`, `RISK_REGISTER.md`.
    - **NOTE:** DO NOT create `TASKS.md` here. Use `/tasks` command for detailed task breakdown.

**After sprint creation:**
- Update todolist: Mark "plan-05-sprint" as `completed`

** ABSOLUTE RULES:**
-  NEVER write code during planning
-  NEVER use generic names like `plan.md`
-  ALWAYS use `docs/` as the source of truth
-  ALWAYS create `docs/` if missing
-  ALWAYS use Obsidian-compatible links (`[[docs/...]]`) across plan artifacts
-  ALWAYS add `## Related` sections for cross-doc navigation

### MANDATORY STOP POINT

After creating/updating the artifacts:

Use the question tool:
```javascript
question({
  questions: [{
    header: "Plan Complete",
    question: "Plan recorded in docs/requirements/ and docs/sprint/Sprint-XX/. What would you like to do next?",
    options: [
      { label: "Generate tasks (/tasks)", description: "Create detailed TASKS.md with INPUT->OUTPUT->VERIFY" },
      { label: "Review artifacts first", description: "Let me check the plan before proceeding" },
      { label: "Adjust the plan", description: "Make changes to the plan" }
    ]
  }]
})
```

**DO NOT CONTINUE WITHOUT EXPLICIT USER APPROVAL.**

---

## Workflow Reference (Migrated)

# /plan - Unified Planning

$ARGUMENTS

---

## Overview

This command routes to the appropriate planning agent based on the scope:
1. **Sprint Planning**: Full sprint breakdown (Stories, Epics, Estimations).
2. **Task Planning**: specific feature or fix breakdown (Steps, Files).

---

## Scenario A: Sprint Planning (Complex/Multi-story)

**Trigger:** User asks for "Sprint Planning", "Backlog", or provides valid context for a full sprint.

### Workflow
1. **Agent**: `project-planner`
2. **Action**: Create/Update `docs/sprint/Sprint-XX/` artifacts.
3. **Steps**:
    - Check current date/sprint context.
    - Use the question tool to ask minimal missing inputs (<=5).
    - Produce `SPRINT_GOAL.md`, `BACKLOG.md`, `RISK_REGISTER.md`.
    - Do not create `TASKS.md` here; `/tasks` owns that artifact.
4. **Output**: Full set of sprint docs.

---

## Scenario B: Task Planning (Feature/Fix)

**Trigger:** User asks to "plan feature X", "fix bug Y", or "plan architecture for Z".

### Workflow
1. **Agent**: `project-planner`
2. **Action**: Create/update `docs/requirements/<feature>/` and align a sprint.
3. **Steps**:
    - **Socratic Gate**: Use the question tool to ask clarifying questions first.
    - **Context Check**: Verify existing code/docs.
    - **Plan Generation**: Create requirements and sprint artifacts.
    - **Verify Standard**: Check against `rules/MASTER.md` Documentation Integrity Protocol.
4. **Naming Rules**:
    - Feature folders use kebab-case (e.g., `auth-fix`, `cart-feature`).
    - Sprint folders use `Sprint-XX` with zero padding.
    - Documentation files use uppercase snake case (`TUDO_MAIUSCULO.md`).

---

## Critical Rules (All Modes)

1. **NO CODE WRITING**: This command ONLY generates documentation/plans.
2. **Ask Before Planning**: Use the question tool to clear ambiguities.
3. **File Locations**:
    - Requirements -> `docs/requirements/<feature>/`
    - Sprint -> `docs/sprint/Sprint-XX/`

---

## Usage Examples

```bash
/plan sprint 16 planning
/plan feature user profile
/plan fix login bug
/plan architecture for notifications
```
