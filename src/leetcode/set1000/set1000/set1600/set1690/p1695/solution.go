package p1695

func maximumUniqueSubarray(nums []int) int {
	n := len(nums)
	sum := make([]int, n+1)
	pos := make(map[int]int)

	var best int

	var prev = -1
	for i := 0; i < n; i++ {
		sum[i+1] = sum[i] + nums[i]
		if j, found := pos[nums[i]]; found {
			prev = max(prev, j)
		}

		tmp := sum[i+1] - sum[prev+1]

		best = max(best, tmp)

		pos[nums[i]] = i
	}

	return best
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
