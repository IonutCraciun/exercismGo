// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import "strings"

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, "-", " ")
	s = strings.ReplaceAll(s, "_", " ")
	list := strings.Split(s, " ")
	var r string
	for _, x := range(list) {
		x = strings.TrimSpace(x)
		if len(x) > 0 {
			r = r + string(x[0])
		}

	}
	return strings.ToUpper(r)
}
