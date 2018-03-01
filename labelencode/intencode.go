package labelencode

import (
	"github.com/averagesecurityguy/structures/set"
)

// IntEncoder is an label encoder for integers.
type IntEncode struct {
	Attributes *set.IntSet
	encode     map[int]int
	decode     map[int]int
}

// Fit finds the unique attributes of the give slice of integers and creates
// a mapping of attributes to ints.
func (e *IntEncode) Fit(data []int) {
	e.Attributes = set.NewIntSet(data)

	// Build our encode and decode maps using our attributes.
	members := e.Attributes.Members()
	for i := range members {
		e.encode[members[i]] = i
		e.decode[i] = members[i]
	}
}

// Encode converts a slice of strings to a slice of ints using the attributes
// of the encoder.
func (e *IntEncode) Encode(data []int) ([]int, error) {
	var encoded []int

	if err := checkIntFit(data, e.Attributes); err != nil {
		return encoded, err
	}

	for i := range data {
		encoded = append(encoded, e.encode[data[i]])
	}

	return encoded, nil
}

// Decode converts a slice of uint64 to a slice of strings using the
// attributes of the encoder.
func (e *IntEncode) Decode(encoded []int) ([]int, error) {
	var decoded []int

	if err := checkIntFit(encoded, e.decodeSet()); err != nil {
		return decoded, err
	}

	for i := range encoded {
		decoded = append(decoded, e.decode[encoded[i]])
	}

	return decoded, nil
}

func (e *IntEncode) decodeSet() *set.IntSet {
	keys := make([]int, len(e.decode))

	for k := range e.decode {
		keys = append(keys, k)
	}

	return set.NewIntSet(keys)
}
