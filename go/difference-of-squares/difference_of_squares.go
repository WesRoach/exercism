package diffsquares

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}

func SquareOfSum(n int) int {
	// sum := 0
	// for i := 1; i < n+1; i++ {
	// 	sum += i
	// }
	sum := (n + 1) * n / 2

	return sum * sum
}

func SumOfSquares(n int) int {
	// sum := 0
	// for i := 1; i < n+1; i++ {
	// 	sum += i * i
	// }
	// return sum

	return n * (n + 1) * (2*n + 1) / 6
}
