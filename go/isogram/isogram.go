package isogram

import (
	"unicode"
)

// IsIsogram function
func IsIsogram(input string) bool {
	for key, val := range input {
		if (val >= 'A' && val <= 'Z') || (val >= 'a' && val <= 'z') {
			for i := key + 1; i < len(input); i++ {
				if unicode.ToLower(val) == unicode.ToLower(rune(input[i])) {
					return false
				}
			}
		}
	}
	return true
}
