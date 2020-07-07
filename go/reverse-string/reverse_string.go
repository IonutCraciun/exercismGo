package reverse



// Reverse function
func Reverse(input string) string {
	runes := make([]rune, len(input))
	if len(input) == 0 {
		return ""
	}
	for k, val := range input {
		runes[len(input)-1-k] = val
	}
	return string(runes)
}