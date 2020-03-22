package p1389

func createTargetArray(nums []int, index []int) []int {
	res := make([]int, 0, len(nums))

	for i := 0; i < len(nums); i++ {
		res = append(res, nums[i])
		if index[i] < len(res)-1 {
			copy(res[index[i]+1:], res[index[i]:])
			res[index[i]] = nums[i]
		}
	}

	return res
}
