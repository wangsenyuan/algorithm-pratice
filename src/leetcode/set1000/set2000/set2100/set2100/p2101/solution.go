package p2101

func maximumDetonation(bombs [][]int) int {
	n := len(bombs)
	g := make([][]int, n)

	add := func(u, v int) {
		if g[u] == nil {
			g[u] = make([]int, 0, 2)
		}
		g[u] = append(g[u], v)
	}

	squareDist := func(u, v int) int64 {
		dx := int64(bombs[v][0] - bombs[u][0])
		dy := int64(bombs[v][1] - bombs[u][1])

		return dx*dx + dy*dy
	}

	canBomb := func(u, v int) bool {
		// explore u, can lead v, xv, yv
		return squareDist(u, v) <= int64(bombs[u][2])*int64(bombs[u][2])
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			if canBomb(i, j) {
				add(i, j)
			}
		}
	}
	vis := make([]int, n)
	var dfs func(u, flag int) int
	dfs = func(u, flag int) int {
		vis[u] = flag
		cnt := 1
		for _, v := range g[u] {
			if vis[v] != flag {
				cnt += dfs(v, flag)
			}
		}
		return cnt
	}

	var best int
	for u := 0; u < n; u++ {
		tmp := dfs(u, u+1)
		best = max(best, tmp)
	}

	return best
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
