package p1899

import "sort"

func earliestAndLatest(n int, firstPlayer int, secondPlayer int) []int {
	firstPlayer--
	secondPlayer--
	dp := make([][][]int, n+1)
	fp := make([][][]int, n+1)

	for i := 1; i <= n; i++ {
		dp[i] = make([][]int, n)
		fp[i] = make([][]int, n)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = make([]int, n)
			fp[i][j] = make([]int, n)
			for u := 0; u < n; u++ {
				dp[i][j][u] = -1
			}
		}
	}

	var dfs func(n int, first int, second int)
	dfs = func(n int, first int, second int) {
		if dp[n][first][second] > 0 {
			return
		}
		if first == n-1-second {
			// only two
			dp[n][first][second] = 1
			fp[n][first][second] = 1
			return
		}
		dp[n][first][second] = 1 << 30
		h := n / 2
		t := 1 << h
		for mask := 0; mask < t; mask++ {

			next := make([]int, 0, (n+1)/2)
			for i := 0; i < h; i++ {

				if mask&(1<<i) == 0 {
					next = append(next, n-1-i)
				} else {
					next = append(next, i)
				}
			}

			if n&1 == 1 {
				next = append(next, h)
			}
			sort.Ints(next)
			a, b := -1, -1
			for i := 0; i < len(next); i++ {
				if next[i] == first {
					a = i
				}
				if next[i] == second {
					b = i
				}
			}
			if a < 0 || b < 0 {
				continue
			}
			dfs(len(next), a, b)
			dp[n][first][second] = min(dp[n][first][second], dp[len(next)][a][b]+1)
			fp[n][first][second] = max(fp[n][first][second], fp[len(next)][a][b]+1)
		}
	}

	dfs(n, firstPlayer, secondPlayer)

	return []int{dp[n][firstPlayer][secondPlayer], fp[n][firstPlayer][secondPlayer]}
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
