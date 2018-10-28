package p931

func minFallingPathSum(A [][]int) int {
	n := len(A)
	B := make([][]int, 2)
	B[0] = make([]int, n)
	B[1] = make([]int, n)
	for i := 0; i < n; i++ {
		B[0][i] = A[0][i]
	}
	var p int
	for i := 1; i < n; i++ {
		q := 1 - p
		for j := 0; j < n; j++ {
			B[q][j] = B[p][j] + A[i][j]
			if j > 0 {
				B[q][j] = min(B[q][j], B[p][j-1]+A[i][j])
			}
			if j < n-1 {
				B[q][j] = min(B[q][j], B[p][j+1]+A[i][j])
			}
		}
		p = q
	}

	ans := B[p][0]
	for j := 1; j < n; j++ {
		ans = min(ans, B[p][j])
	}
	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
