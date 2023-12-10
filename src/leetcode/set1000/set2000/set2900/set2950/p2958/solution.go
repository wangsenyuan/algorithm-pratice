package p2958

func maxSubarrayLength(nums []int, k int) int {
	freq := make(map[int]int)

	var best int
	for l, r := 0, 0; r < len(nums); r++ {
		freq[nums[r]]++
		for freq[nums[r]] > k {
			freq[nums[l]]--
			l++
		}
		best = max(best, r-l+1)
	}

	return best
}
