// Command fetch-spec pulls the live Admin API spec — the OData $metadata and the
// EXOModuleFile cmdlet manifest (the raw Microsoft ExchangeOnline.psm1) — for the
// selected clouds, app-only. Credentials come from ARM_* env vars (or flags); in
// CI it uses a GitHub Actions OIDC token via an Entra federated identity
// credential (ARM_USE_OIDC=true), locally a certificate or client secret.
//
// The Microsoft psm1 is written to -psm1-dir (transient, git-ignored — not
// committed for copyright reasons); the $metadata (schema) to -metadata-dir. A
// downstream step runs generator/extract-catalog.ps1 on the psm1 to (re)produce
// the committed catalog JSON, then the Go generators.
//
// App-only requires the Entra app to have Office 365 Exchange Online ->
// Exchange.ManageAsApp AND a supported directory role on its service principal;
// without the role the service returns "The role assigned to application ... isn't
// supported in this scenario".
package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/terraprovider/go-exoscc/adminapi"
	"github.com/terraprovider/go-msadmin/authx"
)

func main() {
	var (
		cloudSel = flag.String("cloud", "EXO", "EXO | Purview | Both")
		tenant   = flag.String("tenant", "", "tenant ID/domain (default: $ARM_TENANT_ID)")
		clientID = flag.String("client-id", "", "Entra app (client) ID (default: $ARM_CLIENT_ID)")
		auth     = flag.String("auth", "auto", "auto | federated | cert | secret")
		certPEM  = flag.String("cert-pem", "", "PEM cert+key (default: $ARM_CLIENT_CERTIFICATE_PATH)")
		certPass = flag.String("cert-password", "", "PEM key password (default: $ARM_CLIENT_CERTIFICATE_PASSWORD)")
		secret   = flag.String("secret", "", "client secret (default: $ARM_CLIENT_SECRET)")
		authHost = flag.String("authority-host", "", "national-cloud login host override")
		psm1Dir  = flag.String("psm1-dir", ".spec-cache", "where to write the (transient) Microsoft psm1")
		metaDir  = flag.String("metadata-dir", "spec/metadata", "where to write $metadata")
		hostOvr  = flag.String("host", "", "override the cloud base host (e.g. a regional compliance host)")
		anchor   = flag.String("anchor", "", "override X-AnchorMailbox (default derived from org/tid)")
		orgFlag  = flag.String("organization", "", "tenant routing domain (default: auto-discovered from EXO); needed to route Purview app-only")
	)
	flag.Parse()

	cfg := authx.FromEnv().Merge(authx.Config{
		TenantID: *tenant, ClientID: *clientID, ClientSecret: *secret,
		ClientCertificatePath: *certPEM, ClientCertificatePassword: *certPass,
	})
	if *authHost != "" {
		cfg.Environment = *authHost
	}
	switch strings.ToLower(*auth) {
	case "cli":
		cfg.UseCLI = true
	case "msi":
		cfg.UseMSI = true
	case "federated", "oidc":
		cfg.UseOIDC = true
	}
	tp, err := cfg.Build()
	check(err)
	fmt.Printf("auth=%s\n", cfg.Method())

	clouds := map[string]adminapi.Cloud{}
	switch strings.ToLower(*cloudSel) {
	case "exo":
		clouds["EXO"] = adminapi.EXO
	case "purview":
		clouds["Purview"] = adminapi.SCC
	case "both":
		clouds["EXO"] = adminapi.EXO
		clouds["Purview"] = adminapi.SCC
	default:
		fail("fetch-spec: -cloud must be EXO | Purview | Both")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()

	// Resolve the tenant GUID from a token (the Admin API path needs the tid claim).
	var anyResource string
	for _, c := range clouds {
		anyResource = c.Resource
		break
	}
	tok, err := tp.Token(ctx, anyResource)
	check(err)
	tid := claim(tok, "tid")
	if tid == "" {
		fail("fetch-spec: could not read tid from token")
	}
	fmt.Printf("tenant=%s clouds=%s\n", tid, *cloudSel)

	check(os.MkdirAll(*psm1Dir, 0o755))
	check(os.MkdirAll(*metaDir, 0o755))

	// The tenant routing domain drives app-only region resolution (esp. Purview).
	// Auto-discover it from EXO Get-OrganizationConfig if not supplied.
	org := *orgFlag
	if org == "" && *anchor == "" {
		if d, err := discoverOrg(ctx, tp, tid); err != nil {
			fmt.Fprintf(os.Stderr, "WARN could not discover organization (%v); Purview routing may fail\n", err)
		} else {
			org = d
			fmt.Printf("organization=%s (discovered)\n", org)
		}
	}

	failed := 0
	for name, cloud := range clouds {
		if *hostOvr != "" {
			cloud.BaseHost = *hostOvr
		}
		if err := fetchCloud(ctx, name, cloud, tp, tid, org, *anchor, *psm1Dir, *metaDir); err != nil {
			fmt.Fprintf(os.Stderr, "WARN %s: %v\n", name, err)
			failed++
		}
	}
	if failed == len(clouds) {
		os.Exit(1)
	}
}

// discoverOrg reads the tenant routing domain from EXO Get-OrganizationConfig
// (works app-only with a TID anchor), used to build the OAuthUser@<org> anchor
// that routes compliance (Purview) calls to the tenant's region.
func discoverOrg(ctx context.Context, tp adminapi.TokenProvider, tid string) (string, error) {
	cl, err := adminapi.New(adminapi.Options{Cloud: adminapi.EXO, TenantID: tid, Tokens: tp})
	if err != nil {
		return "", err
	}
	res, err := cl.Invoke(ctx, "Get-OrganizationConfig", nil)
	if err != nil {
		return "", err
	}
	if len(res.Value) == 0 {
		return "", fmt.Errorf("empty OrganizationConfig")
	}
	if name, ok := res.Value[0]["Name"].(string); ok && name != "" {
		return name, nil
	}
	return "", fmt.Errorf("OrganizationConfig has no Name")
}

func fetchCloud(ctx context.Context, name string, cloud adminapi.Cloud, tp adminapi.TokenProvider, tid, org, anchor, psm1Dir, metaDir string) error {
	// EXO resolves the tenant on the global host with a TID anchor; only the
	// compliance cloud needs the OAuthUser@<org> anchor for regional routing.
	orgForCloud := ""
	if strings.Contains(cloud.BaseHost, "compliance") {
		orgForCloud = org
	}
	cl, err := adminapi.New(adminapi.Options{Cloud: cloud, TenantID: tid, Tokens: tp, Organization: orgForCloud, Anchor: anchor})
	if err != nil {
		return err
	}

	// EXOModuleFile FIRST: for the compliance cloud this bootstrap pins the correct
	// regional backend and sets the affinity cookie that subsequent calls reuse.
	mf, err := cl.FetchAdmin(ctx, "/AdminApi/v1.0/{tid}/EXOModuleFile", true)
	if err != nil {
		return fmt.Errorf("EXOModuleFile: %w", err)
	}
	psm1, err := extractPsm1(mf)
	if err != nil {
		return err
	}
	psPath := filepath.Join(psm1Dir, name+"-ExchangeOnline.psm1")
	if err := os.WriteFile(psPath, []byte(psm1), 0o644); err != nil {
		return err
	}
	n := len(regexp.MustCompile(`(?m)^function script:[A-Za-z]+-[A-Za-z0-9]+`).FindAllString(psm1, -1))
	fmt.Printf("%s: %s (%d cmdlets)\n", name, psPath, n)

	// $metadata on the now-pinned host. Non-fatal: EXO and SCC share the model, so a
	// compliance-side $metadata hiccup doesn't block the psm1/catalog.
	md, err := cl.FetchAdmin(ctx, "/adminapi/beta/{tid}/$metadata", false)
	if err != nil {
		fmt.Fprintf(os.Stderr, "  %s $metadata skipped: %v\n", name, err)
		return nil
	}
	mdPath := filepath.Join(metaDir, name+"-metadata.xml")
	if err := os.WriteFile(mdPath, md, 0o644); err != nil {
		return err
	}
	fmt.Printf("%s: %s (%d bytes)\n", name, mdPath, len(md))
	return nil
}

// extractPsm1 pulls ExchangeOnline.psm1 out of the EXOModuleFile JSON envelope.
func extractPsm1(body []byte) (string, error) {
	var env struct {
		Value []struct {
			FileName    string `json:"fileName"`
			FileContent string `json:"fileContent"`
		} `json:"value"`
	}
	if err := json.Unmarshal(body, &env); err != nil {
		return "", fmt.Errorf("EXOModuleFile decode: %w", err)
	}
	for _, f := range env.Value {
		if f.FileName == "ExchangeOnline.psm1" {
			return f.FileContent, nil
		}
	}
	return "", fmt.Errorf("ExchangeOnline.psm1 not found in EXOModuleFile response")
}

func claim(jwt, name string) string {
	parts := strings.Split(jwt, ".")
	if len(parts) < 2 {
		return ""
	}
	b, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return ""
	}
	var m map[string]any
	if json.Unmarshal(b, &m) != nil {
		return ""
	}
	if v, ok := m[name].(string); ok {
		return v
	}
	return ""
}

func check(err error) {
	if err != nil {
		fail("fetch-spec: " + err.Error())
	}
}
func fail(msg string) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(1)
}
