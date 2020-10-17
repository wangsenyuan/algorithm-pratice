package p052

func totalNQueens(n int) int {
	A := make([]bool, n)
	B := make([]bool, n)
	C := make([]bool, n)
	D := make([]bool, n)

	var loop func(row int, flag int) int

	loop = func(row int, flag int) int {
		if row == n {
			return 1
		}

		var res int

		for i := 0; i < n; i++ {
			if flag&(1<<i) > 0 {
				continue
			}
			if row >= i && A[row-i] {
				continue
			}
			if row < i && B[i-row] {
				continue
			}
			if row-(n-1-i) >= 0 && C[row-(n-1-i)] {
				continue
			}
			if (n-1-i)-row >= 0 && D[(n-1-i)-row] {
				continue
			}

			if row >= i {
				A[row-i] = true
			}
			if row <= i {
				B[i-row] = true
			}
			if row-(n-1-i) >= 0 {
				C[row-(n-1-i)] = true
			}
			if (n-1-i)-row >= 0 {
				D[(n-1-i)-row] = true
			}

			res += loop(row+1, flag|(1<<i))

			if row >= i {
				A[row-i] = false
			}
			if row <= i {
				B[i-row] = false
			}
			if row-(n-1-i) >= 0 {
				C[row-(n-1-i)] = false
			}
			if (n-1-i)-row >= 0 {
				D[(n-1-i)-row] = false
			}
		}
		return res
	}

	return loop(0, 0)
}
