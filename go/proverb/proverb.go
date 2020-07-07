package proverb

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
	results := make([]string, len(rhyme))
	if len(rhyme) == 0 {
		return []string{}
	} else if len(rhyme) == 1 {
		results[0] = "And all for the want of a " + rhyme[0] + "."
		return results
	}

	for i := 0; i < len(rhyme)-1; i++ {
		results[i] = "For want of a " + rhyme[i] + " the " + rhyme[i+1] + " was lost."
	}
	results[len(rhyme)-1] = "And all for the want of a " + rhyme[0] + "."
	return results
}
