package encode

import (
	"unicode"
	"strconv"
)

// RunLengthDecode func
func RunLengthDecode(input string) string {
	var result, number string
	number = ""
	for _, val := range input {
		if unicode.IsDigit(val) {
			number += string(val)
		} else  { //means it is letter 
			if len(number) == 0 {
				result += string(val)
			} else {
				n, _ := strconv.Atoi(number)
				for i:=0; i < n; i++ {
					result += string(val)
				}
			}
			number = ""
		}
	}
	return result
}

// RunLengthEncode func
func RunLengthEncode(input string) string {
	var result string
	var total int
	if len(input) == 0 {
		return ""
	}
	first := rune(input[0])
	total = 1
	for _, val := range input[1:] {
		if first != val {
			if total != 1 {
				result += strconv.Itoa(total)
			}
			result += string(first)
			first = val
			total = 1
		} else {
			total++
		}
	}
	if total != 1 {
		result += strconv.Itoa(total)
	}
	result += string(first)

	return string(result)
}