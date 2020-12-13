package p5245

import "sort"

func maxHeight(cuboids [][]int) int {
	n := len(cuboids)
	for i := 0; i < n; i++ {
		sort.Ints(cuboids[i])
	}

	sort.Slice(cuboids, func(i, j int) bool {
		a, b := cuboids[i], cuboids[j]

		return a[2] < b[2] || a[2] == b[2] && a[0]+a[1] < b[0]+b[1]
	})

	dp := make([]int, n)

	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			a, b := cuboids[j], cuboids[i]
			if a[0] <= b[0] && a[1] <= b[1] && a[2] <= b[2] {
				dp[i] = max(dp[i], dp[j])
			}
		}
		dp[i] += cuboids[i][2]
	}

	var best int

	for i := 0; i < n; i++ {
		best = max(best, dp[i])
	}
	return best
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
