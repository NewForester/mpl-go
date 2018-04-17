// Package luhn provides a Go solution for the eponymous Exercism slug.
package luhn

import "unicode"

// Valid takes a SIN and returns true if it is valid
func Valid(number string) bool {
	luhnsum := 0
	digitcount := 0

	for _, cc := range number {
		if unicode.IsDigit(cc) {
			digitcount += 1

			luhndigit := int(cc) - '0'
			if digitcount%2 == 0 {
				luhndigit *= 2
				if luhndigit > 9 {
					luhndigit -= 9
				}
			}

			luhnsum += luhndigit
		} else if cc != ' ' {
			return false
		}
	}

	return digitcount > 1 && luhnsum%10 == 0
}

//
// Only two of the solutions examined were similar to mine.  The others
// were all long winded.
//
// I looked but the conversion from rune digit to integer digit is either as
// I have done or by using strconv::Atoi or strconv:ParseInt.
//
