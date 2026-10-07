#!/usr/bin/env bash
# Fetch the Exchange cmdlet reference from MicrosoftDocs/office-docs-powershell at
# the commit pinned in spec/docs-ref (sparse, blobless) into <dir>. Used by
# cmd/annotate-docs; the checkout is transient and never committed.
#
#   ./tools/fetch-docs.sh .spec-cache/docs
set -euo pipefail
ROOT="$(cd "$(dirname "$0")/.." && pwd)"
DIR="${1:?usage: fetch-docs.sh <dir>}"
REF="$(tr -d '[:space:]' < "$ROOT/spec/docs-ref")"

if [ ! -d "$DIR/.git" ]; then
  git clone --quiet --filter=blob:none --no-checkout --sparse \
    https://github.com/MicrosoftDocs/office-docs-powershell.git "$DIR"
  git -C "$DIR" sparse-checkout set exchange/exchange-ps/ExchangePowerShell
fi
git -C "$DIR" fetch --quiet --filter=blob:none origin "$REF"
git -C "$DIR" -c advice.detachedHead=false checkout --quiet "$REF"
echo "docs: office-docs-powershell@$REF -> $DIR"
