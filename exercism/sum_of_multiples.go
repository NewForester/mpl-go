// Package summultiples provides a Go solution to the Exercism sum-of-multiples slug.
package summultiples

// SumMultiples  does what it must by brute force
func SumMultiples(limit int, factors ...int) int {
	sum := 0

	for _, factor := range factors {
		for multiple := factor; multiple < limit; multiple += factor {
			sum += multiple
			for _, bigger := range factors {
				if bigger > factor && multiple%bigger == 0 {
					sum -= multiple
					break
				}
			}
		}
	}

	return sum
}

//
// Half the solutions examined used a loop on limit and then a lopp on factors.
//
// The other half used a sieve and so memory was O(n) as well.
//
// My solution uses constant memory but is O(N**2) where N is the number of factors
//
