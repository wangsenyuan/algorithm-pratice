package p2573

func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
	var res [][]int

	for i, j := 0, 0; i < len(nums1) || j < len(nums2); {
		if i < len(nums1) && j < len(nums2) {
			if nums1[i][0] == nums2[j][0] {
				res = append(res, []int{nums1[i][0], nums1[i][1] + nums2[j][1]})
				i++
				j++
				continue
			}
			if nums1[i][0] < nums2[j][0] {
				res = append(res, nums1[i])
				i++
				continue
			}
			res = append(res, nums2[j])
			j++
			continue
		}
		if i < len(nums1) {
			res = append(res, nums1[i])
			i++
		} else {
			res = append(res, nums2[j])
			j++
		}
	}
	return res
}
