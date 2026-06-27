#!/usr/bin/env bash
# setup.sh — rename module and clean up template markers
# Usage: ./setup.sh mytool github.com/you/mytool

set -euo pipefail

if [ $# -lt 2 ]; then
  echo "Usage: $0 <tool-name> <module-path>"
  echo "Example: $0 portfold github.com/jlcoulter/portfold"
  exit 1
fi

NAME="$1"
MODULE="$2"
OLD_MODULE="github.com/jlcoulter/go-cli-template"

# Replace module path in all Go files
find . -name "*.go" -exec sed -i "s|${OLD_MODULE}|${MODULE}|g" {} +

# Replace in go.mod
sed -i "s|${OLD_MODULE}|${MODULE}|g" go.mod

# Replace in Makefile
sed -i "s|go-cli-template|${NAME}|g" Makefile

# Replace in Dockerfile
sed -i "s|go-cli-template|${NAME}|g" Dockerfile

# Replace in README
sed -i "s|go-cli-template|${NAME}|g" README.md

# Replace in .goreleaser.yml
sed -i "s|go-cli-template|${NAME}|g" .goreleaser.yml

# Replace in CI
sed -i "s|go-cli-template|${NAME}|g" .github/workflows/ci.yml

# Remove example command and internal — keep as reference or delete
echo ""
echo "NOTE: cmd/example.go and internal/example/ contain template examples."
echo "Delete them when you've added your own commands."
echo ""

# Remove setup script
rm -- "$0"

echo "Template configured: name=${NAME}, module=${MODULE}"
echo "Next steps:"
echo "  1. Run: go mod tidy"
echo "  2. Run: make test"
echo "  3. Add your own commands in cmd/"