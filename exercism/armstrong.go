// Package armstrong provides a Go solution for the eponymous Exercism slug.
package armstrong

// IsNumber takes a natural number and returns true if the number is an Armstrong number
func IsNumber(nn int) bool {
	result, _ := calculate(nn, 1)

	return nn == result
}

// calculate returns the Armstrong sum using recursion but lets the boss make the decision
func calculate(nn int, count int) (partial int, total int) {
	digit := nn % 10

	if nn == digit {
		partial, total = 0, count
	} else {
		partial, total = calculate(nn/10, count+1)
	}

	factor := digit
	for ii := total - 1; ii > 0; ii-- {
		factor *= digit
	}

	partial += factor

	return
}

//
// The first question in this exercise is how many digits are there ?
//
// Half of the solutions examined went for strconv to answer this.  This leads
// to a subscript and then a conversion in the other direction.
//
// The second question is whether to use math.pow() and floating point.  The
// strconv solutions tended to, while the other half were more inventive.
//
// The other half used two loops.  I avoid this by using recursion.
//
// The most interesting solution (had been cloned) disguised the two loops
// by hiding one in a subroutine.  It also calculated the power using a recursive
// subroutine.
//
