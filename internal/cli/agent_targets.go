package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/openkit-devtools/openkit/internal/agents"
	"github.com/openkit-devtools/openkit/internal/managedstate"
	"github.com/openkit-devtools/openkit/internal/syncer"
	"github.com/openkit-devtools/openkit/internal/targets"
	"github.com/spf13/cobra"
)

var (
	flagDryRun    bool
	flagOverwrite bool
	flagPrune     bool
)

func init() {
	// Shared flags are duplicated per leaf command for now.
	addAgentTarget("opencode")
	addAgentTarget("claude")
	addAgentTarget("gemini")
	addAgentTarget("codex")
	addAgentTarget("cursor")
}

func addAgentTarget(agentID string) {
	ag := agents.Get(agentID)
	if ag == nil {
		return
	}

	cmd := &cobra.Command{
		Use:   ag.ID,
		Short: fmt.Sprintf("Manage OpenKit content for %s", ag.Name),
	}

	syncCmd := &cobra.Command{
		Use:   "sync",
		Short: "Sync OpenKit content into the current project",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runAgentSync(ag.ID)
		},
	}
	upgradeCmd := &cobra.Command{
		Use:   "upgrade",
		Short: "Upgrade OpenKit content in the current project",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Embedded packs: upgrade == sync.
			return runAgentSync(ag.ID)
		},
	}
	doctorCmd := &cobra.Command{
		Use:   "doctor",
		Short: "Check project health and managed state",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runAgentDoctor(ag.ID)
		},
	}

	for _, c := range []*cobra.Command{syncCmd, upgradeCmd} {
		c.Flags().BoolVar(&flagDryRun, "dry-run", false, "Plan changes without writing")
		c.Flags().BoolVar(&flagOverwrite, "overwrite", false, "Overwrite unmanaged or drifted files")
		c.Flags().BoolVar(&flagPrune, "prune", false, "Remove managed files no longer in the target plan (safe)")
	}

	cmd.AddCommand(syncCmd)
	cmd.AddCommand(upgradeCmd)
	cmd.AddCommand(doctorCmd)
	rootCmd.AddCommand(cmd)
}

func runAgentSync(agentID string) error {
	cyan := color.New(color.FgCyan, color.Bold)
	green := color.New(color.FgGreen, color.Bold)
	yellow := color.New(color.FgYellow)

	projectDir, err := os.Getwd()
	if err != nil {
		exitWithError(fmt.Sprintf("Failed to get current directory: %v", err))
	}

	ag := agents.Get(agentID)
	if ag == nil {
		exitWithError(fmt.Sprintf("Unknown agent '%s'", agentID))
	}

	desired, err := targets.BuildEmbeddedDesired(ag, GetVersion())
	if err != nil {
		exitWithError(fmt.Sprintf("Failed to build desired files: %v", err))
	}

	statePath := managedstate.DefaultPath(projectDir)
	var st *managedstate.State
	if _, err := os.Stat(statePath); err == nil {
		loaded, err := managedstate.Load(statePath)
		if err != nil {
			exitWithError(fmt.Sprintf("Invalid managed state: %v", err))
		}
		st = loaded
	}

	cyan.Printf("\nSyncing OpenKit content for %s\n\n", ag.Name)

	opts := syncer.Options{DryRun: flagDryRun, Overwrite: flagOverwrite, Prune: flagPrune}
	res, nextState, err := syncer.Apply(projectDir, ag.ID, desired.PackID, desired.PackVersion, desired.Files, st, opts)
	if err != nil {
		exitWithError(fmt.Sprintf("Sync failed: %v", err))
	}

	counts := map[string]int{
		"created":     len(res.Plan.Create),
		"updated":     len(res.Plan.Update),
		"overwritten": len(res.Plan.Overwrite),
		"deleted":     len(res.Plan.Delete),
		"skipped":     len(res.Plan.Skip),
		"conflicts":   len(res.Plan.Conflicts),
		"orphaned":    len(res.Plan.Orphaned),
	}

	lines := []string{
		fmt.Sprintf("created: %d", counts["created"]),
		fmt.Sprintf("updated: %d", counts["updated"]),
		fmt.Sprintf("overwritten: %d", counts["overwritten"]),
		fmt.Sprintf("deleted: %d", counts["deleted"]),
		fmt.Sprintf("skipped: %d", counts["skipped"]),
		fmt.Sprintf("conflicts: %d", counts["conflicts"]),
		fmt.Sprintf("orphaned: %d", counts["orphaned"]),
	}
	for _, l := range lines {
		fmt.Println(l)
	}
	if res.BackupsDir != "" {
		fmt.Printf("backups: %s\n", res.BackupsDir)
	}

	if flagDryRun {
		yellow.Println("\n(dry-run) No files were written")
		return nil
	}

	if nextState != nil {
		if err := managedstate.Save(statePath, nextState); err != nil {
			exitWithError(fmt.Sprintf("Failed to write %s: %v", filepath.ToSlash(filepath.Join(".openkit", "managed.json")), err))
		}
	}

	green.Println("\nSync completed")
	return nil
}

func runAgentDoctor(agentID string) error {
	cyan := color.New(color.FgCyan, color.Bold)
	green := color.New(color.FgGreen)
	red := color.New(color.FgRed)
	dim := color.New(color.FgHiBlack)

	projectDir, err := os.Getwd()
	if err != nil {
		exitWithError(fmt.Sprintf("Failed to get current directory: %v", err))
	}

	ag := agents.Get(agentID)
	if ag == nil {
		exitWithError(fmt.Sprintf("Unknown agent '%s'", agentID))
	}

	cyan.Printf("\nOpenKit Doctor (%s)\n", ag.ID)
	cyan.Println("====================")
	fmt.Println()

	checks := []struct {
		name  string
		path  string
		isDir bool
	}{
		{name: ".openkit/managed.json", path: managedstate.DefaultPath(projectDir), isDir: false},
	}

	// Add agent-specific entrypoint checks.
	switch ag.ID {
	case "opencode":
		checks = append(checks,
			struct {
				name  string
				path  string
				isDir bool
			}{name: "opencode.json", path: filepath.Join(projectDir, "opencode.json"), isDir: false},
			struct {
				name  string
				path  string
				isDir bool
			}{name: ".opencode/", path: filepath.Join(projectDir, ".opencode"), isDir: true},
		)
	case "claude":
		checks = append(checks,
			struct {
				name  string
				path  string
				isDir bool
			}{name: ".claude/CLAUDE.md", path: filepath.Join(projectDir, ".claude", "CLAUDE.md"), isDir: false},
			struct {
				name  string
				path  string
				isDir bool
			}{name: ".claude/rules/", path: filepath.Join(projectDir, ".claude", "rules"), isDir: true},
		)
	case "gemini":
		checks = append(checks,
			struct {
				name  string
				path  string
				isDir bool
			}{name: "GEMINI.md", path: filepath.Join(projectDir, "GEMINI.md"), isDir: false},
			struct {
				name  string
				path  string
				isDir bool
			}{name: ".gemini/settings.json", path: filepath.Join(projectDir, ".gemini", "settings.json"), isDir: false},
		)
	case "codex":
		checks = append(checks,
			struct {
				name  string
				path  string
				isDir bool
			}{name: "AGENTS.md", path: filepath.Join(projectDir, "AGENTS.md"), isDir: false},
			struct {
				name  string
				path  string
				isDir bool
			}{name: ".codex/rules/", path: filepath.Join(projectDir, ".codex", "rules"), isDir: true},
			struct {
				name  string
				path  string
				isDir bool
			}{name: ".agents/skills/", path: filepath.Join(projectDir, ".agents", "skills"), isDir: true},
		)
	case "cursor":
		checks = append(checks,
			struct {
				name  string
				path  string
				isDir bool
			}{name: ".cursorrules", path: filepath.Join(projectDir, ".cursorrules"), isDir: false},
			struct {
				name  string
				path  string
				isDir bool
			}{name: ".cursor/rules/", path: filepath.Join(projectDir, ".cursor", "rules"), isDir: true},
			struct {
				name  string
				path  string
				isDir bool
			}{name: ".cursor/skills/", path: filepath.Join(projectDir, ".cursor", "skills"), isDir: true},
		)
	}

	for _, c := range checks {
		fi, err := os.Stat(c.path)
		ok := err == nil
		if ok {
			if c.isDir && !fi.IsDir() {
				ok = false
			}
			if !c.isDir && fi.IsDir() {
				ok = false
			}
		}
		if ok {
			green.Print("  [OK] ")
			fmt.Println(c.name)
		} else {
			red.Print("  [--] ")
			fmt.Println(c.name)
		}
	}

	statePath := managedstate.DefaultPath(projectDir)
	st, err := managedstate.Load(statePath)
	if err != nil {
		fmt.Println()
		dim.Println("Managed state not available; run: openkit <agent> sync")
		return nil
	}

	agst := st.Agents[ag.ID]
	if agst == nil {
		fmt.Println()
		dim.Println("No managed entries for this agent; run: openkit <agent> sync")
		return nil
	}

	drifted := 0
	missing := 0
	for rel, entry := range agst.Files {
		if entry == nil {
			continue
		}
		abs, err := syncer.SafeAbsPath(projectDir, rel)
		if err != nil {
			continue
		}
		b, err := os.ReadFile(abs)
		if err != nil {
			if os.IsNotExist(err) {
				missing++
			}
			continue
		}
		sha := managedstate.Sha256HexBytes(b)
		if sha != entry.InstalledSHA256 {
			drifted++
		}
	}

	fmt.Println()
	fmt.Printf("  Managed files: %d\n", len(agst.Files))
	fmt.Printf("  Drifted:       %d\n", drifted)
	fmt.Printf("  Missing:       %d\n", missing)
	if agst.Pack.ID != "" {
		fmt.Printf("  Pack:          %s@%s\n", agst.Pack.ID, agst.Pack.Version)
	}

	if ag.ID == "gemini" {
		fmt.Println()
		dim.Println("Note: Gemini CLI may ignore project commands unless the repo is trusted")
	}

	return nil
}
