use clap::{Args, Parser, Subcommand};
use regex::Regex;
use serde::{Deserialize, Serialize};
use std::collections::{BTreeMap, HashSet};
use std::fs;
use std::path::{Path, PathBuf};
use std::time::{SystemTime, UNIX_EPOCH};
use walkdir::WalkDir;

#[derive(Parser, Debug)]
#[command(name = "openkit-rs")]
#[command(about = "OpenKit Rust bootstrap CLI")]
struct Cli {
    #[command(subcommand)]
    command: Commands,
}

#[derive(Subcommand, Debug)]
enum Commands {
    Memory(MemoryCommand),
}

#[derive(Args, Debug)]
struct MemoryCommand {
    #[command(subcommand)]
    command: MemorySubcommand,
}

#[derive(Subcommand, Debug)]
enum MemorySubcommand {
    Init(InitArgs),
    Doctor(DoctorArgs),
    Capture(CaptureArgs),
    Review(ReviewArgs),
}

#[derive(Args, Debug)]
struct InitArgs {
    #[arg(long)]
    project: Option<PathBuf>,
    #[arg(long)]
    force: bool,
}

#[derive(Args, Debug)]
struct DoctorArgs {
    #[arg(long)]
    project: Option<PathBuf>,
    #[arg(long)]
    json: bool,
    #[arg(long)]
    write: bool,
}

#[derive(Args, Debug)]
struct CaptureArgs {
    #[arg(long)]
    project: Option<PathBuf>,
    #[arg(long)]
    session_id: Option<String>,
    #[arg(long)]
    summary: Option<String>,
    #[arg(long = "action")]
    actions: Vec<String>,
}

#[derive(Args, Debug)]
struct ReviewArgs {
    #[arg(long)]
    project: Option<PathBuf>,
    #[arg(long)]
    json: bool,
}

#[derive(Serialize, Deserialize)]
struct MemoryConfig {
    version: u8,
    mode: String,
    health_thresholds: HealthThresholds,
    linking: LinkingConfig,
}

#[derive(Serialize, Deserialize)]
struct HealthThresholds {
    healthy: u8,
    warning: u8,
}

#[derive(Serialize, Deserialize)]
struct LinkingConfig {
    require_inline_links: bool,
    require_related_section: bool,
}

#[derive(Serialize, Deserialize)]
struct DerivationState {
    version: u8,
    feature_slug: String,
    decisions: Decisions,
}

#[derive(Serialize, Deserialize)]
struct Decisions {
    runtime_language: String,
    migration_strategy: String,
    tier_policy: TierPolicy,
}

#[derive(Serialize, Deserialize)]
struct TierPolicy {
    tier1: Vec<String>,
    tier2: Vec<String>,
}

#[derive(Serialize, Deserialize)]
struct QueueFile {
    version: u8,
    items: Vec<QueueItem>,
}

#[derive(Serialize, Deserialize)]
struct QueueItem {
    id: String,
    r#type: String,
    status: String,
    title: String,
}

#[derive(Serialize, Deserialize)]
struct SessionSnapshot {
    version: u8,
    session_id: String,
    started_at: String,
    ended_at: String,
    summary: String,
    actions: Vec<String>,
}

#[derive(Serialize)]
struct DoctorReport {
    version: u8,
    score: u8,
    status: String,
    checks: BTreeMap<String, String>,
}

#[derive(Serialize)]
struct ReviewReport {
    sessions: usize,
    observations: usize,
    tensions: usize,
    recommendations: Vec<String>,
}

fn main() -> Result<(), String> {
    let cli = Cli::parse();
    match cli.command {
        Commands::Memory(memory) => match memory.command {
            MemorySubcommand::Init(args) => memory_init(args),
            MemorySubcommand::Doctor(args) => memory_doctor(args),
            MemorySubcommand::Capture(args) => memory_capture(args),
            MemorySubcommand::Review(args) => memory_review(args),
        },
    }
}

fn memory_init(args: InitArgs) -> Result<(), String> {
    let root = project_root(args.project)?;
    let memory_dir = root.join(".openkit/memory");
    let ops_dir = root.join(".openkit/ops");
    let dirs = [
        ops_dir.join("sessions"),
        ops_dir.join("observations"),
        ops_dir.join("tensions"),
        ops_dir.join("health"),
        ops_dir.join("queue"),
        memory_dir.clone(),
    ];

    for dir in dirs {
        fs::create_dir_all(&dir)
            .map_err(|e| format!("failed to create {}: {}", dir.display(), e))?;
    }

    let config = MemoryConfig {
        version: 1,
        mode: "assisted".to_string(),
        health_thresholds: HealthThresholds {
            healthy: 85,
            warning: 70,
        },
        linking: LinkingConfig {
            require_inline_links: true,
            require_related_section: true,
        },
    };

    let derivation = DerivationState {
        version: 1,
        feature_slug: "memory-kernel-rust-cli".to_string(),
        decisions: Decisions {
            runtime_language: "rust".to_string(),
            migration_strategy: "strangler".to_string(),
            tier_policy: TierPolicy {
                tier1: vec!["opencode".to_string()],
                tier2: vec![
                    "claude-code".to_string(),
                    "codex".to_string(),
                    "antigravity".to_string(),
                ],
            },
        },
    };

    let queue = QueueFile {
        version: 1,
        items: vec![QueueItem {
            id: "MK-001".to_string(),
            r#type: "maintenance".to_string(),
            status: "pending".to_string(),
            title: "Resolve stale links in requirements docs".to_string(),
        }],
    };

    write_yaml(
        &memory_dir.join("config.yaml"),
        &config,
        args.force,
    )?;
    write_yaml(
        &memory_dir.join("derivation.yaml"),
        &derivation,
        args.force,
    )?;
    write_yaml(&ops_dir.join("queue.yaml"), &queue, args.force)?;

    println!("Initialized Memory Kernel structure at {}", root.display());
    Ok(())
}

fn memory_doctor(args: DoctorArgs) -> Result<(), String> {
    let root = project_root(args.project)?;
    let docs = root.join("docs");
    let (report, broken_links) = build_doctor_report(&docs)?;

    if args.write {
        let health_file = root.join(".openkit/ops/health/memory-health.json");
        if let Some(parent) = health_file.parent() {
            fs::create_dir_all(parent)
                .map_err(|e| format!("failed to create {}: {}", parent.display(), e))?;
        }
        let payload = serde_json::to_string_pretty(&report)
            .map_err(|e| format!("failed to serialize doctor report: {}", e))?;
        fs::write(&health_file, payload)
            .map_err(|e| format!("failed to write {}: {}", health_file.display(), e))?;
    }

    if args.json {
        let payload = serde_json::to_string_pretty(&report)
            .map_err(|e| format!("failed to serialize doctor report: {}", e))?;
        println!("{}", payload);
    } else {
        println!("Memory Health: {} (score={})", report.status, report.score);
        for (check, status) in report.checks {
            println!("- {}: {}", check, status);
        }
    }

    if !broken_links.is_empty() {
        let preview: Vec<String> = broken_links.into_iter().take(3).collect();
        return Err(format!(
            "memory doctor failed: found broken wikilinks. Examples: {}",
            preview.join(" | ")
        ));
    }

    Ok(())
}

fn memory_capture(args: CaptureArgs) -> Result<(), String> {
    let root = project_root(args.project)?;
    let sessions_dir = root.join(".openkit/ops/sessions");
    fs::create_dir_all(&sessions_dir)
        .map_err(|e| format!("failed to create {}: {}", sessions_dir.display(), e))?;

    let now = timestamp_secs();
    let session_id = args
        .session_id
        .unwrap_or_else(|| format!("mk-{}", now));
    let snapshot = SessionSnapshot {
        version: 1,
        session_id: session_id.clone(),
        started_at: now.to_string(),
        ended_at: now.to_string(),
        summary: args
            .summary
            .unwrap_or_else(|| "OpenKit memory session capture".to_string()),
        actions: if args.actions.is_empty() {
            vec!["capture".to_string()]
        } else {
            args.actions
        },
    };

    let file_path = sessions_dir.join(format!("{}.json", now));
    let payload = serde_json::to_string_pretty(&snapshot)
        .map_err(|e| format!("failed to serialize session snapshot: {}", e))?;
    fs::write(&file_path, payload)
        .map_err(|e| format!("failed to write {}: {}", file_path.display(), e))?;

    println!("Captured session {}", session_id);
    println!("Snapshot: {}", file_path.display());
    Ok(())
}

fn memory_review(args: ReviewArgs) -> Result<(), String> {
    let root = project_root(args.project)?;
    let sessions = count_files(&root.join(".openkit/ops/sessions"));
    let observations = count_files(&root.join(".openkit/ops/observations"));
    let tensions = count_files(&root.join(".openkit/ops/tensions"));

    let mut recommendations = Vec::new();
    if observations >= 10 {
        recommendations.push("Run memory review for accumulated observations".to_string());
    }
    if tensions >= 5 {
        recommendations.push("Resolve repeated tensions before next implementation phase".to_string());
    }
    if sessions >= 5 {
        recommendations.push("Summarize recent sessions into sprint artifacts".to_string());
    }
    if recommendations.is_empty() {
        recommendations.push("Memory operations are within thresholds".to_string());
    }

    let report = ReviewReport {
        sessions,
        observations,
        tensions,
        recommendations,
    };

    if args.json {
        let payload = serde_json::to_string_pretty(&report)
            .map_err(|e| format!("failed to serialize review report: {}", e))?;
        println!("{}", payload);
    } else {
        println!(
            "Memory Review: sessions={}, observations={}, tensions={}",
            report.sessions, report.observations, report.tensions
        );
        for item in report.recommendations {
            println!("- {}", item);
        }
    }

    Ok(())
}

fn build_doctor_report(docs_root: &Path) -> Result<(DoctorReport, Vec<String>), String> {
    let mut checks = BTreeMap::new();

    let inline_ok = check_inline_links(docs_root)?;
    checks.insert(
        "inline_links".to_string(),
        if inline_ok { "pass" } else { "fail" }.to_string(),
    );

    let related_ok = check_related_sections(docs_root)?;
    checks.insert(
        "related_sections".to_string(),
        if related_ok { "pass" } else { "fail" }.to_string(),
    );

    let broken_links = broken_wikilinks(docs_root)?;
    let broken_count = broken_links.len();
    checks.insert(
        "broken_wikilinks".to_string(),
        if broken_count == 0 {
            "pass".to_string()
        } else {
            format!("fail({})", broken_count)
        },
    );

    let stale_status = stale_docs_status(docs_root)?;
    checks.insert("stale_docs".to_string(), stale_status.clone());

    let mut score: i16 = 100;
    if !inline_ok {
        score -= 25;
    }
    if !related_ok {
        score -= 20;
    }
    if broken_count > 0 {
        score -= 30;
    }
    if stale_status == "warn" {
        score -= 10;
    }

    let clamped = score.clamp(0, 100) as u8;
    let status = if clamped >= 85 {
        "healthy"
    } else if clamped >= 70 {
        "warning"
    } else {
        "critical"
    }
    .to_string();

    Ok((
        DoctorReport {
            version: 1,
            score: clamped,
            status,
            checks,
        },
        broken_links,
    ))
}

fn check_inline_links(docs_root: &Path) -> Result<bool, String> {
    let re = Regex::new(r"\[\[[^\]]+\]\]").map_err(|e| e.to_string())?;
    for entry in WalkDir::new(docs_root)
        .into_iter()
        .filter_map(Result::ok)
        .filter(|e| e.file_type().is_file())
    {
        if entry.path().extension().and_then(|s| s.to_str()) != Some("md") {
            continue;
        }
        let content = fs::read_to_string(entry.path()).map_err(|e| {
            format!(
                "failed reading {} for inline link check: {}",
                entry.path().display(),
                e
            )
        })?;
        if !re.is_match(&content) {
            continue;
        }
        if let Some(pos) = content.find("## Related") {
            let before = &content[..pos];
            if re.is_match(before) {
                return Ok(true);
            }
        } else {
            return Ok(true);
        }
    }
    Ok(false)
}

fn check_related_sections(docs_root: &Path) -> Result<bool, String> {
    let required = [
        "HUB-DOCS.md",
        "CONTEXT.md",
        "SECURITY.md",
        "QUALITY_GATES.md",
        "requirements/HUB-REQUIREMENTS.md",
        "sprint/HUB-SPRINTS.md",
    ];
    for rel in required {
        let path = docs_root.join(rel);
        let content = fs::read_to_string(&path)
            .map_err(|e| format!("failed reading {}: {}", path.display(), e))?;
        if !content.contains("## Related") {
            return Ok(false);
        }
    }
    Ok(true)
}

fn broken_wikilinks(docs_root: &Path) -> Result<Vec<String>, String> {
    let re = Regex::new(r"\[\[([^\]]+)\]\]").map_err(|e| e.to_string())?;
    let mut existing = HashSet::new();

    for entry in WalkDir::new(docs_root)
        .into_iter()
        .filter_map(Result::ok)
        .filter(|e| e.file_type().is_file())
    {
        if entry.path().extension().and_then(|s| s.to_str()) == Some("md") {
            let rel = entry
                .path()
                .strip_prefix(docs_root)
                .map_err(|e| e.to_string())?
                .to_string_lossy()
                .replace('\\', "/");
            existing.insert(rel);
        }
    }

    let mut broken = Vec::new();
    for entry in WalkDir::new(docs_root)
        .into_iter()
        .filter_map(Result::ok)
        .filter(|e| e.file_type().is_file())
    {
        if entry.path().extension().and_then(|s| s.to_str()) != Some("md") {
            continue;
        }
        let content = fs::read_to_string(entry.path())
            .map_err(|e| format!("failed reading {}: {}", entry.path().display(), e))?;
        for caps in re.captures_iter(&content) {
            let raw = caps.get(1).map(|m| m.as_str()).unwrap_or_default();
            let link = raw.split('#').next().unwrap_or("").trim();
            if link.is_empty() {
                continue;
            }
            let normalized = link.trim_start_matches("docs/");
            if !existing.contains(normalized) {
                broken.push(format!(
                    "{} -> [[{}]]",
                    entry.path().display(),
                    link
                ));
            }
        }
    }
    Ok(broken)
}

fn stale_docs_status(docs_root: &Path) -> Result<String, String> {
    let mut stale = 0usize;
    let now = SystemTime::now();
    for entry in WalkDir::new(docs_root)
        .into_iter()
        .filter_map(Result::ok)
        .filter(|e| e.file_type().is_file())
    {
        if entry.path().extension().and_then(|s| s.to_str()) != Some("md") {
            continue;
        }
        let meta = fs::metadata(entry.path())
            .map_err(|e| format!("failed stat {}: {}", entry.path().display(), e))?;
        if let Ok(modified) = meta.modified() {
            if let Ok(age) = now.duration_since(modified) {
                if age.as_secs() > 60 * 60 * 24 * 45 {
                    stale += 1;
                }
            }
        }
    }
    if stale > 0 {
        Ok("warn".to_string())
    } else {
        Ok("pass".to_string())
    }
}

fn project_root(project: Option<PathBuf>) -> Result<PathBuf, String> {
    if let Some(path) = project {
        return Ok(path);
    }
    std::env::current_dir().map_err(|e| format!("failed to detect current directory: {}", e))
}

fn write_yaml<T: Serialize>(path: &Path, value: &T, force: bool) -> Result<(), String> {
    if path.exists() && !force {
        return Ok(());
    }
    if let Some(parent) = path.parent() {
        fs::create_dir_all(parent)
            .map_err(|e| format!("failed to create {}: {}", parent.display(), e))?;
    }
    let payload = serde_yaml::to_string(value)
        .map_err(|e| format!("failed to serialize yaml for {}: {}", path.display(), e))?;
    fs::write(path, payload).map_err(|e| format!("failed writing {}: {}", path.display(), e))
}

fn count_files(path: &Path) -> usize {
    if !path.exists() {
        return 0;
    }
    WalkDir::new(path)
        .into_iter()
        .filter_map(Result::ok)
        .filter(|e| e.file_type().is_file())
        .count()
}

fn timestamp_secs() -> u64 {
    SystemTime::now()
        .duration_since(UNIX_EPOCH)
        .map(|d| d.as_secs())
        .unwrap_or(0)
}
