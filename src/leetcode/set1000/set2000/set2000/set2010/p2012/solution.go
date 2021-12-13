package p2012

func sumOfBeauties(nums []int) int {
	n := len(nums)
	right := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		right[i] = nums[i]
		if i+1 < n {
			right[i] = min(right[i], right[i+1])
		}
	}
	var res int
	var left int
	for i := 0; i < n; i++ {
		if i > 0 && i+1 < n {
			if left < nums[i] && nums[i] < right[i+1] {
				res += 2
			} else if nums[i-1] < nums[i] && nums[i] < nums[i+1] {
				res++
			}
		}
		left = max(left, nums[i])
	}
	return res
}
func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
