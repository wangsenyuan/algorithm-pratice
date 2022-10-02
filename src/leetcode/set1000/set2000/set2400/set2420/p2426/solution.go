package p2426

func xorAllNums(nums1 []int, nums2 []int) int {
	m := len(nums1)
	n := len(nums2)
	var res int
	if n&1 == 1 {
		for _, num := range nums1 {
			res ^= num
		}
	}
	if m&1 == 1 {
		for _, num := range nums2 {
			res ^= num
		}
	}
	return res
}
