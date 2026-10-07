package exo

import (
	"reflect"
	"testing"
)

// Typed scalars and lists are bound whenever the caller sets them, so false, 0
// and an empty list (to clear a multi-valued property) reach the API; unset
// fields are omitted.
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
