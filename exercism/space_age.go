// Package space provides a Go solution to the Exercism space-age slug.
package space

type Planet string

// Age takes an age expressed in seconds and a planet name and returns the age in planetary years
func Age(age float64, planet Planet) float64 {

	year := 0.0

	switch planet {
	case "Earth":
		year = 1.0
	case "Mercury":
		year = 0.2408467
	case "Venus":
		year = 0.61519726
	case "Mars":
		year = 1.8808158
	case "Jupiter":
		year = 11.862615
	case "Saturn":
		year = 29.447498
	case "Uranus":
		year = 84.016846
	case "Neptune":
		year = 164.79132
	}

	return age / 31557600.0 / year
}

//
// I have used a a switch here.  Only 1 of the 8 othere examined did likewise.
// The others all ussd a map to set up an associative array.
//
