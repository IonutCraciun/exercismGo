package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency func
func ConcurrentFrequency(s []string) FreqMap {
	result := FreqMap{}
	wordResult := make(chan FreqMap, 3)

	// Go through all the values of the slice 's'
	for _, val := range s {
		// anonymous function / inline that is putting the frequency map into out channel
		// the channel has length 3 -> that means it needs 3 results to be blocked at line 31
		// without length, it will not be real "concurency/parelism"
		go func(val string, out chan FreqMap) {
			out <- Frequency(val)
		}(val, wordResult)
	}

	for i := 0; i < len(s); i++ {
		for k, v := range <-wordResult {
			result[k] += v
		}
	}
	return result
}
