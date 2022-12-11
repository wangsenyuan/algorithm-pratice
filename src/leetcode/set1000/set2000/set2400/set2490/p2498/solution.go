package p2498

import "sort"

func deleteGreatestValue(grid [][]int) int {
	m := len(grid)

	for i := 0; i < m; i++ {
		sort.Ints(grid[i])
	}

	var res int

	for j := len(grid[0]) - 1; j >= 0; j-- {
		var mx int
		for i := 0; i < m; i++ {
			mx = max(grid[i][j], mx)
		}
		res += mx
	}
	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
