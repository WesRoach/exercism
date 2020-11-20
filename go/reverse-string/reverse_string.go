// Package reverse implements functionality for reversing strings
package reverse

// Reverse reverses a given string
func Reverse(word string) string {
	/*
		BenchmarkReverse-8
		1623808
		648 ns/op
		64 B/op
		5 allocs/op
	*/
	if len(word) == 0 {
		return word
	}

	reversed := []rune(word)
	for i, j := 0, len(reversed)-1; i < j; i, j = i+1, j-1 {
		reversed[i], reversed[j] = reversed[j], reversed[i]
	}
	return string(reversed)
}
