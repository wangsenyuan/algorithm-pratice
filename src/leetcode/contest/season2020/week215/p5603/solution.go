package p5603

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	x := getFreq(word1)
	y := getFreq(word2)

	for i := 0; i < 26; i++ {
		if x[i] == y[i] {
			continue
		}
		if x[i] == 0 && y[i] != 0 {
			return false
		}
		if x[i] != 0 && y[i] == 0 {
			return false
		}

		// x[i] != y[i]
		j := i + 1
		for j < 26 && y[j] != x[i] {
			j++
		}
		if j == 26 {
			return false
		}

		y[i], y[j] = y[j], y[i]
	}

	return true
}

func getFreq(s string) []int {
	freq := make([]int, 26)
	for i := 0; i < len(s); i++ {
		x := int(s[i] - 'a')
		freq[x]++
	}
	return freq
}
