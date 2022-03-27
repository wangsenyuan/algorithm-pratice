package p2215

func findDifference(nums1 []int, nums2 []int) [][]int {
	mem1 := toMap(nums1)
	mem2 := toMap(nums2)
	res := make([][]int, 2)
	for i := 0; i < 2; i++ {
		res[i] = make([]int, 0, 1)
	}
	for k := range mem1 {
		if mem2[k] == 0 {
			res[0] = append(res[0], k)
		}
	}
	for k := range mem2 {
		if mem1[k] == 0 {
			res[1] = append(res[1], k)
		}
	}
	return res
}

func toMap(nums []int) map[int]int {
	res := make(map[int]int)
	for _, num := range nums {
		res[num]++
	}
	return res
}
