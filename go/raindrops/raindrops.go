// Package raindrops implements functionality to convert
// an int into rain drop sounds
package raindrops

import (
	"strconv"
)

// Convert an int to the sound of rain
func Convert(n int) string {
	/*
		BenchmarkConvert-8
		2086962
		583 ns/op
		64 B/op
		4 allocs/op
	*/
	words := [4]string{"", "Pling", "Plang", "Plong"}

	factor3, factor5, factor7 := 0, 0, 0
	if n%3 == 0 {
		factor3 = 1
	}
	if n%5 == 0 {
		factor5 = 2
	}
	if n%7 == 0 {
		factor7 = 3
	}

	if !(factor3 > 0 || factor5 > 0 || factor7 > 0) {
		return strconv.Itoa(n)
	}

	return words[factor3] + words[factor5] + words[factor7]
}

// Convert does things
// func Convert(n int) string {
// 	/*
// 		BenchmarkConvert-8
// 		722572
// 		1643 ns/op
// 		496 B/op
// 		24 alocs/op
// 	*/
// 	var str []string
// 	if n%3 == 0 {
// 		str = append(str, "Pling")
// 	}
// 	if n%5 == 0 {
// 		str = append(str, "Plang")
// 	}
// 	if n%7 == 0 {
// 		str = append(str, "Plong")
// 	}
// 	if str == nil {
// 		return strconv.Itoa(n)
// 	}

// 	return strings.Join(str, "")
// }
