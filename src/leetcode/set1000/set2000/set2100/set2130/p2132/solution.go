package p2132

func longestPalindrome(words []string) int {
	cnt := make([][]int, 26)
	for i := 0; i < 26; i++ {
		cnt[i] = make([]int, 26)
	}

	for _, w := range words {
		a, b := int(w[0]-'a'), int(w[1]-'a')
		cnt[a][b]++
	}

	var res int
	for a := 0; a < 26; a++ {
		for b := a + 1; b < 26; b++ {
			res += min(cnt[a][b], cnt[b][a]) * 4
		}
	}
	var res2 int
	for a := 0; a < 26; a++ {
		res += 4 * (cnt[a][a] / 2)
		res2 = max(res2, cnt[a][a]%2)
	}

	return res + 2*res2
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	return a + b - min(a, b)
}
