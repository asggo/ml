package labelencode

import (
	"sort"
	"testing"
)

func TestIntEncode(t *testing.T) {
	var data = []int{1, 2, 2, 6}
	var attr = []int{1, 2, 6}
	var decoded = []int{1, 1, 2, 6}
	var encoded = []int{0, 0, 1, 2}

	e := NewIntEncode()

	e.Fit(data)
	if !equalIntSlices(e.Attributes.Members(), attr) {
		t.Error("Expected", attr, "got", e.Attributes.Members())
	}

	enc, err := e.Encode(decoded)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if !equalIntSlices(encoded, enc) {
		t.Error("Expected", encoded, "got", enc)
	}

	dec, err := e.Decode(enc)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if !equalIntSlices(decoded, dec) {
		t.Error("Expected", data, "got", dec)
	}

	_, err = e.Encode([]int{1, 4, 4})
	if err == nil {
		t.Error("Expected NotFit error got nil")
	}

	_, err = e.Decode([]int{3, 0, 0})
	if err == nil {
		t.Error("Expected NotFit error got nil")
	}
}

func equalIntSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	sort.Ints(a)
	sort.Ints(b)

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
