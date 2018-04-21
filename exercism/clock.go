// Package clock provides a Go solution for the eponymous Exercism slug.
package clock

import "fmt"

type Clock int

// New returns a new, initialised, Clock
func New(hh, mm int) Clock {
	time := hh*60 + mm%(24*60)

	if time < 0 {
		time += 24 * 60
	}

	return Clock{time}
}

// String formats the Clock time for printing and returns the formatting string
func (clock *Clock) String() string {
	return fmt.Sprintf("%02d:%02d", clock.hh, clock.mm)
}

// Add adds the given number of minutes to the Clock time and returns the result as a new Clock
func (clock Clock) Add(mm int) Clock {
	return New(clock.hh, clock.mm+mm)

}

// Subtract subtracts the given number of minutes from the Clock time and returns the result as a new Clock
func (clock Clock) Subtract(mm int) Clock {
	return New(clock.hh, clock.mm-mm)
}
