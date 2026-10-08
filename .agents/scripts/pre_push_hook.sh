#!/usr/bin/env bash
# Pre-Push Quality & Safety Gate for Re-SPEC (Antigravity Hook) 🦖
set -euo pipefail

echo "🦖 [respec-gate] Running pre-push checks..."

# 1. Guard against staging or modifying forbidden files (.env, .riccardo/)
if git status --porcelain | grep -E '(\.env|\.riccardo/)' >/dev/null 2>&1; then
  echo "🛑 [respec-gate] ERROR: Staged or modified protected files detected (.env or .riccardo/)!" >&2
  exit 1
fi

# 2. Run unit tests
echo "🧪 [respec-gate] Running test suite..."
if command -v just >/dev/null 2>&1; then
  just test
else
  go test ./...
fi

echo "✅ [respec-gate] All pre-push checks passed! Ready to push! 🚀"
exit 0
