package p1443

func minTime(n int, edges [][]int, hasApple []bool) int {

	conns := make([][]int, n)

	for i := 0; i < n; i++ {
		conns[i] = make([]int, 0, 10)
	}

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		conns[u] = append(conns[u], v)
		conns[v] = append(conns[v], u)
	}

	var dfs func(p, u int) (int, int)

	dfs = func(p, u int) (int, int) {
		var dist int
		var cnt int

		if hasApple[u] {
			cnt++
		}

		for _, v := range conns[u] {
			if p == v {
				continue
			}
			a, b := dfs(u, v)
			if a == 0 {
				continue
			}
			cnt += a
			// 2 is for reach v and back
			dist += b + 2
		}

		return cnt, dist
	}

	cnt, dist := dfs(-1, 0)

	if cnt == 0 {
		return 0
	}
	return dist
}
