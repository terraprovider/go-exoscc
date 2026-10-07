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
	d := strings.TrimSuffix(strings.TrimSpace(declared), "?")
	if d == "" {
		return KindAny
	}
	if strings.HasSuffix(d, "[]") {
		return KindList
	}
	if i := strings.LastIndex(d, "."); i >= 0 {
		d = d[i+1:]
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

func isIntType(t string) bool {
	t = strings.TrimPrefix(t, "system.")
	switch t {
	case "int", "uint", "short", "ushort", "long", "ulong",
		"int16", "int32", "int64", "uint16", "uint32", "uint64":
		return true
	}
	return false
}
