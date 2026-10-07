# go-exoscc

[![CI](https://github.com/terraprovider/go-exoscc/actions/workflows/ci.yml/badge.svg)](https://github.com/terraprovider/go-exoscc/actions/workflows/ci.yml)

A Go client for the **Exchange Online** and **Security & Compliance (Purview)**
Admin API — the REST/OData service behind the `ExchangeOnlineManagement`
PowerShell module (`Connect-ExchangeOnline` / `Connect-IPPSSession`). Typed
bindings for the cmdlets are **generated** from the module's own exported command
definitions and the OData `$metadata`, so the surface stays in lock-step with the
service.

> Not affiliated with or endorsed by Microsoft. It calls the same documented
> admin cmdlets the PowerShell module does, over the same REST transport.

```bash
go get github.com/terraprovider/go-exoscc
```

## Quickstart

```go
import (
    "github.com/terraprovider/go-exoscc/adminapi"
    "github.com/terraprovider/go-exoscc/exo"
    "github.com/terraprovider/go-exoscc/models"
    "github.com/terraprovider/go-exoscc/msalauth"
)

// App-only with a certificate (recommended):
tp, _ := msalauth.NewConfidentialCertPEM(tenantID, appID, pemBytes, "" /*pw*/, "")

c, _ := adminapi.New(adminapi.Options{
    Cloud:        adminapi.EXO,              // or adminapi.SCC for Purview
    TenantID:     tenantID,
    Tokens:       tp,
    Organization: "contoso.onmicrosoft.com", // app-only routing domain; required for SCC/Purview
    // (delegated instead? set Anchor: "UPN:user@contoso.com")
})

svc := exo.New(c)
res, _ := svc.GetAcceptedDomain(ctx, exo.GetAcceptedDomainParams{})

var domains []models.AcceptedDomain
_ = res.Decode(&domains)               // typed, from $metadata; or use res.Value ([]map[string]any)
```

Purview cmdlets live in the `purview` package and use `adminapi.SCC`.

## Auth (`msalauth`)

All backed by [MSAL for Go](https://github.com/AzureAD/microsoft-authentication-library-for-go):

| Constructor | Flow |
|-------------|------|
| `NewConfidentialCertPEM` / `NewConfidentialCert` | app-only, certificate (recommended) |
| `NewConfidentialSecret` | app-only, client secret¹ |
| `NewConfidentialAssertion` + `GitHubOIDCAssertion` | app-only, **federated** (workload identity / GitHub OIDC, no secret) |
| `NewDelegated` | interactive user (browser) with silent refresh |

Or implement `adminapi.TokenProvider` yourself. Connection pooling and retries are
left to the `*http.Client` you pass in — the client only owns the API-specific
transport (regional redirect, affinity cookie, header parity, OData paging, and
`Rate-Limit-*` surfacing).

¹ App-only auth needs the **Office 365 Exchange Online → `Exchange.ManageAsApp`**
application permission **plus RBAC on the app's service principal** — an Entra
directory role, or (preferred) a service-principal role-group membership
(`New-ServicePrincipal` + `Add-RoleGroupMember`; for Purview do this in a
`Connect-IPPSSession`). Certificate **and** client secret both work. Without a
role you'll get a 403 / "role isn't supported" — that's an assignment gap, not a
token problem.

## Packages

| Package | Contents |
|---------|----------|
| `adminapi` | client core: transport, `InvokeCommand`, paging, errors, `TokenProvider` |
| `exo` | generated Exchange Online bindings (825 cmdlets) |
| `purview` | generated Security & Compliance bindings (418 cmdlets) |
| `models` | types generated from `$metadata` (enums + structs) for `Result.Decode` |
| `msalauth` | MSAL-backed token providers |

`cmd/gen-go`, `cmd/gen-models` (generators), `cmd/annotate-docs` (declared
parameter types), `cmd/fetch-spec` (pulls the live spec), `cmd/verify` (smoke test). Generated files are `zz_generated_*.go` — **do not edit**.

## How the bindings are generated

```
Admin API  --cmd/fetch-spec-->  ExchangeOnline.psm1  --generator/extract-catalog.ps1 (PowerShell AST)-->  spec/catalog/*.json
    │                            $metadata  -----------------------------------------------------------> spec/metadata/*.xml
    └──────────────────────────────────────────────  cmd/gen-go / cmd/gen-models  ──────────────────────>  exo/ purview/ models/
MicrosoftDocs/office-docs-powershell  --cmd/annotate-docs-->  spec/catalog/*.json (declaredType per parameter)
```

- `spec/catalog/*.json` and `spec/metadata/*.xml` are the derived, committed inputs.
- The raw Microsoft `ExchangeOnline.psm1` is fetched **transiently** and never
  committed.
- The psm1 types most parameters as `System.Object`. `cmd/annotate-docs` records
  each parameter's documented .NET type (`declaredType`) from the cmdlet reference
  at the commit pinned in `spec/docs-ref` (`tools/fetch-docs.sh`); known docs
  errors are corrected in `spec/declared-type-overrides.json`.
  `spec.Param.Kind()` maps the types to the binding:

  | Kind | Go field | Sent when |
  |------|----------|-----------|
  | switch | `bool` | `true` |
  | Boolean | `*bool` | non-nil (so `false` can be sent) |
  | Int16/32/64, UInt32/64 | `*int64` | non-nil (so `0` can be sent) |
  | `X[]`, MultiValuedProperty, `*Collection` | `[]string` | non-nil (full replace) |
  | String | `string` | non-empty |
  | anything else (Unlimited, `*IdParameter`, enums, …) | `any` | non-nil |

  Sending an empty list does **not** clear a MultiValuedProperty. To clear or
  edit one in place, Set-* cmdlets have a `<Field>Delta *adminapi.StringDelta`
  companion (`{Add, Remove}`, sent like the module sends `@{Add=…;Remove=…}`: an `Exchange.GenericHashTable` object) for
  every `spec.Param.DeltaCapable()` parameter; it takes precedence over `<Field>`.
- Both **EXO** and **Purview** refresh fully app-only. `fetch-spec` auto-discovers
  the tenant routing domain (via EXO `Get-OrganizationConfig`) and uses the
  `OAuthUser@<domain>` anchor so the compliance calls resolve to the tenant's
  region (the client rewrites the region-redirect's `.admin.` backend to the
  `.ps.compliance.` one). The refreshed cmdlet set tracks the app's role
  membership, so assign broad ones (e.g. Organization Management + Compliance
  Administrator) for the full surface.
- Regenerate locally (app-only cert/secret):

  ```bash
  ./tools/regen.sh -tenant contoso.onmicrosoft.com -client-id <appId> -cert ./app.pem
  ```

## CI

- **ci.yml** — `go build` / `go vet` / `gofmt` / `go test -race` on every push & PR.
- **vuln.yml** — `govulncheck` on PRs and weekly.
- **refresh-spec.yml** — weekly (and on demand) pulls the live spec app-only via a
  **GitHub OIDC federated credential** (no secret), regenerates, and opens a PR on
  change. Configure repo *Variables* `ARM_TENANT_ID` and `ARM_CLIENT_ID`, and add a
  federated identity credential on the Entra app trusting this repo's OIDC subject
  (audience `api://AzureADTokenExchange`).

`cmd/fetch-spec` and `cmd/verify` read the standard Terraform **`ARM_*`** env vars
(`ARM_TENANT_ID`, `ARM_CLIENT_ID`, `ARM_CLIENT_SECRET` /
`ARM_CLIENT_CERTIFICATE_PATH` / `ARM_USE_OIDC`) so the same credentials drive local
testing, CI, and the future Terraform provider.

All actions are SHA-pinned; Dependabot keeps modules and actions current.

## License

[MIT](LICENSE) © glueckanja AG
