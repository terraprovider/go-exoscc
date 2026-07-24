// Command gen-go generates typed Go bindings from a cmdlet catalog produced by
// generator/extract-catalog.ps1 (PowerShell AST -> JSON). One Params struct and
// one *Service method per cmdlet; the body just calls adminapi.Client.Invoke.
//
// Re-runnable: on any API change, re-fetch the psm1, re-run extract-catalog.ps1,
// then re-run this. Output is gofmt'd and marked DO NOT EDIT.
//
//	go run ./cmd/gen-go -catalog spec/catalog/EXO-catalog.json \
//	    -pkg exo -client github.com/terraprovider/go-exoscc/adminapi -out exo/zz_generated_exo.go
package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"go/format"
	"os"
	"sort"
	"strings"
)

type catalog struct {
	Source      string   `json:"source"`
	CmdletCount int      `json:"cmdletCount"`
	Cmdlets     []cmdlet `json:"cmdlets"`
}
type cmdlet struct {
	Cmdlet              string  `json:"cmdlet"`
	Verb                string  `json:"verb"`
	Noun                string  `json:"noun"`
	DefaultParameterSet string  `json:"defaultParameterSet"`
	Parameters          []param `json:"parameters"`
}
type param struct {
	Name        string      `json:"name"`
	Type        string      `json:"type"`
	IsSwitch    bool        `json:"isSwitch"`
	ValidateSet flexStrings `json:"validateSet"`
	Aliases     flexStrings `json:"aliases"`
}

// flexStrings unmarshals either a JSON string or a JSON array of strings — PowerShell's
// ConvertTo-Json collapses single-element arrays to a scalar.
type flexStrings []string

func (f *flexStrings) UnmarshalJSON(b []byte) error {
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
	*f = flexStrings{s}
	return nil
}

func main() {
	var catPath, pkg, clientPath, out string
	flag.StringVar(&catPath, "catalog", "", "path to *-catalog.json")
	flag.StringVar(&pkg, "pkg", "exo", "output package name")
	flag.StringVar(&clientPath, "client", "github.com/terraprovider/go-exoscc/adminapi", "import path of the adminapi client package")
	flag.StringVar(&out, "out", "", "output .go file")
	flag.Parse()
	if catPath == "" || out == "" {
		fmt.Fprintln(os.Stderr, "gen-go: -catalog and -out are required")
		os.Exit(2)
	}

	raw, err := os.ReadFile(catPath)
	must(err)
	var cat catalog
	must(json.Unmarshal(raw, &cat))
	sort.Slice(cat.Cmdlets, func(i, j int) bool { return cat.Cmdlets[i].Cmdlet < cat.Cmdlets[j].Cmdlet })

	var b bytes.Buffer
	fmt.Fprintf(&b, "// Code generated from %s by gen-go. DO NOT EDIT.\n\n", cat.Source)
	fmt.Fprintf(&b, "package %s\n\n", pkg)
	fmt.Fprintf(&b, "import (\n\t%q\n\n\t%q\n)\n\n", "context", clientPath)
	fmt.Fprintf(&b, "// Service exposes the %d cmdlets of %s as typed methods.\n", len(cat.Cmdlets), cat.Source)
	fmt.Fprintf(&b, "type Service struct{ C *adminapi.Client }\n\n")
	fmt.Fprintf(&b, "// New wraps an *adminapi.Client.\nfunc New(c *adminapi.Client) *Service { return &Service{C: c} }\n\n")

	seenField := map[string]bool{}
	for _, cm := range cat.Cmdlets {
		emitCmdlet(&b, cm, seenField)
	}

	src, err := format.Source(b.Bytes())
	if err != nil {
		// write unformatted for debugging, then fail
		_ = os.WriteFile(out+".unformatted", b.Bytes(), 0o644)
		must(fmt.Errorf("gofmt: %w (wrote %s.unformatted)", err, out))
	}
	must(os.WriteFile(out, src, 0o644))
	fmt.Printf("gen-go: %d cmdlets -> %s\n", len(cat.Cmdlets), out)
}

func emitCmdlet(b *bytes.Buffer, cm cmdlet, _ map[string]bool) {
	method := goName(cm.Cmdlet)
	pstruct := method + "Params"

	// Params struct
	fmt.Fprintf(b, "// %sParams are the parameters of %s.\n", method, cm.Cmdlet)
	if cm.DefaultParameterSet != "" {
		fmt.Fprintf(b, "// DefaultParameterSetName: %s\n", cm.DefaultParameterSet)
	}
	fmt.Fprintf(b, "type %s struct {\n", pstruct)
	used := map[string]bool{}
	for _, p := range cm.Parameters {
		field := exportName(p.Name)
		if used[field] { // guard against rare collisions after normalization
			continue
		}
		used[field] = true
		comment := ""
		if len(p.ValidateSet) > 0 {
			comment = " // one of: " + strings.Join(p.ValidateSet, ", ")
		}
		fmt.Fprintf(b, "\t%s %s `ps:%q`%s\n", field, goType(p), p.Name, comment)
	}
	fmt.Fprintf(b, "}\n\n")

	// params() -> map of bound parameters only
	fmt.Fprintf(b, "func (p %s) params() map[string]any {\n\tm := map[string]any{}\n", pstruct)
	used = map[string]bool{}
	for _, p := range cm.Parameters {
		field := exportName(p.Name)
		if used[field] {
			continue
		}
		used[field] = true
		fmt.Fprintf(b, "\t%s\n", boundCheck(field, p.Name, goType(p)))
	}
	fmt.Fprintf(b, "\treturn m\n}\n\n")

	// Service method
	fmt.Fprintf(b, "// %s runs the %s cmdlet.\n", method, cm.Cmdlet)
	fmt.Fprintf(b, "func (s *Service) %s(ctx context.Context, p %s) (*adminapi.Result, error) {\n", method, pstruct)
	fmt.Fprintf(b, "\treturn s.C.Invoke(ctx, %q, p.params())\n}\n\n", cm.Cmdlet)
}

func boundCheck(field, psName, gotype string) string {
	switch gotype {
	case "bool":
		return fmt.Sprintf("if p.%s { m[%q] = true }", field, psName)
	case "string":
		return fmt.Sprintf("if p.%s != \"\" { m[%q] = p.%s }", field, psName, field)
	case "int":
		return fmt.Sprintf("if p.%s != 0 { m[%q] = p.%s }", field, psName, field)
	case "[]string":
		return fmt.Sprintf("if len(p.%s) > 0 { m[%q] = p.%s }", field, psName, field)
	default: // any
		return fmt.Sprintf("if p.%s != nil { m[%q] = p.%s }", field, psName, field)
	}
}

// goType maps a PowerShell parameter type to a Go type.
func goType(p param) string {
	if p.IsSwitch {
		return "bool"
	}
	t := strings.ToLower(p.Type)
	switch {
	case strings.HasSuffix(t, "[]"):
		return "[]string"
	case t == "string" || strings.HasSuffix(t, ".string"):
		return "string"
	case t == "bool" || strings.HasSuffix(t, ".boolean"):
		return "bool"
	case t == "int" || strings.HasSuffix(t, ".int32") || strings.HasSuffix(t, ".int64"):
		return "int"
	case strings.HasSuffix(t, ".guid"):
		return "string"
	default:
		return "any"
	}
}

// goName turns "Get-Mailbox" into "GetMailbox".
func goName(cmdlet string) string {
	parts := strings.FieldsFunc(cmdlet, func(r rune) bool { return r == '-' || r == '_' })
	var sb strings.Builder
	for _, p := range parts {
		sb.WriteString(exportName(p))
	}
	return sb.String()
}

// exportName makes an exported Go identifier from a PowerShell name.
func exportName(s string) string {
	s = strings.Map(func(r rune) rune {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
			return r
		}
		return -1
	}, s)
	if s == "" {
		return "X"
	}
	if s[0] >= '0' && s[0] <= '9' {
		s = "N" + s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func must(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "gen-go:", err)
		os.Exit(1)
	}
}
