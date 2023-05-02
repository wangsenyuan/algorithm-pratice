package p2661

func firstCompleteIndex(arr []int, mat [][]int) int {
	n := len(mat)
	m := len(mat[0])
	row := make([]int, n)
	col := make([]int, m)
	pos := make([]int, m*n+1)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			pos[mat[i][j]] = i*m + j
		}
	}

	for k, x := range arr {
		i, j := pos[x]/m, pos[x]%m
		row[i]++
		col[j]++
		if row[i] == m || col[j] == n {
			return k
		}
	}
	return 0
}
