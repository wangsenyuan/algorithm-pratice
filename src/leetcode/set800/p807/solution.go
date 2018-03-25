package p807

func maxIncreaseKeepingSkyline(grid [][]int) int {
	n := len(grid)

	sr := make([]int, n)
	sc := make([]int, n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] > sr[i] {
				sr[i] = grid[i][j]
			}
			if grid[i][j] > sc[j] {
				sc[j] = grid[i][j]
			}
		}
	}

	var res int

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			x := min(sr[i], sc[j])
			res += x - grid[i][j]
		}
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
