package p1245

func treeDiameter(edges [][]int) int {
	n := len(edges)

	grid := make([][]int, n+1)

	for i := 0; i < n; i++ {
		grid[i] = make([]int, 0, 3)
	}

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		grid[u] = append(grid[u], v)
		grid[v] = append(grid[v], u)
	}

	var dfs func(p int, u int, h int) (int, int)

	dfs = func(p int, u int, h int) (int, int) {

		if p >= 0 && len(grid[u]) == 1 {
			// have parent in
			return u, h
		}
		var i int
		var ih int
		for _, v := range grid[u] {
			if v != p {
				j, jh := dfs(u, v, h+1)
				if jh > ih {
					i = j
					ih = jh
				}
			}
		}
		return i, ih
	}

	x, _ := dfs(-1, 0, 0)

	_, h := dfs(-1, x, 0)

	return h
}
