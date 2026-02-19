use clap::{Args, Parser, Subcommand};
use flate2::read::GzDecoder;
use include_dir::{include_dir, Dir};
use regex::Regex;
use serde::{Deserialize, Serialize};
use sha2::{Digest, Sha256};
use std::collections::{BTreeMap, HashSet};
use std::fs;
use std::io::{self, Write};
use std::path::{Path, PathBuf};
use std::process::Command;
use std::time::{SystemTime, UNIX_EPOCH};
use tar::Archive;
use walkdir::WalkDir;

static TEMPLATE_ROOT: Dir<'_> = include_dir!("$CARGO_MANIFEST_DIR/../internal/templates");

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
    Uninstall(UninstallArgs),
    #[command(about = "Memory Kernel maintenance and repair commands")]
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
struct UninstallArgs {
    #[arg(long)]
    dry_run: bool,
    #[arg(long)]
    yes: bool,
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
    #[arg(hide = true)]
    ai: Option<String>,
    #[arg(long)]
    claude: bool,
    #[arg(long)]
    codex: bool,
    #[arg(long)]
    cursor: bool,
    #[arg(long)]
    gemini: bool,
    #[arg(long)]
    antigravity: bool,
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
        Commands::Uninstall(args) => run_uninstall(args),
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

fn run_uninstall(args: UninstallArgs) -> Result<(), String> {
    let binary_name = if std::env::consts::OS == "windows" {
        "openkit.exe"
    } else {
        "openkit"
    };

    let mut targets = HashSet::new();

    if let Ok(current) = std::env::current_exe() {
        targets.insert(current);
    }

    if let Some(home) = home_dir() {
        let openkit_home = std::env::var_os("OPENKIT_HOME")
            .map(PathBuf::from)
            .unwrap_or_else(|| home.join(".openkit"));
        let openkit_bin = std::env::var_os("OPENKIT_INSTALL_DIR")
            .map(PathBuf::from)
            .unwrap_or_else(|| openkit_home.join("bin"));
        targets.insert(openkit_bin.join(binary_name));

        if let Some(xdg_bin) = std::env::var_os("XDG_BIN_HOME") {
            targets.insert(PathBuf::from(xdg_bin).join(binary_name));
        }
        targets.insert(home.join(".local/bin").join(binary_name));
        targets.insert(home.join("bin").join(binary_name));
    }

    if std::env::consts::OS != "windows" {
        targets.insert(PathBuf::from("/usr/local/bin").join(binary_name));
    }

    let mut existing: Vec<PathBuf> = targets.into_iter().filter(|path| path.exists()).collect();
    existing.sort();

    if existing.is_empty() {
        println!("OpenKit binary not found in known install paths.");
        return Ok(());
    }

    if args.dry_run {
        println!("Dry run: would remove the following paths:");
        for path in existing {
            println!("- {}", path.display());
        }
        return Ok(());
    }

    if !args.yes {
        println!("OpenKit uninstall will remove:");
        for path in &existing {
            println!("- {}", path.display());
        }
        print!("Continue? [y/N] ");
        io::stdout()
            .flush()
            .map_err(|e| format!("failed to flush stdout: {}", e))?;

        let mut input = String::new();
        io::stdin()
            .read_line(&mut input)
            .map_err(|e| format!("failed to read confirmation: {}", e))?;
        let confirmed = matches!(input.trim().to_ascii_lowercase().as_str(), "y" | "yes");
        if !confirmed {
            println!("Uninstall cancelled.");
            return Ok(());
        }
    }

    let mut removed = 0usize;
    for path in existing {
        match fs::remove_file(&path) {
            Ok(_) => {
                removed += 1;
                println!("Removed {}", path.display());
            }
            Err(e) if e.kind() == std::io::ErrorKind::NotFound => {}
            Err(e) if e.kind() == std::io::ErrorKind::PermissionDenied => {
                return Err(format!(
                    "permission denied removing {}. Re-run with elevated permissions",
                    path.display()
                ));
            }
            Err(e) => {
                return Err(format!("failed to remove {}: {}", path.display(), e));
            }
        }
    }

    println!("OpenKit uninstall complete. Removed {} file(s).", removed);
    println!("Project-level files (.openkit/, .opencode/, openkit-memory/) were not modified.");
    Ok(())
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

    if args.dry_run {
        let asset = release_asset_name(std::env::consts::OS, std::env::consts::ARCH)?;
        println!(
            "Dry run: would self-update from latest release at {}/{}",
            repo, asset
        );
        return Ok(());
    }

    if std::env::consts::OS == "windows" {
        let cmd = "irm https://raw.githubusercontent.com/orionlabz/openkit/main/scripts/install.ps1 | iex";
        let status = Command::new("powershell")
            .args(["-NoProfile", "-Command", cmd])
            .status()
            .map_err(|e| format!("failed to execute upgrade installer: {}", e))?;
        if !status.success() {
            return Err("upgrade installer failed".to_string());
        }
        println!("Upgrade completed. Run `openkit --version` to verify.");
        return Ok(());
    }

    run_self_update_unix(&repo, current)
}

fn run_self_update_unix(repo: &str, current: &str) -> Result<(), String> {
    let release = fetch_latest_release(repo)?;
    let latest = release.tag_name.trim_start_matches('v').to_string();
    if latest == current {
        println!("OpenKit is already up to date.");
        return Ok(());
    }

    let asset_name = release_asset_name(std::env::consts::OS, std::env::consts::ARCH)?;
    let asset = release
        .assets
        .iter()
        .find(|a| a.name == asset_name)
        .ok_or_else(|| format!("release asset not found: {}", asset_name))?;

    let checksums = release.assets.iter().find(|a| a.name == "checksums.txt");
    let tmp_dir = std::env::temp_dir().join(format!("openkit-upgrade-{}", timestamp_secs()));
    fs::create_dir_all(&tmp_dir)
        .map_err(|e| format!("failed to create temp dir {}: {}", tmp_dir.display(), e))?;

    let archive_path = tmp_dir.join(&asset_name);
    download_file(&asset.browser_download_url, &archive_path)?;

    if let Some(checksum_asset) = checksums {
        let checksums_path = tmp_dir.join("checksums.txt");
        download_file(&checksum_asset.browser_download_url, &checksums_path)?;
        verify_checksum(&archive_path, &checksums_path, &asset_name)?;
    }

    let extracted = extract_tar_gz_binary(&archive_path, "openkit", &tmp_dir)?;
    replace_current_binary(&extracted)?;

    let _ = fs::remove_dir_all(&tmp_dir);
    println!("Updated OpenKit from {} to {}.", current, latest);
    println!("Run `openkit --version` to verify.");
    Ok(())
}

fn download_file(url: &str, destination: &Path) -> Result<(), String> {
    let status = Command::new("curl")
        .args(["-fsSL", url, "-o"])
        .arg(destination)
        .status()
        .map_err(|e| format!("failed to execute curl: {}", e))?;
    if !status.success() {
        return Err(format!("download failed: {}", url));
    }
    Ok(())
}

fn verify_checksum(archive: &Path, checksums: &Path, asset_name: &str) -> Result<(), String> {
    let expected = checksum_for_asset(checksums, asset_name)?;
    let data = fs::read(archive)
        .map_err(|e| format!("failed to read archive {}: {}", archive.display(), e))?;
    let digest = Sha256::digest(&data);
    let actual = format!("{:x}", digest);
    if actual != expected {
        return Err(format!(
            "checksum mismatch for {}: expected {}, got {}",
            asset_name, expected, actual
        ));
    }
    Ok(())
}

fn checksum_for_asset(checksums_path: &Path, asset_name: &str) -> Result<String, String> {
    let content = fs::read_to_string(checksums_path)
        .map_err(|e| format!("failed to read {}: {}", checksums_path.display(), e))?;
    for line in content.lines() {
        let trimmed = line.trim();
        if trimmed.is_empty() {
            continue;
        }
        let mut parts = trimmed.split_whitespace();
        let hash = parts.next().unwrap_or_default();
        let file = parts.next().unwrap_or_default().trim_start_matches('*');
        if file == asset_name {
            return Ok(hash.to_lowercase());
        }
    }
    Err(format!("checksum entry not found for {}", asset_name))
}

fn extract_tar_gz_binary(
    archive_path: &Path,
    binary_name: &str,
    out_dir: &Path,
) -> Result<PathBuf, String> {
    let file = fs::File::open(archive_path)
        .map_err(|e| format!("failed to open archive {}: {}", archive_path.display(), e))?;
    let reader = GzDecoder::new(file);
    let mut archive = Archive::new(reader);
    let target = out_dir.join(format!("{}-new", binary_name));
    let mut extracted = false;

    let entries = archive
        .entries()
        .map_err(|e| format!("failed to read archive entries: {}", e))?;
    for entry in entries {
        let mut e = entry.map_err(|err| format!("failed reading archive entry: {}", err))?;
        let path = e
            .path()
            .map_err(|err| format!("failed to read archive entry path: {}", err))?;
        if path.file_name().and_then(|s| s.to_str()) == Some(binary_name) {
            e.unpack(&target)
                .map_err(|err| format!("failed to unpack {}: {}", binary_name, err))?;
            extracted = true;
            break;
        }
    }

    if !extracted {
        return Err(format!("binary {} not found in archive", binary_name));
    }

    #[cfg(unix)]
    {
        use std::os::unix::fs::PermissionsExt;
        let mut perms = fs::metadata(&target)
            .map_err(|e| format!("failed to read {} metadata: {}", target.display(), e))?
            .permissions();
        perms.set_mode(0o755);
        fs::set_permissions(&target, perms)
            .map_err(|e| format!("failed to set permissions on {}: {}", target.display(), e))?;
    }

    Ok(target)
}

fn replace_current_binary(new_binary: &Path) -> Result<(), String> {
    let current = std::env::current_exe()
        .map_err(|e| format!("failed to resolve current executable: {}", e))?;
    let target = fs::canonicalize(&current).unwrap_or(current);
    let parent = target
        .parent()
        .ok_or_else(|| format!("failed to resolve parent for {}", target.display()))?;

    let staged = parent.join("openkit.staged");
    let backup = parent.join("openkit.backup");

    fs::copy(new_binary, &staged)
        .map_err(|e| format!("failed to stage binary {}: {}", staged.display(), e))?;

    if backup.exists() {
        let _ = fs::remove_file(&backup);
    }

    fs::rename(&target, &backup).map_err(|e| {
        format!(
            "failed to move current binary {} to backup: {}",
            target.display(),
            e
        )
    })?;

    if let Err(e) = fs::rename(&staged, &target) {
        let _ = fs::rename(&backup, &target);
        return Err(format!(
            "failed to replace binary {}: {}",
            target.display(),
            e
        ));
    }

    let _ = fs::remove_file(&backup);
    Ok(())
}

#[derive(Deserialize)]
struct ReleaseAsset {
    name: String,
    browser_download_url: String,
}

#[derive(Deserialize)]
struct LatestRelease {
    tag_name: String,
    assets: Vec<ReleaseAsset>,
}

fn fetch_latest_release(repo: &str) -> Result<LatestRelease, String> {
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
    serde_json::from_str(&body).map_err(|e| format!("failed to parse release JSON: {}", e))
}

fn release_asset_name(os: &str, arch: &str) -> Result<String, String> {
    let os_name = match os {
        "macos" => "Darwin",
        "linux" => "Linux",
        "windows" => "Windows",
        other => return Err(format!("unsupported OS for upgrade: {}", other)),
    };

    let arch_name = match arch {
        "x86_64" => "x86_64",
        "aarch64" | "arm64" => "arm64",
        other => return Err(format!("unsupported architecture for upgrade: {}", other)),
    };

    let ext = if os_name == "Windows" {
        "zip"
    } else {
        "tar.gz"
    };
    Ok(format!("openkit_{}_{}.{}", os_name, arch_name, ext))
}

fn fetch_latest_tag(repo: &str) -> Result<String, String> {
    let release = fetch_latest_release(repo)?;
    Ok(release.tag_name)
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

    let base = TEMPLATE_ROOT
        .get_dir("base")
        .ok_or_else(|| "missing embedded templates: internal/templates/base".to_string())?;
    copy_embedded_dir(base, &target, args.overwrite)?;

    let memory_rule = TEMPLATE_ROOT
        .get_file("memory/rules/MEMORY_KERNEL.md")
        .ok_or_else(|| {
            "missing embedded template: internal/templates/memory/rules/MEMORY_KERNEL.md"
                .to_string()
        })?;
    write_embedded_file(
        memory_rule,
        &target.join("rules/MEMORY_KERNEL.md"),
        args.overwrite,
    )?;

    if agent == "opencode" {
        let root_config = TEMPLATE_ROOT
            .get_file("root/opencode.json")
            .ok_or_else(|| {
                "missing embedded template: internal/templates/root/opencode.json".to_string()
            })?;
        write_embedded_file(root_config, &project.join("opencode.json"), args.overwrite)?;
    }

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

    let agent = resolve_init_agent(&args)?;

    let docs_dir = project_dir.join("openkit-memory");
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
            "# Agents\n\nSee `openkit-memory/HUB-DOCS.md` for project context and workflow references.\n".to_string(),
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
            "# API\n\n## Surface\n\n- `openkit --version`\n- `openkit check`\n- `openkit init`\n- `openkit upgrade`\n- `openkit uninstall`\n- `openkit <agent> sync|doctor|upgrade`\n- `openkit memory doctor|capture|review` (maintenance)\n- `openkit memory init` (repair only)\n\n## Related\n\n- [[HUB-DOCS.md]]\n- [[CONTEXT.md]]\n".to_string(),
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
    println!("  Memory Kernel: active (initialized by default)");
    if !args.here {
        println!("\nNext steps:");
        println!("  cd {}", project_name);
    } else {
        println!("\nNext steps:");
    }
    println!("  openkit check");
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

    let config_template = TEMPLATE_ROOT
        .get_file("memory/config.yaml")
        .ok_or_else(|| {
            "missing embedded template: internal/templates/memory/config.yaml".to_string()
        })?;
    let derivation_template = TEMPLATE_ROOT
        .get_file("memory/derivation.yaml")
        .ok_or_else(|| {
            "missing embedded template: internal/templates/memory/derivation.yaml".to_string()
        })?;
    let queue_template = TEMPLATE_ROOT.get_file("memory/queue.yaml").ok_or_else(|| {
        "missing embedded template: internal/templates/memory/queue.yaml".to_string()
    })?;

    write_embedded_file(config_template, &memory_dir.join("config.yaml"), args.force)?;
    write_embedded_file(
        derivation_template,
        &memory_dir.join("derivation.yaml"),
        args.force,
    )?;
    write_embedded_file(queue_template, &ops_dir.join("queue.yaml"), args.force)?;

    println!("Initialized Memory Kernel structure at {}", root.display());
    Ok(())
}

fn memory_doctor(args: DoctorArgs) -> Result<(), String> {
    let root = project_root(args.project)?;
    let docs = if root.join("openkit-memory").exists() {
        root.join("openkit-memory")
    } else {
        root.join("docs")
    };
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
            let normalized = link
                .trim_start_matches("openkit-memory/")
                .trim_start_matches("docs/");
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

fn write_embedded_file(
    file: &include_dir::File<'_>,
    path: &Path,
    force: bool,
) -> Result<(), String> {
    if path.exists() && !force {
        return Ok(());
    }
    if let Some(parent) = path.parent() {
        fs::create_dir_all(parent)
            .map_err(|e| format!("failed to create {}: {}", parent.display(), e))?;
    }
    fs::write(path, file.contents())
        .map_err(|e| format!("failed writing {}: {}", path.display(), e))
}

fn copy_embedded_dir(dir: &Dir<'_>, destination: &Path, force: bool) -> Result<(), String> {
    copy_embedded_dir_with_prefix(dir, destination, force, dir.path())
}

fn copy_embedded_dir_with_prefix(
    dir: &Dir<'_>,
    destination: &Path,
    force: bool,
    root_prefix: &Path,
) -> Result<(), String> {
    for file in dir.files() {
        let rel = file.path().strip_prefix(root_prefix).unwrap_or(file.path());
        let target = destination.join(rel);
        write_embedded_file(file, &target, force)?;
    }
    for child in dir.dirs() {
        copy_embedded_dir_with_prefix(child, destination, force, root_prefix)?;
    }
    Ok(())
}

fn resolve_init_agent(args: &ProjectInitArgs) -> Result<String, String> {
    let mut selected = Vec::new();
    if args.claude {
        selected.push("claude");
    }
    if args.codex {
        selected.push("codex");
    }
    if args.cursor {
        selected.push("cursor");
    }
    if args.gemini {
        selected.push("gemini");
    }
    if args.antigravity {
        selected.push("antigravity");
    }

    if let Some(ai) = &args.ai {
        if !selected.is_empty() {
            return Err("use either --ai or one agent flag (--claude/--codex/...)".to_string());
        }
        let supported = [
            "opencode",
            "claude",
            "cursor",
            "gemini",
            "codex",
            "antigravity",
        ];
        if !supported.contains(&ai.as_str()) {
            return Err(format!(
                "unknown agent '{}'. Supported: {}",
                ai,
                supported.join(", ")
            ));
        }
        return Ok(ai.clone());
    }

    if selected.len() > 1 {
        return Err(
            "use only one agent flag at a time (--claude, --codex, --cursor, --gemini, --antigravity)"
                .to_string(),
        );
    }

    Ok(selected.first().copied().unwrap_or("opencode").to_string())
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

fn home_dir() -> Option<PathBuf> {
    std::env::var_os("HOME").map(PathBuf::from)
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn release_asset_name_matches_platform_matrix() {
        assert_eq!(
            release_asset_name("macos", "aarch64").expect("asset name"),
            "openkit_Darwin_arm64.tar.gz"
        );
        assert_eq!(
            release_asset_name("linux", "x86_64").expect("asset name"),
            "openkit_Linux_x86_64.tar.gz"
        );
        assert_eq!(
            release_asset_name("windows", "x86_64").expect("asset name"),
            "openkit_Windows_x86_64.zip"
        );
    }

    #[test]
    fn checksum_for_asset_reads_sha256sum_format() {
        let file = std::env::temp_dir().join(format!("openkit-checksums-{}.txt", timestamp_secs()));
        let payload = "abc123  openkit_Darwin_arm64.tar.gz\ndef456 *openkit_Linux_x86_64.tar.gz\n";
        fs::write(&file, payload).expect("write checksums fixture");

        let hash = checksum_for_asset(&file, "openkit_Linux_x86_64.tar.gz").expect("parse hash");
        assert_eq!(hash, "def456");

        let missing = checksum_for_asset(&file, "openkit_Windows_x86_64.zip");
        assert!(missing.is_err());

        let _ = fs::remove_file(&file);
    }

    #[test]
    fn embedded_templates_are_available() {
        assert!(TEMPLATE_ROOT
            .get_file("base/commands/discover.md")
            .is_some());
        assert!(TEMPLATE_ROOT.get_file("memory/config.yaml").is_some());
    }

    #[test]
    fn copy_embedded_dir_materializes_expected_files() {
        let temp = tempfile::tempdir().expect("tempdir");
        let base = TEMPLATE_ROOT.get_dir("base").expect("base templates");
        copy_embedded_dir(base, temp.path(), true).expect("copy embedded dir");
        assert!(temp.path().join("commands/discover.md").exists());
    }
}
