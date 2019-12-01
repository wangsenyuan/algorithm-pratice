package p1277

func countSquares(matrix [][]int) int {
	m := len(matrix)
	n := len(matrix[0])
	sum := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		sum[i] = make([]int, n+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			sum[i+1][j+1] = sum[i+1][j] + sum[i][j+1] - sum[i][j] + matrix[i][j]
		}
	}

	var res int

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				continue
			}
			x := min(m-i, n-j)
			for k := 1; k <= x; k++ {
				a := sum[i+k][j+k] - sum[i+k][j] - sum[i][j+k] + sum[i][j]
				if a != k*k {
					break
				}
				res++
			}
		}
	}

	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
