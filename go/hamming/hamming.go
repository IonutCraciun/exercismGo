package hamming

import "errors"

// Distance calculates Hamming distance between  two DNA strands a and b
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("lengths of string are not equal")
	}
	differences := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			differences++
		}
	}
	return differences, nil

}
