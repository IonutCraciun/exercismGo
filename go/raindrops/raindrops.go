package raindrops

import "strconv"

// Convert returns a number converted in raindrops sounds
func Convert(x int) string {
	var results string
	if x%3 == 0 {
		results += "Pling"
	}
	if x%5 == 0 {
		results += "Plang"
	}
	if x%7 == 0 {
		results += "Plong"
	}
	if results == "" {
		return strconv.Itoa(x)
	}
	return results
}
