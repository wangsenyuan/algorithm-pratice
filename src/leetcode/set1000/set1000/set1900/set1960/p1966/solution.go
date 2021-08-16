package p1966

func numOfStrings(patterns []string, word string) int {
	var res int
	for _, pat := range patterns {
		if isSubString(pat, word) {
			res++
		}
	}

	return res
}

func isSubString(a, b string) bool {
	lps := kmp(a)
	var i, j int
	for i < len(b) {
		if b[i] == a[j] {
			i++
			j++
		}
		if j == len(a) {
			return true
		}
		if i < len(b) && b[i] != a[j] {
			if j > 0 {
				j = lps[j-1]
			} else {
				i++
			}
		}
	}

	return false
}

func kmp(s string) []int {
	n := len(s)
	lps := make([]int, n)

	for i := 1; i < n; i++ {
		j := lps[i-1]

		for j > 0 && s[i] != s[j] {
			j = lps[j-1]
		}

		lps[i] = j
		if s[i] == s[j] {
			lps[i]++
		}
	}

	return lps
}
