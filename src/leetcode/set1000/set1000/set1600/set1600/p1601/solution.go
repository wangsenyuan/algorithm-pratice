package p1601

func maximumRequests(n int, requests [][]int) int {
	m := len(requests)

	M := 1 << m

	degree := make([]int, n)
	var ans int
	for state := 1; state < M; state++ {
		var cnt int
		for i := 0; i < m; i++ {
			if state&(1<<i) > 0 {
				cnt++
				req := requests[i]
				degree[req[0]]++
				degree[req[1]]--
			}
		}

		var ok = true
		for i := 0; i < n && ok; i++ {
			ok = degree[i] == 0
		}

		if ok {
			ans = max(ans, cnt)
		}
		for i := 0; i < n; i++ {
			degree[i] = 0
		}
	}

	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
