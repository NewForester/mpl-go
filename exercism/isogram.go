// Package isogram provides a Go solution to the eponymous Exercism slug.
package isogram

import "strings"
import "unicode"

// IsIsogram takes a string of letters and returns true if they represent an isogram
func IsIsogram(word string) bool {
	lettersSeen := ""

	for _, cc := range word {
		if unicode.IsLetter(cc) {
			cc = unicode.ToLower(cc)

			if strings.ContainsRune(lettersSeen, cc) {
				return false
			}

			lettersSeen += string(cc)
		}
	}

	return true
}

//
// All solutions examined followed the same pattern:  all terminated early.
//
// I used a plain string, about half used a map instead and one a bit field.
// My string is a plain string ... others used make(...).  I wonder what the
// practical different between the two is.
//
// Does Go not have a set ?  On other tracks, set was popular and map was not.
//
