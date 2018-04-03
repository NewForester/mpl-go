// Package raindrops provides a Go solution to the eponymouse Exercism slug.
package raindrops

import "fmt"

type Rule struct {
	divisor int
	sound   string
}

// Convert takes a natural number and returns the corresponding 'sound' as per the read me
func Convert(nn int) string {
	rules := [3]Rule{{3, "Pling"}, {5, "Plang"}, {7, "Plong"}}

	result := ""

	for _, rule := range rules {
		if nn%rule.divisor == 0 {
			result += rule.sound
		}
	}

	if result == "" {
		result = fmt.Sprintf("%d", nn)
	}

	return result
}

//
// Of the 8, 7 unrolled the loop and 3 had a redundant loop that counted to nn.
// One had a map for no good reason and 6 had a 'faster' conversion from integer
// to string:
//     ret = strconv.Itoa(num)
// but strconv has to be imported.
//
