package p1072

func maxEqualRowsAfterFlips(matrix [][]int) int {
	// row[i][0] = no of zeros
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	m := len(matrix)
	n := len(matrix[0])

	A := make([][]int, m)
	for i := 0; i < m; i++ {
		A[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if matrix[i][j] != matrix[i][0] {
				A[i][j] = 1
			}
		}
	}
	row := make([]int, m)

	for i := 0; i < m; i++ {
		if row[i] == 0 {
			row[i] = 1
			for j := i + 1; j < m; j++ {
				var k int
				for k < n && A[i][k] == A[j][k] {
					k++
				}
				if k == n {
					row[i]++
					row[j] = 1
				}
			}
		}
	}
	var ans int

	for i := 0; i < m; i++ {
		ans = max(ans, row[i])
	}

	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
