package p2129

func possibleToStamp(grid [][]int, stampHeight int, stampWidth int) bool {
	m := len(grid)
	n := len(grid[0])

	sum := make([][]int, m)
	for i := 0; i < m; i++ {
		sum[i] = make([]int, n)
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

	pref := make([][]int, m)
	for i := 0; i < m; i++ {
		pref[i] = make([]int, n)
	}

	getRangeSum := func(a, b, c, d int) int {
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

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				continue
			}
			// try put stamp here
			if i+stampHeight-1 < m && j+stampWidth-1 < n && getRangeSum(i, j, i+stampHeight-1, j+stampWidth-1) == 0 {
				pref[i][j]++
				if i+stampHeight < m {
					pref[i+stampHeight][j]--
				}
				if j+stampWidth < n {
					pref[i][j+stampWidth]--
				}
				if i+stampHeight < m && j+stampWidth < n {
					pref[i+stampHeight][j+stampWidth]++
				}
			}
		}
	}

	cnt := make([]int, n)

	for i := 0; i < m; i++ {
		var cur int
		for j := 0; j < n; j++ {
			cur += pref[i][j] + cnt[j]
			if grid[i][j] == 0 && cur == 0 {
				return false
			}
			cnt[j] += pref[i][j]
		}
	}

	return true
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
