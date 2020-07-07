package accumulate

type converter func(string) string

// Accumulate is doing function on all things from list
func Accumulate(list []string, function converter) []string {
	var result []string
	if len(list) == 0 {
		return list
	}
	for _, val := range list {
		result = append(result, function(val))
	}
	return result
}
