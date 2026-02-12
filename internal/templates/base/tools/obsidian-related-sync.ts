import { tool } from "@opencode-ai/plugin"
import fs from "node:fs/promises"
import path from "node:path"

function normalizeLinks(links: string[]): string[] {
  const seen = new Set<string>()
  const out: string[] = []
  for (const item of links) {
    const value = item.trim()
    if (!value || seen.has(value)) continue
    seen.add(value)
    out.push(value)
  }
  return out
}

function buildRelatedSection(links: string[]): string {
  const lines = ["## Related", ""]
  for (const link of links) lines.push(`- [[${link}]]`)
  return `${lines.join("\n")}\n`
}

export default tool({
  description: "Create or replace Related section in docs",
  args: {
    file: tool.schema.string().describe("Target markdown file path"),
    related: tool.schema
      .array(tool.schema.string())
      .min(1)
      .describe("List of internal docs paths without brackets"),
  },
  async execute(args, context) {
    const target = path.join(context.worktree, args.file)
    const content = await fs.readFile(target, "utf8")
    const section = buildRelatedSection(normalizeLinks(args.related))

    const marker = /\n## Related\n[\s\S]*$/m
    const updated = marker.test(content)
      ? content.replace(marker, `\n${section}`)
      : `${content.trimEnd()}\n\n${section}`

    await fs.writeFile(target, updated, "utf8")
    return `Updated Related section in ${args.file}`
  },
})
