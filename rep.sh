#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

# Root dir (default = current working dir)
ROOT_DIR="${1:-.}"

echo "🔍 Searching for .go files inside: $ROOT_DIR"

# Find and replace
find "$ROOT_DIR" -type f -name "*.go" | while read -r file; do
  if grep -q "github.com/ipfs/boxo" "$file"; then
    echo "⚡ Updating: $file"
    # Use perl for cross-platform safe replacement (avoids BSD/GNU sed issues)
    perl -pi -e 's|github\.com/ipfs/boxo|github.com/nnlgsakib/wwfs-sdk|g' "$file"
  fi
done

echo "✅ Replacement done!"
