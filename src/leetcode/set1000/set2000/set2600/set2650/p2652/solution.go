package p2652

func getSubarrayBeauty(nums []int, k int, x int) []int {
	cnt := make([]int, 101)
	n := len(nums)
	var res []int
	for i := 0; i < n; i++ {
		cnt[nums[i]+50]++
		if i >= k-1 {
			// have k elements
			var c int
			for j := 0; j < 50; j++ {
				c += cnt[j]
				if c >= x {
					res = append(res, j-50)
					c = -1
					break
				}
			}
			if c >= 0 {
				res = append(res, 0)
			}
			cnt[nums[i-k+1]+50]--
		}
	}
	return res
}
