// Command gen-go generates typed Go bindings from a cmdlet catalog produced by
// generator/extract-catalog.ps1 (PowerShell AST -> JSON). One Params struct and
// one *Service method per cmdlet; the body just calls adminapi.Client.Invoke.
//
// Field types come from spec.Param.Kind, so run cmd/annotate-docs on the catalog
// first to type the psm1's System.Object parameters. Re-runnable: on any API
// change, re-fetch the psm1, re-run extract-catalog.ps1 and annotate-docs, then
// re-run this. Output is gofmt'd and marked DO NOT EDIT.
//
//	go run ./cmd/gen-go -catalog spec/catalog/EXO-catalog.json \
//	    -pkg exo -client github.com/terraprovider/go-exoscc/adminapi -out exo/zz_generated_exo.go
package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"os"
	"sort"
	"strings"

	"github.com/terraprovider/go-exoscc/spec"
)

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
	cat, err := spec.Parse(bytes.TrimPrefix(raw, []byte("\xef\xbb\xbf")))
	must(err)
	sort.Slice(cat.Cmdlets, func(i, j int) bool { return cat.Cmdlets[i].Cmdlet < cat.Cmdlets[j].Cmdlet })

	var b bytes.Buffer
	from := cat.Source
	if cat.DocsSource != "" {
		from += " and " + cat.DocsSource
	}
	fmt.Fprintf(&b, "// Code generated from %s by gen-go. DO NOT EDIT.\n\n", from)
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

func emitCmdlet(b *bytes.Buffer, cm spec.Cmdlet, _ map[string]bool) {
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
		fmt.Fprintf(b, "\t%s %s `ps:%q`%s\n", field, goType(p), p.Name, fieldComment(p))
		if hasDelta(cm, p) {
			fmt.Fprintf(b, "\t%sDelta *adminapi.StringDelta `ps:%q` // adds/removes values of %s; takes precedence over it\n", field, p.Name, field)
		}
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
		if hasDelta(cm, p) {
			fmt.Fprintf(b, "\t%s\n", deltaBoundCheck(field, p.Name))
			continue
		}
		fmt.Fprintf(b, "\t%s\n", boundCheck(field, p.Name, goType(p)))
	}
	fmt.Fprintf(b, "\treturn m\n}\n\n")

	// Service method
	fmt.Fprintf(b, "// %s runs the %s cmdlet.\n", method, cm.Cmdlet)
	fmt.Fprintf(b, "func (s *Service) %s(ctx context.Context, p %s) (*adminapi.Result, error) {\n", method, pstruct)
	fmt.Fprintf(b, "\treturn s.C.Invoke(ctx, %q, p.params())\n}\n\n", cm.Cmdlet)
}

// fieldComment documents what the Go type can't: the allowed values, and the
// declared .NET type of an untyped (any) parameter.
func fieldComment(p spec.Param) string {
	var notes []string
	if len(p.ValidateSet) > 0 {
		notes = append(notes, "one of: "+strings.Join(p.ValidateSet, ", "))
	}
	if p.Kind() == spec.KindAny && p.DeclaredType != "" {
		notes = append(notes, p.DeclaredType)
	}
	if len(notes) == 0 {
		return ""
	}
	return " // " + strings.Join(notes, "; ")
}

// boundCheck emits the params() line that binds a field only when the caller set
// it. Pointers and slices are bound whenever non-nil, so false, 0 and an empty
// list (which clears a multi-valued property) can be sent.
func boundCheck(field, psName, gotype string) string {
	switch gotype {
	case "bool": // switch
		return fmt.Sprintf("if p.%s { m[%q] = true }", field, psName)
	case "string":
		return fmt.Sprintf("if p.%s != \"\" { m[%q] = p.%s }", field, psName, field)
	case "*bool", "*int64":
		return fmt.Sprintf("if p.%s != nil { m[%q] = *p.%s }", field, psName, field)
	default: // []string, any
		return fmt.Sprintf("if p.%s != nil { m[%q] = p.%s }", field, psName, field)
	}
}

// hasDelta reports whether a parameter gets a <Field>Delta companion: a
// delta-capable MultiValuedProperty on a Set-* cmdlet (where clearing or
// shrinking a list needs Remove; New-* always sends the full list).
func hasDelta(cm spec.Cmdlet, p spec.Param) bool {
	return strings.EqualFold(cm.Verb, "Set") && p.DeltaCapable()
}

// deltaBoundCheck binds the delta when set, else the full list.
func deltaBoundCheck(field, psName string) string {
	return fmt.Sprintf("if p.%sDelta != nil { m[%q] = *p.%sDelta } else if p.%s != nil { m[%q] = p.%s }",
		field, psName, field, field, psName, field)
}

// goType maps a parameter's spec.Kind to its Go field type.
func goType(p spec.Param) string {
	switch p.Kind() {
	case spec.KindSwitch:
		return "bool"
	case spec.KindBool:
		return "*bool"
	case spec.KindInt64:
		return "*int64"
	case spec.KindString:
		return "string"
	case spec.KindList:
		return "[]string"
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
