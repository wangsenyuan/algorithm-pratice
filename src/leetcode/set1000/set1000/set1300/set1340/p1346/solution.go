package p1346

func maxStudents(seats [][]byte) int {
	m := len(seats)
	n := len(seats[0])

	N := 1 << uint(n)
	dp := make([]int, N)
	fp := make([]int, N)

	for state := 1; state < N; state++ {
		var cnt int

		for i := 0; i < n; i++ {
			if state&(1<<uint(i)) > 0 {
				if seats[0][i] == '#' {
					cnt = -1
					break
				}
				if i > 0 && (state&(1<<uint(i-1)) > 0) {
					// invalid position
					cnt = -1
					break
				}
				cnt++
			}
		}
		dp[state] = cnt
		// dp[0] = max(dp[0], cnt)
		fp[state] = -1
	}

	for r := 0; r < m-1; r++ {
		for state := 0; state < N; state++ {
			for next := 0; next < N; next++ {
				if dp[state] >= 0 {
					var cnt int
					for i := 0; i < n; i++ {
						if next&(1<<uint(i)) == 0 {
							continue
						}
						if seats[r+1][i] == '#' {
							cnt = -1
							break
						}
						if i > 0 && (next&(1<<uint(i-1)) > 0) {
							cnt = -1
							break
						}
						if state&(1<<uint(i-1)) > 0 {
							cnt = -1
							break
						}
						if i+1 < n && state&(1<<uint(i+1)) > 0 {
							cnt = -1
							break
						}
						cnt++
					}
					if cnt < 0 {
						fp[next] = max(fp[next], cnt)
					} else {
						fp[next] = max(fp[next], dp[state]+cnt)
					}
				}
			}
		}

		for state := 1; state < N; state++ {
			dp[0] = max(dp[0], dp[state])
			dp[state] = fp[state]
			fp[state] = -1
		}
	}

	res := dp[0]

	for state := 1; state < N; state++ {
		res = max(res, dp[state])
	}

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
