package p1377

func frogPosition(n int, edges [][]int, t int, target int) float64 {
	conns := make([][]int, n)

	for i := 0; i < n; i++ {
		conns[i] = make([]int, 0, 3)
	}

	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		conns[u] = append(conns[u], v)
		conns[v] = append(conns[v], u)
	}

	depth := make([]int, n)
	var dfs func(p, u int)

	dp := make([]float64, n)

	dp[0] = 1
	leaf := make([]bool, n)

	dfs = func(p, u int) {
		c := len(conns[u])
		if p >= 0 {
			c--
		}
		for _, v := range conns[u] {
			if p == v {
				continue
			}
			dp[v] = dp[u] * (1 / float64(c))
		}
		leaf[u] = true
		for _, v := range conns[u] {
			if p == v {
				continue
			}
			depth[v] = depth[u] + 1
			leaf[u] = false
			dfs(u, v)
		}
	}

	dfs(-1, 0)

	target--

	if depth[target] > t {
		return 0
	}

	if depth[target] == t || leaf[target] {
		return dp[target]
	}

	return 0
}
