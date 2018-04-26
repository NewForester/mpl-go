// Package prime provides a Go solution for the Exercism prime factors slug.
package prime

// Factors takes a natural number and returns its prime factors in ascending order
func Factors(number int64) (factors []int64) {
	factors = []int64{}

	for ; number&1 == 0; number >>= 1 {
		factors = append(factors, 2)
	}

	for candy := int64(3); number > 1; candy += 2 {
		for ; number%candy == 0; number /= candy {
			factors = append(factors, candy)
		}
	}

	return
}

//
// This second solution does not incorporate any code from the nth-prime
// exercise and does not involve the generation of a list of primes.
//
// It does as almost all the other solutions did and uses a brute force
// method of calculating factors.
//
// It is a lot (orders of magitude) faster than my first attempt and this
// appears to be because it cuts out the generation of the list of primes
// and the memory managment / garbage collection that involves.  Beware.
//
// $ go test --bench .
// goos: linux
// goarch: amd64
// BenchmarkPrimeFactors-2              100          16249759 ns/op
// PASS
// ok      _/home/work/exercism/go/prime-factors   1.663s
//
// which I think means the test suite takes 16 ms to run on my machine.
//
