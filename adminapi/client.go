// Package adminapi is a lean Go client for the Exchange Online / Security &
// Compliance ("Purview") Admin API — the REST/OData surface behind the
// ExchangeOnlineManagement PowerShell module (Connect-ExchangeOnline /
// Connect-IPPSSession).
//
// It speaks the same wire protocol as the module: it POSTs to
// {host}/adminapi/beta/{tenantID}/InvokeCommand with a body of
// {"CmdletInput":{"CmdletName","Parameters"}} and sends the module's header set
// so the traffic is indistinguishable from the PowerShell client.
//
// The core is intentionally thin. Transient-error and rate-limit (429 +
// Retry-After) retries are handled by a shared retrying RoundTripper
// (github.com/terraprovider/go-msadmin/retry) installed by default; pass NoRetry to
// opt out. Only the API-specific bits live here — regional-host redirect with
// auth re-add, :446->:443, the affinity cookie jar, the header fidelity set,
// OData error handling and @odata.nextLink paging.
//
// Auth is abstracted behind TokenProvider (an alias of go-msadmin/auth's); MSAL-
// backed implementations and credential resolution live in go-msadmin/msalauth and
// go-msadmin/authx, so this core stays free of any credential-SDK dependency.
package adminapi

import (
	"context"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"strings"

	"github.com/terraprovider/go-msadmin/auth"
	"github.com/terraprovider/go-msadmin/retry"
)

// Cloud selects the service endpoint / token resource.
type Cloud struct {
	Name     string // "EXO" or "SCC"
	BaseHost string // e.g. outlook.office365.com
	Resource string // token audience, e.g. https://outlook.office365.com
}

var (
	// EXO is Exchange Online (commercial cloud).
	EXO = Cloud{Name: "EXO", BaseHost: "outlook.office365.com", Resource: "https://outlook.office365.com"}
	// SCC is Security & Compliance / Purview (commercial cloud).
	SCC = Cloud{Name: "SCC", BaseHost: "ps.compliance.protection.outlook.com", Resource: "https://ps.compliance.protection.outlook.com"}
)

// TokenProvider returns a bearer access token for the given resource. The token's
// audience must match Cloud.Resource. Implementations should cache/refresh.
//
// It is an alias of go-msadmin/auth.TokenProvider — the neutral token abstraction
// shared across the Microsoft admin clients — so credential resolvers built on
// that interface (go-msadmin/authx) plug in here directly.
type TokenProvider = auth.TokenProvider

// StaticTokenProvider serves a fixed, pre-acquired JWT (handy for tests and for
// short-lived scripts). It ignores the resource argument.
type StaticTokenProvider string

func (s StaticTokenProvider) Token(context.Context, string) (string, error) {
	if s == "" {
		return "", fmt.Errorf("adminapi: empty static token")
	}
	return string(s), nil
}

// Options configures a Client.
type Options struct {
	Cloud    Cloud
	TenantID string        // tenant GUID (from the token 'tid' claim)
	Tokens   TokenProvider // required
	// Organization is the tenant's routing domain (e.g. contoso.onmicrosoft.com).
	// For app-only it drives the region: the anchor becomes UPN:OAuthUser@<org>,
	// which the frontend uses to redirect to the tenant's regional backend (this is
	// how the PowerShell module routes app-only, and is required for Purview/SCC).
	Organization string
	// Anchor overrides X-AnchorMailbox. If empty it defaults to
	// UPN:OAuthUser@<Organization> when Organization is set, else TID:<TenantID>.
	Anchor     string
	HTTPClient *http.Client // optional; a cookie-jar client is created if nil

	// RetryConfig tunes the built-in transient-error / rate-limit (429 +
	// Retry-After) retrying transport. The zero value uses sensible defaults.
	RetryConfig retry.Config
	// NoRetry disables the built-in retrying transport (e.g. when supplying an
	// HTTPClient whose transport already retries).
	NoRetry bool

	// Header-fidelity knobs (defaults mimic module 3.10.0 on PowerShell 7.5.2).
	ModuleVersion string // X-ClientModuleVersion / exomodule-version
	PSVersion     string // User-Agent PowerShell/<ver>, ps-version
	AcceptLang    string // Accept-Language
}

// Client talks to one cloud in one tenant.
type Client struct {
	opt          Options
	http         *http.Client
	connectionID string // minted once per client (stable across the session)

	// resolvedHost is the regional forest host learned from the first redirect and
	// pinned for all subsequent calls (compliance vs admin backend matters for RBAC).
	resolvedHost string
}

// New builds a Client. TenantID and Tokens are required.
func New(opt Options) (*Client, error) {
	if opt.Tokens == nil {
		return nil, fmt.Errorf("adminapi: Options.Tokens is required")
	}
	if opt.TenantID == "" {
		return nil, fmt.Errorf("adminapi: Options.TenantID is required")
	}
	if opt.Cloud.BaseHost == "" {
		opt.Cloud = EXO
	}
	hc := opt.HTTPClient
	if hc == nil {
		jar, _ := cookiejar.New(nil)
		hc = &http.Client{Jar: jar}
	} else if hc.Jar == nil {
		// The affinity cookie from the regional redirect must be retained.
		jar, _ := cookiejar.New(nil)
		hc.Jar = jar
	}
	// Never auto-follow: we re-add the bearer + normalize the port manually.
	hc.CheckRedirect = func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse }
	// Transparently retry throttling (429 + Retry-After) and transient 5xx.
	if !opt.NoRetry {
		if _, already := hc.Transport.(*retry.Transport); !already {
			hc.Transport = retry.NewTransport(hc.Transport, opt.RetryConfig)
		}
	}

	if opt.ModuleVersion == "" {
		opt.ModuleVersion = "3.10.0"
	}
	if opt.PSVersion == "" {
		opt.PSVersion = "7.5.2"
	}
	if opt.AcceptLang == "" {
		opt.AcceptLang = "en-US"
	}
	if opt.Anchor == "" {
		if opt.Organization != "" {
			opt.Anchor = "UPN:OAuthUser@" + opt.Organization
		} else {
			opt.Anchor = "TID:" + opt.TenantID
		}
	}
	return &Client{
		opt:          opt,
		http:         hc,
		connectionID: newGUID(),
		resolvedHost: opt.Cloud.BaseHost,
	}, nil
}

func (c *Client) invokeURL(host string) string {
	return fmt.Sprintf("https://%s/adminapi/beta/%s/InvokeCommand", host, c.opt.TenantID)
}

// moduleFileURL is the connect-time bootstrap that pins the regional host.
func (c *Client) moduleFileURL(host string) string {
	return fmt.Sprintf("https://%s/AdminApi/v1.0/%s/EXOModuleFile", host, c.opt.TenantID)
}

func hostOf(rawurl string) string {
	i := strings.Index(rawurl, "://")
	if i < 0 {
		return ""
	}
	rest := rawurl[i+3:]
	if j := strings.IndexAny(rest, "/"); j >= 0 {
		rest = rest[:j]
	}
	return rest
}
