package p1380

func luckyNumbers(matrix [][]int) []int {
	row := make([]int, len(matrix))
	col := make([]int, len(matrix[0]))

	for i := 0; i < len(row); i++ {
		row[i] = 1000001
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			row[i] = min(row[i], matrix[i][j])
			col[j] = max(col[j], matrix[i][j])
		}
	}

	var res []int

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == row[i] && matrix[i][j] == col[j] {
				res = append(res, matrix[i][j])
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

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
