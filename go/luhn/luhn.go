// Package luhn implements SIN validation
package luhn

import "unicode"

// Valid checks if provided string is a valid SIN
func Valid(s string) bool {
	if len(s) <= 1 {
		return false
	}

	// Create int array from strings
	// Only fill intArr when value is valid digit
	// Track maximum used index of intArr
	intArr := make([]int, len(s))
	intArrMaxIdx := 0
	for _, val := range s {
		if unicode.IsDigit(val) {
			intArr[intArrMaxIdx] = int(val - '0')
			intArrMaxIdx++
		} else if !unicode.IsSpace(val) {
			return false
		}
	}

	// Catches " 0" variants
	if intArrMaxIdx <= 1 {
		return false
	}

	// Starting from the right:
	// 1) Double each int
	// 2) If int > 9, subtract 9
	for i := intArrMaxIdx - 2; i >= 0; i = i - 2 {
		intArr[i] = intArr[i] * 2
		if intArr[i] > 9 {
			intArr[i] = intArr[i] - 9
		}
	}

	// Sum converted values in intArr
	sum := 0
	for i := 0; i < intArrMaxIdx; i++ {
		sum += intArr[i]
	}

	// If sum is evenly divisble by 10: true
	if sum%10 == 0 {
		return true
	}
	return false
}
