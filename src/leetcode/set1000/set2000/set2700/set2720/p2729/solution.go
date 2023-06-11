package p2729

func longestSemiRepetitiveSubstring(s string) int {
	n := len(s)
	L := make([]int, n)
	L[0] = 0
	for i := 1; i < n; i++ {
		if s[i] == s[i-1] {
			L[i] = i
		} else {
			L[i] = L[i-1]
		}
	}
	var res int
	R := n - 1
	for i := n - 2; i >= 0; i-- {
		if s[i] == s[i+1] {
			res = max(res, R-L[i]+1)
			R = i
		}
	}

	if res == 0 {
		res = n
	}

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
