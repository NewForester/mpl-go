// Package scrabble provides a Go solution to the Exercism scrabble-score slug.
package scrabble

import "strings"

var scores = map[rune]int{
	'A': 1, 'E': 1, 'I': 1, 'O': 1, 'U': 1,
	'L': 1, 'N': 1, 'R': 1, 'S': 1, 'T': 1,
	'D': 2, 'G': 2,
	'B': 3, 'C': 3, 'M': 3, 'P': 3,
	'F': 4, 'H': 4, 'V': 4, 'W': 4, 'Y': 4,
	'K': 5,
	'J': 8, 'X': 8,
	'Q': 10, 'Z': 10,
}

// Score takes a word or string and returns what it would score in Scrabble (English version)
func Score(word string) int {
	sum := 0

	for _, cc := range strings.ToUpper(word) {
		sum = sum + scores[cc]
	}

	return sum
}

//
// My solution is very similar to the majority of other solutions examined.
// All but one declared a 'dict' and used a simple lookup to calculate the
// score.  The exception used a switch statement.
//
// Several solutions use 'these letters score this" notation following the
// read me.  One of these converted this to the 'this letter scores this'
// within the declaration.  I liked their use of type as in:
type spec = map[string]int
type scrabble = map[rune]int
//
