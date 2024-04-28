package p3129

func numberOfRightTriangles(grid [][]int) int64 {
	n := len(grid)
	m := len(grid[0])

	row := make([]int, n)
	col := make([]int, m)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			row[i] += grid[i][j]
			col[j] += grid[i][j]
		}
	}

	var res int
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				res += (row[i] - 1) * (col[j] - 1)
			}
		}
	}

	return int64(res)
}
