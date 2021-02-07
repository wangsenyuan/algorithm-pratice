package p1749

func maxAbsoluteSum(nums []int) int {
	a := maxSum(nums)
	for i := 0; i < len(nums); i++ {
		nums[i] = -nums[i]
	}
	b := maxSum(nums)

	return max(a, b)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func maxSum(arr []int) int {
	var best int
	var sum int

	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		if sum > best {
			best = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return best
}
