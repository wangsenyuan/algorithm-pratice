package p3027

func minimumTimeToInitialState(word string, k int) int {
	n := len(word)

	lps := make([]int, n)

	for i := 1; i < n; i++ {
		j := lps[i-1]
		for j > 0 && word[i] != word[j] {
			j = lps[j-1]
		}

		if word[i] == word[j] {
			j++
		}
		lps[i] = j
	}

	ans := (n + k - 1) / k

	for i := lps[n-1]; i > 0; i = lps[i-1] {
		if (n-i)%k == 0 {
			ans = (n - i) / k
			break
		}
	}

	return ans
}
