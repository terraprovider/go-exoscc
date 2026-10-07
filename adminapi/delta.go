package adminapi

import "encoding/json"

// StringDelta adds and removes values of a multi-valued property instead of
// replacing it — the Admin API form of PowerShell's -Param @{Add=...;Remove=...}.
// Sending an empty list does not clear a MultiValuedProperty; Remove does.
//
// It is sent the way the module's proxy psm1 sends a [Hashtable] parameter value
// (ConvertTo-HashTable): a JSON object annotated as Exchange.GenericHashTable.
// Empty Add/Remove are omitted.
type StringDelta struct {
	Add    []string
	Remove []string
}

// MarshalJSON emits {"@odata.type":"#Exchange.GenericHashTable","Remove":["a"]}.
func (d StringDelta) MarshalJSON() ([]byte, error) {
	v := struct {
		Type   string   `json:"@odata.type"`
		Add    []string `json:"Add,omitempty"`
		Remove []string `json:"Remove,omitempty"`
	}{"#Exchange.GenericHashTable", d.Add, d.Remove}
	return json.Marshal(v)
}
