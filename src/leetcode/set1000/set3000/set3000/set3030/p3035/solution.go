package p3035

func modifiedMatrix(matrix [][]int) [][]int {
	m := len(matrix[0])
	col := make([]int, m)
	for i := 0; i < m; i++ {
		col[i] = -2
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < m; j++ {
			col[j] = max(col[j], matrix[i][j])
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] < 0 {
				matrix[i][j] = col[j]
			}
		}
	}
	return matrix
}
