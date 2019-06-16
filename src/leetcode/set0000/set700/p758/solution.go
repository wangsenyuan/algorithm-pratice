package p758

func boldWords(words []string, S string) string {
	dict := make(map[string]bool)
	maxLen := 0
	for _, word := range words {
		dict[word] = true
		if len(word) > maxLen {
			maxLen = len(word)
		}
	}
	n := len(S)
	f := make([]int, n)

	for i := 0; i < n; i++ {
		right := i
		for j := 1; i+j <= n && j <= maxLen; j++ {
			if dict[S[i:i+j]] {
				right = i + j
			}
		}
		f[i] = right
	}
	var res string
	for i := 0; i < n; i++ {
		right := f[i]
		for j := i + 1; j < n && j <= right; j++ {
			if f[j] > right {
				right = f[j]
			}
		}
		if right > i {
			res += "<b>" + S[i:right] + "</b>"
			i = right - 1
		} else {
			res += S[i : i+1]
		}
	}
	return res
}
