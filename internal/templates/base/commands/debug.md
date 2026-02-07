---
description: Systematic 4-phase root cause analysis for bugs and issues
subtask: true
---

**DEBUG MODE ACTIVATED**

**Reported issue:** $ARGUMENTS

**If $ARGUMENTS is empty:**
Use the question tool to ask: "Describe the problem you're facing. What is the error? When did it start? What did you expect to happen?"

**BMAD Protocol - 4 Phases:**

---

## PHASE 1: Symptom Analysis

**Goal:** Understand what is happening vs. what should happen

**Tasks:**
1. **Observed symptoms:**
   - Exact error messages
   - Full stack traces
   - Anomalous behavior

2. **Expected behavior:**
   - Use the question tool to ask what should happen.
   - Use the question tool to ask when it last worked correctly.

3. **Gap identified:**
   - Difference between observed and expected

4. **Quantify:**
   - Frequency: [always / sometimes / rare]
   - Scope: [all users / specific users / single user]
   - Environment: [production / staging / local]

**Output:** Documented Symptom Report

---

## PHASE 2: Information Gathering

**Goal:** Collect relevant data and context

**Tasks:**
1. **Logs and Errors:**
   !`tail -n 50 logs/app.log 2>/dev/null || echo "No logs found"`
   - App logs, server logs, browser console

2. **Recent Changes:**
   !`git log --oneline -10`
   !`git diff HEAD~5 --name-only`
   - Recent commits
   - Deployments
   - Config changes
   - Dependency updates

3. **System State:**
   - Database state
   - Relevant environment variables
   - Resource usage (CPU, memory)
   - Network connectivity

**Output:** Data Collection Report

---

## PHASE 3: Hypothesis Formation

**Goal:** Generate and test root-cause hypotheses

**Tasks:**
1. **Hypothesis brainstorm:**
   List all possible causes (do not filter yet):
   - Code bug (logic error, typo)
   - Configuration issue
   - Data corruption
   - External dependency failure
   - Resource exhaustion
   - Race condition
   - Environment mismatch

2. **Prioritization:**
   Order by likelihood (High → Medium → Low)

3. **Hypothesis testing:**
   - Test one at a time
   - Document results
   - Eliminate or confirm

**Output:** Hypothesis Testing Report

---

## PHASE 4: Resolution & Verification

**Goal:** Fix and validate the solution

**Tasks:**
1. **Implement the fix:**
   - Minimal change required
   - Focus on root cause
   - Do not over-engineer

2. **Test the fix:**
   - Verify the issue is resolved
   - Check for side effects
   - Run regression tests

3. **Prevent recurrence:**
   - Add tests
   - Update documentation
   - Consider monitoring

**Output:** Resolution Report

---

## STOP POINTS

**Between each phase:**
> Use the question tool to ask "Phase [X] complete. Document generated. Proceed to the next phase?"

**Before implementing a fix:**
> Use the question tool to ask "Root cause identified: [description]. Implement the fix?"

**After the fix:**
> Use the question tool to ask "Fix applied. Validate that the issue is resolved?"

---

## Debug Session Report

Document the full session:

```markdown
# Debug Session: [Issue Title]

## Phase 1: Symptom Analysis
[Summary]

## Phase 2: Information Gathering
[Collected data]

## Phase 3: Hypothesis Formation
[Tested hypotheses and results]

## Phase 4: Resolution & Verification
[Fix applied and validation]

## Summary
- **Root Cause:** [One sentence]
- **Fix:** [One sentence]
- **Time to Resolve:** [Duration]
- **Lessons Learned:** [Learnings]
```

**Use agent `debugger` and skill `systematic-debugging`.**
