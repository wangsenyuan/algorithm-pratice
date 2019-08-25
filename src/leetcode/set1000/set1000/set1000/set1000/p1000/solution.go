package p1000

const INF = 1 << 29

func mergeStones(stones []int, K int) int {
	n := len(stones)

	sum := make([]int, n)

	for i := 0; i < n; i++ {
		sum[i] = stones[i]
		if i > 0 {
			sum[i] += sum[i-1]
		}
	}

	stoneSum := func(i int, j int) int {
		if i == 0 {
			return sum[j]
		}
		return sum[j] - sum[i-1]
	}

	dp := make([][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]int, K+1)
			for k := 0; k <= K; k++ {
				dp[i][j][k] = -1
			}
		}
	}

	var dfs func(i, j, m int) int

	dfs = func(i, j, m int) int {
		if (j-i+1-m)%(K-1) != 0 {
			//can't merge into m piles
			return INF
		}
		if i == j {
			if m == 1 {
				return 0
			}
			return INF
		}
		if dp[i][j][m] >= 0 {
			return dp[i][j][m]
		}
		res := INF
		if m == 1 {
			//merge into K piles first, then into one
			res = dfs(i, j, K) + stoneSum(i, j)
		} else {
			for mid := i; mid < j; mid++ {
				tmp := dfs(i, mid, 1) + dfs(mid+1, j, m-1)
				res = min(res, tmp)
			}
		}
		dp[i][j][m] = res
		return res
	}

	res := dfs(0, n-1, 1)
	if res >= INF {
		return -1
	}
	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
