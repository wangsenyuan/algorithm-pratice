package p1560

func mostVisited(n int, rounds []int) []int {
	presum := make([]int, n+1)

	presum[rounds[0]-1]++
	presum[rounds[0]]--

	for i := 1; i < len(rounds); i++ {
		a := rounds[i-1]
		b := rounds[i]
		presum[a]++
		if b <= a {
			presum[0]++
		}

		presum[b]--
	}

	for i := 1; i < n; i++ {
		presum[i] += presum[i-1]
	}
	var best int

	for i := 0; i < n; i++ {
		best = max(best, presum[i])
	}

	var res []int

	for i := 0; i < n; i++ {
		if presum[i] == best {
			res = append(res, i+1)
		}
	}
	return res
}

func mostVisited1(n int, rounds []int) []int {
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
