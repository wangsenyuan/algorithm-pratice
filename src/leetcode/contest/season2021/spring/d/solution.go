package d

var dd = []int{-1, 0, 1, 0, -1}

func escapeMaze(maze [][]string) bool {
	l := len(maze)
	m0 := maze[0]
	n := len(m0)
	m := len(m0[0])
	dp := make([][][][]bool, l)
	for t := 0; t < l; t++ {
		dp[t] = make([][][]bool, n*m)
		for i := 0; i < n*m; i++ {
			dp[t][i] = make([][]bool, 2)
			for j := 0; j < 2; j++ {
				dp[t][i][j] = make([]bool, 3)
			}
		}
	}

	dp[0][0][0][0] = true

	for t := 0; t < l-1; t++ {
		for p := 0; p < n*m; p++ {
			u, v := p/m, p%m
			for d1 := 0; d1 < 2; d1++ {
				for d2 := 0; d2 < 3; d2++ {
					if dp[t][p][d1][d2] {
						if d2 == 2 || maze[t+1][u][v] == '.' {
							dp[t+1][p][d1][d2] = true
						}
						for i := 0; i < 4; i++ {
							x, y := u+dd[i], v+dd[i+1]
							if x >= 0 && x < n && y >= 0 && y < m {
								if maze[t+1][x][y] == '.' {
									dp[t+1][x*m+y][d1][min(d2, 1)] = true
								} else {
									if d1 == 0 {
										dp[t+1][x*m+y][1][min(d2, 1)] = true
									}
									if d2 == 0 {
										dp[t+1][x*m+y][d1][2] = true
									}
								}
							}
						}
					}
				}
			}
		}
	}

	for d1 := 0; d1 < 2; d1++ {
		for d2 := 0; d2 < 3; d2++ {
			if dp[l-1][n*m-1][d1][d2] {
				return true
			}
		}
	}
	return false
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
