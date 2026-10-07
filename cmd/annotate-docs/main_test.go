package main

import (
	"bufio"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const page = "# Set-Thing\n\n## SYNTAX\n\n### Identity\n```\nSet-Thing -Identity <ThingIdParameter>\n```\n\n## PARAMETERS\n\n" +
	"### -Identity\n\nThe Identity parameter.\n\n```yaml\nType: ThingIdParameter\nParameter Sets: (All)\n```\n\n" +
	"### -Enabled\n\n```yaml\nType: System.Boolean\nParameter Sets: A\n```\n\n```yaml\nType: String\nParameter Sets: B\n```\n\n" +
	"### -Actions\n\n```yaml\nType: MultiValuedProperty\n```\n\n" +
	"### CommonParameters\n\n```yaml\nType: Bogus\n```\n"

func TestParseTypes(t *testing.T) {
	got, err := parseTypes(bufio.NewScanner(strings.NewReader(page)))
	if err != nil {
		t.Fatal(err)
	}
	want := map[string]string{"identity": "ThingIdParameter", "enabled": "System.Boolean", "actions": "MultiValuedProperty"}
	if len(got) != len(want) {
		t.Fatalf("got %v, want %v", got, want)
	}
	for k, v := range want {
		if got[k] != v {
			t.Errorf("%s = %q, want %q", k, got[k], v)
		}
	}
}

func TestAnnotate(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "Set-Thing.md"), []byte(page), 0o644); err != nil {
		t.Fatal(err)
	}
	pages, err := indexPages(dir)
	if err != nil {
		t.Fatal(err)
	}
	cat := catalog{Source: "EXO-ExchangeOnline.psm1", Cmdlets: []cmdlet{
		{Cmdlet: "Set-Thing", Parameters: []param{{Name: "Enabled"}, {Name: "actions"}, {Name: "Undocumented", DeclaredType: "stale"}}},
		{Cmdlet: "Get-Hidden", Parameters: []param{{Name: "Identity"}}},
	}}
	ovs := []override{
		{Catalog: "EXO", Cmdlet: "Set-Thing", Parameter: "Actions", Type: "System.String"},
		{Catalog: "Purview", Cmdlet: "Not-InThisCatalog", Parameter: "X", Type: "System.String"}, // other catalog: ignored
	}
	st, err := annotate(&cat, pages, ovs)
	if err != nil {
		t.Fatal(err)
	}
	ps := cat.Cmdlets[0].Parameters
	if ps[0].DeclaredType != "System.Boolean" || ps[1].DeclaredType != "System.String" || ps[2].DeclaredType != "" {
		t.Errorf("declared types = %q %q %q", ps[0].DeclaredType, ps[1].DeclaredType, ps[2].DeclaredType)
	}
	if st.params != 4 || st.typed != 2 || st.overridden != 1 || len(st.noPage) != 1 {
		t.Errorf("stats = %+v", st)
	}

	for _, ov := range []override{
		{Catalog: "EXO", Cmdlet: "Set-Thing", Parameter: "Gone", Type: "System.String"},      // stale parameter
		{Catalog: "EXO", Cmdlet: "Set-Missing", Parameter: "Actions", Type: "System.String"}, // misspelled/removed cmdlet
		{Cmdlet: "Set-Thing", Parameter: "Actions", Type: "System.String"},                   // unscoped
	} {
		if _, err := annotate(&cat, pages, []override{ov}); err == nil {
			t.Errorf("override %+v should fail", ov)
		}
	}
}

// The catalog must round-trip byte-for-byte so annotation only adds lines.
func TestRoundTripPreservesCatalog(t *testing.T) {
	for _, f := range []string{"EXO", "Purview"} {
		raw, err := os.ReadFile(filepath.Join("..", "..", "spec", "catalog", f+"-catalog.json"))
		if err != nil {
			t.Fatal(err)
		}
		var cat catalog
		if err := json.Unmarshal(raw, &cat); err != nil {
			t.Fatal(err)
		}
		out, err := encode(cat)
		if err != nil {
			t.Fatal(err)
		}
		if string(out) != string(raw) {
			t.Errorf("%s catalog does not round-trip", f)
		}
	}
}

func TestInheritNewSet(t *testing.T) {
	cat := catalog{Cmdlets: []cmdlet{
		{Cmdlet: "New-Policy", Verb: "New", Noun: "Policy", Parameters: []param{
			{Name: "IPAllowList", Type: "System.Object"},
			{Name: "Mode", Type: "System.Object", DeclaredType: "PolicyMode"},
			{Name: "Typed", Type: "string"},
		}},
		{Cmdlet: "Set-Policy", Verb: "Set", Noun: "Policy", Parameters: []param{
			{Name: "IPAllowList", Type: "System.Object", DeclaredType: "MultiValuedProperty"},
			{Name: "Mode", Type: "System.Object"},
			{Name: "Typed", Type: "System.Object", DeclaredType: "Boolean"},
		}},
		{Cmdlet: "Get-Policy", Verb: "Get", Noun: "Policy", Parameters: []param{
			{Name: "IPAllowList", Type: "System.Object"},
		}},
	}}
	if n := inheritNewSet(&cat); n != 2 {
		t.Errorf("inherited %d, want 2", n)
	}
	got := func(c, p int) string { return cat.Cmdlets[c].Parameters[p].DeclaredType }
	if got(0, 0) != "MultiValuedProperty" || got(1, 1) != "PolicyMode" {
		t.Errorf("New -IPAllowList = %q, Set -Mode = %q", got(0, 0), got(1, 1))
	}
	if got(0, 2) != "" {
		t.Errorf("a concrete psm1 type must not inherit, got %q", got(0, 2))
	}
	if got(2, 0) != "" {
		t.Errorf("Get must not inherit, got %q", got(2, 0))
	}
}
