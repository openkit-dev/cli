# OpenKit CLI Installer for Windows
# Usage: irm https://openkit.dev/install.ps1 | iex

$ErrorActionPreference = 'Stop'

$REPO = "openkit-devtools/openkit"
$BINARY_NAME = "openkit"
$RUNTIME_NAME = "openkit-rs"

$OPENKIT_HOME = if ($env:OPENKIT_HOME) { $env:OPENKIT_HOME } else { Join-Path $HOME ".openkit" }
$INSTALL_DIR = if ($env:OPENKIT_INSTALL_DIR) { $env:OPENKIT_INSTALL_DIR } else { Join-Path $OPENKIT_HOME "bin" }

Write-Host ""
Write-Host "   ___                   _  ___ _   " -ForegroundColor Cyan
Write-Host "  / _ \ _ __   ___ _ __ | |/ (_) |_ " -ForegroundColor Cyan
Write-Host " | | | | '_ \ / _ \ '_ \| ' /| | __|" -ForegroundColor Cyan
Write-Host " | |_| | |_) |  __/ | | | . \| | |_ " -ForegroundColor Cyan
Write-Host "  \___/| .__/ \___|_| |_|_|\_\_|\__|" -ForegroundColor Cyan
Write-Host "       |_|                          " -ForegroundColor Cyan
Write-Host ""
Write-Host "Universal Spec-Driven Development Toolkit" -ForegroundColor White
Write-Host ""

Write-Host "Install location..." -ForegroundColor Cyan
Write-Host "  OPENKIT_HOME: $OPENKIT_HOME"
Write-Host "  BIN_DIR:      $INSTALL_DIR"
Write-Host ""

# Detect architecture
$ARCH = if ([Environment]::Is64BitOperatingSystem) { "x86_64" } else { "i386" }

Write-Host "Detecting system..." -ForegroundColor Cyan
Write-Host "  OS: Windows"
Write-Host "  Architecture: $ARCH"
Write-Host ""

# Get latest release
Write-Host "Fetching latest release..." -ForegroundColor Cyan
try {
    $response = Invoke-RestMethod -Uri "https://api.github.com/repos/$REPO/releases/latest"
    $LATEST_RELEASE = $response.tag_name
    Write-Host "  Latest version: $LATEST_RELEASE"
    Write-Host ""
} catch {
    Write-Host "Failed to fetch latest release" -ForegroundColor Red
    exit 1
}

$FILENAME = "openkit_Windows_${ARCH}.zip"
$DOWNLOAD_URL = "https://github.com/$REPO/releases/download/$LATEST_RELEASE/$FILENAME"

Write-Host "Downloading OpenKit CLI..." -ForegroundColor Cyan
Write-Host "  URL: $DOWNLOAD_URL"

# Create temp directory
$TMP_DIR = New-Item -ItemType Directory -Path "$env:TEMP\openkit-install-$(Get-Random)"

try {
    # Download
    $ZIP_PATH = Join-Path $TMP_DIR $FILENAME
    Invoke-WebRequest -Uri $DOWNLOAD_URL -OutFile $ZIP_PATH -UseBasicParsing

    # Extract
    Write-Host "Extracting..." -ForegroundColor Cyan
    Expand-Archive -Path $ZIP_PATH -DestinationPath $TMP_DIR -Force

    # Create install directory
    if (-not (Test-Path $INSTALL_DIR)) {
        New-Item -ItemType Directory -Path $INSTALL_DIR -Force | Out-Null
    }

    # Install
    Write-Host "Installing to $INSTALL_DIR..." -ForegroundColor Cyan
    $BINARY_PATH = Join-Path $TMP_DIR "$BINARY_NAME.exe"
    $DEST_PATH = Join-Path $INSTALL_DIR "$BINARY_NAME.exe"
    Copy-Item -Path $BINARY_PATH -Destination $DEST_PATH -Force

    # Optional: install Rust runtime sidecar when available
    $RUNTIME_FILENAME = "openkit-rs_Windows_${ARCH}.zip"
    $RUNTIME_URL = "https://github.com/$REPO/releases/download/$LATEST_RELEASE/$RUNTIME_FILENAME"
    Write-Host "Checking memory runtime sidecar..." -ForegroundColor Cyan
    try {
        $RUNTIME_ZIP_PATH = Join-Path $TMP_DIR $RUNTIME_FILENAME
        Invoke-WebRequest -Uri $RUNTIME_URL -OutFile $RUNTIME_ZIP_PATH -UseBasicParsing
        $RUNTIME_TMP = Join-Path $TMP_DIR "runtime"
        Expand-Archive -Path $RUNTIME_ZIP_PATH -DestinationPath $RUNTIME_TMP -Force
        $RUNTIME_BIN = Join-Path $RUNTIME_TMP "$RUNTIME_NAME.exe"
        if (Test-Path $RUNTIME_BIN) {
            Copy-Item -Path $RUNTIME_BIN -Destination (Join-Path $INSTALL_DIR "$RUNTIME_NAME.exe") -Force
            Write-Host "  Installed runtime: $RUNTIME_NAME.exe" -ForegroundColor Gray
        } else {
            Write-Host "  Runtime archive found, but binary missing. Skipping sidecar install." -ForegroundColor Yellow
        }
    } catch {
        Write-Host "  Runtime sidecar not available for Windows/$ARCH. Cargo fallback remains available." -ForegroundColor Yellow
    }

    # Add to PATH
    $PATH = [Environment]::GetEnvironmentVariable("Path", "User")
    if ($PATH -notlike "*$INSTALL_DIR*") {
        Write-Host "Adding to PATH..." -ForegroundColor Cyan
        [Environment]::SetEnvironmentVariable("Path", "$PATH;$INSTALL_DIR", "User")
        $env:Path = "$env:Path;$INSTALL_DIR"
    }

    Write-Host ""
    Write-Host "✓ OpenKit CLI installed successfully!" -ForegroundColor Green
    Write-Host ""
    Write-Host "Run 'openkit --help' to get started" -ForegroundColor White
    Write-Host ""
    Write-Host "Examples:"
    Write-Host "  openkit check              # Check system requirements"
    Write-Host "  openkit init my-app        # Create new project"
    Write-Host "  openkit init --ai claude   # Create project for Claude Code"
    Write-Host ""
    Write-Host "Note: Restart your terminal to use OpenKit CLI" -ForegroundColor Yellow
    Write-Host ""

} finally {
    # Cleanup
    Remove-Item -Path $TMP_DIR -Recurse -Force -ErrorAction SilentlyContinue
}
