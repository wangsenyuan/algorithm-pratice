package p1879

func countGoodSubstrings(s string) int {
	var res int
	for i := 2; i < len(s); i++ {
		if s[i] != s[i-1] && s[i] != s[i-2] && s[i-1] != s[i-2] {
			res++
		}
	}

	return res
}
