// Package acronym provides a Go solution to the eponymouse Exercism slug.
package acronym

import s "strings"
import u "unicode"

// Abbreviate takes a technical term and returns an acronym for that term
func Abbreviate(term string) string {

	result := ""
	capitaliseNext := true
	previousLower := false

	for _, cc := range term {
		if capitaliseNext || previousLower && u.IsUpper(cc) {
			result += string(cc)
		}

		capitaliseNext = cc == ' ' || cc == '-'
		previousLower = u.IsLower(cc)
	}

	return s.ToUpper(result)
}

//
// The other 8 recent solutions all started by splitting the term into to words
// but all appeared to have no code to handle 'HyperText Markup Language'.
//
// So I think none of them work and all fail the 'can I reason about this' test
// but it may be that I am missing something.
//
