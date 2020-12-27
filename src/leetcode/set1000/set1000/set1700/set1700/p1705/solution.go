package p1705

func findBall(grid [][]int) []int {
	m := len(grid)
	n := len(grid[0])

	find := func(j int) int {
		y := j

		for i := 0; i < m; i++ {
			dir := grid[i][y]
			if dir == 1 {
				// \
				if y == n-1 || grid[i][y+1] == -1 {
					return -1
				}
				y++
			} else {
				// /
				if y == 0 || grid[i][y-1] == 1 {
					return -1
				}
				y--
			}
		}
		return y
	}

	ans := make([]int, n)

	for i := 0; i < n; i++ {
		ans[i] = find(i)
	}

	return ans
}
