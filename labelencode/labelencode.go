// Package labelencode implements a Label encoding scheme for strings and
// integers.
package labelencode

import (
	"errors"

	"github.com/averagesecurityguy/structures/set"
)

var NotFit = errors.New("labelencoder: data does not fit given attributes")

// NewIntEncode returs an IntEncode structure.
func NewIntEncode() *IntEncode {
	ie := new(IntEncode)

	ie.encode = make(map[int]int)
	ie.decode = make(map[int]int)

	return ie
}

// NewStringEncode returns a StringEncode structure.
func NewStringEncode() *StringEncode {
	se := new(StringEncode)

	se.encode = make(map[string]int)
	se.decode = make(map[int]string)

	return se
}

// checkStringFit ensures the attributes of the slice of strings are contained
// in the given slice of attributes.
func checkStringFit(data []string, expected *set.StringSet) error {
	// Uniqe and sort the data to get its attributes.
	given := set.NewStringSet(data)

	if !given.Subset(expected) {
		return NotFit
	}

	return nil
}

// checkIntFit ensures the attributes of the slice of ints are contained in
// the given slice of attributes.
func checkIntFit(data []int, expected *set.IntSet) error {
	// Uniqe and sort the data to get its attributes.
	given := set.NewIntSet(data)

	if !given.Subset(expected) {
		return NotFit
	}

	return nil
}
