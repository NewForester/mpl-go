// Exercism Go Track:  a solution for the gigasecond exercise

package gigasecond

// import path for the time package from the standard library
import "time"

// AddGigasecond takes a time.Time value and returns it after adding 1e9 seconds to it.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(1e9 * time.Second)
}
