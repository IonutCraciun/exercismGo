package flatten

// Flatten func
func Flatten(input interface{}) []interface{} {
	// empty []interface{} slice
	result := []interface{}{}

	// get every element of the interface with type assert
	for _, elem := range input.([]interface{}) {
		// if element is null, go to new element
		if elem == nil {
			continue
		}
		// check the type with type assertion
		switch val := elem.(type) {
		case int:
			// if it's int, add to results
			result = append(result, val)
		default:
			// else append to results recursive Flatten(the rest of input)
			// ... becasue Flatten can return multiple results like 1, 2, 3
			result = append(result, Flatten(val.([]interface{}))...)
		}
	}
	return result
}