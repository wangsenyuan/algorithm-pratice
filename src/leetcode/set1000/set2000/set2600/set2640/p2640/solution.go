package p2640

func findPrefixScore(nums []int) []int64 {
	n := len(nums)

	ans := make([]int64, n)

	var x int

	for i := 0; i < n; i++ {
		if nums[i] > x {
			x = nums[i]
		}
		cover := int64(nums[i]) + int64(x)
		ans[i] = cover
		if i > 0 {
			ans[i] += ans[i-1]
		}
	}

	return ans
}
