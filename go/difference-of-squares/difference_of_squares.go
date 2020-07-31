package diffsquares

// SquareOfSum func
func SquareOfSum(n int) int {
	return ((1+n)*n)/2 * ((1+n)*n)/2
}

// SumOfSquares func
func SumOfSquares(n int) int {
	return n*(n+1)*(2*n+1)/6
}

// Difference fun
func Difference(n int ) int {
	return SquareOfSum(n) - SumOfSquares(n)
}