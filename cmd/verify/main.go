// Command verify is a smoke test: it drives the generated bindings against the
// live Admin API and prints real data, proving auth + transport (regional
// redirect, cookie, headers) + response parsing all work.
//
// App-only (ARM_* env or flags):
//
//	go run ./cmd/verify                       # ARM_TENANT_ID / ARM_CLIENT_ID / ARM_CLIENT_SECRET
//	go run ./cmd/verify -tenant <t> -client-id <id>
//
// Or with a pre-acquired delegated token:
//
//	go run ./cmd/verify -token ./token.jwt
package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/terraprovider/go-exoscc/adminapi"
	"github.com/terraprovider/go-exoscc/exo"
	"github.com/terraprovider/go-msadmin/authx"
)

func main() {
	tokenFile := flag.String("token", "", "raw access-token JWT (delegated); omit to use app-only ARM_* creds")
	tenant := flag.String("tenant", "", "tenant (default $ARM_TENANT_ID)")
	clientID := flag.String("client-id", "", "app id (default $ARM_CLIENT_ID)")
	flag.Parse()

	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()

	var tp adminapi.TokenProvider
	var tid, anchor string

	if *tokenFile != "" {
		raw, err := os.ReadFile(*tokenFile)
		check(err)
		jwt := strings.TrimSpace(string(raw))
		tp = adminapi.StaticTokenProvider(jwt)
		tid = claim(jwt, "tid")
		upn := claim(jwt, "upn")
		if upn == "" {
			upn = claim(jwt, "unique_name")
		}
		if upn != "" {
			anchor = "UPN:" + upn
		} else {
			anchor = "TID:" + tid
		}
	} else {
		cfg := authx.FromEnv().Merge(authx.Config{TenantID: *tenant, ClientID: *clientID})
		var err error
		tp, err = cfg.Build()
		check(err)
		fmt.Printf("auth=%s\n", cfg.Method())
		tok, err := tp.Token(ctx, adminapi.EXO.Resource)
		check(err)
		tid = claim(tok, "tid")
		anchor = "TID:" + tid
		fmt.Printf("token: aud=%s appid=%s roles=%s\n", claim(tok, "aud"), claim(tok, "appid"), claimList(tok, "roles"))
	}
	if tid == "" {
		fail("verify: could not determine tenant id")
	}
	fmt.Printf("tenant=%s anchor=%s\n", tid, anchor)

	c, err := adminapi.New(adminapi.Options{Cloud: adminapi.EXO, TenantID: tid, Tokens: tp, Anchor: anchor})
	check(err)
	svc := exo.New(c)

	res, err := svc.GetOrganizationConfig(ctx, exo.GetOrganizationConfigParams{})
	check(err)
	fmt.Printf("Get-OrganizationConfig: %d object(s)\n", len(res.Value))
	if len(res.Value) > 0 {
		fmt.Printf("  Name=%v  AdminDisplayVersion=%v\n", res.Value[0]["Name"], res.Value[0]["AdminDisplayVersion"])
	}
	fmt.Printf("RateLimit: %d/%d\n", res.RateLimit.Remaining, res.RateLimit.Limit)

	ad, err := svc.GetAcceptedDomain(ctx, exo.GetAcceptedDomainParams{})
	if err != nil {
		fmt.Println("Get-AcceptedDomain:", err)
		return
	}
	fmt.Printf("Get-AcceptedDomain: %d domain(s)\n", len(ad.Value))
	for _, d := range ad.Value {
		fmt.Printf("  - %v (default=%v)\n", d["DomainName"], d["Default"])
	}
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

func claimList(jwt, name string) string {
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
	if arr, ok := m[name].([]any); ok {
		var out []string
		for _, v := range arr {
			out = append(out, fmt.Sprint(v))
		}
		return strings.Join(out, ",")
	}
	return ""
}

func check(err error) {
	if err != nil {
		fail("verify: " + err.Error())
	}
}
func fail(msg string) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(1)
}
