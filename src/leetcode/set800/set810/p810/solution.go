package p810

func maxSumTwoNoOverlap(A []int, L int, M int) int {
	// M is the larger sub-array
	n := len(A)
	B := make([]int, n+1)
	for i := 0; i < n; i++ {
		B[i+1] = B[i] + A[i]
	}

	E := make([]int, n)
	F := make([]int, n)
	for i := 0; i < n; i++ {
		if i+1 >= L {
			if i > 0 {
				E[i] = max(E[i-1], B[i+1]-B[i-L+1])
			} else {
				E[i] = B[i+1] - B[i-L+1]
			}
		}
		if i+1 >= M {
			if i > 0 {
				F[i] = max(F[i-1], B[i+1]-B[i-M+1])
			} else {
				F[i] = B[i+1] - B[i-M+1]
			}
		}
	}
	// x is the max L sum from right
	var x int
	var y int

	var best int
	for i := n - 1; i >= 0; i-- {
		if n-i < L {
			x += A[i]
		} else {
			x = max(x, B[i+L]-B[i])
			if i >= M {
				best = max(best, x+F[i-1])
			}
		}

		if n-i < M {
			y += A[i]
		} else {
			y = max(y, B[i+M]-B[i])
			if i >= L {
				best = max(best, y+E[i-1])
			}
		}
	}

	return best
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
