package lcp43

const N = 21

const INF = 1 << 30

func trafficCommand(directions []string) int {
	// e, s, w, n
	dp := make([][][][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([][][]int, N)
		for j := 0; j < N; j++ {
			dp[i][j] = make([][]int, N)
			for k := 0; k < N; k++ {
				dp[i][j][k] = make([]int, N)
				for u := 0; u < N; u++ {
					dp[i][j][k][u] = -1
				}
			}
		}
	}

	valid := func(buf []byte, mask int, p []int) bool {
		for i := 0; i < 4; i++ {
			if (mask>>i)&1 == 1 {
				if p[i] >= len(directions[i]) {
					return false
				}
				buf[i] = directions[i][p[i]]
			} else {
				buf[i] = '.'
			}
		}

		if buf[0] == 'W' {
			// go west
			if buf[1] == 'N' || buf[1] == 'W' || buf[2] == 'N' || buf[3] == 'S' || buf[3] == 'E' || buf[3] == 'W' {
				return false
			}
		}
		if buf[0] == 'S' {
			if buf[1] == 'N' || buf[1] == 'W' || buf[2] == 'E' || buf[2] == 'S' || buf[3] == 'E' || buf[3] == 'S' {
				return false
			}
		}

		if buf[0] == 'N' {
			if buf[1] == 'N' || buf[2] == 'N' {
				return false
			}
		}

		if buf[1] == 'N' {
			if buf[2] == 'E' || buf[2] == 'N' || buf[3] == 'E' {
				return false
			}
		}

		if buf[1] == 'W' {
			if buf[2] == 'E' || buf[2] == 'N' || buf[3] == 'S' || buf[3] == 'W' {
				return false
			}
		}
		if buf[1] == 'E' {
			if buf[2] == 'E' || buf[3] == 'E' {
				return false
			}
		}

		if buf[2] == 'E' {
			if buf[3] == 'S' || buf[3] == 'E' {
				return false
			}
		}

		if buf[2] == 'N' {
			if buf[3] == 'S' || buf[3] == 'E' {
				return false
			}
		}

		if buf[2] == 'S' {
			if buf[3] == 'S' {
				return false
			}
		}

		return true
	}

	var fn func(i, j, k, u int) int
	fn = func(i, j, k, u int) int {
		if dp[i][j][k][u] >= 0 {
			return dp[i][j][k][u]
		}
		if i == len(directions[0]) && j == len(directions[1]) && k == len(directions[2]) && u == len(directions[3]) {
			return 0
		}
		best := INF
		for mask := 1; mask < 1<<4; mask++ {
			buf := make([]byte, 4)
			if valid(buf, mask, []int{i, j, k, u}) {
				ii, jj, kk, uu := i, j, k, u
				if buf[0] != '.' {
					ii++
				}
				if buf[1] != '.' {
					jj++
				}
				if buf[2] != '.' {
					kk++
				}
				if buf[3] != '.' {
					uu++
				}
				best = min(best, fn(ii, jj, kk, uu))
			}
		}
		dp[i][j][k][u] = best + 1
		return best + 1
	}

	return fn(0, 0, 0, 0)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
