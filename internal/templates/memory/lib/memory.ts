// Core memory management logic
// Handles extraction, storage, and retrieval of semantic memories

import { LanceDBStorage } from './storage.ts'
import { EmbeddingService } from './embeddings.ts'

export interface Memory {
  id: string
  project: string
  type: 'decision' | 'pattern' | 'error' | 'spec' | 'context'
  title: string
  content: string
  facts?: string[]
  concepts?: string[]
  files?: string[]
  salience: number
  created_at: number
  accessed_at: number
  access_count: number
  expires_at: number | null
}

export interface MemoryConfig {
  version: string
  embedding: {
    model: string
    runtime: string
  }
  retrieval: {
    max_results: number
    min_similarity: number
    token_budget: number
  }
  curation: {
    ttl_days: number
    max_per_project: number
    prune_unused_after_days: number
  }
  extraction: {
    on_session_idle: boolean
    patterns: string[]
  }
  debug: {
    verbose: boolean
    show_injection_indicator: boolean
  }
}

export class SemanticMemory {
  private storage: LanceDBStorage
  private embedder: EmbeddingService
  private config: MemoryConfig
  private configPath: string
  private sessionCache: Memory[] = []

  constructor({ configPath, dbPath }: { configPath: string; dbPath: string }) {
    this.configPath = configPath
    this.storage = new LanceDBStorage({ dbPath })
    this.embedder = new EmbeddingService({ modelPath: '~/.cache/opencode/models/nomic-embed-text' })
  }

  async initialize() {
    // Load config
    this.config = await this.loadConfig()
    
    // Initialize storage
    await this.storage.initialize()
    
    // Initialize embedder
    await this.embedder.initialize()

    // Note: GC/pruning disabled for now to avoid API issues
    // Can be called manually via pruneMemories() if needed
  }

  private async loadConfig(): Promise<MemoryConfig> {
    try {
      const fs = await import('fs/promises')
      const content = await fs.readFile(this.configPath, 'utf-8')
      return JSON.parse(content)
    } catch (err) {
      // Return defaults if config missing
      return {
        version: '1.0.0',
        embedding: { model: 'nomic-embed-text', runtime: 'onnx' },
        retrieval: { max_results: 10, min_similarity: 0.7, token_budget: 4000 },
        curation: { ttl_days: 90, max_per_project: 500, prune_unused_after_days: 30 },
        extraction: { on_session_idle: true, patterns: ['decision', 'architecture', 'pattern', 'fix', 'solution'] },
        debug: { verbose: false, show_injection_indicator: true }
      }
    }
  }

  async extractFromSession(sessionId: string, client: any) {
    if (!this.config.extraction.on_session_idle) {
      return
    }

    try {
      // Get session messages
      const messages = await client.getSessionMessages?.(sessionId)
      if (!messages || messages.length === 0) {
        return
      }

      // Detect decisions using pattern matching
      const decisions = this.detectDecisions(messages)
      if (decisions.length === 0) {
        if (this.config.debug.verbose) {
          console.log(`[semantic-memory] No decisions detected in session ${sessionId}`)
        }
        return
      }

      // Store each decision as a memory
      for (const decision of decisions) {
        const vector = await this.embedder.embed(decision.content)
        await this.storage.createMemory({
          ...decision,
          vector,
          project: this.getProjectName(),
          expires_at: Date.now() + (this.config.curation.ttl_days * 24 * 60 * 60 * 1000)
        })
      }

      if (this.config.debug.verbose) {
        console.log(`[semantic-memory] Extracted ${decisions.length} memories from session ${sessionId}`)
      }
    } catch (err) {
      console.error('[semantic-memory] Extraction failed:', err)
    }
  }

  private detectDecisions(messages: any[]): Omit<Memory, 'id' | 'vector' | 'project' | 'created_at' | 'accessed_at' | 'access_count' | 'expires_at'>[] {
    const decisions: any[] = []
    const patterns = this.config.extraction.patterns

    for (const msg of messages) {
      if (msg.role !== 'assistant') continue

      const content = msg.content?.toString() || ''
      const lowerContent = content.toLowerCase()

      // Check if message contains decision patterns
      const hasPattern = patterns.some(pattern => lowerContent.includes(pattern.toLowerCase()))
      if (!hasPattern) continue

      // Extract structured information
      const title = this.extractTitle(content)
      const facts = this.extractFacts(content)
      const concepts = this.extractConcepts(content)
      const files = this.extractFiles(content)
      const type = this.classifyType(lowerContent)
      const salience = this.calculateSalience(content, facts, concepts)

      decisions.push({
        type,
        title,
        content: content.slice(0, 2000), // Limit content size
        facts,
        concepts,
        files,
        salience
      })
    }

    return decisions
  }

  private extractTitle(content: string): string {
    // Extract first meaningful line as title
    const lines = content.split('\n').filter(l => l.trim().length > 10)
    if (lines.length === 0) return 'Decision'
    
    let title = lines[0].trim()
    // Remove markdown headers
    title = title.replace(/^#+\s*/, '')
    // Limit length
    return title.slice(0, 100)
  }

  private extractFacts(content: string): string[] {
    const facts: string[] = []
    const lines = content.split('\n')

    for (const line of lines) {
      const trimmed = line.trim()
      // Extract bullet points and numbered lists
      if (/^[-*•]\s/.test(trimmed) || /^\d+\.\s/.test(trimmed)) {
        const fact = trimmed.replace(/^[-*•]\s|^\d+\.\s/, '').trim()
        if (fact.length > 10 && fact.length < 200) {
          facts.push(fact)
        }
      }
    }

    return facts.slice(0, 10) // Limit to 10 facts
  }

  private extractConcepts(content: string): string[] {
    const concepts = new Set<string>()
    
    // Extract code blocks language identifiers
    const codeBlockRegex = /```(\w+)/g
    let match
    while ((match = codeBlockRegex.exec(content)) !== null) {
      concepts.add(match[1])
    }

    // Extract technical terms (capitalized words, acronyms)
    const termRegex = /\b[A-Z][A-Za-z]+\b|\b[A-Z]{2,}\b/g
    const terms = content.match(termRegex) || []
    terms.forEach(term => {
      if (term.length > 2 && term.length < 30) {
        concepts.add(term)
      }
    })

    return Array.from(concepts).slice(0, 20) // Limit to 20 concepts
  }

  private extractFiles(content: string): string[] {
    const files = new Set<string>()
    
    // Extract file paths (basic patterns)
    const fileRegex = /[\w-]+\/[\w\/-]+\.\w+/g
    const matches = content.match(fileRegex) || []
    matches.forEach(file => files.add(file))

    return Array.from(files).slice(0, 10) // Limit to 10 files
  }

  private classifyType(content: string): Memory['type'] {
    if (content.includes('decision') || content.includes('chose') || content.includes('selected')) {
      return 'decision'
    }
    if (content.includes('pattern') || content.includes('architecture') || content.includes('design')) {
      return 'pattern'
    }
    if (content.includes('error') || content.includes('bug') || content.includes('fix')) {
      return 'error'
    }
    if (content.includes('spec') || content.includes('requirement') || content.includes('acceptance')) {
      return 'spec'
    }
    return 'context'
  }

  private calculateSalience(content: string, facts: string[], concepts: string[]): number {
    let score = 0.5 // Base score

    // More facts = more salient
    score += Math.min(facts.length * 0.05, 0.2)

    // More concepts = more salient
    score += Math.min(concepts.length * 0.02, 0.2)

    // Longer content = more salient (up to a point)
    const wordCount = content.split(/\s+/).length
    score += Math.min(wordCount / 1000, 0.1)

    return Math.min(score, 1.0)
  }

  private getProjectName(): string {
    // Extract project name from current working directory
    const cwd = process.cwd()
    const parts = cwd.split('/')
    return parts[parts.length - 1] || 'default'
  }

  async getRelevantContext(query: string, limit?: number): Promise<Memory[]> {
    try {
      const maxResults = limit || this.config.retrieval.max_results
      const vector = await this.embedder.embed(query)
      
      // Search with similarity threshold
      let memories = await this.storage.searchMemories(
        vector,
        maxResults * 2, // Get more results for token budget filtering
        this.config.retrieval.min_similarity
      )

      // Sort by salience * similarity (if similarity available)
      memories.sort((a, b) => b.salience - a.salience)

      // Apply token budget
      memories = this.applyTokenBudget(memories, this.config.retrieval.token_budget)

      // Update access metadata
      for (const memory of memories) {
        await this.storage.updateAccessMetadata(memory.id)
      }

      return memories.slice(0, maxResults)
    } catch (err) {
      console.error('[semantic-memory] Context retrieval failed:', err)
      return []
    }
  }

  private applyTokenBudget(memories: Memory[], budget: number): Memory[] {
    const result: Memory[] = []
    let currentTokens = 0

    for (const memory of memories) {
      // Rough estimate: 1 token ≈ 4 characters
      const memoryTokens = Math.ceil(memory.content.length / 4)
      
      if (currentTokens + memoryTokens <= budget) {
        result.push(memory)
        currentTokens += memoryTokens
      } else {
        break
      }
    }

    return result
  }

  async search(query: string, limit: number): Promise<Memory[]> {
    // Get all memories and filter by text match
    // This is simpler and more reliable than vector search with hash embeddings
    const all = await this.storage.getAllMemories(1000)
    
    if (!query || query.trim() === '') {
      return all.slice(0, limit)
    }
    
    const queryLower = query.toLowerCase()
    const queryTerms = queryLower.split(/\s+/).filter(t => t.length > 2)
    
    // Score memories by text match
    const scored = all.map(mem => {
      const text = `${mem.title} ${mem.content}`.toLowerCase()
      let score = 0
      
      for (const term of queryTerms) {
        if (text.includes(term)) {
          score += 1
        }
      }
      
      // Boost exact phrase match
      if (text.includes(queryLower)) {
        score += 5
      }
      
      return { mem, score }
    })
    
    // Sort by score and return top results
    return scored
      .filter(s => s.score > 0)
      .sort((a, b) => b.score - a.score)
      .slice(0, limit)
      .map(s => s.mem)
  }

  async getAllMemories(limit: number = 1000): Promise<Memory[]> {
    return await this.storage.getAllMemories(limit)
  }

  async getCount(): Promise<number> {
    return await this.storage.count()
  }

  /**
   * Manually create a memory entry
   * Used by memory_save tool for explicit knowledge capture
   */
  async createMemory(input: {
    type: Memory['type']
    title: string
    content: string
    files?: string[]
    facts?: string[]
    concepts?: string[]
    salience?: number
  }): Promise<string> {
    try {
      const vector = await this.embedder.embed(input.content)
      
      const memoryData = {
        type: input.type,
        title: input.title,
        content: input.content,
        files: input.files || [],
        facts: input.facts || this.extractFacts(input.content),
        concepts: input.concepts || this.extractConcepts(input.content),
        salience: input.salience || this.calculateSalience(input.content, input.facts || [], input.concepts || []),
        vector,
        project: this.getProjectName(),
        expires_at: Date.now() + (this.config.curation.ttl_days * 24 * 60 * 60 * 1000)
      }

      const id = await this.storage.createMemory(memoryData)

      if (this.config.debug.verbose) {
        console.log(`[semantic-memory] Created memory: ${id} (${input.type}: ${input.title})`)
      }

      return id
    } catch (err) {
      console.error('[semantic-memory] Failed to create memory:', err)
      throw err
    }
  }

  setSessionCache(memories: Memory[]) {
    this.sessionCache = memories
  }

  getSessionCache(): Memory[] {
    return this.sessionCache
  }

  async pruneMemories() {
    try {
      // Delete expired memories (TTL-based)
      const deletedExpired = await this.storage.deleteExpired()

      // Delete unused memories (access-based)
      const unusedThreshold = Date.now() - (this.config.curation.prune_unused_after_days * 24 * 60 * 60 * 1000)
      const deletedUnused = await this.pruneUnusedMemories(unusedThreshold)

      // Enforce hard cap (LRU-based)
      const deletedOverCap = await this.enforceHardCap()

      if (this.config.debug.verbose) {
        console.log(`[semantic-memory] GC complete: ${deletedExpired} expired, ${deletedUnused} unused, ${deletedOverCap} over-cap`)
      }
    } catch (err) {
      console.error('[semantic-memory] GC failed:', err)
    }
  }

  private async pruneUnusedMemories(threshold: number): Promise<number> {
    // Query all memories
    const allMemories = await this.storage.searchMemories(
      new Array(768).fill(0), // Dummy vector
      10000, // Large limit
      0 // No similarity threshold
    )

    let deleted = 0
    for (const memory of allMemories) {
      // Delete if not accessed recently and low access count
      if (memory.accessed_at < threshold && memory.access_count < 2) {
        await this.storage.deleteMemory(memory.id)
        deleted++
      }
    }

    return deleted
  }

  private async enforceHardCap(): Promise<number> {
    const maxPerProject = this.config.curation.max_per_project
    const projectName = this.getProjectName()

    // Get all memories for current project
    const allMemories = await this.storage.searchMemories(
      new Array(768).fill(0), // Dummy vector
      10000, // Large limit
      0 // No similarity threshold
    )

    const projectMemories = allMemories.filter(m => m.project === projectName)

    if (projectMemories.length <= maxPerProject) {
      return 0
    }

    // Sort by LRU: accessed_at ASC (oldest first)
    projectMemories.sort((a, b) => a.accessed_at - b.accessed_at)

    // Delete oldest memories beyond cap
    const toDelete = projectMemories.slice(0, projectMemories.length - maxPerProject)
    for (const memory of toDelete) {
      await this.storage.deleteMemory(memory.id)
    }

    return toDelete.length
  }
}
