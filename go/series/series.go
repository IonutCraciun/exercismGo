package series

// All func
func All(n int, s string) []string {
	result := []string{}
	for i:=0; i+n <= len(s); i++ {
		result = append(result, s[i:i+n])
	}
	return result
}

// First func
func First(n int, s string) (first string, ok bool) {
	if len(s) < n {
		return "", false
	}
	return s[0:n], true
}

// UnsafeFirst func
func UnsafeFirst(n int, s string) string {
	return s[0:n]
}