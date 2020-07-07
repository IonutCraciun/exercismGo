// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import "strings"

// Hey should have a comment documenting it.
func Hey(remark string) string {


	remark = strings.TrimSpace(remark)

	if strings.ToUpper(remark) == strings.ToLower(remark) {

		if len(remark) == 0 {
			return "Fine. Be that way!"
		}
		if remark[len(remark)-1] == '?' {
			return "Sure."
		}
		return "Whatever."
	}

	if strings.ToUpper(remark) == remark {
		if strings.ToUpper(remark) != strings.ToLower(remark) {
			if remark[len(remark)-1] == '?' {
				return "Calm down, I know what I'm doing!"
			}
			return "Whoa, chill out!"
		}
	} else {
		if remark[len(remark)-1] == '?' {
			return "Sure."
		}
		return "Whatever."
	}
	return "Fine. Be that way!"
}
