// Package collatzconjecture provides a Go solution for the eponymous Exercism slug.
package collatzconjecture

import "errors"

// CollatzConjecture takes a natural number and returns the number of steps required to reach 1
func CollatzConjecture(nn int) (steps int, err error) {
	for {
		switch {
		case nn <= 0:
			err = errors.New("domain error: natural numbers only")
			return
		case nn == 1:
			return
		case nn%2 == 1:
			nn, steps = 3*nn+1, steps+1
		case true:
			nn, steps = nn/2, steps+1
		}
	}
}

//
// This must be new to this track as there were only 8 other solutions.
//
// None used switch (poor man's match).  Half used for n != 1 and ifs.
// 2 others used recursion and the other 2 went berserk over optimisation.
//
