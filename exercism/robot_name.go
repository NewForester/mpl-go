// Package robotname provides a Go solution for the Exercism robot-name slug.
package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

// A source of pseudo random numbers
var randy = rand.New(rand.NewSource(time.Now().UnixNano()))

// The Robot structure has only a name but the name is unique
type Robot struct {
	name string
}

// New returns a new Robot with a unique name
func New() Robot {
	return Robot{name: newName()}
}

// Name returns the Robot's unique name
func (robot Robot) Name() string {
	return robot.name
}

// Reset gives the Robot a new unique name
func (robot *Robot) Reset() {
	robot.name = newName()
}

// newName returns a suitably unique name for a Robot
func newName() string {
	rn := randy.Intn(26 * 26 * 1000)

	return fmt.Sprintf("%c%c%03d", makeLetter(rn/1000/26), makeLetter(rn/1000), rn%1000)
}

// makeletter takes a natural number and returns an uppercase ASCII letter as a rune
func makeLetter(nn int) rune {
	return rune('A' + (nn % 26))
}

//
// Variation in other solutions indicates no obvious solution.
// Only one other used fmt.Sprintf and it used %v.
// About half used lazy name evaluation and half checked for collisions.
// One of the latter is ineffective and only one wiped the old name.
//
// Most seem not to understand the random number generator.
//
