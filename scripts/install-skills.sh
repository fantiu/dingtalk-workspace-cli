#!/bin/sh
set -eu

# Install DWS agent skills from GitHub into detected agent directories.
# Usage:
#   curl -fsSL https://raw.githubusercontent.com/DingTalk-Real-AI/dingtalk-workspace-cli/main/scripts/install-skills.sh | sh
#
# The script downloads the dws-skills.zip release asset from GitHub Releases
# and copies it into every detected agent skills directory in the current
# project.

REPO="DingTalk-Real-AI/dingtalk-workspace-cli"
VERSION="${DWS_VERSION:-latest}"
SKILL_NAME="dws"

# ── Agent directory to install skills into ───────────────────────────────────
# Only install to .agents/skills — most agents can fall back to this directory.
AGENT_DIR=".agents/skills"

# ── Helpers ──────────────────────────────────────────────────────────────────

need_cmd() {
  if ! command -v "$1" >/dev/null 2>&1; then
    printf '❌ Missing required command: %s\n' "$1" >&2
    exit 1
  fi
}

resolve_version() {
  if [ "$VERSION" = "latest" ]; then
    VERSION="$(curl -fsSI "https://github.com/${REPO}/releases/latest" 2>/dev/null \
      | grep -i '^location:' | sed 's|.*/tag/||;s/[[:space:]]*$//')"
    if [ -z "$VERSION" ]; then
      printf '❌ Could not determine the latest version. Set DWS_VERSION explicitly.\n' >&2
      exit 1
    fi
  fi
}

extract_zip() {
  archive="$1"
  dest="$2"
  if command -v unzip >/dev/null 2>&1; then
    unzip -q "$archive" -d "$dest"
    return 0
  fi
  if command -v tar >/dev/null 2>&1 && tar -xf "$archive" -C "$dest" >/dev/null 2>&1; then
    return 0
  fi
  printf '❌ Missing required command: unzip (or tar with zip support)\n' >&2
  exit 1
}

# ── Main ─────────────────────────────────────────────────────────────────────

main() {
  need_cmd curl
  resolve_version

  CWD="$(pwd)"

  printf '\n'
  printf '  ┌──────────────────────────────────────┐\n'
  printf '  │     DWS Skill Installer              │\n'
  printf '  │     DingTalk Workspace CLI            │\n'
  printf '  └──────────────────────────────────────┘\n'
  printf '\n'

  # Download the tarball to a temp directory
  TMPDIR_WORK="$(mktemp -d)"
  trap 'rm -rf "$TMPDIR_WORK"' EXIT INT TERM

  ASSET_URL="https://github.com/${REPO}/releases/download/${VERSION}/dws-skills.zip"
  printf '  ⬇  Downloading skills from GitHub Releases: %s (%s)\n' "$REPO" "$VERSION"
  curl -fsSL "$ASSET_URL" -o "$TMPDIR_WORK/dws-skills.zip"
  extract_zip "$TMPDIR_WORK/dws-skills.zip" "$TMPDIR_WORK/extracted"

  SKILL_SRC="$TMPDIR_WORK/extracted"
  if [ -f "$TMPDIR_WORK/extracted/${SKILL_NAME}/SKILL.md" ]; then
    SKILL_SRC="$TMPDIR_WORK/extracted/${SKILL_NAME}"
  fi

  if [ ! -f "$SKILL_SRC/SKILL.md" ]; then
    printf '  ❌ Skill source not found in release asset\n' >&2
    exit 1
  fi

  # Install to .agents/skills only
  dest="$CWD/$AGENT_DIR/$SKILL_NAME"

  # Remove existing installation
  if [ -d "$dest" ]; then
    rm -rf "$dest"
  fi

  # Copy skill files
  mkdir -p "$dest"
  cp -R "$SKILL_SRC/"* "$dest/"
  file_count="$(find "$dest" -type f | wc -l | tr -d ' ')"

  printf '  ✅ Universal (.agents)\n'
  printf '     → %s/%s (%s files)\n' "$AGENT_DIR" "$SKILL_NAME" "$file_count"

  printf '\n'
  printf '  📖 Skill includes:\n'
  printf '     • SKILL.md — Main skill with product overview and intent routing\n'
  printf '     • references/ — Detailed product command references\n'
  printf '     • scripts/ — Batch operation scripts for all products\n'
  printf '\n'
  printf '  ⚡ Requires: dws CLI installed and on $PATH\n'
  printf '     Install: go install github.com/%s/cmd@latest\n' "$REPO"
  printf '\n'
}

main
