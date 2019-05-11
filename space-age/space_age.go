package space

const EarthYearInSeconds = 31557600

type Planet string

var planetMap = map[Planet]float64{
	"Mercury": EarthYearInSeconds * 0.2408467,
	"Venus":   EarthYearInSeconds * 0.61519726,
	"Mars":    EarthYearInSeconds * 1.8808158,
	"Jupiter": EarthYearInSeconds * 11.862615,
	"Saturn":  EarthYearInSeconds * 29.447498,
	"Uranus":  EarthYearInSeconds * 84.016846,
	"Neptune": EarthYearInSeconds * 164.79132,
	"Earth":   EarthYearInSeconds,
}


func Age(seconds float64, planetName Planet) float64 {
	return seconds / planetMap[planetName]
}
