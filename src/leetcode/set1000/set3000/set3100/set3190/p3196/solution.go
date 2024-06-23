package p3196

func maximumTotalCost(nums []int) int64 {
	// 子数组的长度是不是都是1/2 ？
	// 所以就可以用两个状态就可以了
	a, b := 0, nums[0]

	for i := 1; i < len(nums); i++ {
		cur := max(b+nums[i], a+nums[i-1]-nums[i])
		a, b = b, cur
	}
	return int64(b)
}
