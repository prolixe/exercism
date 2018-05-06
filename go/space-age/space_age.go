package space

type Planet string

const (
	Mercury Planet = "Mercury"
	Venus   Planet = "Venus"
	Earth   Planet = "Earth"
	Mars    Planet = "Mars"
	Jupiter Planet = "Jupiter"
	Saturn  Planet = "Saturn"
	Uranus  Planet = "Uranus"
	Neptune Planet = "Neptune"

	EarthYearInSeconds = 31557600.0
	EarthYear          = 1.0

	/*
		Earth: orbital period 365.25 Earth days, or 31557600 seconds
		Mercury: orbital period 0.2408467 Earth years
		Venus: orbital period 0.61519726 Earth years
		Mars: orbital period 1.8808158 Earth years
		Jupiter: orbital period 11.862615 Earth years
		Saturn: orbital period 29.447498 Earth years
		Uranus: orbital period 84.016846 Earth years
		Neptune: orbital period 164.79132 Earth years
	*/

	MercuryYear = EarthYear * 0.2408467
	VenusYear   = EarthYear * 0.61519726
	MarsYear    = EarthYear * 1.8808158
	JupiterYear = EarthYear * 11.862615
	SaturnYear  = EarthYear * 29.447498
	UranusYear  = EarthYear * 84.016846
	NeptuneYear = EarthYear * 164.79132
)

var orbitRatios = map[Planet]float64{
	Mercury: MercuryYear,
	Venus:   VenusYear,
	Earth:   EarthYear,
	Mars:    MarsYear,
	Jupiter: JupiterYear,
	Saturn:  SaturnYear,
	Uranus:  UranusYear,
	Neptune: NeptuneYear,
}

func Age(seconds float64, p Planet) float64 {
	return seconds / (EarthYearInSeconds * orbitRatios[p])
}
