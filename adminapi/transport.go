package adminapi

import (
	"bytes"
	"compress/flate"
	"compress/gzip"
	"context"
	"crypto/rand"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"strings"
)

// debugEnabled turns on redirect/host tracing to stderr when EXOSCC_DEBUG is set.
func debugEnabled() bool { return os.Getenv("EXOSCC_DEBUG") != "" }

// newGUID returns a random RFC-4122 v4 GUID string (lowercase).
func newGUID() string {
	var b [16]byte
	_, _ = rand.Read(b[:])
	b[6] = (b[6] & 0x0f) | 0x40
	b[8] = (b[8] & 0x3f) | 0x80
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:16])
}

func (c *Client) userAgent() string {
	// Mimic the module's UA: Mozilla/5.0 (<os>) PowerShell/<ver>
	osTok := "Windows NT 10.0; Microsoft Windows"
	switch runtime.GOOS {
	case "darwin":
		osTok = "Macintosh; Darwin"
	case "linux":
		osTok = "Linux"
	}
	return fmt.Sprintf("Mozilla/5.0 (%s) PowerShell/%s", osTok, c.opt.PSVersion)
}

// setInvokeHeaders applies the per-cmdlet header set the module sends, verbatim,
// so the request is wire-indistinguishable from the PowerShell client.
// client-request-id is fresh per call; connection-id is stable per client.
func (c *Client) setInvokeHeaders(r *http.Request, cmdlet, token string) {
	h := r.Header
	h.Set("Authorization", "Bearer "+token)
	h.Set("Content-Type", "application/json")
	h.Set("Accept", "application/json")
	h.Set("Accept-Charset", "UTF-8")
	h.Set("Accept-Encoding", "gzip, deflate, br")
	h.Set("Accept-Language", c.opt.AcceptLang)
	h.Set("Prefer", "odata.maxpagesize=1000")
	h.Set("X-ResponseFormat", "json") // module default is clixml; we want structured JSON
	h.Set("X-SerializationLevel", "Partial")
	h.Set("X-CmdletName", cmdlet)
	h.Set("X-AnchorMailbox", c.opt.Anchor)
	h.Set("X-ClientApplication", "ExoManagementModule")
	h.Set("X-ClientModuleVersion", c.opt.ModuleVersion)
	h.Set("WarningAction", "")
	h.Set("connection-id", c.connectionID)
	h.Set("client-request-id", newGUID())
	h.Set("User-Agent", c.userAgent())
}

// normalizeURL rewrites the backend port 446 to the public HTTPS listener 443.
func normalizeURL(raw string) string {
	u, err := url.Parse(raw)
	if err != nil {
		return raw
	}
	if u.Port() == "446" {
		u.Host = u.Hostname() + ":443"
	}
	return u.String()
}

// rewriteRedirect normalizes a redirect target. For the compliance cloud, the
// region-resolving redirect from the generic host lands on the Exchange *admin*
// backend ({region}.admin.protection.outlook.com); the compliance cmdlets live on
// {region}.ps.compliance.protection.outlook.com, so swap the suffix (keeping the
// region prefix) while leaving EXO untouched.
func (c *Client) rewriteRedirect(loc string) string {
	loc = normalizeURL(loc)
	if strings.Contains(c.opt.Cloud.BaseHost, "compliance") {
		loc = strings.Replace(loc, ".admin.protection.outlook.com", ".ps.compliance.protection.outlook.com", 1)
	}
	return loc
}

// postInvoke sends one InvokeCommand POST, following the regional 302 redirect(s)
// manually (re-adding auth, normalizing :446->:443, retaining cookies via the jar)
// and pinning the resolved host for future calls. Returns the final response.
func (c *Client) postInvoke(ctx context.Context, cmdlet string, body []byte) (*http.Response, error) {
	token, err := c.opt.Tokens.Token(ctx, c.opt.Cloud.Resource)
	if err != nil {
		return nil, fmt.Errorf("adminapi: token: %w", err)
	}
	target := c.invokeURL(c.resolvedHost)
	seen := map[string]int{}
	for hop := 0; hop < 8; hop++ {
		req, err := http.NewRequestWithContext(ctx, http.MethodPost, target, bytes.NewReader(body))
		if err != nil {
			return nil, err
		}
		c.setInvokeHeaders(req, cmdlet, token)
		resp, err := c.http.Do(req)
		if err != nil {
			return nil, err
		}
		if debugEnabled() {
			fmt.Fprintf(os.Stderr, "[POST] %s -> %d\n", target, resp.StatusCode)
		}
		if isRedirect(resp.StatusCode) {
			loc := resp.Header.Get("Location")
			_ = resp.Body.Close()
			if loc == "" {
				return nil, fmt.Errorf("adminapi: %d redirect without Location", resp.StatusCode)
			}
			loc = c.rewriteRedirect(loc)
			if seen[loc] >= 2 {
				return nil, fmt.Errorf("adminapi: redirect loop at %s", loc)
			}
			seen[loc]++
			// Pin the regional host for subsequent calls (strip back to InvokeCommand URL).
			if h := hostOf(loc); h != "" {
				c.resolvedHost = h
			}
			target = loc
			continue
		}
		return resp, nil
	}
	return nil, fmt.Errorf("adminapi: too many redirects")
}

// bootstrapHeaders are the module's connect-time headers, sent on the EXOModuleFile
// GET (they mirror what Connect-* sends). Kept in one place for fidelity.
func (c *Client) bootstrapHeaders() map[string]string {
	return map[string]string{
		"CommandName":                "*",
		"ps-version":                 c.opt.PSVersion,
		"exomodule-version":          c.opt.ModuleVersion,
		"is-cloud-shell-environment": "False",
		"X-InitializeConnectionContextObjectIdDynamic": "true",
		"connection-id":     c.connectionID,
		"client-request-id": newGUID(),
	}
}

// getFollow issues an authenticated GET, following the regional redirect manually
// (re-adding auth, :446->:443, cookie jar) and pinning the resolved host.
func (c *Client) getFollow(ctx context.Context, urlStr string, headers map[string]string) (*http.Response, error) {
	token, err := c.opt.Tokens.Token(ctx, c.opt.Cloud.Resource)
	if err != nil {
		return nil, fmt.Errorf("adminapi: token: %w", err)
	}
	target := normalizeURL(urlStr)
	seen := map[string]int{}
	for hop := 0; hop < 8; hop++ {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, target, nil)
		if err != nil {
			return nil, err
		}
		req.Header.Set("Authorization", "Bearer "+token)
		req.Header.Set("X-AnchorMailbox", c.opt.Anchor)
		req.Header.Set("Accept", "application/json")
		req.Header.Set("Accept-Encoding", "gzip, deflate")
		req.Header.Set("User-Agent", c.userAgent())
		for k, v := range headers {
			req.Header.Set(k, v)
		}
		resp, err := c.http.Do(req)
		if err != nil {
			return nil, err
		}
		if debugEnabled() {
			fmt.Fprintf(os.Stderr, "[GET] %s -> %d\n", target, resp.StatusCode)
		}
		if isRedirect(resp.StatusCode) {
			loc := resp.Header.Get("Location")
			_ = resp.Body.Close()
			if loc == "" {
				return nil, fmt.Errorf("adminapi: %d redirect without Location", resp.StatusCode)
			}
			loc = c.rewriteRedirect(loc)
			if seen[loc] >= 2 {
				return nil, fmt.Errorf("adminapi: redirect loop at %s", loc)
			}
			seen[loc]++
			if h := hostOf(loc); h != "" {
				c.resolvedHost = h
			}
			target = loc
			continue
		}
		return resp, nil
	}
	return nil, fmt.Errorf("adminapi: too many redirects")
}

// FetchAdmin GETs an Admin API path on this cloud and returns the decoded body.
// "{tid}" in pathTemplate is replaced with the tenant ID. Set bootstrap=true for
// the EXOModuleFile endpoint (sends the module's connect headers). Used to pull
// $metadata and the EXOModuleFile cmdlet manifest.
func (c *Client) FetchAdmin(ctx context.Context, pathTemplate string, bootstrap bool) ([]byte, error) {
	path := strings.ReplaceAll(pathTemplate, "{tid}", c.opt.TenantID)
	u := "https://" + c.resolvedHost + path
	h := map[string]string{}
	if bootstrap {
		h = c.bootstrapHeaders()
	}
	// $metadata is CSDL XML; everything else is JSON.
	if strings.HasSuffix(path, "$metadata") {
		h["Accept"] = "application/xml"
	}
	resp, err := c.getFollow(ctx, u, h)
	if err != nil {
		return nil, err
	}
	body, err := readBody(resp)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 400 {
		return body, fmt.Errorf("adminapi: GET %s -> %d: %s", path, resp.StatusCode, snippet(body))
	}
	return body, nil
}

func snippet(b []byte) string {
	if len(b) > 2000 {
		return string(b[:2000])
	}
	return string(b)
}

func isRedirect(code int) bool {
	switch code {
	case http.StatusMovedPermanently, http.StatusFound, http.StatusSeeOther,
		http.StatusTemporaryRedirect, http.StatusPermanentRedirect:
		return true
	}
	return false
}

// readBody returns the decompressed response body. Because we set Accept-Encoding
// ourselves (for header fidelity), net/http does NOT auto-decompress, so we do it
// here based on Content-Encoding — but defensively: some responses (notably
// errors) advertise gzip yet aren't, so we sniff the magic bytes and fall back to
// the raw body rather than failing. brotli ("br") is passed through raw.
func readBody(resp *http.Response) ([]byte, error) {
	defer resp.Body.Close()
	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	switch strings.ToLower(resp.Header.Get("Content-Encoding")) {
	case "gzip":
		if len(raw) >= 2 && raw[0] == 0x1f && raw[1] == 0x8b {
			if zr, e := gzip.NewReader(bytes.NewReader(raw)); e == nil {
				defer zr.Close()
				if out, e := io.ReadAll(zr); e == nil {
					return out, nil
				}
			}
		}
	case "deflate":
		fr := flate.NewReader(bytes.NewReader(raw))
		defer fr.Close()
		if out, e := io.ReadAll(fr); e == nil {
			return out, nil
		}
	}
	return raw, nil
}
