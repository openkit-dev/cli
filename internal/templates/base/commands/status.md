---
description: Display agent and project status. Progress tracking and status board.
---

# /status - Show Status

$ARGUMENTS

## Task

Show current project and agent status.

## Question Tool Protocol (MANDATORY)

When status requires user input or choices:
- Use `question` tool for all multi-option scenarios

**Example:**
```javascript
question({
  questions: [{
      question: "What level of detail?",
      header: "Detail",
      options: [
        { label: "Summary", description: "Brief overview" },
        { label: "Detailed", description: "Full breakdown" }
      ]
    }]
})
```

### What It Shows

1. **Project Info**
   - Project name and path
   - Tech stack
   - Current features

2. **Agent Status Board**
   - Which agents are running
   - Which tasks are completed
   - Pending work

3. **File Statistics**
   - Files created count
   - Files modified count

4. **Preview Status**
   - Is server running
   - URL
   - Health check

---

## Example Output

```
=== Project Status ===

 Project: my-ecommerce
 Path: C:/projects/my-ecommerce
 Type: nextjs-ecommerce
 Status: active

 Tech Stack:
   Framework: next.js
   Database: postgresql
   Auth: clerk
   Payment: stripe

 Features (5):
   • product-listing
   • cart
   • checkout
   • user-auth
   • order-history

⏳ Pending (2):
   • admin-panel
   • email-notifications

 Files: 73 created, 12 modified

=== Agent Status ===

 database-architect → Completed
 backend-specialist → Completed
 frontend-specialist → Dashboard components (60%)
⏳ test-engineer → Waiting

=== Preview ===

 URL: http://localhost:3000
 Health: OK
```

---

## Technical

Status uses these scripts:
- `python .opencode/scripts/session_manager.py status`
- `python .opencode/scripts/auto_preview.py status`
