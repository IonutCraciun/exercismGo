package wordcount

import (
	"regexp"
	"strings"
)

// Frequency type
type Frequency map[string]int

var reSeparator = regexp.MustCompile(`[^\w']+`)

// WordCount fun
func WordCount(s string) Frequency {
	s = strings.ToLower(s)
	result := make(Frequency)
	tokens := reSeparator.Split(s, -1)
	for _, val := range tokens {
		if val != "" {
			val = strings.TrimSpace(val)
			val = strings.Trim(val, "'")
			result[val]++
		}
	}
	return result
}
