package p2606

func maximumCostSubstring(s string, chars string, vals []int) int {
	sc := make([]int, 26)
	for i := 0; i < 26; i++ {
		sc[i] = i + 1
	}
	for i := 0; i < len(chars); i++ {
		x := int(chars[i] - 'a')
		sc[x] = vals[i]
	}

	var best int
	var sum int
	for i := 0; i < len(s); i++ {
		x := int(s[i] - 'a')
		y := sc[x]
		sum += y
		best = max(best, sum)
		if sum < 0 {
			sum = 0
		}
	}
	return best
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
