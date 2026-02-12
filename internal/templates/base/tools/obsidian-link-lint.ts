import { tool } from "@opencode-ai/plugin"
import fs from "node:fs/promises"
import path from "node:path"

type LinkFinding = {
  source: string
  target: string
  exists: boolean
}

const WIKILINK_REGEX = /\[\[([^\]|#]+)(?:#[^\]|]+)?(?:\|[^\]]+)?\]\]/g

async function listMarkdownFiles(dir: string): Promise<string[]> {
  const entries = await fs.readdir(dir, { withFileTypes: true })
  const files: string[] = []
  for (const entry of entries) {
    const full = path.join(dir, entry.name)
    if (entry.isDirectory()) {
      files.push(...(await listMarkdownFiles(full)))
      continue
    }
    if (entry.isFile() && full.endsWith(".md")) files.push(full)
  }
  return files
}

function extractTargets(content: string): string[] {
  const targets: string[] = []
  for (const match of content.matchAll(WIKILINK_REGEX)) {
    const raw = match[1]?.trim()
    if (!raw) continue
    targets.push(raw)
  }
  return targets
}

export default tool({
  description: "Validate Obsidian wikilinks in docs",
  args: {
    root: tool.schema
      .string()
      .default("docs")
      .describe("Relative docs root to scan"),
  },
  async execute(args, context) {
    const root = path.join(context.worktree, args.root)
    const files = await listMarkdownFiles(root)
    const findings: LinkFinding[] = []

    for (const file of files) {
      const content = await fs.readFile(file, "utf8")
      const source = path.relative(context.worktree, file)
      const targets = extractTargets(content)
      for (const target of targets) {
        const abs = path.join(context.worktree, target)
        let exists = true
        try {
          await fs.access(abs)
        } catch {
          exists = false
        }
        findings.push({ source, target, exists })
      }
    }

    const broken = findings.filter((f) => !f.exists)
    const total = findings.length

    const lines: string[] = []
    lines.push(`# Obsidian Link Lint`)
    lines.push(``) 
    lines.push(`- Root: ${args.root}`)
    lines.push(`- Files scanned: ${files.length}`)
    lines.push(`- Internal links found: ${total}`)
    lines.push(`- Broken links: ${broken.length}`)

    if (broken.length) {
      lines.push(``)
      lines.push(`## Broken Links`)
      for (const item of broken) {
        lines.push(`- ${item.source} -> [[${item.target}]]`) 
      }
    }

    return lines.join("\n")
  },
})
