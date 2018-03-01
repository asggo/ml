package labelencode

import (
	"github.com/averagesecurityguy/structures/set"
)

type StringEncode struct {
	Attributes *set.StringSet
	encode     map[string]int
	decode     map[int]string
}

// Fit finds the unique attributes of the give slice of strings and creates
// a mapping of attributes to ints.
func (s *StringEncode) Fit(data []string) {
	// Unique and sort the data to get its attributes.
	s.Attributes = set.NewStringSet(data)

	// Build our encoder and decoder using our attributes.
	members := s.Attributes.Members()
	for i := range members {
		s.encode[members[i]] = i
		s.decode[i] = members[i]
	}
}

// Encode converts a slice of strings to a slice of ints using the attributes
// of the encoder.
func (s *StringEncode) Encode(data []string) ([]int, error) {
	var encoded []int

	if err := checkStringFit(data, s.Attributes); err != nil {
		return encoded, err
	}

	for i := range data {
		encoded = append(encoded, s.encode[data[i]])
	}

	return encoded, nil
}

// Decode converts a slice of uint64 to a slice of strings using the
// attributes of the encoder.
func (s *StringEncode) Decode(encoded []int) ([]string, error) {
	var decoded []string

	if err := checkIntFit(encoded, s.decodeSet()); err != nil {
		return decoded, err
	}

	for i := range encoded {
		decoded = append(decoded, s.decode[encoded[i]])
	}

	return decoded, nil
}

func (s *StringEncode) decodeSet() *set.IntSet {
	keys := make([]int, len(s.decode))

	for k := range s.decode {
		keys = append(keys, k)
	}

	return set.NewIntSet(keys)
}
