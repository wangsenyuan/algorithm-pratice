package p3133

func addedInteger(nums1 []int, nums2 []int) int {
	a := min_of(nums1)
	b := min_of(nums2)
	return b - a
}

func min_of(nums []int) int {
	res := nums[0]
	for _, num := range nums {
		res = min(res, num)
	}
	return res
}
