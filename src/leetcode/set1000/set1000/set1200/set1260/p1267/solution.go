package p1267

func countServers(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	row := make([]int, m)
	col := make([]int, n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			row[i] += grid[i][j]
			col[j] += grid[i][j]
		}
	}

	var res int

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] > 0 && (row[i] > 1 || col[j] > 1) {
				res++
			}
		}
	}

	return res
}
