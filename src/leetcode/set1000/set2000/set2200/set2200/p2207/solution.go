package p2207

func maximumSubsequenceCount(text string, pattern string) int64 {
	x, y := pattern[0], pattern[1]

	var best int64

	cnt := make([]int64, 2)

	n := len(text)

	var sum int64

	for i := n - 1; i >= 0; i-- {
		if text[i] == x {
			sum += cnt[1]
		}
		if text[i] == y {
			cnt[1]++
		}
	}
	// at the very begining
	best = max(best, sum+cnt[1])
	sum = 0
	for i := 0; i < n; i++ {
		if text[i] == y {
			sum += cnt[0]
		}
		if text[i] == x {
			cnt[0]++
		}
	}
	best = max(best, sum+cnt[0])

	return best
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}
