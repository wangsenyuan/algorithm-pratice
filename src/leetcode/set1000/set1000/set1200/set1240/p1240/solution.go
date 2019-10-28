package p1240

func tilingRectangle(n int, m int) int {
	grid := make([][]int, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]int, m)
	}

	findUnfilledTopLeft := func() (int, int) {
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if grid[i][j] == 0 {
					return i, j
				}
			}
		}
		return -1, -1
	}

	allUnFilled := func(x, y, size int) bool {
		for i := x; i < x+size; i++ {
			for j := y; j < y+size; j++ {
				if grid[i][j] == 1 {
					return false
				}
			}
		}
		return true
	}

	fill := func(x, y, size, v int) {
		for i := x; i < x+size; i++ {
			for j := y; j < y+size; j++ {
				grid[i][j] = v
			}
		}
	}

	var loop func(cur, left int)

	best := n * m

	loop = func(cur, left int) {
		best = min(best, cur+left)

		if cur > best || left == 0 {
			return
		}

		x, y := findUnfilledTopLeft()

		for size := min(n-x, m-y); size >= 1; size-- {
			if !allUnFilled(x, y, size) {
				continue
			}
			// can fill size from (r, c)
			fill(x, y, size, 1)

			loop(cur+1, left-size*size)

			fill(x, y, size, 0)
		}
	}

	loop(0, n*m)

	return best
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
