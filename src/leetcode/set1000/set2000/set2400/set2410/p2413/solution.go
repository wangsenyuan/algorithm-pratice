package p2413

func longestContinuousSubstring(s string) int {
	var res int

	n := len(s)

	prev := -n
	var cnt int

	for i := 0; i < n; i++ {
		x := int(s[i]-'a') - i
		if x == prev {
			cnt++
		} else {
			res = max(res, cnt)
			cnt = 1
			prev = x
		}
	}
	res = max(res, cnt)
	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
