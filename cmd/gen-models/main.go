// Command gen-models turns the Admin API OData $metadata (CSDL) into Go types:
// entity types and complex types become structs (json-tagged with the OData
// property names, so adminapi.Result.Decode fills them), and enum types become
// string constants. Re-runnable; output is gofmt'd and marked DO NOT EDIT.
//
//	go run ./cmd/gen-models -metadata spec/metadata/EXO-metadata.xml \
//	    -pkg models -out models/zz_generated_models.go
package main

import (
	"bytes"
	"encoding/xml"
	"flag"
	"fmt"
	"go/format"
	"os"
	"sort"
	"strings"
)

type edmx struct {
	Schemas []schema `xml:"DataServices>Schema"`
}
type schema struct {
	Namespace string      `xml:"Namespace,attr"`
	Enums     []enumType  `xml:"EnumType"`
	Complex   []structDef `xml:"ComplexType"`
	Entities  []structDef `xml:"EntityType"`
}
type enumType struct {
	Name    string `xml:"Name,attr"`
	Members []struct {
		Name string `xml:"Name,attr"`
	} `xml:"Member"`
}
type structDef struct {
	Name       string     `xml:"Name,attr"`
	BaseType   string     `xml:"BaseType,attr"`
	Abstract   string     `xml:"Abstract,attr"`
	Properties []property `xml:"Property"`
	NavProps   []property `xml:"NavigationProperty"`
}
type property struct {
	Name string `xml:"Name,attr"`
	Type string `xml:"Type,attr"`
}

func main() {
	var mdPath, pkg, out string
	flag.StringVar(&mdPath, "metadata", "", "path to $metadata CSDL xml")
	flag.StringVar(&pkg, "pkg", "models", "output package name")
	flag.StringVar(&out, "out", "", "output .go file")
	flag.Parse()
	if mdPath == "" || out == "" {
		fmt.Fprintln(os.Stderr, "gen-models: -metadata and -out are required")
		os.Exit(2)
	}
	raw, err := os.ReadFile(mdPath)
	must(err)
	var doc edmx
	must(xml.Unmarshal(raw, &doc))

	// Classify every defined type so references resolve correctly: enum refs are
	// value (string) types, struct refs are pointers (avoids invalid recursive
	// value types like a self-referential entity), unknown refs fall back to raw JSON.
	kinds := map[string]string{} // name -> "enum" | "struct"
	var enums []enumType
	var structs []structDef
	for _, s := range doc.Schemas {
		for _, e := range s.Enums {
			kinds[e.Name] = "enum"
			enums = append(enums, e)
		}
		for _, c := range s.Complex {
			kinds[c.Name] = "struct"
			structs = append(structs, c)
		}
		for _, e := range s.Entities {
			kinds[e.Name] = "struct"
			structs = append(structs, e)
		}
	}
	sort.Slice(enums, func(i, j int) bool { return enums[i].Name < enums[j].Name })
	sort.Slice(structs, func(i, j int) bool { return structs[i].Name < structs[j].Name })

	var b bytes.Buffer
	fmt.Fprintf(&b, "// Code generated from %s by gen-models. DO NOT EDIT.\n\n", trimPath(mdPath))
	fmt.Fprintf(&b, "package %s\n\n", pkg)
	fmt.Fprintf(&b, "import (\n\t%q\n\t%q\n)\n\n", "encoding/json", "time")
	fmt.Fprintf(&b, "var _ = time.Time{}\nvar _ json.RawMessage\n\n")

	for _, e := range enums {
		fmt.Fprintf(&b, "// %s enum.\ntype %s string\n\nconst (\n", e.Name, e.Name)
		for _, m := range e.Members {
			fmt.Fprintf(&b, "\t%s_%s %s = %q\n", e.Name, exportName(m.Name), e.Name, m.Name)
		}
		fmt.Fprintf(&b, ")\n\n")
	}

	for _, s := range structs {
		fmt.Fprintf(&b, "// %s (OData %s).\ntype %s struct {\n", s.Name, kindOf(s), s.Name)
		if bt := localName(s.BaseType); bt != "" && kinds[bt] == "struct" {
			fmt.Fprintf(&b, "\t%s\n", bt) // embed base type
		}
		seen := map[string]bool{}
		for _, p := range append(append([]property{}, s.Properties...), s.NavProps...) {
			field := exportName(p.Name)
			if field == "" || seen[field] {
				continue
			}
			seen[field] = true
			fmt.Fprintf(&b, "\t%s %s `json:%q`\n", field, goType(p.Type, kinds), p.Name+",omitempty")
		}
		fmt.Fprintf(&b, "}\n\n")
	}

	src, err := format.Source(b.Bytes())
	if err != nil {
		_ = os.WriteFile(out+".unformatted", b.Bytes(), 0o644)
		must(fmt.Errorf("gofmt: %w (wrote %s.unformatted)", err, out))
	}
	must(os.WriteFile(out, src, 0o644))
	fmt.Printf("gen-models: %d enums, %d structs -> %s\n", len(enums), len(structs), out)
}

func kindOf(s structDef) string {
	if len(s.NavProps) > 0 {
		return "EntityType"
	}
	return "ComplexType"
}

// goType maps an Edm/OData type reference to a Go type. kinds maps a local type
// name to "enum" or "struct".
func goType(t string, kinds map[string]string) string {
	if strings.HasPrefix(t, "Collection(") {
		inner := strings.TrimSuffix(strings.TrimPrefix(t, "Collection("), ")")
		return "[]" + goType(inner, kinds)
	}
	if strings.HasPrefix(t, "Edm.") {
		switch t {
		case "Edm.String", "Edm.Guid":
			return "string"
		case "Edm.Boolean":
			return "bool"
		case "Edm.Int16", "Edm.Int32":
			return "int32"
		case "Edm.Int64":
			return "int64"
		case "Edm.Byte":
			return "byte"
		case "Edm.Double", "Edm.Single", "Edm.Decimal":
			return "float64"
		case "Edm.DateTimeOffset", "Edm.Date":
			return "time.Time"
		case "Edm.Binary":
			return "[]byte"
		default:
			return "json.RawMessage"
		}
	}
	// Named type (Exchange.Foo): enum -> value string type; struct -> pointer
	// (idiomatic for nested objects and avoids invalid recursive value types).
	switch n := localName(t); kinds[n] {
	case "enum":
		return n
	case "struct":
		return "*" + n
	default:
		return "json.RawMessage"
	}
}

func localName(t string) string {
	if t == "" {
		return ""
	}
	if i := strings.LastIndex(t, "."); i >= 0 {
		return t[i+1:]
	}
	return t
}

func exportName(s string) string {
	s = strings.Map(func(r rune) rune {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
			return r
		}
		return -1
	}, s)
	if s == "" {
		return ""
	}
	if s[0] >= '0' && s[0] <= '9' {
		s = "N" + s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func trimPath(p string) string {
	if i := strings.LastIndexAny(p, "/\\"); i >= 0 {
		return p[i+1:]
	}
	return p
}

func must(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "gen-models:", err)
		os.Exit(1)
	}
}
