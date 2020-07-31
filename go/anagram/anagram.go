package anagram

import (
	"strings"
	"sort"
)

// Detect func
func Detect(subject string, candidates []string) []string {
	result := []string{}
	for _, val := range candidates { 
		if strings.ToLower(subject) == strings.ToLower(val) {
			continue
		}
		if sortString(strings.ToLower(subject)) == sortString(strings.ToLower(val)) {
			result = append(result, val)
		}
	}
	return result
}

// sortString function
func sortString(s string) string {
	temporary := strings.Split(s, "")
	sort.Strings(temporary)
	return strings.Join(temporary, "")
}