package p2639

func findColumnWidth(grid [][]int) []int {
	m := len(grid)
	n := len(grid[0])

	ans := make([]int, n)

	for j := 0; j < n; j++ {
		for i := 0; i < m; i++ {
			ans[j] = max(ans[j], digitLength(grid[i][j]))
		}
	}

	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func digitLength(num int) int {
	var res int
	if num < 0 {
		res++
		num = -num
	}
	if num == 0 {
		return 1
	}
	for num > 0 {
		res++
		num /= 10
	}
	return res
}
