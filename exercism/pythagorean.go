// Package pythagorean provides a Go solution for the Exercism pythagorean-triple slug.
package pythagorean

// Readme returns the product of the Pythagorean triplet whose sum is 1000 AS DESCRIBED IN THE README
func Readme() int {
	/*
	   From Euclid's formula for Pythagorean triplet was have:

	       aa = 2 * mm * nn

	       bb = mm * mm - nn * nn

	       cc = mm * mm + nn * nn

	   where the constraint is:

	       aa + bb + cc = 1000

	   substituting for aa, bb and cc we have:

	       1000 = (2 * mm * nn) + (mm * mm - nn * nn) + (mm * mm + nn * nn)

	   that can be reduced to:

	       nn = ((1000 / 2) / mm) - mm

	   the rest is brute force (and very quick)
	*/
	for mm := 1; ; mm++ {
		nn := (500 / mm) - mm

		if mm < nn || nn < 0 {
			continue
		}

		aa, bb, cc := euclid(nn, mm)

		if aa+bb+cc == 1000 {
			return aa * bb * cc
		}
	}

	return -1 // does not happen
}

// euclid takes two numbers (in principle coprime) and returns a Pythagorean triplet
func euclid(nn, mm int) (aa, bb, cc int) {
	if mm < nn {
		nn, mm = mm, nn
	}

	aa = 2 * mm * nn

	bb = mm*mm - nn*nn

	if bb < aa {
		aa, bb = bb, aa
	}

	cc = mm*mm + nn*nn

	return
}

type Triplet [3]int

// Range is one of the functions the tests tell us to implement
func Range(min, max int) (triplets []Triplet) {
	for nn := 1; nn < max; nn++ {
		for mm := nn + 1; mm < max; mm++ {
			aa, bb, cc := euclid(nn, mm)

			if aa < min || cc > max {
				continue
			}

			triplets = append(triplets, Triplet{aa, bb, cc})
		}
	}

	return
}

// Sum is the other function the tests tell us to implement
func Sum(tp int) (triplets []Triplet) {
	ff := 1

	for _, dv := range []int{1, 3, 5} {
		tp = tp / dv
		ff = ff * dv

		for mm := 1; ; mm++ {
			nn := (tp / 2 / mm) - mm

			if nn < 0 {
				break
			}

			aa, bb, cc := euclid(nn, mm)

			if aa+bb+cc == tp {
				triplets = append(triplets, Triplet{aa * ff, bb * ff, cc * ff})
			}
		}
	}

	return
}

//
// What a balls up.
//
// I write a routine that finds the answer rather than simply returns it.  Then
// I find that the track wants me to write two other routines that do not find
// the the answer but are there presumably to stop you simply returning it.
//
// I implemented the second first.  It seems, from the solutions submitted by
// other folks, that it is expected you would implement the second in terms of
// the first.
//
// My commonality is in the euclid routine which I created last. My version does
// not handle the k factor (not required to solve the read me problem but might
// help implement a better Sum() function) but does handle negative inputs etc.
//
// All solutions examined look verbose and that goes for mine too.  That is Go
// for you.  All but two used brute force without the benefit of Euclid's
// Theorem.  One of those used the k factor.
//
// Solutions that use the k factor may be a lot faster than those that do not
// but it does seem to complicate matters - use of gcd - many times, (integer)
// sqrt - loop limits and requires duplicate removal.  Hmm ... someday, maybe.
//
