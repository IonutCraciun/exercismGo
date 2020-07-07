package strain

type pred func(int) bool
type pred2 func(string) bool
type pred3 func([]int) bool

// Ints type
type Ints []int

// Strings type
type Strings []string

// Lists type
type Lists [][]int

// Keep method
func (list Ints) Keep(predicate pred) Ints {
	var results Ints
	for _, val := range list {
		if predicate(val) {
			results = append(results, val)
		}

	}
	return results
}

//Discard method
func (list Ints) Discard(predicate pred) Ints {
	var results Ints
	for _, val := range list {
		if !predicate(val) {
			results = append(results, val)
		}

	}
	return results
}

// Keep method
func (list Strings) Keep(predicate pred2) Strings {
	var results Strings
	for _, val := range list {
		if predicate(val) {
			results = append(results, val)
		}
	}
	return results
}

//Discard method
func (list Strings) Discard(predicate pred2) Strings {
	var results Strings
	for _, val := range list {
		if !predicate(val) {
			results = append(results, val)
		}

	}
	return results
}

// Keep method
func (list Lists) Keep(predicate pred3) Lists {
	var results Lists
	for _, val := range list {
		if predicate(val) {
			results = append(results, val)
		}
	}
	return results
}

//Discard method
func (list Lists) Discard(predicate pred3) Lists {
	var results Lists
	for _, val := range list {
		if !predicate(val) {
			results = append(results, val)
		}

	}
	return results
}
