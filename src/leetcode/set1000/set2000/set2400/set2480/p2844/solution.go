package p2844

func onesMinusZeros(grid [][]int) [][]int {
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

	diff := make([][]int, m)
	for i := 0; i < m; i++ {
		diff[i] = make([]int, n)
		for j := 0; j < n; j++ {
			diff[i][j] = row[i] + col[j] - (n - row[i]) - (m - col[j])
		}
	}

	return diff
}
