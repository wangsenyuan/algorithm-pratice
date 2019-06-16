package p746

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return cost[0]
	}

	a, b := cost[0], cost[1]
	for i := 2; i < n; i++ {
		c := cost[i] + min(a, b)
		a, b = b, c
	}
	return min(a, b)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
