package p1800

func maxAscendingSum(nums []int) int {
	var best int = nums[0]
	var sum int = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			sum += nums[i]
			best = max(best, sum)
			continue
		}
		sum = nums[i]
		best = max(best, sum)
	}
	return best
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
