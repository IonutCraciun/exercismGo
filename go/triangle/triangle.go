// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

import "math"

// Kind type of triangle
type Kind string

const (
	// NaT not a triangle
	NaT = "Nat"
	// Equ equilateral
	Equ = "Equ"
	// Iso isosceles
	Iso = "Iso"
	// Sca scalene
	Sca = "Sca"
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	var k Kind
	k = NaT
	if a <= 0 || b <= 0 || c <= 0 {
		return NaT
	}
	if a+b < c || b+c < a || c+a < b {
		return NaT
	}
	if a == math.Inf(1) || a == math.Inf(-1) || a == math.NaN() ||
		b == math.Inf(1) || b == math.Inf(-1) || b == math.NaN() ||
		c == math.Inf(1) || c == math.Inf(-1) || c == math.NaN() {
		return NaT
	}
	if a > 0 && b > 0 && c > 0 {
		k = Sca
	}
	if a == b || b == c || c == a {
		k = Iso
	}
	if a == b && b == c {
		k = Equ
	}

	return k
}
