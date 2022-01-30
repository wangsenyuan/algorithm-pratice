package p2155

func maxScoreIndices(nums []int) []int {
	n := len(nums)
	sum := make([]int, n+1)
	for i := 0; i < n; i++ {
		sum[i+1] = sum[i] + 1 - nums[i]
	}
	best := sum[n]

	ans := make([]int, n+1)
	ans[n] = sum[n]
	var right int
	for i := n - 1; i >= 0; i-- {
		left := sum[i]
		right += nums[i]
		ans[i] = left + right
		best = max(best, ans[i])
	}

	var res []int

	for i := 0; i <= n; i++ {
		if best == ans[i] {
			res = append(res, i)
		}
	}
	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
