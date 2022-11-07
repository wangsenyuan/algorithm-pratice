package p2463

import "sort"

var INF int64 = 1 << int64(60)

func minimumTotalDistance(robot []int, factory [][]int) int64 {
	sort.Ints(robot)

	sort.Slice(factory, func(i, j int) bool {
		return factory[i][0] < factory[j][0]
	})

	m := len(robot)
	n := len(factory)

	dp := make([]int64, m+1)
	for i := 1; i <= m; i++ {
		dp[i] = INF
	}
	// 前i个工厂修理了前j个机器人
	dp[0] = 0

	for i := 1; i <= n; i++ {
		for j := m; j > 0; j-- {

			var cost int64

			for k := 1; k <= j && k <= factory[i-1][1]; k++ {
				// factory 修理前k个机器人
				cost += int64(abs(robot[j-k] - factory[i-1][0]))
				dp[j] = min(dp[j], dp[j-k]+cost)
			}
		}
	}

	return dp[m]
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
