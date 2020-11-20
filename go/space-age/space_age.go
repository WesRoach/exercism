package space

type Planet string

// IT'S AN EARTH YEAR REPRESENTED IN SECONDS
const EARTH_YEAR_IN_SECONDS = 31557600

// Map of seconds per year for each planet
var secondsPerYear = map[Planet]float64{
	"Earth":   EARTH_YEAR_IN_SECONDS,
	"Mercury": EARTH_YEAR_IN_SECONDS * 0.2408467,
	"Venus":   EARTH_YEAR_IN_SECONDS * 0.61519726,
	"Mars":    EARTH_YEAR_IN_SECONDS * 1.8808158,
	"Jupiter": EARTH_YEAR_IN_SECONDS * 11.862615,
	"Saturn":  EARTH_YEAR_IN_SECONDS * 29.447498,
	"Uranus":  EARTH_YEAR_IN_SECONDS * 84.016846,
	"Neptune": EARTH_YEAR_IN_SECONDS * 164.79132,
}

// Given Age in Seconds, & Planet
// Return Solar Age based on Planet
func Age(seconds float64, planet Planet) float64 {
	return seconds / secondsPerYear[planet]
}
