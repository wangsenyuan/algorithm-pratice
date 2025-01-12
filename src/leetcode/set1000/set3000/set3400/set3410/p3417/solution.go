package p3417

func zigzagTraversal(grid [][]int) []int {
	n := len(grid)
	m := len(grid[0])
	var res []int

	for i := 0; i < n; i++ {
		if i&1 == 0 {
			for j := 0; j < m; j++ {
				res = append(res, grid[i][j])
			}
		} else {
			for j := m - 1; j >= 0; j-- {
				res = append(res, grid[i][j])
			}
		}
	}
	var ans []int
	for i := 0; i < len(res); i += 2 {
		ans = append(ans, res[i])
	}
	return ans
}
