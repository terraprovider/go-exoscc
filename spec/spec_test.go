package spec

import "testing"

func TestEXOCatalogLoads(t *testing.T) {
	c, err := EXO()
	if err != nil {
		t.Fatal(err)
	}
	if c.CmdletCount == 0 || len(c.Cmdlets) == 0 {
		t.Fatalf("empty catalog: count=%d cmdlets=%d", c.CmdletCount, len(c.Cmdlets))
	}
	if len(c.Cmdlets) != c.CmdletCount {
		t.Errorf("cmdletCount=%d but %d cmdlets", c.CmdletCount, len(c.Cmdlets))
	}
}

func TestPurviewCatalogLoads(t *testing.T) {
	if _, err := Purview(); err != nil {
		t.Fatal(err)
	}
}

func TestRoleGroupIsCRUDComplete(t *testing.T) {
	c, err := EXO()
	if err != nil {
		t.Fatal(err)
	}
	byNoun := c.ByNoun()
	rg, ok := byNoun["RoleGroup"]
	if !ok {
		t.Fatal("RoleGroup noun missing")
	}
	for _, v := range CRUDVerbs {
		if _, ok := rg[v]; !ok {
			t.Errorf("RoleGroup missing verb %s", v)
		}
	}
	found := false
	for _, n := range c.CRUDComplete() {
		if n == "RoleGroup" {
			found = true
		}
	}
	if !found {
		t.Error("RoleGroup not reported CRUD-complete")
	}
	t.Logf("CRUD-complete nouns: %d", len(c.CRUDComplete()))
}

func TestNewRoleGroupParams(t *testing.T) {
	c, _ := EXO()
	nr := c.ByNoun()["RoleGroup"]["New"]
	var hasName, hasDesc bool
	for _, p := range nr.Parameters {
		switch p.Name {
		case "Name":
			hasName = true
		case "Description":
			hasDesc = true
		}
	}
	if !hasName || !hasDesc {
		t.Errorf("New-RoleGroup params: name=%v desc=%v", hasName, hasDesc)
	}
}
