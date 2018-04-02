// Exercism Go Track:  a solution for the bob exercise
package bob

import s "strings"

// Hey takes something someone says to Bob and returns his answer
func Hey(remark string) string {

	remark = s.TrimSpace(remark)

	if remark == "" {
		return "Fine. Be that way!"
	}

	if s.ToUpper(remark) == remark && s.ToLower(remark) != remark {
		if s.HasSuffix(remark, "?") {
			return "Calm down, I know what I'm doing!"
		}
		return "Whoa, chill out!"
	} else {
		if s.HasSuffix(remark, "?") {
			return "Sure."
		}
		return "Whatever."
	}
}

//
// The eight contemporary solution I examined showed considerable variation.
// A couple were like mine but a couple used regular expression and a couple
// even implemented some of the string library rouines themselves.
//
// What was most striking was that many evalauted what needed to be evaludated
// first so they could use Go's powerful case statement.  Here is an example:
func _Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	upper := strings.ToUpper(remark)
	lower := strings.ToLower(remark)
	question := strings.HasSuffix(remark, "?")
	yelling := upper != lower && upper == remark

	switch {
	case yelling && question:
		return "Calm down, I know what I'm doing!"
	case yelling:
		return "Whoa, chill out!"
	case question:
		return "Sure."
	case remark == "":
		return "Fine. Be that way!"
	}

	return "Whatever."
}
