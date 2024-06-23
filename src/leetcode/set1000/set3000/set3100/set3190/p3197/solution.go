package p3197

const inf = 1 << 30

func minimumSum(grid [][]int) int {
	// bruteforce 第一个矩形的位置（右下角）
	// 然后判断剩余的位置，给另外两个矩形

	// in the range (a, b) (c, d) to cover all 1s
	f := func(a, b, c, d int) int {
		t, l := inf, inf
		x, r := 0, 0
		for i := a; i <= c; i++ {
			for j := b; j <= d; j++ {
				if grid[i][j] == 1 {
					t = min(t, i)
					l = min(l, j)
					x = max(x, i)
					r = max(r, j)
				}
			}
		}
		if t > x {
			return 0
		}
		return (x - t + 1) * (r - l + 1)
	}

	best := inf

	n := len(grid)
	m := len(grid[0])

	for j1 := 0; j1 < m; j1++ {
		for j2 := j1 + 1; j2 < m; j2++ {
			a := f(0, 0, n-1, j1)
			b := f(0, j1+1, n-1, j2)
			c := f(0, j2+1, n-1, m-1)
			best = min(best, a+b+c)
		}
	}

	for i1 := 0; i1 < n; i1++ {
		for i2 := i1 + 1; i2 < n; i2++ {
			a := f(0, 0, i1, m-1)
			b := f(i1+1, 0, i2, m-1)
			c := f(i2+1, 0, n-1, m-1)
			best = min(best, a+b+c)
		}
	}

	for i := 0; i < n; i++ {
		a1 := f(0, 0, i, m-1)
		a2 := f(i+1, 0, n-1, m-1)
		for j := 0; j < m; j++ {
			b1 := f(i+1, 0, n-1, j)
			c1 := f(i+1, j+1, n-1, m-1)
			best = min(best, a1+b1+c1)

			b2 := f(0, 0, i, j)
			c2 := f(0, j+1, i, m-1)
			best = min(best, a2+b2+c2)
		}
	}

	for j := 0; j < m; j++ {
		a1 := f(0, 0, n-1, j)
		a2 := f(0, j+1, n-1, m-1)
		for i := 0; i < n; i++ {
			b1 := f(0, j+1, i, m-1)
			c1 := f(i+1, j+1, n-1, m-1)
			best = min(best, a1+b1+c1)
			b2 := f(0, 0, i, j)
			c2 := f(i+1, 0, n-1, j)
			best = min(best, a2+b2+c2)
		}
	}

	return best
}
