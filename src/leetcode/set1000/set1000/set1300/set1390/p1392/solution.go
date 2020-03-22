package p1392

func longestPrefix(s string) string {
	n := len(s)
	lps := make([]int, n)

	var l int

	for i := 1; i < n; {
		if s[i] == s[l] {
			l++
			lps[i] = l
			i++
			continue
		}
		if l != 0 {
			l = lps[l-1]
		} else {
			i++
		}
	}

	return s[n-lps[n-1] : n]
}
