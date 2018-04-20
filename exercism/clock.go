// Package clock provides a Go solution for the eponymous Exercism slug.
package clock

import "fmt"

type Clock struct {
	hh int
	mm int
}

// New returns a new, initialised, Clock
func New(hh, mm int) (clock Clock) {
	time := hh*60 + mm

	hh = time / 60 % 24
	mm = time % 60

	if mm < 0 {
		hh -= 1
		mm += 60
	}
	if hh < 0 {
		hh += 24
	}

	clock = Clock{hh, mm}

	return
}

// String formats the Clock time for printing and returns the formatting string
func (clock *Clock) String() string {
	return fmt.Sprintf("%02d:%02d", clock.hh, clock.mm)
}

// Add adds the given number of minutes to the Clock time and returns the result as a new Clock
func (clock Clock) Add(mm int) (result Clock) {
	result = New(clock.hh, clock.mm+mm)

	return
}

// Subtract subtracts the given number of minutes from the Clock time and returns the result as a new Clock
func (clock Clock) Subtract(mm int) (result Clock) {
	result = New(clock.hh, clock.mm-mm)

	return
}

//
// This solution is short and neat.
//
// Many of the solution examined were much longer and those that weren't used
// a single integer to store the clock.
//
