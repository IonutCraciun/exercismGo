package pythagorean

// Triplet s
type Triplet [3]int

var triplet = Triplet{3,4,5}

// Range s
func Range(min, max int) []Triplet {
	result := []Triplet{}
	for a := min; a <= max; a++ {
		for b := a; b <= max; b++ {
			for c := b; c <= max; c++ {
				switch {
				case a*a+b*b == c*c:
					result = append(result, Triplet{a, b, c})
					break
				case a*a+b*b < c*c:
					break
				}
			}
		}
	}

	return result
}

// Sum s
func Sum(sum int) []Triplet {
	result := []Triplet{}

	for _, val := range Range(1, sum/2) {
		if val[0]+val[1]+val[2] == sum {
			result = append(result, val)
		}
	}

	return result
}