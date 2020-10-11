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

	var dfs func(p, u int) Pair
	dfs = func(p, u int) Pair {
		vis[u]++
		res := Pair{u, H[u]}
		for _, v := range G[u] {
			if p == v {
				continue
			}
			H[v] = H[u] + 1
			tmp := dfs(u, v)
			if res.second < tmp.second {
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
				r = dfs(-1, i).first
			}
		}

		if cnt == 1 {
			H[r] = 0
			d := dfs(-1, r).second
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

type Pair struct {
	first, second int
}
