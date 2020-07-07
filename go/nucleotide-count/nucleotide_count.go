package dna

import "errors"

type Histogram map[rune]int

type DNA string

// Counts
func (d DNA) Counts() (Histogram, error) {
	h := Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}
	for _, val := range d {
		if testNecletides(val) {
			h[val]++
		} else {
			return nil, errors.New("error")
		}

	}
	return h, nil
}

// testNecledides
func testNecletides(x rune) bool {
	result := false
	var nucletides = []rune{'A', 'C', 'G', 'T'}
	for _, val := range nucletides {
		if val == x {
			result = true
			break
		}
	}
	return result
}
