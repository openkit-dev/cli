use assert_cmd::prelude::*;
use serde_json::Value as JsonValue;
use serde_yaml::Value as YamlValue;
use std::fs;
use std::path::{Path, PathBuf};
use std::process::Command;
use tempfile::tempdir;

fn write_file(path: &Path, content: &str) {
    if let Some(parent) = path.parent() {
        fs::create_dir_all(parent).expect("failed creating parent dir");
    }
    fs::write(path, content).expect("failed writing file");
}

fn fixture(path: &str) -> String {
    let base = PathBuf::from(env!("CARGO_MANIFEST_DIR"))
        .join("tests")
        .join("golden")
        .join(path);
    fs::read_to_string(base).expect("failed reading fixture")
}

fn setup_docs(root: &Path) {
    write_file(
        &root.join("openkit-memory/HUB-DOCS.md"),
        "# HUB\nSee [[CONTEXT.md]]\n## Related\n- [[CONTEXT.md]]\n- [[SECURITY.md]]\n",
    );
    write_file(
        &root.join("openkit-memory/CONTEXT.md"),
        "# CONTEXT\n## Related\n- [[HUB-DOCS.md]]\n",
    );
    write_file(
        &root.join("openkit-memory/SECURITY.md"),
        "# SECURITY\n## Related\n- [[HUB-DOCS.md]]\n",
    );
    write_file(
        &root.join("openkit-memory/QUALITY_GATES.md"),
        "# QUALITY\n## Related\n- [[HUB-DOCS.md]]\n",
    );
    write_file(
        &root.join("openkit-memory/requirements/HUB-REQUIREMENTS.md"),
        "# REQ\n## Related\n- [[HUB-DOCS.md]]\n",
    );
    write_file(
        &root.join("openkit-memory/sprint/HUB-SPRINTS.md"),
        "# SPRINT\n## Related\n- [[HUB-DOCS.md]]\n",
    );
}

#[test]
fn memory_init_creates_contract_files_and_is_idempotent() {
    let temp = tempdir().expect("failed to create temp dir");
    let root = temp.path();

    let mut cmd = Command::new(assert_cmd::cargo::cargo_bin!("openkit"));
    cmd.args(["memory", "init", "--project"])
        .arg(root)
        .assert()
        .success();

    let mut cmd_second = Command::new(assert_cmd::cargo::cargo_bin!("openkit"));
    cmd_second
        .args(["memory", "init", "--project"])
        .arg(root)
        .assert()
        .success();

    let config_actual: YamlValue = serde_yaml::from_str(
        &fs::read_to_string(root.join(".openkit/memory/config.yaml")).expect("missing config.yaml"),
    )
    .expect("invalid config.yaml");
    let config_expected: YamlValue =
        serde_yaml::from_str(&fixture("config.yaml")).expect("invalid fixture config.yaml");
    assert_eq!(config_actual, config_expected);

    let derivation_actual: YamlValue = serde_yaml::from_str(
        &fs::read_to_string(root.join(".openkit/memory/derivation.yaml"))
            .expect("missing derivation.yaml"),
    )
    .expect("invalid derivation.yaml");
    let derivation_expected: YamlValue =
        serde_yaml::from_str(&fixture("derivation.yaml")).expect("invalid fixture derivation.yaml");
    assert_eq!(derivation_actual, derivation_expected);

    let queue_actual: YamlValue = serde_yaml::from_str(
        &fs::read_to_string(root.join(".openkit/ops/queue.yaml")).expect("missing queue.yaml"),
    )
    .expect("invalid queue.yaml");
    let queue_expected: YamlValue =
        serde_yaml::from_str(&fixture("queue.yaml")).expect("invalid fixture queue.yaml");
    assert_eq!(queue_actual, queue_expected);
}

#[test]
fn memory_doctor_matches_golden_for_healthy_docs() {
    let temp = tempdir().expect("failed to create temp dir");
    let root = temp.path();
    setup_docs(root);

    let output = Command::new(assert_cmd::cargo::cargo_bin!("openkit"))
        .args(["memory", "doctor", "--json", "--project"])
        .arg(root)
        .output()
        .expect("failed to run doctor");

    assert!(output.status.success());
    let stdout = String::from_utf8(output.stdout).expect("doctor output not utf8");

    let actual: JsonValue = serde_json::from_str(stdout.trim()).expect("invalid doctor json");
    let expected: JsonValue =
        serde_json::from_str(&fixture("doctor_healthy.json")).expect("invalid fixture json");
    assert_eq!(actual, expected);
}

#[test]
fn memory_doctor_reports_broken_links_with_actionable_error() {
    let temp = tempdir().expect("failed to create temp dir");
    let root = temp.path();
    setup_docs(root);
    write_file(
        &root.join("openkit-memory/BROKEN.md"),
        "# BROKEN\nBad link [[MISSING.md]]\n## Related\n- [[HUB-DOCS.md]]\n",
    );

    let output = Command::new(assert_cmd::cargo::cargo_bin!("openkit"))
        .args(["memory", "doctor", "--json", "--project"])
        .arg(root)
        .output()
        .expect("failed to run doctor");

    assert!(!output.status.success());
    let stderr = String::from_utf8(output.stderr).expect("stderr not utf8");
    assert!(stderr.contains("memory doctor failed: found broken wikilinks"));
    assert!(stderr.contains("BROKEN.md"));
}

#[test]
fn memory_capture_and_review_follow_contract() {
    let temp = tempdir().expect("failed to create temp dir");
    let root = temp.path();
    setup_docs(root);

    let capture = Command::new(assert_cmd::cargo::cargo_bin!("openkit"))
        .args([
            "memory",
            "capture",
            "--project",
            root.to_str().expect("root path is not utf8"),
            "--session-id",
            "session-golden",
            "--summary",
            "Golden capture",
            "--action",
            "init",
            "--action",
            "doctor",
        ])
        .output()
        .expect("failed to run capture");
    assert!(capture.status.success());

    let sessions_dir = root.join(".openkit/ops/sessions");
    let entries = fs::read_dir(&sessions_dir).expect("missing sessions dir");
    let files: Vec<PathBuf> = entries.filter_map(Result::ok).map(|e| e.path()).collect();
    assert_eq!(files.len(), 1);

    let snapshot: JsonValue = serde_json::from_str(
        &fs::read_to_string(&files[0]).expect("missing snapshot file content"),
    )
    .expect("invalid snapshot json");
    assert_eq!(snapshot["version"], 1);
    assert_eq!(snapshot["session_id"], "session-golden");
    assert_eq!(snapshot["summary"], "Golden capture");

    for i in 0..10 {
        write_file(
            &root.join(format!(".openkit/ops/observations/{}.md", i)),
            "obs",
        );
    }
    for i in 0..5 {
        write_file(&root.join(format!(".openkit/ops/tensions/{}.md", i)), "ten");
    }
    for i in 0..4 {
        write_file(
            &root.join(format!(".openkit/ops/sessions/seed-{}.json", i)),
            "{}",
        );
    }

    let review = Command::new(assert_cmd::cargo::cargo_bin!("openkit"))
        .args(["memory", "review", "--json", "--project"])
        .arg(root)
        .output()
        .expect("failed to run review");
    assert!(review.status.success());

    let review_json: JsonValue =
        serde_json::from_str(&String::from_utf8(review.stdout).expect("review stdout not utf8"))
            .expect("invalid review json");
    let expected: JsonValue =
        serde_json::from_str(&fixture("review_thresholds.json")).expect("invalid review fixture");

    assert_eq!(review_json, expected);
}

#[test]
fn check_command_returns_json_schema() {
    let output = Command::new(assert_cmd::cargo::cargo_bin!("openkit"))
        .args(["check", "--json"])
        .output()
        .expect("failed to run check --json");

    assert!(output.status.success());
    let payload = String::from_utf8(output.stdout).expect("check stdout not utf8");
    let data: JsonValue = serde_json::from_str(&payload).expect("invalid check json");

    assert!(data.get("platform").and_then(|v| v.as_str()).is_some());
    assert!(data.get("ready").and_then(|v| v.as_bool()).is_some());

    let agents = data
        .get("agents")
        .and_then(|v| v.as_array())
        .expect("agents must be array");
    let tools = data
        .get("tools")
        .and_then(|v| v.as_array())
        .expect("tools must be array");

    assert_eq!(agents.len(), 5);
    assert_eq!(tools.len(), 4);
}

#[test]
fn init_command_creates_baseline_project_artifacts() {
    let temp = tempdir().expect("failed to create temp dir");
    let root = temp.path();
    let project_name = "demo-app";

    let output = Command::new(assert_cmd::cargo::cargo_bin!("openkit"))
        .current_dir(root)
        .args(["init", project_name, "--agent", "opencode", "--no-git"])
        .output()
        .expect("failed to run init command");

    assert!(output.status.success());

    let project_root = root.join(project_name);
    assert!(project_root.join("AGENTS.md").exists());
    assert!(project_root.join("memory/HUB-DOCS.md").exists());
    assert!(project_root.join("memory/CONTEXT.md").exists());
    assert!(project_root.join("memory/SECURITY.md").exists());
    assert!(project_root
        .join("memory/QUALITY_GATES.md")
        .exists());
    assert!(project_root
        .join("memory/requirements/HUB-REQUIREMENTS.md")
        .exists());
    assert!(project_root
        .join("memory/requirements/bootstrap/HUB-BOOTSTRAP.md")
        .exists());
    assert!(project_root
        .join("memory/sprint/HUB-SPRINTS.md")
        .exists());
    assert!(project_root
        .join("memory/sprint/Sprint-01/HUB-SPRINT-01.md")
        .exists());
    assert!(project_root.join("opencode.json").exists());
    assert!(project_root.join(".opencode/commands/discover.md").exists());
    assert!(project_root
        .join(".opencode/prompts/orchestrator.md")
        .exists());
    assert!(project_root.join(".opencode/rules/MASTER.md").exists());
    assert!(project_root
        .join(".opencode/skills/clean-code/SKILL.md")
        .exists());
    assert!(project_root
        .join(".opencode/rules/MEMORY_KERNEL.md")
        .exists());
    assert!(project_root.join(".openkit/memory/config.yaml").exists());
    assert!(project_root.join(".opencode/OPENKIT.md").exists());
}

#[test]
fn init_without_name_initializes_current_directory() {
    let temp = tempdir().expect("failed to create temp dir");
    let root = temp.path();

    let output = Command::new(assert_cmd::cargo::cargo_bin!("openkit"))
        .current_dir(root)
        .args(["init", "--agent", "opencode", "--no-git"])
        .output()
        .expect("failed to run init in current directory");

    assert!(output.status.success());
    assert!(root.join("AGENTS.md").exists());
    assert!(root.join("memory/HUB-DOCS.md").exists());
    assert!(root.join(".opencode/OPENKIT.md").exists());

    let stdout = String::from_utf8(output.stdout).expect("stdout not utf8");
    assert!(!stdout.contains("\n  cd "));
}

#[test]
fn init_overwrite_replaces_existing_agent_pack() {
    let temp = tempdir().expect("failed to create temp dir");
    let root = temp.path();

    write_file(&root.join(".opencode/OPENKIT.md"), "legacy marker\n");
    write_file(&root.join(".opencode/legacy.txt"), "legacy content\n");

    let no_overwrite = Command::new(assert_cmd::cargo::cargo_bin!("openkit"))
        .current_dir(root)
        .args(["init", "--agent", "opencode", "--no-git"])
        .output()
        .expect("failed to run init without overwrite");
    assert!(no_overwrite.status.success());

    let preserved = fs::read_to_string(root.join(".opencode/OPENKIT.md"))
        .expect("missing marker after init without overwrite");
    assert_eq!(preserved, "legacy marker\n");

    let with_overwrite = Command::new(assert_cmd::cargo::cargo_bin!("openkit"))
        .current_dir(root)
        .args(["init", "--agent", "opencode", "--overwrite", "--no-git"])
        .output()
        .expect("failed to run init with overwrite");
    assert!(with_overwrite.status.success());

    let replaced = fs::read_to_string(root.join(".opencode/OPENKIT.md"))
        .expect("missing marker after overwrite init");
    assert!(replaced.contains("Generated by Rust OpenKit runtime"));
    assert!(!root.join(".opencode/legacy.txt").exists());
}

#[test]
fn sync_and_doctor_json_work() {
    let temp = tempdir().expect("failed to create temp dir");
    let root = temp.path();

    let sync = Command::new(assert_cmd::cargo::cargo_bin!("openkit"))
        .current_dir(root)
        .args(["sync", "--agent", "opencode", "--overwrite"])
        .output()
        .expect("failed to run sync");
    assert!(sync.status.success());
    assert!(root.join(".opencode/OPENKIT.md").exists());

    let doctor = Command::new(assert_cmd::cargo::cargo_bin!("openkit"))
        .current_dir(root)
        .args(["doctor", "--agent", "opencode", "--json"])
        .output()
        .expect("failed to run doctor");
    assert!(doctor.status.success());

    let payload = String::from_utf8(doctor.stdout).expect("doctor stdout not utf8");
    let data: JsonValue = serde_json::from_str(&payload).expect("invalid doctor json");
    assert_eq!(data["agent"], "opencode");
    assert_eq!(data["status"], "healthy");
}

#[test]
fn init_then_memory_doctor_is_healthy() {
    let temp = tempdir().expect("failed to create temp dir");
    let root = temp.path();

    let init = Command::new(assert_cmd::cargo::cargo_bin!("openkit"))
        .current_dir(root)
        .args(["init", "doctor-app", "--agent", "opencode", "--no-git"])
        .output()
        .expect("failed to run init");
    assert!(init.status.success());

    let doctor = Command::new(assert_cmd::cargo::cargo_bin!("openkit"))
        .current_dir(root.join("doctor-app"))
        .args(["memory", "doctor", "--json"])
        .output()
        .expect("failed to run memory doctor");
    assert!(doctor.status.success());

    let payload = String::from_utf8(doctor.stdout).expect("doctor stdout not utf8");
    let data: JsonValue = serde_json::from_str(&payload).expect("invalid doctor json");
    assert_eq!(data["status"], "healthy");
}

#[test]
fn init_with_codex_flag_materializes_codex_pack() {
    let temp = tempdir().expect("failed to create temp dir");
    let root = temp.path();

    let output = Command::new(assert_cmd::cargo::cargo_bin!("openkit"))
        .current_dir(root)
        .args(["init", "codex-app", "--agent", "codex", "--no-git"])
        .output()
        .expect("failed to run init with --agent codex");

    assert!(output.status.success());

    let project_root = root.join("codex-app");
    assert!(project_root.join(".codex/commands/discover.md").exists());
    assert!(project_root.join(".codex/rules/MEMORY_KERNEL.md").exists());
    assert!(project_root.join(".openkit/memory/config.yaml").exists());
}

#[test]
fn upgrade_dry_run_succeeds() {
    let output = Command::new(assert_cmd::cargo::cargo_bin!("openkit"))
        .args(["upgrade", "--dry-run"])
        .output()
        .expect("failed to run upgrade --dry-run");
    assert!(output.status.success());
}

#[test]
fn uninstall_dry_run_lists_targets_and_succeeds() {
    let output = Command::new(assert_cmd::cargo::cargo_bin!("openkit"))
        .args(["uninstall", "--dry-run"])
        .output()
        .expect("failed to run uninstall --dry-run");

    assert!(output.status.success());
    let stdout = String::from_utf8(output.stdout).expect("uninstall stdout not utf8");
    assert!(stdout.contains("Dry run:") || stdout.contains("not found in known install paths"));
}
