package p2428

func maxSum(grid [][]int) int {
	var res int

	m := len(grid)
	n := len(grid[0])

	calc := func(x, y int) int {
		var sum int
		sum += grid[x-1][y-1] + grid[x-1][y] + grid[x-1][y+1]
		sum += grid[x][y]
		sum += grid[x+1][y-1] + grid[x+1][y] + grid[x+1][y+1]
		return sum
	}

	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			res = max(res, calc(i, j))
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
