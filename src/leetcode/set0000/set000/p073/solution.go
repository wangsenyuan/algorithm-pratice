package main

func setZeroes(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])
	col0 := false
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			col0 = true
		}
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0], matrix[0][j] = 0, 0
			}
		}
	}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 1; j-- {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
		if col0 {
			matrix[i][0] = 0
		}
	}
}
