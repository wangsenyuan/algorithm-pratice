package p1886

func findRotation(mat [][]int, target [][]int) bool {
	n := len(mat)

	if n == 1 {
		return mat[0][0] == target[0][0]
	}

	back := make([][]int, n)
	for i := 0; i < n; i++ {
		back[i] = make([]int, n)
	}

	check := func() bool {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if mat[i][j] != target[i][j] {
					return false
				}
			}
		}
		return true
	}

	for k := 0; k < 4; k++ {
		rotate(mat, n, back)

		if check() {
			return true
		}
	}

	return false
}

func rotate(mat [][]int, n int, back [][]int) {

	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			back[c][n-1-r] = mat[r][c]
		}
	}

	for r := 0; r < n; r++ {
		copy(mat[r], back[r])
	}
}
