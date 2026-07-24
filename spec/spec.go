// Package spec exposes the cmdlet catalogs that drive code generation. The
// catalogs are extracted from the ExchangeOnlineManagement PowerShell module's
// own command definitions (see generator/extract-catalog.ps1) and embedded here
// so every consumer — the Go binding generator (cmd/gen-go) and the Terraform
// resource generator (terraform-provider-exo/cmd/gen-tf) — reads the same typed,
// versioned source of truth.
package spec

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"sort"
)

//go:embed catalog/EXO-catalog.json
var exoJSON []byte

//go:embed catalog/Purview-catalog.json
var purviewJSON []byte

// Catalog is a parsed cmdlet catalog for one service.
type Catalog struct {
	Source      string   `json:"source"`
	CmdletCount int      `json:"cmdletCount"`
	Cmdlets     []Cmdlet `json:"cmdlets"`
}

// Cmdlet is a single PowerShell cmdlet (a Verb-Noun) and its parameters.
type Cmdlet struct {
	Cmdlet              string  `json:"cmdlet"`
	Verb                string  `json:"verb"`
	Noun                string  `json:"noun"`
	DefaultParameterSet string  `json:"defaultParameterSet"`
	Parameters          []Param `json:"parameters"`
}

// Param is one cmdlet parameter.
type Param struct {
	Name          string      `json:"name"`
	Type          string      `json:"type"`
	IsSwitch      bool        `json:"isSwitch"`
	ParameterSets []ParamSet  `json:"parameterSets"`
	ValidateSet   FlexStrings `json:"validateSet"`
	Aliases       FlexStrings `json:"aliases"`
}

// ParamSet is a parameter's membership in one PowerShell parameter set.
type ParamSet struct {
	Name      string `json:"name"`
	Mandatory bool   `json:"mandatory"`
	Position  *int   `json:"position"`
}

// Mandatory reports whether the parameter is mandatory in any parameter set.
func (p Param) Mandatory() bool {
	for _, s := range p.ParameterSets {
		if s.Mandatory {
			return true
		}
	}
	return false
}

// FlexStrings unmarshals either a JSON string or a JSON array of strings —
// PowerShell's ConvertTo-Json collapses single-element arrays to a scalar.
type FlexStrings []string

func (f *FlexStrings) UnmarshalJSON(b []byte) error {
	b = bytes.TrimSpace(b)
	if len(b) == 0 || string(b) == "null" {
		return nil
	}
	if b[0] == '[' {
		var a []string
		if err := json.Unmarshal(b, &a); err != nil {
			return err
		}
		*f = a
		return nil
	}
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	*f = FlexStrings{s}
	return nil
}

// EXO returns the Exchange Online catalog.
func EXO() (*Catalog, error) { return parse(exoJSON) }

// Purview returns the Security & Compliance (Purview) catalog.
func Purview() (*Catalog, error) { return parse(purviewJSON) }

func parse(raw []byte) (*Catalog, error) {
	var c Catalog
	if err := json.Unmarshal(raw, &c); err != nil {
		return nil, err
	}
	return &c, nil
}

// ByNoun groups the catalog's cmdlets by noun, keyed verb -> cmdlet.
func (c *Catalog) ByNoun() map[string]map[string]Cmdlet {
	out := map[string]map[string]Cmdlet{}
	for _, cm := range c.Cmdlets {
		if out[cm.Noun] == nil {
			out[cm.Noun] = map[string]Cmdlet{}
		}
		out[cm.Noun][cm.Verb] = cm
	}
	return out
}

// CRUDVerbs are the verbs that make a noun a full create/read/update/delete
// resource.
var CRUDVerbs = []string{"New", "Get", "Set", "Remove"}

// CRUDComplete returns the sorted nouns that have all four CRUD verbs
// (New + Get + Set + Remove) and are therefore full Terraform-resource candidates.
func (c *Catalog) CRUDComplete() []string {
	byNoun := c.ByNoun()
	var nouns []string
	for noun, verbs := range byNoun {
		complete := true
		for _, v := range CRUDVerbs {
			if _, ok := verbs[v]; !ok {
				complete = false
				break
			}
		}
		if complete {
			nouns = append(nouns, noun)
		}
	}
	sort.Strings(nouns)
	return nouns
}
