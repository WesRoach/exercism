// Package strand implements functionality for DNA/RNA conversion
package strand

var dnaToRNA = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

// ToRNA converts DNA sequence into RNA sequence
func ToRNA(dna string) string {
	/*
		BenchmarkRNATranscription-8
		1857578
		624 ns/op
		100 B/op
		11 allocs/op
	*/
	rnaStrand := make([]rune, len(dna))
	for idx, nucleo := range dna {
		rnaStrand[idx] = dnaToRNA[nucleo]
	}
	return string(rnaStrand)
}
