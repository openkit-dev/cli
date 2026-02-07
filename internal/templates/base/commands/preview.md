---
description: Preview server start, stop, and status check. Local development server management.
---

# /preview - Preview Management

$ARGUMENTS

## Task

Manage preview server: start, stop, status check.

### Commands

```
/preview           - Show current status
/preview start     - Start server
/preview stop      - Stop server
/preview restart   - Restart
/preview check     - Health check
```

---

## Usage Examples

### Start Server
```
/preview start

Response:
 Starting preview...
   Port: 3000
   Type: Next.js

 Preview ready!
   URL: http://localhost:3000
```

### Status Check
```
/preview

Response:
=== Preview Status ===

 URL: http://localhost:3000
 Project: C:/projects/my-app
 Type: nextjs
 Health: OK
```

### Port Conflict
```
/preview start

Response:
 Port 3000 is in use.

Use the question tool to ask:
```javascript
question({
  questions: [{
      question: "What would you like to do?",
      header: "Port Conflict",
      options: [
        { label: "Start on port 3001", description: "Use alternative port" },
        { label: "Close app on 3000", description: "Free up the port" },
        { label: "Specify different port", description: "Custom port number" }
      ],
      default: 0
    }]
})
```
```

---

## Technical

Auto preview uses `auto_preview.py` script which wraps Docker Compose:

```bash
python .opencode/scripts/auto_preview.py start
python .opencode/scripts/auto_preview.py stop
python .opencode/scripts/auto_preview.py status
```
