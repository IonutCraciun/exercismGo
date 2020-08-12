package secret

/*
1 = wink
10 = double blink
100 = close your eyes
1000 = jump
10000 = Reverse the order of the operations in the secret handshake.

*/

import (
	"fmt"
	"strconv"
)

type move struct {
	i  uint
	mv string
}

var moves = []move{
	{1, "wink"},
	{2, "double blink"},
	{4, "close your eyes"},
	{8, "jump"},
}

// Handshake func
func Handshake(input uint) (result []string) {
	for _, val := range moves {
		if input&val.i == val.i {
			result = append(result, val.mv)
		}
	}
	if input&16 == 16 {
		result = reverseString(result)
	}
	return result
}

func reverseString(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
