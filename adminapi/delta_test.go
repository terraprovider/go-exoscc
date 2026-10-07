package adminapi

import (
	"encoding/json"
	"testing"
)

func TestStringDeltaMarshal(t *testing.T) {
	cases := []struct {
		d    StringDelta
		want string
	}{
		{StringDelta{Remove: []string{"a"}}, `{"@odata.type":"#Exchange.GenericHashTable","Remove":["a"]}`},
		{StringDelta{Add: []string{"b"}, Remove: []string{"a"}}, `{"@odata.type":"#Exchange.GenericHashTable","Add":["b"],"Remove":["a"]}`},
	}
	for _, c := range cases {
		// As it is sent: a value inside the Parameters map.
		b, err := json.Marshal(map[string]any{"AllowList": c.d})
		if err != nil {
			t.Fatal(err)
		}
		if want := `{"AllowList":` + c.want + `}`; string(b) != want {
			t.Errorf("got %s, want %s", b, want)
		}
	}
}
