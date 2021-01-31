package p1743

func restoreArray(adjacentPairs [][]int) []int {
	n := len(adjacentPairs)
	n++
	// color each number, when we set one
	next := make([]int, 2*n)
	nodes := make(map[int]int)
	to := make([]int, 2*n)
	var cnt int
	ins := func(u, v int) {
		cnt++
		next[cnt] = nodes[u]
		nodes[u] = cnt
		to[cnt] = v
	}

	degree := make(map[int]int)

	for _, pair := range adjacentPairs {
		u, v := pair[0], pair[1]
		ins(u, v)
		ins(v, u)
		degree[u]++
		degree[v]++
	}

	res := make([]int, n)
	var dfs func(p, u, j int)
	dfs = func(p, u, j int) {
		res[j] = u
		for i := nodes[u]; i > 0; i = next[i] {
			if to[i] == p {
				continue
			}
			dfs(u, to[i], j+1)
		}
	}

	for k, v := range degree {
		if v == 1 {
			// one ends
			dfs(-1, k, 0)
			break
		}
	}
	return res
}
