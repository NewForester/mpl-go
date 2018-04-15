// Package diffsquares provides a Go solution to the Exercism difference-of-squares slug.
package diffsquares

// SquareOfSums takes a natural number, nn, and returns the square of the sum of the natural numbers [1..nn]
func SquareOfSums(nn int) int {
	sum := nn * (nn + 1) / 2
	return sum * sum
}

// SumOfSquares takes a natural number, nn, and returns the sum of the squares of the natural numbers [1..nn]
func SumOfSquares(nn int) int {
	return nn * (nn + 1) * (2*nn + 1) / 6
}

// Difference returns SquareOSums(nn) - SumOfSquares(nn)
func Difference(nn int) int {
	return SquareOfSums(nn) - SumOfSquares(nn)
}

//
// Only two of the other solution examined used the formula and one of those
// was criticised.
//
// The other used loops.  One even managed two where one would do.
//
