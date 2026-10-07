package main

import (
	"testing"

	"github.com/terraprovider/go-exoscc/spec"
)

func TestGoType(t *testing.T) {
	cases := []struct {
		p    spec.Param
		want string
	}{
		{spec.Param{Type: "switch", IsSwitch: true}, "bool"},
		{spec.Param{Type: "string"}, "string"},
		{spec.Param{Type: "System.Object[]"}, "[]string"},
		{spec.Param{Type: "int"}, "*int64"},
		{spec.Param{Type: "bool"}, "*bool"},
		{spec.Param{Type: "System.Object", DeclaredType: "System.Boolean"}, "*bool"},
		{spec.Param{Type: "System.Object", DeclaredType: "MultiValuedProperty"}, "[]string"},
		{spec.Param{Type: "System.Object", DeclaredType: "Unlimited"}, "any"},
		{spec.Param{Type: "System.Object"}, "any"},
	}
	for _, c := range cases {
		if got := goType(c.p); got != c.want {
			t.Errorf("goType(%+v) = %q, want %q", c.p, got, c.want)
		}
	}
}

func TestFieldComment(t *testing.T) {
	cases := []struct {
		p    spec.Param
		want string
	}{
		{spec.Param{Type: "System.Object", DeclaredType: "Unlimited"}, " // Unlimited"},
		{spec.Param{Type: "System.Object", DeclaredType: "System.Boolean"}, ""},
		{spec.Param{Type: "string", ValidateSet: spec.FlexStrings{"A", "B"}}, " // one of: A, B"},
		{spec.Param{Type: "System.Object", DeclaredType: "Mode", ValidateSet: spec.FlexStrings{"On"}}, " // one of: On; Mode"},
	}
	for _, c := range cases {
		if got := fieldComment(c.p); got != c.want {
			t.Errorf("fieldComment(%+v) = %q, want %q", c.p, got, c.want)
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
		{"Enabled", "Enabled", "*bool", `if p.Enabled != nil { m["Enabled"] = *p.Enabled }`},
		{"ResultSize", "ResultSize", "*int64", `if p.ResultSize != nil { m["ResultSize"] = *p.ResultSize }`},
		{"Roles", "Roles", "[]string", `if p.Roles != nil { m["Roles"] = p.Roles }`},
		{"Identity", "Identity", "any", `if p.Identity != nil { m["Identity"] = p.Identity }`},
	}
	for _, c := range cases {
		if got := boundCheck(c.field, c.psName, c.gotype); got != c.want {
			t.Errorf("boundCheck(%q,%q,%q) = %q, want %q", c.field, c.psName, c.gotype, got, c.want)
		}
	}
}
