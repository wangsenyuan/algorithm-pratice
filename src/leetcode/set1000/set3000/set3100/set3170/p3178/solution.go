package p3178

const mod = 1e9 + 7

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func valueAfterKSeconds(n int, k int) int {
	// f(i, j) = f(i - 1, j) + f(i, j - 1)
	f := make([][]int, k+1)
	for i := 0; i <= k; i++ {
		f[i] = make([]int, n)
	}
	for j := 0; j < n; j++ {
		f[0][j] = 1
	}

	for r := 1; r <= k; r++ {
		f[r][0] = 1
		for j := 1; j < n; j++ {
			f[r][j] = add(f[r][j-1], f[r-1][j])
		}
	}
	return f[k][n-1]
}
