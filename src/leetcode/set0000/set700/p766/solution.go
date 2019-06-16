package p766

func isToeplitzMatrix(matrix [][]int) bool {
	for i := 0; i < len(matrix); i++ {
		for j := 1; i+j < len(matrix) && j < len(matrix[i]); j++ {
			if matrix[i+j][j] != matrix[i][0] {
				return false
			}
		}
	}

	for j := 0; j < len(matrix[0]); j++ {
		for i := 1; i < len(matrix) && j+i < len(matrix[0]); i++ {
			if matrix[i][i+j] != matrix[0][j] {
				return false
			}
		}
	}
	return true
}
