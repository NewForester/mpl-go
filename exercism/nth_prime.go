// Package prime provides a Go solution to the Exercism nth-prime slug.
package prime

import "math"

// Nth return the nth prime
func Nth(nn int) (int, bool) {
	if nn < 1 {
		return nn, false
	}

	primes := make([]int, 0, nn)
	candidate := 2
	primes = append(primes, candidate)

	for len(primes) < nn {
		candidate = candidate + 1

		limit := int(math.Sqrt(float64(candidate)))

		isprime := true
		for _, prime := range primes {
			if prime > limit {
				break
			}
			if candidate%prime == 0 {
				isprime = false
				break
			}
		}

		if isprime {
			primes = append(primes, candidate)
		}
	}

	return candidate, true
}

//
// Of the 8 other submissions examined, 2 made no attempt to solve the problem
// and half of those that did use an Sieve of Eratosthenes.
//
// The problem with the sieve is knowing how big it must be in order find the
// nth prime.  For this you can use the prime number theorem.
//
// My solution has no unique features and there is nothing missing other than
// the sieve.  Adding 'the upper bound is the square root' (7 lines) made my
// solution about average length and average complexity.  Some folk used two
// or even three routines and all that did was add useful routine names.
//
