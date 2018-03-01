package labelencode

import (
	"sort"
	"testing"
)

func TestStringEncode(t *testing.T) {
	var data = []string{"paris", "paris", "tokyo", "amsterdam"}
	var attr = []string{"amsterdam", "paris", "tokyo"}
	var decoded = []string{"tokyo", "tokyo", "paris"}
	var encoded = []int{2, 2, 1}

	e := NewStringEncode()

	e.Fit(data)
	if !equalStringSlices(e.Attributes.Members(), attr) {
		t.Error("Expected", attr, "got", e.Attributes.Members())
	}

	enc, err := e.Encode(decoded)
	if err != nil {
		t.Error(err)
	}

	if !equalIntSlices(encoded, enc) {
		t.Error("Expected", encoded, "got", enc)
	}

	dec, err := e.Decode(enc)
	if err != nil {
		t.Error(err)
	}
	if !equalStringSlices(decoded, dec) {
		t.Error("Expected", decoded, "got", dec)
	}

	_, err = e.Encode([]string{"london", "paris"})
	if err == nil {
		t.Error("Expected NotFit error got nil")
	}

	_, err = e.Decode([]int{4, 2, 1})
	if err == nil {
		t.Error("Expected NotFit error got nil")
	}
}

func equalStringSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	sort.Strings(a)
	sort.Strings(b)

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
