package grains

import (
	"errors"
)

func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, errors.New("invalid number")
	}
	// return uint64(math.Pow(2, float64(n-1))), nil
	return 1 << (n - 1), nil
}

func Total() uint64 {
	return 1<<64 - 1
}
