package p2465

import "sort"

func distinctAverages(nums []int) int {
	sort.Ints(nums)
	cnt := make(map[int]bool)

	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		cnt[nums[i]+nums[j]] = true
	}

	return len(cnt)
}
