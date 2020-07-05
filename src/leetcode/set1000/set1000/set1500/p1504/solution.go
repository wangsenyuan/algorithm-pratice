package main

func numSubmat(mat [][]int) int {
	m := len(mat)
	n := len(mat[0])

	S := make([][]int, m)
	for i := 0; i < m; i++ {
		S[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			S[i][j] = mat[i][j]
			if i > 0 {
				S[i][j] += S[i-1][j]
			}
			if j > 0 {
				S[i][j] += S[i][j-1]
			}
			if i > 0 && j > 0 {
				S[i][j] -= S[i-1][j-1]
			}
		}
	}
	var res int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for u := i; u < m; u++ {
				for v := j; v < n; v++ {
					X := S[u][v]
					if i > 0 {
						X -= S[i-1][v]
					}
					if j > 0 {
						X -= S[u][j-1]
					}
					if i > 0 && j > 0 {
						X += S[i-1][j-1]
					}
					if X != (u-i+1)*(v-j+1) {
						break
					}
					res++
				}
			}
		}
	}
	return res
}
