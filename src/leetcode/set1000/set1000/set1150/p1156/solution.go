package p1156

func maxRepOpt1(text string) int {
	cnt := make([]int, 26)

	for i := 0; i < len(text); i++ {
		cnt[text[i]-'a']++
	}

	var best int
	for i := 0; i < len(text); {
		j := i + 1
		for j < len(text) && text[j] == text[i] {
			j++
		}
		best = max(best, j-i)

		if j == len(text) {
			break
		}
		// record pos
		k := j

		if cnt[text[i]-'a'] > j-i {
			// has more text[i] to swap in
			best = max(best, j-i+1)
		} else {
			// no way to get a better answer for text[i], skip it
			i = k
			continue
		}

		j++

		for j < len(text) && text[j] == text[i] {
			j++
		}
		best = max(best, j-i-1)

		if cnt[text[i]-'a'] >= j-i {
			best = max(best, j-i)
		}

		i = k
	}

	return best
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
