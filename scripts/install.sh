#!/usr/bin/env bash
set -e

REPO="atevia/ship"
BIN_NAME="ship"
INSTALL_PATH="/usr/local/bin"

echo "🚀 Installing Atevia Ship..."

ARCH=$(uname -m)
OS=$(uname -s | tr '[:upper:]' '[:lower:]')

if [[ "$ARCH" != "x86_64" ]]; then
  echo "❌ Unsupported architecture: $ARCH"
  exit 1
fi

TMP_DIR=$(mktemp -d)
cd "$TMP_DIR"

echo "⬇ Downloading latest release..."
curl -s https://api.github.com/repos/$REPO/releases/latest \
  | grep browser_download_url \
  | grep linux_amd64 \
  | cut -d '"' -f 4 \
  | xargs curl -L -o ship

chmod +x ship
sudo mv ship $INSTALL_PATH/ship

echo "✅ ship installed successfully"
echo "👉 Run: ship help"