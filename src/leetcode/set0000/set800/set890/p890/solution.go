package p890

func possibleBipartition(N int, dislikes [][]int) bool {
	rel := make([]map[int]bool, N)
	for i := 0; i < N; i++ {
		rel[i] = make(map[int]bool)
	}

	for _, dislike := range dislikes {
		a, b := dislike[0]-1, dislike[1]-1
		rel[a][b] = true
		rel[b][a] = true
	}
	color := make([]int, N)
	var dfs func(u int, c int) bool

	dfs = func(u, c int) bool {
		if color[u] != 0 {
			return color[u] == c
		}
		color[u] = c
		for v := range rel[u] {
			if !dfs(v, -c) {
				return false
			}
		}
		return true
	}

	for i := 0; i < N; i++ {
		if color[i] == 0 {
			if !dfs(i, 1) {
				return false
			}
		}
	}
	return true
}
