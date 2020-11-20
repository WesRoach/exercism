// Package hamming implements a simple function for calculating
// Hamming distance between two strings.
package hamming

import "errors"

// Distance calculates the Hamming distance between two strings
// Sequences of mis-matched length will return an error
func Distance(a, b string) (int, error) {
	// Fail case
	if len(a) != len(b) {
		return 0, errors.New("a, b unequal lengths")
	}

	hammingDistance := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			hammingDistance++
		}
	}
	return hammingDistance, nil

}
