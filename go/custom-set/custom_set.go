package stringset

//package main

import (
	"fmt"
)

// Set type
type Set map[string]bool

// New set
func New() Set {
	return make(Set)
}

// NewFromSlice set
func NewFromSlice(input []string) Set {
	set := New()
	if len(input) > 0 {
		for _, val := range input {
			set[val] = true
		}
	}
	return set
}

// IsEmpty func
func (s Set) IsEmpty() bool {
	if len(s) == 0 {
		return true
	}
	return false
}

// String func
func (s Set) String() (result string) {
	if s.IsEmpty() {
		return "{}"
	}
	result = "{"
	for key := range s {
		result += fmt.Sprintf("\"%v\", ", key)
	}
	result = result[0:len(result)-2] + "}"
	return result
}

// Has func
func (s Set) Has(input string) bool {
	_, ok := s[input]
	return ok
}

// Subset func
func Subset(s1, s2 Set) bool {
	for key := range s1 {
		if !s2.Has(key) {
			return false
		}
	}
	return true
}

// Disjoint func
func Disjoint(s1, s2 Set) bool {
	for key := range s1 {
		if s2.Has(key) {
			return false
		}
	}
	return true
}

// Equal func
func Equal(s1, s2 Set) bool {
	if len(s1) != len(s2) {
		return false
	}

	for key := range s1 {
		if !s2.Has(key) {
			return false
		}
	}
	return true
}

// Add func
func (s Set) Add(input string) {
	s[input] = true
}

// Intersection func
func Intersection(s1, s2 Set) Set {
	result := New()
	for key := range s1 {
		if s2.Has(key) {
			result[key] = true
		}
	}
	return result
}

// Difference func
func Difference(s1, s2 Set) Set {
	result := New()
	for key := range s1 {
		if !s2.Has(key) {
			result.Add(key)
		}
	}
	return result
}

// Union func
func Union(s1, s2 Set) Set {

	for key := range s1 {
		if !s2.Has(key) {
			s2.Add(key)
		}
	}
	return s2
}

// func main() {
// 	s := NewFromSlice([]string{"1", "2", "3"})
// 	ss := NewFromSlice([]string{"5", "1"})
// 	s1 := NewFromSlice([]string{"1", "2", "3"})
// 	s2 := NewFromSlice([]string{"1", "2", "3", "4"})
// 	clean := New()
// 	fmt.Println(clean)
// 	fmt.Println(s)

// 	if s.Has("2") {
// 		fmt.Println("it has 2")
// 	}

// 	if Subset(ss, s) {
// 		fmt.Println("is subset")
// 	}

// 	if Disjoint(s, ss) {
// 		fmt.Println("are disjoint")
// 	}

// 	if Equal(s1, s2) {
// 		fmt.Println("Equal true")
// 	}

// 	ss.Add("4")
// 	fmt.Println(ss)

// 	fmt.Println(Intersection(s, ss))

// 	fmt.Println(Difference(s, ss))

// }
