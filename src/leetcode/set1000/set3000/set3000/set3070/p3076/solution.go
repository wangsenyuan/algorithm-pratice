package p3076

func resultArray(nums []int) []int {
	var a1 []int
	var a2 []int
	a1 = append(a1, nums[0])
	a2 = append(a2, nums[1])
	for i := 2; i < len(nums); i++ {
		if a1[len(a1)-1] > a2[len(a2)-1] {
			a1 = append(a1, nums[i])
		} else {
			a2 = append(a2, nums[i])
		}
	}
	return append(a1, a2...)
}
