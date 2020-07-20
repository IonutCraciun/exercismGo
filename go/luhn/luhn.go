package luhn

import (
	"strings"
	"unicode"
)

// Valid func
func Valid(input string) bool {
	input = strings.TrimSpace(input)
	if len(input) <= 1 {
		return false
	}
	reverseString := reverse(input)
	sum := 0

	reverseString = strings.ReplaceAll(reverseString, " ", "")
	for key, val := range reverseString {
		if !unicode.IsDigit(val) {
			return false
		}
		// ascii to int
		val = val - '0'
		if key%2 == 1 {
			if int(val)*2 > 9 {
				sum += (int(val) * 2) - 9

				continue
			}
			sum += int(val) * 2
			continue
		}
		sum += int(val)
	}
	if sum%10 != 0 {
		return false
	}
	return true
}

// reverse fun
func reverse(input string) (result string) {
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
