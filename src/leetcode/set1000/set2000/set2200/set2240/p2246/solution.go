package p2246

func longestPath(parent []int, s string) int {
	n := len(parent)

	dp := make([]int, n)
	fp := make([]int, n)

	deg := make([]int, n)

	for i := 0; i < n; i++ {
		p := parent[i]
		if p >= 0 {
			deg[p]++
		}
		fp[i] = 1
		dp[i] = 1
	}

	que := make([]int, n)
	var front, end int
	for i := 0; i < n; i++ {
		if deg[i] == 0 {
			que[end] = i
			end++
		}
	}

	for front < end {
		u := que[front]
		front++
		p := parent[u]
		if p < 0 {
			break
		}
		if s[p] != s[u] {
			dp[p] = max(dp[p], fp[u]+fp[p])
			fp[p] = max(fp[p], fp[u]+1)
		}
		deg[p]--
		if deg[p] == 0 {
			que[end] = p
			end++
		}
	}
	var best int

	for i := 0; i < n; i++ {
		best = max(best, dp[i])
	}
	return best
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
