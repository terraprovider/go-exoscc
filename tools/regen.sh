#!/usr/bin/env bash
# Refresh go-exoscc from a live tenant (app-only) and regenerate the SDK.
#
#   ./tools/regen.sh -tenant contoso.onmicrosoft.com -client-id <appId> -cert ./app.pem [-cloud Both]
#   ./tools/regen.sh -tenant <t> -client-id <appId> -secret <secret>     [-cloud EXO]
#
# The raw Microsoft psm1 and the docs checkout land in .spec-cache (git-ignored,
# not committed); only the derived catalog JSON, $metadata and generated Go change.
set -euo pipefail
ROOT="$(cd "$(dirname "$0")/.." && pwd)"

TENANT="" CID="" CERT="" SECRET="" CLOUD="Both"
while [ $# -gt 0 ]; do
  case "$1" in
    -tenant) TENANT="$2"; shift 2;;
    -client-id) CID="$2"; shift 2;;
    -cert) CERT="$2"; shift 2;;
    -secret) SECRET="$2"; shift 2;;
    -cloud) CLOUD="$2"; shift 2;;
    *) echo "unknown arg: $1" >&2; exit 2;;
  esac
done
[ -n "$TENANT" ] && [ -n "$CID" ] || { echo "usage: regen.sh -tenant <t> -client-id <id> (-cert <pem> | -secret <s>) [-cloud EXO|Purview|Both]" >&2; exit 2; }

cd "$ROOT"
if [ -n "$CERT" ]; then AUTH=(-auth cert -cert-pem "$CERT")
elif [ -n "$SECRET" ]; then AUTH=(-auth secret -secret "$SECRET")
else echo "provide -cert <pem> or -secret <secret>" >&2; exit 2; fi

echo "==> 1/4 fetch live spec ($CLOUD)"
go run ./cmd/fetch-spec "${AUTH[@]}" -cloud "$CLOUD" -tenant "$TENANT" -client-id "$CID" \
  -psm1-dir .spec-cache -metadata-dir spec/metadata

echo "==> 2/4 psm1 -> catalog JSON (PowerShell AST)"
[ -f .spec-cache/EXO-ExchangeOnline.psm1 ]     && pwsh ./generator/extract-catalog.ps1 -Psm1 .spec-cache/EXO-ExchangeOnline.psm1     -Out spec/catalog/EXO-catalog.json
[ -f .spec-cache/Purview-ExchangeOnline.psm1 ] && pwsh ./generator/extract-catalog.ps1 -Psm1 .spec-cache/Purview-ExchangeOnline.psm1 -Out spec/catalog/Purview-catalog.json

echo "==> 3/4 declared parameter types from the docs (pinned in spec/docs-ref)"
./tools/fetch-docs.sh .spec-cache/docs
for c in EXO Purview; do
  if [ -f "spec/catalog/$c-catalog.json" ]; then
    go run ./cmd/annotate-docs -catalog "spec/catalog/$c-catalog.json" \
      -docs .spec-cache/docs/exchange/exchange-ps/ExchangePowerShell -docs-ref "$(cat spec/docs-ref)"
  fi
done

echo "==> 4/4 catalog + \$metadata -> Go"
go run ./cmd/gen-go -catalog spec/catalog/EXO-catalog.json -pkg exo \
  -client github.com/terraprovider/go-exoscc/adminapi -out exo/zz_generated_exo.go
[ -f spec/catalog/Purview-catalog.json ] && go run ./cmd/gen-go -catalog spec/catalog/Purview-catalog.json -pkg purview \
  -client github.com/terraprovider/go-exoscc/adminapi -out purview/zz_generated_purview.go
go run ./cmd/gen-models -metadata spec/metadata/EXO-metadata.xml -pkg models -out models/zz_generated_models.go

go build ./... && go test ./...
echo "==> done."
