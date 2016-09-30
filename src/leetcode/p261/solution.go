package main

func validTree(n int, edges [][]int) bool {
	g := make([][]int, n)

	for _, edge := range edges {
		a, b := edge[0], edge[1]
		g[a] = updateNeighbors(g[a], b)
		g[b] = updateNeighbors(g[b], a)
	}

	var visit func(p, v int, visited []bool) bool

	visit = func(p, v int, visited []bool) bool {
		visited[v] = true

		for _, w := range g[v] {
			if w == p {
				continue
			}
			if visited[w] {
				return false
			}
			if !visit(v, w, visited) {
				return false
			}
		}
		return true
	}

	visited := make([]bool, n)
	if !visit(-1, 0, visited) {
		return false
	}

	for i := 0; i < n; i++ {
		if !visited[i] {
			return false
		}
	}

	return true
}

func updateNeighbors(row []int, x int) []int {
	if row == nil {
		row = []int{x}
	} else {
		row = append(row, x)
	}
	return row
}
