package p1444

const MOD = 1000000007

func ways(pizza []string, k int) int {
	r := len(pizza)
	c := len(pizza[0])

	dp := make([][][]int, r)
	for i := 0; i < r; i++ {
		dp[i] = make([][]int, c)
		for j := 0; j < c; j++ {
			dp[i][j] = make([]int, k+1)
			for u := 0; u <= k; u++ {
				dp[i][j][u] = -1
			}
		}
	}

	cnt := make([][]int, r)
	for i := 0; i < r; i++ {
		cnt[i] = make([]int, c)
	}

	for i := r - 1; i >= 0; i-- {
		for j := c - 1; j >= 0; j-- {
			if pizza[i][j] == 'A' {
				cnt[i][j] = 1
			}
			if i+1 < r {
				cnt[i][j] += cnt[i+1][j]
			}
			if j+1 < c {
				cnt[i][j] += cnt[i][j+1]
			}

			if i+1 < r && j+1 < c {
				cnt[i][j] -= cnt[i+1][j+1]
			}
		}
	}

	var cut func(i, j, k int) int

	cut = func(i, j, k int) int {
		if cnt[i][j] == 0 {
			return 0
		}

		if k == 1 {
			// only on piece
			return 1
		}

		if dp[i][j][k] >= 0 {
			return dp[i][j][k]
		}
		dp[i][j][k] = 0

		// try horizontal
		for ii := i + 1; ii < r; ii++ {
			if cnt[i][j] == cnt[ii][j] {
				// no apples between i & ii
				continue
			}
			modAdd(&dp[i][j][k], cut(ii, j, k-1))
		}

		// try vertical

		for jj := j + 1; jj < c; jj++ {
			if cnt[i][j] == cnt[i][jj] {
				continue
			}
			modAdd(&dp[i][j][k], cut(i, jj, k-1))
		}

		return dp[i][j][k]
	}

	return cut(0, 0, k)
}

func modAdd(a *int, b int) {
	*a += b
	if *a >= MOD {
		*a -= MOD
	}
}
