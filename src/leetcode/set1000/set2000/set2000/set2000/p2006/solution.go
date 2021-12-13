package p2006

import "sort"

func countKDifference(nums []int, k int) int {
	sort.Ints(nums)
	cnt := make(map[int]int)
	var res int

	for i := 0; i < len(nums); i++ {
		res += cnt[nums[i]-k]
		cnt[nums[i]]++
	}

	return res
}
