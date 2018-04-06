// Package reverse provides a Go solution to the Exercism reverse-string slug.
package reverse

// String takes a string and returns another with the characters reversed with respect to the first
func String(input string) string {
	result := ""

	for _, cc := range input {
		result = string(cc) + result
	}

	return result
}

//
// Hmm ... my solution is shorter than most, if not all the other 8 solutions
// examined, but is clearly C, not Go and has the inefficiencies of repeated
// concatenation.
//
// Most converted to a rune array and used a decrementing index: there appears
// to be no 'reverse iterator'.  The clever solutions swapped runes: simpler
// solutions simply copied to a new rune array.  One used strings.Builder(),
// which is very Go but also probably overkill here.
//
