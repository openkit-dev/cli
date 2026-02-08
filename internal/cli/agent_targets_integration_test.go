package cli

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/openkit-devtools/openkit/internal/agents"
	"github.com/openkit-devtools/openkit/internal/managedstate"
	"github.com/openkit-devtools/openkit/internal/syncer"
	"github.com/openkit-devtools/openkit/internal/targets"
)

func TestGeminiSync_WritesFilesAndManagedState(t *testing.T) {
	tempDir := t.TempDir()
	oldWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("getwd: %v", err)
	}
	t.Cleanup(func() { _ = os.Chdir(oldWd) })
	if err := os.Chdir(tempDir); err != nil {
		t.Fatalf("chdir: %v", err)
	}

	flagDryRun = false
	flagOverwrite = false
	flagPrune = false

	if err := runAgentSync("gemini"); err != nil {
		t.Fatalf("runAgentSync: %v", err)
	}

	requireFile(t, tempDir, "GEMINI.md")
	requireFile(t, tempDir, filepath.ToSlash(filepath.Join(".gemini", "settings.json")))
	requireFile(t, tempDir, filepath.ToSlash(filepath.Join(".gemini", "commands", "openkit", "specify.toml")))
	requireFile(t, tempDir, filepath.ToSlash(filepath.Join(".gemini", "rules", "MASTER.md")))
	requireFile(t, tempDir, filepath.ToSlash(filepath.Join(".gemini", "skills", "clean-code", "SKILL.md")))
	requireFile(t, tempDir, filepath.ToSlash(filepath.Join(".openkit", "managed.json")))

	st, err := managedstate.Load(managedstate.DefaultPath(tempDir))
	if err != nil {
		t.Fatalf("load managed state: %v", err)
	}
	if st.Agents["gemini"] == nil {
		t.Fatalf("expected gemini agent in managed state")
	}
	if st.Agents["gemini"].Pack.Version == "" {
		t.Fatalf("expected pack version to be set")
	}
	if st.Agents["gemini"].Files[".gemini/settings.json"] == nil {
		t.Fatalf("expected .gemini/settings.json managed entry")
	}

	// Second run should produce a no-op plan.
	if err := runAgentSync("gemini"); err != nil {
		t.Fatalf("second runAgentSync: %v", err)
	}

	ag := agents.Get("gemini")
	if ag == nil {
		t.Fatalf("missing gemini agent")
	}
	desired, err := targets.BuildEmbeddedDesired(ag, GetVersion())
	if err != nil {
		t.Fatalf("BuildEmbeddedDesired: %v", err)
	}
	st2, err := managedstate.Load(managedstate.DefaultPath(tempDir))
	if err != nil {
		t.Fatalf("reload managed state: %v", err)
	}
	plan, err := syncer.BuildPlan(tempDir, ag.ID, desired.Files, st2, syncer.Options{DryRun: true})
	if err != nil {
		t.Fatalf("BuildPlan: %v", err)
	}
	if len(plan.Create) != 0 || len(plan.Update) != 0 || len(plan.Overwrite) != 0 || len(plan.Conflicts) != 0 || len(plan.Delete) != 0 {
		t.Fatalf("expected no-op plan, got create=%d update=%d overwrite=%d conflicts=%d delete=%d",
			len(plan.Create), len(plan.Update), len(plan.Overwrite), len(plan.Conflicts), len(plan.Delete))
	}
}

func TestCursorSync_WritesFilesAndManagedState(t *testing.T) {
	tempDir := t.TempDir()
	oldWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("getwd: %v", err)
	}
	t.Cleanup(func() { _ = os.Chdir(oldWd) })
	if err := os.Chdir(tempDir); err != nil {
		t.Fatalf("chdir: %v", err)
	}

	flagDryRun = false
	flagOverwrite = false
	flagPrune = false

	if err := runAgentSync("cursor"); err != nil {
		t.Fatalf("runAgentSync: %v", err)
	}

	requireFile(t, tempDir, ".cursorrules")
	requireFile(t, tempDir, filepath.ToSlash(filepath.Join(".cursor", "rules", "openkit.mdc")))
	requireFile(t, tempDir, filepath.ToSlash(filepath.Join(".cursor", "skills", "clean-code", "SKILL.md")))
	requireFile(t, tempDir, filepath.ToSlash(filepath.Join(".openkit", "managed.json")))

	st, err := managedstate.Load(managedstate.DefaultPath(tempDir))
	if err != nil {
		t.Fatalf("load managed state: %v", err)
	}
	if st.Agents["cursor"] == nil {
		t.Fatalf("expected cursor agent in managed state")
	}
	if st.Agents["cursor"].Pack.Version == "" {
		t.Fatalf("expected pack version to be set")
	}
	if st.Agents["cursor"].Files[".cursorrules"] == nil {
		t.Fatalf("expected .cursorrules managed entry")
	}
	if st.Agents["cursor"].Files[".cursor/rules/openkit.mdc"] == nil {
		t.Fatalf("expected .cursor/rules/openkit.mdc managed entry")
	}

	// Second run should produce a no-op plan.
	if err := runAgentSync("cursor"); err != nil {
		t.Fatalf("second runAgentSync: %v", err)
	}

	ag := agents.Get("cursor")
	if ag == nil {
		t.Fatalf("missing cursor agent")
	}
	desired, err := targets.BuildEmbeddedDesired(ag, GetVersion())
	if err != nil {
		t.Fatalf("BuildEmbeddedDesired: %v", err)
	}
	st2, err := managedstate.Load(managedstate.DefaultPath(tempDir))
	if err != nil {
		t.Fatalf("reload managed state: %v", err)
	}
	plan, err := syncer.BuildPlan(tempDir, ag.ID, desired.Files, st2, syncer.Options{DryRun: true})
	if err != nil {
		t.Fatalf("BuildPlan: %v", err)
	}
	if len(plan.Create) != 0 || len(plan.Update) != 0 || len(plan.Overwrite) != 0 || len(plan.Conflicts) != 0 || len(plan.Delete) != 0 {
		t.Fatalf("expected no-op plan, got create=%d update=%d overwrite=%d conflicts=%d delete=%d",
			len(plan.Create), len(plan.Update), len(plan.Overwrite), len(plan.Conflicts), len(plan.Delete))
	}
}

func TestCodexSync_WritesFilesAndManagedState(t *testing.T) {
	tempDir := t.TempDir()
	oldWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("getwd: %v", err)
	}
	t.Cleanup(func() { _ = os.Chdir(oldWd) })
	if err := os.Chdir(tempDir); err != nil {
		t.Fatalf("chdir: %v", err)
	}

	flagDryRun = false
	flagOverwrite = false
	flagPrune = false

	if err := runAgentSync("codex"); err != nil {
		t.Fatalf("runAgentSync: %v", err)
	}

	requireFile(t, tempDir, "AGENTS.md")
	requireFile(t, tempDir, filepath.ToSlash(filepath.Join(".codex", "rules", "openkit.rules")))
	requireFile(t, tempDir, filepath.ToSlash(filepath.Join(".agents", "skills", "clean-code", "SKILL.md")))
	requireFile(t, tempDir, filepath.ToSlash(filepath.Join(".openkit", "managed.json")))

	st, err := managedstate.Load(managedstate.DefaultPath(tempDir))
	if err != nil {
		t.Fatalf("load managed state: %v", err)
	}
	if st.Agents["codex"] == nil {
		t.Fatalf("expected codex agent in managed state")
	}
	if st.Agents["codex"].Pack.Version == "" {
		t.Fatalf("expected pack version to be set")
	}
	if st.Agents["codex"].Files["AGENTS.md"] == nil {
		t.Fatalf("expected AGENTS.md managed entry")
	}
	if st.Agents["codex"].Files[".codex/rules/openkit.rules"] == nil {
		t.Fatalf("expected .codex/rules/openkit.rules managed entry")
	}

	// Second run should produce a no-op plan.
	if err := runAgentSync("codex"); err != nil {
		t.Fatalf("second runAgentSync: %v", err)
	}

	ag := agents.Get("codex")
	if ag == nil {
		t.Fatalf("missing codex agent")
	}
	desired, err := targets.BuildEmbeddedDesired(ag, GetVersion())
	if err != nil {
		t.Fatalf("BuildEmbeddedDesired: %v", err)
	}
	st2, err := managedstate.Load(managedstate.DefaultPath(tempDir))
	if err != nil {
		t.Fatalf("reload managed state: %v", err)
	}
	plan, err := syncer.BuildPlan(tempDir, ag.ID, desired.Files, st2, syncer.Options{DryRun: true})
	if err != nil {
		t.Fatalf("BuildPlan: %v", err)
	}
	if len(plan.Create) != 0 || len(plan.Update) != 0 || len(plan.Overwrite) != 0 || len(plan.Conflicts) != 0 || len(plan.Delete) != 0 {
		t.Fatalf("expected no-op plan, got create=%d update=%d overwrite=%d conflicts=%d delete=%d",
			len(plan.Create), len(plan.Update), len(plan.Overwrite), len(plan.Conflicts), len(plan.Delete))
	}
}

func requireFile(t *testing.T, root, rel string) {
	t.Helper()
	abs := filepath.Join(root, filepath.FromSlash(rel))
	fi, err := os.Stat(abs)
	if err != nil {
		t.Fatalf("expected file %s: %v", rel, err)
	}
	if fi.IsDir() {
		t.Fatalf("expected file, got dir: %s", rel)
	}
}
