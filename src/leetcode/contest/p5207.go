package contest

import "sort"

func equalSubstring(s string, t string, maxCost int) int {
	n := len(s)
	cost := make([]int, n+1)
	var best int
	for i := 0; i < n; i++ {
		cost[i+1] = cost[i] + calCost(s[i], t[i])
		j := sort.Search(i+1, func(j int) bool {
			return cost[j]+maxCost >= cost[i+1]
		})
		// cost[j] + maxCost >= cost[i+1]
		// cost[i+1] - cost[j] <= maxCost
		best = max(best, i-j+1)
	}

	return best
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func calCost(a, b byte) int {
	r := int(a) - int(b)
	if r < 0 {
		return -r
	}
	return r
}
