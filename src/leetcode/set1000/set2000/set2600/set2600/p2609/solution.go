package p2609

func findTheLongestBalancedSubstring(s string) int {
	var res int
	n := len(s)
	for i, j := 0, 0; i < n; i++ {
		if s[i] == '0' {
			continue
		}
		// s[i] = '1'
		k := i
		for i < n && s[i] == '1' {
			i++
		}

		res = max(res, min(k-j, i-k)*2)
		j = i
		i--
	}

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	return a + b - max(a, b)
}
