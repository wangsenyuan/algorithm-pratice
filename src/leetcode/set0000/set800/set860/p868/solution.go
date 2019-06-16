package p868

func transpose(A [][]int) [][]int {
	n := len(A)
	if n == 0 {
		return nil
	}

	m := len(A[0])
	if m == 0 {
		return nil
	}

	B := make([][]int, m)
	for i := 0; i < m; i++ {
		B[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			B[j][i] = A[i][j]
		}
	}

	return B
}
