package spec

import "strings"

// Kind classifies a parameter for code generation. It is the single mapping
// shared by cmd/gen-go and downstream generators (e.g. terraform-provider-exo's
// gen-tf), so the Go binding type and the Terraform attribute type agree.
type Kind string

const (
	KindSwitch Kind = "switch" // bool; sent only when true
	KindBool   Kind = "bool"   // *bool; sent when non-nil (false is sendable)
	KindInt64  Kind = "int64"  // *int64; sent when non-nil (0 is sendable)
	KindString Kind = "string" // string; sent when non-empty
	KindList   Kind = "list"   // []string; sent when non-nil (empty clears)
	KindAny    Kind = "any"    // any; sent when non-nil
)

// Kind returns the parameter's binding kind. A concrete psm1 type constraint
// always wins; DeclaredType only refines parameters the psm1 types as
// System.Object.
func (p Param) Kind() Kind {
	if p.IsSwitch {
		return KindSwitch
	}
	t := strings.ToLower(p.Type)
	switch {
	case strings.HasSuffix(t, "[]"), strings.HasPrefix(t, "system.collections.generic.list["):
		return KindList
	case t == "string" || t == "system.string" || t == "guid" || t == "system.guid":
		return KindString
	case t == "bool" || t == "system.boolean":
		return KindBool
	case isIntType(t):
		return KindInt64
	case t == "system.object":
		return declaredKind(p.DeclaredType)
	default:
		return KindAny
	}
}

// declaredKind maps a docs type (e.g. "System.Boolean", "MultiValuedProperty",
// "RecipientIdParameter[]") onto a Kind. Unmapped types (Unlimited,
// ByteQuantifiedSize, EnhancedTimeSpan, enums, *IdParameter, …) stay KindAny.
func declaredKind(declared string) Kind {
	d := normalizeDeclared(declared)
	if d == "" {
		return KindAny
	}
	if strings.HasSuffix(d, "[]") {
		return KindList
	}
	switch {
	case d == "Boolean":
		return KindBool
	case isIntType(strings.ToLower(d)):
		return KindInt64
	case d == "MultiValuedProperty", strings.HasSuffix(d, "Collection"):
		return KindList
	default:
		return KindAny
	}
}

// normalizeDeclared strips docs noise from a declared type: surrounding
// whitespace and angle brackets ("<MultiValuedProperty>"), a nullable "?", and
// the namespace ("System.Boolean" -> "Boolean").
func normalizeDeclared(declared string) string {
	d := strings.Trim(strings.TrimSpace(declared), "<>")
	d = strings.TrimSuffix(d, "?")
	if i := strings.LastIndex(d, "."); i >= 0 && !strings.HasSuffix(d, "[]") {
		d = d[i+1:]
	}
	return d
}

// DeltaCapable reports whether the parameter is a MultiValuedProperty that the
// psm1 passes through untyped (System.Object), so it can carry an Add/Remove
// delta (adminapi.StringDelta) as well as a full list. cmd/gen-go adds a
// <Field>Delta companion for these on Set-* cmdlets.
func (p Param) DeltaCapable() bool {
	return !p.IsSwitch && strings.EqualFold(p.Type, "System.Object") &&
		normalizeDeclared(p.DeclaredType) == "MultiValuedProperty"
}

func isIntType(t string) bool {
	t = strings.TrimPrefix(t, "system.")
	switch t {
	case "int", "uint", "short", "ushort", "long", "ulong",
		"int16", "int32", "int64", "uint16", "uint32", "uint64":
		return true
	}
	return false
}
