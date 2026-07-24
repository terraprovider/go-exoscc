package main

import "testing"

func TestGoType(t *testing.T) {
	cases := []struct {
		p    param
		want string
	}{
		{param{Type: "switch", IsSwitch: true}, "bool"},
		{param{Type: "string"}, "string"},
		{param{Type: "System.String"}, "string"},
		{param{Type: "System.Object[]"}, "[]string"},
		{param{Type: "string[]"}, "[]string"},
		{param{Type: "System.Int32"}, "int"},
		{param{Type: "System.Boolean"}, "bool"},
		{param{Type: "System.Guid"}, "string"},
		{param{Type: "System.Object"}, "any"},
		{param{Type: "Microsoft.Exchange.Whatever"}, "any"},
	}
	for _, c := range cases {
		if got := goType(c.p); got != c.want {
			t.Errorf("goType(%+v) = %q, want %q", c.p, got, c.want)
		}
	}
}

func TestExportName(t *testing.T) {
	cases := map[string]string{
		"Identity": "Identity",
		"anr":      "Anr",
		"2FA":      "N2FA",
		"Foo-Bar":  "FooBar",
		"":         "X",
	}
	for in, want := range cases {
		if got := exportName(in); got != want {
			t.Errorf("exportName(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestGoName(t *testing.T) {
	cases := map[string]string{
		"Get-Mailbox":                  "GetMailbox",
		"New-ManagementRoleAssignment": "NewManagementRoleAssignment",
		"Get-RoleGroupMember":          "GetRoleGroupMember",
	}
	for in, want := range cases {
		if got := goName(in); got != want {
			t.Errorf("goName(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestBoundCheck(t *testing.T) {
	cases := []struct {
		field, psName, gotype, want string
	}{
		{"Archive", "Archive", "bool", `if p.Archive { m["Archive"] = true }`},
		{"Anr", "Anr", "string", `if p.Anr != "" { m["Anr"] = p.Anr }`},
		{"ResultSize", "ResultSize", "int", `if p.ResultSize != 0 { m["ResultSize"] = p.ResultSize }`},
		{"Roles", "Roles", "[]string", `if len(p.Roles) > 0 { m["Roles"] = p.Roles }`},
		{"Identity", "Identity", "any", `if p.Identity != nil { m["Identity"] = p.Identity }`},
	}
	for _, c := range cases {
		if got := boundCheck(c.field, c.psName, c.gotype); got != c.want {
			t.Errorf("boundCheck(%q,%q,%q) = %q, want %q", c.field, c.psName, c.gotype, got, c.want)
		}
	}
}

func TestFlexStringsUnmarshal(t *testing.T) {
	var single flexStrings
	if err := single.UnmarshalJSON([]byte(`"Only"`)); err != nil || len(single) != 1 || single[0] != "Only" {
		t.Fatalf("single: %v %v", single, err)
	}
	var many flexStrings
	if err := many.UnmarshalJSON([]byte(`["A","B"]`)); err != nil || len(many) != 2 {
		t.Fatalf("many: %v %v", many, err)
	}
	var none flexStrings
	if err := none.UnmarshalJSON([]byte(`null`)); err != nil || none != nil {
		t.Fatalf("null: %v %v", none, err)
	}
}
