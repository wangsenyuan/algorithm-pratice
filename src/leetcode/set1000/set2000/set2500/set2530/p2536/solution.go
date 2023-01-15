package p2536

func rangeAddQueries(n int, queries [][]int) [][]int {
	mat := make([][]int, n)
	for i := 0; i < n; i++ {
		mat[i] = make([]int, n)
	}

	for _, q := range queries {
		r1, c1, r2, c2 := q[0], q[1], q[2], q[3]
		mat[r1][c1]++
		if c2+1 < n {
			mat[r1][c2+1]--
		}
		if r2+1 < n {
			mat[r2+1][c1]--
		}
		if r2+1 < n && c2+1 < n {
			mat[r2+1][c2+1]++
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i > 0 {
				mat[i][j] += mat[i-1][j]
			}
			if j > 0 {
				mat[i][j] += mat[i][j-1]
			}
			if i > 0 && j > 0 {
				mat[i][j] -= mat[i-1][j-1]
			}
		}
	}
	return mat
}
