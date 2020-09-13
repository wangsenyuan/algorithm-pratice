package p1582

func numSpecial(mat [][]int) int {
	m := len(mat)
	n := len(mat[0])

	rc := make([]int, m)
	cc := make([]int, n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			rc[i] += mat[i][j]
			cc[j] += mat[i][j]
		}
	}

	var res int

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 1 && rc[i] == 1 && cc[j] == 1 {
				res++
			}
		}
	}

	return res
}
