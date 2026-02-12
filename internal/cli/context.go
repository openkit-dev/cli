package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/fatih/color"
	"github.com/openkit-devtools/openkit/internal/detection"
	"github.com/spf13/cobra"
)

var (
	contextAutoYes  bool
	contextType     string
	contextOverlays []string
)

var contextCmd = &cobra.Command{
	Use:   "context",
	Short: "Analyze codebase and generate context documentation",
	Long: `Analyze the current codebase to detect project type and generate
appropriate context documentation based on detected project characteristics.

This command:
- Detects project type from file structure and patterns
- Shows detection results with evidence
- Generates context docs based on project type templates
- Suggests and applies optional overlays for additional capabilities`,
	Run: func(cmd *cobra.Command, args []string) {
		runContext()
	},
}

func init() {
	rootCmd.AddCommand(contextCmd)
	contextCmd.Flags().BoolVar(&contextAutoYes, "yes", false, "Skip prompts and use detected/default options")
	contextCmd.Flags().StringVar(&contextType, "type", "", "Force project type (e.g., cli-tool, web-fullstack, library)")
	contextCmd.Flags().StringSliceVar(&contextOverlays, "overlays", nil, "Comma-separated overlays to apply")
}

func runContext() {
	cyan := color.New(color.FgCyan, color.Bold)
	green := color.New(color.FgGreen, color.Bold)
	yellow := color.New(color.FgYellow)
	white := color.New(color.FgWhite)

	// Get current directory
	projectDir, err := os.Getwd()
	if err != nil {
		exitWithError(fmt.Sprintf("Failed to get current directory: %v", err))
	}

	cyan.Println("\nDetecting project type...")

	// Load project type registry
	registry, err := detection.LoadRegistry(projectDir)
	if err != nil {
		exitWithError(fmt.Sprintf("Failed to load project type registry: %v", err))
	}

	// Detect project type
	result, err := registry.Detect()
	if err != nil {
		exitWithError(fmt.Sprintf("Detection failed: %v", err))
	}

	forcedType := strings.TrimSpace(contextType)
	if forcedType != "" {
		result.ProjectType = forcedType
		result.Evidence = append(result.Evidence, "Project type forced via --type flag")
	}

	// Handle unknown project type
	if result.ProjectType == "unknown" {
		yellow.Println("\nCould not automatically detect project type")
		var selectedType string
		if forcedType != "" {
			selectedType = forcedType
		} else {
			fmt.Println("\nPlease select a project type from the available options:")
			selectedType = selectProjectType(registry)
		}

		// Get config for selected type
		config, ok := registry.GetConfig(selectedType)
		if !ok {
			exitWithError("Failed to get project type configuration")
		}

		// Generate context with manual selection
		if err := generateContext(projectDir, registry, config, nil); err != nil {
			exitWithError(fmt.Sprintf("Failed to generate context: %v", err))
		}

		green.Println("\nContext documentation generated successfully!")
		return
	}

	// Get detected project type config
	config, ok := registry.GetConfig(result.ProjectType)
	if !ok {
		exitWithError("Failed to get project type configuration")
	}

	// Show detection results
	cyan.Printf("\nDetected: %s", config.Name)
	if result.Confidence > 0 {
		white.Printf(" (confidence: %.0f%%)\n", result.Confidence)
	} else {
		fmt.Println()
	}

	if config.Description != "" {
		fmt.Printf("  %s\n", config.Description)
	}

	// Show evidence
	if len(result.Evidence) > 0 {
		fmt.Println("\nEvidence:")
		for _, e := range result.Evidence {
			fmt.Printf("  • %s\n", e)
		}
	}

	// Show conflicts if any
	if len(result.Conflicts) > 0 {
		yellow.Println("\nConflicts detected:")
		for _, c := range result.Conflicts {
			fmt.Printf("  • %s\n", c)
		}
	}

	// Show suggested overlays
	suggestedOverlays := result.SuggestedOverlays
	if len(suggestedOverlays) > 0 {
		fmt.Println("\nSuggested overlays:")
		for _, overlayID := range suggestedOverlays {
			overlay, ok := registry.GetOverlayConfig(overlayID)
			if !ok {
				continue
			}
			fmt.Printf("  • %s", overlay.Name)
			if overlay.Description != "" {
				fmt.Printf(": %s", overlay.Description)
			}
			fmt.Println()
		}
	}

	// Ask for confirmation
	if !contextAutoYes {
		fmt.Println()
		if !confirm(fmt.Sprintf("Generate context for %s?", config.Name)) {
			fmt.Println("Context generation cancelled.")
			return
		}
	}

	// Ask for overlay selection
	selectedOverlays := resolveOverlays(registry, suggestedOverlays)

	// Generate context documentation
	if err := generateContext(projectDir, registry, config, selectedOverlays); err != nil {
		exitWithError(fmt.Sprintf("Failed to generate context: %v", err))
	}

	green.Println("\n✓ Context documentation generated successfully!")
	fmt.Println("\nGenerated files:")
	docsDir := filepath.Join(projectDir, "docs")
	for _, doc := range config.BaseContext.Docs {
		docPath := filepath.Join(docsDir, doc.Name)
		fmt.Printf("  • %s\n", docPath)
	}

	if len(selectedOverlays) > 0 {
		fmt.Println("\nApplied overlays:")
		for _, overlayID := range selectedOverlays {
			overlay, ok := registry.GetOverlayConfig(overlayID)
			if !ok {
				continue
			}
			for _, doc := range overlay.Adds.Docs {
				docPath := filepath.Join(docsDir, doc.Name)
				fmt.Printf("  • %s (from %s)\n", docPath, overlay.Name)
			}
		}
	}

	fmt.Printf("\nNext: Run your AI agent to analyze the context\n")
}

// selectProjectType prompts user to select a project type
func selectProjectType(registry *detection.Registry) string {
	configs := registry.GetAllConfigs()
	sort.Slice(configs, func(i, j int) bool { return configs[i].ID < configs[j].ID })

	fmt.Println()
	for i, config := range configs {
		fmt.Printf("  [%d] %s", i+1, config.Name)
		if config.Description != "" {
			fmt.Printf(" - %s", config.Description)
		}
		fmt.Println()
	}

	for {
		fmt.Printf("\nEnter selection (1-%d): ", len(configs))
		var choice int
		if _, err := fmt.Scanf("%d\n", &choice); err != nil {
			printError("Invalid input")
			continue
		}

		if choice < 1 || choice > len(configs) {
			printError(fmt.Sprintf("Please enter a number between 1 and %d", len(configs)))
			continue
		}

		return configs[choice-1].ID
	}
}

func resolveOverlays(registry *detection.Registry, suggested []string) []string {
	if len(contextOverlays) > 0 {
		var out []string
		for _, id := range contextOverlays {
			id = strings.TrimSpace(id)
			if id == "" {
				continue
			}
			if _, ok := registry.GetOverlayConfig(id); ok {
				out = append(out, id)
			}
		}
		return out
	}

	if contextAutoYes {
		return suggested
	}

	return selectOverlays(registry, suggested)
}

// selectOverlays prompts user to select overlays
func selectOverlays(registry *detection.Registry, suggested []string) []string {
	if len(suggested) == 0 {
		return nil
	}

	cyan := color.New(color.FgCyan)
	cyan.Println("\nSelect overlays to apply (press Enter for all suggested):")

	for i, overlayID := range suggested {
		overlay, ok := registry.GetOverlayConfig(overlayID)
		if !ok {
			continue
		}
		fmt.Printf("  [%d] %s", i+1, overlay.Name)
		if overlay.Description != "" {
			fmt.Printf(": %s", overlay.Description)
		}
		fmt.Println()
	}
	fmt.Printf("  [0] All suggested overlays\n")

	for {
		fmt.Printf("\nEnter selection (0-%d): ", len(suggested))
		var choice int
		if _, err := fmt.Scanf("%d\n", &choice); err != nil {
			// Empty input means select all
			fmt.Println()
			return suggested
		}

		if choice == 0 {
			return suggested
		}

		if choice < 1 || choice > len(suggested) {
			printError(fmt.Sprintf("Please enter a number between 0 and %d", len(suggested)))
			continue
		}

		return []string{suggested[choice-1]}
	}
}

// generateContext generates context documentation based on project type and overlays
func generateContext(projectDir string, registry *detection.Registry, config *detection.ProjectTypeConfig, selectedOverlays []string) error {
	docsDir := filepath.Join(projectDir, "docs")

	// Create docs directory if it doesn't exist
	if err := os.MkdirAll(docsDir, 0755); err != nil {
		return fmt.Errorf("failed to create docs directory: %w", err)
	}

	// Generate base context docs
	for _, docTemplate := range config.BaseContext.Docs {
		if err := generateDocFromTemplate(docsDir, &docTemplate); err != nil {
			return fmt.Errorf("failed to generate %s: %w", docTemplate.Name, err)
		}
	}

	// Apply overlays
	for _, overlayID := range selectedOverlays {
		overlay, ok := registry.GetOverlayConfig(overlayID)
		if !ok {
			continue
		}

		// Add docs from overlay
		for _, doc := range overlay.Adds.Docs {
			if err := generateDocFromTemplate(docsDir, &doc); err != nil {
				return fmt.Errorf("failed to generate overlay doc %s: %w", doc.Name, err)
			}
		}

		// Modify existing docs if specified
		if hasOverlayModifications(overlay.Modifies) {
			if err := applyModifications(docsDir, overlay.Modifies); err != nil {
				return fmt.Errorf("failed to apply modifications: %w", err)
			}
		}
	}

	// Generate backward-compatible files for web-fullstack
	if config.ID == "web-fullstack" {
		if err := generateWebFullstackCompatFiles(docsDir); err != nil {
			return fmt.Errorf("failed to generate web-fullstack compat files: %w", err)
		}
	}

	return nil
}

// generateDocFromTemplate generates a document from a template reference
func generateDocFromTemplate(docsDir string, docTemplate interface{}) error {
	var name, template string
	var required bool

	switch t := docTemplate.(type) {
	case *detection.DocTemplate:
		name = t.Name
		template = t.Template
		required = t.Required
	case *detection.DocRef:
		name = t.Name
		template = t.Template
		required = t.Required
	default:
		return fmt.Errorf("invalid template type")
	}

	// Map template names to template content
	templateContent, ok := getTemplateContent(template)
	if !ok {
		return fmt.Errorf("template not found: %s", template)
	}

	docPath := filepath.Join(docsDir, name)

	// Check if file exists and is required
	if _, err := os.Stat(docPath); err == nil {
		if required {
			// File exists, skip
			return nil
		}
	}

	return os.WriteFile(docPath, []byte(templateContent), 0644)
}

// applyModifications applies overlay modifications to existing docs
func applyModifications(docsDir string, modifications detection.OverlayModifies) error {
	// Helper function to modify a doc file
	modifyDoc := func(filename string, mods *detection.DocModifications) error {
		if mods == nil {
			return nil
		}

		docPath := filepath.Join(docsDir, filename)
		content, err := os.ReadFile(docPath)
		if err != nil {
			if os.IsNotExist(err) {
				return nil // File doesn't exist, skip
			}
			return err
		}

		contentStr := string(content)

		// Add sections
		if mods.AddSection != nil {
			for sectionName, sectionContent := range mods.AddSection {
				contentStr += fmt.Sprintf("\n## %s\n\n%s\n", sectionName, sectionContent)
			}
		}

		// Add items (append to existing lists)
		if len(mods.AddItems) > 0 {
			for _, item := range mods.AddItems {
				contentStr += fmt.Sprintf("\n- %s", item)
			}
		}

		return os.WriteFile(docPath, []byte(contentStr), 0644)
	}

	// Apply QUALITY_GATES modifications
	if modifications.QUALITY_GATES != nil {
		if err := modifyDoc("QUALITY_GATES.md", modifications.QUALITY_GATES); err != nil {
			return err
		}
	}

	// Apply ACTION_ITEMS modifications
	if modifications.ACTION_ITEMS != nil {
		if err := modifyDoc("ACTION_ITEMS.md", modifications.ACTION_ITEMS); err != nil {
			return err
		}
	}

	return nil
}

// generateWebFullstackCompatFiles generates backward-compatible files for web-fullstack projects
func generateWebFullstackCompatFiles(docsDir string) error {
	// Generate BACKEND.md if it doesn't exist
	backendPath := filepath.Join(docsDir, "BACKEND.md")
	if _, err := os.Stat(backendPath); os.IsNotExist(err) {
		templateContent, ok := getTemplateContent("base-backend")
		if !ok {
			return fmt.Errorf("base-backend template not found")
		}
		if err := os.WriteFile(backendPath, []byte(templateContent), 0644); err != nil {
			return err
		}
	}

	// Generate FRONTEND.md if it doesn't exist
	frontendPath := filepath.Join(docsDir, "FRONTEND.md")
	if _, err := os.Stat(frontendPath); os.IsNotExist(err) {
		templateContent, ok := getTemplateContent("base-frontend")
		if !ok {
			return fmt.Errorf("base-frontend template not found")
		}
		if err := os.WriteFile(frontendPath, []byte(templateContent), 0644); err != nil {
			return err
		}
	}

	return nil
}

// confirm prompts user for yes/no confirmation
func confirm(prompt string) bool {
	fmt.Printf("\n%s [y/N]: ", prompt)
	var response string
	if _, err := fmt.Scanln(&response); err != nil {
		return false
	}
	response = strings.ToLower(strings.TrimSpace(response))

	return response == "y" || response == "yes"
}

// printError prints an error message
func printError(msg string) {
	red := color.New(color.FgRed, color.Bold)
	red.Printf("Error: %s\n", msg)
}

// getTemplateContent returns template content by template name
// This maps template names from project type configs to actual template files
func getTemplateContent(templateName string) (string, bool) {
	// Map of template names to template content
	templateMap := map[string]string{
		// Base templates
		"base-context":      baseContextTemplate,
		"base-security":     baseSecurityTemplate,
		"base-quality":      baseQualityTemplate,
		"base-action-items": baseActionItemsTemplate,

		// Project type specific templates
		"base-backend":     baseBackendTemplate,
		"base-frontend":    baseFrontendTemplate,
		"cli-architecture": cliArchitectureTemplate,

		// Library templates
		"library-public-api": libraryPublicApiTemplate,
		"library-versioning": libraryVersioningTemplate,
		"library-breaking":   libraryBreakingTemplate,

		// Overlay templates
		"testing-overview": testingOverviewTemplate,
	}

	content, ok := templateMap[templateName]
	return content, ok
}

func hasOverlayModifications(mod detection.OverlayModifies) bool {
	return mod.QUALITY_GATES != nil || mod.ACTION_ITEMS != nil
}

// Template contents
const (
	baseContextTemplate = "# CONTEXT\n\n" +
		"**Created**: [YYYY-MM-DD]\n" +
		"**Scope**: [Full | Backend | Frontend]\n\n" +
		"## Executive Summary (10 bullets)\n\n" +
		"- [Key fact 1]\n" +
		"- [Key fact 2]\n\n" +
		"## Repository Map\n\n" +
		"| Area | Path(s) | Notes |\n" +
		"|---|---|---|\n" +
		"| CLI | | |\n" +
		"| Backend | | |\n" +
		"| Frontend | | |\n\n" +
		"## Key Flows\n\n" +
		"1. [Flow name]: [entry] -> [core] -> [outputs]\n\n" +
		"## Evidence\n\n" +
		"- `path/to/file.ext`: short snippet / reason\n\n" +
		"## Terminology\n\n" +
		"> For standard terminology definitions, see `docs/GLOSSARY.md`\n\n" +
		"| Term | Definition (project-specific) |\n" +
		"|------|-------------------------------|\n" +
		"| [Term 1] | [Definition specific to this project] |\n" +
		"| [Term 2] | [Definition specific to this project] |\n"

	baseSecurityTemplate = "# SECURITY\n\n" +
		"Last updated: [YYYY-MM-DD]\n\n" +
		"## Security Principles\n\n" +
		"1. **Never trust user input** - Validate all inputs at the boundary\n" +
		"2. **Principle of least privilege** - Minimal access required\n" +
		"3. **Defense in depth** - Multiple layers of security controls\n" +
		"4. **Secure by default** - Security features enabled by default\n\n" +
		"## Security Checklist\n\n" +
		"- [ ] Input validation on all user inputs\n" +
		"- [ ] Output encoding to prevent XSS\n" +
		"- [ ] Parameterized queries to prevent SQL injection\n" +
		"- [ ] Authentication and authorization on all protected routes\n" +
		"- [ ] HTTPS/TLS for all communications\n" +
		"- [ ] Secrets managed securely (no hardcoded credentials)\n" +
		"- [ ] Dependency vulnerability scanning\n" +
		"- [ ] Security headers configured\n" +
		"- [ ] Rate limiting on API endpoints\n" +
		"- [ ] Logging and monitoring for security events\n\n" +
		"## Security Resources\n\n" +
		"- OWASP Top 10: https://owasp.org/www-project-top-ten/\n" +
		"- CWE/SANS Top 25: https://cwe.mitre.org/top25/\n\n" +
		"## Reporting Security Issues\n\n" +
		"If you discover a security vulnerability, please report it privately to [security@yourdomain.com].\n"

	baseQualityTemplate = "# QUALITY GATES\n\n" +
		"Quality gates that must be passed before merging changes.\n\n" +
		"## Code Quality\n\n" +
		"- [ ] Code follows project style guide\n" +
		"- [ ] No linting errors or warnings\n" +
		"- [ ] TypeScript/Python types properly defined\n" +
		"- [ ] Code complexity within acceptable limits\n" +
		"- [ ] No TODO or FIXME comments in production code\n\n" +
		"## Testing\n\n" +
		"- [ ] Unit tests pass (minimum 80% coverage)\n" +
		"- [ ] Integration tests pass\n" +
		"- [ ] New features have test coverage\n" +
		"- [ ] Critical paths have end-to-end tests\n\n" +
		"## Documentation\n\n" +
		"- [ ] API documentation updated\n" +
		"- [ ] README updated (if applicable)\n" +
		"- [ ] CHANGELOG entry added\n" +
		"- [ ] Code comments for complex logic\n\n" +
		"## Security\n\n" +
		"- [ ] No hardcoded secrets or credentials\n" +
		"- [ ] No vulnerabilities in dependencies\n" +
		"- [ ] Input validation on all user inputs\n" +
		"- [ ] Proper error handling (no stack traces to users)\n\n" +
		"## Performance\n\n" +
		"- [ ] No significant performance regressions\n" +
		"- [ ] Database queries optimized (no N+1)\n" +
		"- [ ] Assets properly optimized and cached\n\n" +
		"## Deployment\n\n" +
		"- [ ] Environment variables documented\n" +
		"- [ ] Migration scripts tested\n" +
		"- [ ] Rollback plan documented\n"

	baseActionItemsTemplate = "# ACTION ITEMS\n\n" +
		"Track cross-cutting action items and decisions that impact multiple areas.\n\n" +
		"## Current Action Items\n\n" +
		"| ID | Description | Owner | Status | Priority |\n" +
		"|----|-------------|-------|--------|----------|\n" +
		"| [ID] | [Description] | [Owner] | [Open/Done] | [High/Medium/Low] |\n\n" +
		"## Completed Action Items\n\n" +
		"| ID | Description | Completed Date |\n" +
		"|----|-------------|----------------|\n" +
		"| [ID] | [Description] | [YYYY-MM-DD] |\n"

	baseBackendTemplate = "# BACKEND\n\n" +
		"Backend architecture and implementation details.\n\n" +
		"## Technology Stack\n\n" +
		"| Component | Technology | Version |\n" +
		"|-----------|-----------|---------|\n" +
		"| Framework | | |\n" +
		"| Runtime | | |\n" +
		"| Database | | |\n" +
		"| ORM | | |\n\n" +
		"## API Architecture\n\n" +
		"### REST Endpoints\n\n" +
		"| Method | Path | Description | Auth |\n" +
		"|--------|------|-------------|------|\n" +
		"| GET | /api/endpoint | Description | Required |\n\n" +
		"### Authentication/Authorization\n\n" +
		"- [ ] Authentication method: [JWT/Session/OAuth]\n" +
		"- [ ] Authorization: [Role-based/Permission-based]\n\n" +
		"## Data Models\n\n" +
		"### [Model Name]\n\n" +
		"| Field | Type | Description |\n" +
		"|-------|------|-------------|\n" +
		"| id | UUID | Primary key |\n\n" +
		"## Key Services\n\n" +
		"| Service | Responsibility |\n" +
		"|---------|---------------|\n" +
		"| [Service] | [Description] |\n\n" +
		"## Database Schema\n\n" +
		"See `docs/database/SCHEMA.md` for full schema details.\n\n" +
		"## Environment Variables\n\n" +
		"| Variable | Description | Required |\n" +
		"|----------|-------------|----------|\n" +
		"| DATABASE_URL | PostgreSQL connection string | Yes |\n\n" +
		"## Common Patterns\n\n" +
		"1. **Error Handling**: Centralized error handling middleware\n" +
		"2. **Validation**: Input validation at API boundary\n" +
		"3. **Logging**: Structured logging with correlation IDs\n"

	baseFrontendTemplate = "# FRONTEND\n\n" +
		"Frontend architecture and implementation details.\n\n" +
		"## Technology Stack\n\n" +
		"| Component | Technology | Version |\n" +
		"|-----------|-----------|---------|\n" +
		"| Framework | | |\n" +
		"| Build Tool | | |\n" +
		"| Styling | | |\n" +
		"| State Management | | |\n\n" +
		"## Project Structure\n\n" +
		"```\n" +
		"frontend/\n" +
		"├── src/\n" +
		"│   ├── components/     # Reusable UI components\n" +
		"│   ├── pages/          # Page components\n" +
		"│   ├── hooks/          # Custom React hooks\n" +
		"│   ├── lib/            # Utility functions\n" +
		"│   └── styles/         # Global styles\n" +
		"```\n\n" +
		"## Key Features\n\n" +
		"| Feature | Description | Status |\n" +
		"|---------|-------------|--------|\n" +
		"| [Feature] | [Description] | [Done/In Progress/Planned] |\n\n" +
		"## Component Library\n\n" +
		"| Component | Props | Usage |\n" +
		"|-----------|-------|-------|\n" +
		"| [Component] | [Prop list] | [Where used] |\n\n" +
		"## State Management\n\n" +
		"| Store/Context | Purpose |\n" +
		"|---------------|---------|\n" +
		"| [Store] | [Description] |\n\n" +
		"## API Integration\n\n" +
		"| API Call | Method | Endpoint | Purpose |\n" +
		"|----------|--------|----------|---------|\n" +
		"| [Function] | GET/POST | /api/... | [Description] |\n\n" +
		"## Styling Guidelines\n\n" +
		"- Use Tailwind utility classes\n" +
		"- Follow component-based architecture\n" +
		"- Mobile-first responsive design\n" +
		"- Accessible color contrast ratios\n\n" +
		"## Environment Variables\n\n" +
		"| Variable | Description | Required |\n" +
		"|----------|-------------|----------|\n" +
		"| VITE_API_URL | Backend API URL | Yes |\n"

	cliArchitectureTemplate = "# CLI ARCHITECTURE\n\n" +
		"CLI tool architecture and command structure.\n\n" +
		"## Command Structure\n\n" +
		"```\n" +
		"myapp\n" +
		"├── command1\n" +
		"│   └── subcommand\n" +
		"├── command2\n" +
		"└── --global-flag\n" +
		"```\n\n" +
		"## Root Commands\n\n" +
		"| Command | Short | Description |\n" +
		"|---------|-------|-------------|\n" +
		"| init | | Initialize new project |\n" +
		"| build | | Build the application |\n\n" +
		"## Key Components\n\n" +
		"| Component | Responsibility |\n" +
		"|-----------|---------------|\n" +
		"| Command Runner | Parse and execute commands |\n" +
		"| Configuration | Load and validate config |\n" +
		"| Output Formatter | Format output (JSON/text/table) |\n\n" +
		"## Configuration\n\n" +
		"| File | Format | Purpose |\n" +
		"|------|--------|---------|\n" +
		"| config.yaml | YAML | User configuration |\n\n" +
		"## Environment Variables\n\n" +
		"| Variable | Description | Default |\n" +
		"|----------|-------------|---------|\n" +
		"| MYAPP_CONFIG | Path to config file | ~/.config/myapp/config.yaml |\n" +
		"| MYAPP_DEBUG | Enable debug logging | false |\n\n" +
		"## Error Handling\n\n" +
		"- Use structured error types\n" +
		"- Provide helpful error messages\n" +
		"- Suggest fixes when possible\n\n" +
		"## Output Formats\n\n" +
		"- **text**: Human-readable output (default)\n" +
		"- **json**: Machine-readable JSON\n" +
		"- **table**: Tabular data\n"

	libraryPublicApiTemplate = "# PUBLIC API\n\n" +
		"Public API documentation for library consumers.\n\n" +
		"## Overview\n\n" +
		"[Library name] provides [brief description].\n\n" +
		"## Installation\n\n" +
		"```\n" +
		"# Go\n" +
		"go get github.com/example/library\n\n" +
		"# Node\n" +
		"npm install @example/library\n" +
		"```\n\n" +
		"## Quick Start\n\n" +
		"```\n" +
		"go\n" +
		"package main\n\n" +
		"import \"github.com/example/library\"\n" +
		"\n" +
		"func main() {\n" +
		"    // Your code here\n" +
		"}\n" +
		"```\n\n" +
		"## Core API\n\n" +
		"### FunctionName\n\n" +
		"Description of what the function does.\n\n" +
		"```\n" +
		"go\n" +
		"func FunctionName(param Type) ReturnType\n" +
		"```\n\n" +
		"**Parameters:**\n\n" +
		"| Param | Type | Description |\n" +
		"|-------|------|-------------|\n" +
		"| param | Type | Description |\n\n" +
		"**Returns:**\n\n" +
		"| Type | Description |\n" +
		"|------|-------------|\n" +
		"| ReturnType | Description |\n\n" +
		"**Example:**\n\n" +
		"```\n" +
		"go\n" +
		"result := library.FunctionName(value)\n" +
		"```\n\n" +
		"## Types\n\n" +
		"### TypeName\n\n" +
		"Description of the type.\n\n" +
		"```\n" +
		"go\n" +
		"type TypeName struct {\n" +
		"    Field Type // Description\n" +
		"}\n" +
		"```\n\n" +
		"## Constants\n\n" +
		"| Constant | Value | Description |\n" +
		"|----------|-------|-------------|\n" +
		"| CONSTANT | value | Description |\n\n" +
		"## Errors\n\n" +
		"| Error | Description |\n" +
		"|-------|-------------|\n" +
		"| ErrNotFound | Item not found |\n\n" +
		"## Best Practices\n\n" +
		"1. [Best practice 1]\n" +
		"2. [Best practice 2]\n"

	libraryVersioningTemplate = "# VERSIONING\n\n" +
		"Library versioning and release guidelines.\n\n" +
		"## Version Scheme\n\n" +
		"This project follows [Semantic Versioning 2.0.0](https://semver.org/).\n\n" +
		"- **MAJOR**: Incompatible API changes\n" +
		"- **MINOR**: Backwards-compatible functionality additions\n" +
		"- **PATCH**: Backwards-compatible bug fixes\n\n" +
		"## Current Version\n\n" +
		"Latest release: `v1.0.0`\n\n" +
		"## Release Process\n\n" +
		"1. Update version in `go.mod`/`package.json`\n" +
		"2. Create release notes in `CHANGELOG.md`\n" +
		"3. Tag commit with `git tag v1.2.3`\n" +
		"4. Push tags with `git push --tags`\n" +
		"5. Create GitHub release\n\n" +
		"## Breaking Changes\n\n" +
		"Breaking changes are announced in the `BREAKING_CHANGES.md` file and\n" +
		"in the release notes. Migration guides are provided.\n\n" +
		"## Deprecation Policy\n\n" +
		"- Features are deprecated for at least one minor version before removal\n" +
		"- Deprecation warnings are logged at runtime\n" +
		"- Deprecation notices are added to documentation\n\n" +
		"## Backwards Compatibility\n\n" +
		"We maintain backwards compatibility for:\n" +
		"- Public API functions\n" +
		"- Struct fields exported by the library\n" +
		"- Configuration file formats (with migrations)\n"

	libraryBreakingTemplate = "# BREAKING CHANGES\n\n" +
		"Track breaking changes and migration guides.\n\n" +
		"## Version 2.0.0 (Planned)\n\n" +
		"### Breaking Change: [Name]\n\n" +
		"**What changed:**\n" +
		"- Description of the change\n\n" +
		"**Why changed:**\n" +
		"- Reason for the change\n\n" +
		"**Migration:**\n" +
		"```\n" +
		"go\n" +
		"// Old code\n" +
		"oldCode()\n\n" +
		"// New code\n" +
		"newCode()\n" +
		"```\n\n" +
		"### Breaking Change: [Another Change]\n\n" +
		"[Details...]\n\n" +
		"---\n\n" +
		"## Version 1.5.0 (Released)\n\n" +
		"### Breaking Change: [Name]\n\n" +
		"[Details...]\n\n" +
		"## Previous Breaking Changes\n\n" +
		"See `CHANGELOG.md` for full history of changes.\n"

	testingOverviewTemplate = "# TESTING\n\n" +
		"Testing strategy and guidelines.\n\n" +
		"## Testing Pyramid\n\n" +
		"```\n" +
		"         /\\        \n" +
		"        /  \\       \n" +
		"       /E2E \\      (10%)\n" +
		"      /------\\    \n" +
		"     /Integration\\ (20%)\n" +
		"    /------------\\\n" +
		"   /   Unit Tests \\  (70%)\n" +
		"  /----------------\\\n" +
		"```\n\n" +
		"## Unit Tests\n\n" +
		"- Test individual functions and methods in isolation\n" +
		"- Mock external dependencies\n" +
		"- Fast execution (milliseconds)\n\n" +
		"### Running Unit Tests\n\n" +
		"```\n" +
		"# Go\n" +
		"go test ./...\n\n" +
		"# Node\n" +
		"npm test\n" +
		"```\n\n" +
		"### Coverage Goals\n\n" +
		"- Core business logic: 90%+\n" +
		"- API endpoints: 85%+\n" +
		"- Utilities: 95%+\n\n" +
		"## Integration Tests\n\n" +
		"- Test interactions between components\n" +
		"- Use real database (test instance)\n" +
		"- Slower execution (seconds)\n\n" +
		"### Running Integration Tests\n\n" +
		"```\n" +
		"# With test database\n" +
		"go test -tags=integration ./...\n\n" +
		"# Node\n" +
		"npm run test:integration\n" +
		"```\n\n" +
		"## End-to-End Tests\n\n" +
		"- Test complete user workflows\n" +
		"- Run in production-like environment\n" +
		"- Slowest execution (minutes)\n\n" +
		"### Running E2E Tests\n\n" +
		"```\n" +
		"# Playwright\n" +
		"npm run test:e2e\n\n" +
		"# Cypress\n" +
		"npx cypress run\n" +
		"```\n\n" +
		"## Test Organization\n\n" +
		"```\n" +
		"tests/\n" +
		"├── unit/           # Fast, isolated tests\n" +
		"├── integration/    # Component interaction tests\n" +
		"└── e2e/            # Full workflow tests\n" +
		"```\n\n" +
		"## Best Practices\n\n" +
		"1. **Arrange-Act-Assert** pattern for test structure\n" +
		"2. Descriptive test names that explain behavior\n" +
		"3. Test both happy paths and error cases\n" +
		"4. Use fixtures for test data\n" +
		"5. Keep tests independent and deterministic\n\n" +
		"## Common Test Scenarios\n\n" +
		"- [ ] Valid input processing\n" +
		"- [ ] Invalid input handling\n" +
		"- [ ] Error conditions\n" +
		"- [ ] Edge cases and boundaries\n" +
		"- [ ] Authentication/authorization\n" +
		"- [ ] Performance benchmarks\n"
)
