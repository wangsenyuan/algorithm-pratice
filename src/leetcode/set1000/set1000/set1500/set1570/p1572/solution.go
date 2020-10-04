package p1572

func diagonalSum(mat [][]int) int {
	m := len(mat)
	if m == 0 {
		return 0
	}
	var res int
	for i := 0; i < m; i++ {
		res += mat[i][i]
		j := m - 1 - i
		if j != i {
			res += mat[i][j]
		}
	}

	return res
}
