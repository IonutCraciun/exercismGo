// Package twofer exposes utilities to ensure everyone gets one
package twofer

import "fmt"

// ShareWith receives a string and return a response based on the content of the string
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
