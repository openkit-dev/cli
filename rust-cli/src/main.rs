use clap::{Args, Parser, Subcommand};
use regex::Regex;
use serde::{Deserialize, Serialize};
use std::collections::{BTreeMap, HashSet};
use std::fs;
use std::path::{Path, PathBuf};
use std::process::Command;
use std::time::{SystemTime, UNIX_EPOCH};
use walkdir::WalkDir;

#[derive(Parser, Debug)]
#[command(name = "openkit")]
#[command(about = "OpenKit Rust CLI")]
#[command(version)]
struct Cli {
    #[command(subcommand)]
    command: Commands,
}

#[derive(Subcommand, Debug)]
enum Commands {
    Check(CheckArgs),
    Init(ProjectInitArgs),
    Upgrade(UpgradeArgs),
    Memory(MemoryCommand),
    Opencode(AgentCommand),
    Claude(AgentCommand),
    Cursor(AgentCommand),
    Gemini(AgentCommand),
    Codex(AgentCommand),
    Antigravity(AgentCommand),
}

#[derive(Args, Debug)]
struct CheckArgs {
    #[arg(long)]
    json: bool,
}

#[derive(Args, Debug)]
struct UpgradeArgs {
    #[arg(long)]
    check: bool,
    #[arg(long)]
    dry_run: bool,
}

#[derive(Args, Debug)]
struct AgentCommand {
    #[command(subcommand)]
    command: AgentSubcommand,
}

#[derive(Subcommand, Debug)]
enum AgentSubcommand {
    Sync(AgentSyncArgs),
    Doctor(AgentDoctorArgs),
    Upgrade(AgentSyncArgs),
}

#[derive(Args, Debug, Clone)]
struct AgentSyncArgs {
    #[arg(long)]
    dry_run: bool,
    #[arg(long)]
    overwrite: bool,
    #[arg(long)]
    prune: bool,
}

#[derive(Args, Debug)]
struct AgentDoctorArgs {
    #[arg(long)]
    json: bool,
}

#[derive(Args, Debug)]
struct MemoryCommand {
    #[command(subcommand)]
    command: MemorySubcommand,
}

#[derive(Subcommand, Debug)]
enum MemorySubcommand {
    Init(MemoryInitArgs),
    Doctor(DoctorArgs),
    Capture(CaptureArgs),
    Review(ReviewArgs),
}

#[derive(Args, Debug)]
struct MemoryInitArgs {
    #[arg(long)]
    project: Option<PathBuf>,
    #[arg(long)]
    force: bool,
}

#[derive(Args, Debug)]
struct ProjectInitArgs {
    project_name: Option<String>,
    #[arg(long)]
    ai: Option<String>,
    #[arg(long)]
    here: bool,
    #[arg(long)]
    force: bool,
    #[arg(long = "no-git")]
    no_git: bool,
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

#[derive(Serialize)]
struct CheckEntry {
    name: String,
    command: String,
    status: String,
    version: Option<String>,
    note: Option<String>,
}

#[derive(Serialize)]
struct CheckReport {
    platform: String,
    agents: Vec<CheckEntry>,
    tools: Vec<CheckEntry>,
    ready: bool,
}

#[derive(Serialize)]
struct AgentDoctorReport {
    agent: String,
    config_dir: String,
    status: String,
    notes: Vec<String>,
}

fn main() -> Result<(), String> {
    let cli = Cli::parse();
    match cli.command {
        Commands::Check(args) => run_check(args),
        Commands::Init(args) => run_init(args),
        Commands::Upgrade(args) => run_upgrade(args),
        Commands::Opencode(cmd) => run_agent_command("opencode", cmd),
        Commands::Claude(cmd) => run_agent_command("claude", cmd),
        Commands::Cursor(cmd) => run_agent_command("cursor", cmd),
        Commands::Gemini(cmd) => run_agent_command("gemini", cmd),
        Commands::Codex(cmd) => run_agent_command("codex", cmd),
        Commands::Antigravity(cmd) => run_agent_command("antigravity", cmd),
        Commands::Memory(memory) => match memory.command {
            MemorySubcommand::Init(args) => memory_init(args),
            MemorySubcommand::Doctor(args) => memory_doctor(args),
            MemorySubcommand::Capture(args) => memory_capture(args),
            MemorySubcommand::Review(args) => memory_review(args),
        },
    }
}

fn run_upgrade(args: UpgradeArgs) -> Result<(), String> {
    let repo = std::env::var("OPENKIT_REPO").unwrap_or_else(|_| "orionlabz/openkit".to_string());
    let current = env!("CARGO_PKG_VERSION");

    if args.check {
        let latest = fetch_latest_tag(&repo)?;
        println!("Current: {}", current);
        println!("Latest:  {}", latest);
        if latest.trim_start_matches('v') == current {
            println!("OpenKit is up to date.");
        } else {
            println!("Update available.");
        }
        return Ok(());
    }

    if std::env::consts::OS == "windows" {
        let cmd = "irm https://raw.githubusercontent.com/orionlabz/openkit/main/scripts/install.ps1 | iex";
        if args.dry_run {
            println!("Dry run: powershell -NoProfile -Command \"{}\"", cmd);
            return Ok(());
        }
        let status = Command::new("powershell")
            .args(["-NoProfile", "-Command", cmd])
            .status()
            .map_err(|e| format!("failed to execute upgrade installer: {}", e))?;
        if !status.success() {
            return Err("upgrade installer failed".to_string());
        }
    } else {
        let cmd = "curl -fsSL https://raw.githubusercontent.com/orionlabz/openkit/main/scripts/install.sh | bash";
        if args.dry_run {
            println!("Dry run: {}", cmd);
            return Ok(());
        }
        let status = Command::new("sh")
            .args(["-c", cmd])
            .status()
            .map_err(|e| format!("failed to execute upgrade installer: {}", e))?;
        if !status.success() {
            return Err("upgrade installer failed".to_string());
        }
    }

    println!("Upgrade completed. Run `openkit --version` to verify.");
    Ok(())
}

fn fetch_latest_tag(repo: &str) -> Result<String, String> {
    let url = format!("https://api.github.com/repos/{}/releases/latest", repo);
    let output = Command::new("curl")
        .args([
            "-fsSL",
            "-H",
            "Accept: application/vnd.github+json",
            "-H",
            "User-Agent: openkit-rust-cli",
            &url,
        ])
        .output()
        .map_err(|e| format!("failed to execute curl: {}", e))?;
    if !output.status.success() {
        return Err("failed to fetch latest release data".to_string());
    }

    let body = String::from_utf8_lossy(&output.stdout);
    for line in body.lines() {
        if line.contains("\"tag_name\"") {
            let value = line
                .split(':')
                .nth(1)
                .unwrap_or("")
                .trim()
                .trim_matches(',')
                .trim_matches('"')
                .to_string();
            if !value.is_empty() {
                return Ok(value);
            }
        }
    }

    Err("tag_name not found in release response".to_string())
}

fn run_agent_command(agent: &str, cmd: AgentCommand) -> Result<(), String> {
    match cmd.command {
        AgentSubcommand::Sync(args) => run_agent_sync(agent, args),
        AgentSubcommand::Upgrade(args) => {
            let mut upgrade_args = args;
            upgrade_args.overwrite = true;
            run_agent_sync(agent, upgrade_args)
        }
        AgentSubcommand::Doctor(args) => run_agent_doctor(agent, args),
    }
}

fn run_agent_sync(agent: &str, args: AgentSyncArgs) -> Result<(), String> {
    let project = std::env::current_dir()
        .map_err(|e| format!("failed to detect current directory: {}", e))?;
    run_agent_sync_at(&project, agent, args)
}

fn run_agent_sync_at(project: &Path, agent: &str, args: AgentSyncArgs) -> Result<(), String> {
    let target = project.join(agent_dir(agent));
    let marker = target.join("OPENKIT.md");

    if args.dry_run {
        println!("Plan: create/update {}", marker.display());
        return Ok(());
    }

    fs::create_dir_all(&target)
        .map_err(|e| format!("failed to create {}: {}", target.display(), e))?;

    let content = format!(
        "# OpenKit Agent Config\n\nAgent: {}\n\nGenerated by Rust OpenKit runtime.\n",
        agent
    );
    write_text(&marker, &content, args.overwrite)?;

    println!("Synced agent configuration for {}", agent);
    println!("Config: {}", target.display());
    if args.prune {
        println!("Note: --prune acknowledged (no-op in baseline Rust parity implementation)");
    }
    Ok(())
}

fn run_agent_doctor(agent: &str, args: AgentDoctorArgs) -> Result<(), String> {
    let project = std::env::current_dir()
        .map_err(|e| format!("failed to detect current directory: {}", e))?;
    let config_dir = project.join(agent_dir(agent));
    let marker = config_dir.join("OPENKIT.md");

    let mut notes = Vec::new();
    let status = if config_dir.exists() && marker.exists() {
        notes.push("configuration directory and marker file found".to_string());
        "healthy"
    } else {
        notes.push("configuration missing; run sync".to_string());
        "missing"
    }
    .to_string();

    let report = AgentDoctorReport {
        agent: agent.to_string(),
        config_dir: config_dir.display().to_string(),
        status,
        notes,
    };

    if args.json {
        let payload = serde_json::to_string_pretty(&report)
            .map_err(|e| format!("failed to serialize doctor report: {}", e))?;
        println!("{}", payload);
    } else {
        println!("Agent doctor: {}", report.agent);
        println!("Status: {}", report.status);
        println!("Path: {}", report.config_dir);
        for note in report.notes {
            println!("- {}", note);
        }
    }

    Ok(())
}

fn agent_dir(agent: &str) -> &'static str {
    match agent {
        "opencode" => ".opencode",
        "claude" => ".claude",
        "cursor" => ".cursor",
        "gemini" => ".gemini",
        "codex" => ".codex",
        "antigravity" => ".antigravity",
        _ => ".agents",
    }
}

fn run_init(args: ProjectInitArgs) -> Result<(), String> {
    let cwd =
        std::env::current_dir().map_err(|e| format!("failed to get current directory: {}", e))?;

    let (project_dir, project_name) = if args.here {
        let name = cwd
            .file_name()
            .and_then(|n| n.to_str())
            .unwrap_or("project")
            .to_string();
        (cwd.clone(), name)
    } else {
        let name = args.project_name.clone().ok_or_else(|| {
            "project name required. Use: openkit init <name> or openkit init --here".to_string()
        })?;
        (cwd.join(&name), name)
    };

    if project_dir.exists() && !args.here && !args.force {
        return Err(format!(
            "directory '{}' already exists. Use --force to overwrite",
            project_name
        ));
    }

    fs::create_dir_all(&project_dir).map_err(|e| {
        format!(
            "failed to create project directory {}: {}",
            project_dir.display(),
            e
        )
    })?;

    let agent = args.ai.unwrap_or_else(|| "opencode".to_string());
    let supported_agents = [
        "opencode",
        "claude",
        "cursor",
        "gemini",
        "codex",
        "antigravity",
    ];
    if !supported_agents.contains(&agent.as_str()) {
        return Err(format!(
            "unknown agent '{}'. Supported: {}",
            agent,
            supported_agents.join(", ")
        ));
    }

    let docs_dir = project_dir.join("docs");
    let req_dir = docs_dir.join("requirements");
    let sprint_dir = docs_dir.join("sprint");

    fs::create_dir_all(&req_dir)
        .map_err(|e| format!("failed to create {}: {}", req_dir.display(), e))?;
    fs::create_dir_all(&sprint_dir)
        .map_err(|e| format!("failed to create {}: {}", sprint_dir.display(), e))?;

    let sprint01 = sprint_dir.join("Sprint-01");
    let req_bootstrap = req_dir.join("bootstrap");
    fs::create_dir_all(&sprint01)
        .map_err(|e| format!("failed to create {}: {}", sprint01.display(), e))?;
    fs::create_dir_all(&req_bootstrap)
        .map_err(|e| format!("failed to create {}: {}", req_bootstrap.display(), e))?;

    let files = vec![
        (
            project_dir.join("AGENTS.md"),
            "# Agents\n\nSee `docs/HUB-DOCS.md` for project context and workflow references.\n".to_string(),
        ),
        (
            docs_dir.join("HUB-DOCS.md"),
            "# Documentation Index\n\n## Context\n\nCentral hub for project documentation and discovery artifacts.\n\n## Navigation\n\n- [[CONTEXT.md]]\n- [[SECURITY.md]]\n- [[QUALITY_GATES.md]]\n- [[ACTION_ITEMS.md]]\n- [[API.md]]\n- [[GLOSSARY.md]]\n- [[requirements/HUB-REQUIREMENTS.md]]\n- [[sprint/HUB-SPRINTS.md]]\n\n## Related\n\n- [[CONTEXT.md]]\n- [[requirements/HUB-REQUIREMENTS.md]]\n- [[sprint/HUB-SPRINTS.md]]\n".to_string(),
        ),
        (
            docs_dir.join("CONTEXT.md"),
            format!("# CONTEXT\n\nProject initialized by OpenKit Rust runtime for agent `{}`.\n\n## Related\n\n- [[HUB-DOCS.md]]\n- [[SECURITY.md]]\n- [[QUALITY_GATES.md]]\n", agent),
        ),
        (
            docs_dir.join("SECURITY.md"),
            "# SECURITY\n\n## Baseline\n\n- Keep secrets out of VCS.\n- Enforce dependency and release checks.\n\n## Related\n\n- [[HUB-DOCS.md]]\n- [[QUALITY_GATES.md]]\n".to_string(),
        ),
        (
            docs_dir.join("QUALITY_GATES.md"),
            "# QUALITY_GATES\n\n- `cargo fmt --check`\n- `cargo clippy -- -D warnings`\n- `cargo test`\n\n## Related\n\n- [[HUB-DOCS.md]]\n- [[SECURITY.md]]\n".to_string(),
        ),
        (
            docs_dir.join("ACTION_ITEMS.md"),
            "# ACTION_ITEMS\n\nTrack cross-scope follow-ups and carry-over tasks.\n\n## Related\n\n- [[HUB-DOCS.md]]\n- [[sprint/HUB-SPRINTS.md]]\n".to_string(),
        ),
        (
            docs_dir.join("API.md"),
            "# API\n\n## Surface\n\n- `openkit --version`\n- `openkit check`\n- `openkit init`\n- `openkit <agent> sync|doctor|upgrade`\n- `openkit memory init|doctor|capture|review`\n\n## Related\n\n- [[HUB-DOCS.md]]\n- [[CONTEXT.md]]\n".to_string(),
        ),
        (
            docs_dir.join("GLOSSARY.md"),
            "# GLOSSARY\n\n- **Memory Kernel**: Docs-first persistent project memory model.\n- **Parity Matrix**: Command migration tracking from legacy runtime.\n\n## Related\n\n- [[HUB-DOCS.md]]\n".to_string(),
        ),
        (
            req_dir.join("HUB-REQUIREMENTS.md"),
            "# Requirements Hub\n\n## Features\n\n- [[requirements/bootstrap/HUB-BOOTSTRAP.md]]\n\n## Related\n\n- [[HUB-DOCS.md]]\n- [[sprint/HUB-SPRINTS.md]]\n".to_string(),
        ),
        (
            req_bootstrap.join("HUB-BOOTSTRAP.md"),
            "# Feature Hub: bootstrap\n\n- [[requirements/bootstrap/PROBLEM_STATEMENT.md]]\n- [[requirements/bootstrap/USER_STORIES.md]]\n- [[requirements/bootstrap/ACCEPTANCE_CRITERIA.md]]\n- [[requirements/bootstrap/DATA_CONTRACTS.md]]\n- [[requirements/bootstrap/RISKS.md]]\n- [[requirements/bootstrap/PLAN.md]]\n\n## Related\n\n- [[requirements/HUB-REQUIREMENTS.md]]\n- [[sprint/Sprint-01/HUB-SPRINT-01.md]]\n".to_string(),
        ),
        (
            req_bootstrap.join("PROBLEM_STATEMENT.md"),
            "# PROBLEM_STATEMENT\n\nDefine the baseline OpenKit project workflow and memory model.\n\n## Related\n\n- [[requirements/bootstrap/HUB-BOOTSTRAP.md]]\n".to_string(),
        ),
        (
            req_bootstrap.join("USER_STORIES.md"),
            "# USER_STORIES\n\n- As a developer, I want reliable command tooling and docs-driven context.\n\n## Related\n\n- [[requirements/bootstrap/HUB-BOOTSTRAP.md]]\n".to_string(),
        ),
        (
            req_bootstrap.join("ACCEPTANCE_CRITERIA.md"),
            "# ACCEPTANCE_CRITERIA\n\n- [ ] Core commands validated\n- [ ] Docs graph healthy\n\n## Related\n\n- [[requirements/bootstrap/HUB-BOOTSTRAP.md]]\n".to_string(),
        ),
        (
            req_bootstrap.join("DATA_CONTRACTS.md"),
            "# DATA_CONTRACTS\n\n- `.openkit/memory/config.yaml`\n- `.openkit/ops/queue.yaml`\n\n## Related\n\n- [[requirements/bootstrap/HUB-BOOTSTRAP.md]]\n".to_string(),
        ),
        (
            req_bootstrap.join("RISKS.md"),
            "# RISKS\n\n- Runtime parity gaps\n- Documentation drift\n\n## Related\n\n- [[requirements/bootstrap/HUB-BOOTSTRAP.md]]\n".to_string(),
        ),
        (
            req_bootstrap.join("PLAN.md"),
            "# PLAN\n\n1. Build core command parity\n2. Validate docs and release gates\n\n## Related\n\n- [[requirements/bootstrap/HUB-BOOTSTRAP.md]]\n- [[sprint/Sprint-01/TASKS.md]]\n".to_string(),
        ),
        (
            sprint_dir.join("HUB-SPRINTS.md"),
            "# Sprint Hub\n\n- [[sprint/Sprint-01/HUB-SPRINT-01.md]]\n\n## Related\n\n- [[HUB-DOCS.md]]\n- [[requirements/HUB-REQUIREMENTS.md]]\n".to_string(),
        ),
        (
            sprint01.join("HUB-SPRINT-01.md"),
            "# Sprint-01 Hub\n\n- [[sprint/Sprint-01/SPRINT_GOAL.md]]\n- [[sprint/Sprint-01/BACKLOG.md]]\n- [[sprint/Sprint-01/TASKS.md]]\n- [[sprint/Sprint-01/RISK_REGISTER.md]]\n\n## Related\n\n- [[sprint/HUB-SPRINTS.md]]\n- [[requirements/bootstrap/HUB-BOOTSTRAP.md]]\n".to_string(),
        ),
        (
            sprint01.join("SPRINT_GOAL.md"),
            "# SPRINT_GOAL\n\nDeliver baseline command and docs workflow.\n\n## Related\n\n- [[sprint/Sprint-01/TASKS.md]]\n".to_string(),
        ),
        (
            sprint01.join("BACKLOG.md"),
            "# BACKLOG\n\n- Bootstrap docs and command baseline\n\n## Related\n\n- [[sprint/Sprint-01/TASKS.md]]\n".to_string(),
        ),
        (
            sprint01.join("TASKS.md"),
            "# TASKS\n\n- [ ] Validate `openkit check`\n- [ ] Validate `openkit memory doctor`\n\n## Related\n\n- [[sprint/Sprint-01/HUB-SPRINT-01.md]]\n".to_string(),
        ),
        (
            sprint01.join("RISK_REGISTER.md"),
            "# RISK_REGISTER\n\n- Command parity drift\n\n## Related\n\n- [[sprint/Sprint-01/HUB-SPRINT-01.md]]\n".to_string(),
        ),
    ];

    for (path, content) in files {
        write_text(&path, &content, args.force)?;
    }

    run_agent_sync_at(
        &project_dir,
        &agent,
        AgentSyncArgs {
            dry_run: false,
            overwrite: true,
            prune: false,
        },
    )?;

    memory_init(MemoryInitArgs {
        project: Some(project_dir.clone()),
        force: true,
    })?;

    if !args.no_git {
        let git_dir = project_dir.join(".git");
        if !git_dir.exists() {
            let output = Command::new("git")
                .arg("init")
                .current_dir(&project_dir)
                .output();
            if let Err(err) = output {
                eprintln!("warning: git initialization failed: {}", err);
            }
        }
    }

    println!("\nProject initialized successfully!");
    println!("  Project: {}", project_name);
    println!("  Agent: {}", agent);
    if !args.here {
        println!("\nNext steps:");
        println!("  cd {}", project_name);
    } else {
        println!("\nNext steps:");
    }
    println!("  openkit check");
    println!("  openkit memory doctor --json");
    Ok(())
}

fn run_check(args: CheckArgs) -> Result<(), String> {
    let agents = vec![
        ("OpenCode", "opencode", vec!["--version"]),
        ("Claude Code", "claude", vec!["--version"]),
        ("Cursor", "cursor", vec!["--version"]),
        ("Gemini CLI", "gemini", vec!["--version"]),
        ("Codex CLI", "codex", vec!["--version"]),
    ];

    let tools = vec![
        ("Git", "git", vec!["--version"]),
        ("Node.js", "node", vec!["--version"]),
        ("Python", "python3", vec!["--version"]),
        ("Go", "go", vec!["version"]),
    ];

    let agent_results: Vec<CheckEntry> = agents
        .into_iter()
        .map(|(name, cmd, args)| check_command(name, cmd, args))
        .collect();
    let tool_results: Vec<CheckEntry> = tools
        .into_iter()
        .map(|(name, cmd, args)| check_command(name, cmd, args))
        .collect();

    let ready = agent_results.iter().any(|entry| entry.status == "ok");
    let report = CheckReport {
        platform: format!("{}/{}", std::env::consts::OS, std::env::consts::ARCH),
        agents: agent_results,
        tools: tool_results,
        ready,
    };

    if args.json {
        let payload = serde_json::to_string_pretty(&report)
            .map_err(|e| format!("failed to serialize check report: {}", e))?;
        println!("{}", payload);
        return Ok(());
    }

    println!("\nSystem Check");
    println!("============\n");
    println!("  Platform: {}\n", report.platform);

    println!("AI Agents");
    println!("---------");
    for entry in &report.agents {
        print_check_entry(entry);
    }

    println!("\nDevelopment Tools");
    println!("-----------------");
    for entry in &report.tools {
        print_check_entry(entry);
    }

    println!();
    if report.ready {
        println!("  Ready to use OpenKit!");
    } else {
        println!("  No AI agents detected. Install one of the supported agents.");
    }
    println!();

    Ok(())
}

fn check_command(name: &str, command: &str, args: Vec<&str>) -> CheckEntry {
    let output = Command::new(command).args(args).output();
    match output {
        Err(err) if err.kind() == std::io::ErrorKind::NotFound => CheckEntry {
            name: name.to_string(),
            command: command.to_string(),
            status: "missing".to_string(),
            version: None,
            note: Some("not found".to_string()),
        },
        Ok(out) if out.status.success() => {
            let version = String::from_utf8_lossy(&out.stdout)
                .lines()
                .next()
                .unwrap_or("")
                .chars()
                .take(80)
                .collect::<String>();
            CheckEntry {
                name: name.to_string(),
                command: command.to_string(),
                status: "ok".to_string(),
                version: Some(version),
                note: None,
            }
        }
        _ => CheckEntry {
            name: name.to_string(),
            command: command.to_string(),
            status: "error".to_string(),
            version: None,
            note: Some("error running command".to_string()),
        },
    }
}

fn print_check_entry(entry: &CheckEntry) {
    let marker = match entry.status.as_str() {
        "ok" => "[OK]",
        "missing" => "[--]",
        _ => "[!!]",
    };
    let message = if let Some(version) = &entry.version {
        version.as_str()
    } else {
        entry.note.as_deref().unwrap_or("")
    };
    println!("  {} {:<15} {}", marker, entry.name, message);
}

fn memory_init(args: MemoryInitArgs) -> Result<(), String> {
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

    write_yaml(&memory_dir.join("config.yaml"), &config, args.force)?;
    write_yaml(&memory_dir.join("derivation.yaml"), &derivation, args.force)?;
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
    let session_id = args.session_id.unwrap_or_else(|| format!("mk-{}", now));
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
        recommendations
            .push("Resolve repeated tensions before next implementation phase".to_string());
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
                broken.push(format!("{} -> [[{}]]", entry.path().display(), link));
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

fn write_text(path: &Path, content: &str, force: bool) -> Result<(), String> {
    if path.exists() && !force {
        return Ok(());
    }
    if let Some(parent) = path.parent() {
        fs::create_dir_all(parent)
            .map_err(|e| format!("failed to create {}: {}", parent.display(), e))?;
    }
    fs::write(path, content).map_err(|e| format!("failed writing {}: {}", path.display(), e))
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
