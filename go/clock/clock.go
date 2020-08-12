package clock

import (
	"fmt"
)

// Clock type
type Clock struct {
	minutes int
}

// New -> Mic's New returns a new Clock
func New(hour, minute int) Clock {
	var minutes = (hour * 60) + minute
	minutes %= 24 * 60
	if minutes < 0 {
		minutes += 24 * 60
	}

	return Clock{minutes: minutes}
}

// String func
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.minutes/60, c.minutes%60)
}

// Add func
func (c Clock) Add(minutes int) Clock {
	return New(c.minutes/60, c.minutes%60+minutes)
}

// Subtract func
func (c Clock) Subtract(minutes int) (result Clock) {
	return New(c.minutes/60, c.minutes%60-minutes)
}
