package p1464

func maxProduct(nums []int) int {

	var res int

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			tmp := (nums[i] - 1) * (nums[j] - 1)
			res = max(res, tmp)
		}
	}

	return res

}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
