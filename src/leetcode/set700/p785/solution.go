package p785

func isBipartite(graph [][]int) bool {
	n := len(graph)

	colors := make([]int, n)

	var dfs func(v int, color int) bool

	dfs = func(v int, color int) bool {
		colors[v] = color

		for _, w := range graph[v] {
			if colors[w] == color {
				return false
			}
			if colors[w] == 0 {
				tmp := dfs(w, -color)
				if !tmp {
					return false
				}
			}
		}

		return true
	}

	for i := 0; i < n; i++ {
		if colors[i] == 0 {
			tmp := dfs(i, 1)
			if !tmp {
				return false
			}
		}
	}
	return true
}
