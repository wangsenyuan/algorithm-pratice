package main

func oddString(words []string) string {
	n := len(words[0])
	diff := calcDiff(words[0])
	for j, word := range words {
		if j == 0 {
			continue
		}
		diff1 := calcDiff(word)
		for i := 0; i < n-1; i++ {
			if diff1[i] != diff[i] {
				if j > 1 {
					return word
				}
				// j == 1
				diff2 := calcDiff(words[j+1])
				for ii := 0; ii < n-1; ii++ {
					if diff2[ii] != diff1[ii] {
						return word
					}
				}
				return words[0]
			}
		}
	}
	return words[0]
}

func calcDiff(word string) []int {
	n := len(word)
	diff := make([]int, n-1)
	for i := 0; i+1 < n; i++ {
		diff[i] = int(word[i+1] - word[i])
	}
	return diff
}
