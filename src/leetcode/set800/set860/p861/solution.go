package p861

func matrixScore(A [][]int) int {
	R := len(A)
	C := len(A[0])

	var ans int

	for c := 0; c < C; c++ {
		var tmp int
		for r := 0; r < R; r++ {
			tmp += A[r][c] ^ A[r][0]
		}
		ans += max(tmp, R-tmp) * (1 << uint(C-1-c))
	}

	return ans
}

func matrixScore1(A [][]int) int {
	n := len(A)
	m := len(A[0])
	M := 1 << uint(m)
	var best int

	nums := rowToNum(A, n)
	nums2 := rowToNum2(A, n)

	for state := 0; state < M; state++ {
		var tmp int
		for i := 0; i < n; i++ {
			a, b := nums[i], nums2[i]

			a ^= state
			b ^= state
			tmp += max(a, b)
		}
		best = max(best, tmp)
	}

	return best
}

func rowToNum(A [][]int, n int) []int {
	res := make([]int, n)

	for i := 0; i < n; i++ {
		var tmp int
		m := len(A[i])
		for j := 0; j < m; j++ {
			if A[i][j] == 1 {
				tmp |= 1 << uint(m-1-j)
			}
		}
		res[i] = tmp
	}

	return res
}

func rowToNum2(A [][]int, n int) []int {
	res := make([]int, n)

	for i := 0; i < n; i++ {
		var tmp int
		m := len(A[i])
		for j := 0; j < m; j++ {
			if A[i][j] == 0 {
				tmp |= 1 << uint(m-1-j)
			}
		}
		res[i] = tmp
	}

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
