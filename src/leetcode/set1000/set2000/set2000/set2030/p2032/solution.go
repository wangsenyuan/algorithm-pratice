package p2032

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	occ1 := occur(nums1)
	occ2 := occur(nums2)
	occ3 := occur(nums3)

	var res []int

	for k := range occ1 {
		if occ2[k] || occ3[k] {
			res = append(res, k)
		}
	}
	for k := range occ2 {
		if occ3[k] && !occ1[k] {
			res = append(res, k)
		}
	}
	return res
}

func occur(arr []int) map[int]bool {
	res := make(map[int]bool)
	for _, num := range arr {
		res[num] = true
	}
	return res
}
