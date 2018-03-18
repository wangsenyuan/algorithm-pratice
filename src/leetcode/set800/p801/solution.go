package p801

func minSwap(A []int, B []int) int {
	n := len(A)
	if n < 2 {
		return 0
	}

	f := make([][]int, n)

	for i := 0; i < n; i++ {
		f[i] = make([]int, 2)
		f[i][0] = 1 << 30
		f[i][1] = 1 << 30
	}
	// noswap
	f[0][0] = 0
	// swap
	f[0][1] = 1

	for i := 1; i < n; i++ {
		a, b := A[i], B[i]
		c, d := A[i-1], B[i-1]
		if f[i-1][0] >= 0 {
			// no swap before
			if a > c && b > d {
				// no swap here
				f[i][0] = f[i-1][0]
			}
			if a > d && b > c {
				// try swap here
				f[i][1] = f[i-1][0] + 1
			}
		}
		if f[i-1][1] > 0 {
			// no swap here
			if a > d && b > c && f[i-1][1] < f[i][0] {
				f[i][0] = f[i-1][1]
			}
			if a > c && b > d && f[i-1][1]+1 < f[i][1] {
				f[i][1] = f[i-1][1] + 1
			}
		}
	}

	return min(f[n-1][0], f[n-1][1])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
