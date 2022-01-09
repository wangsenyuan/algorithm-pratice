package p2129

func possibleToStamp(grid [][]int, stampHeight int, stampWidth int) bool {
	m := len(grid)
	n := len(grid[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; {
			if grid[i][j] == 1 {
				j++
				continue
			}
			k := j
			for j < n && grid[i][j] == 0 {
				j++
			}
			if j-k < stampWidth {
				return false
			}
		}
	}

	for j := 0; j < n; j++ {
		for i := 0; i < m; {
			if grid[i][j] == 1 {
				i++
				continue
			}
			k := i
			for i < m && grid[i][j] == 0 {
				i++
			}
			if i-k < stampHeight {
				return false
			}
		}
	}

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

	get := func(i, j int) int {
		if i < 0 || i >= m || j < 0 || j >= n {
			return 1
		}
		return grid[i][j]
	}

	// lets check corners

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				continue
			}

			if get(i, j-1) == 1 && get(i-1, j) == 1 {
				tmp := getRangeSum(i, j, i+stampHeight-1, j+stampWidth-1)
				if tmp > 0 {
					return false
				}
			}

			if get(i, j+1) == 1 && get(i-1, j) == 1 {
				tmp := getRangeSum(i, j-stampWidth+1, i+stampHeight-1, j)
				if tmp > 0 {
					return false
				}
			}

			if get(i+1, j) == 1 && get(i, j-1) == 1 {
				tmp := getRangeSum(i-stampHeight+1, j, i, j+stampWidth-1)
				if tmp > 0 {
					return false
				}
			}

			if get(i+1, j) == 1 && get(i, j+1) == 1 {
				tmp := getRangeSum(i-stampHeight+1, j-stampWidth+1, i, j)
				if tmp > 0 {
					return false
				}
			}
		}
	}

	return true
}
