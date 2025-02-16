package p3456

func hasSpecialSubstring(s string, k int) bool {
	n := len(s)
	for i := 0; i < n; {
		j := i
		for i < n && s[i] == s[j] {
			i++
		}
		if i-j == k {
			return true
		}
	}
	return false
}
