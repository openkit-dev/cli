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
        &root.join("docs/HUB-DOCS.md"),
        "# HUB\nSee [[CONTEXT.md]]\n## Related\n- [[CONTEXT.md]]\n- [[SECURITY.md]]\n",
    );
    write_file(
        &root.join("docs/CONTEXT.md"),
        "# CONTEXT\n## Related\n- [[HUB-DOCS.md]]\n",
    );
    write_file(
        &root.join("docs/SECURITY.md"),
        "# SECURITY\n## Related\n- [[HUB-DOCS.md]]\n",
    );
    write_file(
        &root.join("docs/QUALITY_GATES.md"),
        "# QUALITY\n## Related\n- [[HUB-DOCS.md]]\n",
    );
    write_file(
        &root.join("docs/requirements/HUB-REQUIREMENTS.md"),
        "# REQ\n## Related\n- [[HUB-DOCS.md]]\n",
    );
    write_file(
        &root.join("docs/sprint/HUB-SPRINTS.md"),
        "# SPRINT\n## Related\n- [[HUB-DOCS.md]]\n",
    );
}

#[test]
fn memory_init_creates_contract_files_and_is_idempotent() {
    let temp = tempdir().expect("failed to create temp dir");
    let root = temp.path();

    let mut cmd = Command::new(assert_cmd::cargo::cargo_bin!("openkit-rs"));
    cmd.args(["memory", "init", "--project"])
        .arg(root)
        .assert()
        .success();

    let mut cmd_second = Command::new(assert_cmd::cargo::cargo_bin!("openkit-rs"));
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
    let derivation_expected: YamlValue = serde_yaml::from_str(&fixture("derivation.yaml"))
        .expect("invalid fixture derivation.yaml");
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

    let output = Command::new(assert_cmd::cargo::cargo_bin!("openkit-rs"))
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
        &root.join("docs/BROKEN.md"),
        "# BROKEN\nBad link [[MISSING.md]]\n## Related\n- [[HUB-DOCS.md]]\n",
    );

    let output = Command::new(assert_cmd::cargo::cargo_bin!("openkit-rs"))
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

    let capture = Command::new(assert_cmd::cargo::cargo_bin!("openkit-rs"))
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
    let files: Vec<PathBuf> = entries
        .filter_map(Result::ok)
        .map(|e| e.path())
        .collect();
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
        write_file(
            &root.join(format!(".openkit/ops/tensions/{}.md", i)),
            "ten",
        );
    }
    for i in 0..4 {
        write_file(
            &root.join(format!(".openkit/ops/sessions/seed-{}.json", i)),
            "{}",
        );
    }

    let review = Command::new(assert_cmd::cargo::cargo_bin!("openkit-rs"))
        .args(["memory", "review", "--json", "--project"])
        .arg(root)
        .output()
        .expect("failed to run review");
    assert!(review.status.success());

    let review_json: JsonValue = serde_json::from_str(
        &String::from_utf8(review.stdout).expect("review stdout not utf8"),
    )
    .expect("invalid review json");
    let expected: JsonValue =
        serde_json::from_str(&fixture("review_thresholds.json")).expect("invalid review fixture");

    assert_eq!(review_json, expected);
}
