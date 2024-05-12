package p3143

const oo = 1 << 40

func maxScore(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])

	// 每个格子左上角区域的最小值
	col := make([]int, m)
	// 到r行的时候，每一列的最小值
	copy(col, grid[0])

	var best int = -oo

	for i := 0; i < n; i++ {
		tmp := oo
		for j := 0; j < m; j++ {
			best = max(best, grid[i][j]-tmp)
			if i > 0 {
				best = max(best, grid[i][j]-col[j])
			}
			col[j] = min(col[j], grid[i][j])
			tmp = min(tmp, col[j])
		}
	}
	return best
}
