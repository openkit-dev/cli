---
description: Universal orchestrator for complex multi-agent missions with 15 specialist agents
subtask: false
---

<!-- Tools are configured in opencode.json under "orchestrator" agent -->
<!-- All tools available: read, grep, glob, list, bash, edit, write, patch, webfetch, skill, todowrite, todoread, question -->

**UNIVERSAL ORCHESTRATOR MODE ACTIVATED**

**User mission:** $ARGUMENTS

**If $ARGUMENTS is empty:**
Use the question tool to ask:
```javascript
question({
  questions: [{
      question: "Describe the complex task you need executed.",
      header: "Task Description"
    }]
})
```

**Examples:**
- 'Build a full e-commerce with Stripe checkout'
- 'Refactor the auth system to OAuth2'
- 'Add real-time chat with WebSockets'
- 'Migrate MySQL to PostgreSQL'

---

## COMPLEXITY ANALYSIS

**Determine the mode automatically:**

### Mode 1: Router (Specific Tasks)
**Triggers:**
- Specific keywords: "specify", "clarify", "plan", "tasks", "implement", "test", "debug", "design", "deploy", "analyze", "checklist"
- Single-domain tasks
- Simple direct actions

**Action:** Route to specialized command
- "specify..." → Trigger `/specify`
- "clarify..." → Trigger `/clarify`
- "plan..." → Trigger `/plan`
- "tasks..." → Trigger `/tasks`
- "analyze..." → Trigger `/analyze`
- "checklist..." → Trigger `/checklist`
- "test..." → Trigger `/test`
- "debug..." → Trigger `/debug`

### Mode 2: Orchestrator (Complex Tasks)
**Triggers:**
- Multiple domains involved (backend + frontend + database + security)
- Keywords: "complete", "full", "system", "application", "platform"
- Large feature or new project
- Does not fit a single command

**Action:** Start multi-agent orchestration

---

## ORCHESTRATOR MODE - PROTOCOL

### Chain of Thought (MANDATORY)

**Analyze the mission:** $ARGUMENTS

1. **What did the user ask for?**
   - Main task
   - Explicit requirements
   - Implicit requirements

2. **What is the implicit objective?**
   - Problem to solve
   - Value delivered
   - Define success criteria.

3. **Which specialists are required?**
   - Database? → `database-architect`
   - Backend? → `backend-specialist`
   - Frontend? → `frontend-specialist`
   - Mobile? → `mobile-developer`
   - Security? → `security-auditor`
   - Testing? → `test-engineer`
   - DevOps? → `devops-engineer`

4. **Why is this complex enough for orchestration?**
   - Multiple domains
   - Task dependencies
   - Coordination needed

---

## PHASE 1: PLANNING

**Agent:** `project-planner`
**Model:** `openai/codex` (for high complexity)

**Before starting Phase 1:**
1. **Read existing todolist:**
   ```javascript
   todoread()
   ```
2. **Create orchestration todolist:**
   ```javascript
   todowrite({
     todos: [
       {
         id: "phase1-analysis",
         content: "Phase 1: Analyze mission complexity",
         status: "in_progress",
         priority: "high"
       },
       {
         id: "phase1-planning",
         content: "Phase 1: Create planning artifacts",
         status: "pending",
         priority: "high"
       },
       {
         id: "phase2-foundation",
         content: "Phase 2: Foundation (DB + Security)",
         status: "pending",
         priority: "high"
       },
       {
         id: "phase2-core",
         content: "Phase 2: Core (Backend + Frontend)",
         status: "pending",
         priority: "high"
       },
       {
         id: "phase2-polish",
         content: "Phase 2: Polish (Tests + Performance)",
         status: "pending",
         priority: "medium"
       },
        {
          id: "phase3-verification",
          content: "Phase 3: Final verification",
          status: "pending",
          priority: "high"
        }
     ]
   })
   ```

**Tasks:**
1. Deep requirement analysis
2. Identify all components
3. Define tech stack
4. Define file structure
5. **Task breakdown** with:
   - IDs
   - Clear names
   - Assigned agents
   - Required skills
   - Priorities (P0-P3)
   - Dependencies
   - INPUT→OUTPUT→VERIFY criteria

**Output:** Artifacts in `docs/requirements/` and `docs/sprint/Sprint-XX/`

**Required:**
1. Ensure `docs/` exists at the root (create if missing).
2. Use the question tool to ask whether to use the latest sprint or create a new one.
3. If no sprint exists, create `docs/sprint/Sprint-01/`.
4. Update `SPRINT_GOAL.md`, `BACKLOG.md`, `TASKS.md`, `RISK_REGISTER.md`.

**After planning artifacts:**
- Update todolist: Mark "phase1-analysis" and "phase1-planning" as `completed`
- Mark "phase2-foundation" as `in_progress`

### MANDATORY STOP POINT

> **STOP:** Use the question tool to ask " PLAN RECORDED in `docs/`
>
> Plan summary:
> - Total tasks: [X]
> - Required agents: [list]
> - Estimate: [time]
>
> **Review** plan. Type 'yes' to proceed to implementation or request changes."

**DO NOT CONTINUE WITHOUT EXPLICIT APPROVAL.**

---

## PHASE 2: IMPLEMENTATION

**After plan approval:**

### Execution Order (P0 → P1 → P2 → P3)

**P0 - Foundation (Critical)**
Run in parallel:
-  `database-architect` (if DB is needed)
-  `security-auditor` (always for auth/security)

** STOP:**
- Update todolist: Mark "phase2-foundation" as `completed`
- Mark "phase2-core" as `in_progress`
- " P0 (Foundation) completed. Proceed to P1 (Core Backend)?"

**P1 - Core Backend**
Run:
-  `backend-specialist`

** STOP:**
- Update todolist: Mark "phase2-core" as `completed`
- Mark "phase2-polish" as `in_progress`
- " P1 (Core Backend) completed. Proceed to P2 (UI/UX)?"

**P2 - UI/UX**
Choose ONE (never both):
-  `frontend-specialist` (WEB projects)
-  `mobile-developer` (MOBILE projects)

** STOP:**
- Update todolist: Mark "phase2-polish" as `completed`
- Mark "phase3-verification" as `in_progress`
- " P2 (UI/UX) completed. Proceed to P3 (Polish)?"

**P3 - Polish**
Run in parallel:
-  `test-engineer`
-  `performance-optimizer` (if needed)
-  `seo-specialist` (if needed)
-  `devops-engineer` (for deploy/setup)

### FINAL STOP POINT

> **STOP:** Use the question tool to ask " IMPLEMENTATION COMPLETE!
>
> All phases executed:
> - P0: Foundation 
> - P1: Core Backend 
> - P2: UI/UX 
> - P3: Polish 
>
 > **Run final verification (Phase 3) with all scripts?**
 >
 > Reply:
 > - 'yes' → Run full Phase 3
 > - 'later' → Wait for manual `/test`"

---

## PHASE 3: VERIFICATION

**Before running verification:**
- Update todolist: Ensure "phase3-verification" is `in_progress`

**MANDATORY SCRIPTS (run in order):**

### P0 - Critical
```bash
npm run lint
npx tsc --noEmit
```
!`python .opencode/skills/vulnerability-scanner/scripts/security_scan.py .`

### P1 - Quality
!`python .opencode/skills/frontend-design/scripts/ux_audit.py .`

### P2 - Performance
!`python .opencode/skills/performance-profiling/scripts/lighthouse_audit.py http://localhost:3000`

### P3 - E2E Tests
!`python .opencode/skills/webapp-testing/scripts/playwright_runner.py http://localhost:3000 --screenshot`

### All-in-One
!`python .opencode/scripts/verify_all.py . --url http://localhost:3000`

**After all scripts pass:**
- Update todolist: Mark "phase3-verification" as `completed`

### FINAL STOP POINT

 > **STOP:** Use the question tool to ask "PHASE 3 COMPLETE!
>
> Results:
> - Lint: [PASS/FAIL]
> - Security: [PASS/WARN/FAIL]
> - Tests: [PASS/FAIL]
> - Performance: [PASS/WARN/FAIL]
>
> All tasks in todolist marked as completed.
>
> **Mark the project complete and update the sprint Tasks?**
>
> Reply 'yes' to finish or describe the adjustments needed."

---

## AVAILABLE SPECIALIST AGENTS

| Agent | Domain | File |
|--------|---------|---------|
| `project-planner` | Planning | `.opencode/prompts/project-planner.md` |
| `database-architect` | Schema/Migrations | `.opencode/prompts/database-architect.md` |
| `security-auditor` | Vulnerabilities | `.opencode/prompts/security-auditor.md` |
| `backend-specialist` | Python/FastAPI | `.opencode/prompts/backend-specialist.md` |
| `frontend-specialist` | React/TanStack | `.opencode/prompts/frontend-specialist.md` |
| `mobile-developer` | iOS/Android/RN | `.opencode/prompts/mobile-developer.md` |
| `test-engineer` | QA/E2E | `.opencode/prompts/test-engineer.md` |
| `performance-optimizer` | Web Vitals | `.opencode/prompts/performance-optimizer.md` |
| `seo-specialist` | Ranking/Geo | `.opencode/prompts/seo-specialist.md` |
| `devops-engineer` | Docker/CI/CD | `.opencode/prompts/devops-engineer.md` |
| `debugger` | Root Cause | `.opencode/prompts/debugger.md` |
| `explorer-agent` | Code Analysis | `.opencode/prompts/explorer-agent.md` |
| `penetration-tester` | Offensive Sec | `.opencode/prompts/penetration-tester.md` |
| `product-owner` | Requirements | `.opencode/prompts/product-owner.md` |
| `documentation-writer` | Docs | `.opencode/prompts/documentation-writer.md` |

---

## FINAL DOCUMENTATION

After successful completion:

 1. **Update** `docs/sprint/Sprint-XX/TASKS.md` with:
    -  All tasks marked complete
    -  Phase 3 completion date
    -  Verification results

2. **Create** a user summary:
   ```markdown
   ##  Mission Complete!

   ### Summary
   - **Task:** [Description]
   - **Total time:** [Duration]
   - **Agents used:** [List]

   ### Deliverables
   - [List of files created/modified]

   ### Quality
   -  All checks passed
   -  Code ready for production

   ### Next Steps (optional)
   - [Suggested follow-ups]
   ```

---

## CRITICAL RULES

1.  **Minimum 3 agents:** fewer than that is not orchestration
2.  **Strict phases:** Planning → STOP → Implementation → STOP → Verification
3.  **Documentation:** always update the docs plan
 4.  **Quality:** Phase 3 is mandatory
5.  **Communication:** keep the user informed at each phase

---

**Welcome to the OpenKit-OpenCode Hybrid System!**

---

## Workflow Reference (Migrated)

# /engineer - Engineering Hub

$ARGUMENTS

---

## Overview

The `/engineer` command is the central hub for development operations. It operates in two modes:

1.  **Router Mode**: For specific phases (Planning, Implementation, Testing), it routes to specialized sub-commands.
2.  **Orchestrator Mode**: For complex, multi-domain tasks, it coordinates multiple specialized agents.

---

## 1. Router Mode (Specialized Workflows)

Use these sub-commands for focused tasks:

| Command | Purpose | Use When |
| :--- | :--- | :--- |
| `/specify` | **Specification** | Create feature specification. |
| `/clarify` | **Clarification** | Resolve spec ambiguities. |
| `/plan` | **Planning** | Creating Sprint plans or Task plans. |
| `/tasks` | **Tasking** | Generate executable tasks. |
| `/analyze` | **Analysis** | Validate spec/plan/tasks. |
| `/checklist` | **Checklist** | Spec/plan readiness checks. |
| `/impl` | **Implementation** | Writing code, fixing bugs, adding features. |
| `/test` | **Testing** | Generating or running tests. |
| `/deploy` | **Deployment** | Deploying to production or staging. |
| `/doc` | **Documentation** | Writing/Updating docs only. |
| `/debug` | **Debugging** | Investigating complex errors. |
| `/ui-ux` | **Design** | Creating design systems or UI components. |
| `/context`| **Analysis** | Generating context packs. |

**Usage:**
```bash
/engineer plan feature login
/engineer impl fix auth bug
/engineer test user service
```

---

## 2. Orchestrator Mode (Complex Missions)

**Trigger**: When the task is complex, requires multiple domains (e.g., "Build a full SaaS dashboard"), or doesn't fit a single sub-command.

**Goal**: Coordinate specialized agents (Frontend, Backend, Database, Security) to solve the problem.

### Critical Rules
- **Documentation**: All plans and artifacts MUST follow `rules/MASTER.md` Documentation Integrity Protocol.
- **Discovery Gate:** Always execute `/context` (or `/brainstorm`) before authoring a new plan or touching code.
- **Planning Gate:** `/impl` or specialist agents cannot run until `/specify`, `/clarify`, `/plan`, and `/tasks` have produced the required docs.
- **Minimum 3 Agents**: If you use fewer than 3, you are not orchestrating.
- **2-Phase Strict Execution**:
    - **Phase 1: Planning** (`project-planner` only). STOP for approval.
    - **Phase 2: Implementation** (Parallel agents based on approval).

### Orchestration Protocol

#### Step 1: Analyze & Plan (Phase 1)
1.  ** Chain of Thought (MANDATORY)**:
    - *Model*: Select `GLM 4.7 Thinking` for orchestration.
    - *Input*: Determine what the user strictly asked.
    - *Intent*: Determine the implicit goal.
    - *Domains*: Identify required specialists (`frontend-specialist`, `backend-specialist`, etc.).
    - *Reasoning*: Explain why this is complex enough for orchestration.

2.  **Identify Domains**: Security, Backend, Frontend, Database, etc.
3.  **Agent**: Use `project-planner` to produce SDD artifacts:
    - `docs/requirements/<feature>/PROBLEM_STATEMENT.md`
    - `docs/requirements/<feature>/USER_STORIES.md`
    - `docs/requirements/<feature>/ACCEPTANCE_CRITERIA.md`
    - `docs/requirements/<feature>/RISKS.md`
    - `docs/requirements/<feature>/PLAN.md`
    - `docs/sprint/Sprint-XX/TASKS.md`
4.  **STOP**: Use the question tool to ask "Plan recorded in docs. Proceed to implementation?".

#### Step 2: Execute (Phase 2)
After approval, invoke agents in **PARALLEL** groups:
1.  **Foundation**: `database-architect`, `security-auditor`.
2.  **Core**: `backend-specialist`, `frontend-specialist`.
3.  **Polish**: `test-engineer`, `devops-engineer`.

#### Step 3: Verify & Report
1.  **Execute Scripts**:
    - `security_scan.py`
    - `lint_runner.py`
2.  **Synthesize**: Create a final report summarizing all agent contributions.

---

## Available Specialist Agents

| Agent | Domain |
| :--- | :--- |
| `project-planner` | Planning & Task Breakdown |
| `frontend-specialist` | UI/UX, React, CSS |
| `backend-specialist` | API, DB, Logic |
| `database-architect` | Schema, Migrations |
| `security-auditor` | Vulnerabilities, Auth |
| `test-engineer` | Unit/E2E Testing |
| `devops-engineer` | Kubernetes, Docker, Deploy |

---

## Example Usage

**Router:**
> "/engineer test coverage" -> Routes to `/test coverage`

**Orchestrator:**
> "/engineer build a secure e-commerce checkout with stripe"
> 1. Detects complexity.
> 2. Starts **Orchestration Mode**.
> 3. Plans with `project-planner`.
> 4. Executes with `backend-specialist` (API), `frontend-specialist` (UI), `security-auditor` (PCI compliance).
