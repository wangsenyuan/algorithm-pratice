package p2373

func largestLocal(grid [][]int) [][]int {
	m := len(grid)
	n := len(grid[0])
	res := make([][]int, m-2)
	for i := 0; i < m-2; i++ {
		res[i] = make([]int, n-2)

		for j := 0; j < n-2; j++ {
			for u := 0; u < 3; u++ {
				for v := 0; v < 3; v++ {
					res[i][j] = max(res[i][j], grid[i+u][j+v])
				}
			}
		}
	}

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
