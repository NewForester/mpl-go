// Package prime provides a Go solution for the Exercism prime factors slug.
package prime

import "math"

// Factors takes a natural number and returns its prime factors in ascending order
func Factors(number int64) (factors []int64) {
	factors = []int64{}

	primes := []int64{2, 3, 5, 7, 11}
	prime := int64(1)

	for number > 1 {
		prime, primes = nextPrime(prime, primes)

		for number%prime == 0 {
			factors = append(factors, prime)

			number = number / prime
		}
	}

	return
}

/// nextPrime() returns the next prime, extending the list of primes as necessary
func nextPrime(prime int64, primes []int64) (int64, []int64) {
	var candidate int64 = 1

	for _, candidate = range primes {
		if candidate > prime {
			return candidate, primes
		}
	}

	for isprime := false; !isprime; {
		candidate = candidate + 1

		limit := int64(math.Sqrt(float64(candidate)))

		for _, prime := range primes {
			if prime > limit {
				isprime = true
				break
			}
			if candidate%prime == 0 {
				break
			}
		}
	}

	primes = append(primes, candidate)

	return candidate, primes
}

//
// My solution incorporates code from the nth-prime exercise and involves the
// construction of a list of primes.
//
// Only one other solution did this but the list of primes was generated in
// parallel.  This solution was a order of magnitude faster than mine but I
// am not sure that this has anything to do with the parallelism.
//
// All the other solutions did not generate a list of primes.  They are much
// simpler and I expect much faster.
//
