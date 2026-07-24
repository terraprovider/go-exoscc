package adminapi

import (
	"encoding/json"
	"net/http"
	"regexp"
	"testing"
	"time"
)

func TestNormalizeURL(t *testing.T) {
	cases := map[string]string{
		"https://deu01b.admin.protection.outlook.com:446/adminapi/beta/x/InvokeCommand": "https://deu01b.admin.protection.outlook.com:443/adminapi/beta/x/InvokeCommand",
		"https://outlook.office365.com/adminapi/beta/x/InvokeCommand":                   "https://outlook.office365.com/adminapi/beta/x/InvokeCommand",
		"https://host:443/p": "https://host:443/p",
	}
	for in, want := range cases {
		if got := normalizeURL(in); got != want {
			t.Errorf("normalizeURL(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestIsRedirect(t *testing.T) {
	for _, c := range []int{301, 302, 303, 307, 308} {
		if !isRedirect(c) {
			t.Errorf("isRedirect(%d) = false, want true", c)
		}
	}
	for _, c := range []int{200, 400, 404, 500} {
		if isRedirect(c) {
			t.Errorf("isRedirect(%d) = true, want false", c)
		}
	}
}

func TestNewGUID(t *testing.T) {
	re := regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`)
	seen := map[string]bool{}
	for i := 0; i < 100; i++ {
		g := newGUID()
		if !re.MatchString(g) {
			t.Fatalf("newGUID() = %q, not a v4 GUID", g)
		}
		if seen[g] {
			t.Fatalf("newGUID() produced a duplicate: %q", g)
		}
		seen[g] = true
	}
}

func TestParseRateLimit(t *testing.T) {
	h := http.Header{}
	h.Set("Rate-Limit-Limit", "3000")
	h.Set("Rate-Limit-Remaining", "2998")
	h.Set("Rate-Limit-Reset", "2026-07-23T23:11:45Z")
	rl := parseRateLimit(h)
	if rl.Limit != 3000 || rl.Remaining != 2998 {
		t.Fatalf("limit/remaining = %d/%d, want 3000/2998", rl.Limit, rl.Remaining)
	}
	if rl.Reset.IsZero() || rl.Reset.Year() != 2026 {
		t.Fatalf("reset = %v, want parsed 2026 time", rl.Reset)
	}
	_ = time.Now
}

func TestOdataResponseAndDecode(t *testing.T) {
	body := `{
	  "@odata.context": "https://h/adminapi/beta/x/$metadata#RoleGroup",
	  "@adminapi.warnings": ["heads up"],
	  "value": [
	    {"Name":"OrganizationManagement","Guid":"00000000-0000-0000-0000-000000000001"},
	    {"Name":"ComplianceAdministrator","Guid":"00000000-0000-0000-0000-000000000002"}
	  ]
	}`
	var env odataResponse
	if err := json.Unmarshal([]byte(body), &env); err != nil {
		t.Fatal(err)
	}
	if len(env.Value) != 2 || env.Value[0]["Name"] != "OrganizationManagement" {
		t.Fatalf("unexpected value: %+v", env.Value)
	}
	if len(env.Warnings) != 1 || env.Warnings[0] != "heads up" {
		t.Fatalf("warnings = %+v", env.Warnings)
	}

	res := &Result{Value: env.Value}
	type rg struct {
		Name string `json:"Name"`
		Guid string `json:"Guid"`
	}
	var got []rg
	if err := res.Decode(&got); err != nil {
		t.Fatal(err)
	}
	if len(got) != 2 || got[1].Name != "ComplianceAdministrator" {
		t.Fatalf("decoded = %+v", got)
	}
}

func TestParseAPIError(t *testing.T) {
	body := `{"error":{"code":"Forbidden","message":"User is not allowed to call Get-RoleGroup","innererror":{"x":1}}}`
	err := parseAPIError(403, []byte(body))
	ae, ok := err.(*APIError)
	if !ok {
		t.Fatalf("want *APIError, got %T", err)
	}
	if ae.Status != 403 || ae.Code != "Forbidden" {
		t.Fatalf("APIError = %+v", ae)
	}
	if ae.Error() == "" || len(ae.InnerRaw) == 0 {
		t.Fatalf("error string / inner missing: %q %s", ae.Error(), ae.InnerRaw)
	}
}

func TestNewValidatesRequired(t *testing.T) {
	if _, err := New(Options{TenantID: "t"}); err == nil {
		t.Error("expected error when Tokens is nil")
	}
	if _, err := New(Options{Tokens: StaticTokenProvider("x")}); err == nil {
		t.Error("expected error when TenantID is empty")
	}
	c, err := New(Options{TenantID: "t", Tokens: StaticTokenProvider("x")})
	if err != nil {
		t.Fatal(err)
	}
	if c.opt.Cloud.BaseHost != EXO.BaseHost {
		t.Errorf("default cloud not EXO: %+v", c.opt.Cloud)
	}
	if c.resolvedHost != EXO.BaseHost || c.connectionID == "" {
		t.Errorf("client not initialized: host=%q connID=%q", c.resolvedHost, c.connectionID)
	}
}
