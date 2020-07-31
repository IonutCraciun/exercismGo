package grains

import (
	"fmt"
)

// ErrorInvalidInput error  >> error commented
// var ErrorInvalidInput = errors.New("Invalid input")

// Square fun
func Square(input int) (r uint64, err error) {
	if input > 64 || input < 1 {
		return 0, fmt.Errorf("Invalid input: %d", input)
	}
	r = 1 << uint(input-1)
	return
}

// Total func
func Total() (r uint64) {
	// it's (2^n) - 1 or 2^0 + 2^1 + .... + 2^63   --tt
	r = (1 << 64) - 1
	return
}
