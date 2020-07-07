package space

type Planet string

func Age(seconds float64, planet Planet) float64 {
    yearEarth := seconds / 31557600
    switch planet {
    case "Mercury":
        {
            return yearEarth / 0.2408467
        }
    case "Earth":
        {
            return yearEarth
        }
    case "Venus":
        {
            return yearEarth / 0.61519726
        }
    case "Mars":
        {
            return yearEarth / 1.8808158
        }
    case "Jupiter":
        {
            return yearEarth / 11.862615
        }
    case "Saturn":
        {
            return yearEarth / 29.447498
        }
    case "Uranus":
        {
            return yearEarth / 84.016846
        }
    case "Neptune":
        {
            return yearEarth / 164.79132
        }
    default:
        {
            return 0.0
        }
    }
}

