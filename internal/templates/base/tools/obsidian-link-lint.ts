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

async function pathExists(filePath: string): Promise<boolean> {
  try {
    await fs.access(filePath)
    return true
  } catch {
    return false
  }
}

async function hasAnyFileWithName(dir: string, name: string): Promise<boolean> {
  const files = await listMarkdownFiles(dir)
  return files.some((file) => path.basename(file) === name)
}

async function resolveTargetExists(
  vaultRoot: string,
  sourceFile: string,
  target: string,
): Promise<boolean> {
  if (!target || target === "x") return true

  const normalized = target.replace(/^\/+/, "")
  const candidates = new Set<string>()

  candidates.add(path.join(vaultRoot, normalized))
  candidates.add(path.join(path.dirname(sourceFile), normalized))

  if (!normalized.endsWith(".md")) {
    candidates.add(path.join(vaultRoot, `${normalized}.md`))
    candidates.add(path.join(path.dirname(sourceFile), `${normalized}.md`))
  }

  for (const candidate of candidates) {
    if (await pathExists(candidate)) return true
  }

  const fallbackName = normalized.endsWith(".md")
    ? path.basename(normalized)
    : `${path.basename(normalized)}.md`

  return hasAnyFileWithName(vaultRoot, fallbackName)
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
    const vaultRoot = path.join(context.worktree, args.root)
    const files = await listMarkdownFiles(vaultRoot)
    const findings: LinkFinding[] = []

    for (const file of files) {
      const content = await fs.readFile(file, "utf8")
      const source = path.relative(context.worktree, file)
      const targets = extractTargets(content)
      for (const target of targets) {
        const exists = await resolveTargetExists(vaultRoot, file, target)
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
