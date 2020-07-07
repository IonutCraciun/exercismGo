package strand

/*
G -> C
C -> G
T -> A
A -> U
*/

// ToRNA function
func ToRNA(dna string) string {
	var result string
	for _, val := range dna {
		switch val {
		case 'G':
			result = result + "C"
		case 'C':
			result = result + "G"
		case 'T':
			result = result + "A"
		case 'A':
			result = result + "U"
		}
	}
	return result
}
