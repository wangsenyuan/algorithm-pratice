package p3264

import "sort"

func getFinalState(nums []int, k int, multiplier int) []int {

	n := len(nums)

	id := make([]int, n)
	for i := 0; i < n; i++ {
		id[i] = i
	}

	for i := 0; i < k; i++ {
		sort.Slice(id, func(a, b int) bool {
			return nums[id[a]] < nums[id[b]] || nums[id[a]] == nums[id[b]] && id[a] < id[b]
		})
		nums[id[0]] *= multiplier
	}

	return nums
}
