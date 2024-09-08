package p3280

func findMaximumScore(nums []int) int64 {
	n := len(nums)
	// score += (j - i) * nums[i]
	// dp[k] 表示k到n-1的结果
	// 那么假设对于i的时候，把k计入的时候
	// dp[k] + k * num[i]最大的结果
	// 如果 num[i] >= nums[i+1], 似乎使用nums[i+1]肯定不是一个好主意
	// 如果 nums[i] < nums[i+1], 那么使用nums[i+1]肯定不会更差
	var res int
	j := 0
	for i := 1; i < n; i++ {
		res += nums[j]

		if nums[i] > nums[j] {
			j = i
		}
	}

	return int64(res)
}
