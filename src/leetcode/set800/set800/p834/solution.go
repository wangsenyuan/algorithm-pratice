package p834

func sumOfDistancesInTree(N int, edges [][]int) []int {
	ns := make([]map[int]bool, N)

	for i := 0; i < N; i++ {
		ns[i] = make(map[int]bool)
	}

	for _, edge := range edges {
		a, b := edge[0], edge[1]
		ns[a][b] = true
		ns[b][a] = true
	}

	cc := make([]int, N)
	pp := make([]int, N)
	ans := make([]int, N)
	var dfs func(p, u, h int)

	dfs = func(p, u, h int) {
		pp[u] = p
		ans[0] += h
		cnt := 1
		for v := range ns[u] {
			if p == v {
				continue
			}
			dfs(u, v, h+1)
			cnt += cc[v]
		}

		cc[u] = cnt
	}

	dfs(-1, 0, 0)

	var cal func(p, u int)

	cal = func(p, u int) {

		for v := range ns[u] {
			if p == v {
				continue
			}
			vcnt := cc[v]
			acnt := N - vcnt
			ans[v] = ans[u] + acnt - vcnt
			cal(u, v)
		}
	}

	cal(-1, 0)

	return ans
}
