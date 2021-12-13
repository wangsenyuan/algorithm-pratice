package p2098

import "sort"

func maxSubsequence(nums []int, k int) []int {
	n := len(nums)
	pos := make([]int, n)
	for i := 0; i < n; i++ {
		pos[i] = i
	}

	sort.Slice(pos, func(i, j int) bool {
		return nums[pos[i]] > nums[pos[j]]
	})

	pos = pos[:k]

	sort.Ints(pos)

	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = nums[pos[i]]
	}

	return res
}
