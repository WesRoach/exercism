// Package acronym implements functionality for converting
// strings into acronyms
package acronym

import (
	"regexp"
	"strings"
)

// regex: capture words (include ' in word)
var r = regexp.MustCompile("[a-zA-Z']+")

// Abbreviate given string using the first letter of each word
func Abbreviate(s string) string {
	/*
		BenchmarkAcronym-8
		51724
		22407 ns/op
		4211 B/op
		77 allocs/op
	*/

	// find [start, end] of each word
	// given: "do not"
	// returns: [[0 2], [3 6]] (start_of_word, end_of_word + 1)
	indices := r.FindAllStringIndex(s, -1)

	// Create indices-length array of strings, fill
	acronym := make([]string, len(indices))

	// Place first letter of each word in -acronym-
	for i := range indices {
		acronym[i] = string(s[indices[i][0]])
	}

	return strings.ToUpper(strings.Join(acronym[:], ""))
}
