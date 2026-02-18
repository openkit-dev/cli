#!/bin/bash
set -e

# OpenKit CLI Installer for macOS/Linux/WSL
# Usage: curl -fsSL https://openkit.dev/install | bash

REPO="openkit-devtools/openkit"
BINARY_NAME="openkit"
RUNTIME_NAME="openkit-rs"

OPENKIT_HOME_DEFAULT="$HOME/.openkit"
OPENKIT_HOME="${OPENKIT_HOME:-$OPENKIT_HOME_DEFAULT}"
OPENKIT_BIN_DIR_DEFAULT="$OPENKIT_HOME/bin"
OPENKIT_BIN_DIR="${OPENKIT_INSTALL_DIR:-$OPENKIT_BIN_DIR_DEFAULT}"

is_in_path() {
  case ":$PATH:" in
    *":$1:"*) return 0 ;;
    *) return 1 ;;
  esac
}

ensure_dir() {
  mkdir -p "$1"
}

try_symlink_into_path() {
  src="$1"

  # Prefer a user-writable bin already in PATH.
  candidates=""
  if [ -n "${XDG_BIN_HOME:-}" ]; then
    candidates="$candidates $XDG_BIN_HOME"
  fi
  candidates="$candidates $HOME/.local/bin"
  candidates="$candidates $HOME/bin"

  for d in $candidates; do
    if is_in_path "$d"; then
      ensure_dir "$d"
      if [ -w "$d" ]; then
        ln -sf "$src" "$d/$BINARY_NAME"
        return 0
      fi
    fi
  done

  return 1
}

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
CYAN='\033[0;36m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${CYAN}"
echo "   ___                   _  ___ _   "
echo "  / _ \ _ __   ___ _ __ | |/ (_) |_ "
echo " | | | | '_ \ / _ \ '_ \| ' /| | __|"
echo " | |_| | |_) |  __/ | | | . \| | |_ "
echo "  \___/| .__/ \___|_| |_|_|\_\_|\__|"
echo "       |_|                          "
echo -e "${NC}"
echo "Universal Spec-Driven Development Toolkit"
echo ""

echo -e "${CYAN}Install location...${NC}"
echo "  OPENKIT_HOME: $OPENKIT_HOME"
echo "  BIN_DIR:      $OPENKIT_BIN_DIR"
echo ""

# Detect OS and architecture
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case "$OS" in
  linux)
    OS="Linux"
    ;;
  darwin)
    OS="Darwin"
    ;;
  *)
    echo -e "${RED}Unsupported OS: $OS${NC}"
    exit 1
    ;;
esac

case "$ARCH" in
  x86_64)
    ARCH="x86_64"
    ;;
  aarch64|arm64)
    ARCH="arm64"
    ;;
  *)
    echo -e "${RED}Unsupported architecture: $ARCH${NC}"
    exit 1
    ;;
esac

echo -e "${CYAN}Detecting system...${NC}"
echo "  OS: $OS"
echo "  Architecture: $ARCH"
echo ""

# Get latest release
echo -e "${CYAN}Fetching latest release...${NC}"
LATEST_RELEASE=$(curl -s "https://api.github.com/repos/$REPO/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')

if [ -z "$LATEST_RELEASE" ]; then
  echo -e "${RED}Failed to fetch latest release${NC}"
  exit 1
fi

echo "  Latest version: $LATEST_RELEASE"
echo ""

FILENAME="openkit_${OS}_${ARCH}.tar.gz"
if [ "$OS" = "windows" ]; then
  FILENAME="openkit_${OS}_${ARCH}.zip"
fi

DOWNLOAD_URL="https://github.com/$REPO/releases/download/$LATEST_RELEASE/$FILENAME"

echo -e "${CYAN}Downloading OpenKit CLI...${NC}"
echo "  URL: $DOWNLOAD_URL"

# Create temp directory
TMP_DIR=$(mktemp -d)
trap "rm -rf $TMP_DIR" EXIT

cd "$TMP_DIR"

# Download
if ! curl -fsSL "$DOWNLOAD_URL" -o "$FILENAME"; then
  echo -e "${RED}Failed to download $FILENAME${NC}"
  exit 1
fi

# Extract
echo -e "${CYAN}Extracting...${NC}"
tar -xzf "$FILENAME"

INSTALL_DIR="$OPENKIT_BIN_DIR"

# Install
echo -e "${CYAN}Installing to $INSTALL_DIR...${NC}"

ensure_dir "$INSTALL_DIR"
if [ ! -w "$INSTALL_DIR" ]; then
  echo -e "${RED}Install dir is not writable: $INSTALL_DIR${NC}"
  exit 1
fi

install -m 755 "$BINARY_NAME" "$INSTALL_DIR/$BINARY_NAME"

# Optional: install Rust runtime sidecar when available
RUNTIME_FILENAME="openkit-rs_${OS}_${ARCH}.tar.gz"
RUNTIME_URL="https://github.com/$REPO/releases/download/$LATEST_RELEASE/$RUNTIME_FILENAME"
echo -e "${CYAN}Checking memory runtime sidecar...${NC}"
if curl -fsSL "$RUNTIME_URL" -o "$RUNTIME_FILENAME"; then
  tar -xzf "$RUNTIME_FILENAME"
  if [ -f "$RUNTIME_NAME" ]; then
    install -m 755 "$RUNTIME_NAME" "$INSTALL_DIR/$RUNTIME_NAME"
    echo "  Installed runtime: $RUNTIME_NAME"
  else
    echo "  Runtime archive found, but binary missing. Skipping sidecar install."
  fi
else
  echo "  Runtime sidecar not available for $OS/$ARCH. Cargo fallback remains available."
fi

# Verify installation
if command -v "$BINARY_NAME" &> /dev/null; then
  echo ""
  echo -e "${GREEN}✓ OpenKit CLI installed successfully!${NC}"
  echo ""
  echo "Run 'openkit --help' to get started"
  echo ""
  echo "Examples:"
  echo "  openkit check              # Check system requirements"
  echo "  openkit init my-app        # Create new project"
  echo "  openkit init --ai claude   # Create project for Claude Code"
  echo ""
else
  # Try to make it available without requiring PATH edits.
  if try_symlink_into_path "$INSTALL_DIR/$BINARY_NAME" && command -v "$BINARY_NAME" &> /dev/null; then
    echo ""
    echo -e "${GREEN}✓ OpenKit CLI installed successfully!${NC}"
    echo ""
    echo "Run 'openkit --help' to get started"
    echo ""
    exit 0
  fi

  echo -e "${YELLOW}Installed, but '$BINARY_NAME' is not on PATH yet.${NC}"
  echo ""
  echo "Option A (recommended): add OpenKit bin dir to PATH"
  echo "  echo 'export PATH=\"$INSTALL_DIR:\$PATH\"' >> ~/.zshrc"
  echo "  source ~/.zshrc"
  echo ""
  echo "Option B: run it directly"
  echo "  $INSTALL_DIR/$BINARY_NAME --help"
  echo ""
  exit 0
fi
