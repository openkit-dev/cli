package detection

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// Pattern represents a file/directory/content pattern to match
type Pattern struct {
	Type       string  `json:"type"`             // "file", "directory", "content"
	Pattern    string  `json:"pattern"`          // Glob pattern or regex
	MinMatches int     `json:"min_matches"`      // Minimum matches required
	Weight     float64 `json:"weight"`           // Confidence boost weight
	Reason     string  `json:"reason,omitempty"` // Why this pattern matters
}

// OverlayConfig represents an overlay configuration
type OverlayConfig struct {
	ID          string           `json:"id"`
	Name        string           `json:"name"`
	Description string           `json:"description,omitempty"`
	Condition   OverlayCondition `json:"condition"`
	Adds        OverlayAdds      `json:"adds"`
	Modifies    OverlayModifies  `json:"modifies,omitempty"`
}

// OverlayCondition defines when an overlay should be active
type OverlayCondition struct {
	Any []ConditionRule `json:"any,omitempty"`
	All []ConditionRule `json:"all,omitempty"`
}

// ConditionRule defines a single condition
type ConditionRule struct {
	HasFile      string `json:"has-file,omitempty"`
	HasDirectory string `json:"has-directory,omitempty"`
	HasPattern   string `json:"has-pattern,omitempty"`
}

// OverlayAdds defines resources added by an overlay
type OverlayAdds struct {
	Docs   []DocRef   `json:"docs,omitempty"`
	Skills []SkillRef `json:"skills,omitempty"`
	Agents []string   `json:"agents,omitempty"`
}

// OverlayModifies defines modifications to existing docs/skills
type OverlayModifies struct {
	QUALITY_GATES *DocModifications `json:"QUALITY_GATES.md,omitempty"`
	ACTION_ITEMS  *DocModifications `json:"ACTION_ITEMS.md,omitempty"`
}

// DocRef represents a document reference
type DocRef struct {
	Name         string                 `json:"name"`
	Template     string                 `json:"template"`
	Required     bool                   `json:"required"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Condition    *OverlayCondition      `json:"condition,omitempty"`
}

// DocModifications defines modifications to existing docs
type DocModifications struct {
	AddSection map[string]string `json:"add-section,omitempty"`
	AddItems   []string          `json:"add-items,omitempty"`
}

// DetectionRules defines how to detect a project type
type DetectionRules struct {
	RequiredPatterns    []Pattern `json:"required_patterns"`
	SuggestingPatterns  []Pattern `json:"suggesting_patterns"`
	ConflictingPatterns []Pattern `json:"conflicting_patterns"`
	ConfidenceThreshold float64   `json:"confidence_threshold"` // Minimum score to match (default: 0.7)
}

// ProjectTypeConfig represents a project type configuration
type ProjectTypeConfig struct {
	ID              string         `json:"id"`
	Name            string         `json:"name"`
	Description     string         `json:"description,omitempty"`
	Detection       DetectionRules `json:"detection"`
	BaseContext     BaseContext    `json:"base_context"`
	SuggestOverlays []string       `json:"suggest_overlays,omitempty"`
}

// BaseContext defines base context documentation and skills
type BaseContext struct {
	Docs          []DocTemplate `json:"docs"`
	Skills        []SkillRef    `json:"skills,omitempty"`
	SkillsExclude []string      `json:"skills_exclude,omitempty"`
}

// DocTemplate represents a documentation template
type DocTemplate struct {
	Name         string                 `json:"name"`                // Output filename
	Template     string                 `json:"template"`            // Template name
	Required     bool                   `json:"required"`            // Must exist
	RenameTo     string                 `json:"rename_to,omitempty"` // Rename to different name
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
}

// SkillRef represents a skill reference
type SkillRef struct {
	ID       string `json:"id"`
	Required bool   `json:"required"`
	Priority string `json:"priority,omitempty"`
}

// DetectionResult represents the result of project type detection
type DetectionResult struct {
	ProjectType       string   `json:"project_type"`
	Confidence        float64  `json:"confidence"`
	Evidence          []string `json:"evidence"`
	SuggestedOverlays []string `json:"suggested_overlays"`
	Conflicts         []string `json:"conflicts,omitempty"`
}

// Registry holds all known project types
type Registry struct {
	types      map[string]*ProjectTypeConfig
	overlays   map[string]*OverlayConfig
	projectDir string
}

// LoadRegistry loads project types and overlays from .opencode directory
func LoadRegistry(projectDir string) (*Registry, error) {
	reg := &Registry{
		types:      make(map[string]*ProjectTypeConfig),
		overlays:   make(map[string]*OverlayConfig),
		projectDir: projectDir,
	}

	loadDefaultConfigs(reg)

	// Load project types
	typesDir := filepath.Join(projectDir, ".opencode", "project-types")
	if entries, err := os.ReadDir(typesDir); err == nil {
		for _, entry := range entries {
			if entry.IsDir() {
				continue
			}
			if !strings.HasSuffix(entry.Name(), ".json") {
				continue
			}

			filePath := filepath.Join(typesDir, entry.Name())
			data, err := os.ReadFile(filePath)
			if err != nil {
				return nil, fmt.Errorf("failed to read %s: %w", entry.Name(), err)
			}

			var config ProjectTypeConfig
			if err := json.Unmarshal(data, &config); err != nil {
				return nil, fmt.Errorf("failed to parse %s: %w", entry.Name(), err)
			}

			reg.types[config.ID] = &config
		}
	}

	// Load overlays
	overlaysDir := filepath.Join(projectDir, ".opencode", "overlays")
	if entries, err := os.ReadDir(overlaysDir); err == nil {
		for _, entry := range entries {
			if entry.IsDir() {
				continue
			}
			if !strings.HasSuffix(entry.Name(), ".json") {
				continue
			}

			filePath := filepath.Join(overlaysDir, entry.Name())
			data, err := os.ReadFile(filePath)
			if err != nil {
				return nil, fmt.Errorf("failed to read overlay %s: %w", entry.Name(), err)
			}

			var config OverlayConfig
			if err := json.Unmarshal(data, &config); err != nil {
				return nil, fmt.Errorf("failed to parse overlay %s: %w", entry.Name(), err)
			}

			if config.ID == "" {
				config.ID = strings.TrimSuffix(entry.Name(), ".json")
			}
			reg.overlays[config.ID] = &config
		}
	}

	return reg, nil
}

// Detect detects the project type from the current directory
func (reg *Registry) Detect() (*DetectionResult, error) {
	if reg.projectDir == "" {
		return nil, fmt.Errorf("project directory not set")
	}

	scores := make(map[string]float64)
	evidence := make(map[string][]string)
	conflicts := make(map[string][]string)

	// Score each project type
	for typeID, config := range reg.types {
		score := reg.scoreType(config)
		if score > 0 {
			scores[typeID] = score
			evidence[typeID] = reg.collectEvidence(config)
		}

		// Check for conflicts
		for _, conflict := range config.Detection.ConflictingPatterns {
			if reg.matchesPattern(conflict) {
				conflicts[typeID] = append(conflicts[typeID], conflict.Pattern+" ("+conflict.Reason+")")
			}
		}
	}

	// Find best match (deterministic): score -> required pattern count -> id
	var bestType string
	var bestScore float64
	bestRequiredCount := -1
	typeIDs := make([]string, 0, len(scores))
	for typeID := range scores {
		typeIDs = append(typeIDs, typeID)
	}
	sort.Strings(typeIDs)

	for _, typeID := range typeIDs {
		score := scores[typeID]
		reqCount := len(reg.types[typeID].Detection.RequiredPatterns)
		if score > bestScore || (score == bestScore && reqCount > bestRequiredCount) {
			bestScore = score
			bestRequiredCount = reqCount
			bestType = typeID
		}
	}

	if bestType == "" {
		return &DetectionResult{
			ProjectType: "unknown",
			Confidence:  0,
			Evidence:    []string{"No matching project type detected"},
		}, nil
	}

	// Check for conflicts
	if len(conflicts[bestType]) > 0 {
		return &DetectionResult{
			ProjectType: "unknown",
			Confidence:  0,
			Evidence:    []string{fmt.Sprintf("Conflicting patterns: %v", conflicts[bestType])},
		}, nil
	}

	// Get suggested overlays
	suggestedOverlays := reg.types[bestType].SuggestOverlays

	return &DetectionResult{
		ProjectType:       bestType,
		Confidence:        bestScore,
		Evidence:          evidence[bestType],
		SuggestedOverlays: suggestedOverlays,
	}, nil
}

// scoreType calculates a confidence score for a project type
func (reg *Registry) scoreType(config *ProjectTypeConfig) float64 {
	var score float64

	// Check required patterns (must all match)
	requiredMatches := 0
	for _, pattern := range config.Detection.RequiredPatterns {
		if reg.matchesPattern(pattern) {
			requiredMatches++
		}
	}

	// If required patterns don't all match, fail immediately
	if requiredMatches != len(config.Detection.RequiredPatterns) {
		return 0
	}

	// Base score for passing required patterns
	score = 100.0

	// Add weights for suggesting patterns
	for _, pattern := range config.Detection.SuggestingPatterns {
		if reg.matchesPattern(pattern) {
			score += pattern.Weight * 100.0
		}
	}

	return score
}

// matchesPattern checks if a pattern matches the project
func (reg *Registry) matchesPattern(pattern Pattern) bool {
	switch pattern.Type {
	case "file":
		return reg.hasMatchingFiles(pattern.Pattern, pattern.MinMatches)
	case "directory":
		return reg.hasDirectory(pattern.Pattern)
	case "content":
		return reg.hasMatchingContent(pattern.Pattern)
	}
	return false
}

// hasMatchingFiles checks if files matching pattern exist
func (reg *Registry) hasMatchingFiles(pattern string, minMatches int) bool {
	if minMatches <= 0 {
		minMatches = 1
	}
	matches, err := filepath.Glob(filepath.Join(reg.projectDir, pattern))
	if err != nil {
		return false
	}
	return len(matches) >= minMatches
}

// hasDirectory checks if a directory exists
func (reg *Registry) hasDirectory(pattern string) bool {
	path := filepath.Join(reg.projectDir, pattern)
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

// hasMatchingContent checks if pattern exists in file contents
func (reg *Registry) hasMatchingContent(pattern string) bool {
	found := false
	_ = filepath.Walk(reg.projectDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if info.IsDir() {
			return nil
		}
		// Skip non-text files
		if !isTextFile(path) {
			return nil
		}
		content, err := os.ReadFile(path)
		if err != nil {
			return nil
		}
		if strings.Contains(string(content), pattern) {
			found = true
			return filepath.SkipDir
		}
		return nil
	})
	return found
}

func loadDefaultConfigs(reg *Registry) {
	reg.types["cli-tool"] = &ProjectTypeConfig{
		ID:          "cli-tool",
		Name:        "CLI Tool",
		Description: "Command-line interface tool",
		Detection: DetectionRules{
			RequiredPatterns: []Pattern{
				{Type: "file", Pattern: "go.mod", MinMatches: 1},
				{Type: "directory", Pattern: "cmd/", MinMatches: 1},
			},
			SuggestingPatterns: []Pattern{{Type: "file", Pattern: "internal/cli/*.go", MinMatches: 1, Weight: 0.8}},
		},
		BaseContext:     BaseContext{Docs: []DocTemplate{{Name: "CONTEXT.md", Template: "base-context", Required: true}, {Name: "CLI_ARCHITECTURE.md", Template: "cli-architecture", Required: true}, {Name: "SECURITY.md", Template: "base-security", Required: true}, {Name: "QUALITY_GATES.md", Template: "base-quality", Required: true}, {Name: "ACTION_ITEMS.md", Template: "base-action-items", Required: true}}},
		SuggestOverlays: []string{"testing-overlay", "security-overlay", "ci-cd-overlay", "documentation-overlay"},
	}

	reg.types["web-fullstack"] = &ProjectTypeConfig{
		ID:          "web-fullstack",
		Name:        "Web Full-Stack",
		Description: "Web application with frontend and backend",
		Detection: DetectionRules{
			RequiredPatterns: []Pattern{{Type: "directory", Pattern: "frontend/", MinMatches: 1}, {Type: "directory", Pattern: "backend/", MinMatches: 1}},
		},
		BaseContext:     BaseContext{Docs: []DocTemplate{{Name: "CONTEXT.md", Template: "base-context", Required: true}, {Name: "BACKEND.md", Template: "base-backend", Required: true}, {Name: "FRONTEND.md", Template: "base-frontend", Required: true}, {Name: "SECURITY.md", Template: "base-security", Required: true}, {Name: "QUALITY_GATES.md", Template: "base-quality", Required: true}, {Name: "ACTION_ITEMS.md", Template: "base-action-items", Required: true}}},
		SuggestOverlays: []string{"testing-overlay", "security-overlay", "ci-cd-overlay", "documentation-overlay"},
	}

	reg.types["library"] = &ProjectTypeConfig{
		ID:              "library",
		Name:            "Library / SDK",
		Description:     "Reusable library or SDK",
		Detection:       DetectionRules{RequiredPatterns: []Pattern{{Type: "file", Pattern: "go.mod", MinMatches: 1}}, SuggestingPatterns: []Pattern{{Type: "file", Pattern: "README.md", MinMatches: 1, Weight: 0.4}}},
		BaseContext:     BaseContext{Docs: []DocTemplate{{Name: "CONTEXT.md", Template: "base-context", Required: true}, {Name: "PUBLIC_API.md", Template: "library-public-api", Required: true}, {Name: "VERSIONING.md", Template: "library-versioning", Required: true}, {Name: "BREAKING_CHANGES.md", Template: "library-breaking", Required: true}, {Name: "SECURITY.md", Template: "base-security", Required: true}, {Name: "QUALITY_GATES.md", Template: "base-quality", Required: true}, {Name: "ACTION_ITEMS.md", Template: "base-action-items", Required: true}}},
		SuggestOverlays: []string{"testing-overlay", "documentation-overlay"},
	}

	reg.overlays["testing-overlay"] = &OverlayConfig{ID: "testing-overlay", Name: "Testing Documentation", Adds: OverlayAdds{Docs: []DocRef{{Name: "TESTING.md", Template: "testing-overview", Required: false}}}}
	reg.overlays["security-overlay"] = &OverlayConfig{ID: "security-overlay", Name: "Security Enhanced"}
	reg.overlays["ci-cd-overlay"] = &OverlayConfig{ID: "ci-cd-overlay", Name: "CI/CD Configuration"}
	reg.overlays["documentation-overlay"] = &OverlayConfig{ID: "documentation-overlay", Name: "Documentation Enhancement"}
}

// isTextFile checks if a file is likely a text file
func isTextFile(path string) bool {
	ext := strings.ToLower(filepath.Ext(path))
	textExtensions := map[string]bool{
		".go":   true,
		".py":   true,
		".ts":   true,
		".js":   true,
		".json": true,
		".md":   true,
		".yaml": true,
		".yml":  true,
		".toml": true,
		".txt":  true,
		".c":    true,
		".h":    true,
		".rs":   true,
		".java": true,
		".sh":   true,
		".bash": true,
	}
	return textExtensions[ext]
}

// collectEvidence collects evidence for why a type matched
func (reg *Registry) collectEvidence(config *ProjectTypeConfig) []string {
	var evidence []string

	// Required patterns evidence
	for _, pattern := range config.Detection.RequiredPatterns {
		if reg.matchesPattern(pattern) {
			evidence = append(evidence, fmt.Sprintf("Required: %s (%s)", pattern.Type, pattern.Pattern))
		}
	}

	// Suggesting patterns evidence
	for _, pattern := range config.Detection.SuggestingPatterns {
		if reg.matchesPattern(pattern) {
			evidence = append(evidence, fmt.Sprintf("Suggested: %s (%s)", pattern.Type, pattern.Pattern))
		}
	}

	return evidence
}

// GetConfig returns a project type configuration by ID
func (reg *Registry) GetConfig(typeID string) (*ProjectTypeConfig, bool) {
	config, ok := reg.types[typeID]
	return config, ok
}

// GetAllConfigs returns all project type configurations
func (reg *Registry) GetAllConfigs() []*ProjectTypeConfig {
	configs := make([]*ProjectTypeConfig, 0, len(reg.types))
	for _, config := range reg.types {
		configs = append(configs, config)
	}
	return configs
}

// GetOverlayConfig returns an overlay configuration by ID
func (reg *Registry) GetOverlayConfig(overlayID string) (*OverlayConfig, bool) {
	config, ok := reg.overlays[overlayID]
	return config, ok
}

// GetAllOverlayConfigs returns all overlay configurations
func (reg *Registry) GetAllOverlayConfigs() []*OverlayConfig {
	configs := make([]*OverlayConfig, 0, len(reg.overlays))
	for _, config := range reg.overlays {
		configs = append(configs, config)
	}
	return configs
}
