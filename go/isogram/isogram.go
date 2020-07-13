package isogram

import (
	"unicode"
)

// IsIsogram function
func IsIsogram(input string) bool {
	set := make(map[rune]bool)
	for _, val := range input {
		if !unicode.IsLetter(val) {
			continue
		}
		val = unicode.ToLower(val)
		if set[val] {
			return false
		}
		set[val] = true
	}
	return true
}
