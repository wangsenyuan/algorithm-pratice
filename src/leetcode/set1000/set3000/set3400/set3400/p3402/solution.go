package p3402

func minimumOperations(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])
	var res int
	for j := 0; j < m; j++ {
		prev := grid[0][j]
		for i := 1; i < n; i++ {
			tmp := max(grid[i][j], prev+1)
			res += tmp - grid[i][j]
			prev = tmp
		}
	}
	return res
}
