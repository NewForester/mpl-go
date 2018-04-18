// Package grains provides a Go solution for the eponymous Exercism slug.
package grains

import "errors"

// Square returns the number of grains on square N
func Square(nn int) (grains uint64, err error) {
	if nn >= 1 && nn <= 64 {
		grains = 1 << uint(nn-1)
	} else {
		err = errors.New("domain error")
	}

	return
}

// Total returns the total number of grains on the chessboard
func Total() uint64 {
	return 0xffffffffffffffff
}

//
// A very simple exercise made a tad more interesting by the tests insisting
// on the error return.
//
// Most of those examined used math.pow in their Square and a loop in their
// Total.  Few used Go's named return variables.
//
