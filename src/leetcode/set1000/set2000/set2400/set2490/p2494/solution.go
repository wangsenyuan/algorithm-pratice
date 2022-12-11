package p2494

func maximumValue(strs []string) int {
	var best int

	for _, s := range strs {
		best = max(best, change(s))
	}

	return best
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func change(s string) int {
	var sum int
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			sum = sum*10 + int(s[i]-'0')
		} else {
			return len(s)
		}
	}
	return sum
}
