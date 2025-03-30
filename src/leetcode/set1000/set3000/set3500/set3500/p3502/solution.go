package p3502

func minCosts(cost []int) []int {
	n := len(cost)
	// ans[i] = min(ans[0], ans[1], ... ans[i-1], f(i))
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = cost[i]
		if i > 0 {
			ans[i] = min(ans[i], ans[i-1])
		}
	}
	return ans
}
