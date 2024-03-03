package p3075

func countSubmatrices(grid [][]int, k int) int {
	n := len(grid)
	m := len(grid[0])
	sum := make([][]int, n)
	var res int
	for i := 0; i < n; i++ {
		sum[i] = make([]int, m)
		for j := 0; j < m; j++ {
			sum[i][j] = grid[i][j]
			if i > 0 {
				sum[i][j] += sum[i-1][j]
			}
			if j > 0 {
				sum[i][j] += sum[i][j-1]
			}
			if i > 0 && j > 0 {
				sum[i][j] -= sum[i-1][j-1]
			}
			if sum[i][j] <= k {
				res++
			}
		}
	}
	return res
}
