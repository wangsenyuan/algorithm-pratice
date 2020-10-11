package p1605

func countSubgraphsForEachDiameter(n int, edges [][]int) []int {
	res := make([]int, n-1)
	N := 1 << (n - 1)
	G := make([][]int, n)
	for i := 0; i < n; i++ {
		G[i] = make([]int, 0, n)
	}
	vis := make([]int, n)
	for i := 0; i < n; i++ {
		vis[i] = -1
	}
	H := make([]int, n)

	var dfs func(u int) int
	dfs = func(u int) int {
		r := u
		vis[u]++
		for _, v := range G[u] {
			if vis[v] == 0 {
				H[v] = H[u] + 1
				tmp := dfs(v)
				if H[tmp] > H[r] {
					r = tmp
				}
			}
		}
		return r
	}

	var dfs2 func(p, u int) int

	dfs2 = func(p, u int) int {
		res := H[u]

		for _, v := range G[u] {
			if v == p {
				continue
			}
			H[v] = H[u] + 1
			tmp := dfs2(u, v)
			if tmp > res {
				res = tmp
			}
		}
		return res
	}

	for state := 1; state < N; state++ {
		for i := 0; i < n-1; i++ {
			if state&(1<<i) > 0 {
				e := edges[i]
				u, v := e[0]-1, e[1]-1
				G[u] = append(G[u], v)
				G[v] = append(G[v], u)
				vis[u] = 0
				vis[v] = 0
			}
		}

		var cnt int
		var r int
		for i := 0; i < n && cnt < 2; i++ {
			if vis[i] == 0 {
				cnt++
				H[i] = 0
				r = dfs(i)
			}
		}

		if cnt == 1 {
			H[r] = 0
			d := dfs2(-1, r)
			if d > 0 {
				res[d-1]++
			}
		}

		for i := 0; i < n; i++ {
			vis[i] = -1
			G[i] = G[i][:0]
		}
	}

	return res
}
