package p1560

func mostVisited(n int, rounds []int) []int {
	cnt := make([]int, n+1)

	cnt[rounds[0]]++
	for i := 1; i < len(rounds); i++ {
		a := rounds[i-1]
		b := rounds[i]
		if a+1 <= b {
			for j := a + 1; j <= b; j++ {
				cnt[j]++
			}
		} else {
			for j := a + 1; j <= n; j++ {
				cnt[j]++
			}
			for j := 1; j <= b; j++ {
				cnt[j]++
			}
		}
	}
	var best int
	for i := 1; i <= n; i++ {
		best = max(best, cnt[i])
	}

	var res []int

	for i := 1; i <= n; i++ {
		if best == cnt[i] {
			res = append(res, i)
		}
	}
	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
