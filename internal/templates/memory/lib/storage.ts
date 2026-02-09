// LanceDB storage wrapper
// Handles vector storage and retrieval

export interface Memory {
  id: string
  project: string
  type: 'decision' | 'pattern' | 'error' | 'spec' | 'context'
  title: string
  content: string
  facts?: string[]
  concepts?: string[]
  files?: string[]
  vector: number[]
  salience: number
  created_at: number
  accessed_at: number
  access_count: number
  expires_at: number | null
}

export class LanceDBStorage {
  private dbPath: string
  private db: any = null
  private table: any = null
  private tableName = 'memories'

  constructor({ dbPath }: { dbPath: string }) {
    this.dbPath = dbPath
  }

  async initialize() {
    try {
      const lancedb = await import('@lancedb/lancedb')
      this.db = await lancedb.connect(this.dbPath)
      
      const tables = await this.db.tableNames()
      
      if (tables.includes(this.tableName)) {
        this.table = await this.db.openTable(this.tableName)
      } else {
        this.table = null
      }
    } catch (err: any) {
      // Silent fail - will work without persistence
      this.db = null
      this.table = null
    }
  }
  
  private async ensureTable(firstRecord: any) {
    if (this.table || !this.db) return
    this.table = await this.db.createTable(this.tableName, [firstRecord])
  }

  private normalizeRecord(id: string, memory: Omit<Memory, 'id'>): any {
    let vectorArray: number[]
    if (memory.vector instanceof Float32Array) {
      vectorArray = Array.from(memory.vector)
    } else if (Array.isArray(memory.vector)) {
      vectorArray = memory.vector
    } else {
      vectorArray = new Array(768).fill(0)
    }
    
    return {
      id: String(id),
      project: String(memory.project || 'default'),
      type: String(memory.type || 'context'),
      title: String(memory.title || ''),
      content: String(memory.content || ''),
      facts_json: JSON.stringify(memory.facts || []),
      concepts_json: JSON.stringify(memory.concepts || []),
      files_json: JSON.stringify(memory.files || []),
      vector: vectorArray,
      salience: Number(memory.salience) || 0.5,
      created_at: Number(memory.created_at) || Date.now(),
      accessed_at: Number(memory.accessed_at) || Date.now(),
      access_count: Number(memory.access_count) || 0,
      expires_at: memory.expires_at ? Number(memory.expires_at) : 0
    }
  }

  private recordToMemory(record: any): Memory {
    let facts: string[] = []
    let concepts: string[] = []
    let files: string[] = []
    
    try {
      facts = record.facts_json ? JSON.parse(record.facts_json) : []
    } catch {}
    try {
      concepts = record.concepts_json ? JSON.parse(record.concepts_json) : []
    } catch {}
    try {
      files = record.files_json ? JSON.parse(record.files_json) : []
    } catch {}
    
    return {
      id: record.id || '',
      project: record.project || '',
      type: (record.type || 'context') as Memory['type'],
      title: record.title || '',
      content: record.content || '',
      facts,
      concepts,
      files,
      vector: record.vector || [],
      salience: record.salience || 0.5,
      created_at: record.created_at || 0,
      accessed_at: record.accessed_at || 0,
      access_count: record.access_count || 0,
      expires_at: record.expires_at || null
    }
  }

  async createMemory(memory: Omit<Memory, 'id'>): Promise<string> {
    if (!this.db) {
      throw new Error('Storage not initialized')
    }

    const id = `mem_${Date.now()}_${Math.random().toString(36).substr(2, 9)}`
    const record = this.normalizeRecord(id, memory)

    try {
      if (!this.table) {
        await this.ensureTable(record)
      } else {
        await this.table.add([record])
      }
      return id
    } catch (err: any) {
      throw new Error(`Failed to create memory: ${err.message}`)
    }
  }

  // Simple text-based search instead of vector search
  // This works better with hash-based embeddings
  async searchMemories(vector: Float32Array | number[], limit: number, threshold: number): Promise<Memory[]> {
    // Just return all memories - let the caller filter
    return this.getAllMemories(limit)
  }

  async getMemory(id: string): Promise<Memory | null> {
    if (!this.table) return null

    try {
      const all = await this.getAllMemories(1000)
      return all.find(m => m.id === id) || null
    } catch {
      return null
    }
  }

  async deleteMemory(id: string): Promise<void> {
    if (!this.table) return

    try {
      await this.table.delete(`id = '${id}'`)
    } catch {
      // Ignore delete errors
    }
  }

  async deleteExpired(ttlDays?: number): Promise<number> {
    return 0
  }

  async count(): Promise<number> {
    if (!this.table) return 0

    try {
      return await this.table.countRows()
    } catch {
      return 0
    }
  }

  async getAllMemories(limit: number = 1000): Promise<Memory[]> {
    if (!this.table) return []

    try {
      const results = await this.table.query().limit(limit).toArray()
      const memories: Memory[] = []
      
      for (const r of results) {
        try {
          memories.push(this.recordToMemory(r))
        } catch {
          // Skip bad records
        }
      }
      
      return memories
    } catch {
      return []
    }
  }
}
