// Package triangle implements triangle functionality
package triangle

import "math"

// Kind is a category of triangle
type Kind int

// Define triangle category constants
const (
	NaT Kind = iota // 0; not a triangle
	Equ             // 1; equilateral
	Iso             // 2; isosceles
	Sca             // 3; scalene
)

// KindFromSides returns the type of triangle, Kind
func KindFromSides(a, b, c float64) Kind {
	/*
		BenchmarkKind-8
		1892742
		630 ns/op
		0 B/op
		0 allocs/op
	*/
	for _, v := range []float64{a, b, c} {
		if v <= 0 || math.IsNaN(v) || math.IsInf(v, 0) {
			return NaT
		}
	}

	if (a+b < c) || (b+c < a) || (a+c < b) {
		return NaT
	} else if a == b && b == c {
		return Equ
	} else if a == b || b == c || c == a {
		return Iso
	} else {
		return Sca
	}
}
