package darts

import "math"

// Score func
func Score(x, y float64) int {
	length := math.Sqrt(x*x + y*y)
	if length <= 1 {
		return 10
	} else if length <= 5 {
		return 5
	} else if length <= 10 {
		return 1
	}
	return 0
}
