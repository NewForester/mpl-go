// Package twofer provides a Go solution to the eponymous Exercism slug.
package twofer

import "fmt"

// ShareWith takes an optional name and returns a string "One for <name>, one for me".  The name defaults to 'you'.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}

//
// So simple the other 8 solution examined looked very similar.
// A couple used concatenation instead of fmt.Sprintf();
// a couple had two return points and one used a 'const' string
// defined outside the function.
//
