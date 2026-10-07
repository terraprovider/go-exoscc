// Command annotate-docs records each parameter's declared .NET type in a cmdlet
// catalog. The psm1 proxy module types most parameters as System.Object; the
// cmdlet reference in MicrosoftDocs/office-docs-powershell
// (exchange/exchange-ps/ExchangePowerShell/<Cmdlet>.md) gives the real type in
// the yaml block of every "### -<Param>" section. That type is written to the
// catalog as "declaredType", and spec.Param.Kind uses it to type the bindings.
//
// A System.Object parameter with no documented type inherits the declared type
// of the same-named parameter on the noun's other New/Set cmdlet, so one
// Terraform attribute that writes both gets one type.
//
// Known docs errors are corrected in spec/declared-type-overrides.json. The docs
// commit is pinned in spec/docs-ref and recorded in the catalog's "docsSource".
// Re-runnable and idempotent; run it after extract-catalog.ps1 (tools/regen.sh
// does):
//
//	./tools/fetch-docs.sh .spec-cache/docs
//	go run ./cmd/annotate-docs -catalog spec/catalog/EXO-catalog.json \
//	    -docs .spec-cache/docs/exchange/exchange-ps/ExchangePowerShell \
//	    -docs-ref "$(cat spec/docs-ref)" -overrides spec/declared-type-overrides.json
package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// The catalog is round-tripped through these structs. Field order matches
// extract-catalog.ps1's output, and fields the generator does not interpret stay
// raw, so re-encoding is byte-identical apart from the added annotations.
type catalog struct {
	Source        string   `json:"source"`
	CmdletCount   int      `json:"cmdletCount"`
	GeneratedFrom string   `json:"generatedFrom"`
	DocsSource    string   `json:"docsSource,omitempty"`
	Cmdlets       []cmdlet `json:"cmdlets"`
}

type cmdlet struct {
	Cmdlet              string          `json:"cmdlet"`
	Verb                string          `json:"verb"`
	Noun                string          `json:"noun"`
	DefaultParameterSet json.RawMessage `json:"defaultParameterSet"`
	Parameters          []param         `json:"parameters"`
}

type param struct {
	Name          string          `json:"name"`
	Type          string          `json:"type"`
	DeclaredType  string          `json:"declaredType,omitempty"`
	IsSwitch      bool            `json:"isSwitch"`
	ParameterSets json.RawMessage `json:"parameterSets"`
	ValidateSet   json.RawMessage `json:"validateSet"`
	Aliases       json.RawMessage `json:"aliases"`
}

// override corrects one parameter whose documented type is wrong. Catalog
// scopes it to one catalog ("EXO" or "Purview", the prefix of the catalog's
// source), so an override that matches nothing in its catalog is caught.
type override struct {
	Catalog   string `json:"catalog"`
	Cmdlet    string `json:"cmdlet"`
	Parameter string `json:"parameter"`
	Type      string `json:"type"`
	Reason    string `json:"reason"`
}

func main() {
	var catPath, docsDir, docsRef, ovPath string
	flag.StringVar(&catPath, "catalog", "", "path to *-catalog.json (rewritten in place)")
	flag.StringVar(&docsDir, "docs", "", "office-docs-powershell exchange/exchange-ps/ExchangePowerShell directory")
	flag.StringVar(&docsRef, "docs-ref", "", "office-docs-powershell commit the -docs checkout is at")
	flag.StringVar(&ovPath, "overrides", "spec/declared-type-overrides.json", "docs-error overrides")
	flag.Parse()
	if catPath == "" || docsDir == "" || docsRef == "" {
		fmt.Fprintln(os.Stderr, "annotate-docs: -catalog, -docs and -docs-ref are required")
		os.Exit(2)
	}

	raw, err := os.ReadFile(catPath)
	must(err)
	var cat catalog
	must(json.Unmarshal(bytes.TrimPrefix(raw, []byte("\xef\xbb\xbf")), &cat))

	var ovs []override
	ovRaw, err := os.ReadFile(ovPath)
	must(err)
	must(json.Unmarshal(ovRaw, &ovs))

	pages, err := indexPages(docsDir)
	must(err)

	stats, err := annotate(&cat, pages, ovs)
	must(err)
	cat.DocsSource = "MicrosoftDocs/office-docs-powershell@" + docsRef

	out, err := encode(cat)
	must(err)
	must(os.WriteFile(catPath, out, 0o644))
	fmt.Printf("annotate-docs: %s: %d/%d params typed (%d inherited New<->Set), %d cmdlets without a docs page, %d overrides\n",
		catPath, stats.typed, stats.params, stats.inherited, len(stats.noPage), stats.overridden)
}

// encode writes the catalog in extract-catalog.ps1's layout (2-space indent, no
// HTML escaping, trailing newline).
func encode(cat catalog) ([]byte, error) {
	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "  ")
	err := enc.Encode(cat)
	return b.Bytes(), err
}

type stats struct {
	params, typed, overridden, inherited int
	noPage                               []string
}

// annotate sets DeclaredType on every parameter from its docs page (or an
// override). An override that matches no parameter in its catalog is an error,
// so the list can't
// silently go stale.
func annotate(cat *catalog, pages map[string]string, ovs []override) (stats, error) {
	var st stats
	name := cat.name()
	if name == "" {
		return st, fmt.Errorf("catalog source %q has no <Catalog>- prefix", cat.Source)
	}
	var mine []override
	for _, ov := range ovs {
		if ov.Catalog == "" {
			return st, fmt.Errorf("override %s -%s has no catalog", ov.Cmdlet, ov.Parameter)
		}
		if strings.EqualFold(ov.Catalog, name) {
			mine = append(mine, ov)
		}
	}
	ovs = mine
	used := make([]bool, len(ovs))
	for i := range cat.Cmdlets {
		cm := &cat.Cmdlets[i]
		var types map[string]string
		if path, ok := pages[strings.ToLower(cm.Cmdlet)]; ok {
			var err error
			if types, err = parsePage(path); err != nil {
				return st, err
			}
		} else {
			st.noPage = append(st.noPage, cm.Cmdlet)
		}
		for j := range cm.Parameters {
			p := &cm.Parameters[j]
			st.params++
			p.DeclaredType = types[strings.ToLower(p.Name)]
			for k, ov := range ovs {
				if strings.EqualFold(ov.Cmdlet, cm.Cmdlet) && strings.EqualFold(ov.Parameter, p.Name) {
					p.DeclaredType = ov.Type
					used[k] = true
					st.overridden++
				}
			}
			if p.DeclaredType != "" {
				st.typed++
			}
		}
	}
	st.inherited = inheritNewSet(cat)
	st.typed += st.inherited
	for k, u := range used {
		if !u {
			return st, fmt.Errorf("override %s -%s matches no parameter in the %s catalog", ovs[k].Cmdlet, ovs[k].Parameter, name)
		}
	}
	return st, nil
}

// inheritNewSet gives an undocumented System.Object parameter of New-<Noun> or
// Set-<Noun> the declared type of the same-named parameter on the other verb.
// Sources are only parameters typed from the docs or an override: a target is
// untyped by definition, so nothing is inherited twice.
func inheritNewSet(cat *catalog) int {
	other := map[string]string{"new": "set", "set": "new"}
	declared := map[string]string{} // "verb noun param" (lower) -> declared type
	for _, cm := range cat.Cmdlets {
		if _, ok := other[strings.ToLower(cm.Verb)]; !ok {
			continue
		}
		for _, p := range cm.Parameters {
			if p.DeclaredType != "" {
				declared[strings.ToLower(cm.Verb+" "+cm.Noun+" "+p.Name)] = p.DeclaredType
			}
		}
	}
	n := 0
	for i := range cat.Cmdlets {
		cm := &cat.Cmdlets[i]
		ov, ok := other[strings.ToLower(cm.Verb)]
		if !ok {
			continue
		}
		for j := range cm.Parameters {
			p := &cm.Parameters[j]
			if p.DeclaredType != "" || !strings.EqualFold(p.Type, "System.Object") {
				continue
			}
			if t := declared[strings.ToLower(ov+" "+cm.Noun+" "+p.Name)]; t != "" {
				p.DeclaredType = t
				n++
			}
		}
	}
	return n
}

// name is the catalog's service name, the prefix of its source
// ("EXO-ExchangeOnline.psm1" -> "EXO").
func (c *catalog) name() string {
	name, _, ok := strings.Cut(c.Source, "-")
	if !ok {
		return ""
	}
	return name
}

// indexPages maps lower-cased cmdlet name -> docs page path.
func indexPages(dir string) (map[string]string, error) {
	ents, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	out := map[string]string{}
	for _, e := range ents {
		if name, ok := strings.CutSuffix(e.Name(), ".md"); ok && !e.IsDir() {
			out[strings.ToLower(name)] = filepath.Join(dir, e.Name())
		}
	}
	return out, nil
}

var (
	headingRe = regexp.MustCompile(`^### -(\S+)\s*$`)
	typeRe    = regexp.MustCompile(`^Type:\s*(\S+)\s*$`)
)

// parsePage returns lower-cased parameter name -> declared type. A parameter
// documented in several parameter sets takes the type of its first yaml block.
func parsePage(path string) (map[string]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return parseTypes(bufio.NewScanner(f))
}

func parseTypes(sc *bufio.Scanner) (map[string]string, error) {
	out := map[string]string{}
	cur := ""
	for sc.Scan() {
		line := sc.Text()
		if m := headingRe.FindStringSubmatch(line); m != nil {
			cur = strings.ToLower(m[1])
			continue
		}
		if strings.HasPrefix(line, "## ") || strings.HasPrefix(line, "### ") {
			cur = "" // left the parameter section (e.g. "### CommonParameters")
			continue
		}
		if m := typeRe.FindStringSubmatch(line); m != nil && cur != "" {
			if _, seen := out[cur]; !seen {
				out[cur] = m[1]
			}
		}
	}
	return out, sc.Err()
}

func must(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "annotate-docs:", err)
		os.Exit(1)
	}
}
