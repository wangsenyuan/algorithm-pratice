package a

var C [7][7]int

func init() {
	C[0][0] = 1
	for i := 1; i < 7; i++ {
		C[i][0] = 1
		C[i][i] = 1
		for j := 1; j < i; j++ {
			C[i][j] = C[i-1][j] + C[i-1][j-1]
		}
	}
}

func nCr(n, r int) int {
	if n < r {
		return 0
	}
	return C[n][r]
}

func paintingPlan(n int, k int) int {
	if n*n == k {
		return 1
	}
	// x, y => x * n + y * n - x * y = k
	// (x + y) * n- x * y = k
	var res int

	for x := 0; x < n; x++ {
		for y := 0; y < n; y++ {
			if (x+y)*n-x*y == k {
				res += nCr(n, x) * nCr(n, y)
			}
		}
	}

	return res
}
