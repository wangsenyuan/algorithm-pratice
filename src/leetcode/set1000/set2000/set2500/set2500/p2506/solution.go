package p2506

func isPossible(n int, edges [][]int) bool {
	deg := make([]int, n+1)
	adj := make([]map[int]bool, n+1)
	for i := 1; i <= n; i++ {
		adj[i] = make(map[int]bool)
	}
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		deg[u]++
		deg[v]++
		adj[u][v] = true
		adj[v][u] = true
	}

	var odds []int

	for i := 1; i <= n; i++ {
		if deg[i]%2 == 1 {
			odds = append(odds, i)
		}
	}

	if len(odds) > 4 {
		return false
	}

	if len(odds) == 0 {
		return true
	}

	if len(odds)%2 == 1 {
		return false
	}

	if len(odds) == 2 {
		u := odds[0]
		v := odds[1]
		// find x,
		if !adj[u][v] {
			return true
		}
		for i := 1; i <= n; i++ {
			if !adj[u][i] && !adj[i][v] {
				return true
			}
		}
		return false
	}

	// len(odds) = 4

	a, b, c, d := odds[0], odds[1], odds[2], odds[3]
	if !adj[a][b] && !adj[c][d] {
		return true
	}

	if !adj[a][c] && !adj[b][d] {
		return true
	}

	if !adj[a][d] && !adj[b][c] {
		return true
	}

	return false
}
