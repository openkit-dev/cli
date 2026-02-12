---
trigger: always_on
priority: P0
applies_to: [orchestrator, all-agents, all-commands, all-skills]
---

# TOOL USAGE PROTOCOL — MANDATORY FOR ALL AGENTS

> Definitive guide for correct tool selection and usage patterns across the entire system.
> Based on official OpenCode documentation: https://opencode.ai/docs/tools/

---

## Purpose

This rule enforces **consistent tool usage** across:
- **Orchestrator** (`.opencode/prompts/orchestrator.md`)
- **All specialist prompts** (`.opencode/prompts/*.md`)
- **All commands** (`.opencode/commands/*.md`)
- **All skills** (`.opencode/skills/*/SKILL.md`)

**Philosophy:** Use the right tool for the right job. Never improvise.

---

## Tool Configuration

**Tools are configured in `opencode.json`**, NOT in prompt markdown files.

**Example:**
```json
{
  "agent": {
    "orchestrator": {
      "tools": {
        "read": true,
        "grep": true,
        "glob": true,
        "list": true,
        "bash": true,
        "edit": true,
        "write": true,
        "patch": true,
        "webfetch": true,
        "skill": true,
        "todowrite": true,
        "todoread": true,
        "question": true
      },
      "permission": {
        "edit": "ask",
        "bash": "ask",
        "webfetch": "ask"
      }
    }
  }
}
```

**Permission values:**
- `"allow"` - Tool can be used without asking
- `"ask"` - User must approve each use (default for most tools)
- `"deny"` - Tool cannot be used

**Orchestrator has ALL tools enabled** to coordinate complex multi-agent missions.

---

## Complete Tool Reference

### Built-in Tools Available

| Tool | Purpose | Permission | Use When |
|------|---------|-----------|----------|
| `bash` | Execute shell commands | `bash` | Git, npm, build, test scripts |
| `read` | Read file contents | `read` | View file contents |
| `edit` | Modify existing files | `edit` | Surgical file changes |
| `write` | Create/overwrite files | `edit` | New files or full rewrites |
| `patch` | Apply patch files | `edit` | Apply diffs |
| `glob` | Find files by pattern | `glob` | Search files by name |
| `grep` | Search file contents | `grep` | Find code patterns |
| `list` | List directory contents | `list` | Browse directories |
| `skill` | Load skill instructions | `skill` | Apply domain knowledge |
| `todowrite` | Create/update todo lists | `todowrite` | Track complex tasks |
| `todoread` | Read todo lists | `todoread` | Check task status |
| `question` | Ask user questions | `question` | Get user input |
| `webfetch` | Fetch web content | `webfetch` | Read documentation online |
| `lsp` (experimental) | Code intelligence | `lsp` | Definitions, references, hover |

**Note:** All tools controlled by `edit` permission: `edit`, `write`, `patch`, `multiedit`

---

## Tool Selection Decision Tree

### 1. **File Discovery & Search**

| Goal | Tool | Example | Why |
|------|------|---------|-----|
| Find files by **name pattern** | `glob` | `glob("**/*.tsx")` | Fast, works with any codebase size |
| Find files by **content pattern** | `grep` | `grep("import.*useState")` | Searches file contents with regex |
| List **directory contents** | `list` | `list("src/")` | Browse specific directory |
| Find specific **known file** | `read` | `read("src/App.tsx")` | Direct access when path is known |
| Search **specific file content** | `read` + analyze | `read("src/App.tsx")` | When scope is 1-3 files |
| Multi-round **exploration** | `task(explorer-agent)` | Complex discovery | Needs multiple iterations |

**NEVER:**
- Use `bash find` instead of `glob`
- Use `bash grep` instead of `grep` tool
- Use `bash cat` instead of `read` tool
- Use `bash ls` instead of `list` tool

**ALWAYS:**
- Use `glob` for file patterns
- Use `grep` for content search across many files
- Use `list` for directory browsing
- Use `read` for known file paths

---

### 2. **File Operations**

| Goal | Tool | Example | Why |
|------|------|---------|-----|
| **Read** existing file | `read` | `read("src/config.ts")` | Mandatory before any edit |
| **Edit** existing file | `edit` | `edit("file.ts", "old", "new")` | Surgical changes, preserves formatting |
| **Create** new file | `write` | `write("new.ts", "content")` | For new files only |
| **Overwrite** existing file | `write` (after `read`) | Must read first | Requires prior read |
| **Apply patch** | `patch` | `patch("file.ts", patchContent)` | Apply diffs from external sources |
| **Multiple edits** same file | Multiple `edit` calls | Sequential edits | Safer than full rewrite |
| **Batch edits** across files | Parallel `edit` calls | Independent changes | Maximize performance |

** CRITICAL RULES:**

1. **Always `read` before `edit`/`write`:**
   ```javascript
   //  CORRECT
   read("file.ts")
   edit("file.ts", oldString, newString)
   
   //  WRONG - Will fail
   edit("file.ts", oldString, newString)
   ```

2. **Preserve line numbers context:**
   - Read output shows: `00042|   const value = "old"`
   - The prefix `00042|` is for reference ONLY
   - Your `oldString` must ONLY contain: `   const value = "old"` (exact indentation, no prefix)

3. **Use `edit` for existing files, `write` for new files:**
   ```javascript
   //  For existing files
   read("existing.ts")
   edit("existing.ts", "old", "new")
   
   //  For new files
   write("new-file.ts", "content")
   
   //  WRONG - Overwrites unnecessarily
   read("existing.ts")
   write("existing.ts", modifiedContent)
   ```

4. **Permissions:**
   - All file modification tools (`edit`, `write`, `patch`) are controlled by `edit` permission
   - Check agent frontmatter for permission settings

---

### 3. **Code Execution & Shell**

| Goal | Tool | Example | Why |
|------|------|---------|-----|
| Run **git** commands | `bash` | `bash("git status")` | Git operations |
| Run **npm/yarn** commands | `bash` | `bash("npm install")` | Package management |
| Run **build/test** scripts | `bash` | `bash("npm run build")` | CI/CD operations |
| Run **Python** scripts | `bash` | `bash("python script.py .")` | Skill scripts, verification |
| Execute **Docker** commands | `bash` | `bash("docker compose up")` | Container management |
| **Change directory** for command | `workdir` param | `bash("pytest", workdir="/backend")` | Never use `cd &&` |

** NEVER:**
- Use `cd <dir> && <command>` → Use `workdir` parameter instead
- Use `bash` for file reading → Use `read` tool
- Use `bash` for file editing → Use `edit` tool
- Use `bash` for file search → Use `glob`/`grep` tools
- Use `bash echo` to output text → Just output text directly
- Use `bash find` → Use `glob` tool
- Use `bash grep` → Use `grep` tool
- Use `bash cat/head/tail` → Use `read` tool
- Use `bash ls` → Use `list` tool

** ALWAYS:**
- Quote paths with spaces: `bash('rm "path with spaces/file.txt"')`
- Use `workdir` instead of `cd`:
  ```javascript
  //  CORRECT
  bash("pytest tests", workdir="/backend")
  
  //  WRONG
  bash("cd /backend && pytest tests")
  ```

---

### 4. **User Interaction**

#### 4.1 **question Tool (MANDATORY)**

Use `question` tool when user must choose from **multiple options**:

| Situation | Use `question` | Example |
|-----------|----------------|---------|
| Multiple framework options |  YES | "Which framework: FastAPI or Django?" |
| Deployment target selection |  YES | "Deploy to staging or production?" |
| Feature preference |  YES | "Use REST or GraphQL?" |
| STOP points requiring approval |  YES | "Proceed with implementation?" |
| Trade-off decisions |  YES | "Option A (fast) or Option B (secure)?" |
| Any multi-option choice |  YES | Present options for user selection |

**Format:**
```javascript
question({
  questions: [{
    question: "Clear, concise question text",
    header: "Short title (max 30 chars)",
    options: [
      { label: "Option 1", description: "Why choose this" },
      { label: "Option 2", description: "Alternative reason" }
    ],
    multiple: false // or true for multi-select
  }]
})
```

**Features:**
- Automatic "Type your own answer" option (when `custom: true`, which is default)
- Multiple questions in one call
- Multi-select support (`multiple: true`)
- Navigation between questions

** NEVER use inline questions for multiple options:**
```markdown
 WRONG:
What would you like to do?
1. Option A
2. Option B
3. Option C

Please reply with your choice.
```

** ALWAYS use `question` tool:**
```javascript
 CORRECT:
question({
  questions: [{
    question: "What would you like to do?",
    header: "Action Selection",
    options: [
      { label: "Option A", description: "Does X" },
      { label: "Option B", description: "Does Y" },
      { label: "Option C", description: "Does Z" }
    ]
  }]
})
```

#### 4.2 **Inline Questions (Allowed)**

Simple yes/no confirmations where context is clear:
-  "Delete this file?" (after showing what will be deleted)
-  "Continue with the plan?" (after presenting it)

---

### 5. **Task Tracking**

#### 5.1 **todowrite/todoread Tools (MANDATORY)**

**See complete protocol:** `.opencode/rules/TODOLIST_PROTOCOL.md`

Use todolist tools for **complex multi-step tasks**:

| Situation | Use `todowrite` | Example |
|-----------|----------------|---------|
| Multi-step tasks (3+ steps) |  YES | "Build feature with 5 components" |
| Complex planning phases |  YES | Sprint planning with 5+ stories |
| Orchestration mode |  YES | Coordinating 3+ agents |
| Tracking execution |  YES | Marking progress across phases |
| Parallel workflows |  YES | Managing dependencies |
| STOP points with outcomes |  YES | "P0 complete. Proceed to P1?" |

**MANDATORY:** After creating todolist, use `question` tool to request user approval before execution.

**Lifecycle:**

1. **Read existing tasks FIRST (if available):**
   ```javascript
   //  CORRECT - Check existing state
   todoread()
   // Then create or update based on what exists
   
   //  WRONG - Creates duplicates
   todowrite({ todos: [...] })
   ```

2. **Create todolist at START of complex tasks:**
   ```javascript
   todowrite({
     todos: [
       {
         id: "task-1",
         content: "Brief description (max 50 chars)",
         status: "pending",
         priority: "high"
       },
       {
         id: "task-2",
         content: "Another task",
         status: "pending",
         priority: "medium"
       }
     ]
   })
   ```

3. **Request approval (MANDATORY):**
   ```javascript
   // MUST use question tool for approval
   question({
     questions: [{
       question: "Task list created with 2 tasks. Proceed with execution?",
       header: "Approve Execution",
       options: [
         { label: "Yes, proceed", description: "Start executing tasks" },
         { label: "Cancel", description: "Don't execute" }
       ]
     }]
   })
   ```

4. **Update as you progress (NO emojis):**
   ```javascript
   // Mark as in_progress when starting
   todowrite({
     todos: [
       {
         id: "task-1",
         content: "Task description",
         status: "in_progress", // ← Changed
         priority: "high"
       },
       {
         id: "task-2",
         content: "Another task",
         status: "pending",
         priority: "medium"
       }
     ]
   })
   
   // Mark as completed when done (NO emojis!)
   todowrite({
     todos: [
       {
         id: "task-1",
         content: "Task description", // ← NO emoji
         status: "completed", // ← Changed
         priority: "high"
       },
       {
         id: "task-2",
         content: "Another task",
         status: "in_progress", // ← Start next
         priority: "medium"
       }
     ]
   })
   ```

**Status values:** `pending`, `in_progress`, `completed`, `cancelled`  
**Priority values:** `high`, `medium`, `low`

** CRITICAL RULES:**

1. **NO emojis in task content** (keep content clean)
2. **MANDATORY approval** with `question` tool after creating todolist
3. **Only ONE task** can be `in_progress` at a time
4. **Complete current tasks** before starting new ones
5. **Mark tasks complete IMMEDIATELY** after finishing
6. **Always `todoread`** before `todowrite` (when available)

** NEVER:**
- Use emojis in task content (e.g., "Task ")
- Skip approval with `question` tool
- Use inline tracking for complex tasks (3+ steps)
- Forget to `todoread` before `todowrite` (when tool is available)
- Have multiple tasks with `in_progress` status
- Skip updating todolist during execution

** ALWAYS:**
- Keep task content emoji-free
- Use `question` tool for approval after creating todolist
- Use `todoread` first (if available)
- Create todolist at START of complex tasks
- Mark tasks complete IMMEDIATELY
- Only ONE task `in_progress` at a time

**See complete protocol:** `.opencode/rules/TODOLIST_PROTOCOL.md`

#### 5.2 **Inline Tracking (Allowed)**

Simple, immediate operations:
-  "Installing dependencies..." (single npm install)
-  "Running tests..." (single test command)

**Note:** `todoread` and `todowrite` are **disabled for subagents by default** but can be enabled in agent config.

---

### 6. **Knowledge Loading**

| Goal | Tool | Example | Why |
|------|------|---------|-----|
| Load **skill instructions** | `skill` | `skill("clean-code")` | Domain-specific workflows |
| Load **agent capabilities** | `read` + prompt file | `read(".opencode/prompts/backend-specialist.md")` | Agent context |
| Invoke **specialist agent** | `task` | `task(subagent_type="backend-specialist", ...)` | Delegate complex tasks |

**Skill Tool:**
- Loads a `SKILL.md` file and returns content
- Skills provide domain-specific instructions
- Located in `.opencode/skills/*/SKILL.md`

```javascript
// Load a skill
skill("python-patterns")
// Returns full skill content for context
```

** NEVER:**
- Invoke agent without reading agent file first
- Load skill without announcing: `Applying knowledge of @[skill-name]...`

** ALWAYS:**
- Read agent file before `task` invocation
- Load skills listed in agent frontmatter
- Announce skill/agent usage to user

---

### 7. **Web Operations**

| Goal | Tool | Example | Why |
|------|------|---------|-----|
| Fetch **web content** | `webfetch` | `webfetch("https://docs.com", "markdown")` | API docs, external resources |
| Fetch **GitHub PR/Issues** | `bash` with `gh` CLI | `bash("gh pr view 123")` | GitHub operations |

**Webfetch Tool:**
- Fetches URL content and converts to format
- Formats: `markdown` (default), `text`, `html`
- HTTP URLs auto-upgrade to HTTPS
- Results may be summarized if very large

```javascript
webfetch("https://example.com/docs", "markdown")
```

** NEVER:**
- Use `webfetch` for local file operations
- Use `webfetch` when specialized tool exists (like `gh` for GitHub)

---

### 8. **Code Intelligence (Experimental)**

| Goal | Tool | Example | Why |
|------|------|---------|-----|
| Go to **definition** | `lsp` | `lsp("goToDefinition", ...)` | Find symbol definition |
| Find **references** | `lsp` | `lsp("findReferences", ...)` | Find all usages |
| Get **hover info** | `lsp` | `lsp("hover", ...)` | Type/doc information |
| Find **symbols** | `lsp` | `lsp("documentSymbol", ...)` | List file symbols |
| Call **hierarchy** | `lsp` | `lsp("prepareCallHierarchy", ...)` | Function call graph |

**LSP Tool (Experimental):**
- Only available with `OPENCODE_EXPERIMENTAL_LSP_TOOL=true`
- Requires LSP servers configured in project
- Supported operations:
  - `goToDefinition`
  - `findReferences`
  - `hover`
  - `documentSymbol`
  - `workspaceSymbol`
  - `goToImplementation`
  - `prepareCallHierarchy`
  - `incomingCalls`
  - `outgoingCalls`

**Note:** Configure LSP servers in your OpenCode runtime configuration. See [LSP Servers docs](https://opencode.ai/docs/lsp).

---

## Tool Usage Patterns by Agent Type

### Orchestrator Pattern
```javascript
// 1. Discovery
todoread() // Check existing tasks (if available)
glob("**/*.md") // Find project files
grep("import.*React", "**/*.tsx") // Find patterns
list("src/") // Browse directories

// 2. Planning
todowrite({ todos: [...] }) // Create todolist
question({ questions: [...] }) // Get user approval

// 3. Execution
task(subagent_type="backend-specialist", ...) // Invoke specialist
task(subagent_type="frontend-specialist", ...) // Parallel invocation

// 4. Verification
bash("python .opencode/scripts/verify_all.py .") // Run checks
todowrite({ todos: [{ status: "completed" }] }) // Update todolist
```

### Specialist Agent Pattern
```javascript
// 1. Context
read(".opencode/prompts/backend-specialist.md") // Read prompt file
skill("python-patterns") // Load skills

// 2. Discovery
glob("backend/**/*.py") // Find relevant files
grep("def.*login", "backend/**/*.py") // Find patterns
read("backend/main.py") // Read specific files

// 3. Implementation
edit("backend/main.py", oldString, newString) // Edit existing
write("backend/new-feature.py", content) // Create new

// 4. Validation
bash("python -m pytest tests/") // Run tests
```

### Command Pattern
```javascript
// 1. Input
question({ questions: [...] }) // Get user input if needed

// 2. Execution
bash("npm run lint") // Execute command
read("output.log") // Read results

// 3. Reporting
// Output results directly to user
```

---

## Tool Priority Matrix

When multiple tools could work, use this priority:

| Scenario | Priority 1 | Priority 2 | Priority 3 |
|----------|-----------|-----------|-----------|
| **File search** | `glob` | `grep` | `bash find` (never) |
| **Content search** | `grep` | `read` (1-3 files) | `bash grep` (never) |
| **File read** | `read` | `bash cat` (never) | - |
| **File edit** | `edit` | `write` (after read) | `bash sed` (never) |
| **Directory list** | `list` | `glob` | `bash ls` (never) |
| **Command exec** | `bash` | - | - |
| **User questions** | `question` tool | Inline (yes/no only) | - |
| **Task tracking** | `todowrite`/`todoread` | Inline (simple tasks) | - |
| **Apply patches** | `patch` | `edit` (manual) | - |
| **Load knowledge** | `skill` | `read` (manual) | - |
| **Web content** | `webfetch` | `bash curl` | - |

---

## Common Mistakes & Fixes

### Mistake 1: Using bash for file operations
```javascript
//  WRONG
bash("cat src/App.tsx")
bash("grep 'useState' src/**/*.tsx")
bash("find . -name '*.tsx'")
bash("ls src/")

//  CORRECT
read("src/App.tsx")
grep("useState", "src/**/*.tsx")
glob("**/*.tsx")
list("src/")
```

### Mistake 2: Editing without reading
```javascript
//  WRONG
edit("file.ts", "old", "new")

//  CORRECT
read("file.ts")
edit("file.ts", "old", "new")
```

### Mistake 3: Inline questions for multiple options
```javascript
//  WRONG
"Which framework do you prefer? Reply with: 1. FastAPI, 2. Django, 3. Flask"

//  CORRECT
question({
  questions: [{
    question: "Which framework do you prefer?",
    header: "Framework Selection",
    options: [
      { label: "FastAPI", description: "Modern, fast, async" },
      { label: "Django", description: "Batteries included" },
      { label: "Flask", description: "Lightweight, flexible" }
    ]
  }]
})
```

### Mistake 4: Not using todolist for complex tasks
```javascript
//  WRONG (for 5-step feature)
// Just start coding without tracking

//  CORRECT
todoread() // Check existing (if available)
todowrite({
  todos: [
    { id: "step1", content: "Design schema", status: "pending", priority: "high" },
    { id: "step2", content: "Create API", status: "pending", priority: "high" },
    { id: "step3", content: "Build UI", status: "pending", priority: "medium" },
    { id: "step4", content: "Write tests", status: "pending", priority: "medium" },
    { id: "step5", content: "Deploy", status: "pending", priority: "low" }
  ]
})
```

### Mistake 5: Using cd && command
```javascript
//  WRONG
bash("cd backend && pytest tests")

//  CORRECT
bash("pytest tests", workdir="/Users/user/project/backend")
```

### Mistake 6: Not checking todoread before todowrite
```javascript
//  WRONG (creates duplicates)
todowrite({ todos: [...] })

//  CORRECT
todoread() // Check what exists first
// Then create or update based on existing state
todowrite({ todos: [...] })
```

### Mistake 7: Multiple tasks in_progress
```javascript
//  WRONG
todowrite({
  todos: [
    { id: "task1", status: "in_progress", ... },
    { id: "task2", status: "in_progress", ... } // ← BAD!
  ]
})

//  CORRECT - Only ONE in_progress
todowrite({
  todos: [
    { id: "task1", status: "in_progress", ... },
    { id: "task2", status: "pending", ... } // ← Wait
  ]
})
```

---

## Validation Checklist

Before any tool invocation, verify:

- [ ] Is this the **right tool** for the job? (Check decision tree)
- [ ] Have I **read the file** before editing? (For `edit`/`write`)
- [ ] Am I using **`question` tool** for multiple options? (Not inline)
- [ ] Am I using **`todolist` tools** for complex tasks? (3+ steps)
- [ ] Have I **`todoread`** before **`todowrite`**? (When available)
- [ ] Do I have only **ONE task** `in_progress`? (For todolist)
- [ ] Am I using **`workdir`** instead of `cd &&`? (For bash)
- [ ] Have I **loaded required skills** for this agent? (Check frontmatter)
- [ ] Am I using **correct tool priority**? (glob > grep > bash find)

---

## Reference

- **Official Docs**: https://opencode.ai/docs/tools/
- **Master Rules**: `.opencode/rules/MASTER.md`
- **Obsidian Linking**: `.opencode/rules/OBSIDIAN_LINKING.md`
- **Docs File Glossary**: `.opencode/rules/DOCS_FILE_GLOSSARY.md`
- **Agent Prompts**: `.opencode/prompts/*.md`
- **Command Files**: `.opencode/commands/*.md`
- **Skill Files**: `.opencode/skills/*/SKILL.md`
- **Permissions**: https://opencode.ai/docs/permissions/

---

## Ignore Patterns

**Internally**, tools like `grep`, `glob`, and `list` use [ripgrep](https://github.com/BurntSushi/ripgrep), which respects `.gitignore` patterns by default.

**To include ignored files**, create a `.ignore` file:

```
# .ignore file
!node_modules/
!dist/
!build/
```

This allows tools to search within normally-ignored directories.

---

**This rule is MANDATORY and supersedes any conflicting instructions in individual agents, commands, or skills.**

**Last Updated:** 2026-02-04  
**Based on:** OpenCode Official Documentation v1.0
