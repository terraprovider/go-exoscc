package exo

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/terraprovider/go-exoscc/adminapi"
)

// Typed scalars and lists are bound whenever the caller sets them, so false, 0
// and an empty list reach the API; unset fields are omitted.
func TestParamsSendsZeroValues(t *testing.T) {
	off := false
	got := SetExternalInOutlookParams{Enabled: &off, AllowList: []string{}}.params()
	want := map[string]any{"Enabled": false, "AllowList": []string{}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("params() = %#v, want %#v", got, want)
	}
	if got := (SetExternalInOutlookParams{}).params(); len(got) != 0 {
		t.Errorf("unset params() = %#v, want empty", got)
	}
}

// A delta takes precedence over the full list and goes out as the annotated
// StringFieldDeltaUpdateData complex value.
func TestParamsSendsDelta(t *testing.T) {
	p := SetExternalInOutlookParams{
		AllowList:      []string{"ignored"},
		AllowListDelta: &adminapi.StringDelta{Remove: []string{"a"}},
	}
	b, err := json.Marshal(p.params())
	if err != nil {
		t.Fatal(err)
	}
	want := `{"AllowList":{"@odata.type":"#Exchange.StringFieldDeltaUpdateData","Remove":["a"]}}`
	if string(b) != want {
		t.Errorf("params() = %s, want %s", b, want)
	}
}
