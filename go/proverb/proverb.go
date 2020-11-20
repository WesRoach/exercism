// Package proverb implements functionality for creating proverbs
// from lists of strings
package proverb

import "fmt"

const (
	multiLine = "For want of a %v the %v was lost."
	finalLine = "And all for the want of a %v."
)

// Proverb outputs a proverb given a list of strings
func Proverb(rhyme []string) []string {
	/*
		BenchmarkProverb-8
		210519
		4989 ns/op
		1472 B/op
	 	51 allocs/op
	*/
	lengthRhyme := len(rhyme)
	proverb := make([]string, lengthRhyme)
	for idx := range rhyme {
		if idx == lengthRhyme-1 {
			proverb[idx] = fmt.Sprintf(finalLine, rhyme[0])
			return proverb
		}

		proverb[idx] = fmt.Sprintf(multiLine, rhyme[idx], rhyme[idx+1])
	}
	return proverb
}
