package p2956

func findIntersectionValues(nums1 []int, nums2 []int) []int {
	res := []int{0, 0}

	for _, x := range nums1 {
		for _, y := range nums2 {
			if x == y {
				res[0]++
				break
			}
		}
	}

	for _, x := range nums2 {
		for _, y := range nums1 {
			if x == y {
				res[1]++
				break
			}
		}
	}

	return res
}
