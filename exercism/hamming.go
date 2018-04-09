// Package hamming provides a Go solution to the eponymous Exercism slug.
package hamming

import "errors"

// Distance returns the hamming distance between to strands of DNA
func Distance(a, b string) (int, error) {
	count := 0

	if len(a) != len(b) {
		return 0, errors.New("Strands of different lengths")
	}

	for ii := len(b) - 1; ii >= 0; ii = ii - 1 {
		if a[ii] != b[ii] {
			count += 1
		}
	}

	return count, nil
}

//
// A simple problem and all solutions take the same approach.
//
// Those more familiar with Go used for ii := range a, which is better.
//
// Two converted only b to runes and I cannot see why only b and not a would
// need converting.  Furthermore, I would guess any conversion would need to be
// done before comparing the lengths.  (Rust forced conversion but to bytes).
//
