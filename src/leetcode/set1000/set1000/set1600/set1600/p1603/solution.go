package p1603

func maximalNetworkRank(n int, roads [][]int) int {
	degree := make([]int, n)
	adj := make([]map[int]bool, n)
	for i := 0; i < n; i++ {
		adj[i] = make(map[int]bool)
	}
	for _, road := range roads {
		u, v := road[0], road[1]
		degree[u]++
		degree[v]++
		adj[u][v] = true
		adj[v][u] = true
	}

	var best int

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j || degree[i]+degree[j] <= best {
				continue
			}
			tmp := degree[i] + degree[j]
			if adj[i][j] {
				tmp--
			}
			if tmp > best {
				best = tmp
			}
		}
	}
	return best
}
