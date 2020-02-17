package p1351

func countNegatives(grid [][]int) int {
	m := len(grid)

	if m == 0 {
		return 0
	}

	n := len(grid[0])

	if n == 0 {
		return 0
	}

	var res int

	row, col := 0, n-1

	for row < m && col >= 0 {
		num := grid[row][col]
		if num < 0 {
			col--
		} else {
			res += col + 1
			row++
		}
	}

	return m*n - res

}
