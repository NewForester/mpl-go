// Exercism Go Track:  a solution for the leap (year) exercise

package leap

// IsLeapYear takes an integer representing a year and returns true if that year is a leap year
func IsLeapYear(year int) bool {
	if year%400 == 0 {
		return true
	} else if year%100 == 0 {
		return false
	} else if year%4 == 0 {
		return true
	} else {
		return false
	}
}

//
// This is a simple exercise that is crying out for pattern matching.  Go has none.
//
// Most solution used a single expression that is less easy to reason about.
// One used nested if.
//
// The go style guide rightly points out that the `else` keywords are redundant
// and that leads to more indentation than necessary.  Not so here.
//
// Remove them here and the code will not compile - you need to replace `else`
// with a new line, which makes the code longer.
//
// I miss pattern matching with guards.
//
