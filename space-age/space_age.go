package space

type Planet string

const secondsInEarthYear = 31557600 //units: seconds/earth year

// The map below calculates
// type period_map map[string]float64

func Age(seconds float64, planet Planet) float64 {

	planets_dict := map[Planet]float64{ // units: earth years/planet year
		"Mercury": 0.2408467,
		"Venus":   0.61519726,
		"Earth":   1.0,
		"Mars":    1.8808158,
		"Jupiter": 11.862615,
		"Saturn":  29.447498,
		"Uranus":  84.016846,
		"Neptune": 164.79132,
	}
	return (seconds) / (secondsInEarthYear * planets_dict[planet])
}
