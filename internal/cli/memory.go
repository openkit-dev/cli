package cli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// Memory represents a stored memory entry
type Memory struct {
	ID          string   `json:"id"`
	Project     string   `json:"project"`
	Type        string   `json:"type"`
	Title       string   `json:"title"`
	Content     string   `json:"content"`
	Facts       []string `json:"facts,omitempty"`
	Concepts    []string `json:"concepts,omitempty"`
	Files       []string `json:"files,omitempty"`
	Salience    float64  `json:"salience"`
	CreatedAt   int64    `json:"created_at"`
	AccessedAt  int64    `json:"accessed_at"`
	AccessCount int      `json:"access_count"`
	ExpiresAt   *int64   `json:"expires_at,omitempty"`
}

// SessionMetrics represents metrics for a session
type SessionMetrics struct {
	SessionID           string `json:"sessionId"`
	StartTime           int64  `json:"startTime"`
	MemoriesLoaded      int    `json:"memoriesLoaded"`
	MemoriesInjected    int    `json:"memoriesInjected"`
	TokensInjected      int    `json:"tokensInjected"`
	CompactionTriggered bool   `json:"compactionTriggered"`
	ExtractionTriggered bool   `json:"extractionTriggered"`
}

// MemoryConfig represents the memory plugin configuration
type MemoryConfig struct {
	Version   string `json:"version"`
	Embedding struct {
		Model   string `json:"model"`
		Runtime string `json:"runtime"`
	} `json:"embedding"`
	Retrieval struct {
		MaxResults    int     `json:"max_results"`
		MinSimilarity float64 `json:"min_similarity"`
		TokenBudget   int     `json:"token_budget"`
	} `json:"retrieval"`
	Curation struct {
		TTLDays              int `json:"ttl_days"`
		MaxPerProject        int `json:"max_per_project"`
		PruneUnusedAfterDays int `json:"prune_unused_after_days"`
	} `json:"curation"`
	Extraction struct {
		OnSessionIdle bool     `json:"on_session_idle"`
		Patterns      []string `json:"patterns"`
	} `json:"extraction"`
	Debug struct {
		Verbose                bool `json:"verbose"`
		ShowInjectionIndicator bool `json:"show_injection_indicator"`
	} `json:"debug"`
}

var (
	memoryDir    = ".opencode/memory"
	configFile   = "config.json"
	metricsFile  = "metrics.json"
	memoriesFile = "memories.json" // Fallback for non-LanceDB inspection
)

var memoryCmd = &cobra.Command{
	Use:   "memory",
	Short: "Manage semantic memory for AI context optimization",
	Long: `Manage the semantic memory system that captures and persists context
across AI agent sessions. This helps reduce token usage and maintain
relevant context between conversations.

The memory system automatically extracts decisions, patterns, and important
context from your sessions and retrieves them when relevant.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := cmd.Help(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	},
}

var memoryListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all stored memories",
	Long:  `Display all memories stored in the semantic memory database.`,
	Run: func(cmd *cobra.Command, args []string) {
		runMemoryList(cmd)
	},
}

var memorySearchCmd = &cobra.Command{
	Use:   "search <query>",
	Short: "Search memories by content",
	Long:  `Search for memories containing the specified text or matching keywords.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		query := strings.Join(args, " ")
		runMemorySearch(query, cmd)
	},
}

var memoryStatsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Show memory statistics",
	Long: `Display statistics about the semantic memory system including
total memories, token usage, session history, and configuration.`,
	Run: func(cmd *cobra.Command, args []string) {
		runMemoryStats()
	},
}

var memoryPruneCmd = &cobra.Command{
	Use:   "prune",
	Short: "Clean up old and unused memories",
	Long: `Remove expired memories and enforce storage limits.
This runs the garbage collection process that normally runs automatically.`,
	Run: func(cmd *cobra.Command, args []string) {
		runMemoryPrune(cmd)
	},
}

var memoryExportCmd = &cobra.Command{
	Use:   "export [file]",
	Short: "Export memories to JSON file",
	Long:  `Export all memories to a JSON file for backup or analysis.`,
	Run: func(cmd *cobra.Command, args []string) {
		outputFile := "memories-export.json"
		if len(args) > 0 {
			outputFile = args[0]
		}
		runMemoryExport(outputFile)
	},
}

var memoryConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Show or modify memory configuration",
	Long:  `Display the current memory configuration or modify settings.`,
	Run: func(cmd *cobra.Command, args []string) {
		runMemoryConfig(cmd)
	},
}

var memoryDebugCmd = &cobra.Command{
	Use:   "debug",
	Short: "Debug memory system status",
	Long:  `Show detailed debug information about the memory system.`,
	Run: func(cmd *cobra.Command, args []string) {
		runMemoryDebug()
	},
}

// Flags
var (
	listType      string
	listLimit     int
	searchLimit   int
	pruneForce    bool
	pruneDryRun   bool
	configSet     string
	configVerbose bool
)

func init() {
	// Add memory command to root
	rootCmd.AddCommand(memoryCmd)

	// Add subcommands
	memoryCmd.AddCommand(memoryListCmd)
	memoryCmd.AddCommand(memorySearchCmd)
	memoryCmd.AddCommand(memoryStatsCmd)
	memoryCmd.AddCommand(memoryPruneCmd)
	memoryCmd.AddCommand(memoryExportCmd)
	memoryCmd.AddCommand(memoryConfigCmd)
	memoryCmd.AddCommand(memoryDebugCmd)

	// List flags
	memoryListCmd.Flags().StringVarP(&listType, "type", "t", "", "Filter by type (decision, pattern, error, spec, context)")
	memoryListCmd.Flags().IntVarP(&listLimit, "limit", "n", 20, "Maximum number of memories to show")

	// Search flags
	memorySearchCmd.Flags().IntVarP(&searchLimit, "limit", "n", 10, "Maximum number of results")

	// Prune flags
	memoryPruneCmd.Flags().BoolVarP(&pruneForce, "force", "f", false, "Skip confirmation")
	memoryPruneCmd.Flags().BoolVar(&pruneDryRun, "dry-run", false, "Show what would be deleted without deleting")

	// Config flags
	memoryConfigCmd.Flags().StringVar(&configSet, "set", "", "Set a config value (e.g., --set debug.verbose=true)")
	memoryConfigCmd.Flags().BoolVarP(&configVerbose, "verbose", "v", false, "Enable verbose mode")
}

// Helper functions

func getMemoryDir() string {
	cwd, err := os.Getwd()
	if err != nil {
		return memoryDir
	}
	return filepath.Join(cwd, memoryDir)
}

func loadConfig() (*MemoryConfig, error) {
	configPath := filepath.Join(getMemoryDir(), configFile)
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("config not found: %w", err)
	}

	var config MemoryConfig
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}

	return &config, nil
}

func loadMetrics() ([]SessionMetrics, error) {
	metricsPath := filepath.Join(getMemoryDir(), metricsFile)
	data, err := os.ReadFile(metricsPath)
	if err != nil {
		return nil, err
	}

	var metrics []SessionMetrics
	if err := json.Unmarshal(data, &metrics); err != nil {
		return nil, err
	}

	return metrics, nil
}

func loadMemories() ([]Memory, error) {
	// Try to load from memories.json (fallback/export format)
	memoriesPath := filepath.Join(getMemoryDir(), memoriesFile)
	data, err := os.ReadFile(memoriesPath)
	if err != nil {
		// No memories file - this is expected if using LanceDB
		return nil, nil
	}

	var memories []Memory
	if err := json.Unmarshal(data, &memories); err != nil {
		return nil, err
	}

	return memories, nil
}

func checkMemoryInstalled() bool {
	configPath := filepath.Join(getMemoryDir(), configFile)
	_, err := os.Stat(configPath)
	return err == nil
}

func estimateTokens(text string) int {
	return (len(text) + 3) / 4 // Rough estimate: 1 token ~ 4 chars
}

func formatTime(timestamp int64) string {
	t := time.Unix(timestamp/1000, 0)
	return t.Format("2006-01-02 15:04")
}

// Bridge response types
type BridgeListResponse struct {
	Memories []Memory `json:"memories"`
	Count    int      `json:"count"`
	Error    string   `json:"error,omitempty"`
}

type BridgeStatsResponse struct {
	Total        int            `json:"total"`
	TotalTokens  int            `json:"totalTokens"`
	ByType       map[string]int `json:"byType"`
	OldestAccess *int64         `json:"oldestAccess"`
	NewestAccess *int64         `json:"newestAccess"`
	Error        string         `json:"error,omitempty"`
}

type BridgeExportResponse struct {
	Exported int    `json:"exported"`
	File     string `json:"file"`
	Error    string `json:"error,omitempty"`
}

type BridgePruneResponse struct {
	Deleted     int    `json:"deleted"`
	WouldDelete int    `json:"wouldDelete"`
	Expired     int    `json:"expired"`
	Unused      int    `json:"unused"`
	OverCap     int    `json:"overCap"`
	DryRun      bool   `json:"dryRun"`
	Error       string `json:"error,omitempty"`
}

// callBridge executes the TypeScript bridge script and returns the output
func callBridge(args ...string) ([]byte, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("failed to get working directory: %w", err)
	}

	bridgePath := filepath.Join(cwd, ".opencode", "plugins", "semantic-memory", "scripts", "bridge.ts")

	// Check if bridge exists
	if _, err := os.Stat(bridgePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("bridge script not found at %s", bridgePath)
	}

	// Add --db flag with the correct path
	dbPath := filepath.Join(cwd, ".opencode", "memory", "index.lance")
	fullArgs := append([]string{bridgePath}, args...)
	fullArgs = append(fullArgs, "--db", dbPath)

	// Try npx tsx first, fall back to node --loader
	cmd := exec.Command("npx", append([]string{"tsx"}, fullArgs...)...)
	cmd.Dir = cwd

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()
	if err != nil {
		// Check if it's a JSON error from the bridge
		if stderr.Len() > 0 {
			var errResp struct {
				Error string `json:"error"`
			}
			if json.Unmarshal(stderr.Bytes(), &errResp) == nil && errResp.Error != "" {
				return nil, fmt.Errorf("%s", errResp.Error)
			}
		}
		return nil, fmt.Errorf("bridge execution failed: %w\nstderr: %s", err, stderr.String())
	}

	return stdout.Bytes(), nil
}

// loadMemoriesFromBridge loads memories using the Node.js bridge
func loadMemoriesFromBridge(memType string, limit int) ([]Memory, error) {
	args := []string{"list"}
	if memType != "" {
		args = append(args, "--type", memType)
	}
	if limit > 0 {
		args = append(args, "--limit", fmt.Sprintf("%d", limit))
	}

	output, err := callBridge(args...)
	if err != nil {
		return nil, err
	}

	var resp BridgeListResponse
	if err := json.Unmarshal(output, &resp); err != nil {
		return nil, fmt.Errorf("failed to parse bridge response: %w", err)
	}

	if resp.Error != "" {
		return nil, fmt.Errorf("%s", resp.Error)
	}

	return resp.Memories, nil
}

// searchMemoriesFromBridge searches memories using the Node.js bridge
func searchMemoriesFromBridge(query string, limit int) ([]Memory, error) {
	args := []string{"search", query, "--limit", fmt.Sprintf("%d", limit)}

	output, err := callBridge(args...)
	if err != nil {
		return nil, err
	}

	var resp BridgeListResponse
	if err := json.Unmarshal(output, &resp); err != nil {
		return nil, fmt.Errorf("failed to parse bridge response: %w", err)
	}

	if resp.Error != "" {
		return nil, fmt.Errorf("%s", resp.Error)
	}

	return resp.Memories, nil
}

// Command implementations

func runMemoryList(cmd *cobra.Command) {
	cyan := color.New(color.FgCyan, color.Bold)
	dim := color.New(color.FgHiBlack)
	yellow := color.New(color.FgYellow)

	if !checkMemoryInstalled() {
		exitWithError("Memory plugin not installed. Run 'openkit init --memory' to enable.")
	}

	// Try bridge first (reads from LanceDB)
	memories, err := loadMemoriesFromBridge(listType, listLimit)
	if err != nil {
		// Fall back to JSON file
		memories, err = loadMemories()
		if err != nil {
			printWarning("Could not load memories: " + err.Error())
			printInfo("Make sure npm dependencies are installed in .opencode/")
			return
		}

		// Filter by type if specified (only needed for fallback)
		if listType != "" {
			var filtered []Memory
			for _, m := range memories {
				if m.Type == listType {
					filtered = append(filtered, m)
				}
			}
			memories = filtered
		}

		// Sort by accessed_at (most recent first)
		sort.Slice(memories, func(i, j int) bool {
			return memories[i].AccessedAt > memories[j].AccessedAt
		})

		// Limit results
		if len(memories) > listLimit {
			memories = memories[:listLimit]
		}
	}

	if len(memories) == 0 {
		printInfo("No memories stored yet.")
		printInfo("Memories are automatically extracted from your OpenCode sessions.")
		return
	}

	cyan.Printf("\nMemories (%d total)\n", len(memories))
	cyan.Println(strings.Repeat("=", 50))
	fmt.Println()

	for _, m := range memories {
		typeColor := color.New(color.FgGreen)
		switch m.Type {
		case "decision":
			typeColor = color.New(color.FgCyan)
		case "error":
			typeColor = color.New(color.FgRed)
		case "pattern":
			typeColor = color.New(color.FgMagenta)
		case "spec":
			typeColor = color.New(color.FgBlue)
		}

		typeColor.Printf("  [%s] ", strings.ToUpper(m.Type))
		fmt.Printf("%s\n", m.Title)

		// Truncate content
		content := m.Content
		if len(content) > 100 {
			content = content[:100] + "..."
		}
		dim.Printf("    %s\n", content)

		yellow.Printf("    Salience: %.2f | Accessed: %dx | ", m.Salience, m.AccessCount)
		dim.Printf("Last: %s\n", formatTime(m.AccessedAt))

		if len(m.Files) > 0 {
			dim.Printf("    Files: %s\n", strings.Join(m.Files, ", "))
		}
		fmt.Println()
	}
}

func runMemorySearch(query string, cmd *cobra.Command) {
	cyan := color.New(color.FgCyan, color.Bold)
	dim := color.New(color.FgHiBlack)
	yellow := color.New(color.FgYellow)

	if !checkMemoryInstalled() {
		exitWithError("Memory plugin not installed. Run 'openkit init --memory' to enable.")
	}

	// Try bridge first (reads from LanceDB)
	results, err := searchMemoriesFromBridge(query, searchLimit)
	if err != nil {
		// Fall back to JSON file
		memories, loadErr := loadMemories()
		if loadErr != nil || memories == nil {
			printWarning("Could not search memories: " + err.Error())
			printInfo("Make sure npm dependencies are installed in .opencode/")
			return
		}

		// Simple text search (not semantic - that requires embeddings)
		queryLower := strings.ToLower(query)
		results = nil

		for _, m := range memories {
			if strings.Contains(strings.ToLower(m.Title), queryLower) ||
				strings.Contains(strings.ToLower(m.Content), queryLower) ||
				containsAny(m.Concepts, queryLower) ||
				containsAny(m.Facts, queryLower) {
				results = append(results, m)
			}
		}

		// Limit results
		if len(results) > searchLimit {
			results = results[:searchLimit]
		}
	}

	cyan.Printf("\nSearch Results for '%s' (%d found)\n", query, len(results))
	cyan.Println(strings.Repeat("=", 50))
	fmt.Println()

	if len(results) == 0 {
		printInfo("No memories match your search.")
		printInfo("Try different keywords or use the memory_query tool in OpenCode for semantic search.")
		return
	}

	for _, m := range results {
		typeColor := color.New(color.FgGreen)
		switch m.Type {
		case "decision":
			typeColor = color.New(color.FgCyan)
		case "error":
			typeColor = color.New(color.FgRed)
		case "pattern":
			typeColor = color.New(color.FgMagenta)
		}

		typeColor.Printf("  [%s] ", strings.ToUpper(m.Type))
		fmt.Printf("%s\n", m.Title)

		content := m.Content
		if len(content) > 150 {
			content = content[:150] + "..."
		}
		dim.Printf("    %s\n", content)
		yellow.Printf("    Salience: %.2f | Tokens: ~%d\n", m.Salience, estimateTokens(m.Content))
		fmt.Println()
	}
}

func containsAny(slice []string, query string) bool {
	for _, s := range slice {
		if strings.Contains(strings.ToLower(s), query) {
			return true
		}
	}
	return false
}

func runMemoryStats() {
	cyan := color.New(color.FgCyan, color.Bold)
	green := color.New(color.FgGreen)
	yellow := color.New(color.FgYellow)
	dim := color.New(color.FgHiBlack)

	if !checkMemoryInstalled() {
		exitWithError("Memory plugin not installed. Run 'openkit init --memory' to enable.")
	}

	cyan.Println("\nSemantic Memory Statistics")
	cyan.Println(strings.Repeat("=", 50))
	fmt.Println()

	// Load config
	config, err := loadConfig()
	if err != nil {
		printWarning("Could not load config: " + err.Error())
	}

	// Load metrics
	metrics, _ := loadMetrics()

	// Try to get stats from bridge (LanceDB)
	var totalMemories int
	var totalTokens int
	var byType map[string]int
	var bridgeWorked bool

	output, bridgeErr := callBridge("stats")
	if bridgeErr == nil {
		var stats BridgeStatsResponse
		if json.Unmarshal(output, &stats) == nil && stats.Error == "" {
			totalMemories = stats.Total
			totalTokens = stats.TotalTokens
			byType = stats.ByType
			bridgeWorked = true
		}
	}

	// Fall back to JSON if bridge failed
	if !bridgeWorked {
		memories, _ := loadMemories()
		if memories != nil {
			totalMemories = len(memories)
			byType = make(map[string]int)
			for _, m := range memories {
				totalTokens += estimateTokens(m.Content)
				byType[m.Type]++
			}
			bridgeWorked = true
		}
	}

	// Memory stats
	cyan.Println("Storage")
	cyan.Println("-------")

	if bridgeWorked && totalMemories > 0 {
		fmt.Printf("  Total Memories: %d\n", totalMemories)
		fmt.Printf("  Total Tokens:   ~%d\n", totalTokens)
		fmt.Println()

		cyan.Println("By Type")
		cyan.Println("-------")
		for t, count := range byType {
			fmt.Printf("  %-12s %d\n", t+":", count)
		}
	} else if bridgeWorked {
		dim.Println("  No memories stored yet.")
		dim.Println("  Memories are automatically extracted from your OpenCode sessions.")
	} else {
		dim.Println("  Could not read memory data.")
		dim.Println("  Make sure npm dependencies are installed: cd .opencode && npm install")
	}
	fmt.Println()

	// Session stats
	cyan.Println("Session History")
	cyan.Println("---------------")
	fmt.Println()

	if len(metrics) > 0 {
		totalInjected := 0
		totalTokensInjected := 0
		compactionCount := 0

		for _, m := range metrics {
			totalInjected += m.MemoriesInjected
			totalTokensInjected += m.TokensInjected
			if m.CompactionTriggered {
				compactionCount++
			}
		}

		avgTokens := 0
		if len(metrics) > 0 {
			avgTokens = totalTokensInjected / len(metrics)
		}

		fmt.Printf("  Sessions Tracked:     %d\n", len(metrics))
		fmt.Printf("  Compaction Events:    %d (%.0f%%)\n", compactionCount, float64(compactionCount)/float64(len(metrics))*100)
		fmt.Printf("  Total Memories Injected: %d\n", totalInjected)
		fmt.Printf("  Total Tokens Injected:   ~%d\n", totalTokensInjected)
		fmt.Printf("  Avg Tokens per Session:  ~%d\n", avgTokens)

		// Estimated savings
		estimatedFullContext := 8000 // Rough estimate
		if avgTokens > 0 && avgTokens < estimatedFullContext {
			savings := estimatedFullContext - avgTokens
			savingsPercent := float64(savings) / float64(estimatedFullContext) * 100
			green.Printf("\n  Estimated Savings: ~%d tokens/session (%.0f%%)\n", savings, savingsPercent)
		}
	} else {
		dim.Println("  No session metrics recorded yet")
	}
	fmt.Println()

	// Configuration
	cyan.Println("Configuration")
	cyan.Println("-------------")

	if config != nil {
		fmt.Printf("  Token Budget:    %d\n", config.Retrieval.TokenBudget)
		fmt.Printf("  Max Results:     %d\n", config.Retrieval.MaxResults)
		fmt.Printf("  Min Similarity:  %.2f\n", config.Retrieval.MinSimilarity)
		fmt.Printf("  TTL:             %d days\n", config.Curation.TTLDays)
		fmt.Printf("  Max per Project: %d\n", config.Curation.MaxPerProject)
		fmt.Printf("  Verbose:         %v\n", config.Debug.Verbose)
	} else {
		dim.Println("  Could not load configuration")
	}
	fmt.Println()

	// Health check
	cyan.Println("Health")
	cyan.Println("------")

	dbPath := filepath.Join(getMemoryDir(), "index.lance")
	if _, err := os.Stat(dbPath); err == nil {
		green.Println("  [OK] LanceDB directory exists")
	} else {
		yellow.Println("  [--] LanceDB not initialized (will be created on first memory)")
	}

	if config != nil {
		green.Println("  [OK] Configuration loaded")
	} else {
		yellow.Println("  [--] Configuration missing or invalid")
	}

	if metrics != nil && len(metrics) > 0 {
		green.Printf("  [OK] %d sessions tracked\n", len(metrics))
	} else {
		yellow.Println("  [--] No session metrics yet")
	}

	fmt.Println()
}

func runMemoryPrune(cmd *cobra.Command) {
	cyan := color.New(color.FgCyan, color.Bold)
	yellow := color.New(color.FgYellow)
	green := color.New(color.FgGreen)

	if !checkMemoryInstalled() {
		exitWithError("Memory plugin not installed. Run 'openkit init --memory' to enable.")
	}

	cyan.Println("\nMemory Garbage Collection")
	cyan.Println(strings.Repeat("=", 50))
	fmt.Println()

	config, err := loadConfig()
	if err != nil {
		exitWithError("Could not load config: " + err.Error())
	}

	// Build prune args
	args := []string{"prune",
		"--ttl", fmt.Sprintf("%d", config.Curation.TTLDays),
		"--unused", fmt.Sprintf("%d", config.Curation.PruneUnusedAfterDays),
		"--max", fmt.Sprintf("%d", config.Curation.MaxPerProject),
	}

	// First do a dry run to show what would be deleted
	dryRunArgs := append(args, "--dry-run")
	output, bridgeErr := callBridge(dryRunArgs...)
	if bridgeErr != nil {
		printWarning("Could not analyze memories: " + bridgeErr.Error())
		printInfo("Make sure npm dependencies are installed: cd .opencode && npm install")
		return
	}

	var resp BridgePruneResponse
	if err := json.Unmarshal(output, &resp); err != nil {
		printWarning("Could not parse prune response: " + err.Error())
		return
	}

	if resp.Error != "" {
		printWarning("Prune error: " + resp.Error)
		return
	}

	fmt.Printf("  Expired (TTL > %d days):       %d memories\n", config.Curation.TTLDays, resp.Expired)
	fmt.Printf("  Unused (> %d days, <2 access): %d memories\n", config.Curation.PruneUnusedAfterDays, resp.Unused)
	fmt.Printf("  Over Cap (> %d):               %d memories\n", config.Curation.MaxPerProject, resp.OverCap)
	fmt.Println()

	totalToDelete := resp.WouldDelete

	if totalToDelete == 0 {
		printSuccess("No memories need pruning.")
		return
	}

	if pruneDryRun {
		yellow.Printf("  [DRY RUN] Would delete %d memories\n", totalToDelete)
		return
	}

	if !pruneForce {
		yellow.Printf("  This will delete %d memories. Continue? [y/N] ", totalToDelete)
		var response string
		_, _ = fmt.Scanln(&response)
		if strings.ToLower(response) != "y" {
			printInfo("Cancelled.")
			return
		}
	}

	// Actually prune
	output, bridgeErr = callBridge(args...)
	if bridgeErr != nil {
		printWarning("Prune failed: " + bridgeErr.Error())
		return
	}

	if err := json.Unmarshal(output, &resp); err != nil {
		printWarning("Could not parse prune response: " + err.Error())
		return
	}

	if resp.Error != "" {
		printWarning("Prune error: " + resp.Error)
		return
	}

	green.Printf("  Deleted %d memories\n", resp.Deleted)
}

func runMemoryExport(outputFile string) {
	cyan := color.New(color.FgCyan, color.Bold)

	if !checkMemoryInstalled() {
		exitWithError("Memory plugin not installed. Run 'openkit init --memory' to enable.")
	}

	cyan.Println("\nExport Memories")
	cyan.Println(strings.Repeat("=", 50))
	fmt.Println()

	// Use bridge to export from LanceDB
	output, bridgeErr := callBridge("export", outputFile)
	if bridgeErr != nil {
		// Fall back to JSON file
		memories, err := loadMemories()
		if err != nil || memories == nil {
			printWarning("Could not export memories: " + bridgeErr.Error())
			printInfo("Make sure npm dependencies are installed: cd .opencode && npm install")
			return
		}

		data, err := json.MarshalIndent(memories, "", "  ")
		if err != nil {
			exitWithError("Failed to serialize memories: " + err.Error())
		}

		if err := os.WriteFile(outputFile, data, 0644); err != nil {
			exitWithError("Failed to write file: " + err.Error())
		}

		printSuccess(fmt.Sprintf("Exported %d memories to %s", len(memories), outputFile))
		return
	}

	var resp BridgeExportResponse
	if err := json.Unmarshal(output, &resp); err != nil {
		printWarning("Could not parse export response: " + err.Error())
		return
	}

	if resp.Error != "" {
		printWarning("Export error: " + resp.Error)
		return
	}

	printSuccess(fmt.Sprintf("Exported %d memories to %s", resp.Exported, resp.File))
}

func runMemoryConfig(cmd *cobra.Command) {
	cyan := color.New(color.FgCyan, color.Bold)
	dim := color.New(color.FgHiBlack)

	if !checkMemoryInstalled() {
		exitWithError("Memory plugin not installed. Run 'openkit init --memory' to enable.")
	}

	configPath := filepath.Join(getMemoryDir(), configFile)

	// If --verbose flag, toggle verbose mode
	if configVerbose {
		config, err := loadConfig()
		if err != nil {
			exitWithError("Could not load config: " + err.Error())
		}

		config.Debug.Verbose = !config.Debug.Verbose

		data, err := json.MarshalIndent(config, "", "  ")
		if err != nil {
			exitWithError("Failed to serialize config: " + err.Error())
		}

		if err := os.WriteFile(configPath, data, 0644); err != nil {
			exitWithError("Failed to write config: " + err.Error())
		}

		if config.Debug.Verbose {
			printSuccess("Verbose mode enabled")
		} else {
			printSuccess("Verbose mode disabled")
		}
		return
	}

	// If --set flag, update config
	if configSet != "" {
		parts := strings.SplitN(configSet, "=", 2)
		if len(parts) != 2 {
			exitWithError("Invalid --set format. Use: --set key=value")
		}

		printWarning("Config modification via CLI not fully implemented.")
		printInfo("Edit the config file directly: " + configPath)
		return
	}

	// Otherwise, show config
	cyan.Println("\nMemory Configuration")
	cyan.Println(strings.Repeat("=", 50))
	fmt.Println()

	config, err := loadConfig()
	if err != nil {
		exitWithError("Could not load config: " + err.Error())
	}

	fmt.Println("  Embedding:")
	fmt.Printf("    model:   %s\n", config.Embedding.Model)
	fmt.Printf("    runtime: %s\n", config.Embedding.Runtime)
	fmt.Println()

	fmt.Println("  Retrieval:")
	fmt.Printf("    max_results:    %d\n", config.Retrieval.MaxResults)
	fmt.Printf("    min_similarity: %.2f\n", config.Retrieval.MinSimilarity)
	fmt.Printf("    token_budget:   %d\n", config.Retrieval.TokenBudget)
	fmt.Println()

	fmt.Println("  Curation:")
	fmt.Printf("    ttl_days:              %d\n", config.Curation.TTLDays)
	fmt.Printf("    max_per_project:       %d\n", config.Curation.MaxPerProject)
	fmt.Printf("    prune_unused_after:    %d days\n", config.Curation.PruneUnusedAfterDays)
	fmt.Println()

	fmt.Println("  Extraction:")
	fmt.Printf("    on_session_idle: %v\n", config.Extraction.OnSessionIdle)
	fmt.Printf("    patterns:        %s\n", strings.Join(config.Extraction.Patterns, ", "))
	fmt.Println()

	fmt.Println("  Debug:")
	fmt.Printf("    verbose:              %v\n", config.Debug.Verbose)
	fmt.Printf("    show_injection_indicator: %v\n", config.Debug.ShowInjectionIndicator)
	fmt.Println()

	dim.Printf("  Config file: %s\n", configPath)
	fmt.Println()
}

func runMemoryDebug() {
	cyan := color.New(color.FgCyan, color.Bold)
	green := color.New(color.FgGreen)
	yellow := color.New(color.FgYellow)
	red := color.New(color.FgRed)
	dim := color.New(color.FgHiBlack)

	cyan.Println("\nMemory System Debug")
	cyan.Println(strings.Repeat("=", 50))
	fmt.Println()

	memDir := getMemoryDir()

	// Check installation
	cyan.Println("Installation Check")
	cyan.Println("------------------")

	checks := []struct {
		name string
		path string
	}{
		{"Plugin directory", filepath.Join(filepath.Dir(memDir), "plugins", "semantic-memory")},
		{"Plugin index.ts", filepath.Join(filepath.Dir(memDir), "plugins", "semantic-memory", "index.ts")},
		{"Memory directory", memDir},
		{"Config file", filepath.Join(memDir, configFile)},
		{"LanceDB directory", filepath.Join(memDir, "index.lance")},
		{"Metrics file", filepath.Join(memDir, metricsFile)},
	}

	allOk := true
	for _, check := range checks {
		if _, err := os.Stat(check.path); err == nil {
			green.Printf("  [OK] %s\n", check.name)
		} else {
			if strings.Contains(check.name, "LanceDB") || strings.Contains(check.name, "Metrics") {
				yellow.Printf("  [--] %s (will be created on first use)\n", check.name)
			} else {
				red.Printf("  [!!] %s - NOT FOUND\n", check.name)
				allOk = false
			}
		}
	}
	fmt.Println()

	// Config validation
	cyan.Println("Configuration")
	cyan.Println("-------------")

	config, err := loadConfig()
	if err != nil {
		red.Printf("  [!!] Error loading config: %s\n", err)
	} else {
		green.Println("  [OK] Config loaded successfully")
		fmt.Printf("       Version: %s\n", config.Version)
		fmt.Printf("       Verbose: %v\n", config.Debug.Verbose)
	}
	fmt.Println()

	// Session metrics
	cyan.Println("Session Metrics")
	cyan.Println("---------------")

	metrics, err := loadMetrics()
	if err != nil {
		yellow.Println("  [--] No metrics file (sessions not tracked yet)")
	} else if len(metrics) == 0 {
		yellow.Println("  [--] Metrics file empty")
	} else {
		green.Printf("  [OK] %d sessions tracked\n", len(metrics))

		// Show last 3 sessions
		start := len(metrics) - 3
		if start < 0 {
			start = 0
		}
		fmt.Println()
		dim.Println("  Recent sessions:")
		for _, m := range metrics[start:] {
			fmt.Printf("    %s | Loaded: %d | Injected: %d | Tokens: ~%d\n",
				m.SessionID[:8]+"...", m.MemoriesLoaded, m.MemoriesInjected, m.TokensInjected)
		}
	}
	fmt.Println()

	// Recommendations
	cyan.Println("Recommendations")
	cyan.Println("---------------")

	if !allOk {
		yellow.Println("  - Run 'openkit init --memory' to install the plugin")
	}

	if config != nil && !config.Debug.Verbose {
		dim.Println("  - Enable verbose mode for detailed logs: openkit memory config --verbose")
	}

	if len(metrics) < 5 {
		dim.Println("  - Run more OpenCode sessions to accumulate metrics")
	}

	if config != nil && config.Retrieval.MinSimilarity > 0.8 {
		dim.Println("  - Consider lowering min_similarity if context injection is rare")
	}

	fmt.Println()
}
