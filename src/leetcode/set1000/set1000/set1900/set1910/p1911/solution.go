package p1911

func maxAlternatingSum(nums []int) int64 {
	var ans int64
	var buy int64
	var sell int64
	for i := 0; i < len(nums); i++ {
		buy, sell = max(buy, sell+int64(nums[i])), max(sell, buy-int64(nums[i]))
		ans = max(ans, buy)
	}
	return ans
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}
