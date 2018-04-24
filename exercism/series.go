// Package series provides a Go solution for the eponymous Exercism slug.
package series

// All takes a length and a string and returns all substrings of the given length
func All(nn int, ss string) []string {
	limit := len(ss) - nn + 1
	rr := make([]string, 0, limit)
	for ii := 0; ii < limit; ii++ {
		rr = append(rr, ss[ii:ii+nn])
	}
	return rr
}

// UnsafeFirst takes a length and a string and returns the first substring of the given length
func UnsafeFirst(nn int, ss string) string {
	return ss[:nn]
}

// First is as for Unsafe but return a boolean:  if the length is great than the length of the string, false is returned
func First(nn int, ss string) (first string, ok bool) {
	ok = len(ss) >= nn

	if ok {
		first = ss[:nn]
	}

	return
}

//
// All solutions examined followed the same pattern.  Some, like mine,
// estimated the size of the array of strings, others did not.  The only
// slightly unusual one did not used append but did a straight assignment.
// Not the clearest of code that one.
//
