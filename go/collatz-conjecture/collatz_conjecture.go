package collatzconjecture

import "errors"

// CollatzConjecture function that returns the number of steps for collats conjecture
func CollatzConjecture(x int) (steps int, error error) {
	steps = 0

	if x <= 0 {
		return 0, errors.New("negative number or 0")
	}
	for x != 1 {
		if x%2 == 0 {
			x = x / 2
			steps++
		} else {
			x = x*3 + 1
			steps++
		}
	}
	return steps, nil
}
