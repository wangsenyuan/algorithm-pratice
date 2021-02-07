package p1750

func minimumLength(s string) int {
	n := len(s)
	i, j := 0, n-1

	for i < j && s[i] == s[j] {
		for i < j && s[i] == s[j] {
			i++
		}
		if i == j {
			return 0
		}
		for i < j && s[i-1] == s[j] {
			j--
		}
	}

	return j - i + 1
}
