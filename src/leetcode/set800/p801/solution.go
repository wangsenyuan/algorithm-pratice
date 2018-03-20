package p801

func minSwap(A []int, B []int) int {
	n := len(A)
	if n < 2 {
		return 0
	}
	// noswap
	x := 0
	// swap
	y := 1

	INF := 1 << 30
	for i := 1; i < n; i++ {
		xx, yy := INF, INF
		a, b := A[i], B[i]
		c, d := A[i-1], B[i-1]
		if x < INF {
			// no swap before
			if a > c && b > d {
				// no swap here
				xx = x
			}
			if a > d && b > c {
				// try swap here
				yy = x + 1
			}
		}
		if y < INF {
			// no swap here
			if a > d && b > c && y < xx {
				xx = y
			}
			if a > c && b > d && y+1 < yy {
				yy = y + 1
			}
		}
		x, y = xx, yy
	}

	return min(x, y)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
