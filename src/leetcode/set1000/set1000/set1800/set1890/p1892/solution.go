package p1892

func largestMagicSquare(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	sum := make([][]int, m)
	for i := 0; i < m; i++ {
		sum[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			sum[i][j] = grid[i][j]
			if i > 0 {
				sum[i][j] += sum[i-1][j]
			}
			if j > 0 {
				sum[i][j] += sum[i][j-1]
			}
			if i > 0 && j > 0 {
				sum[i][j] -= sum[i-1][j-1]
			}
		}
	}

	total := func(a, b, c, d int) int {
		res := sum[c][d]
		if a > 0 {
			res -= sum[a-1][d]
		}
		if b > 0 {
			res -= sum[c][b-1]
		}
		if a > 0 && b > 0 {
			res += sum[a-1][b-1]
		}
		return res
	}

	var best int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			var x int
			for k := 1; i+k <= m && j+k <= n; k++ {
				x += grid[i+k-1][j+k-1]
				// 左框为 j, 右框为 j + k-1 - k
				if k > best && total(i, j, i+k-1, j+k-1) == x*k {
					// a candidate
					ok := true
					var y int
					for l := 0; l < k && ok; l++ {
						y += grid[i+l][j+k-1-l]
						row := total(i+l, j, i+l, j+k-1)
						if row != x {
							ok = false
							break
						}
						col := total(i, j+l, i+k-1, j+l)
						if col != x {
							ok = false
							break
						}
					}
					if ok && x == y {
						best = k
					}
				}
			}
		}
	}
	return best
}
