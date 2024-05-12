package p3147

func satisfiesConditions(grid [][]int) bool {
	n := len(grid)
	m := len(grid[0])

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i+1 < n && grid[i][j] != grid[i+1][j] {
				return false
			}
			if j+1 < m && grid[i][j] == grid[i][j+1] {
				return false
			}
		}
	}

	return true
}
