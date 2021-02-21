package p1763

func longestNiceSubstring(s string) string {
	n := len(s)
	for l := n; l > 0; l-- {
		for i := 0; i+l <= n; i++ {
			if nice(s[i : i+l]) {
				return s[i : i+l]
			}
		}
	}
	return ""
}

func nice(s string) bool {
	cnt := make([]bool, 26)
	cnt2 := make([]bool, 26)
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			cnt[int(s[i])-'a'] = true
		}
		if s[i] >= 'A' && s[i] <= 'Z' {
			cnt2[int(s[i]-'A')] = true
		}
	}

	for i := 0; i < 26; i++ {
		if cnt[i] != cnt2[i] {
			return false
		}
	}
	return true
}
