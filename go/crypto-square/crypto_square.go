package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

// Encode func
func Encode(input string) (output string) {
	output = normalize(input)
	numCols := int(math.Ceil(math.Sqrt(float64(len(output)))))

	result := make([]string, numCols)
	for i, r := range output {
		result[i%numCols] += string(r)
	}
	output = strings.ToLower(strings.Join(result, " "))
	return
}

func normalize(input string) (output string) {
	var runeSlice []rune
	for _, val := range input {
		if unicode.IsLetter(rune(val)) || unicode.IsDigit(rune(val)) {
			runeSlice = append(runeSlice, val)
		}
	}
	output = string(runeSlice)
	return
}