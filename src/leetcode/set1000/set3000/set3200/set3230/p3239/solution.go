package p3239

func minFlips(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])
	var res1 int
	for i := 0; i < n; i++ {
		for j, k := 0, m-1; j < k; j, k = j+1, k-1 {
			if grid[i][j] != grid[i][k] {
				res1++
			}
		}
	}

	var res2 int

	for j := 0; j < m; j++ {
		for i, k := 0, n-1; i < k; i, k = i+1, k-1 {
			if grid[i][j] != grid[k][j] {
				res2++
			}
		}
	}

	return min(res1, res2)
}
