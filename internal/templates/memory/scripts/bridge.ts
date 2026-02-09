#!/usr/bin/env npx tsx
// LanceDB Bridge Script for OpenKit CLI
// This script allows the Go CLI to interact with LanceDB
// Usage: npx tsx bridge.ts <command> [args...]

import * as fs from 'fs/promises'
import * as path from 'path'

interface Memory {
  id: string
  project: string
  type: 'decision' | 'pattern' | 'error' | 'spec' | 'context'
  title: string
  content: string
  facts_json?: string
  concepts_json?: string
  files_json?: string
  vector: number[]
  salience: number
  created_at: number
  accessed_at: number
  access_count: number
  expires_at: number | null
}

interface OutputMemory {
  id: string
  project: string
  type: string
  title: string
  content: string
  facts: string[]
  concepts: string[]
  files: string[]
  salience: number
  created_at: number
  accessed_at: number
  access_count: number
  expires_at: number | null
}

function parseJsonField(json: string | undefined): string[] {
  if (!json) return []
  try {
    return JSON.parse(json)
  } catch {
    return []
  }
}

function toOutputMemory(m: Memory): OutputMemory {
  return {
    id: m.id,
    project: m.project,
    type: m.type,
    title: m.title,
    content: m.content,
    facts: parseJsonField(m.facts_json),
    concepts: parseJsonField(m.concepts_json),
    files: parseJsonField(m.files_json),
    salience: m.salience,
    created_at: m.created_at,
    accessed_at: m.accessed_at,
    access_count: m.access_count,
    expires_at: m.expires_at
  }
}

async function getDb(dbPath: string) {
  const lancedb = await import('@lancedb/lancedb')
  return await lancedb.connect(dbPath)
}

async function listMemories(dbPath: string, options: { type?: string; limit?: number }) {
  try {
    const db = await getDb(dbPath)
    const tables = await db.tableNames()
    
    if (!tables.includes('memories')) {
      console.log(JSON.stringify({ memories: [], count: 0 }))
      return
    }

    const table = await db.openTable('memories')
    let results = await table.query().limit(options.limit || 100).toArray()

    // Filter by type if specified
    if (options.type) {
      results = results.filter((r: any) => r.type === options.type)
    }

    // Sort by accessed_at descending
    results.sort((a: any, b: any) => (b.accessed_at || 0) - (a.accessed_at || 0))

    const memories = results.map((r: any) => toOutputMemory(r as Memory))
    console.log(JSON.stringify({ memories, count: memories.length }))
  } catch (err: any) {
    console.error(JSON.stringify({ error: err.message }))
    process.exit(1)
  }
}

async function searchMemories(dbPath: string, query: string, limit: number) {
  try {
    const db = await getDb(dbPath)
    const tables = await db.tableNames()
    
    if (!tables.includes('memories')) {
      console.log(JSON.stringify({ memories: [], count: 0 }))
      return
    }

    const table = await db.openTable('memories')
    const allResults = await table.query().limit(1000).toArray()

    // Simple text search (semantic search would require embeddings)
    const queryLower = query.toLowerCase()
    const queryTerms = queryLower.split(/\s+/).filter(t => t.length > 2)

    const scored = allResults.map((r: any) => {
      const text = `${r.title || ''} ${r.content || ''}`.toLowerCase()
      let score = 0
      for (const term of queryTerms) {
        if (text.includes(term)) score += 1
      }
      return { record: r, score }
    })

    const filtered = scored
      .filter(s => s.score > 0)
      .sort((a, b) => b.score - a.score)
      .slice(0, limit)
      .map(s => toOutputMemory(s.record as Memory))

    console.log(JSON.stringify({ memories: filtered, count: filtered.length }))
  } catch (err: any) {
    console.error(JSON.stringify({ error: err.message }))
    process.exit(1)
  }
}

async function getStats(dbPath: string) {
  try {
    const db = await getDb(dbPath)
    const tables = await db.tableNames()
    
    if (!tables.includes('memories')) {
      console.log(JSON.stringify({
        total: 0,
        totalTokens: 0,
        byType: {},
        oldestAccess: null,
        newestAccess: null
      }))
      return
    }

    const table = await db.openTable('memories')
    const results = await table.query().limit(10000).toArray()

    const byType: Record<string, number> = {}
    let totalTokens = 0
    let oldestAccess = Infinity
    let newestAccess = 0

    for (const r of results as Memory[]) {
      byType[r.type] = (byType[r.type] || 0) + 1
      totalTokens += Math.ceil((r.content?.length || 0) / 4)
      if (r.accessed_at < oldestAccess) oldestAccess = r.accessed_at
      if (r.accessed_at > newestAccess) newestAccess = r.accessed_at
    }

    console.log(JSON.stringify({
      total: results.length,
      totalTokens,
      byType,
      oldestAccess: oldestAccess === Infinity ? null : oldestAccess,
      newestAccess: newestAccess === 0 ? null : newestAccess
    }))
  } catch (err: any) {
    console.error(JSON.stringify({ error: err.message }))
    process.exit(1)
  }
}

async function exportMemories(dbPath: string, outputFile: string) {
  try {
    const db = await getDb(dbPath)
    const tables = await db.tableNames()
    
    if (!tables.includes('memories')) {
      await fs.writeFile(outputFile, JSON.stringify([], null, 2))
      console.log(JSON.stringify({ exported: 0, file: outputFile }))
      return
    }

    const table = await db.openTable('memories')
    const results = await table.query().limit(10000).toArray()
    const memories = results.map((r: any) => toOutputMemory(r as Memory))

    await fs.writeFile(outputFile, JSON.stringify(memories, null, 2))
    console.log(JSON.stringify({ exported: memories.length, file: outputFile }))
  } catch (err: any) {
    console.error(JSON.stringify({ error: err.message }))
    process.exit(1)
  }
}

async function pruneMemories(dbPath: string, options: { ttlDays: number; unusedDays: number; maxCount: number; dryRun: boolean }) {
  try {
    const db = await getDb(dbPath)
    const tables = await db.tableNames()
    
    if (!tables.includes('memories')) {
      console.log(JSON.stringify({ deleted: 0, expired: 0, unused: 0, overCap: 0 }))
      return
    }

    const table = await db.openTable('memories')
    const results = await table.query().limit(10000).toArray() as Memory[]

    const now = Date.now()
    const ttlMs = options.ttlDays * 24 * 60 * 60 * 1000
    const unusedMs = options.unusedDays * 24 * 60 * 60 * 1000

    const toDelete: string[] = []
    let expired = 0
    let unused = 0
    let overCap = 0

    // Check expired
    for (const m of results) {
      if (m.expires_at && m.expires_at < now) {
        toDelete.push(m.id)
        expired++
      }
    }

    // Check unused (not accessed recently and low access count)
    for (const m of results) {
      if (!toDelete.includes(m.id)) {
        if (m.accessed_at < (now - unusedMs) && m.access_count < 2) {
          toDelete.push(m.id)
          unused++
        }
      }
    }

    // Check over cap (sort by salience, keep highest)
    const remaining = results.filter(m => !toDelete.includes(m.id))
    if (remaining.length > options.maxCount) {
      remaining.sort((a, b) => b.salience - a.salience)
      const toRemove = remaining.slice(options.maxCount)
      for (const m of toRemove) {
        toDelete.push(m.id)
        overCap++
      }
    }

    if (!options.dryRun && toDelete.length > 0) {
      for (const id of toDelete) {
        try {
          await table.delete(`id = '${id}'`)
        } catch {
          // Ignore individual delete errors
        }
      }
    }

    console.log(JSON.stringify({
      deleted: options.dryRun ? 0 : toDelete.length,
      wouldDelete: options.dryRun ? toDelete.length : 0,
      expired,
      unused,
      overCap,
      dryRun: options.dryRun
    }))
  } catch (err: any) {
    console.error(JSON.stringify({ error: err.message }))
    process.exit(1)
  }
}

async function deleteMemory(dbPath: string, id: string) {
  try {
    const db = await getDb(dbPath)
    const tables = await db.tableNames()
    
    if (!tables.includes('memories')) {
      console.log(JSON.stringify({ deleted: false, error: 'No memories table' }))
      return
    }

    const table = await db.openTable('memories')
    await table.delete(`id = '${id}'`)
    console.log(JSON.stringify({ deleted: true, id }))
  } catch (err: any) {
    console.error(JSON.stringify({ error: err.message }))
    process.exit(1)
  }
}

// Main entry point
async function main() {
  const args = process.argv.slice(2)
  const command = args[0]
  
  // Default paths - can be overridden with --db flag
  let dbPath = '.opencode/memory/index.lance'
  const dbIndex = args.indexOf('--db')
  if (dbIndex !== -1 && args[dbIndex + 1]) {
    dbPath = args[dbIndex + 1]
  }

  switch (command) {
    case 'list': {
      const typeIndex = args.indexOf('--type')
      const limitIndex = args.indexOf('--limit')
      await listMemories(dbPath, {
        type: typeIndex !== -1 ? args[typeIndex + 1] : undefined,
        limit: limitIndex !== -1 ? parseInt(args[limitIndex + 1]) : 100
      })
      break
    }

    case 'search': {
      const query = args[1] || ''
      const limitIndex = args.indexOf('--limit')
      const limit = limitIndex !== -1 ? parseInt(args[limitIndex + 1]) : 10
      await searchMemories(dbPath, query, limit)
      break
    }

    case 'stats': {
      await getStats(dbPath)
      break
    }

    case 'export': {
      const outputFile = args[1] || 'memories-export.json'
      await exportMemories(dbPath, outputFile)
      break
    }

    case 'prune': {
      const ttlIndex = args.indexOf('--ttl')
      const unusedIndex = args.indexOf('--unused')
      const maxIndex = args.indexOf('--max')
      const dryRun = args.includes('--dry-run')
      
      await pruneMemories(dbPath, {
        ttlDays: ttlIndex !== -1 ? parseInt(args[ttlIndex + 1]) : 90,
        unusedDays: unusedIndex !== -1 ? parseInt(args[unusedIndex + 1]) : 30,
        maxCount: maxIndex !== -1 ? parseInt(args[maxIndex + 1]) : 1000,
        dryRun
      })
      break
    }

    case 'delete': {
      const id = args[1]
      if (!id) {
        console.error(JSON.stringify({ error: 'Memory ID required' }))
        process.exit(1)
      }
      await deleteMemory(dbPath, id)
      break
    }

    default:
      console.error(JSON.stringify({ 
        error: 'Unknown command',
        usage: 'bridge.ts <list|search|stats|export|prune|delete> [args...]',
        commands: {
          list: '--type <type> --limit <n>',
          search: '<query> --limit <n>',
          stats: '',
          export: '<output-file>',
          prune: '--ttl <days> --unused <days> --max <count> --dry-run',
          delete: '<memory-id>'
        }
      }))
      process.exit(1)
  }
}

main().catch(err => {
  console.error(JSON.stringify({ error: err.message }))
  process.exit(1)
})
