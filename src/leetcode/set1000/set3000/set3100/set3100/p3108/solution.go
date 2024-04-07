package p3108

func longestMonotonicSubarray(nums []int) int {
	res := maxInc(nums)
	reverse(nums)
	res = max(res, maxInc(nums))
	return res
}

func reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func maxInc(nums []int) int {
	n := len(nums)
	cur := 1
	var best = 1
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			cur++
			best = max(best, cur)
		} else {
			cur = 1
		}
	}
	return best
}
