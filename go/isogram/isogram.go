// Package isogram implements functionality to determine Isograminity
package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram returns true/false if supplied string is an Isogram
func IsIsogram(word string) bool {
	/*
			BenchmarkIsIsogram-8
			129034
			8424 ns/op
			1469 B/op
		 	16 allocs/op
	*/
	wordTracker := make(map[rune]bool)
	if len(word) == 0 {
		return true
	}
	for _, letter := range strings.ToLower(word) {
		if unicode.IsLetter(letter) {
			if _, ok := wordTracker[letter]; ok {
				return false
			}
			wordTracker[letter] = true
		}
	}
	return true
}
