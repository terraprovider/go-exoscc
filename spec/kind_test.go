package spec

import "testing"

func TestKind(t *testing.T) {
	cases := []struct {
		p    Param
		want Kind
	}{
		{Param{Type: "switch", IsSwitch: true}, KindSwitch},
		{Param{Type: "string"}, KindString},
		{Param{Type: "System.String"}, KindString},
		{Param{Type: "guid"}, KindString},
		{Param{Type: "bool"}, KindBool},
		{Param{Type: "System.Boolean"}, KindBool},
		{Param{Type: "int"}, KindInt64},
		{Param{Type: "uint"}, KindInt64},
		{Param{Type: "uint32"}, KindInt64},
		{Param{Type: "ushort"}, KindInt64},
		{Param{Type: "System.Int64"}, KindInt64},
		{Param{Type: "string[]"}, KindList},
		{Param{Type: "System.Object[]"}, KindList},
		{Param{Type: "System.Collections.Generic.List[string]"}, KindList},
		{Param{Type: "double"}, KindAny},
		{Param{Type: "hashtable"}, KindAny},
		// System.Object is refined by the declared (docs) type.
		{Param{Type: "System.Object"}, KindAny},
		{Param{Type: "System.Object", DeclaredType: "System.Boolean"}, KindBool},
		{Param{Type: "System.Object", DeclaredType: "Boolean"}, KindBool},
		{Param{Type: "System.Object", DeclaredType: "Int32"}, KindInt64},
		{Param{Type: "System.Object", DeclaredType: "System.UInt32"}, KindInt64},
		{Param{Type: "System.Object", DeclaredType: "Int64?"}, KindInt64},
		{Param{Type: "System.Object", DeclaredType: "MultiValuedProperty"}, KindList},
		{Param{Type: "System.Object", DeclaredType: "ProxyAddressCollection"}, KindList},
		{Param{Type: "System.Object", DeclaredType: "RecipientIdParameter[]"}, KindList},
		{Param{Type: "System.Object", DeclaredType: "<MultiValuedProperty>"}, KindList},
		// Arrays of structured elements are not string lists.
		{Param{Type: "System.Object", DeclaredType: "System.Collections.Hashtable[]"}, KindAny},
		{Param{Type: "System.Object", DeclaredType: "PswsHashtable[]"}, KindAny},
		{Param{Type: "System.Object[]", DeclaredType: "PswsHashtable[]"}, KindAny},
		{Param{Type: "hashtable[]"}, KindAny},
		{Param{Type: "byte[]"}, KindAny},
		{Param{Type: "guid[]"}, KindList},
		{Param{Type: "string[]", DeclaredType: "Hashtable[]"}, KindList},
		{Param{Type: "System.Object", DeclaredType: "Unlimited"}, KindAny},
		{Param{Type: "System.Object", DeclaredType: "System.String"}, KindAny},
		{Param{Type: "System.Object", DeclaredType: "MailboxIdParameter"}, KindAny},
		// A concrete psm1 type wins over the declared type.
		{Param{Type: "string", DeclaredType: "MultiValuedProperty"}, KindString},
		{Param{Type: "System.Object[]", DeclaredType: "String"}, KindList},
	}
	for _, c := range cases {
		if got := c.p.Kind(); got != c.want {
			t.Errorf("%+v.Kind() = %q, want %q", c.p, got, c.want)
		}
	}
}

func TestDeltaCapable(t *testing.T) {
	cases := []struct {
		p    Param
		want bool
	}{
		{Param{Type: "System.Object", DeclaredType: "MultiValuedProperty"}, true},
		{Param{Type: "System.Object", DeclaredType: "<MultiValuedProperty>"}, true},
		{Param{Type: "System.Object[]", DeclaredType: "MultiValuedProperty"}, false},
		{Param{Type: "System.Object", DeclaredType: "ProxyAddressCollection"}, false},
		{Param{Type: "System.Object", DeclaredType: "String[]"}, false},
		{Param{Type: "System.Object"}, false},
	}
	for _, c := range cases {
		if got := c.p.DeltaCapable(); got != c.want {
			t.Errorf("%+v.DeltaCapable() = %v, want %v", c.p, got, c.want)
		}
	}
}

func TestCatalogsAreAnnotated(t *testing.T) {
	for name, load := range map[string]func() (*Catalog, error){"EXO": EXO, "Purview": Purview} {
		c, err := load()
		if err != nil {
			t.Fatal(err)
		}
		if c.DocsSource == "" {
			t.Errorf("%s catalog has no docsSource; run cmd/annotate-docs", name)
		}
	}
}

func TestDeclaredTypesTypeKnownParams(t *testing.T) {
	c, err := EXO()
	if err != nil {
		t.Fatal(err)
	}
	want := map[string]map[string]Kind{
		"Set-ExternalInOutlook": {"Enabled": KindBool},
		"Set-TransportConfig":   {"AllowLegacyTLSClients": KindBool},
		// docs error, corrected in declared-type-overrides.json
		"Set-AntiPhishPolicy": {"TargetedDomainProtectionAction": KindAny},
	}
	for _, cm := range c.Cmdlets {
		for _, p := range cm.Parameters {
			if k, ok := want[cm.Cmdlet][p.Name]; ok {
				if got := p.Kind(); got != k {
					t.Errorf("%s -%s (%s / %s) = %q, want %q", cm.Cmdlet, p.Name, p.Type, p.DeclaredType, got, k)
				}
				delete(want[cm.Cmdlet], p.Name)
			}
		}
	}
	for cm, ps := range want {
		for p := range ps {
			t.Errorf("%s -%s not found", cm, p)
		}
	}
}

func TestFlexStringsUnmarshal(t *testing.T) {
	var single FlexStrings
	if err := single.UnmarshalJSON([]byte(`"Only"`)); err != nil || len(single) != 1 || single[0] != "Only" {
		t.Fatalf("single: %v %v", single, err)
	}
	var many FlexStrings
	if err := many.UnmarshalJSON([]byte(`["A","B"]`)); err != nil || len(many) != 2 {
		t.Fatalf("many: %v %v", many, err)
	}
	var none FlexStrings
	if err := none.UnmarshalJSON([]byte(`null`)); err != nil || none != nil {
		t.Fatalf("null: %v %v", none, err)
	}
}
