package p3074

func minimumOperationsToWriteY(grid [][]int) int {
	ans := len(grid) * len(grid)
	for a := 0; a <= 2; a++ {
		for b := 0; b <= 2; b++ {
			if a == b {
				continue
			}
			ans = min(ans, solve(grid, []int{a, b}))
		}
	}
	return ans
}

func solve(grid [][]int, expect []int) int {
	n := len(grid)

	check := func(i, j int) bool {
		if i < n/2 {
			return j == i || j == n-1-i
		}
		// i >= n/2
		return j == n/2
	}

	var res int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if check(i, j) && grid[i][j] != expect[0] {
				res++
			} else if !check(i, j) && grid[i][j] != expect[1] {
				res++
			}
		}
	}
	return res
}
