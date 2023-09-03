package p2841

func maxSum(nums []int, m int, k int) int64 {
	freq := make(map[int]int)
	var res int64
	n := len(nums)
	var sum int64
	for i := 0; i < n; i++ {
		sum += int64(nums[i])
		freq[nums[i]]++
		if i >= k-1 {
			if len(freq) >= m {
				res = max(res, sum)
			}
			freq[nums[i-k+1]]--
			if freq[nums[i-k+1]] == 0 {
				delete(freq, nums[i-k+1])
			}
			sum -= int64(nums[i-k+1])
		}
	}
	return res
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}
