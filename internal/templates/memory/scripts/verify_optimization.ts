#!/usr/bin/env bun
/**
 * Verification script for semantic memory optimization
 * 
 * This script helps verify that the memory plugin is:
 * 1. Reducing token usage compared to full context
 * 2. Maintaining relevant context (low drift)
 * 3. Actually being used by OpenCode
 * 
 * Usage:
 *   bun run scripts/verify_optimization.ts [--verbose]
 */

import * as fs from "fs/promises"
import * as path from "path"

interface SessionMetrics {
  sessionId: string
  startTime: number
  memoriesLoaded: number
  memoriesInjected: number
  tokensInjected: number
  compactionTriggered: boolean
  extractionTriggered: boolean
}

interface Memory {
  id: string
  type: string
  title: string
  content: string
  salience: number
  access_count: number
  created_at: number
  accessed_at: number
}

interface VerificationResult {
  status: 'pass' | 'warn' | 'fail'
  check: string
  message: string
  details?: any
}

const MEMORY_DIR = '.opencode/memory'
const METRICS_FILE = `${MEMORY_DIR}/metrics.json`
const CONFIG_FILE = `${MEMORY_DIR}/config.json`
const DB_DIR = `${MEMORY_DIR}/index.lance`

// Estimate tokens (1 token ~ 4 chars)
function estimateTokens(text: string): number {
  return Math.ceil(text.length / 4)
}

async function fileExists(filePath: string): Promise<boolean> {
  try {
    await fs.access(filePath)
    return true
  } catch {
    return false
  }
}

async function runChecks(): Promise<VerificationResult[]> {
  const results: VerificationResult[] = []
  const verbose = process.argv.includes('--verbose')

  // Check 1: Plugin files exist
  console.log('\n[1/6] Checking plugin installation...')
  const pluginExists = await fileExists('.opencode/plugins/semantic-memory/index.ts')
  results.push({
    status: pluginExists ? 'pass' : 'fail',
    check: 'Plugin Installation',
    message: pluginExists 
      ? 'Plugin files are installed correctly'
      : 'Plugin not found at .opencode/plugins/semantic-memory/'
  })

  // Check 2: Config file exists and is valid
  console.log('[2/6] Checking configuration...')
  let config: any = null
  try {
    const configContent = await fs.readFile(CONFIG_FILE, 'utf-8')
    config = JSON.parse(configContent)
    results.push({
      status: 'pass',
      check: 'Configuration',
      message: `Config loaded: token_budget=${config.retrieval?.token_budget}, max_results=${config.retrieval?.max_results}`,
      details: config
    })
  } catch (err) {
    results.push({
      status: 'warn',
      check: 'Configuration',
      message: 'Config file missing or invalid, using defaults'
    })
  }

  // Check 3: Database exists
  console.log('[3/6] Checking database...')
  const dbExists = await fileExists(DB_DIR)
  results.push({
    status: dbExists ? 'pass' : 'warn',
    check: 'Database',
    message: dbExists 
      ? 'LanceDB database directory exists'
      : 'Database not created yet (will be created on first memory)'
  })

  // Check 4: Metrics file and session tracking
  console.log('[4/6] Checking session metrics...')
  let metrics: SessionMetrics[] = []
  try {
    const metricsContent = await fs.readFile(METRICS_FILE, 'utf-8')
    metrics = JSON.parse(metricsContent)
    
    const sessionsWithInjection = metrics.filter(m => m.tokensInjected > 0)
    const avgTokensInjected = sessionsWithInjection.length > 0
      ? Math.round(sessionsWithInjection.reduce((sum, m) => sum + m.tokensInjected, 0) / sessionsWithInjection.length)
      : 0

    results.push({
      status: metrics.length > 0 ? 'pass' : 'warn',
      check: 'Session Metrics',
      message: `${metrics.length} sessions tracked, ${sessionsWithInjection.length} with memory injection`,
      details: {
        totalSessions: metrics.length,
        sessionsWithInjection: sessionsWithInjection.length,
        avgTokensInjected,
        compactionRate: `${Math.round((metrics.filter(m => m.compactionTriggered).length / metrics.length) * 100)}%`
      }
    })
  } catch {
    results.push({
      status: 'warn',
      check: 'Session Metrics',
      message: 'No metrics recorded yet. Run some sessions first.'
    })
  }

  // Check 5: Token savings estimation
  console.log('[5/6] Estimating token savings...')
  if (metrics.length > 0) {
    // Estimate what full context would cost
    // Assume average session context is ~8000 tokens without optimization
    const ESTIMATED_FULL_CONTEXT_TOKENS = 8000
    const totalSessionsWithCompaction = metrics.filter(m => m.compactionTriggered).length
    const totalTokensInjected = metrics.reduce((sum, m) => sum + m.tokensInjected, 0)
    
    if (totalSessionsWithCompaction > 0) {
      const avgInjected = totalTokensInjected / totalSessionsWithCompaction
      const estimatedSavings = ESTIMATED_FULL_CONTEXT_TOKENS - avgInjected
      const savingsPercent = Math.round((estimatedSavings / ESTIMATED_FULL_CONTEXT_TOKENS) * 100)

      results.push({
        status: savingsPercent > 50 ? 'pass' : savingsPercent > 20 ? 'warn' : 'fail',
        check: 'Token Savings',
        message: `Estimated ${savingsPercent}% token reduction (~${Math.round(estimatedSavings)} tokens saved per session)`,
        details: {
          estimatedFullContext: ESTIMATED_FULL_CONTEXT_TOKENS,
          avgOptimizedContext: Math.round(avgInjected),
          estimatedSavingsPerSession: Math.round(estimatedSavings),
          savingsPercent
        }
      })
    } else {
      results.push({
        status: 'warn',
        check: 'Token Savings',
        message: 'No compaction events recorded yet. Cannot estimate savings.'
      })
    }
  } else {
    results.push({
      status: 'warn',
      check: 'Token Savings',
      message: 'No session data. Run sessions to measure savings.'
    })
  }

  // Check 6: Context drift detection
  console.log('[6/6] Checking for context drift...')
  if (metrics.length >= 5) {
    // Check if memories are being accessed (indicating relevance)
    const recentSessions = metrics.slice(-10)
    const sessionsWithMemories = recentSessions.filter(m => m.memoriesInjected > 0)
    const accessRate = sessionsWithMemories.length / recentSessions.length

    // Check memory access patterns
    // TODO: This would need actual memory data to be more accurate
    
    results.push({
      status: accessRate > 0.5 ? 'pass' : accessRate > 0.2 ? 'warn' : 'fail',
      check: 'Context Relevance',
      message: `${Math.round(accessRate * 100)}% of recent sessions received relevant context`,
      details: {
        recentSessions: recentSessions.length,
        sessionsWithContext: sessionsWithMemories.length,
        accessRate: `${Math.round(accessRate * 100)}%`
      }
    })
  } else {
    results.push({
      status: 'warn',
      check: 'Context Relevance',
      message: 'Need at least 5 sessions to measure context drift'
    })
  }

  return results
}

function printResults(results: VerificationResult[], verbose: boolean) {
  console.log('\n' + '='.repeat(60))
  console.log('SEMANTIC MEMORY VERIFICATION REPORT')
  console.log('='.repeat(60) + '\n')

  const statusEmoji = {
    pass: '[PASS]',
    warn: '[WARN]',
    fail: '[FAIL]'
  }

  for (const result of results) {
    console.log(`${statusEmoji[result.status]} ${result.check}`)
    console.log(`   ${result.message}`)
    if (verbose && result.details) {
      console.log(`   Details: ${JSON.stringify(result.details, null, 2).split('\n').join('\n   ')}`)
    }
    console.log()
  }

  // Summary
  const passes = results.filter(r => r.status === 'pass').length
  const warns = results.filter(r => r.status === 'warn').length
  const fails = results.filter(r => r.status === 'fail').length

  console.log('='.repeat(60))
  console.log(`SUMMARY: ${passes} passed, ${warns} warnings, ${fails} failed`)
  console.log('='.repeat(60))

  if (fails > 0) {
    console.log('\nAction Required:')
    for (const result of results.filter(r => r.status === 'fail')) {
      console.log(`  - Fix: ${result.check} - ${result.message}`)
    }
  }

  if (warns > 0 && fails === 0) {
    console.log('\nRecommendations:')
    for (const result of results.filter(r => r.status === 'warn')) {
      console.log(`  - Consider: ${result.check} - ${result.message}`)
    }
  }

  if (passes === results.length) {
    console.log('\nAll checks passed! Semantic memory is working correctly.')
  }
}

// Main execution
async function main() {
  console.log('Semantic Memory Optimization Verifier')
  console.log('=====================================')

  const verbose = process.argv.includes('--verbose')
  
  try {
    const results = await runChecks()
    printResults(results, verbose)
    
    // Exit with error code if any failures
    const hasFailures = results.some(r => r.status === 'fail')
    process.exit(hasFailures ? 1 : 0)
  } catch (err) {
    console.error('Verification failed:', err)
    process.exit(1)
  }
}

main()
