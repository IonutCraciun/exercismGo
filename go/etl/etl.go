package etl

import (
	"strings"
)

// Transform is
func Transform(input map[int][]string) map[string]int {
	var result = make(map[string]int)
	for key, val := range input {
		for _, item := range val {
			result[strings.ToLower(item)] = key
		}
	}
	return result
}
