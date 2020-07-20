package pangram

import (
	"unicode"
)

// IsPangram func
func IsPangram(s string) bool {
	m := make(map[rune]int)
	if len(s) == 0 {
		return false
	}
	for _, val := range s{
		if !unicode.IsLetter(val) {
			continue
		}
		m[unicode.ToLower(val)]++
	}
	if len(m) == 26 {
		return true
	}
	return false
	
}