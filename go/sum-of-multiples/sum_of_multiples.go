package summultiples

// SumMultiples fun
func SumMultiples(limit int, divisors ...int) (sum int) {
	for i:=1; i < limit; i++ {
		for _, val := range divisors {
			if val != 0 && i % val == 0 {
				sum += i
				break
			}
		}
	}
	return
}