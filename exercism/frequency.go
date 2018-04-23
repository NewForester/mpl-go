// Package letter provides a Go solution for the Exercism parallel letter frequency slug.
package letter

type FreqMap map[rune]int

// Frequency takes a string and returns a rune map containing letter frequencies - without concurrency.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ChannelFrequency wraps Frequency to returm the rune map via a channel
func ChannelFrequency(s string, c chan FreqMap) {
	c <- Frequency(s)
}

// ConcurrentFrequency takes an array of strings and returns one rune map containing letter frequencies - with concurrency.
func ConcurrentFrequency(ss []string) FreqMap {
	m := FreqMap{}
	channel := make(chan FreqMap)
	for _, s := range ss {
		go ChannelFrequency(s, channel)
	}
	for range ss {
		ms := <- channel
		for k, v := range ms {
			m[k] += v
		}
	}
	return m
}

//
// My first piece of Go code to use concurrent execution.  This may be a toy
// that does not actually execute the test cases any faster than a sequential
// solution does.
//
// My solution follows the same form as all the others examined, which give me
// some confidence that it is not horrendously bad.  Most used a lambda function
// where I added a new function.  So:
//
go func(s string) {channel <- Frequency(s)} (s)
//
// Overall, I found adding the concurrency simple to do and it (literally)
// only turned a simple 14 line routines into a 35 line program.
//
