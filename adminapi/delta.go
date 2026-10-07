package adminapi

import "encoding/json"

// StringDelta adds and removes values of a multi-valued property instead of
// replacing it — the Admin API form of PowerShell's -Param @{Add=...;Remove=...}.
// Sending an empty list does not clear a MultiValuedProperty; Remove does.
//
// It is sent as the $metadata complex type Exchange.StringFieldDeltaUpdateData.
// Empty Add/Remove are omitted.
type StringDelta struct {
	Add    []string
	Remove []string
}

// MarshalJSON emits the OData-annotated complex value, e.g.
// {"@odata.type":"#Exchange.StringFieldDeltaUpdateData","Remove":["a"]}.
// CmdletInput.Parameters is an open type, so a dynamic complex value carries its
// type annotation.
func (d StringDelta) MarshalJSON() ([]byte, error) {
	v := struct {
		Type   string   `json:"@odata.type"`
		Add    []string `json:"Add,omitempty"`
		Remove []string `json:"Remove,omitempty"`
	}{"#Exchange.StringFieldDeltaUpdateData", d.Add, d.Remove}
	return json.Marshal(v)
}
