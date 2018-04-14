// Package proverb provides a Go solution to the eponymous Exercism slug.
package proverb

import "fmt"

// Proverb takes a list of words and uses them to generate a proverb about loss
func Proverb(rhyme []string) []string {
	var proverb []string // proverb := make([]string, 0, len(rhyme))

	if len(rhyme) == 0 {
		return proverb
	}

	previous := rhyme[0]

	for _, current := range rhyme[1:] {
		proverb = append(proverb, fmt.Sprintf("For want of a %s the %s was lost.", previous, current))

		previous = current
	}

	return append(proverb, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
}

//
// The were only 4 other solutions to examined.  One used []interface{} in a
// way I do not recognise.
//
// My solution has the same form as the others examined but with some unique
// features.  This version incorporates a couple of changes to show I didn't
// need to do make (as it seems Go will do that as necessary).
//
