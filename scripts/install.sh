#!/usr/bin/env sh
set -eu

REPO_OWNER="RewriteToday"
REPO_NAME="cli"
BINARY_NAME="rewrite"
INSTALL_DIR="${INSTALL_DIR:-/usr/local/bin}"
VERSION="${VERSION:-latest}"

need_cmd() {
  command -v "$1" >/dev/null 2>&1 || {
    echo "missing required command: $1" >&2
    exit 1
  }
}

fetch_url() {
  url="$1"
  out="$2"

  if command -v curl >/dev/null 2>&1; then
    curl -fsSL "$url" -o "$out"
    return
  fi

  if command -v wget >/dev/null 2>&1; then
    wget -qO "$out" "$url"
    return
  fi

  echo "curl or wget is required" >&2
  exit 1
}

resolve_latest_version() {
  latest_url="https://github.com/${REPO_OWNER}/${REPO_NAME}/releases/latest"

  if command -v curl >/dev/null 2>&1; then
    final_url=$(curl -fsSLI -o /dev/null -w '%{url_effective}' "$latest_url")
  elif command -v wget >/dev/null 2>&1; then
    final_url=$(wget -qO- --max-redirect=0 "$latest_url" 2>&1 | awk '/^  Location: /{print $2}' | tr -d '\r')
  else
    echo "curl or wget is required" >&2
    exit 1
  fi

  tag=$(basename "$final_url")
  if [ -z "$tag" ]; then
    echo "failed to resolve latest release tag" >&2
    exit 1
  fi

  echo "$tag"
}

normalize_os() {
  os=$(uname -s | tr '[:upper:]' '[:lower:]')
  case "$os" in
    linux) echo "linux" ;;
    darwin) echo "darwin" ;;
    mingw*|msys*|cygwin*) echo "windows" ;;
    *)
      echo "unsupported OS: $os" >&2
      exit 1
      ;;
  esac
}

normalize_arch() {
  arch=$(uname -m)
  case "$arch" in
    x86_64|amd64) echo "amd64" ;;
    aarch64|arm64) echo "arm64" ;;
    *)
      echo "unsupported architecture: $arch" >&2
      exit 1
      ;;
  esac
}

install_binary() {
  bin_path="$1"
  chmod +x "$bin_path"

  if [ -w "$INSTALL_DIR" ]; then
    mv "$bin_path" "$INSTALL_DIR/$BINARY_NAME"
  else
    need_cmd sudo
    sudo mv "$bin_path" "$INSTALL_DIR/$BINARY_NAME"
  fi
}

main() {
  need_cmd uname
  need_cmd tar

  if [ "$VERSION" = "latest" ]; then
    VERSION=$(resolve_latest_version)
  fi

  os=$(normalize_os)
  arch=$(normalize_arch)

  asset_ext="tar.gz"
  if [ "$os" = "windows" ]; then
    need_cmd unzip
    asset_ext="zip"
  fi

  asset_name="${BINARY_NAME}_${VERSION}_${os}_${arch}.${asset_ext}"
  download_url="https://github.com/${REPO_OWNER}/${REPO_NAME}/releases/download/${VERSION}/${asset_name}"

  tmp_dir=$(mktemp -d)
  trap 'rm -rf "$tmp_dir"' EXIT INT TERM

  archive_path="$tmp_dir/$asset_name"
  fetch_url "$download_url" "$archive_path"

  if [ "$asset_ext" = "zip" ]; then
    unzip -q "$archive_path" -d "$tmp_dir"
  else
    tar -xzf "$archive_path" -C "$tmp_dir"
  fi

  bin_path="$tmp_dir/$BINARY_NAME"
  if [ ! -f "$bin_path" ]; then
    echo "binary not found in downloaded archive" >&2
    exit 1
  fi

  install_binary "$bin_path"

  echo "installed ${BINARY_NAME} ${VERSION} to ${INSTALL_DIR}/${BINARY_NAME}"
}

main "$@"
