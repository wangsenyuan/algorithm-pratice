package p3240

func minFlips(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])
	var res int

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			// n - 1 - i, j
			tot := 1
			cnt := grid[i][j]
			if n-1-i != i {
				tot++
				cnt += grid[n-1-i][j]
			}

			if m-1-j != j {
				tot++
				cnt += grid[i][m-1-j]
			}

			if n-1-i != i && m-1-j != j {
				tot++
				cnt += grid[n-1-i][m-1-j]
			}
			if tot < 4 {
				continue
			}
			if tot == 4 {
				if cnt == 1 || cnt == 3 {
					res++
				} else if cnt == 2 {
					res += 2
				}
			}
			grid[i][j] = 0
			grid[n-1-i][j] = 0
			grid[i][m-1-j] = 0
			grid[n-1-i][m-1-j] = 0
		}
	}
	var one int
	var diff int
	if n%2 == 1 {
		// 中间那一行
		for i, j := 0, m-1; i < j; i, j = i+1, j-1 {
			if grid[n/2][i] != grid[n/2][j] {
				diff++
			} else if grid[n/2][i] == 1 {
				one++
			}
		}
	}
	if m%2 == 1 {
		for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
			if grid[i][m/2] != grid[j][m/2] {
				diff++
			} else if grid[i][m/2] == 1 {
				one++
			}
		}
	}
	res += diff
	if n%2 == 1 && m%2 == 1 {
		res += grid[n/2][m/2]
	}
	if diff == 0 && one%2 == 1 {
		res += 2
	}

	return res
}
