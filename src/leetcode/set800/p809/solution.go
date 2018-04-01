package p809

func expressiveWords(S string, words []string) int {
	a, b := groupsOfWord(S)

	check := func(word string) bool {
		c, d := groupsOfWord(word)
		if len(c) != len(a) {
			return false
		}

		for i := 0; i < len(a); i++ {
			if a[i] != c[i] {
				return false
			}
			if b[i] < d[i] {
				return false
			}
			if b[i] != d[i] && b[i] < 3 {
				return false
			}
		}
		return true
	}

	var ans int

	for _, word := range words {
		if check(word) {
			ans++
		}
	}
	return ans
}

func groupsOfWord(s string) (letters []byte, cnt []int) {
	n := len(s)

	letters = make([]byte, 0, n)
	cnt = make([]int, 0, n)

	for i, j := 0, 0; i <= n; i++ {
		if i < n && s[i] == s[j] {
			continue
		}
		letters = append(letters, s[j])
		cnt = append(cnt, i-j)
		j = i
	}
	return
}
