package p245

func shortestWordDistance(words []string, word1 string, word2 string) int {
	if word1 == word2 {
		return sameWordDist(words, word1)
	}

	a, b := -1, -1
	c := len(words)
	for i, word := range words {
		if word == word1 {
			a = i
		} else if word == word2 {
			b = i
		}

		if a >= 0 && b >= 0 && abs(a-b) < c {
			c = abs(a - b)
		}
	}

	return c
}

func sameWordDist(words []string, target string) int {
	prev := -1
	res := len(words)

	for i, word := range words {
		if word != target {
			continue
		}

		if prev >= 0 && i-prev < res {
			res = i - prev
		}
		prev = i
	}

	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
