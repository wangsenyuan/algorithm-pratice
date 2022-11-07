package p2461

func maximumSubarraySum(nums []int, k int) int64 {
	n := len(nums)

	var best int64

	cnt := make(map[int]int)

	var sum int64

	for i := 0; i < n; i++ {
		sum += int64(nums[i])
		cnt[nums[i]]++
		if i >= k-1 {
			if i >= k {
				sum -= int64(nums[i-k])
				cnt[nums[i-k]]--
				if cnt[nums[i-k]] == 0 {
					delete(cnt, nums[i-k])
				}
			}
			if len(cnt) == k {
				best = max(best, sum)
			}
		}
	}

	return best
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}
