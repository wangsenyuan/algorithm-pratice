package p1466

func minReorder(n int, connections [][]int) int {

	conn := make([][]int, n)
	for i := 0; i < n; i++ {
		conn[i] = make([]int, 0, 3)
	}

	for _, e := range connections {
		u, v := e[0], e[1]
		conn[u] = append(conn[u], v)
		conn[v] = append(conn[v], u)
	}
	P := make([]int, n)
	var dfs func(p, u int)
	dfs = func(p, u int) {
		P[u] = p
		for _, v := range conn[u] {
			if p == v {
				continue
			}
			dfs(u, v)
		}
	}

	dfs(-1, 0)
	var res int
	for _, e := range connections {
		u, v := e[0], e[1]
		if P[v] == u {
			res++
		}
	}

	return res
}
